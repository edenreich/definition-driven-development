/*
 * Pets API
 *
 * A simple API to manage pets
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// CatsAPIController binds http requests to an api service and writes the service results to the http response
type CatsAPIController struct {
	service      CatsAPIServicer
	errorHandler ErrorHandler
}

// CatsAPIOption for how the controller is set up.
type CatsAPIOption func(*CatsAPIController)

// WithCatsAPIErrorHandler inject ErrorHandler into controller
func WithCatsAPIErrorHandler(h ErrorHandler) CatsAPIOption {
	return func(c *CatsAPIController) {
		c.errorHandler = h
	}
}

// NewCatsAPIController creates a default api controller
func NewCatsAPIController(s CatsAPIServicer, opts ...CatsAPIOption) Router {
	controller := &CatsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the CatsAPIController
func (c *CatsAPIController) Routes() Routes {
	return Routes{
		"CreateCat": Route{
			strings.ToUpper("Post"),
			"/cats",
			c.CreateCat,
		},
		"GetCat": Route{
			strings.ToUpper("Get"),
			"/cats/{id}",
			c.GetCat,
		},
		"GetCats": Route{
			strings.ToUpper("Get"),
			"/cats",
			c.GetCats,
		},
		"UpdateCat": Route{
			strings.ToUpper("Put"),
			"/cats/{id}",
			c.UpdateCat,
		},
	}
}

// CreateCat - Create a new cat
func (c *CatsAPIController) CreateCat(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.CreateCat(r.Context(), createdCatParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w) //nolint: errcheck
}

// GetCat - Get information about a specific cat
func (c *CatsAPIController) GetCat(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	if idParam == "" {
		c.errorHandler(w, r, &RequiredError{"id"}, nil)
		return
	}
	result, err := c.service.GetCat(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w) //nolint: errcheck
}

// GetCats - Get a list of all cats
func (c *CatsAPIController) GetCats(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetCats(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w) //nolint: errcheck
}

// UpdateCat - Update information about a specific cat
func (c *CatsAPIController) UpdateCat(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	if idParam == "" {
		c.errorHandler(w, r, &RequiredError{"id"}, nil)
		return
	}
	updatedCatParam := UpdatedCat{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&updatedCatParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUpdatedCatRequired(updatedCatParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUpdatedCatConstraints(updatedCatParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateCat(r.Context(), idParam, updatedCatParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w) //nolint: errcheck
}
