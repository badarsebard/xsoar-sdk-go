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

// CreateUpdateRequest Request to create new or update a classifier
type CreateUpdateRequest struct {
	DefaultIncidentType *string                 `json:"defaultIncidentType,omitempty"`
	KeyTypeMap          *map[string]interface{} `json:"keyTypeMap,omitempty"`
	Name                *string                 `json:"name,omitempty"`
	Transformer         *map[string]interface{} `json:"transformer,omitempty"`
	Mapping             *map[string]interface{} `json:"mapping,omitempty"`
	Type                *string                 `json:"type,omitempty"`
	Id                  *string                 `json:"id,omitempty"`
	PropagationLabels   *[]string               `json:"propagationLabels,omitempty"`
}

// NewCreateUpdateRequest instantiates a new CreateUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUpdateRequest() *CreateUpdateRequest {
	this := CreateUpdateRequest{}
	return &this
}

// NewCreateUpdateRequestWithDefaults instantiates a new CreateUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUpdateRequestWithDefaults() *CreateUpdateRequest {
	this := CreateUpdateRequest{}
	return &this
}

// GetDefaultIncidentType returns the DefaultIncidentType field value if set, zero value otherwise.
func (o *CreateUpdateRequest) GetDefaultIncidentType() string {
	if o == nil || o.DefaultIncidentType == nil {
		var ret string
		return ret
	}
	return *o.DefaultIncidentType
}

// GetDefaultIncidentTypeOk returns a tuple with the DefaultIncidentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateRequest) GetDefaultIncidentTypeOk() (*string, bool) {
	if o == nil || o.DefaultIncidentType == nil {
		return nil, false
	}
	return o.DefaultIncidentType, true
}

// HasDefaultIncidentType returns a boolean if a field has been set.
func (o *CreateUpdateRequest) HasDefaultIncidentType() bool {
	if o != nil && o.DefaultIncidentType != nil {
		return true
	}

	return false
}

// SetDefaultIncidentType gets a reference to the given string and assigns it to the DefaultIncidentType field.
func (o *CreateUpdateRequest) SetDefaultIncidentType(v string) {
	o.DefaultIncidentType = &v
}

// GetKeyTypeMap returns the KeyTypeMap field value if set, zero value otherwise.
func (o *CreateUpdateRequest) GetKeyTypeMap() map[string]interface{} {
	if o == nil || o.KeyTypeMap == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.KeyTypeMap
}

// GetKeyTypeMapOk returns a tuple with the KeyTypeMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateRequest) GetKeyTypeMapOk() (*map[string]interface{}, bool) {
	if o == nil || o.KeyTypeMap == nil {
		return nil, false
	}
	return o.KeyTypeMap, true
}

// HasKeyTypeMap returns a boolean if a field has been set.
func (o *CreateUpdateRequest) HasKeyTypeMap() bool {
	if o != nil && o.KeyTypeMap != nil {
		return true
	}

	return false
}

// SetKeyTypeMap gets a reference to the given map[string]interface{} and assigns it to the KeyTypeMap field.
func (o *CreateUpdateRequest) SetKeyTypeMap(v map[string]interface{}) {
	o.KeyTypeMap = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateUpdateRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateUpdateRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateUpdateRequest) SetName(v string) {
	o.Name = &v
}

// GetTransformer returns the Transformer field value if set, zero value otherwise.
func (o *CreateUpdateRequest) GetTransformer() map[string]interface{} {
	if o == nil || o.Transformer == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Transformer
}

// GetTransformerOk returns a tuple with the Transformer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateRequest) GetTransformerOk() (*map[string]interface{}, bool) {
	if o == nil || o.Transformer == nil {
		return nil, false
	}
	return o.Transformer, true
}

// HasTransformer returns a boolean if a field has been set.
func (o *CreateUpdateRequest) HasTransformer() bool {
	if o != nil && o.Transformer != nil {
		return true
	}

	return false
}

// SetTransformer gets a reference to the given map[string]interface{} and assigns it to the Transformer field.
func (o *CreateUpdateRequest) SetTransformer(v map[string]interface{}) {
	o.Transformer = &v
}

// GetMapping returns the Mapping field value if set, zero value otherwise.
func (o *CreateUpdateRequest) GetMapping() map[string]interface{} {
	if o == nil || o.Mapping == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateRequest) GetMappingOk() (*map[string]interface{}, bool) {
	if o == nil || o.Mapping == nil {
		return nil, false
	}
	return o.Mapping, true
}

// HasMapping returns a boolean if a field has been set.
func (o *CreateUpdateRequest) HasMapping() bool {
	if o != nil && o.Mapping != nil {
		return true
	}

	return false
}

// SetMapping gets a reference to the given map[string]interface{} and assigns it to the Mapping field.
func (o *CreateUpdateRequest) SetMapping(v map[string]interface{}) {
	o.Mapping = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateUpdateRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateUpdateRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateUpdateRequest) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateUpdateRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateUpdateRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateUpdateRequest) SetId(v string) {
	o.Id = &v
}

// GetPropagationLabels returns the PropagationLabels field value if set, zero value otherwise.
func (o *CreateUpdateRequest) GetPropagationLabels() []string {
	if o == nil || o.PropagationLabels == nil {
		var ret []string
		return ret
	}
	return *o.PropagationLabels
}

// GetPropagationLabelsOk returns a tuple with the PropagationLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateRequest) GetPropagationLabelsOk() (*[]string, bool) {
	if o == nil || o.PropagationLabels == nil {
		return nil, false
	}
	return o.PropagationLabels, true
}

// HasPropagationLabels returns a boolean if a field has been set.
func (o *CreateUpdateRequest) HasPropagationLabels() bool {
	if o != nil && o.PropagationLabels != nil {
		return true
	}

	return false
}

// SetPropagationLabels gets a reference to the given []string and assigns it to the PropagationLabels field.
func (o *CreateUpdateRequest) SetPropagationLabels(v []string) {
	o.PropagationLabels = &v
}

func (o CreateUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultIncidentType != nil {
		toSerialize["defaultIncidentType"] = o.DefaultIncidentType
	}
	if o.KeyTypeMap != nil {
		toSerialize["keyTypeMap"] = o.KeyTypeMap
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Transformer != nil {
		toSerialize["transformer"] = o.Transformer
	}
	if o.Mapping != nil {
		toSerialize["mapping"] = o.Mapping
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PropagationLabels != nil {
		toSerialize["propagationLabels"] = o.PropagationLabels
	}
	return json.Marshal(toSerialize)
}

type NullableCreateUpdateRequest struct {
	value *CreateUpdateRequest
	isSet bool
}

func (v NullableCreateUpdateRequest) Get() *CreateUpdateRequest {
	return v.value
}

func (v *NullableCreateUpdateRequest) Set(val *CreateUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUpdateRequest(val *CreateUpdateRequest) *NullableCreateUpdateRequest {
	return &NullableCreateUpdateRequest{value: val, isSet: true}
}

func (v NullableCreateUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
