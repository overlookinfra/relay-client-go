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
	"time"
)

// InviteAllOf An invitation to become a member of an account
type InviteAllOf struct {
	// The time that the invite was accepted
	AcceptedAt NullableTime `json:"accepted_at,omitempty"`
	Acceptor *UserSummary `json:"acceptor,omitempty"`
	Inviter *UserSummary `json:"inviter,omitempty"`
	// The initial roles to grant to the user
	Roles *[]RoleSummary `json:"roles,omitempty"`
	// The time that the invite was last sent
	SentAt NullableTime `json:"sent_at,omitempty"`
}

// NewInviteAllOf instantiates a new InviteAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteAllOf() *InviteAllOf {
	this := InviteAllOf{}
	return &this
}

// NewInviteAllOfWithDefaults instantiates a new InviteAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteAllOfWithDefaults() *InviteAllOf {
	this := InviteAllOf{}
	return &this
}

// GetAcceptedAt returns the AcceptedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteAllOf) GetAcceptedAt() time.Time {
	if o == nil || o.AcceptedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.AcceptedAt.Get()
}

// GetAcceptedAtOk returns a tuple with the AcceptedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteAllOf) GetAcceptedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AcceptedAt.Get(), o.AcceptedAt.IsSet()
}

// HasAcceptedAt returns a boolean if a field has been set.
func (o *InviteAllOf) HasAcceptedAt() bool {
	if o != nil && o.AcceptedAt.IsSet() {
		return true
	}

	return false
}

// SetAcceptedAt gets a reference to the given NullableTime and assigns it to the AcceptedAt field.
func (o *InviteAllOf) SetAcceptedAt(v time.Time) {
	o.AcceptedAt.Set(&v)
}
// SetAcceptedAtNil sets the value for AcceptedAt to be an explicit nil
func (o *InviteAllOf) SetAcceptedAtNil() {
	o.AcceptedAt.Set(nil)
}

// UnsetAcceptedAt ensures that no value is present for AcceptedAt, not even an explicit nil
func (o *InviteAllOf) UnsetAcceptedAt() {
	o.AcceptedAt.Unset()
}

// GetAcceptor returns the Acceptor field value if set, zero value otherwise.
func (o *InviteAllOf) GetAcceptor() UserSummary {
	if o == nil || o.Acceptor == nil {
		var ret UserSummary
		return ret
	}
	return *o.Acceptor
}

// GetAcceptorOk returns a tuple with the Acceptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteAllOf) GetAcceptorOk() (*UserSummary, bool) {
	if o == nil || o.Acceptor == nil {
		return nil, false
	}
	return o.Acceptor, true
}

// HasAcceptor returns a boolean if a field has been set.
func (o *InviteAllOf) HasAcceptor() bool {
	if o != nil && o.Acceptor != nil {
		return true
	}

	return false
}

// SetAcceptor gets a reference to the given UserSummary and assigns it to the Acceptor field.
func (o *InviteAllOf) SetAcceptor(v UserSummary) {
	o.Acceptor = &v
}

// GetInviter returns the Inviter field value if set, zero value otherwise.
func (o *InviteAllOf) GetInviter() UserSummary {
	if o == nil || o.Inviter == nil {
		var ret UserSummary
		return ret
	}
	return *o.Inviter
}

// GetInviterOk returns a tuple with the Inviter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteAllOf) GetInviterOk() (*UserSummary, bool) {
	if o == nil || o.Inviter == nil {
		return nil, false
	}
	return o.Inviter, true
}

// HasInviter returns a boolean if a field has been set.
func (o *InviteAllOf) HasInviter() bool {
	if o != nil && o.Inviter != nil {
		return true
	}

	return false
}

// SetInviter gets a reference to the given UserSummary and assigns it to the Inviter field.
func (o *InviteAllOf) SetInviter(v UserSummary) {
	o.Inviter = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *InviteAllOf) GetRoles() []RoleSummary {
	if o == nil || o.Roles == nil {
		var ret []RoleSummary
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteAllOf) GetRolesOk() (*[]RoleSummary, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *InviteAllOf) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []RoleSummary and assigns it to the Roles field.
func (o *InviteAllOf) SetRoles(v []RoleSummary) {
	o.Roles = &v
}

// GetSentAt returns the SentAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteAllOf) GetSentAt() time.Time {
	if o == nil || o.SentAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SentAt.Get()
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteAllOf) GetSentAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SentAt.Get(), o.SentAt.IsSet()
}

// HasSentAt returns a boolean if a field has been set.
func (o *InviteAllOf) HasSentAt() bool {
	if o != nil && o.SentAt.IsSet() {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given NullableTime and assigns it to the SentAt field.
func (o *InviteAllOf) SetSentAt(v time.Time) {
	o.SentAt.Set(&v)
}
// SetSentAtNil sets the value for SentAt to be an explicit nil
func (o *InviteAllOf) SetSentAtNil() {
	o.SentAt.Set(nil)
}

// UnsetSentAt ensures that no value is present for SentAt, not even an explicit nil
func (o *InviteAllOf) UnsetSentAt() {
	o.SentAt.Unset()
}

func (o InviteAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AcceptedAt.IsSet() {
		toSerialize["accepted_at"] = o.AcceptedAt.Get()
	}
	if o.Acceptor != nil {
		toSerialize["acceptor"] = o.Acceptor
	}
	if o.Inviter != nil {
		toSerialize["inviter"] = o.Inviter
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.SentAt.IsSet() {
		toSerialize["sent_at"] = o.SentAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInviteAllOf struct {
	value *InviteAllOf
	isSet bool
}

func (v NullableInviteAllOf) Get() *InviteAllOf {
	return v.value
}

func (v *NullableInviteAllOf) Set(val *InviteAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteAllOf(val *InviteAllOf) *NullableInviteAllOf {
	return &NullableInviteAllOf{value: val, isSet: true}
}

func (v NullableInviteAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


