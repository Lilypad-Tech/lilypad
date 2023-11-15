package solver

import (
	"archive/tar"
	"context"
	"encoding/json"
	"fmt"
	"io"
	corehttp "net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/http"
	"github.com/bacalhau-project/lilypad/pkg/solver/store"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type solverServer struct {
	options    http.ServerOptions
	controller *SolverController
	store      store.SolverStore
}

func NewSolverServer(
	options http.ServerOptions,
	controller *SolverController,
	store store.SolverStore,
) (*solverServer, error) {
	server := &solverServer{
		options:    options,
		controller: controller,
		store:      store,
	}
	return server, nil
}

/*
 *
 *
 *

 Routes

 *
 *
 *
*/

func (solverServer *solverServer) ListenAndServe(ctx context.Context, cm *system.CleanupManager) error {
	router := mux.NewRouter()

	subrouter := router.PathPrefix(http.API_SUB_PATH).Subrouter()

	subrouter.Use(http.CorsMiddleware)

	subrouter.HandleFunc("/job_offers", http.GetHandler(solverServer.getJobOffers)).Methods("GET")
	subrouter.HandleFunc("/job_offers", http.PostHandler(solverServer.addJobOffer)).Methods("POST")

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
	signerAddress, err := http.GetAddressFromHeaders(req)
	if err != nil {
		log.Error().Err(err).Msgf("have error parsing user address")
		return nil, err
	}
	// only the job creator can post a job offer
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
	signerAddress, err := http.GetAddressFromHeaders(req)
	if err != nil {
		log.Error().Err(err).Msgf("have error parsing user address")
		return nil, err
	}
	// only the job creator can post a job offer
	if signerAddress != resourceOffer.ResourceProvider {
		return nil, fmt.Errorf("resource provider address does not match signer address")
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
	signerAddress, err := http.GetAddressFromHeaders(req)
	if err != nil {
		log.Error().Err(err).Msgf("have error parsing user address")
		return nil, err
	}
	// only the resource provider can add a result
	if signerAddress != deal.ResourceProvider {
		return nil, fmt.Errorf("resource provider address does not match signer address")
	}
	err = data.CheckResult(results)
	if err != nil {
		log.Error().Err(err).Msgf("Error checking resource offer")
		return nil, err
	}
	results.DealID = id
	return solverServer.store.AddResult(results)
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
	signerAddress, err := http.GetAddressFromHeaders(req)
	if err != nil {
		log.Error().Err(err).Msgf("have error parsing user address")
		return nil, err
	}
	// only the job creator can post a job offer
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
	signerAddress, err := http.GetAddressFromHeaders(req)
	if err != nil {
		log.Error().Err(err).Msgf("have error parsing user address")
		return nil, err
	}
	// only the job creator can post a job offer
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
	signerAddress, err := http.GetAddressFromHeaders(req)
	if err != nil {
		log.Error().Err(err).Msgf("have error parsing user address")
		return nil, err
	}
	// only the job creator can post a job offer
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
		filesPath := GetDealsFilePath(id)
		// check if the filesPath directory exists
		if _, err := os.Stat(filesPath); os.IsNotExist(err) {
			return &http.HTTPError{
				Message:    err.Error(),
				StatusCode: corehttp.StatusNotFound,
			}
		}
		buf, err := system.GetTarBuffer(filesPath)
		if err != nil {
			return &http.HTTPError{
				Message:    err.Error(),
				StatusCode: corehttp.StatusInternalServerError,
			}
		}
		res.Header().Set("Content-Disposition", "attachment; filename=archive.tar")
		res.Header().Set("Content-Type", "application/x-tar")
		io.Copy(res, buf)
		return nil
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
		signerAddress, err := http.GetAddressFromHeaders(req)
		if err != nil {
			log.Error().Err(err).Msgf("have error parsing user address")
			return err
		}
		// only the resource provider can add a result
		if signerAddress != deal.ResourceProvider {
			return fmt.Errorf("resource provider address does not match signer address")
		}
		tr := tar.NewReader(req.Body)
		uploadPath, err := EnsureDealsFilePath(id)
		if err != nil {
			return err
		}
		for {
			header, err := tr.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
			target := filepath.Join(uploadPath, header.Name)
			switch header.Typeflag {
			case tar.TypeDir:
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			case tar.TypeReg:
				f, err := os.Create(target)
				if err != nil {
					return err
				}
				defer f.Close()
				if _, err := io.Copy(f, tr); err != nil {
					return err
				}
			}
		}
		return nil
	}()

	if err != nil {
		log.Ctx(req.Context()).Error().Msgf("error for route: %s", err.Error())
		corehttp.Error(res, err.Error(), corehttp.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(res).Encode(data.Result{
		// TODO: we need to be putting this in IPFS and calculating the CID
		DataID: id,
	})
	if err != nil {
		log.Ctx(req.Context()).Error().Msgf("error for json encoding: %s", err.Error())
		corehttp.Error(res, err.Error(), corehttp.StatusInternalServerError)
		return
	}
}
