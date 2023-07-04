/*
SLOs

OpenAPI schema for SLOs endpoints

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slo

import (
	"encoding/json"
)

// checks if the CreateCompositeSloResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCompositeSloResponse{}

// CreateCompositeSloResponse struct for CreateCompositeSloResponse
type CreateCompositeSloResponse struct {
	Id string `json:"id"`
}

// NewCreateCompositeSloResponse instantiates a new CreateCompositeSloResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCompositeSloResponse(id string) *CreateCompositeSloResponse {
	this := CreateCompositeSloResponse{}
	this.Id = id
	return &this
}

// NewCreateCompositeSloResponseWithDefaults instantiates a new CreateCompositeSloResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCompositeSloResponseWithDefaults() *CreateCompositeSloResponse {
	this := CreateCompositeSloResponse{}
	return &this
}

// GetId returns the Id field value
func (o *CreateCompositeSloResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateCompositeSloResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateCompositeSloResponse) SetId(v string) {
	o.Id = v
}

func (o CreateCompositeSloResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCompositeSloResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableCreateCompositeSloResponse struct {
	value *CreateCompositeSloResponse
	isSet bool
}

func (v NullableCreateCompositeSloResponse) Get() *CreateCompositeSloResponse {
	return v.value
}

func (v *NullableCreateCompositeSloResponse) Set(val *CreateCompositeSloResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCompositeSloResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCompositeSloResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCompositeSloResponse(val *CreateCompositeSloResponse) *NullableCreateCompositeSloResponse {
	return &NullableCreateCompositeSloResponse{value: val, isSet: true}
}

func (v NullableCreateCompositeSloResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCompositeSloResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
