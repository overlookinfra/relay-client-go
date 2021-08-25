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
	"time"
)

// ApprovalWorkflowRunStepState struct for ApprovalWorkflowRunStepState
type ApprovalWorkflowRunStepState struct {
	// Time at which the step execution ended
	EndedAt NullableTime `json:"ended_at,omitempty"`
	// Time at which step execution started
	StartedAt NullableTime `json:"started_at,omitempty"`
	// Workflow run step status
	Status string `json:"status"`
	// Workflow run step approval for manual gates
	Approval *string `json:"approval,omitempty"`
	// Time at which the step was approved or rejected
	ApprovalSubmittedAt NullableTime `json:"approval_submitted_at,omitempty"`
	Approver *UserSummary `json:"approver,omitempty"`
	// Type of step
	Type string `json:"type"`
}

// NewApprovalWorkflowRunStepState instantiates a new ApprovalWorkflowRunStepState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalWorkflowRunStepState(status string, type_ string) *ApprovalWorkflowRunStepState {
	this := ApprovalWorkflowRunStepState{}
	this.Status = status
	var approval string = "waiting"
	this.Approval = &approval
	this.Type = type_
	return &this
}

// NewApprovalWorkflowRunStepStateWithDefaults instantiates a new ApprovalWorkflowRunStepState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalWorkflowRunStepStateWithDefaults() *ApprovalWorkflowRunStepState {
	this := ApprovalWorkflowRunStepState{}
	var approval string = "waiting"
	this.Approval = &approval
	return &this
}

// GetEndedAt returns the EndedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApprovalWorkflowRunStepState) GetEndedAt() time.Time {
	if o == nil || o.EndedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndedAt.Get()
}

// GetEndedAtOk returns a tuple with the EndedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalWorkflowRunStepState) GetEndedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndedAt.Get(), o.EndedAt.IsSet()
}

// HasEndedAt returns a boolean if a field has been set.
func (o *ApprovalWorkflowRunStepState) HasEndedAt() bool {
	if o != nil && o.EndedAt.IsSet() {
		return true
	}

	return false
}

// SetEndedAt gets a reference to the given NullableTime and assigns it to the EndedAt field.
func (o *ApprovalWorkflowRunStepState) SetEndedAt(v time.Time) {
	o.EndedAt.Set(&v)
}
// SetEndedAtNil sets the value for EndedAt to be an explicit nil
func (o *ApprovalWorkflowRunStepState) SetEndedAtNil() {
	o.EndedAt.Set(nil)
}

// UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
func (o *ApprovalWorkflowRunStepState) UnsetEndedAt() {
	o.EndedAt.Unset()
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApprovalWorkflowRunStepState) GetStartedAt() time.Time {
	if o == nil || o.StartedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalWorkflowRunStepState) GetStartedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// HasStartedAt returns a boolean if a field has been set.
func (o *ApprovalWorkflowRunStepState) HasStartedAt() bool {
	if o != nil && o.StartedAt.IsSet() {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given NullableTime and assigns it to the StartedAt field.
func (o *ApprovalWorkflowRunStepState) SetStartedAt(v time.Time) {
	o.StartedAt.Set(&v)
}
// SetStartedAtNil sets the value for StartedAt to be an explicit nil
func (o *ApprovalWorkflowRunStepState) SetStartedAtNil() {
	o.StartedAt.Set(nil)
}

// UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
func (o *ApprovalWorkflowRunStepState) UnsetStartedAt() {
	o.StartedAt.Unset()
}

// GetStatus returns the Status field value
func (o *ApprovalWorkflowRunStepState) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApprovalWorkflowRunStepState) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApprovalWorkflowRunStepState) SetStatus(v string) {
	o.Status = v
}

// GetApproval returns the Approval field value if set, zero value otherwise.
func (o *ApprovalWorkflowRunStepState) GetApproval() string {
	if o == nil || o.Approval == nil {
		var ret string
		return ret
	}
	return *o.Approval
}

// GetApprovalOk returns a tuple with the Approval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalWorkflowRunStepState) GetApprovalOk() (*string, bool) {
	if o == nil || o.Approval == nil {
		return nil, false
	}
	return o.Approval, true
}

