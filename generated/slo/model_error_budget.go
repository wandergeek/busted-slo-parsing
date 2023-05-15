/*
SLOs

OpenAPI schema for SLOs endpoints

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slo

import (
	"encoding/json"
)

// checks if the ErrorBudget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorBudget{}

// ErrorBudget struct for ErrorBudget
type ErrorBudget struct {
	// The initial error budget, as 1 - objective
	Initial *float32 `json:"initial,omitempty"`
	// The error budget consummed, as a percentage of the initial value.
	Consumed *float32 `json:"consumed,omitempty"`
	// The error budget remaining, as a percentage of the initial value.
	Remaining *float32 `json:"remaining,omitempty"`
	// Only for SLO defined with occurrences budgeting method and calendar aligned time window.
	IsEstimated *bool `json:"isEstimated,omitempty"`
}

// NewErrorBudget instantiates a new ErrorBudget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorBudget() *ErrorBudget {
	this := ErrorBudget{}
	return &this
}

// NewErrorBudgetWithDefaults instantiates a new ErrorBudget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorBudgetWithDefaults() *ErrorBudget {
	this := ErrorBudget{}
	return &this
}

// GetInitial returns the Initial field value if set, zero value otherwise.
func (o *ErrorBudget) GetInitial() float32 {
	if o == nil || IsNil(o.Initial) {
		var ret float32
		return ret
	}
	return *o.Initial
}

// GetInitialOk returns a tuple with the Initial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorBudget) GetInitialOk() (*float32, bool) {
	if o == nil || IsNil(o.Initial) {
		return nil, false
	}
	return o.Initial, true
}

// HasInitial returns a boolean if a field has been set.
func (o *ErrorBudget) HasInitial() bool {
	if o != nil && !IsNil(o.Initial) {
		return true
	}

	return false
}

// SetInitial gets a reference to the given float32 and assigns it to the Initial field.
func (o *ErrorBudget) SetInitial(v float32) {
	o.Initial = &v
}

// GetConsumed returns the Consumed field value if set, zero value otherwise.
func (o *ErrorBudget) GetConsumed() float32 {
	if o == nil || IsNil(o.Consumed) {
		var ret float32
		return ret
	}
	return *o.Consumed
}

// GetConsumedOk returns a tuple with the Consumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorBudget) GetConsumedOk() (*float32, bool) {
	if o == nil || IsNil(o.Consumed) {
		return nil, false
	}
	return o.Consumed, true
}

// HasConsumed returns a boolean if a field has been set.
func (o *ErrorBudget) HasConsumed() bool {
	if o != nil && !IsNil(o.Consumed) {
		return true
	}

	return false
}

// SetConsumed gets a reference to the given float32 and assigns it to the Consumed field.
func (o *ErrorBudget) SetConsumed(v float32) {
	o.Consumed = &v
}

// GetRemaining returns the Remaining field value if set, zero value otherwise.
func (o *ErrorBudget) GetRemaining() float32 {
	if o == nil || IsNil(o.Remaining) {
		var ret float32
		return ret
	}
	return *o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorBudget) GetRemainingOk() (*float32, bool) {
	if o == nil || IsNil(o.Remaining) {
		return nil, false
	}
	return o.Remaining, true
}

// HasRemaining returns a boolean if a field has been set.
func (o *ErrorBudget) HasRemaining() bool {
	if o != nil && !IsNil(o.Remaining) {
		return true
	}

	return false
}

// SetRemaining gets a reference to the given float32 and assigns it to the Remaining field.
func (o *ErrorBudget) SetRemaining(v float32) {
	o.Remaining = &v
}

// GetIsEstimated returns the IsEstimated field value if set, zero value otherwise.
func (o *ErrorBudget) GetIsEstimated() bool {
	if o == nil || IsNil(o.IsEstimated) {
		var ret bool
		return ret
	}
	return *o.IsEstimated
}

// GetIsEstimatedOk returns a tuple with the IsEstimated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorBudget) GetIsEstimatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEstimated) {
		return nil, false
	}
	return o.IsEstimated, true
}

// HasIsEstimated returns a boolean if a field has been set.
func (o *ErrorBudget) HasIsEstimated() bool {
	if o != nil && !IsNil(o.IsEstimated) {
		return true
	}

	return false
}

// SetIsEstimated gets a reference to the given bool and assigns it to the IsEstimated field.
func (o *ErrorBudget) SetIsEstimated(v bool) {
	o.IsEstimated = &v
}

func (o ErrorBudget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorBudget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Initial) {
		toSerialize["initial"] = o.Initial
	}
	if !IsNil(o.Consumed) {
		toSerialize["consumed"] = o.Consumed
	}
	if !IsNil(o.Remaining) {
		toSerialize["remaining"] = o.Remaining
	}
	if !IsNil(o.IsEstimated) {
		toSerialize["isEstimated"] = o.IsEstimated
	}
	return toSerialize, nil
}

type NullableErrorBudget struct {
	value *ErrorBudget
	isSet bool
}

func (v NullableErrorBudget) Get() *ErrorBudget {
	return v.value
}

func (v *NullableErrorBudget) Set(val *ErrorBudget) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorBudget) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorBudget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorBudget(val *ErrorBudget) *NullableErrorBudget {
	return &NullableErrorBudget{value: val, isSet: true}
}

func (v NullableErrorBudget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorBudget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

