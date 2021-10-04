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

// InvPlaybookTaskData struct for InvPlaybookTaskData
type InvPlaybookTaskData struct {
	AddAfter            *bool                  `json:"addAfter,omitempty"`
	AutomationScript    *string                `json:"automationScript,omitempty"`
	Description         *string                `json:"description,omitempty"`
	Loop                *TaskLoop              `json:"loop,omitempty"`
	Name                *string                `json:"name,omitempty"`
	NeighborInvPBTaskId *string                `json:"neighborInvPBTaskId,omitempty"`
	PlaybookId          *string                `json:"playbookId,omitempty"`
	ScriptArguments     *map[string]AdvanceArg `json:"scriptArguments,omitempty"`
	SeparateContext     *bool                  `json:"separateContext,omitempty"`
	Tags                *[]string              `json:"tags,omitempty"`
	// TaskType is the Task in the playbook context as a node
	Type *string `json:"type,omitempty"`
}

// NewInvPlaybookTaskData instantiates a new InvPlaybookTaskData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvPlaybookTaskData() *InvPlaybookTaskData {
	this := InvPlaybookTaskData{}
	return &this
}

// NewInvPlaybookTaskDataWithDefaults instantiates a new InvPlaybookTaskData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvPlaybookTaskDataWithDefaults() *InvPlaybookTaskData {
	this := InvPlaybookTaskData{}
	return &this
}

// GetAddAfter returns the AddAfter field value if set, zero value otherwise.
func (o *InvPlaybookTaskData) GetAddAfter() bool {
	if o == nil || o.AddAfter == nil {
		var ret bool
		return ret
	}
	return *o.AddAfter
}

// GetAddAfterOk returns a tuple with the AddAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvPlaybookTaskData) GetAddAfterOk() (*bool, bool) {
	if o == nil || o.AddAfter == nil {
		return nil, false
	}
	return o.AddAfter, true
}

// HasAddAfter returns a boolean if a field has been set.
func (o *InvPlaybookTaskData) HasAddAfter() bool {
	if o != nil && o.AddAfter != nil {
		return true
	}

	return false
}

// SetAddAfter gets a reference to the given bool and assigns it to the AddAfter field.
func (o *InvPlaybookTaskData) SetAddAfter(v bool) {
	o.AddAfter = &v
}

// GetAutomationScript returns the AutomationScript field value if set, zero value otherwise.
func (o *InvPlaybookTaskData) GetAutomationScript() string {
	if o == nil || o.AutomationScript == nil {
		var ret string
		return ret
	}
	return *o.AutomationScript
}

// GetAutomationScriptOk returns a tuple with the AutomationScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvPlaybookTaskData) GetAutomationScriptOk() (*string, bool) {
	if o == nil || o.AutomationScript == nil {
		return nil, false
	}
	return o.AutomationScript, true
}

// HasAutomationScript returns a boolean if a field has been set.
func (o *InvPlaybookTaskData) HasAutomationScript() bool {
	if o != nil && o.AutomationScript != nil {
		return true
	}

	return false
}

// SetAutomationScript gets a reference to the given string and assigns it to the AutomationScript field.
func (o *InvPlaybookTaskData) SetAutomationScript(v string) {
	o.AutomationScript = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InvPlaybookTaskData) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvPlaybookTaskData) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InvPlaybookTaskData) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InvPlaybookTaskData) SetDescription(v string) {
	o.Description = &v
}

// GetLoop returns the Loop field value if set, zero value otherwise.
func (o *InvPlaybookTaskData) GetLoop() TaskLoop {
	if o == nil || o.Loop == nil {
		var ret TaskLoop
		return ret
	}
	return *o.Loop
}

// GetLoopOk returns a tuple with the Loop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvPlaybookTaskData) GetLoopOk() (*TaskLoop, bool) {
	if o == nil || o.Loop == nil {
		return nil, false
	}
	return o.Loop, true
}

// HasLoop returns a boolean if a field has been set.
func (o *InvPlaybookTaskData) HasLoop() bool {
	if o != nil && o.Loop != nil {
		return true
	}

	return false
}

// SetLoop gets a reference to the given TaskLoop and assigns it to the Loop field.
func (o *InvPlaybookTaskData) SetLoop(v TaskLoop) {
	o.Loop = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InvPlaybookTaskData) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvPlaybookTaskData) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InvPlaybookTaskData) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InvPlaybookTaskData) SetName(v string) {
	o.Name = &v
}

// GetNeighborInvPBTaskId returns the NeighborInvPBTaskId field value if set, zero value otherwise.
func (o *InvPlaybookTaskData) GetNeighborInvPBTaskId() string {
	if o == nil || o.NeighborInvPBTaskId == nil {
		var ret string
		return ret
	}
	return *o.NeighborInvPBTaskId
}

// GetNeighborInvPBTaskIdOk returns a tuple with the NeighborInvPBTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvPlaybookTaskData) GetNeighborInvPBTaskIdOk() (*string, bool) {
	if o == nil || o.NeighborInvPBTaskId == nil {
		return nil, false
	}
	return o.NeighborInvPBTaskId, true
}

