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

// DockerImagesResult DockerImagesResult for the get images request
type DockerImagesResult struct {
	Images *[]DockerImage `json:"images,omitempty"`
}

// NewDockerImagesResult instantiates a new DockerImagesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDockerImagesResult() *DockerImagesResult {
	this := DockerImagesResult{}
	return &this
}

// NewDockerImagesResultWithDefaults instantiates a new DockerImagesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDockerImagesResultWithDefaults() *DockerImagesResult {
	this := DockerImagesResult{}
	return &this
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *DockerImagesResult) GetImages() []DockerImage {
	if o == nil || o.Images == nil {
		var ret []DockerImage
		return ret
	}
	return *o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerImagesResult) GetImagesOk() (*[]DockerImage, bool) {
	if o == nil || o.Images == nil {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *DockerImagesResult) HasImages() bool {
	if o != nil && o.Images != nil {
		return true
	}

	return false
}

// SetImages gets a reference to the given []DockerImage and assigns it to the Images field.
func (o *DockerImagesResult) SetImages(v []DockerImage) {
	o.Images = &v
}

func (o DockerImagesResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	return json.Marshal(toSerialize)
}

type NullableDockerImagesResult struct {
	value *DockerImagesResult
	isSet bool
}

func (v NullableDockerImagesResult) Get() *DockerImagesResult {
	return v.value
}

func (v *NullableDockerImagesResult) Set(val *DockerImagesResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDockerImagesResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDockerImagesResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDockerImagesResult(val *DockerImagesResult) *NullableDockerImagesResult {
	return &NullableDockerImagesResult{value: val, isSet: true}
}

func (v NullableDockerImagesResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDockerImagesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
