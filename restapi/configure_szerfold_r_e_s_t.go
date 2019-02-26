// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/go-swagger/go-swagger/cmd/swagger/restapi/operations"
	"github.com/go-swagger/go-swagger/cmd/swagger/restapi/operations/daily"
	"github.com/go-swagger/go-swagger/cmd/swagger/restapi/operations/todos"
)

//go:generate swagger generate server --target ../../swagger --name SzerfoldREST --spec ../../../../../../../szerfoldAPI/swagger.yml

func configureFlags(api *operations.SzerfoldRESTAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SzerfoldRESTAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.DailyAddOneHandler = daily.AddOneHandlerFunc(func(params daily.AddOneParams) middleware.Responder {
		return middleware.NotImplemented("operation daily.AddOne has not yet been implemented")
	})
	api.TodosDestroyOneHandler = todos.DestroyOneHandlerFunc(func(params todos.DestroyOneParams) middleware.Responder {
		return middleware.NotImplemented("operation todos.DestroyOne has not yet been implemented")
	})
	api.DailyGetDailyHandler = daily.GetDailyHandlerFunc(func(params daily.GetDailyParams) middleware.Responder {
		return middleware.NotImplemented("operation daily.GetDaily has not yet been implemented")
	})
	api.TodosUpdateOneHandler = todos.UpdateOneHandlerFunc(func(params todos.UpdateOneParams) middleware.Responder {
		return middleware.NotImplemented("operation todos.UpdateOne has not yet been implemented")
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
