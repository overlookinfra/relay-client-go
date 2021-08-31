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

// UserIdentifier A unique identifier for an account user
type UserIdentifier struct {
	// The unique ID of this user
	Id string `json:"id"`
}

// NewUserIdentifier instantiates a new UserIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIdentifier(id string) *UserIdentifier {
	this := UserIdentifier{}
	this.Id = id
	return &this
}

// NewUserIdentifierWithDefaults instantiates a new UserIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdentifierWithDefaults() *UserIdentifier {
	this := UserIdentifier{}
	return &this
}

// GetId returns the Id field value
func (o *UserIdentifier) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserIdentifier) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserIdentifier) SetId(v string) {
	o.Id = v
}

func (o UserIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableUserIdentifier struct {
	value *UserIdentifier
	isSet bool
}

func (v NullableUserIdentifier) Get() *UserIdentifier {
	return v.value
}

func (v *NullableUserIdentifier) Set(val *UserIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIdentifier(val *UserIdentifier) *NullableUserIdentifier {
	return &NullableUserIdentifier{value: val, isSet: true}
}

func (v NullableUserIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


