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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// DefaultAPIController binds http requests to an api service and writes the service results to the http response
type DefaultAPIController struct {
	service DefaultAPIServicer
	errorHandler ErrorHandler
}

// DefaultAPIOption for how the controller is set up.
type DefaultAPIOption func(*DefaultAPIController)

// WithDefaultAPIErrorHandler inject ErrorHandler into controller
func WithDefaultAPIErrorHandler(h ErrorHandler) DefaultAPIOption {
	return func(c *DefaultAPIController) {
		c.errorHandler = h
	}
}

// NewDefaultAPIController creates a default api controller
func NewDefaultAPIController(s DefaultAPIServicer, opts ...DefaultAPIOption) Router {
	controller := &DefaultAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DefaultAPIController
func (c *DefaultAPIController) Routes() Routes {
	return Routes{
		"CatsGet": Route{
			strings.ToUpper("Get"),
			"/cats",
			c.CatsGet,
		},
		"CatsIdGet": Route{
			strings.ToUpper("Get"),
			"/cats/{id}",
			c.CatsIdGet,
		},
		"CatsIdPut": Route{
			strings.ToUpper("Put"),
			"/cats/{id}",
			c.CatsIdPut,
		},
		"CatsPost": Route{
			strings.ToUpper("Post"),
			"/cats",
			c.CatsPost,
		},
	}
}

// CatsGet - Get a list of all cats
func (c *DefaultAPIController) CatsGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.CatsGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CatsIdGet - Get information about a specific cat
func (c *DefaultAPIController) CatsIdGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	if idParam == "" {
		c.errorHandler(w, r, &RequiredError{"id"}, nil)
		return
	}
	result, err := c.service.CatsIdGet(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CatsIdPut - Update information about a specific cat
func (c *DefaultAPIController) CatsIdPut(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	if idParam == "" {
		c.errorHandler(w, r, &RequiredError{"id"}, nil)
		return
	}
	catParam := Cat{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&catParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCatRequired(catParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertCatConstraints(catParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CatsIdPut(r.Context(), idParam, catParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CatsPost - Create a new cat
func (c *DefaultAPIController) CatsPost(w http.ResponseWriter, r *http.Request) {
	createdCatParam := CreatedCat{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&createdCatParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCreatedCatRequired(createdCatParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertCreatedCatConstraints(createdCatParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CatsPost(r.Context(), createdCatParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}