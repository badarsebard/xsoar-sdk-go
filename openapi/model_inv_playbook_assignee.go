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

// InvPlaybookAssignee struct for InvPlaybookAssignee
type InvPlaybookAssignee struct {
	Assignee *string `json:"assignee,omitempty"`
	InTaskID *string `json:"inTaskID,omitempty"`
	InvId    *string `json:"invId,omitempty"`
	Version  *int64  `json:"version,omitempty"`
}

// NewInvPlaybookAssignee instantiates a new InvPlaybookAssignee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvPlaybookAssignee() *InvPlaybookAssignee {
	this := InvPlaybookAssignee{}
	return &this
}

// NewInvPlaybookAssigneeWithDefaults instantiates a new InvPlaybookAssignee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvPlaybookAssigneeWithDefaults() *InvPlaybookAssignee {
	this := InvPlaybookAssignee{}
	return &this
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *InvPlaybookAssignee) GetAssignee() string {
	if o == nil || o.Assignee == nil {
		var ret string
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvPlaybookAssignee) GetAssigneeOk() (*string, bool) {
	if o == nil || o.Assignee == nil {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *InvPlaybookAssignee) HasAssignee() bool {
	if o != nil && o.Assignee != nil {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given string and assigns it to the Assignee field.
func (o *InvPlaybookAssignee) SetAssignee(v string) {
	o.Assignee = &v
}

// GetInTaskID returns the InTaskID field value if set, zero value otherwise.
func (o *InvPlaybookAssignee) GetInTaskID() string {
	if o == nil || o.InTaskID == nil {
		var ret string
		return ret
	}
	return *o.InTaskID
}

// GetInTaskIDOk returns a tuple with the InTaskID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvPlaybookAssignee) GetInTaskIDOk() (*string, bool) {
	if o == nil || o.InTaskID == nil {
		return nil, false
	}
	return o.InTaskID, true
}

// HasInTaskID returns a boolean if a field has been set.
func (o *InvPlaybookAssignee) HasInTaskID() bool {
	if o != nil && o.InTaskID != nil {
		return true
	}

	return false
}

// SetInTaskID gets a reference to the given string and assigns it to the InTaskID field.
func (o *InvPlaybookAssignee) SetInTaskID(v string) {
	o.InTaskID = &v
}

// GetInvId returns the InvId field value if set, zero value otherwise.
func (o *InvPlaybookAssignee) GetInvId() string {
	if o == nil || o.InvId == nil {
		var ret string
		return ret
	}
	return *o.InvId
}

// GetInvIdOk returns a tuple with the InvId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvPlaybookAssignee) GetInvIdOk() (*string, bool) {
	if o == nil || o.InvId == nil {
		return nil, false
	}
	return o.InvId, true
}

// HasInvId returns a boolean if a field has been set.
func (o *InvPlaybookAssignee) HasInvId() bool {
	if o != nil && o.InvId != nil {
		return true
	}

	return false
}

// SetInvId gets a reference to the given string and assigns it to the InvId field.
func (o *InvPlaybookAssignee) SetInvId(v string) {
	o.InvId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *InvPlaybookAssignee) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvPlaybookAssignee) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *InvPlaybookAssignee) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *InvPlaybookAssignee) SetVersion(v int64) {
	o.Version = &v
}

func (o InvPlaybookAssignee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Assignee != nil {
		toSerialize["assignee"] = o.Assignee
	}
	if o.InTaskID != nil {
		toSerialize["inTaskID"] = o.InTaskID
	}
	if o.InvId != nil {
		toSerialize["invId"] = o.InvId
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableInvPlaybookAssignee struct {
	value *InvPlaybookAssignee
	isSet bool
}

func (v NullableInvPlaybookAssignee) Get() *InvPlaybookAssignee {
	return v.value
}

func (v *NullableInvPlaybookAssignee) Set(val *InvPlaybookAssignee) {
	v.value = val
	v.isSet = true
}

func (v NullableInvPlaybookAssignee) IsSet() bool {
	return v.isSet
}

func (v *NullableInvPlaybookAssignee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvPlaybookAssignee(val *InvPlaybookAssignee) *NullableInvPlaybookAssignee {
	return &NullableInvPlaybookAssignee{value: val, isSet: true}
}

func (v NullableInvPlaybookAssignee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvPlaybookAssignee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
