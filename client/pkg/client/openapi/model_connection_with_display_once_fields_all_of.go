/*
 * Relay API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v20200615
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ConnectionWithDisplayOnceFieldsAllOf A connection with optional display-once fields
type ConnectionWithDisplayOnceFieldsAllOf struct {
	// The additional fields to show to the user
	DisplayOnceFields *map[string]interface{} `json:"display_once_fields,omitempty"`
}

// NewConnectionWithDisplayOnceFieldsAllOf instantiates a new ConnectionWithDisplayOnceFieldsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionWithDisplayOnceFieldsAllOf() *ConnectionWithDisplayOnceFieldsAllOf {
	this := ConnectionWithDisplayOnceFieldsAllOf{}
	return &this
}

// NewConnectionWithDisplayOnceFieldsAllOfWithDefaults instantiates a new ConnectionWithDisplayOnceFieldsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionWithDisplayOnceFieldsAllOfWithDefaults() *ConnectionWithDisplayOnceFieldsAllOf {
	this := ConnectionWithDisplayOnceFieldsAllOf{}
	return &this
}

// GetDisplayOnceFields returns the DisplayOnceFields field value if set, zero value otherwise.
func (o *ConnectionWithDisplayOnceFieldsAllOf) GetDisplayOnceFields() map[string]interface{} {
	if o == nil || o.DisplayOnceFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.DisplayOnceFields
}

// GetDisplayOnceFieldsOk returns a tuple with the DisplayOnceFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionWithDisplayOnceFieldsAllOf) GetDisplayOnceFieldsOk() (*map[string]interface{}, bool) {
	if o == nil || o.DisplayOnceFields == nil {
		return nil, false
	}
	return o.DisplayOnceFields, true
}

// HasDisplayOnceFields returns a boolean if a field has been set.
func (o *ConnectionWithDisplayOnceFieldsAllOf) HasDisplayOnceFields() bool {
	if o != nil && o.DisplayOnceFields != nil {
		return true
	}

	return false
}

// SetDisplayOnceFields gets a reference to the given map[string]interface{} and assigns it to the DisplayOnceFields field.
func (o *ConnectionWithDisplayOnceFieldsAllOf) SetDisplayOnceFields(v map[string]interface{}) {
	o.DisplayOnceFields = &v
}

func (o ConnectionWithDisplayOnceFieldsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayOnceFields != nil {
		toSerialize["display_once_fields"] = o.DisplayOnceFields
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionWithDisplayOnceFieldsAllOf struct {
	value *ConnectionWithDisplayOnceFieldsAllOf
	isSet bool
}

func (v NullableConnectionWithDisplayOnceFieldsAllOf) Get() *ConnectionWithDisplayOnceFieldsAllOf {
	return v.value
}

func (v *NullableConnectionWithDisplayOnceFieldsAllOf) Set(val *ConnectionWithDisplayOnceFieldsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionWithDisplayOnceFieldsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionWithDisplayOnceFieldsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionWithDisplayOnceFieldsAllOf(val *ConnectionWithDisplayOnceFieldsAllOf) *NullableConnectionWithDisplayOnceFieldsAllOf {
	return &NullableConnectionWithDisplayOnceFieldsAllOf{value: val, isSet: true}
}

func (v NullableConnectionWithDisplayOnceFieldsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionWithDisplayOnceFieldsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


