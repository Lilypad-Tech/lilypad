package solver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	corehttp "net/http"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/data"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/http"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/metricsDashboard"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/solver/stats"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/solver/store"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/system"
	"github.com/go-chi/httprate"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel/sdk/trace"
)

type solverServer struct {
	options       http.ServerOptions
	controller    *SolverController
	store         store.SolverStore
	stats         stats.Stats
	services      data.ServiceConfig
	versionConfig *system.VersionConfig
	log           *zerolog.Logger
}

func NewSolverServer(
	options http.ServerOptions,
	controller *SolverController,
	store store.SolverStore,
	stats stats.Stats,
	versionConfig *system.VersionConfig,
	services data.ServiceConfig,
) (*solverServer, error) {
	server := &solverServer{
		options:       options,
		controller:    controller,
		store:         store,
		stats:         stats,
		versionConfig: versionConfig,
		log:           system.GetLogger(system.SolverService),
	}

	metricsDashboard.Init(services.APIHost)

	return server, nil
}

/*
*
*
*

# Routes

*
*
*
*/
func (server *solverServer) ListenAndServe(ctx context.Context, cm *system.CleanupManager, tracerProvider *trace.TracerProvider) error {
	router := mux.NewRouter()

	subrouter := router.PathPrefix(http.API_SUB_PATH).Subrouter()

	subrouter.Use(http.CorsMiddleware)
	subrouter.Use(otelmux.Middleware("solver", otelmux.WithTracerProvider(tracerProvider)))

	subrouter.Use(httprate.Limit(
		server.options.RateLimiter.RequestLimit,
		time.Duration(server.options.RateLimiter.WindowLength)*time.Second,
		httprate.WithKeyFuncs(httprate.KeyByRealIP, httprate.KeyByEndpoint),
	))

	subrouter.HandleFunc("/job_offers", http.GetHandler(server.getJobOffers)).Methods("GET")
	subrouter.HandleFunc("/job_offers", http.PostHandler(server.addJobOffer)).Methods("POST")
	subrouter.HandleFunc("/job_offers/with_files", server.addJobOfferWithFiles).Methods("POST")

	subrouter.HandleFunc("/job_offers/{id}", http.GetHandler(server.getJobOffer)).Methods("GET")
	subrouter.HandleFunc("/job_offers/{id}/files", http.GetHandler(server.jobOfferDownloadFiles)).Methods("GET")

	subrouter.HandleFunc("/resource_offers", http.GetHandler(server.getResourceOffers)).Methods("GET")
	subrouter.HandleFunc("/resource_offers", http.PostHandler(server.addResourceOffer)).Methods("POST")

	subrouter.HandleFunc("/deals", http.GetHandler(server.getDeals)).Methods("GET")
	subrouter.HandleFunc("/deals/{id}", http.GetHandler(server.getDeal)).Methods("GET")

	subrouter.HandleFunc("/deals/{id}/input_files", http.GetHandler(server.downloadInputFiles)).Methods("GET")
	subrouter.HandleFunc("/deals/{id}/files", http.GetHandler(server.downloadFiles)).Methods("GET")
	subrouter.HandleFunc("/deals/{id}/files", server.uploadFiles).Methods("POST")

	subrouter.HandleFunc("/deals/{id}/result", http.GetHandler(server.getResult)).Methods("GET")
	subrouter.HandleFunc("/deals/{id}/result", http.PostHandler(server.addResult)).Methods("POST")

	subrouter.HandleFunc("/deals/{id}/txs/resource_provider", http.PostHandler(server.updateTransactionsResourceProvider)).Methods("POST")
	subrouter.HandleFunc("/deals/{id}/txs/job_creator", http.PostHandler(server.updateTransactionsJobCreator)).Methods("POST")
	subrouter.HandleFunc("/deals/{id}/txs/mediator", http.PostHandler(server.updateTransactionsMediator)).Methods("POST")

	subrouter.HandleFunc("/validation_token", http.GetHandler(server.getValidationToken)).Methods("GET")

	//anura subrouter
	anuraMiddleware := func(next corehttp.Handler) corehttp.Handler {
		return corehttp.HandlerFunc(func(w corehttp.ResponseWriter, r *corehttp.Request) {
			_, err := http.CheckAnuraSignature(r, server.options.AccessControl.AnuraAddresses)
			if err != nil {
				corehttp.Error(w, "Unauthorized", corehttp.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	anurarouter := router.PathPrefix(http.API_SUB_PATH + "/anura").Subrouter()
	anurarouter.Use(http.CorsMiddleware)
	anurarouter.Use(otelmux.Middleware("solver", otelmux.WithTracerProvider(tracerProvider)))
	anurarouter.Use(anuraMiddleware)

	anurarouter.HandleFunc("/job_offers", http.PostHandler(server.addJobOffer)).Methods("POST")
	anurarouter.HandleFunc("/job_offers/with_files", server.addJobOfferWithFiles).Methods("POST")
	anurarouter.HandleFunc("/job_offers", http.GetHandler(server.getJobOffers)).Methods("GET")
	anurarouter.HandleFunc("/job_offers/{id}", http.GetHandler(server.getJobOffer)).Methods("GET")
	anurarouter.HandleFunc("/job_offers/{id}/files", http.GetHandler(server.jobOfferDownloadFiles)).Methods("GET")

	// this will fan out to all connected web socket connections
	// we read all events coming from inside the solver controller
	// and write them to anyone who is connected to us
	websocketEventChannel := make(chan []byte)

	server.log.Debug().Msgf("begin server.controller.subscribeEvents")
	server.controller.subscribeEvents(func(ev SolverEvent) {
		evBytes, err := json.Marshal(ev)
		if err != nil {
			server.log.Error().Err(err).Msg("error marshalling event")
		}
		websocketEventChannel <- evBytes
	})

	http.StartWebSocketServer(
		subrouter,
		http.WEBSOCKET_SUB_PATH,
		websocketEventChannel,
		ctx,
		server.connectCB,
		server.disconnectCB,
	)

	srv := &corehttp.Server{
		Addr:              fmt.Sprintf("%s:%d", server.options.Host, server.options.Port),
		WriteTimeout:      time.Minute * 15,
		ReadTimeout:       time.Minute * 15,
		ReadHeaderTimeout: time.Minute * 15,
		IdleTimeout:       time.Minute * 60,
		Handler:           router,
	}

	// Create a channel to receive errors from ListenAndServe
	serverErrors := make(chan error, 1)

	// Run ListenAndServe in a goroutine because it blocks
	go func() {
		serverErrors <- srv.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		return err
	case <-ctx.Done():
		// Create a context with a timeout for the server to close
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Attempt to gracefully shut down the server
		if err := srv.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("failed to stop server: %w", err)
		}
	}

	return nil
}

// WS connect events
func (server *solverServer) connectCB(connParams http.WSConnectionParams) {
	if connParams.Type == "ResourceProvider" {
		metricsDashboard.TrackNodeConnectionEvent(metricsDashboard.NodeConnectionParams{
			Event:       "Connect",
			ID:          connParams.ID,
			CountryCode: connParams.CountryCode,
			IP:          connParams.IP,
		})
	}
}

func (server *solverServer) disconnectCB(connParams http.WSConnectionParams) {
	if connParams.Type == "ResourceProvider" {
		metricsDashboard.TrackNodeConnectionEvent(metricsDashboard.NodeConnectionParams{
			Event:       "Disconnect",
			ID:          connParams.ID,
			CountryCode: connParams.CountryCode,
			IP:          connParams.IP,
		})
		server.controller.removeUnmatchedResourceOffers(connParams.ID)
	}
}

/*
*
*
*

# Lists

*
*
*
*/
func (server *solverServer) getJobOffers(res corehttp.ResponseWriter, req *corehttp.Request) ([]data.JobOfferContainer, error) {
	query := store.GetJobOffersQuery{}
	// if there is a job_creator query param then assign it
	if jobCreator := req.URL.Query().Get("job_creator"); jobCreator != "" {
		query.JobCreator = jobCreator
	}
	if notMatched := req.URL.Query().Get("not_matched"); notMatched == "true" {
		query.NotMatched = true
	}
	if active := req.URL.Query().Get("active"); active == "true" {
		query.Active = true
	}
	if cancelled := req.URL.Query().Get("cancelled"); cancelled != "" {
		if val, err := strconv.ParseBool(cancelled); err == nil {
			query.Cancelled = &val
		} else {
			return nil, fmt.Errorf("invalid cancelled filter value: %s", cancelled)
		}
	}

	return server.store.GetJobOffers(query)
}

func (server *solverServer) getResourceOffers(res corehttp.ResponseWriter, req *corehttp.Request) ([]data.ResourceOfferContainer, error) {
	query := store.GetResourceOffersQuery{}
	// if there is a job_creator query param then assign it
	if resourceProvider := req.URL.Query().Get("resource_provider"); resourceProvider != "" {
		query.ResourceProvider = resourceProvider
	}
	if active := req.URL.Query().Get("active"); active == "true" {
		query.Active = true
	}
	if notMatched := req.URL.Query().Get("not_matched"); notMatched == "true" {
		query.NotMatched = true
	}
	return server.store.GetResourceOffers(query)
}

func (server *solverServer) getDeals(res corehttp.ResponseWriter, req *corehttp.Request) ([]data.DealContainer, error) {
	query := store.GetDealsQuery{}
	// if there is a job_creator query param then assign it
	if jobCreator := req.URL.Query().Get("job_creator"); jobCreator != "" {
		query.JobCreator = jobCreator
	}
	if resourceProvider := req.URL.Query().Get("resource_provider"); resourceProvider != "" {
		query.ResourceProvider = resourceProvider
	}
	if state := req.URL.Query().Get("state"); state != "" {
		query.State = state
	}
	return server.store.GetDeals(query)
}

/*
*
*
*

	Getters

*
*
*
*/

func (server *solverServer) getJobOffer(res corehttp.ResponseWriter, req *corehttp.Request) (data.JobOfferContainer, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	jobOffer, err := server.store.GetJobOffer(id)
	if err != nil {
		return data.JobOfferContainer{}, err
	}
	return *jobOffer, nil
}

func (server *solverServer) getDeal(res corehttp.ResponseWriter, req *corehttp.Request) (data.DealContainer, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	deal, err := server.store.GetDeal(id)
	if err != nil {
		return data.DealContainer{}, err
	}
	if deal == nil {
		return data.DealContainer{}, fmt.Errorf("deal not found")
	}
	return *deal, nil
}

func (server *solverServer) getResult(res corehttp.ResponseWriter, req *corehttp.Request) (data.Result, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	result, err := server.store.GetResult(id)
	if err != nil {
		return data.Result{}, err
	}
	if result == nil {
		return data.Result{}, fmt.Errorf("result not found")
	}
	return *result, nil
}

/*
*
*
*

	Adders

*
*
*
*/
func (server *solverServer) addJobOffer(jobOffer data.JobOffer, res corehttp.ResponseWriter, req *corehttp.Request) (*data.JobOfferContainer, error) {
	signerAddress, err := http.CheckSignature(req)
	if err != nil {
		server.log.Error().Err(err).Msgf("error checking signature")
		return nil, err
	}

	// Only the job creator can post their job offer
	if signerAddress != jobOffer.JobCreator {
		return nil, fmt.Errorf("job creator address does not match signer address")
	}

	if server.options.AccessControl.EnableVersionCheck && !http.IsAnura(req, server.options.AccessControl.AnuraAddresses) {
		versionHeader, _ := http.GetVersionFromHeaders(req)
		minVersion, ok := server.versionConfig.IsSupported(versionHeader)
		if !ok {
			server.log.Debug().Str("cid", jobOffer.ID).
				Str("address", jobOffer.JobCreator).
				Str("version", versionHeader).
				Str("minVersion", minVersion).
				Msg("job offer rejected because job creator is running an unsupported version")
			return nil, fmt.Errorf("Please update to minimum supported version %s or newer: https://github.com/Lilypad-Tech/lilypad/releases", minVersion)
		}
	}

	offerRecent := isTimestampRecent(jobOffer.CreatedAt, server.options.AccessControl.OfferTimestampDiffSeconds*1000)
	if !offerRecent {
		server.log.Debug().Str("cid", jobOffer.ID).Str("address", jobOffer.JobCreator).Msg("job offer rejected because timestamp was not recent")
		return nil, errors.New("job offer rejected because CreatedAt time is not recent, check your computer's time settings and network connection")
	}

	err = data.CheckJobOffer(jobOffer)
	if err != nil {
		server.log.Error().Err(err).Msg("Error checking job offer")
		return nil, err
	}

	return server.controller.addJobOffer(jobOffer)
}

func (server *solverServer) addJobOfferWithFiles(res corehttp.ResponseWriter, req *corehttp.Request) {
	// Check signature and validate the job offer
	signerAddress, err := http.CheckSignature(req)
	if err != nil {
		server.log.Error().Err(err).Msg("error checking signature")
		corehttp.Error(res, "Unauthorized", corehttp.StatusUnauthorized)
		return
	}

	// Set a maximum size for the entire request body
	// This includes the job offer, which only accounts for a small portion of the request body
	maxRequestSize := int64(server.options.Storage.MaximumFileInputsSizeMB << 20)
	req.Body = corehttp.MaxBytesReader(res, req.Body, maxRequestSize)

	// Parse the multipart form with max memory size
	err = req.ParseMultipartForm(int64(server.options.Storage.MaximumFileInputsMemoryMB << 20))
	if err != nil {
		server.log.Error().Err(err).Msg("error parsing multipart form")
		corehttp.Error(res, "Error parsing form data", corehttp.StatusBadRequest)
		return
	}

	// Get the job offer JSON part
	jobOfferPart, _, err := req.FormFile("job_offer.json")
	if err != nil {
		server.log.Error().Err(err).Msg("error getting job offer")
		corehttp.Error(res, "Missing job_offer.json part", corehttp.StatusBadRequest)
		return
	}
	defer jobOfferPart.Close()

	var jobOffer data.JobOffer
	err = json.NewDecoder(jobOfferPart).Decode(&jobOffer)
	if err != nil {
		server.log.Error().Err(err).Msg("error decoding job offer JSON")
		corehttp.Error(res, "Invalid job offer JSON", corehttp.StatusBadRequest)
		return
	}

	// Only the job creator can post their job offer
	if signerAddress != jobOffer.JobCreator {
		server.log.Error().
			Str("signer", signerAddress).
			Str("jobCreator", jobOffer.JobCreator).
			Msg("job creator address does not match signer address")
		corehttp.Error(res, "Job creator address does not match signer address", corehttp.StatusUnauthorized)
		return
	}

	// Check version
	if server.options.AccessControl.EnableVersionCheck && !http.IsAnura(req, server.options.AccessControl.AnuraAddresses) {
		versionHeader, _ := http.GetVersionFromHeaders(req)
		minVersion, ok := server.versionConfig.IsSupported(versionHeader)
		if !ok {
			server.log.Debug().Str("cid", jobOffer.ID).
				Str("address", jobOffer.JobCreator).
				Str("version", versionHeader).
				Str("minVersion", minVersion).
				Msg("job offer rejected because job creator is running an unsupported version")
			corehttp.Error(res,
				fmt.Sprintf("Please update to minimum supported version %s or newer: https://github.com/Lilypad-Tech/lilypad/releases", minVersion),
				corehttp.StatusForbidden)
			return
		}
	}

	// Check recency
	offerRecent := isTimestampRecent(jobOffer.CreatedAt, server.options.AccessControl.OfferTimestampDiffSeconds*1000)
	if !offerRecent {
		server.log.Debug().Str("cid", jobOffer.ID).Str("address", jobOffer.JobCreator).Msg("job offer rejected because timestamp was not recent")
		corehttp.Error(res, "Job offer rejected because CreatedAt time is not recent, check your computer's time settings and network connection", corehttp.StatusBadRequest)
		return
	}

	err = data.CheckJobOffer(jobOffer)
	if err != nil {
		server.log.Error().Err(err).Msg("Error checking job offer")
		corehttp.Error(res, fmt.Sprintf("Invalid job offer: %s", err.Error()), corehttp.StatusBadRequest)
		return
	}

	// Get the inputs file part
	inputsFile, _, err := req.FormFile("inputs.tar")
	if err != nil {
		server.log.Error().Err(err).Msg("error getting inputs.tar part")
		corehttp.Error(res, "Missing inputs.tar part", corehttp.StatusBadRequest)
		return
	}
	defer inputsFile.Close()

	// Create directory for job offer inputs file
	jobID, err := data.GetJobOfferID(jobOffer)
	if err != nil {
		server.log.Error().Err(err).Msg("unable to compute job ID")
		corehttp.Error(res, "Could not compute a job ID", corehttp.StatusInternalServerError)
		return
	}
	dirPath, err := EnsureInputsArchivePath(jobID)
	if err != nil {
		server.log.Error().Err(err).Str("cid", jobID).Msg("error creating job offer directory")
		corehttp.Error(res, "Error creating job offer directory", corehttp.StatusInternalServerError)
		return
	}

	// Create the file
	filename := "inputs.tar"
	filePath := filepath.Join(dirPath, filename)
	f, err := os.Create(filePath)
	if err != nil {
		server.log.Error().Err(err).Str("filePath", filePath).Msg("error creating file")
		corehttp.Error(res, "Error saving input file", corehttp.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Copy the data to the file
	_, err = io.Copy(f, inputsFile)
	if err != nil {
		server.log.Error().Err(err).Msg("error copying input file data")
		corehttp.Error(res, "Error saving input file data", corehttp.StatusInternalServerError)
		return
	}

	// Add the job for matching
	jobOfferContainer, err := server.controller.addJobOffer(jobOffer)
	if err != nil {
		server.log.Error().Err(err).Msg("error adding job offer")
		corehttp.Error(res, fmt.Sprintf("Error adding job offer: %s", err.Error()), corehttp.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(res).Encode(jobOfferContainer)
	if err != nil {
		server.log.Error().Err(err).Msg("error encoding job offer container response")
		corehttp.Error(res, "Error encoding response", corehttp.StatusInternalServerError)
		return
	}
}

func (server *solverServer) addResourceOffer(resourceOffer data.ResourceOffer, res corehttp.ResponseWriter, req *corehttp.Request) (*data.ResourceOfferContainer, error) {
	signerAddress, err := http.CheckSignature(req)
	if err != nil {
		server.log.Error().Err(err).Msg("error checking signature")
		return nil, err
	}
	// Only the resource provider can post their resource offer
	if signerAddress != resourceOffer.ResourceProvider {
		return nil, fmt.Errorf("resource provider address does not match signer address")
	}

	// Resource provider must be in allowlist when enabled
	if server.options.AccessControl.EnableResourceProviderAllowlist {
		allowedProviders, err := server.store.GetAllowedResourceProviders()
		if err != nil {
			server.log.Error().Err(err).Msgf("Unable to load resource provider allowlist")
			return nil, err
		}

		if !slices.Contains(allowedProviders, resourceOffer.ResourceProvider) {
			server.log.Debug().Str("address", resourceOffer.ResourceProvider).Msg("resource provider not in allowlist")
			return nil, errors.New("resource provider not in beta program, request beta program access here: https://forms.gle/XaE3rRuXVLxTnZto7")
		}
	}

	if server.options.AccessControl.EnableVersionCheck {
		versionHeader, _ := http.GetVersionFromHeaders(req)
		minVersion, ok := server.versionConfig.IsSupported(versionHeader)
		if !ok {
			server.log.Debug().Str("cid", resourceOffer.ID).
				Str("address", resourceOffer.ResourceProvider).
				Str("version", versionHeader).
				Str("minVersion", minVersion).
				Msg("resource offer rejected because resource provider is running an unsupported version")
			return nil, fmt.Errorf("Please update to minimum supported version %s or newer: https://github.com/Lilypad-Tech/lilypad/releases", minVersion)
		}
	}

	offerRecent := isTimestampRecent(resourceOffer.CreatedAt, server.options.AccessControl.OfferTimestampDiffSeconds*1000)
	if !offerRecent {
		server.log.Debug().Str("cid", resourceOffer.ID).Str("address", resourceOffer.ResourceProvider).Msg("resource offer rejected because timestamp was not recent")
		return nil, errors.New("resource offer rejected because CreatedAt time is not recent, check your computer's time settings and network connection")
	}

	err = data.CheckResourceOffer(resourceOffer)
	if err != nil {
		server.log.Error().Err(err).Msg("Error checking resource offer")
		return nil, err
	}
	return server.controller.addResourceOffer(resourceOffer)
}

func (server *solverServer) addResult(results data.Result, res corehttp.ResponseWriter, req *corehttp.Request) (*data.Result, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	deal, err := server.store.GetDeal(id)
	if err != nil {
		server.log.Error().Err(err).Msg("error loading deal")
		return nil, err
	}
	if deal == nil {
		return nil, fmt.Errorf("deal not found")
	}
	if deal.State == data.GetAgreementStateIndex("JobTimedOut") {
		server.log.Trace().
			Str("cid", deal.ID).
			Str("address", deal.ResourceProvider).
			Msg("attempted results post for timed out job for deal")
		return nil, fmt.Errorf("job with deal ID %s timed out", deal.ID)
	}
	signerAddress, err := http.CheckSignature(req)
	if err != nil {
		server.log.Error().Err(err).Msg("error checking signature")
		return nil, err
	}
	// Only the resource provider in a deal can add a result
	if signerAddress != deal.ResourceProvider {
		return nil, fmt.Errorf("resource provider address does not match signer address")
	}
	err = data.CheckResult(results)
	if err != nil {
		server.log.Error().Err(err).Str("cid", results.DealID).Msgf("Error checking result for deal")
		return nil, err
	}
	results.DealID = id

	storedResult, err := server.store.AddResult(results)
	if err != nil {
		return nil, err
	}

	err = server.updateJobStates(id, "ResultsSubmitted")
	if err != nil {
		return nil, err
	}

	return storedResult, nil
}

/*
*
*
*

	Updaters

*
*
*
*/
func (server *solverServer) updateTransactionsResourceProvider(payload data.DealTransactionsResourceProvider, res corehttp.ResponseWriter, req *corehttp.Request) (*data.DealContainer, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	deal, err := server.store.GetDeal(id)
	if err != nil {
		server.log.Error().Err(err).Str("cid", id).Msg("error loading deal")
		return nil, err
	}
	if deal == nil {
		server.log.Error().Err(err).Str("cid", id).Msg("deal not found")
		return nil, fmt.Errorf("deal not found")
	}
	signerAddress, err := http.CheckSignature(req)
	if err != nil {
		server.log.Error().Err(err).Msg("error checking signature")
		return nil, err
	}
	// Only the resource provider in a deal can update its transactions
	if signerAddress != deal.ResourceProvider {
		return nil, fmt.Errorf("resource provider address does not match signer address")
	}
	return server.controller.updateDealTransactionsResourceProvider(id, payload)
}

func (server *solverServer) updateTransactionsJobCreator(payload data.DealTransactionsJobCreator, res corehttp.ResponseWriter, req *corehttp.Request) (*data.DealContainer, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	deal, err := server.store.GetDeal(id)
	if err != nil {
		server.log.Error().Err(err).Str("cid", id).Msg("error loading deal")
		return nil, err
	}
	if deal == nil {
		server.log.Error().Err(err).Str("cid", id).Msg("deal not found")
		return nil, fmt.Errorf("deal not found")
	}
	signerAddress, err := http.CheckSignature(req)
	if err != nil {
		server.log.Error().Err(err).Msg("error checking signature")
		return nil, err
	}
	// Only the job creator in a deal can update its transactions
	if signerAddress != deal.JobCreator {
		return nil, fmt.Errorf("job creator address does not match signer address")
	}
	return server.controller.updateDealTransactionsJobCreator(id, payload)
}

func (server *solverServer) updateTransactionsMediator(payload data.DealTransactionsMediator, res corehttp.ResponseWriter, req *corehttp.Request) (*data.DealContainer, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	deal, err := server.store.GetDeal(id)
	if err != nil {
		server.log.Error().Err(err).Str("cid", id).Msg("error loading deal")
		return nil, err
	}
	if deal == nil {
		server.log.Error().Err(err).Str("cid", id).Msg("deal not found")
		return nil, fmt.Errorf("deal not found")
	}
	signerAddress, err := http.CheckSignature(req)
	if err != nil {
		server.log.Error().Err(err).Msg("error checking signature")
		return nil, err
	}
	// Only the mediator in a deal can update its transactions
	if signerAddress != deal.Mediator {
		return nil, fmt.Errorf("job creator address does not match mediator address")
	}
	return server.controller.updateDealTransactionsMediator(id, payload)
}

/*
*
*
*

	Files

*
*
*
*/

// We use EmptyResponse to provide a type for the http.GetHandler wrapper,
// but the client expects a file stream and will ignore it.
type EmptyResponse struct{}

func (server *solverServer) downloadInputFiles(res corehttp.ResponseWriter, req *corehttp.Request) (EmptyResponse, error) {
	vars := mux.Vars(req)
	id := vars["id"]

	deal, err := server.store.GetDeal(id)
	if err != nil {
		server.log.Error().Err(err).Str("cid", id).Msg("error loading deal")
		return EmptyResponse{}, &http.HTTPError{
			Message:    "error loading deal",
			StatusCode: corehttp.StatusInternalServerError,
		}
	}
	if deal == nil {
		return EmptyResponse{}, &http.HTTPError{
			Message:    "deal not found",
			StatusCode: corehttp.StatusNotFound,
		}
	}

	// Check authorization
	signerAddress, err := http.CheckSignature(req)
	if err != nil {
		server.log.Error().Err(err).Msg("error checking signature")
		return EmptyResponse{}, &http.HTTPError{
			Message:    "unauthorized",
			StatusCode: corehttp.StatusUnauthorized,
		}
	}
	if signerAddress != deal.ResourceProvider {
		server.log.Debug().
			Str("signer", signerAddress).
			Str("resourceProvider", deal.ResourceProvider).
			Msg("signer address does not match resource provider address")
		return EmptyResponse{}, &http.HTTPError{
			Message:    "not authorized: only the resource provider running this job can download input files",
			StatusCode: corehttp.StatusUnauthorized,
		}
	}

	// Serve the input files
	jobID := deal.Deal.JobOffer.ID
	dirPath, err := EnsureInputsArchivePath(jobID)
	if err != nil {
		server.log.Error().Err(err).Str("cid", jobID).Msg("error getting file inputs directory")
		return EmptyResponse{}, &http.HTTPError{
			Message:    "error accessing input files",
			StatusCode: corehttp.StatusInternalServerError,
		}
	}
	if err := server.handleFileDownload(dirPath, res, func() {
		server.log.Debug().Str("cid", jobID).Str("address", deal.ResourceProvider).Msg("job input files downloaded successfully")
	}); err != nil {
		return EmptyResponse{}, err
	}

	return EmptyResponse{}, nil
}

func (server *solverServer) downloadFiles(res corehttp.ResponseWriter, req *corehttp.Request) (EmptyResponse, error) {
	vars := mux.Vars(req)
	id := vars["id"]

	// Get the deal
	deal, err := server.store.GetDeal(id)
	if err != nil {
		return EmptyResponse{}, &http.HTTPError{
			Message:    "error loading deal",
			StatusCode: corehttp.StatusInternalServerError,
		}
	}
	if deal == nil {
		return EmptyResponse{}, &http.HTTPError{
			Message:    "deal not found",
			StatusCode: corehttp.StatusNotFound,
		}
	}

	// Check authorization
	signerAddress, err := http.CheckSignature(req)
	if err != nil {
		return EmptyResponse{}, &http.HTTPError{
			Message:    "not authorized",
			StatusCode: corehttp.StatusUnauthorized,
		}
	}
	if signerAddress != deal.JobCreator {
		server.log.Debug().
			Str("signer", signerAddress).
			Str("address", deal.JobCreator).
			Msg("signer address does not match job creator address")
		return EmptyResponse{}, &http.HTTPError{
			Message:    "not authorized: job creator address does not match signer address",
			StatusCode: corehttp.StatusUnauthorized,
		}
	}

	if err := server.handleFileDownload(GetDealsFilePath(id), res, func() {
		downloadAt := int(time.Now().UnixNano() / int64(time.Millisecond))
		deal, err = server.store.UpdateDealDownloadTime(deal.ID, downloadAt)
		if err != nil {
			server.log.Error().Str("dealID", deal.ID).
				Int("downloadAt", downloadAt).Err(err).Msg("failed to record deal downloadAt time")
		}

		server.stats.PostJobRun(deal)
		server.stats.PostReputation(deal.ResourceProvider,
			stats.NewReputationBuilder().
				WithJobCompletedNoValidation(true).
				Build(),
		)
	}); err != nil {
		return EmptyResponse{}, err
	}

	return EmptyResponse{}, nil
}

func (server *solverServer) handleFileDownload(dirPath string, res corehttp.ResponseWriter, onCompletion func()) *http.HTTPError {
	// Read directory contents
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return &http.HTTPError{
			Message:    fmt.Sprintf("error reading directory: %s", err.Error()),
			StatusCode: corehttp.StatusNotFound,
		}
	}

	// Find the first regular file
	var targetFile os.DirEntry
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			server.log.Warn().Err(err).Msg("expected file renamed or moved")
			continue
		}
		if info.Mode().IsRegular() {
			targetFile = file
			break
		}
	}

	if targetFile == nil {
		return &http.HTTPError{
			Message:    "no regular files found in directory",
			StatusCode: corehttp.StatusNotFound,
		}
	}

	// Get the actual filename and path
	filename := targetFile.Name()
	filePath := filepath.Join(dirPath, filename)

	file, err := os.Open(filePath)
	if err != nil {
		return &http.HTTPError{
			Message:    err.Error(),
			StatusCode: corehttp.StatusInternalServerError,
		}
	}
	defer file.Close()

	// Set appropriate headers using the actual filename
	res.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	res.Header().Set("Content-Type", "application/x-tar")

	_, err = io.Copy(res, file)
	if err != nil {
		return &http.HTTPError{
			Message:    err.Error(),
			StatusCode: corehttp.StatusInternalServerError,
		}
	}

	onCompletion()
	return nil
}

func (server *solverServer) uploadFiles(res corehttp.ResponseWriter, req *corehttp.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	err := func() error {
		deal, err := server.store.GetDeal(id)
		if err != nil {
			server.log.Error().Err(err).Str("cid", id).Msg("error loading deal")
			return err
		}
		if deal == nil {
			server.log.Error().Str("cid", id).Msg("deal not found")
			return err
		}

		if deal.State == data.GetAgreementStateIndex("JobTimedOut") {
			server.log.Trace().
				Str("cid", deal.ID).
				Str("address", deal.ResourceProvider).
				Msg("attempted file upload for timed out job for deal")
			return fmt.Errorf("job with deal ID %s timed out", deal.ID)
		}

		signerAddress, err := http.CheckSignature(req)
		if err != nil {
			server.log.Error().Err(err).Msg("error checking signature")
			return err
		}
		// Only the resource provider in a deal can upload job outputs
		if signerAddress != deal.ResourceProvider {
			return fmt.Errorf("resource provider address does not match signer address")
		}

		// Get the directory path
		dirPath, err := EnsureDealsFilePath(id)
		if err != nil {
			return err
		}

		contentDisposition := req.Header.Get("Content-Disposition")
		filename := "results.tar" // default filename

		if contentDisposition != "" {
			// Look for the filename parameter
			if parts := strings.Split(contentDisposition, "filename="); len(parts) > 1 {
				filename = strings.Trim(parts[1], `"`)
			}
		}

		// Create the file path with original filename
		filePath := filepath.Join(dirPath, filename)

		// Create the file
		f, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer f.Close()

		// Copy the data
		_, err = io.Copy(f, req.Body)
		if err != nil {
			return err
		}

		uploadAt := int(time.Now().UnixNano() / int64(time.Millisecond))
		_, err = server.store.UpdateDealUploadTime(deal.ID, uploadAt)
		if err != nil {
			server.log.Error().Str("dealID", deal.ID).
				Int("uploadAt", uploadAt).Err(err).Msg("failed to record deal uploadAt time")
		}

		return nil
	}()

	if err != nil {
		server.log.Error().Err(err).Msg("error for route")
		corehttp.Error(res, err.Error(), corehttp.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(res).Encode(data.Result{
		DataID: id,
	})
	if err != nil {
		server.log.Error().Err(err).Msg("error encoding json")
		corehttp.Error(res, err.Error(), corehttp.StatusInternalServerError)
		return
	}
}

func (server *solverServer) jobOfferDownloadFiles(res corehttp.ResponseWriter, req *corehttp.Request) (EmptyResponse, error) {
	vars := mux.Vars(req)
	id := vars["id"]

	jobOffer, err := server.store.GetJobOffer(id)
	if err != nil {
		server.log.Error().Err(err).Msg("error loading job offer")
		return EmptyResponse{}, &http.HTTPError{
			Message:    err.Error(),
			StatusCode: corehttp.StatusInternalServerError,
		}
	}
	if jobOffer == nil {
		return EmptyResponse{}, &http.HTTPError{
			Message:    err.Error(),
			StatusCode: corehttp.StatusNotFound,
		}
	}

	signerAddress, err := http.CheckSignature(req)
	if err != nil {
		server.log.Error().Err(err).Msgf("error checking signature")
		return EmptyResponse{}, &http.HTTPError{
			Message:    errors.New("not authorized").Error(),
			StatusCode: corehttp.StatusUnauthorized,
		}
	}

	if signerAddress != jobOffer.JobCreator {
		server.log.Debug().
			Str("signer", signerAddress).
			Str("address", jobOffer.JobCreator).
			Msg("signer address does not match job creator address")
		return EmptyResponse{}, &http.HTTPError{
			Message:    errors.New("not authorized").Error(),
			StatusCode: corehttp.StatusUnauthorized,
		}
	}

	server.updateJobStates(jobOffer.DealID, "ResultsAccepted")

	// Retrieve deal for stats reporting
	deal, err := server.store.GetDeal(jobOffer.DealID)
	if err != nil {
		return EmptyResponse{}, &http.HTTPError{
			Message:    "error loading deal",
			StatusCode: corehttp.StatusInternalServerError,
		}
	}
	if deal == nil {
		return EmptyResponse{}, &http.HTTPError{
			Message:    "deal not found",
			StatusCode: corehttp.StatusNotFound,
		}
	}

	if err := server.handleFileDownload(GetDealsFilePath(jobOffer.DealID), res, func() {
		downloadAt := int(time.Now().UnixNano() / int64(time.Millisecond))
		deal, err = server.store.UpdateDealDownloadTime(deal.ID, downloadAt)
		if err != nil {
			server.log.Error().Str("dealID", deal.ID).
				Int("downloadAt", downloadAt).Err(err).Msg("failed to record deal downloadAt time")
		}

		server.stats.PostJobRun(deal)
		server.stats.PostReputation(deal.ResourceProvider,
			stats.NewReputationBuilder().
				WithJobCompletedNoValidation(true).
				Build(),
		)
	}); err != nil {
		return EmptyResponse{}, err
	}

	return EmptyResponse{}, nil
}

// Validation Service

func (server *solverServer) getValidationToken(res corehttp.ResponseWriter, req *corehttp.Request) (*http.ValidationToken, error) {
	// Check signature
	signerAddress, err := http.CheckSignature(req)
	if err != nil {
		server.log.Warn().Err(err).Msg("error checking signature")
		return nil, err
	}

	// Create token with signer address sub
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   "rp_" + signerAddress,
		"iss":   "kafka-auth",
		"aud":   "kafka-broker",
		"scope": "kafka-cluster",
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Duration(server.options.AccessControl.ValidationTokenExpiration) * time.Second).Unix(),
		"jti":   uuid.New().String(),
	})

	// Add the key ID to the token header
	token.Header["kid"] = server.options.AccessControl.ValidationTokenKid

	// Sign the token
	secret := []byte(server.options.AccessControl.ValidationTokenSecret)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		server.log.Error().Err(err).Msg("failed to sign token")
		return nil, errors.New("failed to sign token")
	}

	// Respond with the JWT
	return &http.ValidationToken{JWT: tokenString}, nil
}

func (server *solverServer) updateJobStates(dealID string, state string) error {
	deal, err := server.store.GetDeal(dealID)
	if err != nil {
		return err
	}

	_, err = server.controller.updateDealState(deal.Deal.ID, data.GetAgreementStateIndex(state))
	if err != nil {
		return err
	}
	// update the job offer state
	_, err = server.controller.updateJobOfferState(deal.Deal.JobOffer.ID, deal.ID, data.GetAgreementStateIndex(state))
	if err != nil {
		return err
	}
	// update the resource offer state
	_, err = server.controller.updateResourceOfferState(deal.Deal.ResourceOffer.ID, deal.ID, data.GetAgreementStateIndex(state))
	if err != nil {
		return err
	}

	server.controller.writeEvent(SolverEvent{
		EventType: DealStateUpdated,
		Deal:      deal,
	})

	return nil
}
