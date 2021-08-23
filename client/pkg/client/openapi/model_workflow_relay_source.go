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

// WorkflowRelaySource The relay source of a workflow
type WorkflowRelaySource struct {
	// The type discriminator for this repository source
	Type *string `json:"type,omitempty"`
}

// NewWorkflowRelaySource instantiates a new WorkflowRelaySource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRelaySource() *WorkflowRelaySource {
	this := WorkflowRelaySource{}
	return &this
}

// NewWorkflowRelaySourceWithDefaults instantiates a new WorkflowRelaySource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRelaySourceWithDefaults() *WorkflowRelaySource {
	this := WorkflowRelaySource{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowRelaySource) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRelaySource) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowRelaySource) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowRelaySource) SetType(v string) {
	o.Type = &v
}

func (o WorkflowRelaySource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowRelaySource struct {
	value *WorkflowRelaySource
	isSet bool
}

func (v NullableWorkflowRelaySource) Get() *WorkflowRelaySource {
	return v.value
}

func (v *NullableWorkflowRelaySource) Set(val *WorkflowRelaySource) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRelaySource) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRelaySource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRelaySource(val *WorkflowRelaySource) *NullableWorkflowRelaySource {
	return &NullableWorkflowRelaySource{value: val, isSet: true}
}

func (v NullableWorkflowRelaySource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRelaySource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


