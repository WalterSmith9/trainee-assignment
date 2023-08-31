// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/WalterSmith9/trainee-assignment/src/swagger/generated/restapi/operations"
	"github.com/WalterSmith9/trainee-assignment/src/swagger/generated/restapi/operations/segment"
	"github.com/WalterSmith9/trainee-assignment/src/swagger/generated/restapi/operations/user"
)

//go:generate swagger generate server --target ..\..\generated --name TraineeAssignment --spec ..\..\..\..\swagger.yaml --template-dir ./templates/tmpl --principal interface{}

func configureFlags(api *operations.TraineeAssignmentAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TraineeAssignmentAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.UserGetSegmentsHandler == nil {
		api.UserGetSegmentsHandler = user.GetSegmentsHandlerFunc(func(params user.GetSegmentsParams) middleware.Responder {
			return middleware.NotImplemented("operation user.GetSegments has not yet been implemented")
		})
	}
	if api.SegmentSegmentCreateHandler == nil {
		api.SegmentSegmentCreateHandler = segment.SegmentCreateHandlerFunc(func(params segment.SegmentCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation segment.SegmentCreate has not yet been implemented")
		})
	}
	if api.SegmentSegmentDeleteHandler == nil {
		api.SegmentSegmentDeleteHandler = segment.SegmentDeleteHandlerFunc(func(params segment.SegmentDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation segment.SegmentDelete has not yet been implemented")
		})
	}
	if api.UserUserEditHandler == nil {
		api.UserUserEditHandler = user.UserEditHandlerFunc(func(params user.UserEditParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserEdit has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
