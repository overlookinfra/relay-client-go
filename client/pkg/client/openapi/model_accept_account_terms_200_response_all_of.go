/*
Relay API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v20200615
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AcceptAccountTerms200ResponseAllOf The response type for accepting the terms and conditions
type AcceptAccountTerms200ResponseAllOf struct {
	AcceptedTerms *AcceptedTerms `json:"accepted_terms,omitempty"`
}

// NewAcceptAccountTerms200ResponseAllOf instantiates a new AcceptAccountTerms200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptAccountTerms200ResponseAllOf() *AcceptAccountTerms200ResponseAllOf {
	this := AcceptAccountTerms200ResponseAllOf{}
	return &this
}

// NewAcceptAccountTerms200ResponseAllOfWithDefaults instantiates a new AcceptAccountTerms200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptAccountTerms200ResponseAllOfWithDefaults() *AcceptAccountTerms200ResponseAllOf {
	this := AcceptAccountTerms200ResponseAllOf{}
	return &this
}

// GetAcceptedTerms returns the AcceptedTerms field value if set, zero value otherwise.
func (o *AcceptAccountTerms200ResponseAllOf) GetAcceptedTerms() AcceptedTerms {
	if o == nil || o.AcceptedTerms == nil {
		var ret AcceptedTerms
		return ret
	}
	return *o.AcceptedTerms
}

// GetAcceptedTermsOk returns a tuple with the AcceptedTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptAccountTerms200ResponseAllOf) GetAcceptedTermsOk() (*AcceptedTerms, bool) {
	if o == nil || o.AcceptedTerms == nil {
		return nil, false
	}
	return o.AcceptedTerms, true
}

// HasAcceptedTerms returns a boolean if a field has been set.
func (o *AcceptAccountTerms200ResponseAllOf) HasAcceptedTerms() bool {
	if o != nil && o.AcceptedTerms != nil {
		return true
	}

	return false
}

// SetAcceptedTerms gets a reference to the given AcceptedTerms and assigns it to the AcceptedTerms field.
func (o *AcceptAccountTerms200ResponseAllOf) SetAcceptedTerms(v AcceptedTerms) {
	o.AcceptedTerms = &v
}

func (o AcceptAccountTerms200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AcceptedTerms != nil {
		toSerialize["accepted_terms"] = o.AcceptedTerms
	}
	return json.Marshal(toSerialize)
}

type NullableAcceptAccountTerms200ResponseAllOf struct {
	value *AcceptAccountTerms200ResponseAllOf
	isSet bool
}

func (v NullableAcceptAccountTerms200ResponseAllOf) Get() *AcceptAccountTerms200ResponseAllOf {
	return v.value
}

func (v *NullableAcceptAccountTerms200ResponseAllOf) Set(val *AcceptAccountTerms200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptAccountTerms200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptAccountTerms200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptAccountTerms200ResponseAllOf(val *AcceptAccountTerms200ResponseAllOf) *NullableAcceptAccountTerms200ResponseAllOf {
	return &NullableAcceptAccountTerms200ResponseAllOf{value: val, isSet: true}
}

func (v NullableAcceptAccountTerms200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptAccountTerms200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
