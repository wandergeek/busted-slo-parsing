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

// checks if the IndicatorPropertiesApmAvailabilityParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndicatorPropertiesApmAvailabilityParams{}

// IndicatorPropertiesApmAvailabilityParams An object containing the indicator parameters.
type IndicatorPropertiesApmAvailabilityParams struct {
	// The APM service name
	Service string `json:"service"`
	// The APM service environment or \"*\"
	Environment string `json:"environment"`
	// The APM transaction type or \"*\"
	TransactionType string `json:"transactionType"`
	// The APM transaction name or \"*\"
	TransactionName string `json:"transactionName"`
	// KQL query used for filtering the data
	Filter *string `json:"filter,omitempty"`
	// The index used by APM metrics
	Index string `json:"index"`
}

// NewIndicatorPropertiesApmAvailabilityParams instantiates a new IndicatorPropertiesApmAvailabilityParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndicatorPropertiesApmAvailabilityParams(service string, environment string, transactionType string, transactionName string, index string) *IndicatorPropertiesApmAvailabilityParams {
	this := IndicatorPropertiesApmAvailabilityParams{}
	this.Service = service
	this.Environment = environment
	this.TransactionType = transactionType
	this.TransactionName = transactionName
	this.Index = index
	return &this
}

// NewIndicatorPropertiesApmAvailabilityParamsWithDefaults instantiates a new IndicatorPropertiesApmAvailabilityParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndicatorPropertiesApmAvailabilityParamsWithDefaults() *IndicatorPropertiesApmAvailabilityParams {
	this := IndicatorPropertiesApmAvailabilityParams{}
	return &this
}

// GetService returns the Service field value
func (o *IndicatorPropertiesApmAvailabilityParams) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *IndicatorPropertiesApmAvailabilityParams) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *IndicatorPropertiesApmAvailabilityParams) SetService(v string) {
	o.Service = v
}

// GetEnvironment returns the Environment field value
func (o *IndicatorPropertiesApmAvailabilityParams) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *IndicatorPropertiesApmAvailabilityParams) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *IndicatorPropertiesApmAvailabilityParams) SetEnvironment(v string) {
	o.Environment = v
}

// GetTransactionType returns the TransactionType field value
func (o *IndicatorPropertiesApmAvailabilityParams) GetTransactionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value
// and a boolean to check if the value has been set.
func (o *IndicatorPropertiesApmAvailabilityParams) GetTransactionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionType, true
}

// SetTransactionType sets field value
func (o *IndicatorPropertiesApmAvailabilityParams) SetTransactionType(v string) {
	o.TransactionType = v
}

// GetTransactionName returns the TransactionName field value
func (o *IndicatorPropertiesApmAvailabilityParams) GetTransactionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionName
}

// GetTransactionNameOk returns a tuple with the TransactionName field value
// and a boolean to check if the value has been set.
func (o *IndicatorPropertiesApmAvailabilityParams) GetTransactionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionName, true
}

// SetTransactionName sets field value
func (o *IndicatorPropertiesApmAvailabilityParams) SetTransactionName(v string) {
	o.TransactionName = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *IndicatorPropertiesApmAvailabilityParams) GetFilter() string {
	if o == nil || IsNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorPropertiesApmAvailabilityParams) GetFilterOk() (*string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *IndicatorPropertiesApmAvailabilityParams) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *IndicatorPropertiesApmAvailabilityParams) SetFilter(v string) {
	o.Filter = &v
}

// GetIndex returns the Index field value
func (o *IndicatorPropertiesApmAvailabilityParams) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *IndicatorPropertiesApmAvailabilityParams) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *IndicatorPropertiesApmAvailabilityParams) SetIndex(v string) {
	o.Index = v
}

func (o IndicatorPropertiesApmAvailabilityParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndicatorPropertiesApmAvailabilityParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["service"] = o.Service
	toSerialize["environment"] = o.Environment
	toSerialize["transactionType"] = o.TransactionType
	toSerialize["transactionName"] = o.TransactionName
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	toSerialize["index"] = o.Index
	return toSerialize, nil
}

type NullableIndicatorPropertiesApmAvailabilityParams struct {
	value *IndicatorPropertiesApmAvailabilityParams
	isSet bool
}

func (v NullableIndicatorPropertiesApmAvailabilityParams) Get() *IndicatorPropertiesApmAvailabilityParams {
	return v.value
}

func (v *NullableIndicatorPropertiesApmAvailabilityParams) Set(val *IndicatorPropertiesApmAvailabilityParams) {
	v.value = val
	v.isSet = true
}

func (v NullableIndicatorPropertiesApmAvailabilityParams) IsSet() bool {
	return v.isSet
}

func (v *NullableIndicatorPropertiesApmAvailabilityParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndicatorPropertiesApmAvailabilityParams(val *IndicatorPropertiesApmAvailabilityParams) *NullableIndicatorPropertiesApmAvailabilityParams {
	return &NullableIndicatorPropertiesApmAvailabilityParams{value: val, isSet: true}
}

func (v NullableIndicatorPropertiesApmAvailabilityParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndicatorPropertiesApmAvailabilityParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
