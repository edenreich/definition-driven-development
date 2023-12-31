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
	"context"
	"errors"
	"net/http"
)

// CatsAPIService is a service that implements the logic for the CatsAPIServicer
// This service should implement the business logic for every endpoint for the CatsAPI API.
// Include any external packages or services that will be required by this service.
type CatsAPIService struct {
}

// NewCatsAPIService creates a default api service
func NewCatsAPIService() CatsAPIServicer {
	return &CatsAPIService{}
}

// CreateCat - Create a new cat
func (s *CatsAPIService) CreateCat(ctx context.Context, createdCat CreatedCat) (ImplResponse, error) {
	// TODO - update CreateCat with the required logic for this service method.
	// Add api_cats_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, {}) or use other options such as http.Ok ...
	// return Response(201, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateCat method not implemented")
}

// GetCat - Get information about a specific cat
func (s *CatsAPIService) GetCat(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update GetCat with the required logic for this service method.
	// Add api_cats_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	// return Response(200, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetCat method not implemented")
}

// GetCats - Get a list of all cats
func (s *CatsAPIService) GetCats(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetCats with the required logic for this service method.
	// Add api_cats_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []Cat{}) or use other options such as http.Ok ...
	// return Response(200, []Cat{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetCats method not implemented")
}

// UpdateCat - Update information about a specific cat
func (s *CatsAPIService) UpdateCat(ctx context.Context, id string, updatedCat UpdatedCat) (ImplResponse, error) {
	// TODO - update UpdateCat with the required logic for this service method.
	// Add api_cats_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	// return Response(200, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateCat method not implemented")
}
