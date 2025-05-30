package jobcreator

import (
	"fmt"
	"time"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/data"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/module"
	"github.com/rs/zerolog/log"

	nanoid "github.com/matoous/go-nanoid/v2"
)

// this will load the module in the offer options
// and hoist the machine spec from the module into the offer
func getJobOfferFromOptions(options JobCreatorOfferOptions, jobCreatorAddress string) (data.JobOffer, error) {
	// process the given module so we know what spec the job is asking for
	// this will also validate the module the user is asking for
	loadedModule, err := module.LoadModule(options.Module, options.Inputs)
	if err != nil {
		return data.JobOffer{}, fmt.Errorf("error loading module: %s opts=%+v", err.Error(), options)
	}

	if module.HasInputFiles(loadedModule.InputFiles) {
		err = module.ValidateInputFiles(options.InputsPath, loadedModule.InputFiles)
		if err != nil {
			log.Error().Err(err).
				Str("inputsPath", options.InputsPath).
				Any("inputFiles", loadedModule.InputFiles).
				Msg("failed to validate input files")
			return data.JobOffer{}, fmt.Errorf("error reading input files: %s", err.Error())
		}
	}

	// Generate a nonce to make sure the job offer is unique
	nonce, err := nanoid.New()
	if err != nil {
		return data.JobOffer{}, fmt.Errorf("error generating job offer nonce: %v", err)
	}

	return data.JobOffer{
		// assign CreatedAt to the current millisecond timestamp
		CreatedAt:  int(time.Now().UnixNano() / int64(time.Millisecond)),
		Nonce:      nonce,
		JobCreator: jobCreatorAddress,
		Module:     options.Module,
		Spec:       loadedModule.Machine,
		Inputs:     options.Inputs,
		InputFiles: loadedModule.InputFiles,
		Mode:       options.Mode,
		Pricing:    options.Pricing,
		Timeouts:   options.Timeouts,
		Services:   options.Services,
		Target:     options.Target,
	}, nil
}
