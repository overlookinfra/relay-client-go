/*
Relay API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v20200615
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// ConnectionWithDisplayOnceFields struct for ConnectionWithDisplayOnceFields
type ConnectionWithDisplayOnceFields struct {
	// The unique ID of this connection
	Id string `json:"id"`
	// A descriptive connection name
	Name string `json:"name"`
	// This connection's type identifier
	Type string `json:"type"`
	// The set of capabilities to enable for a connection
	Capabilities []ConnectionProviderCapability `json:"capabilities"`
	// Time of creation
	CreatedAt time.Time `json:"created_at"`
	// Time of last update
	UpdatedAt    time.Time              `json:"updated_at"`
	Auth         ConnectionAuthStatus   `json:"auth"`
	Availability ConnectionAvailability `json:"availability"`
	// The workflows being used by this connection
	Workflows []ConnectionWorkflowSummary `json:"workflows,omitempty"`
	// The additional fields to show to the user
	DisplayOnceFields map[string]interface{} `json:"display_once_fields,omitempty"`
}

// NewConnectionWithDisplayOnceFields instantiates a new ConnectionWithDisplayOnceFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionWithDisplayOnceFields(id string, name string, type_ string, capabilities []ConnectionProviderCapability, createdAt time.Time, updatedAt time.Time, auth ConnectionAuthStatus, availability ConnectionAvailability) *ConnectionWithDisplayOnceFields {
	this := ConnectionWithDisplayOnceFields{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.Capabilities = capabilities
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Auth = auth
	this.Availability = availability
	return &this
}

// NewConnectionWithDisplayOnceFieldsWithDefaults instantiates a new ConnectionWithDisplayOnceFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionWithDisplayOnceFieldsWithDefaults() *ConnectionWithDisplayOnceFields {
	this := ConnectionWithDisplayOnceFields{}
	return &this
}

// GetId returns the Id field value
func (o *ConnectionWithDisplayOnceFields) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConnectionWithDisplayOnceFields) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConnectionWithDisplayOnceFields) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ConnectionWithDisplayOnceFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConnectionWithDisplayOnceFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConnectionWithDisplayOnceFields) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *ConnectionWithDisplayOnceFields) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConnectionWithDisplayOnceFields) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConnectionWithDisplayOnceFields) SetType(v string) {
	o.Type = v
}

// GetCapabilities returns the Capabilities field value
func (o *ConnectionWithDisplayOnceFields) GetCapabilities() []ConnectionProviderCapability {
	if o == nil {
		var ret []ConnectionProviderCapability
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *ConnectionWithDisplayOnceFields) GetCapabilitiesOk() ([]ConnectionProviderCapability, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// SetCapabilities sets field value
func (o *ConnectionWithDisplayOnceFields) SetCapabilities(v []ConnectionProviderCapability) {
	o.Capabilities = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ConnectionWithDisplayOnceFields) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ConnectionWithDisplayOnceFields) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ConnectionWithDisplayOnceFields) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ConnectionWithDisplayOnceFields) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ConnectionWithDisplayOnceFields) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ConnectionWithDisplayOnceFields) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetAuth returns the Auth field value
func (o *ConnectionWithDisplayOnceFields) GetAuth() ConnectionAuthStatus {
	if o == nil {
		var ret ConnectionAuthStatus
		return ret
	}

	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *ConnectionWithDisplayOnceFields) GetAuthOk() (*ConnectionAuthStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value
func (o *ConnectionWithDisplayOnceFields) SetAuth(v ConnectionAuthStatus) {
	o.Auth = v
}

// GetAvailability returns the Availability field value
func (o *ConnectionWithDisplayOnceFields) GetAvailability() ConnectionAvailability {
	if o == nil {
		var ret ConnectionAvailability
		return ret
	}

	return o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value
// and a boolean to check if the value has been set.
func (o *ConnectionWithDisplayOnceFields) GetAvailabilityOk() (*ConnectionAvailability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Availability, true
}

// SetAvailability sets field value
func (o *ConnectionWithDisplayOnceFields) SetAvailability(v ConnectionAvailability) {
	o.Availability = v
}

// GetWorkflows returns the Workflows field value if set, zero value otherwise.
func (o *ConnectionWithDisplayOnceFields) GetWorkflows() []ConnectionWorkflowSummary {
	if o == nil || o.Workflows == nil {
		var ret []ConnectionWorkflowSummary
		return ret
	}
	return o.Workflows
}

// GetWorkflowsOk returns a tuple with the Workflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionWithDisplayOnceFields) GetWorkflowsOk() ([]ConnectionWorkflowSummary, bool) {
	if o == nil || o.Workflows == nil {
		return nil, false
	}
	return o.Workflows, true
}

// HasWorkflows returns a boolean if a field has been set.
func (o *ConnectionWithDisplayOnceFields) HasWorkflows() bool {
	if o != nil && o.Workflows != nil {
		return true
	}

	return false
}

// SetWorkflows gets a reference to the given []ConnectionWorkflowSummary and assigns it to the Workflows field.
func (o *ConnectionWithDisplayOnceFields) SetWorkflows(v []ConnectionWorkflowSummary) {
	o.Workflows = v
}

// GetDisplayOnceFields returns the DisplayOnceFields field value if set, zero value otherwise.
func (o *ConnectionWithDisplayOnceFields) GetDisplayOnceFields() map[string]interface{} {
	if o == nil || o.DisplayOnceFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.DisplayOnceFields
}

// GetDisplayOnceFieldsOk returns a tuple with the DisplayOnceFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionWithDisplayOnceFields) GetDisplayOnceFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.DisplayOnceFields == nil {
		return nil, false
	}
	return o.DisplayOnceFields, true
}

// HasDisplayOnceFields returns a boolean if a field has been set.
func (o *ConnectionWithDisplayOnceFields) HasDisplayOnceFields() bool {
	if o != nil && o.DisplayOnceFields != nil {
		return true
	}

	return false
}

// SetDisplayOnceFields gets a reference to the given map[string]interface{} and assigns it to the DisplayOnceFields field.
func (o *ConnectionWithDisplayOnceFields) SetDisplayOnceFields(v map[string]interface{}) {
	o.DisplayOnceFields = v
}

func (o ConnectionWithDisplayOnceFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["capabilities"] = o.Capabilities
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["auth"] = o.Auth
	}
	if true {
		toSerialize["availability"] = o.Availability
	}
	if o.Workflows != nil {
		toSerialize["workflows"] = o.Workflows
	}
	if o.DisplayOnceFields != nil {
		toSerialize["display_once_fields"] = o.DisplayOnceFields
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionWithDisplayOnceFields struct {
	value *ConnectionWithDisplayOnceFields
	isSet bool
}

func (v NullableConnectionWithDisplayOnceFields) Get() *ConnectionWithDisplayOnceFields {
	return v.value
}

func (v *NullableConnectionWithDisplayOnceFields) Set(val *ConnectionWithDisplayOnceFields) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionWithDisplayOnceFields) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionWithDisplayOnceFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionWithDisplayOnceFields(val *ConnectionWithDisplayOnceFields) *NullableConnectionWithDisplayOnceFields {
	return &NullableConnectionWithDisplayOnceFields{value: val, isSet: true}
}

func (v NullableConnectionWithDisplayOnceFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionWithDisplayOnceFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
