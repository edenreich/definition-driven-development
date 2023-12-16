# \CatsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCat**](CatsAPI.md#CreateCat) | **Post** /cats | Create a new cat
[**GetCat**](CatsAPI.md#GetCat) | **Get** /cats/{id} | Get information about a specific cat
[**GetCats**](CatsAPI.md#GetCats) | **Get** /cats | Get a list of all cats
[**UpdateCat**](CatsAPI.md#UpdateCat) | **Put** /cats/{id} | Update information about a specific cat



## CreateCat

> CreateCat(ctx).CreatedCat(createdCat).Execute()

Create a new cat

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/edenreich/definition-driven-development"
)

func main() {
	createdCat := *openapiclient.NewCreatedCat("Name_example", "Breed_example", int32(123)) // CreatedCat | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatsAPI.CreateCat(context.Background()).CreatedCat(createdCat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatsAPI.CreateCat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdCat** | [**CreatedCat**](CreatedCat.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCat

> GetCat(ctx, id).Execute()

Get information about a specific cat

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/edenreich/definition-driven-development"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatsAPI.GetCat(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatsAPI.GetCat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCats

> []Cat GetCats(ctx).Execute()

Get a list of all cats

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/edenreich/definition-driven-development"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatsAPI.GetCats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatsAPI.GetCats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCats`: []Cat
	fmt.Fprintf(os.Stdout, "Response from `CatsAPI.GetCats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatsRequest struct via the builder pattern


### Return type

[**[]Cat**](Cat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCat

> UpdateCat(ctx, id).UpdatedCat(updatedCat).Execute()

Update information about a specific cat

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/edenreich/definition-driven-development"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updatedCat := *openapiclient.NewUpdatedCat("Name_example", "Breed_example", int32(123)) // UpdatedCat | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatsAPI.UpdateCat(context.Background(), id).UpdatedCat(updatedCat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatsAPI.UpdateCat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedCat** | [**UpdatedCat**](UpdatedCat.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

