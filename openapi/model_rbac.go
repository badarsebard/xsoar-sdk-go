/*
Cortex XSOAR API

This is the public REST API to integrate with the Cortex XSOAR server. HTTP request can be sent using any HTTP-client.  For an example dedicated client take a look at: https://github.com/demisto/demisto-py.  Requests must include API-key that can be generated in the Cortex XSOAR web client under 'Settings' -> 'Integrations' -> 'API keys'   Optimistic Locking and Versioning\\:  When using Cortex XSOAR REST API, you will need to make sure to work on the latest version of the item (incident, entry, etc.), otherwise, you will get a DB version error (which not allow you to override a newer item). In addition, you can pass 'version\\: -1' to force data override (make sure that other users data might be lost).  Assume that Alice and Bob both read the same data from Cortex XSOAR server, then they both changed the data, and then both tried to write the new versions back to the server. Whose changes should be saved? Alice’s? Bob’s? To solve this, each data item in Cortex XSOAR has a numeric incremental version. If Alice saved an item with version 4 and Bob trying to save the same item with version 3, Cortex XSOAR will rollback Bob request and returns a DB version conflict error. Bob will need to get the latest item and work on it so Alice work will not get lost.  Example request using 'curl'\\:  ``` curl 'https://hostname:443/incidents/search' -H 'content-type: application/json' -H 'accept: application/json' -H 'Authorization: <API Key goes here>' --data-binary '{\"filter\":{\"query\":\"-status:closed -category:job\",\"period\":{\"by\":\"day\",\"fromValue\":7}}}' --compressed ```

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// RBAC RBAC holds roles data for domain entities
type RBAC struct {
	AllRead      *bool `json:"allRead,omitempty"`
	AllReadWrite *bool `json:"allReadWrite,omitempty"`
	// Who has created this event - relevant only for manual incidents
	DbotCreatedBy *string `json:"dbotCreatedBy,omitempty"`
	// Internal field to make queries on role faster
	HasRole              *bool `json:"hasRole,omitempty"`
	PreviousAllRead      *bool `json:"previousAllRead,omitempty"`
	PreviousAllReadWrite *bool `json:"previousAllReadWrite,omitempty"`
	// Do not change this field manually
	PreviousRoles *[]string `json:"previousRoles,omitempty"`
	// The role assigned to this investigation
	Roles                      *[]string `json:"roles,omitempty"`
	XsoarHasReadOnlyRole       *bool     `json:"xsoarHasReadOnlyRole,omitempty"`
	XsoarPreviousReadOnlyRoles *[]string `json:"xsoarPreviousReadOnlyRoles,omitempty"`
	XsoarReadOnlyRoles         *[]string `json:"xsoarReadOnlyRoles,omitempty"`
}

// NewRBAC instantiates a new RBAC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRBAC() *RBAC {
	this := RBAC{}
	return &this
}

// NewRBACWithDefaults instantiates a new RBAC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRBACWithDefaults() *RBAC {
	this := RBAC{}
	return &this
}

// GetAllRead returns the AllRead field value if set, zero value otherwise.
func (o *RBAC) GetAllRead() bool {
	if o == nil || o.AllRead == nil {
		var ret bool
		return ret
	}
	return *o.AllRead
}

// GetAllReadOk returns a tuple with the AllRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBAC) GetAllReadOk() (*bool, bool) {
	if o == nil || o.AllRead == nil {
		return nil, false
	}
	return o.AllRead, true
}

// HasAllRead returns a boolean if a field has been set.
func (o *RBAC) HasAllRead() bool {
	if o != nil && o.AllRead != nil {
		return true
	}

	return false
}

// SetAllRead gets a reference to the given bool and assigns it to the AllRead field.
func (o *RBAC) SetAllRead(v bool) {
	o.AllRead = &v
}

// GetAllReadWrite returns the AllReadWrite field value if set, zero value otherwise.
func (o *RBAC) GetAllReadWrite() bool {
	if o == nil || o.AllReadWrite == nil {
		var ret bool
		return ret
	}
	return *o.AllReadWrite
}

// GetAllReadWriteOk returns a tuple with the AllReadWrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBAC) GetAllReadWriteOk() (*bool, bool) {
	if o == nil || o.AllReadWrite == nil {
		return nil, false
	}
	return o.AllReadWrite, true
}

// HasAllReadWrite returns a boolean if a field has been set.
func (o *RBAC) HasAllReadWrite() bool {
	if o != nil && o.AllReadWrite != nil {
		return true
	}

	return false
}

// SetAllReadWrite gets a reference to the given bool and assigns it to the AllReadWrite field.
func (o *RBAC) SetAllReadWrite(v bool) {
	o.AllReadWrite = &v
}

// GetDbotCreatedBy returns the DbotCreatedBy field value if set, zero value otherwise.
func (o *RBAC) GetDbotCreatedBy() string {
	if o == nil || o.DbotCreatedBy == nil {
		var ret string
		return ret
	}
	return *o.DbotCreatedBy
}

// GetDbotCreatedByOk returns a tuple with the DbotCreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBAC) GetDbotCreatedByOk() (*string, bool) {
	if o == nil || o.DbotCreatedBy == nil {
		return nil, false
	}
	return o.DbotCreatedBy, true
}

// HasDbotCreatedBy returns a boolean if a field has been set.
func (o *RBAC) HasDbotCreatedBy() bool {
	if o != nil && o.DbotCreatedBy != nil {
		return true
	}

	return false
}

// SetDbotCreatedBy gets a reference to the given string and assigns it to the DbotCreatedBy field.
func (o *RBAC) SetDbotCreatedBy(v string) {
	o.DbotCreatedBy = &v
}

// GetHasRole returns the HasRole field value if set, zero value otherwise.
func (o *RBAC) GetHasRole() bool {
	if o == nil || o.HasRole == nil {
		var ret bool
		return ret
	}
	return *o.HasRole
}

// GetHasRoleOk returns a tuple with the HasRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBAC) GetHasRoleOk() (*bool, bool) {
	if o == nil || o.HasRole == nil {
		return nil, false
	}
	return o.HasRole, true
}

// HasHasRole returns a boolean if a field has been set.
func (o *RBAC) HasHasRole() bool {
	if o != nil && o.HasRole != nil {
		return true
	}

	return false
}

// SetHasRole gets a reference to the given bool and assigns it to the HasRole field.
func (o *RBAC) SetHasRole(v bool) {
	o.HasRole = &v
}

// GetPreviousAllRead returns the PreviousAllRead field value if set, zero value otherwise.
func (o *RBAC) GetPreviousAllRead() bool {
	if o == nil || o.PreviousAllRead == nil {
		var ret bool
		return ret
	}
	return *o.PreviousAllRead
}

// GetPreviousAllReadOk returns a tuple with the PreviousAllRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBAC) GetPreviousAllReadOk() (*bool, bool) {
	if o == nil || o.PreviousAllRead == nil {
		return nil, false
	}
	return o.PreviousAllRead, true
}

// HasPreviousAllRead returns a boolean if a field has been set.
func (o *RBAC) HasPreviousAllRead() bool {
	if o != nil && o.PreviousAllRead != nil {
		return true
	}

	return false
}

// SetPreviousAllRead gets a reference to the given bool and assigns it to the PreviousAllRead field.
func (o *RBAC) SetPreviousAllRead(v bool) {
	o.PreviousAllRead = &v
}

// GetPreviousAllReadWrite returns the PreviousAllReadWrite field value if set, zero value otherwise.
func (o *RBAC) GetPreviousAllReadWrite() bool {
	if o == nil || o.PreviousAllReadWrite == nil {
		var ret bool
		return ret
	}
	return *o.PreviousAllReadWrite
}

// GetPreviousAllReadWriteOk returns a tuple with the PreviousAllReadWrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBAC) GetPreviousAllReadWriteOk() (*bool, bool) {
	if o == nil || o.PreviousAllReadWrite == nil {
		return nil, false
	}
	return o.PreviousAllReadWrite, true
}

// HasPreviousAllReadWrite returns a boolean if a field has been set.
func (o *RBAC) HasPreviousAllReadWrite() bool {
	if o != nil && o.PreviousAllReadWrite != nil {
		return true
	}

	return false
}

// SetPreviousAllReadWrite gets a reference to the given bool and assigns it to the PreviousAllReadWrite field.
func (o *RBAC) SetPreviousAllReadWrite(v bool) {
	o.PreviousAllReadWrite = &v
}

// GetPreviousRoles returns the PreviousRoles field value if set, zero value otherwise.
func (o *RBAC) GetPreviousRoles() []string {
	if o == nil || o.PreviousRoles == nil {
		var ret []string
		return ret
	}
	return *o.PreviousRoles
}

// GetPreviousRolesOk returns a tuple with the PreviousRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBAC) GetPreviousRolesOk() (*[]string, bool) {
	if o == nil || o.PreviousRoles == nil {
		return nil, false
	}
	return o.PreviousRoles, true
}

// HasPreviousRoles returns a boolean if a field has been set.
func (o *RBAC) HasPreviousRoles() bool {
	if o != nil && o.PreviousRoles != nil {
		return true
	}

	return false
}

// SetPreviousRoles gets a reference to the given []string and assigns it to the PreviousRoles field.
func (o *RBAC) SetPreviousRoles(v []string) {
	o.PreviousRoles = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *RBAC) GetRoles() []string {
	if o == nil || o.Roles == nil {
		var ret []string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBAC) GetRolesOk() (*[]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *RBAC) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *RBAC) SetRoles(v []string) {
	o.Roles = &v
}

// GetXsoarHasReadOnlyRole returns the XsoarHasReadOnlyRole field value if set, zero value otherwise.
func (o *RBAC) GetXsoarHasReadOnlyRole() bool {
	if o == nil || o.XsoarHasReadOnlyRole == nil {
		var ret bool
		return ret
	}
	return *o.XsoarHasReadOnlyRole
}

// GetXsoarHasReadOnlyRoleOk returns a tuple with the XsoarHasReadOnlyRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBAC) GetXsoarHasReadOnlyRoleOk() (*bool, bool) {
	if o == nil || o.XsoarHasReadOnlyRole == nil {
		return nil, false
	}
	return o.XsoarHasReadOnlyRole, true
}

// HasXsoarHasReadOnlyRole returns a boolean if a field has been set.
func (o *RBAC) HasXsoarHasReadOnlyRole() bool {
	if o != nil && o.XsoarHasReadOnlyRole != nil {
		return true
	}

	return false
}

// SetXsoarHasReadOnlyRole gets a reference to the given bool and assigns it to the XsoarHasReadOnlyRole field.
func (o *RBAC) SetXsoarHasReadOnlyRole(v bool) {
	o.XsoarHasReadOnlyRole = &v
}

// GetXsoarPreviousReadOnlyRoles returns the XsoarPreviousReadOnlyRoles field value if set, zero value otherwise.
func (o *RBAC) GetXsoarPreviousReadOnlyRoles() []string {
	if o == nil || o.XsoarPreviousReadOnlyRoles == nil {
		var ret []string
		return ret
	}
	return *o.XsoarPreviousReadOnlyRoles
}

// GetXsoarPreviousReadOnlyRolesOk returns a tuple with the XsoarPreviousReadOnlyRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBAC) GetXsoarPreviousReadOnlyRolesOk() (*[]string, bool) {
	if o == nil || o.XsoarPreviousReadOnlyRoles == nil {
		return nil, false
	}
	return o.XsoarPreviousReadOnlyRoles, true
}

// HasXsoarPreviousReadOnlyRoles returns a boolean if a field has been set.
func (o *RBAC) HasXsoarPreviousReadOnlyRoles() bool {
	if o != nil && o.XsoarPreviousReadOnlyRoles != nil {
		return true
	}

	return false
}

// SetXsoarPreviousReadOnlyRoles gets a reference to the given []string and assigns it to the XsoarPreviousReadOnlyRoles field.
func (o *RBAC) SetXsoarPreviousReadOnlyRoles(v []string) {
	o.XsoarPreviousReadOnlyRoles = &v
}

// GetXsoarReadOnlyRoles returns the XsoarReadOnlyRoles field value if set, zero value otherwise.
func (o *RBAC) GetXsoarReadOnlyRoles() []string {
	if o == nil || o.XsoarReadOnlyRoles == nil {
		var ret []string
		return ret
	}
	return *o.XsoarReadOnlyRoles
}

// GetXsoarReadOnlyRolesOk returns a tuple with the XsoarReadOnlyRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBAC) GetXsoarReadOnlyRolesOk() (*[]string, bool) {
	if o == nil || o.XsoarReadOnlyRoles == nil {
		return nil, false
	}
	return o.XsoarReadOnlyRoles, true
}

// HasXsoarReadOnlyRoles returns a boolean if a field has been set.
func (o *RBAC) HasXsoarReadOnlyRoles() bool {
	if o != nil && o.XsoarReadOnlyRoles != nil {
		return true
	}

	return false
}

// SetXsoarReadOnlyRoles gets a reference to the given []string and assigns it to the XsoarReadOnlyRoles field.
func (o *RBAC) SetXsoarReadOnlyRoles(v []string) {
	o.XsoarReadOnlyRoles = &v
}

func (o RBAC) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllRead != nil {
		toSerialize["allRead"] = o.AllRead
	}
	if o.AllReadWrite != nil {
		toSerialize["allReadWrite"] = o.AllReadWrite
	}
	if o.DbotCreatedBy != nil {
		toSerialize["dbotCreatedBy"] = o.DbotCreatedBy
	}
	if o.HasRole != nil {
		toSerialize["hasRole"] = o.HasRole
	}
	if o.PreviousAllRead != nil {
		toSerialize["previousAllRead"] = o.PreviousAllRead
	}
	if o.PreviousAllReadWrite != nil {
		toSerialize["previousAllReadWrite"] = o.PreviousAllReadWrite
	}
	if o.PreviousRoles != nil {
		toSerialize["previousRoles"] = o.PreviousRoles
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.XsoarHasReadOnlyRole != nil {
		toSerialize["xsoarHasReadOnlyRole"] = o.XsoarHasReadOnlyRole
	}
	if o.XsoarPreviousReadOnlyRoles != nil {
		toSerialize["xsoarPreviousReadOnlyRoles"] = o.XsoarPreviousReadOnlyRoles
	}
	if o.XsoarReadOnlyRoles != nil {
		toSerialize["xsoarReadOnlyRoles"] = o.XsoarReadOnlyRoles
	}
	return json.Marshal(toSerialize)
}

type NullableRBAC struct {
	value *RBAC
	isSet bool
}

func (v NullableRBAC) Get() *RBAC {
	return v.value
}

func (v *NullableRBAC) Set(val *RBAC) {
	v.value = val
	v.isSet = true
}

func (v NullableRBAC) IsSet() bool {
	return v.isSet
}

func (v *NullableRBAC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRBAC(val *RBAC) *NullableRBAC {
	return &NullableRBAC{value: val, isSet: true}
}

func (v NullableRBAC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRBAC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
