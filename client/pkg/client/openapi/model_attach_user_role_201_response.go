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

// AttachUserRole201Response struct for AttachUserRole201Response
type AttachUserRole201Response struct {
	Access *EntityAccess `json:"access,omitempty"`
	Role   *RoleSummary  `json:"role,omitempty"`
}

// NewAttachUserRole201Response instantiates a new AttachUserRole201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachUserRole201Response() *AttachUserRole201Response {
	this := AttachUserRole201Response{}
	return &this
}

// NewAttachUserRole201ResponseWithDefaults instantiates a new AttachUserRole201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachUserRole201ResponseWithDefaults() *AttachUserRole201Response {
	this := AttachUserRole201Response{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *AttachUserRole201Response) GetAccess() EntityAccess {
	if o == nil || o.Access == nil {
		var ret EntityAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachUserRole201Response) GetAccessOk() (*EntityAccess, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *AttachUserRole201Response) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given EntityAccess and assigns it to the Access field.
func (o *AttachUserRole201Response) SetAccess(v EntityAccess) {
	o.Access = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *AttachUserRole201Response) GetRole() RoleSummary {
	if o == nil || o.Role == nil {
		var ret RoleSummary
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachUserRole201Response) GetRoleOk() (*RoleSummary, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *AttachUserRole201Response) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given RoleSummary and assigns it to the Role field.
func (o *AttachUserRole201Response) SetRole(v RoleSummary) {
	o.Role = &v
}

func (o AttachUserRole201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableAttachUserRole201Response struct {
	value *AttachUserRole201Response
	isSet bool
}

func (v NullableAttachUserRole201Response) Get() *AttachUserRole201Response {
	return v.value
}

func (v *NullableAttachUserRole201Response) Set(val *AttachUserRole201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachUserRole201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachUserRole201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachUserRole201Response(val *AttachUserRole201Response) *NullableAttachUserRole201Response {
	return &NullableAttachUserRole201Response{value: val, isSet: true}
}

func (v NullableAttachUserRole201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachUserRole201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
