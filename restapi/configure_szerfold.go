// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"SzerfoldAPI/own_code"
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"SzerfoldAPI/restapi/operations"
	"SzerfoldAPI/restapi/operations/daily"

	models "SzerfoldAPI/models"
)

//go:generate swagger generate server --target ../../SzerfoldAPI --name Szerfold --spec ../swagger.yml --principal models.Principal
// swagger generate server -A SzerfoldAPI -P models.Principal -f ./swagger.yml

func configureFlags(api *operations.SzerfoldAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SzerfoldAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "x-token" header is set
	api.KeyAuth = func(token string) (*models.Principal, error) {
		conf := own_code.ReadConfiguration()
		if token == conf["Xtoken"] {
			principal := models.Principal(token)
			return &principal, nil
		} else {
			return nil, errors.New(401, "You shall not pass!")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	api.DailyAddOneHandler = daily.AddOneHandlerFunc(func(params daily.AddOneParams, principal *models.Principal) middleware.Responder {
		addedDaily := own_code.AddNewDaily(*params.Body)
		return daily.NewAddOneCreated().WithPayload(&addedDaily)
	})
	api.DailyDestroyOneHandler = daily.DestroyOneHandlerFunc(func(params daily.DestroyOneParams, principal *models.Principal) middleware.Responder {
		own_code.DeleteDailyByID(params.PathID)
		return daily.NewDestroyOneNoContent()
	})
	api.DailyGetDailyHandler = daily.GetDailyHandlerFunc(func(params daily.GetDailyParams) middleware.Responder {
		var dailies []*models.Daily
		if params.Since != nil {
			dailies = own_code.GetLatestDailies(params.Since)
		} else {
			dailies = own_code.GetDailiesbyId(params.From)
		}
		return daily.NewGetDailyOK().WithPayload(dailies)
	})
	api.DailyUpdateOneHandler = daily.UpdateOneHandlerFunc(func(params daily.UpdateOneParams, principal *models.Principal) middleware.Responder {
		updatedDaily := own_code.UpdateDaily(*params.Daily, params.ID)
		return daily.NewUpdateOneOK().WithPayload(&updatedDaily)
	})
	api.DailyGetOneHandler = daily.GetOneHandlerFunc(func(params daily.GetOneParams) middleware.Responder {
		foundDaily := own_code.GetDailyByID(params.ID)
		return daily.NewGetOneOK().WithPayload(&foundDaily)
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
