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
	"strings"
	"time"

	"github.com/go-chi/httprate"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/http"
	"github.com/lilypad-tech/lilypad/pkg/metricsDashboard"
	"github.com/lilypad-tech/lilypad/pkg/solver/store"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel/sdk/trace"
)

type solverServer struct {
	options    http.ServerOptions
	controller *SolverController
	store      store.SolverStore
	services   data.ServiceConfig
}

func NewSolverServer(
	options http.ServerOptions,
	controller *SolverController,
	store store.SolverStore,
	services data.ServiceConfig,
) (*solverServer, error) {
	server := &solverServer{
		options:    options,
		controller: controller,
		store:      store,
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
func (solverServer *solverServer) ListenAndServe(ctx context.Context, cm *system.CleanupManager, tracerProvider *trace.TracerProvider) error {
	router := mux.NewRouter()

	subrouter := router.PathPrefix(http.API_SUB_PATH).Subrouter()

	subrouter.Use(http.CorsMiddleware)
	subrouter.Use(otelmux.Middleware("solver", otelmux.WithTracerProvider(tracerProvider)))

	exemptIPs := solverServer.options.RateLimiter.ExemptedIPs
	// TODO: re-enable exempt IP rate limiting
	// subrouter.Use(httprate.Limit(
	// 	solverServer.options.RateLimiter.RequestLimit,
	// 	time.Duration(solverServer.options.RateLimiter.WindowLength)*time.Second,
	// 	httprate.WithKeyFuncs(
	// 		exemptIPKeyFunc(exemptIPs),
	// 		httprate.KeyByEndpoint,
	// 	),
	// 	httprate.WithLimitHandler(func(w corehttp.ResponseWriter, r *corehttp.Request) {

	// 		key, _ := exemptIPKeyFunc(exemptIPs)(r)
	// 		if strings.HasPrefix(key, "exempt-") {
	// 			return
	// 		}

	// 		corehttp.Error(w, "Too Many Requests", corehttp.StatusTooManyRequests)
	// 	}),
	// ))

	subrouter.Use(httprate.Limit(
		solverServer.options.RateLimiter.RequestLimit,
		time.Duration(solverServer.options.RateLimiter.WindowLength)*time.Second,
		httprate.WithKeyFuncs(httprate.KeyByRealIP, httprate.KeyByEndpoint),
	))

	log.Info().Strs("exemptIPs", exemptIPs).Msg("Loaded rate limit exemptions")

	subrouter.HandleFunc("/job_offers", http.GetHandler(solverServer.getJobOffers)).Methods("GET")
	subrouter.HandleFunc("/job_offers", http.PostHandler(solverServer.addJobOffer)).Methods("POST")

	subrouter.HandleFunc("/job_offers/{id}", http.GetHandler(solverServer.getJobOffer)).Methods("GET")
	subrouter.HandleFunc("/job_offers/{id}/files", solverServer.jobOfferDownloadFiles).Methods("GET")

	subrouter.HandleFunc("/resource_offers", http.GetHandler(solverServer.getResourceOffers)).Methods("GET")
	subrouter.HandleFunc("/resource_offers", http.PostHandler(solverServer.addResourceOffer)).Methods("POST")

	subrouter.HandleFunc("/deals", http.GetHandler(solverServer.getDeals)).Methods("GET")
	subrouter.HandleFunc("/deals/{id}", http.GetHandler(solverServer.getDeal)).Methods("GET")

	subrouter.HandleFunc("/deals/{id}/files", solverServer.downloadFiles).Methods("GET")
	subrouter.HandleFunc("/deals/{id}/files", solverServer.uploadFiles).Methods("POST")

	subrouter.HandleFunc("/deals/{id}/result", http.GetHandler(solverServer.getResult)).Methods("GET")
	subrouter.HandleFunc("/deals/{id}/result", http.PostHandler(solverServer.addResult)).Methods("POST")

	subrouter.HandleFunc("/deals/{id}/txs/resource_provider", http.PostHandler(solverServer.updateTransactionsResourceProvider)).Methods("POST")
	subrouter.HandleFunc("/deals/{id}/txs/job_creator", http.PostHandler(solverServer.updateTransactionsJobCreator)).Methods("POST")
	subrouter.HandleFunc("/deals/{id}/txs/mediator", http.PostHandler(solverServer.updateTransactionsMediator)).Methods("POST")

	subrouter.HandleFunc("/validation_token", http.GetHandler(solverServer.getValidationToken)).Methods("GET")

	// this will fan out to all connected web socket connections
	// we read all events coming from inside the solver controller
	// and write them to anyone who is connected to us
	websocketEventChannel := make(chan []byte)

	log.Debug().Msgf("begin solverServer.controller.subscribeEvents")
	solverServer.controller.subscribeEvents(func(ev SolverEvent) {
		evBytes, err := json.Marshal(ev)
		if err != nil {
			log.Error().Msgf("Error marshalling event: %s", err.Error())
		}
		websocketEventChannel <- evBytes
	})

	http.StartWebSocketServer(
		subrouter,
		http.WEBSOCKET_SUB_PATH,
		websocketEventChannel,
		ctx,
		solverServer.connectCB,
		solverServer.disconnectCB,
	)

	srv := &corehttp.Server{
		Addr:              fmt.Sprintf("%s:%d", solverServer.options.Host, solverServer.options.Port),
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
func (solverServer *solverServer) connectCB(connParams http.WSConnectionParams) {
	if connParams.Type == "ResourceProvider" {
		metricsDashboard.TrackNodeConnectionEvent(metricsDashboard.NodeConnectionParams{
			Event:       "Connect",
			ID:          connParams.ID,
			CountryCode: connParams.CountryCode,
			IP:          connParams.IP,
		})
	}
}

func (solverServer *solverServer) disconnectCB(connParams http.WSConnectionParams) {
	if connParams.Type == "ResourceProvider" {
		metricsDashboard.TrackNodeConnectionEvent(metricsDashboard.NodeConnectionParams{
			Event:       "Disconnect",
			ID:          connParams.ID,
			CountryCode: connParams.CountryCode,
			IP:          connParams.IP,
		})
		solverServer.controller.removeUnmatchedResourceOffers(connParams.ID)
	}
}

func exemptIPKeyFunc(exemptIPs []string) func(r *corehttp.Request) (string, error) {
	return func(r *corehttp.Request) (string, error) {
		ip, err := httprate.KeyByRealIP(r)
		if err != nil {
			log.Error().Err(err).Msg("error getting real ip")
			return "", err
		}

		// Check if the IP is in the exempt list
		for _, exemptIP := range exemptIPs {
			if http.CanonicalizeIP(exemptIP) == ip {
				return "exempt-" + ip, nil
			}
		}

		return ip, nil
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
func (solverServer *solverServer) getJobOffers(res corehttp.ResponseWriter, req *corehttp.Request) ([]data.JobOfferContainer, error) {
	query := store.GetJobOffersQuery{}
	// if there is a job_creator query param then assign it
	if jobCreator := req.URL.Query().Get("job_creator"); jobCreator != "" {
		query.JobCreator = jobCreator
	}
	if notMatched := req.URL.Query().Get("not_matched"); notMatched == "true" {
		query.NotMatched = true
	}
	if includeCancelled := req.URL.Query().Get("include_cancelled"); includeCancelled == "true" {
		query.IncludeCancelled = true
	}
	return solverServer.store.GetJobOffers(query)
}

func (solverServer *solverServer) getResourceOffers(res corehttp.ResponseWriter, req *corehttp.Request) ([]data.ResourceOfferContainer, error) {
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
	return solverServer.store.GetResourceOffers(query)
}

func (solverServer *solverServer) getDeals(res corehttp.ResponseWriter, req *corehttp.Request) ([]data.DealContainer, error) {
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
	return solverServer.store.GetDeals(query)
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

func (solverServer *solverServer) getJobOffer(res corehttp.ResponseWriter, req *corehttp.Request) (data.JobOfferContainer, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	jobOffer, err := solverServer.store.GetJobOffer(id)
	if err != nil {
		return data.JobOfferContainer{}, err
	}
	return *jobOffer, nil
}

func (solverServer *solverServer) getDeal(res corehttp.ResponseWriter, req *corehttp.Request) (data.DealContainer, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	deal, err := solverServer.store.GetDeal(id)
	if err != nil {
		return data.DealContainer{}, err
	}
	if deal == nil {
		return data.DealContainer{}, fmt.Errorf("deal not found")
	}
	return *deal, nil
}

func (solverServer *solverServer) getResult(res corehttp.ResponseWriter, req *corehttp.Request) (data.Result, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	result, err := solverServer.store.GetResult(id)
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
func (solverServer *solverServer) addJobOffer(jobOffer data.JobOffer, res corehttp.ResponseWriter, req *corehttp.Request) (*data.JobOfferContainer, error) {
	signerAddress, err := http.CheckSignature(req)
	if err != nil {
		log.Error().Err(err).Msgf("error checking signature")
		return nil, err
	}
	// Only the job creator can post their job offer
	if signerAddress != jobOffer.JobCreator {
		return nil, fmt.Errorf("job creator address does not match signer address")
	}
	err = data.CheckJobOffer(jobOffer)
	if err != nil {
		log.Error().Err(err).Msgf("Error checking job offer")
		return nil, err
	}
	return solverServer.controller.addJobOffer(jobOffer)
}

func (solverServer *solverServer) addResourceOffer(resourceOffer data.ResourceOffer, res corehttp.ResponseWriter, req *corehttp.Request) (*data.ResourceOfferContainer, error) {
	versionHeader, _ := http.GetVersionFromHeaders(req)
	log.Debug().Msgf("resource provider adding offer with version header %s", versionHeader)

	signerAddress, err := http.CheckSignature(req)
	if err != nil {
		log.Error().Err(err).Msgf("error checking signature")
		return nil, err
	}
	// Only the resource provider can post their resource offer
	if signerAddress != resourceOffer.ResourceProvider {
		return nil, fmt.Errorf("resource provider address does not match signer address")
	}

	// Resource provider must be in allowlist when enabled
	if solverServer.options.AccessControl.EnableResourceProviderAllowlist {
		allowedProviders, err := solverServer.store.GetAllowedResourceProviders()
		if err != nil {
			log.Error().Err(err).Msgf("Unable to load resource provider allowlist: %s", err)
			return nil, err
		}

		if !slices.Contains(allowedProviders, resourceOffer.ResourceProvider) {
			log.Debug().Msgf("resource provider not in allowlist %s", resourceOffer.ResourceProvider)
			return nil, errors.New("resource provider not in beta program, request beta program access here: https://forms.gle/XaE3rRuXVLxTnZto7")
		}
	}

	err = data.CheckResourceOffer(resourceOffer)
	if err != nil {
		log.Error().Err(err).Msgf("Error checking resource offer")
		return nil, err
	}
	return solverServer.controller.addResourceOffer(resourceOffer)
}

func (solverServer *solverServer) addResult(results data.Result, res corehttp.ResponseWriter, req *corehttp.Request) (*data.Result, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	deal, err := solverServer.store.GetDeal(id)
	if err != nil {
		log.Error().Err(err).Msgf("error loading deal")
		return nil, err
	}
	if deal == nil {
		return nil, fmt.Errorf("deal not found")
	}
	signerAddress, err := http.CheckSignature(req)
	if err != nil {
		log.Error().Err(err).Msgf("error checking signature")
		return nil, err
	}
	// Only the resource provider in a deal can add a result
	if signerAddress != deal.ResourceProvider {
		return nil, fmt.Errorf("resource provider address does not match signer address")
	}
	err = data.CheckResult(results)
	if err != nil {
		log.Error().Err(err).Msgf("Error checking resource offer")
		return nil, err
	}
	results.DealID = id

	storedResult, err := solverServer.store.AddResult(results)
	if err != nil {
		return nil, err
	}

	err = solverServer.updateJobStates(id, "ResultsSubmitted")
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
func (solverServer *solverServer) updateTransactionsResourceProvider(payload data.DealTransactionsResourceProvider, res corehttp.ResponseWriter, req *corehttp.Request) (*data.DealContainer, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	deal, err := solverServer.store.GetDeal(id)
	if err != nil {
		log.Error().Err(err).Msgf("error loading deal")
		return nil, err
	}
	if deal == nil {
		log.Error().Err(err).Msgf("deal not found")
		return nil, fmt.Errorf("deal not found")
	}
	signerAddress, err := http.CheckSignature(req)
	if err != nil {
		log.Error().Err(err).Msgf("error checking signature")
		return nil, err
	}
	// Only the resource provider in a deal can update its transactions
	if signerAddress != deal.ResourceProvider {
		return nil, fmt.Errorf("resource provider address does not match signer address")
	}
	return solverServer.controller.updateDealTransactionsResourceProvider(id, payload)
}

func (solverServer *solverServer) updateTransactionsJobCreator(payload data.DealTransactionsJobCreator, res corehttp.ResponseWriter, req *corehttp.Request) (*data.DealContainer, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	deal, err := solverServer.store.GetDeal(id)
	if err != nil {
		log.Error().Err(err).Msgf("error loading deal")
		return nil, err
	}
	if deal == nil {
		log.Error().Err(err).Msgf("deal not found")
		return nil, fmt.Errorf("deal not found")
	}
	signerAddress, err := http.CheckSignature(req)
	if err != nil {
		log.Error().Err(err).Msgf("error checking signature")
		return nil, err
	}
	// Only the job creator in a deal can update its transactions
	if signerAddress != deal.JobCreator {
		return nil, fmt.Errorf("job creator address does not match signer address")
	}
	return solverServer.controller.updateDealTransactionsJobCreator(id, payload)
}

func (solverServer *solverServer) updateTransactionsMediator(payload data.DealTransactionsMediator, res corehttp.ResponseWriter, req *corehttp.Request) (*data.DealContainer, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	deal, err := solverServer.store.GetDeal(id)
	if err != nil {
		log.Error().Err(err).Msgf("error loading deal")
		return nil, err
	}
	if deal == nil {
		log.Error().Err(err).Msgf("deal not found")
		return nil, fmt.Errorf("deal not found")
	}
	signerAddress, err := http.CheckSignature(req)
	if err != nil {
		log.Error().Err(err).Msgf("error checking signature")
		return nil, err
	}
	// Only the mediator in a deal can update its transactions
	if signerAddress != deal.Mediator {
		return nil, fmt.Errorf("job creator address does not match mediator address")
	}
	return solverServer.controller.updateDealTransactionsMediator(id, payload)
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

func (solverServer *solverServer) downloadFiles(res corehttp.ResponseWriter, req *corehttp.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	err := func() *http.HTTPError {
		deal, err := solverServer.store.GetDeal(id)
		if err != nil {
			log.Error().Err(err).Msgf("error loading deal")
			return &http.HTTPError{
				Message:    err.Error(),
				StatusCode: corehttp.StatusInternalServerError,
			}
		}
		if deal == nil {
			return &http.HTTPError{
				Message:    err.Error(),
				StatusCode: corehttp.StatusNotFound,
			}
		}

		signerAddress, err := http.CheckSignature(req)
		if err != nil {
			log.Error().Err(err).Msgf("error checking signature")
			return &http.HTTPError{
				Message:    errors.New("not authorized").Error(),
				StatusCode: corehttp.StatusUnauthorized,
			}
		}
		// Only the job creator in a deal can download job outputs
		if signerAddress != deal.JobCreator {
			log.Error().Err(err).Msgf("job creator address does not match signer address")
			return &http.HTTPError{
				Message:    errors.New("not authorized").Error(),
				StatusCode: corehttp.StatusUnauthorized,
			}
		}
		return solverServer.handleFileDownload(GetDealsFilePath(deal.ID), res)
	}()

	if err != nil {
		log.Ctx(req.Context()).Error().Msgf("error for route: %s", err.Error())
		corehttp.Error(res, err.Error(), err.StatusCode)
		return
	}
}

func (solverServer *solverServer) uploadFiles(res corehttp.ResponseWriter, req *corehttp.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	err := func() error {
		deal, err := solverServer.store.GetDeal(id)
		if err != nil {
			log.Error().Err(err).Msgf("error loading deal")
			return err
		}
		if deal == nil {
			log.Error().Msgf("deal not found")
			return err
		}
		signerAddress, err := http.CheckSignature(req)
		if err != nil {
			log.Error().Err(err).Msgf("error checking signature")
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

		return nil
	}()

	if err != nil {
		log.Ctx(req.Context()).Error().Msgf("error for route: %s", err.Error())
		corehttp.Error(res, err.Error(), corehttp.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(res).Encode(data.Result{
		DataID: id,
	})
	if err != nil {
		log.Ctx(req.Context()).Error().Msgf("error for json encoding: %s", err.Error())
		corehttp.Error(res, err.Error(), corehttp.StatusInternalServerError)
		return
	}
}

func (solverServer *solverServer) jobOfferDownloadFiles(res corehttp.ResponseWriter, req *corehttp.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	err := func() *http.HTTPError {
		jobOffer, err := solverServer.store.GetJobOffer(id)
		if err != nil {
			log.Error().Err(err).Msgf("error loading job offer")
			return &http.HTTPError{
				Message:    err.Error(),
				StatusCode: corehttp.StatusInternalServerError,
			}
		}
		if jobOffer == nil {
			return &http.HTTPError{
				Message:    err.Error(),
				StatusCode: corehttp.StatusNotFound,
			}
		}

		signerAddress, err := http.CheckSignature(req)
		if err != nil {
			log.Error().Err(err).Msgf("error checking signature")
			return &http.HTTPError{
				Message:    errors.New("not authorized").Error(),
				StatusCode: corehttp.StatusUnauthorized,
			}
		}

		if signerAddress != jobOffer.JobCreator {
			log.Error().Err(err).Msgf("job creator address does not match signer address")
			return &http.HTTPError{
				Message:    errors.New("not authorized").Error(),
				StatusCode: corehttp.StatusUnauthorized,
			}
		}

		solverServer.updateJobStates(jobOffer.DealID, "ResultsAccepted")

		return solverServer.handleFileDownload(GetDealsFilePath(jobOffer.DealID), res)
	}()

	if err != nil {
		log.Ctx(req.Context()).Error().Msgf("error for route: %s", err.Error())
		corehttp.Error(res, err.Error(), err.StatusCode)
	}
}

func (solverServer *solverServer) handleFileDownload(dirPath string, res corehttp.ResponseWriter) *http.HTTPError {
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

	// Open and serve the file
	file, err := os.Open(filePath)
	if err != nil {
		return &http.HTTPError{
			Message:    err.Error(),
			StatusCode: corehttp.StatusInternalServerError,
		}
	}
	defer file.Close()

	res.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	res.Header().Set("Content-Type", "application/x-tar")

	_, err = io.Copy(res, file)
	if err != nil {
		return &http.HTTPError{
			Message:    err.Error(),
			StatusCode: corehttp.StatusInternalServerError,
		}
	}

	return nil
}

// Validation Service

func (solverServer *solverServer) getValidationToken(res corehttp.ResponseWriter, req *corehttp.Request) (*http.ValidationToken, error) {
	// Check signature
	signerAddress, err := http.CheckSignature(req)
	if err != nil {
		log.Warn().Err(err).Msgf("error checking signature")
		return nil, err
	}

	// Create token with signer address sub
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   "rp_" + signerAddress,
		"iss":   "kafka-auth",
		"aud":   "kafka-broker",
		"scope": "kafka-cluster",
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Duration(solverServer.options.AccessControl.ValidationTokenExpiration) * time.Second).Unix(),
		"jti":   uuid.New().String(),
	})

	// Add the key ID to the token header
	token.Header["kid"] = solverServer.options.AccessControl.ValidationTokenKid

	// Sign the token
	secret := []byte(solverServer.options.AccessControl.ValidationTokenSecret)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Error().Err(err).Msgf("failed to sign token")
		return nil, errors.New("failed to sign token")
	}

	// Respond with the JWT
	return &http.ValidationToken{JWT: tokenString}, nil
}

func (solverServer *solverServer) updateJobStates(dealID string, state string) error {
	deal, err := solverServer.store.GetDeal(dealID)
	if err != nil {
		return err
	}

	_, err = solverServer.controller.updateDealState(deal.Deal.ID, data.GetAgreementStateIndex(state))
	if err != nil {
		return err
	}
	// update the job offer state
	_, err = solverServer.controller.updateJobOfferState(deal.Deal.JobOffer.ID, deal.ID, data.GetAgreementStateIndex(state))
	if err != nil {
		return err
	}
	// update the resource offer state
	_, err = solverServer.controller.updateResourceOfferState(deal.Deal.ResourceOffer.ID, deal.ID, data.GetAgreementStateIndex(state))
	if err != nil {
		return err
	}

	solverServer.controller.writeEvent(SolverEvent{
		EventType: DealStateUpdated,
		Deal:      deal,
	})

	return nil
}
