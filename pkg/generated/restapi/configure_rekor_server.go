// This file is safe to edit. Once it exists it will not be overwritten

//
// Copyright 2025 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package restapi

import (
	"crypto/tls"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/sigstore/rekorv2/pkg/generated/restapi/operations"
	"github.com/sigstore/rekorv2/pkg/generated/restapi/operations/entries"
	"github.com/sigstore/rekorv2/pkg/generated/restapi/operations/pubkey"
	"github.com/sigstore/rekorv2/pkg/generated/restapi/operations/tlog"
)

//go:generate swagger generate server --target ../../generated --name RekorServer --spec ../../../openapi.yaml --principal interface{} --exclude-main

func configureFlags(_ *operations.RekorServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.RekorServerAPI) http.Handler {
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

	api.ApplicationXPemFileProducer = runtime.ProducerFunc(func(_ io.Writer, _ interface{}) error {
		return errors.NotImplemented("applicationXPemFile producer has not yet been implemented")
	})
	api.JSONProducer = runtime.JSONProducer()

	if api.EntriesCreateLogEntryHandler == nil {
		api.EntriesCreateLogEntryHandler = entries.CreateLogEntryHandlerFunc(func(_ entries.CreateLogEntryParams) middleware.Responder {
			return middleware.NotImplemented("operation entries.CreateLogEntry has not yet been implemented")
		})
	}
	if api.TlogGetLogInfoHandler == nil {
		api.TlogGetLogInfoHandler = tlog.GetLogInfoHandlerFunc(func(_ tlog.GetLogInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation tlog.GetLogInfo has not yet been implemented")
		})
	}
	if api.PubkeyGetPublicKeyHandler == nil {
		api.PubkeyGetPublicKeyHandler = pubkey.GetPublicKeyHandlerFunc(func(_ pubkey.GetPublicKeyParams) middleware.Responder {
			return middleware.NotImplemented("operation pubkey.GetPublicKey has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(_ *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(_ *http.Server, _, _ string) {
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
