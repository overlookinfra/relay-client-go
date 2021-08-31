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

// UserToken User token for authentication
type UserToken struct {
	// ID of the user token
	Id *string `json:"id,omitempty"`
	// Name of the user token
	Name string `json:"name"`
	// Type of token
	Type *string `json:"type,omitempty"`
	User UserSummary `json:"user"`
}

// NewUserToken instantiates a new UserToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserToken(name string, user UserSummary) *UserToken {
	this := UserToken{}
	this.Name = name
	this.User = user
	return &this
}

// NewUserTokenWithDefaults instantiates a new UserToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTokenWithDefaults() *UserToken {
	this := UserToken{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserToken) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserToken) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserToken) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserToken) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *UserToken) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserToken) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserToken) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserToken) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserToken) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserToken) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserToken) SetType(v string) {
	o.Type = &v
}

// GetUser returns the User field value
func (o *UserToken) GetUser() UserSummary {
	if o == nil {
		var ret UserSummary
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserToken) GetUserOk() (*UserSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserToken) SetUser(v UserSummary) {
	o.User = v
}

func (o UserToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableUserToken struct {
	value *UserToken
	isSet bool
}

func (v NullableUserToken) Get() *UserToken {
	return v.value
}

func (v *NullableUserToken) Set(val *UserToken) {
	v.value = val
	v.isSet = true
}

func (v NullableUserToken) IsSet() bool {
	return v.isSet
}

func (v *NullableUserToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserToken(val *UserToken) *NullableUserToken {
	return &NullableUserToken{value: val, isSet: true}
}

func (v NullableUserToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


