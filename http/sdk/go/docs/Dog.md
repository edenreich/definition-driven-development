# Dog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Breed** | **string** |  | 
**Age** | **int32** |  | 

## Methods

### NewDog

`func NewDog(name string, breed string, age int32, ) *Dog`

NewDog instantiates a new Dog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDogWithDefaults

`func NewDogWithDefaults() *Dog`

NewDogWithDefaults instantiates a new Dog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Dog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Dog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Dog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Dog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dog) SetName(v string)`

SetName sets Name field to given value.


### GetBreed

`func (o *Dog) GetBreed() string`

GetBreed returns the Breed field if non-nil, zero value otherwise.

### GetBreedOk

`func (o *Dog) GetBreedOk() (*string, bool)`

GetBreedOk returns a tuple with the Breed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreed

`func (o *Dog) SetBreed(v string)`

SetBreed sets Breed field to given value.


### GetAge

`func (o *Dog) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Dog) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Dog) SetAge(v int32)`

SetAge sets Age field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


