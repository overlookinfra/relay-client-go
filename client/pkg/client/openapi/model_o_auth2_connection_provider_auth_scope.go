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

// OAuth2ConnectionProviderAuthScope Information about an OAuth 2.0 scope for a connection
type OAuth2ConnectionProviderAuthScope struct {
	// The set of capabilities to enable for a connection
	Capabilities []ConnectionProviderCapability `json:"capabilities,omitempty"`
	// If true, this scope will be automatically added to any authentication request
	Implicit *bool `json:"implicit,omitempty"`
	// The scope name
	Name string `json:"name"`
	// For organizational purposes, the name of another scope that includes this scope
	SubsumedBy *string `json:"subsumed_by,omitempty"`
}

// NewOAuth2ConnectionProviderAuthScope instantiates a new OAuth2ConnectionProviderAuthScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ConnectionProviderAuthScope(name string) *OAuth2ConnectionProviderAuthScope {
	this := OAuth2ConnectionProviderAuthScope{}
	this.Name = name
	return &this
}

// NewOAuth2ConnectionProviderAuthScopeWithDefaults instantiates a new OAuth2ConnectionProviderAuthScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ConnectionProviderAuthScopeWithDefaults() *OAuth2ConnectionProviderAuthScope {
	this := OAuth2ConnectionProviderAuthScope{}
	return &this
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *OAuth2ConnectionProviderAuthScope) GetCapabilities() []ConnectionProviderCapability {
	if o == nil || o.Capabilities == nil {
		var ret []ConnectionProviderCapability
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ConnectionProviderAuthScope) GetCapabilitiesOk() ([]ConnectionProviderCapability, bool) {
	if o == nil || o.Capabilities == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *OAuth2ConnectionProviderAuthScope) HasCapabilities() bool {
	if o != nil && o.Capabilities != nil {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []ConnectionProviderCapability and assigns it to the Capabilities field.
func (o *OAuth2ConnectionProviderAuthScope) SetCapabilities(v []ConnectionProviderCapability) {
	o.Capabilities = v
}

// GetImplicit returns the Implicit field value if set, zero value otherwise.
func (o *OAuth2ConnectionProviderAuthScope) GetImplicit() bool {
	if o == nil || o.Implicit == nil {
		var ret bool
		return ret
	}
	return *o.Implicit
}

// GetImplicitOk returns a tuple with the Implicit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ConnectionProviderAuthScope) GetImplicitOk() (*bool, bool) {
	if o == nil || o.Implicit == nil {
		return nil, false
	}
	return o.Implicit, true
}

// HasImplicit returns a boolean if a field has been set.
func (o *OAuth2ConnectionProviderAuthScope) HasImplicit() bool {
	if o != nil && o.Implicit != nil {
		return true
	}

	return false
}

// SetImplicit gets a reference to the given bool and assigns it to the Implicit field.
func (o *OAuth2ConnectionProviderAuthScope) SetImplicit(v bool) {
	o.Implicit = &v
}

// GetName returns the Name field value
func (o *OAuth2ConnectionProviderAuthScope) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OAuth2ConnectionProviderAuthScope) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OAuth2ConnectionProviderAuthScope) SetName(v string) {
	o.Name = v
}

// GetSubsumedBy returns the SubsumedBy field value if set, zero value otherwise.
func (o *OAuth2ConnectionProviderAuthScope) GetSubsumedBy() string {
	if o == nil || o.SubsumedBy == nil {
		var ret string
		return ret
	}
	return *o.SubsumedBy
}

// GetSubsumedByOk returns a tuple with the SubsumedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ConnectionProviderAuthScope) GetSubsumedByOk() (*string, bool) {
	if o == nil || o.SubsumedBy == nil {
		return nil, false
	}
	return o.SubsumedBy, true
}

// HasSubsumedBy returns a boolean if a field has been set.
func (o *OAuth2ConnectionProviderAuthScope) HasSubsumedBy() bool {
	if o != nil && o.SubsumedBy != nil {
		return true
	}

	return false
}

// SetSubsumedBy gets a reference to the given string and assigns it to the SubsumedBy field.
func (o *OAuth2ConnectionProviderAuthScope) SetSubsumedBy(v string) {
	o.SubsumedBy = &v
}

func (o OAuth2ConnectionProviderAuthScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Capabilities != nil {
		toSerialize["capabilities"] = o.Capabilities
	}
	if o.Implicit != nil {
		toSerialize["implicit"] = o.Implicit
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.SubsumedBy != nil {
		toSerialize["subsumed_by"] = o.SubsumedBy
	}
	return json.Marshal(toSerialize)
}

type NullableOAuth2ConnectionProviderAuthScope struct {
	value *OAuth2ConnectionProviderAuthScope
	isSet bool
}

func (v NullableOAuth2ConnectionProviderAuthScope) Get() *OAuth2ConnectionProviderAuthScope {
	return v.value
}

func (v *NullableOAuth2ConnectionProviderAuthScope) Set(val *OAuth2ConnectionProviderAuthScope) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ConnectionProviderAuthScope) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ConnectionProviderAuthScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ConnectionProviderAuthScope(val *OAuth2ConnectionProviderAuthScope) *NullableOAuth2ConnectionProviderAuthScope {
	return &NullableOAuth2ConnectionProviderAuthScope{value: val, isSet: true}
}

func (v NullableOAuth2ConnectionProviderAuthScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ConnectionProviderAuthScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