// HasApproval returns a boolean if a field has been set.
func (o *ApprovalWorkflowRunStepState) HasApproval() bool {
	if o != nil && o.Approval != nil {
		return true
	}

	return false
}

// SetApproval gets a reference to the given string and assigns it to the Approval field.
func (o *ApprovalWorkflowRunStepState) SetApproval(v string) {
	o.Approval = &v
}

// GetApprovalSubmittedAt returns the ApprovalSubmittedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApprovalWorkflowRunStepState) GetApprovalSubmittedAt() time.Time {
	if o == nil || o.ApprovalSubmittedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ApprovalSubmittedAt.Get()
}

// GetApprovalSubmittedAtOk returns a tuple with the ApprovalSubmittedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalWorkflowRunStepState) GetApprovalSubmittedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApprovalSubmittedAt.Get(), o.ApprovalSubmittedAt.IsSet()
}

// HasApprovalSubmittedAt returns a boolean if a field has been set.
func (o *ApprovalWorkflowRunStepState) HasApprovalSubmittedAt() bool {
	if o != nil && o.ApprovalSubmittedAt.IsSet() {
		return true
	}

	return false
}

// SetApprovalSubmittedAt gets a reference to the given NullableTime and assigns it to the ApprovalSubmittedAt field.
func (o *ApprovalWorkflowRunStepState) SetApprovalSubmittedAt(v time.Time) {
	o.ApprovalSubmittedAt.Set(&v)
}
// SetApprovalSubmittedAtNil sets the value for ApprovalSubmittedAt to be an explicit nil
func (o *ApprovalWorkflowRunStepState) SetApprovalSubmittedAtNil() {
	o.ApprovalSubmittedAt.Set(nil)
}

// UnsetApprovalSubmittedAt ensures that no value is present for ApprovalSubmittedAt, not even an explicit nil
func (o *ApprovalWorkflowRunStepState) UnsetApprovalSubmittedAt() {
	o.ApprovalSubmittedAt.Unset()
}

// GetApprover returns the Approver field value if set, zero value otherwise.
func (o *ApprovalWorkflowRunStepState) GetApprover() UserSummary {
	if o == nil || o.Approver == nil {
		var ret UserSummary
		return ret
	}
	return *o.Approver
}

// GetApproverOk returns a tuple with the Approver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalWorkflowRunStepState) GetApproverOk() (*UserSummary, bool) {
	if o == nil || o.Approver == nil {
		return nil, false
	}
	return o.Approver, true
}

// HasApprover returns a boolean if a field has been set.
func (o *ApprovalWorkflowRunStepState) HasApprover() bool {
	if o != nil && o.Approver != nil {
		return true
	}

	return false
}

// SetApprover gets a reference to the given UserSummary and assigns it to the Approver field.
func (o *ApprovalWorkflowRunStepState) SetApprover(v UserSummary) {
	o.Approver = &v
}

// GetType returns the Type field value
func (o *ApprovalWorkflowRunStepState) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApprovalWorkflowRunStepState) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApprovalWorkflowRunStepState) SetType(v string) {
	o.Type = v
}

func (o ApprovalWorkflowRunStepState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EndedAt.IsSet() {
		toSerialize["ended_at"] = o.EndedAt.Get()
	}
	if o.StartedAt.IsSet() {
		toSerialize["started_at"] = o.StartedAt.Get()
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Approval != nil {
		toSerialize["approval"] = o.Approval
	}
	if o.ApprovalSubmittedAt.IsSet() {
		toSerialize["approval_submitted_at"] = o.ApprovalSubmittedAt.Get()
	}
	if o.Approver != nil {
		toSerialize["approver"] = o.Approver
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableApprovalWorkflowRunStepState struct {
	value *ApprovalWorkflowRunStepState
	isSet bool
}

func (v NullableApprovalWorkflowRunStepState) Get() *ApprovalWorkflowRunStepState {
	return v.value
}

func (v *NullableApprovalWorkflowRunStepState) Set(val *ApprovalWorkflowRunStepState) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalWorkflowRunStepState) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalWorkflowRunStepState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalWorkflowRunStepState(val *ApprovalWorkflowRunStepState) *NullableApprovalWorkflowRunStepState {
	return &NullableApprovalWorkflowRunStepState{value: val, isSet: true}
}

func (v NullableApprovalWorkflowRunStepState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalWorkflowRunStepState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


