/*
Cat API

A simple API to manage cats

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
	"fmt"
)

// checks if the CreatedCat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatedCat{}

// CreatedCat struct for CreatedCat
type CreatedCat struct {
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	Breed string `json:"breed"`
	Age int32 `json:"age"`
}

type _CreatedCat CreatedCat

// NewCreatedCat instantiates a new CreatedCat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatedCat(name string, breed string, age int32) *CreatedCat {
	this := CreatedCat{}
	this.Name = name
	this.Breed = breed
	this.Age = age
	return &this
}

// NewCreatedCatWithDefaults instantiates a new CreatedCat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatedCatWithDefaults() *CreatedCat {
	this := CreatedCat{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreatedCat) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedCat) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreatedCat) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreatedCat) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *CreatedCat) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreatedCat) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreatedCat) SetName(v string) {
	o.Name = v
}

// GetBreed returns the Breed field value
func (o *CreatedCat) GetBreed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Breed
}

// GetBreedOk returns a tuple with the Breed field value
// and a boolean to check if the value has been set.
func (o *CreatedCat) GetBreedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Breed, true
}

// SetBreed sets field value
func (o *CreatedCat) SetBreed(v string) {
	o.Breed = v
}

// GetAge returns the Age field value
func (o *CreatedCat) GetAge() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Age
}

// GetAgeOk returns a tuple with the Age field value
// and a boolean to check if the value has been set.
func (o *CreatedCat) GetAgeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Age, true
}

// SetAge sets field value
func (o *CreatedCat) SetAge(v int32) {
	o.Age = v
}

func (o CreatedCat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatedCat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["breed"] = o.Breed
	toSerialize["age"] = o.Age
	return toSerialize, nil
}

func (o *CreatedCat) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"breed",
		"age",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreatedCat := _CreatedCat{}

	err = json.Unmarshal(bytes, &varCreatedCat)

	if err != nil {
		return err
	}

	*o = CreatedCat(varCreatedCat)

	return err
}

type NullableCreatedCat struct {
	value *CreatedCat
	isSet bool
}

func (v NullableCreatedCat) Get() *CreatedCat {
	return v.value
}

func (v *NullableCreatedCat) Set(val *CreatedCat) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatedCat) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatedCat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatedCat(val *CreatedCat) *NullableCreatedCat {
	return &NullableCreatedCat{value: val, isSet: true}
}

func (v NullableCreatedCat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatedCat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


