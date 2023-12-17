# \DogsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDogs**](DogsAPI.md#GetDogs) | **Get** /dogs | Get a list of all dogs



## GetDogs

> []Dog GetDogs(ctx).Execute()

Get a list of all dogs

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
	resp, r, err := apiClient.DogsAPI.GetDogs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DogsAPI.GetDogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDogs`: []Dog
	fmt.Fprintf(os.Stdout, "Response from `DogsAPI.GetDogs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDogsRequest struct via the builder pattern


### Return type

[**[]Dog**](Dog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

