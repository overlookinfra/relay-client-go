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

// NotificationsResponse struct for NotificationsResponse
type NotificationsResponse struct {
	Access *EntityAccess `json:"access,omitempty"`
	// A list of notifications
	Notifications *[]Notification `json:"notifications,omitempty"`
}

// NewNotificationsResponse instantiates a new NotificationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsResponse() *NotificationsResponse {
	this := NotificationsResponse{}
	return &this
}

// NewNotificationsResponseWithDefaults instantiates a new NotificationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsResponseWithDefaults() *NotificationsResponse {
	this := NotificationsResponse{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *NotificationsResponse) GetAccess() EntityAccess {
	if o == nil || o.Access == nil {
		var ret EntityAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsResponse) GetAccessOk() (*EntityAccess, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *NotificationsResponse) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given EntityAccess and assigns it to the Access field.
func (o *NotificationsResponse) SetAccess(v EntityAccess) {
	o.Access = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *NotificationsResponse) GetNotifications() []Notification {
	if o == nil || o.Notifications == nil {
		var ret []Notification
		return ret
	}
	return *o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsResponse) GetNotificationsOk() (*[]Notification, bool) {
	if o == nil || o.Notifications == nil {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *NotificationsResponse) HasNotifications() bool {
	if o != nil && o.Notifications != nil {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []Notification and assigns it to the Notifications field.
func (o *NotificationsResponse) SetNotifications(v []Notification) {
	o.Notifications = &v
}

func (o NotificationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.Notifications != nil {
		toSerialize["notifications"] = o.Notifications
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationsResponse struct {
	value *NotificationsResponse
	isSet bool
}

func (v NullableNotificationsResponse) Get() *NotificationsResponse {
	return v.value
}

func (v *NullableNotificationsResponse) Set(val *NotificationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsResponse(val *NotificationsResponse) *NullableNotificationsResponse {
	return &NullableNotificationsResponse{value: val, isSet: true}
}

func (v NullableNotificationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


