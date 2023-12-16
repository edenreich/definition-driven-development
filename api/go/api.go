/*
 * Cat API
 *
 * A simple API to manage cats
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
)



// CatsAPIRouter defines the required methods for binding the api requests to a responses for the CatsAPI
// The CatsAPIRouter implementation should parse necessary information from the http request,
// pass the data to a CatsAPIServicer to perform the required actions, then write the service results to the http response.
type CatsAPIRouter interface { 
	CatsGet(http.ResponseWriter, *http.Request)
	CatsIdGet(http.ResponseWriter, *http.Request)
	CatsIdPut(http.ResponseWriter, *http.Request)
	CatsPost(http.ResponseWriter, *http.Request)
}


// CatsAPIServicer defines the api actions for the CatsAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type CatsAPIServicer interface { 
	CatsGet(context.Context) (ImplResponse, error)
	CatsIdGet(context.Context, string) (ImplResponse, error)
	CatsIdPut(context.Context, string, Cat) (ImplResponse, error)
	CatsPost(context.Context, CreatedCat) (ImplResponse, error)
}