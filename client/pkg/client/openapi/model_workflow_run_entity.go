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

// WorkflowRunEntity struct for WorkflowRunEntity
type WorkflowRunEntity struct {
	Access *EntityAccess `json:"access,omitempty"`
	Run *WorkflowRun `json:"run,omitempty"`
}

// NewWorkflowRunEntity instantiates a new WorkflowRunEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRunEntity() *WorkflowRunEntity {
	this := WorkflowRunEntity{}
	return &this
}

// NewWorkflowRunEntityWithDefaults instantiates a new WorkflowRunEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRunEntityWithDefaults() *WorkflowRunEntity {
	this := WorkflowRunEntity{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *WorkflowRunEntity) GetAccess() EntityAccess {
	if o == nil || o.Access == nil {
		var ret EntityAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunEntity) GetAccessOk() (*EntityAccess, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *WorkflowRunEntity) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given EntityAccess and assigns it to the Access field.
func (o *WorkflowRunEntity) SetAccess(v EntityAccess) {
	o.Access = &v
}

// GetRun returns the Run field value if set, zero value otherwise.
func (o *WorkflowRunEntity) GetRun() WorkflowRun {
	if o == nil || o.Run == nil {
		var ret WorkflowRun
		return ret
	}
	return *o.Run
}

// GetRunOk returns a tuple with the Run field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunEntity) GetRunOk() (*WorkflowRun, bool) {
	if o == nil || o.Run == nil {
		return nil, false
	}
	return o.Run, true
}

// HasRun returns a boolean if a field has been set.
func (o *WorkflowRunEntity) HasRun() bool {
	if o != nil && o.Run != nil {
		return true
	}

	return false
}

// SetRun gets a reference to the given WorkflowRun and assigns it to the Run field.
func (o *WorkflowRunEntity) SetRun(v WorkflowRun) {
	o.Run = &v
}

func (o WorkflowRunEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.Run != nil {
		toSerialize["run"] = o.Run
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowRunEntity struct {
	value *WorkflowRunEntity
	isSet bool
}

func (v NullableWorkflowRunEntity) Get() *WorkflowRunEntity {
	return v.value
}

func (v *NullableWorkflowRunEntity) Set(val *WorkflowRunEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRunEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRunEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRunEntity(val *WorkflowRunEntity) *NullableWorkflowRunEntity {
	return &NullableWorkflowRunEntity{value: val, isSet: true}
}

func (v NullableWorkflowRunEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRunEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


