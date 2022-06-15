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

// UpdateUser200ResponseAllOf The response type for updating an account user
type UpdateUser200ResponseAllOf struct {
	User *User `json:"user,omitempty"`
}

// NewUpdateUser200ResponseAllOf instantiates a new UpdateUser200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUser200ResponseAllOf() *UpdateUser200ResponseAllOf {
	this := UpdateUser200ResponseAllOf{}
	return &this
}

// NewUpdateUser200ResponseAllOfWithDefaults instantiates a new UpdateUser200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUser200ResponseAllOfWithDefaults() *UpdateUser200ResponseAllOf {
	this := UpdateUser200ResponseAllOf{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UpdateUser200ResponseAllOf) GetUser() User {
	if o == nil || o.User == nil {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUser200ResponseAllOf) GetUserOk() (*User, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UpdateUser200ResponseAllOf) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *UpdateUser200ResponseAllOf) SetUser(v User) {
	o.User = &v
}

func (o UpdateUser200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateUser200ResponseAllOf struct {
	value *UpdateUser200ResponseAllOf
	isSet bool
}

func (v NullableUpdateUser200ResponseAllOf) Get() *UpdateUser200ResponseAllOf {
	return v.value
}

func (v *NullableUpdateUser200ResponseAllOf) Set(val *UpdateUser200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUser200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUser200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUser200ResponseAllOf(val *UpdateUser200ResponseAllOf) *NullableUpdateUser200ResponseAllOf {
	return &NullableUpdateUser200ResponseAllOf{value: val, isSet: true}
}

func (v NullableUpdateUser200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUser200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
