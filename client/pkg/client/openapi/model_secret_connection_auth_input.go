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

// SecretConnectionAuthInput Connection authentication values with form-based configuration
type SecretConnectionAuthInput struct {
	// The fields to use for authenticating
	Fields map[string]interface{} `json:"fields"`
	// The type of authentication provided by the connection
	Type string `json:"type"`
}

// NewSecretConnectionAuthInput instantiates a new SecretConnectionAuthInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretConnectionAuthInput(fields map[string]interface{}, type_ string) *SecretConnectionAuthInput {
	this := SecretConnectionAuthInput{}
	this.Fields = fields
	this.Type = type_
	return &this
}

// NewSecretConnectionAuthInputWithDefaults instantiates a new SecretConnectionAuthInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretConnectionAuthInputWithDefaults() *SecretConnectionAuthInput {
	this := SecretConnectionAuthInput{}
	return &this
}

// GetFields returns the Fields field value
func (o *SecretConnectionAuthInput) GetFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *SecretConnectionAuthInput) GetFieldsOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *SecretConnectionAuthInput) SetFields(v map[string]interface{}) {
	o.Fields = v
}

// GetType returns the Type field value
func (o *SecretConnectionAuthInput) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecretConnectionAuthInput) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SecretConnectionAuthInput) SetType(v string) {
	o.Type = v
}

func (o SecretConnectionAuthInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fields"] = o.Fields
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSecretConnectionAuthInput struct {
	value *SecretConnectionAuthInput
	isSet bool
}

func (v NullableSecretConnectionAuthInput) Get() *SecretConnectionAuthInput {
	return v.value
}

func (v *NullableSecretConnectionAuthInput) Set(val *SecretConnectionAuthInput) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretConnectionAuthInput) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretConnectionAuthInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretConnectionAuthInput(val *SecretConnectionAuthInput) *NullableSecretConnectionAuthInput {
	return &NullableSecretConnectionAuthInput{value: val, isSet: true}
}

func (v NullableSecretConnectionAuthInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretConnectionAuthInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


