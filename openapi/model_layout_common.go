/*
Cortex XSOAR API

This is the public REST API to integrate with the Cortex XSOAR server. HTTP request can be sent using any HTTP-client.  For an example dedicated client take a look at: https://github.com/demisto/demisto-py.  Requests must include API-key that can be generated in the Cortex XSOAR web client under 'Settings' -> 'Integrations' -> 'API keys'   Optimistic Locking and Versioning\\:  When using Cortex XSOAR REST API, you will need to make sure to work on the latest version of the item (incident, entry, etc.), otherwise, you will get a DB version error (which not allow you to override a newer item). In addition, you can pass 'version\\: -1' to force data override (make sure that other users data might be lost).  Assume that Alice and Bob both read the same data from Cortex XSOAR server, then they both changed the data, and then both tried to write the new versions back to the server. Whose changes should be saved? Alice’s? Bob’s? To solve this, each data item in Cortex XSOAR has a numeric incremental version. If Alice saved an item with version 4 and Bob trying to save the same item with version 3, Cortex XSOAR will rollback Bob request and returns a DB version conflict error. Bob will need to get the latest item and work on it so Alice work will not get lost.  Example request using 'curl'\\:  ``` curl 'https://hostname:443/incidents/search' -H 'content-type: application/json' -H 'accept: application/json' -H 'Authorization: <API Key goes here>' --data-binary '{\"filter\":{\"query\":\"-status:closed -category:job\",\"period\":{\"by\":\"day\",\"fromValue\":7}}}' --compressed ```

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// LayoutCommon struct for LayoutCommon
type LayoutCommon struct {
	CommitMessage                     *string              `json:"commitMessage,omitempty"`
	FromServerVersion                 *Version             `json:"fromServerVersion,omitempty"`
	Highlight                         *map[string][]string `json:"highlight,omitempty"`
	Id                                *string              `json:"id,omitempty"`
	ItemVersion                       *Version             `json:"itemVersion,omitempty"`
	Kind                              *string              `json:"kind,omitempty"`
	Locked                            *bool                `json:"locked,omitempty"`
	Modified                          *time.Time           `json:"modified,omitempty"`
	Name                              *string              `json:"name,omitempty"`
	NumericId                         *int64               `json:"numericId,omitempty"`
	PackID                            *string              `json:"packID,omitempty"`
	PackPropagationLabels             *[]string            `json:"packPropagationLabels,omitempty"`
	PrevKind                          *string              `json:"prevKind,omitempty"`
	PrevTypeId                        *string              `json:"prevTypeId,omitempty"`
	PrimaryTerm                       *int64               `json:"primaryTerm,omitempty"`
	PropagationLabels                 *[]string            `json:"propagationLabels,omitempty"`
	SequenceNumber                    *int64               `json:"sequenceNumber,omitempty"`
	ShouldCommit                      *bool                `json:"shouldCommit,omitempty"`
	SortValues                        *[]string            `json:"sortValues,omitempty"`
	System                            *bool                `json:"system,omitempty"`
	ToServerVersion                   *Version             `json:"toServerVersion,omitempty"`
	TypeId                            *string              `json:"typeId,omitempty"`
	VcShouldIgnore                    *bool                `json:"vcShouldIgnore,omitempty"`
	VcShouldKeepItemLegacyProdMachine *bool                `json:"vcShouldKeepItemLegacyProdMachine,omitempty"`
	Version                           *int64               `json:"version,omitempty"`
}

// NewLayoutCommon instantiates a new LayoutCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLayoutCommon() *LayoutCommon {
	this := LayoutCommon{}
	return &this
}

// NewLayoutCommonWithDefaults instantiates a new LayoutCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLayoutCommonWithDefaults() *LayoutCommon {
	this := LayoutCommon{}
	return &this
}

// GetCommitMessage returns the CommitMessage field value if set, zero value otherwise.
func (o *LayoutCommon) GetCommitMessage() string {
	if o == nil || o.CommitMessage == nil {
		var ret string
		return ret
	}
	return *o.CommitMessage
}

// GetCommitMessageOk returns a tuple with the CommitMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetCommitMessageOk() (*string, bool) {
	if o == nil || o.CommitMessage == nil {
		return nil, false
	}
	return o.CommitMessage, true
}

// HasCommitMessage returns a boolean if a field has been set.
func (o *LayoutCommon) HasCommitMessage() bool {
	if o != nil && o.CommitMessage != nil {
		return true
	}

	return false
}

// SetCommitMessage gets a reference to the given string and assigns it to the CommitMessage field.
func (o *LayoutCommon) SetCommitMessage(v string) {
	o.CommitMessage = &v
}

// GetFromServerVersion returns the FromServerVersion field value if set, zero value otherwise.
func (o *LayoutCommon) GetFromServerVersion() Version {
	if o == nil || o.FromServerVersion == nil {
		var ret Version
		return ret
	}
	return *o.FromServerVersion
}

// GetFromServerVersionOk returns a tuple with the FromServerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetFromServerVersionOk() (*Version, bool) {
	if o == nil || o.FromServerVersion == nil {
		return nil, false
	}
	return o.FromServerVersion, true
}

// HasFromServerVersion returns a boolean if a field has been set.
func (o *LayoutCommon) HasFromServerVersion() bool {
	if o != nil && o.FromServerVersion != nil {
		return true
	}

	return false
}

// SetFromServerVersion gets a reference to the given Version and assigns it to the FromServerVersion field.
func (o *LayoutCommon) SetFromServerVersion(v Version) {
	o.FromServerVersion = &v
}

// GetHighlight returns the Highlight field value if set, zero value otherwise.
func (o *LayoutCommon) GetHighlight() map[string][]string {
	if o == nil || o.Highlight == nil {
		var ret map[string][]string
		return ret
	}
	return *o.Highlight
}

// GetHighlightOk returns a tuple with the Highlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetHighlightOk() (*map[string][]string, bool) {
	if o == nil || o.Highlight == nil {
		return nil, false
	}
	return o.Highlight, true
}

// HasHighlight returns a boolean if a field has been set.
func (o *LayoutCommon) HasHighlight() bool {
	if o != nil && o.Highlight != nil {
		return true
	}

	return false
}

// SetHighlight gets a reference to the given map[string][]string and assigns it to the Highlight field.
func (o *LayoutCommon) SetHighlight(v map[string][]string) {
	o.Highlight = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LayoutCommon) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LayoutCommon) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LayoutCommon) SetId(v string) {
	o.Id = &v
}

// GetItemVersion returns the ItemVersion field value if set, zero value otherwise.
func (o *LayoutCommon) GetItemVersion() Version {
	if o == nil || o.ItemVersion == nil {
		var ret Version
		return ret
	}
	return *o.ItemVersion
}

// GetItemVersionOk returns a tuple with the ItemVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetItemVersionOk() (*Version, bool) {
	if o == nil || o.ItemVersion == nil {
		return nil, false
	}
	return o.ItemVersion, true
}

// HasItemVersion returns a boolean if a field has been set.
func (o *LayoutCommon) HasItemVersion() bool {
	if o != nil && o.ItemVersion != nil {
		return true
	}

	return false
}

// SetItemVersion gets a reference to the given Version and assigns it to the ItemVersion field.
func (o *LayoutCommon) SetItemVersion(v Version) {
	o.ItemVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *LayoutCommon) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *LayoutCommon) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *LayoutCommon) SetKind(v string) {
	o.Kind = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *LayoutCommon) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *LayoutCommon) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *LayoutCommon) SetLocked(v bool) {
	o.Locked = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *LayoutCommon) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *LayoutCommon) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *LayoutCommon) SetModified(v time.Time) {
	o.Modified = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LayoutCommon) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LayoutCommon) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LayoutCommon) SetName(v string) {
	o.Name = &v
}

// GetNumericId returns the NumericId field value if set, zero value otherwise.
func (o *LayoutCommon) GetNumericId() int64 {
	if o == nil || o.NumericId == nil {
		var ret int64
		return ret
	}
	return *o.NumericId
}

// GetNumericIdOk returns a tuple with the NumericId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetNumericIdOk() (*int64, bool) {
	if o == nil || o.NumericId == nil {
		return nil, false
	}
	return o.NumericId, true
}

// HasNumericId returns a boolean if a field has been set.
func (o *LayoutCommon) HasNumericId() bool {
	if o != nil && o.NumericId != nil {
		return true
	}

	return false
}

// SetNumericId gets a reference to the given int64 and assigns it to the NumericId field.
func (o *LayoutCommon) SetNumericId(v int64) {
	o.NumericId = &v
}

// GetPackID returns the PackID field value if set, zero value otherwise.
func (o *LayoutCommon) GetPackID() string {
	if o == nil || o.PackID == nil {
		var ret string
		return ret
	}
	return *o.PackID
}

// GetPackIDOk returns a tuple with the PackID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetPackIDOk() (*string, bool) {
	if o == nil || o.PackID == nil {
		return nil, false
	}
	return o.PackID, true
}

// HasPackID returns a boolean if a field has been set.
func (o *LayoutCommon) HasPackID() bool {
	if o != nil && o.PackID != nil {
		return true
	}

	return false
}

// SetPackID gets a reference to the given string and assigns it to the PackID field.
func (o *LayoutCommon) SetPackID(v string) {
	o.PackID = &v
}

// GetPackPropagationLabels returns the PackPropagationLabels field value if set, zero value otherwise.
func (o *LayoutCommon) GetPackPropagationLabels() []string {
	if o == nil || o.PackPropagationLabels == nil {
		var ret []string
		return ret
	}
	return *o.PackPropagationLabels
}

// GetPackPropagationLabelsOk returns a tuple with the PackPropagationLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetPackPropagationLabelsOk() (*[]string, bool) {
	if o == nil || o.PackPropagationLabels == nil {
		return nil, false
	}
	return o.PackPropagationLabels, true
}

// HasPackPropagationLabels returns a boolean if a field has been set.
func (o *LayoutCommon) HasPackPropagationLabels() bool {
	if o != nil && o.PackPropagationLabels != nil {
		return true
	}

	return false
}

// SetPackPropagationLabels gets a reference to the given []string and assigns it to the PackPropagationLabels field.
func (o *LayoutCommon) SetPackPropagationLabels(v []string) {
	o.PackPropagationLabels = &v
}

// GetPrevKind returns the PrevKind field value if set, zero value otherwise.
func (o *LayoutCommon) GetPrevKind() string {
	if o == nil || o.PrevKind == nil {
		var ret string
		return ret
	}
	return *o.PrevKind
}

// GetPrevKindOk returns a tuple with the PrevKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetPrevKindOk() (*string, bool) {
	if o == nil || o.PrevKind == nil {
		return nil, false
	}
	return o.PrevKind, true
}

// HasPrevKind returns a boolean if a field has been set.
func (o *LayoutCommon) HasPrevKind() bool {
	if o != nil && o.PrevKind != nil {
		return true
	}

	return false
}

// SetPrevKind gets a reference to the given string and assigns it to the PrevKind field.
func (o *LayoutCommon) SetPrevKind(v string) {
	o.PrevKind = &v
}

// GetPrevTypeId returns the PrevTypeId field value if set, zero value otherwise.
func (o *LayoutCommon) GetPrevTypeId() string {
	if o == nil || o.PrevTypeId == nil {
		var ret string
		return ret
	}
	return *o.PrevTypeId
}

// GetPrevTypeIdOk returns a tuple with the PrevTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetPrevTypeIdOk() (*string, bool) {
	if o == nil || o.PrevTypeId == nil {
		return nil, false
	}
	return o.PrevTypeId, true
}

// HasPrevTypeId returns a boolean if a field has been set.
func (o *LayoutCommon) HasPrevTypeId() bool {
	if o != nil && o.PrevTypeId != nil {
		return true
	}

	return false
}

// SetPrevTypeId gets a reference to the given string and assigns it to the PrevTypeId field.
func (o *LayoutCommon) SetPrevTypeId(v string) {
	o.PrevTypeId = &v
}

// GetPrimaryTerm returns the PrimaryTerm field value if set, zero value otherwise.
func (o *LayoutCommon) GetPrimaryTerm() int64 {
	if o == nil || o.PrimaryTerm == nil {
		var ret int64
		return ret
	}
	return *o.PrimaryTerm
}

// GetPrimaryTermOk returns a tuple with the PrimaryTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetPrimaryTermOk() (*int64, bool) {
	if o == nil || o.PrimaryTerm == nil {
		return nil, false
	}
	return o.PrimaryTerm, true
}

// HasPrimaryTerm returns a boolean if a field has been set.
func (o *LayoutCommon) HasPrimaryTerm() bool {
	if o != nil && o.PrimaryTerm != nil {
		return true
	}

	return false
}

// SetPrimaryTerm gets a reference to the given int64 and assigns it to the PrimaryTerm field.
func (o *LayoutCommon) SetPrimaryTerm(v int64) {
	o.PrimaryTerm = &v
}

// GetPropagationLabels returns the PropagationLabels field value if set, zero value otherwise.
func (o *LayoutCommon) GetPropagationLabels() []string {
	if o == nil || o.PropagationLabels == nil {
		var ret []string
		return ret
	}
	return *o.PropagationLabels
}

// GetPropagationLabelsOk returns a tuple with the PropagationLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetPropagationLabelsOk() (*[]string, bool) {
	if o == nil || o.PropagationLabels == nil {
		return nil, false
	}
	return o.PropagationLabels, true
}

// HasPropagationLabels returns a boolean if a field has been set.
func (o *LayoutCommon) HasPropagationLabels() bool {
	if o != nil && o.PropagationLabels != nil {
		return true
	}

	return false
}

// SetPropagationLabels gets a reference to the given []string and assigns it to the PropagationLabels field.
func (o *LayoutCommon) SetPropagationLabels(v []string) {
	o.PropagationLabels = &v
}

// GetSequenceNumber returns the SequenceNumber field value if set, zero value otherwise.
func (o *LayoutCommon) GetSequenceNumber() int64 {
	if o == nil || o.SequenceNumber == nil {
		var ret int64
		return ret
	}
	return *o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetSequenceNumberOk() (*int64, bool) {
	if o == nil || o.SequenceNumber == nil {
		return nil, false
	}
	return o.SequenceNumber, true
}

// HasSequenceNumber returns a boolean if a field has been set.
func (o *LayoutCommon) HasSequenceNumber() bool {
	if o != nil && o.SequenceNumber != nil {
		return true
	}

	return false
}

// SetSequenceNumber gets a reference to the given int64 and assigns it to the SequenceNumber field.
func (o *LayoutCommon) SetSequenceNumber(v int64) {
	o.SequenceNumber = &v
}

// GetShouldCommit returns the ShouldCommit field value if set, zero value otherwise.
func (o *LayoutCommon) GetShouldCommit() bool {
	if o == nil || o.ShouldCommit == nil {
		var ret bool
		return ret
	}
	return *o.ShouldCommit
}

// GetShouldCommitOk returns a tuple with the ShouldCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetShouldCommitOk() (*bool, bool) {
	if o == nil || o.ShouldCommit == nil {
		return nil, false
	}
	return o.ShouldCommit, true
}

// HasShouldCommit returns a boolean if a field has been set.
func (o *LayoutCommon) HasShouldCommit() bool {
	if o != nil && o.ShouldCommit != nil {
		return true
	}

	return false
}

// SetShouldCommit gets a reference to the given bool and assigns it to the ShouldCommit field.
func (o *LayoutCommon) SetShouldCommit(v bool) {
	o.ShouldCommit = &v
}

// GetSortValues returns the SortValues field value if set, zero value otherwise.
func (o *LayoutCommon) GetSortValues() []string {
	if o == nil || o.SortValues == nil {
		var ret []string
		return ret
	}
	return *o.SortValues
}

// GetSortValuesOk returns a tuple with the SortValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetSortValuesOk() (*[]string, bool) {
	if o == nil || o.SortValues == nil {
		return nil, false
	}
	return o.SortValues, true
}

// HasSortValues returns a boolean if a field has been set.
func (o *LayoutCommon) HasSortValues() bool {
	if o != nil && o.SortValues != nil {
		return true
	}

	return false
}

// SetSortValues gets a reference to the given []string and assigns it to the SortValues field.
func (o *LayoutCommon) SetSortValues(v []string) {
	o.SortValues = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *LayoutCommon) GetSystem() bool {
	if o == nil || o.System == nil {
		var ret bool
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetSystemOk() (*bool, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *LayoutCommon) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given bool and assigns it to the System field.
func (o *LayoutCommon) SetSystem(v bool) {
	o.System = &v
}

// GetToServerVersion returns the ToServerVersion field value if set, zero value otherwise.
func (o *LayoutCommon) GetToServerVersion() Version {
	if o == nil || o.ToServerVersion == nil {
		var ret Version
		return ret
	}
	return *o.ToServerVersion
}

// GetToServerVersionOk returns a tuple with the ToServerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetToServerVersionOk() (*Version, bool) {
	if o == nil || o.ToServerVersion == nil {
		return nil, false
	}
	return o.ToServerVersion, true
}

// HasToServerVersion returns a boolean if a field has been set.
func (o *LayoutCommon) HasToServerVersion() bool {
	if o != nil && o.ToServerVersion != nil {
		return true
	}

	return false
}

// SetToServerVersion gets a reference to the given Version and assigns it to the ToServerVersion field.
func (o *LayoutCommon) SetToServerVersion(v Version) {
	o.ToServerVersion = &v
}

// GetTypeId returns the TypeId field value if set, zero value otherwise.
func (o *LayoutCommon) GetTypeId() string {
	if o == nil || o.TypeId == nil {
		var ret string
		return ret
	}
	return *o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetTypeIdOk() (*string, bool) {
	if o == nil || o.TypeId == nil {
		return nil, false
	}
	return o.TypeId, true
}

// HasTypeId returns a boolean if a field has been set.
func (o *LayoutCommon) HasTypeId() bool {
	if o != nil && o.TypeId != nil {
		return true
	}

	return false
}

// SetTypeId gets a reference to the given string and assigns it to the TypeId field.
func (o *LayoutCommon) SetTypeId(v string) {
	o.TypeId = &v
}

// GetVcShouldIgnore returns the VcShouldIgnore field value if set, zero value otherwise.
func (o *LayoutCommon) GetVcShouldIgnore() bool {
	if o == nil || o.VcShouldIgnore == nil {
		var ret bool
		return ret
	}
	return *o.VcShouldIgnore
}

// GetVcShouldIgnoreOk returns a tuple with the VcShouldIgnore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetVcShouldIgnoreOk() (*bool, bool) {
	if o == nil || o.VcShouldIgnore == nil {
		return nil, false
	}
	return o.VcShouldIgnore, true
}

// HasVcShouldIgnore returns a boolean if a field has been set.
func (o *LayoutCommon) HasVcShouldIgnore() bool {
	if o != nil && o.VcShouldIgnore != nil {
		return true
	}

	return false
}

// SetVcShouldIgnore gets a reference to the given bool and assigns it to the VcShouldIgnore field.
func (o *LayoutCommon) SetVcShouldIgnore(v bool) {
	o.VcShouldIgnore = &v
}

// GetVcShouldKeepItemLegacyProdMachine returns the VcShouldKeepItemLegacyProdMachine field value if set, zero value otherwise.
func (o *LayoutCommon) GetVcShouldKeepItemLegacyProdMachine() bool {
	if o == nil || o.VcShouldKeepItemLegacyProdMachine == nil {
		var ret bool
		return ret
	}
	return *o.VcShouldKeepItemLegacyProdMachine
}

// GetVcShouldKeepItemLegacyProdMachineOk returns a tuple with the VcShouldKeepItemLegacyProdMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetVcShouldKeepItemLegacyProdMachineOk() (*bool, bool) {
	if o == nil || o.VcShouldKeepItemLegacyProdMachine == nil {
		return nil, false
	}
	return o.VcShouldKeepItemLegacyProdMachine, true
}

// HasVcShouldKeepItemLegacyProdMachine returns a boolean if a field has been set.
func (o *LayoutCommon) HasVcShouldKeepItemLegacyProdMachine() bool {
	if o != nil && o.VcShouldKeepItemLegacyProdMachine != nil {
		return true
	}

	return false
}

// SetVcShouldKeepItemLegacyProdMachine gets a reference to the given bool and assigns it to the VcShouldKeepItemLegacyProdMachine field.
func (o *LayoutCommon) SetVcShouldKeepItemLegacyProdMachine(v bool) {
	o.VcShouldKeepItemLegacyProdMachine = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LayoutCommon) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutCommon) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LayoutCommon) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *LayoutCommon) SetVersion(v int64) {
	o.Version = &v
}

func (o LayoutCommon) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CommitMessage != nil {
		toSerialize["commitMessage"] = o.CommitMessage
	}
	if o.FromServerVersion != nil {
		toSerialize["fromServerVersion"] = o.FromServerVersion
	}
	if o.Highlight != nil {
		toSerialize["highlight"] = o.Highlight
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ItemVersion != nil {
		toSerialize["itemVersion"] = o.ItemVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NumericId != nil {
		toSerialize["numericId"] = o.NumericId
	}
	if o.PackID != nil {
		toSerialize["packID"] = o.PackID
	}
	if o.PackPropagationLabels != nil {
		toSerialize["packPropagationLabels"] = o.PackPropagationLabels
	}
	if o.PrevKind != nil {
		toSerialize["prevKind"] = o.PrevKind
	}
	if o.PrevTypeId != nil {
		toSerialize["prevTypeId"] = o.PrevTypeId
	}
	if o.PrimaryTerm != nil {
		toSerialize["primaryTerm"] = o.PrimaryTerm
	}
	if o.PropagationLabels != nil {
		toSerialize["propagationLabels"] = o.PropagationLabels
	}
	if o.SequenceNumber != nil {
		toSerialize["sequenceNumber"] = o.SequenceNumber
	}
	if o.ShouldCommit != nil {
		toSerialize["shouldCommit"] = o.ShouldCommit
	}
	if o.SortValues != nil {
		toSerialize["sortValues"] = o.SortValues
	}
	if o.System != nil {
		toSerialize["system"] = o.System
	}
	if o.ToServerVersion != nil {
		toSerialize["toServerVersion"] = o.ToServerVersion
	}
	if o.TypeId != nil {
		toSerialize["typeId"] = o.TypeId
	}
	if o.VcShouldIgnore != nil {
		toSerialize["vcShouldIgnore"] = o.VcShouldIgnore
	}
	if o.VcShouldKeepItemLegacyProdMachine != nil {
		toSerialize["vcShouldKeepItemLegacyProdMachine"] = o.VcShouldKeepItemLegacyProdMachine
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableLayoutCommon struct {
	value *LayoutCommon
	isSet bool
}

func (v NullableLayoutCommon) Get() *LayoutCommon {
	return v.value
}

func (v *NullableLayoutCommon) Set(val *LayoutCommon) {
	v.value = val
	v.isSet = true
}

func (v NullableLayoutCommon) IsSet() bool {
	return v.isSet
}

func (v *NullableLayoutCommon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLayoutCommon(val *LayoutCommon) *NullableLayoutCommon {
	return &NullableLayoutCommon{value: val, isSet: true}
}

func (v NullableLayoutCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLayoutCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
