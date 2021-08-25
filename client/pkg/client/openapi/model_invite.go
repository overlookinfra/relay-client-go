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

// Invite struct for Invite
type Invite struct {
	// The unique identifier of this invite
	Id string `json:"id"`
	// The email address of the invitee
	Email string `json:"email"`
	// The full name of the invitee
	Name string `json:"name"`
	// Time of creation
	CreatedAt time.Time `json:"created_at"`
	// Time of last update
	UpdatedAt time.Time `json:"updated_at"`
	// The time that the invite was accepted
	AcceptedAt NullableTime `json:"accepted_at,omitempty"`
	Acceptor *UserSummary `json:"acceptor,omitempty"`
	Inviter *UserSummary `json:"inviter,omitempty"`
	// The initial roles to grant to the user
	Roles *[]RoleSummary `json:"roles,omitempty"`
	// The time that the invite was last sent
	SentAt NullableTime `json:"sent_at,omitempty"`
}

// NewInvite instantiates a new Invite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvite(id string, email string, name string, createdAt time.Time, updatedAt time.Time) *Invite {
	this := Invite{}
	this.Id = id
	this.Email = email
	this.Name = name
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewInviteWithDefaults instantiates a new Invite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteWithDefaults() *Invite {
	this := Invite{}
	return &this
}

// GetId returns the Id field value
func (o *Invite) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Invite) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Invite) SetId(v string) {
	o.Id = v
}

// GetEmail returns the Email field value
func (o *Invite) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *Invite) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *Invite) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value
func (o *Invite) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Invite) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Invite) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Invite) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Invite) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Invite) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Invite) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Invite) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Invite) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetAcceptedAt returns the AcceptedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Invite) GetAcceptedAt() time.Time {
	if o == nil || o.AcceptedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.AcceptedAt.Get()
}

// GetAcceptedAtOk returns a tuple with the AcceptedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Invite) GetAcceptedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AcceptedAt.Get(), o.AcceptedAt.IsSet()
}

// HasAcceptedAt returns a boolean if a field has been set.
func (o *Invite) HasAcceptedAt() bool {
	if o != nil && o.AcceptedAt.IsSet() {
		return true
	}

	return false
}

// SetAcceptedAt gets a reference to the given NullableTime and assigns it to the AcceptedAt field.
func (o *Invite) SetAcceptedAt(v time.Time) {
	o.AcceptedAt.Set(&v)
}
// SetAcceptedAtNil sets the value for AcceptedAt to be an explicit nil
func (o *Invite) SetAcceptedAtNil() {
	o.AcceptedAt.Set(nil)
}

// UnsetAcceptedAt ensures that no value is present for AcceptedAt, not even an explicit nil
func (o *Invite) UnsetAcceptedAt() {
	o.AcceptedAt.Unset()
}

// GetAcceptor returns the Acceptor field value if set, zero value otherwise.
func (o *Invite) GetAcceptor() UserSummary {
	if o == nil || o.Acceptor == nil {
		var ret UserSummary
		return ret
	}
	return *o.Acceptor
}

// GetAcceptorOk returns a tuple with the Acceptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invite) GetAcceptorOk() (*UserSummary, bool) {
	if o == nil || o.Acceptor == nil {
		return nil, false
	}
	return o.Acceptor, true
}

// HasAcceptor returns a boolean if a field has been set.
func (o *Invite) HasAcceptor() bool {
	if o != nil && o.Acceptor != nil {
		return true
	}

	return false
}

// SetAcceptor gets a reference to the given UserSummary and assigns it to the Acceptor field.
func (o *Invite) SetAcceptor(v UserSummary) {
	o.Acceptor = &v
}

// GetInviter returns the Inviter field value if set, zero value otherwise.
func (o *Invite) GetInviter() UserSummary {
	if o == nil || o.Inviter == nil {
		var ret UserSummary
		return ret
	}
	return *o.Inviter
}

// GetInviterOk returns a tuple with the Inviter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invite) GetInviterOk() (*UserSummary, bool) {
	if o == nil || o.Inviter == nil {
		return nil, false
	}
	return o.Inviter, true
}

// HasInviter returns a boolean if a field has been set.
func (o *Invite) HasInviter() bool {
	if o != nil && o.Inviter != nil {
		return true
	}

	return false
}

// SetInviter gets a reference to the given UserSummary and assigns it to the Inviter field.
func (o *Invite) SetInviter(v UserSummary) {
	o.Inviter = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *Invite) GetRoles() []RoleSummary {
	if o == nil || o.Roles == nil {
		var ret []RoleSummary
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invite) GetRolesOk() (*[]RoleSummary, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *Invite) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []RoleSummary and assigns it to the Roles field.
func (o *Invite) SetRoles(v []RoleSummary) {
	o.Roles = &v
}

// GetSentAt returns the SentAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Invite) GetSentAt() time.Time {
	if o == nil || o.SentAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SentAt.Get()
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Invite) GetSentAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SentAt.Get(), o.SentAt.IsSet()
}

// HasSentAt returns a boolean if a field has been set.
func (o *Invite) HasSentAt() bool {
	if o != nil && o.SentAt.IsSet() {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given NullableTime and assigns it to the SentAt field.
func (o *Invite) SetSentAt(v time.Time) {
	o.SentAt.Set(&v)
}
// SetSentAtNil sets the value for SentAt to be an explicit nil
func (o *Invite) SetSentAtNil() {
	o.SentAt.Set(nil)
}

// UnsetSentAt ensures that no value is present for SentAt, not even an explicit nil
func (o *Invite) UnsetSentAt() {
	o.SentAt.Unset()
}

func (o Invite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
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

type NullableInvite struct {
	value *Invite
	isSet bool
}

func (v NullableInvite) Get() *Invite {
	return v.value
}

func (v *NullableInvite) Set(val *Invite) {
	v.value = val
	v.isSet = true
}

func (v NullableInvite) IsSet() bool {
	return v.isSet
}

func (v *NullableInvite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvite(val *Invite) *NullableInvite {
	return &NullableInvite{value: val, isSet: true}
}

func (v NullableInvite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


