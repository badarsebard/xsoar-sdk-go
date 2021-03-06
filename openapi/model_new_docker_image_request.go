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

// NewDockerImageRequest NewDockerImageRequest creation request
type NewDockerImageRequest struct {
	Base         *string   `json:"base,omitempty"`
	Dependencies *[]string `json:"dependencies,omitempty"`
	Name         *string   `json:"name,omitempty"`
	Packages     *[]string `json:"packages,omitempty"`
}

// NewNewDockerImageRequest instantiates a new NewDockerImageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewDockerImageRequest() *NewDockerImageRequest {
	this := NewDockerImageRequest{}
	return &this
}

// NewNewDockerImageRequestWithDefaults instantiates a new NewDockerImageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewDockerImageRequestWithDefaults() *NewDockerImageRequest {
	this := NewDockerImageRequest{}
	return &this
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *NewDockerImageRequest) GetBase() string {
	if o == nil || o.Base == nil {
		var ret string
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewDockerImageRequest) GetBaseOk() (*string, bool) {
	if o == nil || o.Base == nil {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *NewDockerImageRequest) HasBase() bool {
	if o != nil && o.Base != nil {
		return true
	}

	return false
}

// SetBase gets a reference to the given string and assigns it to the Base field.
func (o *NewDockerImageRequest) SetBase(v string) {
	o.Base = &v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *NewDockerImageRequest) GetDependencies() []string {
	if o == nil || o.Dependencies == nil {
		var ret []string
		return ret
	}
	return *o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewDockerImageRequest) GetDependenciesOk() (*[]string, bool) {
	if o == nil || o.Dependencies == nil {
		return nil, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *NewDockerImageRequest) HasDependencies() bool {
	if o != nil && o.Dependencies != nil {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given []string and assigns it to the Dependencies field.
func (o *NewDockerImageRequest) SetDependencies(v []string) {
	o.Dependencies = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NewDockerImageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewDockerImageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NewDockerImageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NewDockerImageRequest) SetName(v string) {
	o.Name = &v
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *NewDockerImageRequest) GetPackages() []string {
	if o == nil || o.Packages == nil {
		var ret []string
		return ret
	}
	return *o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewDockerImageRequest) GetPackagesOk() (*[]string, bool) {
	if o == nil || o.Packages == nil {
		return nil, false
	}
	return o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *NewDockerImageRequest) HasPackages() bool {
	if o != nil && o.Packages != nil {
		return true
	}

	return false
}

// SetPackages gets a reference to the given []string and assigns it to the Packages field.
func (o *NewDockerImageRequest) SetPackages(v []string) {
	o.Packages = &v
}

func (o NewDockerImageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Base != nil {
		toSerialize["base"] = o.Base
	}
	if o.Dependencies != nil {
		toSerialize["dependencies"] = o.Dependencies
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Packages != nil {
		toSerialize["packages"] = o.Packages
	}
	return json.Marshal(toSerialize)
}

type NullableNewDockerImageRequest struct {
	value *NewDockerImageRequest
	isSet bool
}

func (v NullableNewDockerImageRequest) Get() *NewDockerImageRequest {
	return v.value
}

func (v *NullableNewDockerImageRequest) Set(val *NewDockerImageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNewDockerImageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNewDockerImageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewDockerImageRequest(val *NewDockerImageRequest) *NullableNewDockerImageRequest {
	return &NullableNewDockerImageRequest{value: val, isSet: true}
}

func (v NullableNewDockerImageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewDockerImageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
