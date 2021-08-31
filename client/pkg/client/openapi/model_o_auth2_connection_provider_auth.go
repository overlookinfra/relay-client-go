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

// OAuth2ConnectionProviderAuth Connection authentication information for the OAuth 2.0 protocol
type OAuth2ConnectionProviderAuth struct {
	// Scopes to authenticate with this connection
	Scopes *[]OAuth2ConnectionProviderAuthScope `json:"scopes,omitempty"`
	// The type of authentication provided by the connection
	Type string `json:"type"`
}

// NewOAuth2ConnectionProviderAuth instantiates a new OAuth2ConnectionProviderAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ConnectionProviderAuth(type_ string) *OAuth2ConnectionProviderAuth {
	this := OAuth2ConnectionProviderAuth{}
	this.Type = type_
	return &this
}

// NewOAuth2ConnectionProviderAuthWithDefaults instantiates a new OAuth2ConnectionProviderAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ConnectionProviderAuthWithDefaults() *OAuth2ConnectionProviderAuth {
	this := OAuth2ConnectionProviderAuth{}
	return &this
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *OAuth2ConnectionProviderAuth) GetScopes() []OAuth2ConnectionProviderAuthScope {
	if o == nil || o.Scopes == nil {
		var ret []OAuth2ConnectionProviderAuthScope
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ConnectionProviderAuth) GetScopesOk() (*[]OAuth2ConnectionProviderAuthScope, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *OAuth2ConnectionProviderAuth) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []OAuth2ConnectionProviderAuthScope and assigns it to the Scopes field.
func (o *OAuth2ConnectionProviderAuth) SetScopes(v []OAuth2ConnectionProviderAuthScope) {
	o.Scopes = &v
}

// GetType returns the Type field value
func (o *OAuth2ConnectionProviderAuth) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OAuth2ConnectionProviderAuth) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OAuth2ConnectionProviderAuth) SetType(v string) {
	o.Type = v
}

func (o OAuth2ConnectionProviderAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableOAuth2ConnectionProviderAuth struct {
	value *OAuth2ConnectionProviderAuth
	isSet bool
}

func (v NullableOAuth2ConnectionProviderAuth) Get() *OAuth2ConnectionProviderAuth {
	return v.value
}

func (v *NullableOAuth2ConnectionProviderAuth) Set(val *OAuth2ConnectionProviderAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ConnectionProviderAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ConnectionProviderAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ConnectionProviderAuth(val *OAuth2ConnectionProviderAuth) *NullableOAuth2ConnectionProviderAuth {
	return &NullableOAuth2ConnectionProviderAuth{value: val, isSet: true}
}

func (v NullableOAuth2ConnectionProviderAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ConnectionProviderAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


