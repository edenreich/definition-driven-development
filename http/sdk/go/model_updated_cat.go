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

// checks if the UpdatedCat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatedCat{}

// UpdatedCat struct for UpdatedCat
type UpdatedCat struct {
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	Breed string `json:"breed"`
	Age int32 `json:"age"`
}

type _UpdatedCat UpdatedCat

// NewUpdatedCat instantiates a new UpdatedCat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatedCat(name string, breed string, age int32) *UpdatedCat {
	this := UpdatedCat{}
	this.Name = name
	this.Breed = breed
	this.Age = age
	return &this
}

// NewUpdatedCatWithDefaults instantiates a new UpdatedCat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatedCatWithDefaults() *UpdatedCat {
	this := UpdatedCat{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdatedCat) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatedCat) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdatedCat) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdatedCat) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *UpdatedCat) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdatedCat) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdatedCat) SetName(v string) {
	o.Name = v
}

// GetBreed returns the Breed field value
func (o *UpdatedCat) GetBreed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Breed
}

// GetBreedOk returns a tuple with the Breed field value
// and a boolean to check if the value has been set.
func (o *UpdatedCat) GetBreedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Breed, true
}

// SetBreed sets field value
func (o *UpdatedCat) SetBreed(v string) {
	o.Breed = v
}

// GetAge returns the Age field value
func (o *UpdatedCat) GetAge() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Age
}

// GetAgeOk returns a tuple with the Age field value
// and a boolean to check if the value has been set.
func (o *UpdatedCat) GetAgeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Age, true
}

// SetAge sets field value
func (o *UpdatedCat) SetAge(v int32) {
	o.Age = v
}

func (o UpdatedCat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatedCat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["breed"] = o.Breed
	toSerialize["age"] = o.Age
	return toSerialize, nil
}

func (o *UpdatedCat) UnmarshalJSON(bytes []byte) (err error) {
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

	varUpdatedCat := _UpdatedCat{}

	err = json.Unmarshal(bytes, &varUpdatedCat)

	if err != nil {
		return err
	}

	*o = UpdatedCat(varUpdatedCat)

	return err
}

type NullableUpdatedCat struct {
	value *UpdatedCat
	isSet bool
}

func (v NullableUpdatedCat) Get() *UpdatedCat {
	return v.value
}

func (v *NullableUpdatedCat) Set(val *UpdatedCat) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatedCat) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatedCat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatedCat(val *UpdatedCat) *NullableUpdatedCat {
	return &NullableUpdatedCat{value: val, isSet: true}
}

func (v NullableUpdatedCat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatedCat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

