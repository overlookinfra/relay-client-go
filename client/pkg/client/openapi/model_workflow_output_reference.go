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

// WorkflowOutputReference A workflow step output reference in yaml
type WorkflowOutputReference struct {
	From string `json:"from"`
	Name string `json:"name"`
}

// NewWorkflowOutputReference instantiates a new WorkflowOutputReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowOutputReference(from string, name string) *WorkflowOutputReference {
	this := WorkflowOutputReference{}
	this.From = from
	this.Name = name
	return &this
}

// NewWorkflowOutputReferenceWithDefaults instantiates a new WorkflowOutputReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowOutputReferenceWithDefaults() *WorkflowOutputReference {
	this := WorkflowOutputReference{}
	return &this
}

// GetFrom returns the From field value
func (o *WorkflowOutputReference) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *WorkflowOutputReference) GetFromOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *WorkflowOutputReference) SetFrom(v string) {
	o.From = v
}

// GetName returns the Name field value
func (o *WorkflowOutputReference) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkflowOutputReference) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkflowOutputReference) SetName(v string) {
	o.Name = v
}

func (o WorkflowOutputReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["from"] = o.From
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowOutputReference struct {
	value *WorkflowOutputReference
	isSet bool
}

func (v NullableWorkflowOutputReference) Get() *WorkflowOutputReference {
	return v.value
}

func (v *NullableWorkflowOutputReference) Set(val *WorkflowOutputReference) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowOutputReference) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowOutputReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowOutputReference(val *WorkflowOutputReference) *NullableWorkflowOutputReference {
	return &NullableWorkflowOutputReference{value: val, isSet: true}
}

func (v NullableWorkflowOutputReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowOutputReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