// HasNeighborInvPBTaskId returns a boolean if a field has been set.
func (o *InvPlaybookTaskData) HasNeighborInvPBTaskId() bool {
	if o != nil && o.NeighborInvPBTaskId != nil {
		return true
	}

	return false
}

// SetNeighborInvPBTaskId gets a reference to the given string and assigns it to the NeighborInvPBTaskId field.
func (o *InvPlaybookTaskData) SetNeighborInvPBTaskId(v string) {
	o.NeighborInvPBTaskId = &v
}

// GetPlaybookId returns the PlaybookId field value if set, zero value otherwise.
func (o *InvPlaybookTaskData) GetPlaybookId() string {
	if o == nil || o.PlaybookId == nil {
		var ret string
		return ret
	}
	return *o.PlaybookId
}

// GetPlaybookIdOk returns a tuple with the PlaybookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvPlaybookTaskData) GetPlaybookIdOk() (*string, bool) {
	if o == nil || o.PlaybookId == nil {
		return nil, false
	}
	return o.PlaybookId, true
}

// HasPlaybookId returns a boolean if a field has been set.
func (o *InvPlaybookTaskData) HasPlaybookId() bool {
	if o != nil && o.PlaybookId != nil {
		return true
	}

	return false
}

// SetPlaybookId gets a reference to the given string and assigns it to the PlaybookId field.
func (o *InvPlaybookTaskData) SetPlaybookId(v string) {
	o.PlaybookId = &v
}

// GetScriptArguments returns the ScriptArguments field value if set, zero value otherwise.
func (o *InvPlaybookTaskData) GetScriptArguments() map[string]AdvanceArg {
	if o == nil || o.ScriptArguments == nil {
		var ret map[string]AdvanceArg
		return ret
	}
	return *o.ScriptArguments
}

// GetScriptArgumentsOk returns a tuple with the ScriptArguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvPlaybookTaskData) GetScriptArgumentsOk() (*map[string]AdvanceArg, bool) {
	if o == nil || o.ScriptArguments == nil {
		return nil, false
	}
	return o.ScriptArguments, true
}

// HasScriptArguments returns a boolean if a field has been set.
func (o *InvPlaybookTaskData) HasScriptArguments() bool {
	if o != nil && o.ScriptArguments != nil {
		return true
	}

	return false
}

// SetScriptArguments gets a reference to the given map[string]AdvanceArg and assigns it to the ScriptArguments field.
func (o *InvPlaybookTaskData) SetScriptArguments(v map[string]AdvanceArg) {
	o.ScriptArguments = &v
}

// GetSeparateContext returns the SeparateContext field value if set, zero value otherwise.
func (o *InvPlaybookTaskData) GetSeparateContext() bool {
	if o == nil || o.SeparateContext == nil {
		var ret bool
		return ret
	}
	return *o.SeparateContext
}

// GetSeparateContextOk returns a tuple with the SeparateContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvPlaybookTaskData) GetSeparateContextOk() (*bool, bool) {
	if o == nil || o.SeparateContext == nil {
		return nil, false
	}
	return o.SeparateContext, true
}

// HasSeparateContext returns a boolean if a field has been set.
func (o *InvPlaybookTaskData) HasSeparateContext() bool {
	if o != nil && o.SeparateContext != nil {
		return true
	}

	return false
}

// SetSeparateContext gets a reference to the given bool and assigns it to the SeparateContext field.
func (o *InvPlaybookTaskData) SetSeparateContext(v bool) {
	o.SeparateContext = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InvPlaybookTaskData) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvPlaybookTaskData) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InvPlaybookTaskData) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InvPlaybookTaskData) SetTags(v []string) {
	o.Tags = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InvPlaybookTaskData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvPlaybookTaskData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InvPlaybookTaskData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InvPlaybookTaskData) SetType(v string) {
	o.Type = &v
}

func (o InvPlaybookTaskData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddAfter != nil {
		toSerialize["addAfter"] = o.AddAfter
	}
	if o.AutomationScript != nil {
		toSerialize["automationScript"] = o.AutomationScript
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Loop != nil {
		toSerialize["loop"] = o.Loop
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NeighborInvPBTaskId != nil {
		toSerialize["neighborInvPBTaskId"] = o.NeighborInvPBTaskId
	}
	if o.PlaybookId != nil {
		toSerialize["playbookId"] = o.PlaybookId
	}
	if o.ScriptArguments != nil {
		toSerialize["scriptArguments"] = o.ScriptArguments
	}
	if o.SeparateContext != nil {
		toSerialize["separateContext"] = o.SeparateContext
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableInvPlaybookTaskData struct {
	value *InvPlaybookTaskData
	isSet bool
}

func (v NullableInvPlaybookTaskData) Get() *InvPlaybookTaskData {
	return v.value
}

func (v *NullableInvPlaybookTaskData) Set(val *InvPlaybookTaskData) {
	v.value = val
	v.isSet = true
}

func (v NullableInvPlaybookTaskData) IsSet() bool {
	return v.isSet
}

func (v *NullableInvPlaybookTaskData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvPlaybookTaskData(val *InvPlaybookTaskData) *NullableInvPlaybookTaskData {
	return &NullableInvPlaybookTaskData{value: val, isSet: true}
}

func (v NullableInvPlaybookTaskData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvPlaybookTaskData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
