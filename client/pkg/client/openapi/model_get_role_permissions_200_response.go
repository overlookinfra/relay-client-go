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

// GetRolePermissions200Response The response type for listing the permissions assigned to a role
type GetRolePermissions200Response struct {
	// The permissions assigned to this role
	Permissions []RolePermission `json:"permissions,omitempty"`
}

// NewGetRolePermissions200Response instantiates a new GetRolePermissions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRolePermissions200Response() *GetRolePermissions200Response {
	this := GetRolePermissions200Response{}
	return &this
}

// NewGetRolePermissions200ResponseWithDefaults instantiates a new GetRolePermissions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRolePermissions200ResponseWithDefaults() *GetRolePermissions200Response {
	this := GetRolePermissions200Response{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *GetRolePermissions200Response) GetPermissions() []RolePermission {
	if o == nil || o.Permissions == nil {
		var ret []RolePermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRolePermissions200Response) GetPermissionsOk() ([]RolePermission, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *GetRolePermissions200Response) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []RolePermission and assigns it to the Permissions field.
func (o *GetRolePermissions200Response) SetPermissions(v []RolePermission) {
	o.Permissions = v
}

func (o GetRolePermissions200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableGetRolePermissions200Response struct {
	value *GetRolePermissions200Response
	isSet bool
}

func (v NullableGetRolePermissions200Response) Get() *GetRolePermissions200Response {
	return v.value
}

func (v *NullableGetRolePermissions200Response) Set(val *GetRolePermissions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRolePermissions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRolePermissions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRolePermissions200Response(val *GetRolePermissions200Response) *NullableGetRolePermissions200Response {
	return &NullableGetRolePermissions200Response{value: val, isSet: true}
}

func (v NullableGetRolePermissions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRolePermissions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
