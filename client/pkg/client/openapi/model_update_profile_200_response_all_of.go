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

// UpdateProfile200ResponseAllOf The response type for updating your user profile
type UpdateProfile200ResponseAllOf struct {
	User *UserProfile `json:"user,omitempty"`
}

// NewUpdateProfile200ResponseAllOf instantiates a new UpdateProfile200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProfile200ResponseAllOf() *UpdateProfile200ResponseAllOf {
	this := UpdateProfile200ResponseAllOf{}
	return &this
}

// NewUpdateProfile200ResponseAllOfWithDefaults instantiates a new UpdateProfile200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProfile200ResponseAllOfWithDefaults() *UpdateProfile200ResponseAllOf {
	this := UpdateProfile200ResponseAllOf{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UpdateProfile200ResponseAllOf) GetUser() UserProfile {
	if o == nil || o.User == nil {
		var ret UserProfile
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfile200ResponseAllOf) GetUserOk() (*UserProfile, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UpdateProfile200ResponseAllOf) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserProfile and assigns it to the User field.
func (o *UpdateProfile200ResponseAllOf) SetUser(v UserProfile) {
	o.User = &v
}

func (o UpdateProfile200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateProfile200ResponseAllOf struct {
	value *UpdateProfile200ResponseAllOf
	isSet bool
}

func (v NullableUpdateProfile200ResponseAllOf) Get() *UpdateProfile200ResponseAllOf {
	return v.value
}

func (v *NullableUpdateProfile200ResponseAllOf) Set(val *UpdateProfile200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProfile200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProfile200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProfile200ResponseAllOf(val *UpdateProfile200ResponseAllOf) *NullableUpdateProfile200ResponseAllOf {
	return &NullableUpdateProfile200ResponseAllOf{value: val, isSet: true}
}

func (v NullableUpdateProfile200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProfile200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
