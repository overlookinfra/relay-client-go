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

// NotificationSummaryAllOf A summary of a notification
type NotificationSummaryAllOf struct {
	// The attributes of this notification
	Attributes []string `json:"attributes,omitempty"`
	// The type of event that created the event
	Type string `json:"type"`
}

// NewNotificationSummaryAllOf instantiates a new NotificationSummaryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSummaryAllOf(type_ string) *NotificationSummaryAllOf {
	this := NotificationSummaryAllOf{}
	this.Type = type_
	return &this
}

// NewNotificationSummaryAllOfWithDefaults instantiates a new NotificationSummaryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSummaryAllOfWithDefaults() *NotificationSummaryAllOf {
	this := NotificationSummaryAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NotificationSummaryAllOf) GetAttributes() []string {
	if o == nil || o.Attributes == nil {
		var ret []string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSummaryAllOf) GetAttributesOk() ([]string, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NotificationSummaryAllOf) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []string and assigns it to the Attributes field.
func (o *NotificationSummaryAllOf) SetAttributes(v []string) {
	o.Attributes = v
}

// GetType returns the Type field value
func (o *NotificationSummaryAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotificationSummaryAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NotificationSummaryAllOf) SetType(v string) {
	o.Type = v
}

func (o NotificationSummaryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationSummaryAllOf struct {
	value *NotificationSummaryAllOf
	isSet bool
}

func (v NullableNotificationSummaryAllOf) Get() *NotificationSummaryAllOf {
	return v.value
}

func (v *NullableNotificationSummaryAllOf) Set(val *NotificationSummaryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSummaryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSummaryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSummaryAllOf(val *NotificationSummaryAllOf) *NullableNotificationSummaryAllOf {
	return &NullableNotificationSummaryAllOf{value: val, isSet: true}
}

func (v NullableNotificationSummaryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSummaryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
