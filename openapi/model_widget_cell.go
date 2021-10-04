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

// WidgetCell struct for WidgetCell
type WidgetCell struct {
	ForceRange        *bool   `json:"forceRange,omitempty"`
	H                 *int64  `json:"h,omitempty"`
	I                 *string `json:"i,omitempty"`
	Id                *string `json:"id,omitempty"`
	ReflectDimensions *bool   `json:"reflectDimensions,omitempty"`
	W                 *int64  `json:"w,omitempty"`
	Widget            *Widget `json:"widget,omitempty"`
	X                 *int64  `json:"x,omitempty"`
	Y                 *int64  `json:"y,omitempty"`
}

// NewWidgetCell instantiates a new WidgetCell object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetCell() *WidgetCell {
	this := WidgetCell{}
	return &this
}

// NewWidgetCellWithDefaults instantiates a new WidgetCell object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetCellWithDefaults() *WidgetCell {
	this := WidgetCell{}
	return &this
}

// GetForceRange returns the ForceRange field value if set, zero value otherwise.
func (o *WidgetCell) GetForceRange() bool {
	if o == nil || o.ForceRange == nil {
		var ret bool
		return ret
	}
	return *o.ForceRange
}

// GetForceRangeOk returns a tuple with the ForceRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetCell) GetForceRangeOk() (*bool, bool) {
	if o == nil || o.ForceRange == nil {
		return nil, false
	}
	return o.ForceRange, true
}

// HasForceRange returns a boolean if a field has been set.
func (o *WidgetCell) HasForceRange() bool {
	if o != nil && o.ForceRange != nil {
		return true
	}

	return false
}

// SetForceRange gets a reference to the given bool and assigns it to the ForceRange field.
func (o *WidgetCell) SetForceRange(v bool) {
	o.ForceRange = &v
}

// GetH returns the H field value if set, zero value otherwise.
func (o *WidgetCell) GetH() int64 {
	if o == nil || o.H == nil {
		var ret int64
		return ret
	}
	return *o.H
}

// GetHOk returns a tuple with the H field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetCell) GetHOk() (*int64, bool) {
	if o == nil || o.H == nil {
		return nil, false
	}
	return o.H, true
}

// HasH returns a boolean if a field has been set.
func (o *WidgetCell) HasH() bool {
	if o != nil && o.H != nil {
		return true
	}

	return false
}

// SetH gets a reference to the given int64 and assigns it to the H field.
func (o *WidgetCell) SetH(v int64) {
	o.H = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *WidgetCell) GetI() string {
	if o == nil || o.I == nil {
		var ret string
		return ret
	}
	return *o.I
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetCell) GetIOk() (*string, bool) {
	if o == nil || o.I == nil {
		return nil, false
	}
	return o.I, true
}

// HasI returns a boolean if a field has been set.
func (o *WidgetCell) HasI() bool {
	if o != nil && o.I != nil {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *WidgetCell) SetI(v string) {
	o.I = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WidgetCell) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetCell) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WidgetCell) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WidgetCell) SetId(v string) {
	o.Id = &v
}

// GetReflectDimensions returns the ReflectDimensions field value if set, zero value otherwise.
func (o *WidgetCell) GetReflectDimensions() bool {
	if o == nil || o.ReflectDimensions == nil {
		var ret bool
		return ret
	}
	return *o.ReflectDimensions
}

// GetReflectDimensionsOk returns a tuple with the ReflectDimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetCell) GetReflectDimensionsOk() (*bool, bool) {
	if o == nil || o.ReflectDimensions == nil {
		return nil, false
	}
	return o.ReflectDimensions, true
}

// HasReflectDimensions returns a boolean if a field has been set.
func (o *WidgetCell) HasReflectDimensions() bool {
	if o != nil && o.ReflectDimensions != nil {
		return true
	}

	return false
}

// SetReflectDimensions gets a reference to the given bool and assigns it to the ReflectDimensions field.
func (o *WidgetCell) SetReflectDimensions(v bool) {
	o.ReflectDimensions = &v
}

// GetW returns the W field value if set, zero value otherwise.
func (o *WidgetCell) GetW() int64 {
	if o == nil || o.W == nil {
		var ret int64
		return ret
	}
	return *o.W
}

// GetWOk returns a tuple with the W field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetCell) GetWOk() (*int64, bool) {
	if o == nil || o.W == nil {
		return nil, false
	}
	return o.W, true
}

// HasW returns a boolean if a field has been set.
func (o *WidgetCell) HasW() bool {
	if o != nil && o.W != nil {
		return true
	}

	return false
}

// SetW gets a reference to the given int64 and assigns it to the W field.
func (o *WidgetCell) SetW(v int64) {
	o.W = &v
}

// GetWidget returns the Widget field value if set, zero value otherwise.
func (o *WidgetCell) GetWidget() Widget {
	if o == nil || o.Widget == nil {
		var ret Widget
		return ret
	}
	return *o.Widget
}

// GetWidgetOk returns a tuple with the Widget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetCell) GetWidgetOk() (*Widget, bool) {
	if o == nil || o.Widget == nil {
		return nil, false
	}
	return o.Widget, true
}

// HasWidget returns a boolean if a field has been set.
func (o *WidgetCell) HasWidget() bool {
	if o != nil && o.Widget != nil {
		return true
	}

	return false
}

// SetWidget gets a reference to the given Widget and assigns it to the Widget field.
func (o *WidgetCell) SetWidget(v Widget) {
	o.Widget = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *WidgetCell) GetX() int64 {
	if o == nil || o.X == nil {
		var ret int64
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetCell) GetXOk() (*int64, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *WidgetCell) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given int64 and assigns it to the X field.
func (o *WidgetCell) SetX(v int64) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *WidgetCell) GetY() int64 {
	if o == nil || o.Y == nil {
		var ret int64
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetCell) GetYOk() (*int64, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *WidgetCell) HasY() bool {
	if o != nil && o.Y != nil {
		return true
	}

	return false
}

// SetY gets a reference to the given int64 and assigns it to the Y field.
func (o *WidgetCell) SetY(v int64) {
	o.Y = &v
}

func (o WidgetCell) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ForceRange != nil {
		toSerialize["forceRange"] = o.ForceRange
	}
	if o.H != nil {
		toSerialize["h"] = o.H
	}
	if o.I != nil {
		toSerialize["i"] = o.I
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ReflectDimensions != nil {
		toSerialize["reflectDimensions"] = o.ReflectDimensions
	}
	if o.W != nil {
		toSerialize["w"] = o.W
	}
	if o.Widget != nil {
		toSerialize["widget"] = o.Widget
	}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.Y != nil {
		toSerialize["y"] = o.Y
	}
	return json.Marshal(toSerialize)
}

type NullableWidgetCell struct {
	value *WidgetCell
	isSet bool
}

func (v NullableWidgetCell) Get() *WidgetCell {
	return v.value
}

func (v *NullableWidgetCell) Set(val *WidgetCell) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetCell) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetCell) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetCell(val *WidgetCell) *NullableWidgetCell {
	return &NullableWidgetCell{value: val, isSet: true}
}

func (v NullableWidgetCell) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetCell) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
