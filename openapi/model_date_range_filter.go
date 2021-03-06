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

// DateRangeFilter DateRangeFilter provides common fields for date filtering
type DateRangeFilter struct {
	FromDate        *time.Time `json:"fromDate,omitempty"`
	FromDateLicense *time.Time `json:"fromDateLicense,omitempty"`
	Period          *Period    `json:"period,omitempty"`
	// A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years.
	TimeFrame *int64     `json:"timeFrame,omitempty"`
	ToDate    *time.Time `json:"toDate,omitempty"`
}

// NewDateRangeFilter instantiates a new DateRangeFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateRangeFilter() *DateRangeFilter {
	this := DateRangeFilter{}
	return &this
}

// NewDateRangeFilterWithDefaults instantiates a new DateRangeFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateRangeFilterWithDefaults() *DateRangeFilter {
	this := DateRangeFilter{}
	return &this
}

// GetFromDate returns the FromDate field value if set, zero value otherwise.
func (o *DateRangeFilter) GetFromDate() time.Time {
	if o == nil || o.FromDate == nil {
		var ret time.Time
		return ret
	}
	return *o.FromDate
}

// GetFromDateOk returns a tuple with the FromDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateRangeFilter) GetFromDateOk() (*time.Time, bool) {
	if o == nil || o.FromDate == nil {
		return nil, false
	}
	return o.FromDate, true
}

// HasFromDate returns a boolean if a field has been set.
func (o *DateRangeFilter) HasFromDate() bool {
	if o != nil && o.FromDate != nil {
		return true
	}

	return false
}

// SetFromDate gets a reference to the given time.Time and assigns it to the FromDate field.
func (o *DateRangeFilter) SetFromDate(v time.Time) {
	o.FromDate = &v
}

// GetFromDateLicense returns the FromDateLicense field value if set, zero value otherwise.
func (o *DateRangeFilter) GetFromDateLicense() time.Time {
	if o == nil || o.FromDateLicense == nil {
		var ret time.Time
		return ret
	}
	return *o.FromDateLicense
}

// GetFromDateLicenseOk returns a tuple with the FromDateLicense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateRangeFilter) GetFromDateLicenseOk() (*time.Time, bool) {
	if o == nil || o.FromDateLicense == nil {
		return nil, false
	}
	return o.FromDateLicense, true
}

// HasFromDateLicense returns a boolean if a field has been set.
func (o *DateRangeFilter) HasFromDateLicense() bool {
	if o != nil && o.FromDateLicense != nil {
		return true
	}

	return false
}

// SetFromDateLicense gets a reference to the given time.Time and assigns it to the FromDateLicense field.
func (o *DateRangeFilter) SetFromDateLicense(v time.Time) {
	o.FromDateLicense = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *DateRangeFilter) GetPeriod() Period {
	if o == nil || o.Period == nil {
		var ret Period
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateRangeFilter) GetPeriodOk() (*Period, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *DateRangeFilter) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given Period and assigns it to the Period field.
func (o *DateRangeFilter) SetPeriod(v Period) {
	o.Period = &v
}

// GetTimeFrame returns the TimeFrame field value if set, zero value otherwise.
func (o *DateRangeFilter) GetTimeFrame() int64 {
	if o == nil || o.TimeFrame == nil {
		var ret int64
		return ret
	}
	return *o.TimeFrame
}

// GetTimeFrameOk returns a tuple with the TimeFrame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateRangeFilter) GetTimeFrameOk() (*int64, bool) {
	if o == nil || o.TimeFrame == nil {
		return nil, false
	}
	return o.TimeFrame, true
}

// HasTimeFrame returns a boolean if a field has been set.
func (o *DateRangeFilter) HasTimeFrame() bool {
	if o != nil && o.TimeFrame != nil {
		return true
	}

	return false
}

// SetTimeFrame gets a reference to the given int64 and assigns it to the TimeFrame field.
func (o *DateRangeFilter) SetTimeFrame(v int64) {
	o.TimeFrame = &v
}

// GetToDate returns the ToDate field value if set, zero value otherwise.
func (o *DateRangeFilter) GetToDate() time.Time {
	if o == nil || o.ToDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ToDate
}

// GetToDateOk returns a tuple with the ToDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateRangeFilter) GetToDateOk() (*time.Time, bool) {
	if o == nil || o.ToDate == nil {
		return nil, false
	}
	return o.ToDate, true
}

// HasToDate returns a boolean if a field has been set.
func (o *DateRangeFilter) HasToDate() bool {
	if o != nil && o.ToDate != nil {
		return true
	}

	return false
}

// SetToDate gets a reference to the given time.Time and assigns it to the ToDate field.
func (o *DateRangeFilter) SetToDate(v time.Time) {
	o.ToDate = &v
}

func (o DateRangeFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FromDate != nil {
		toSerialize["fromDate"] = o.FromDate
	}
	if o.FromDateLicense != nil {
		toSerialize["fromDateLicense"] = o.FromDateLicense
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.TimeFrame != nil {
		toSerialize["timeFrame"] = o.TimeFrame
	}
	if o.ToDate != nil {
		toSerialize["toDate"] = o.ToDate
	}
	return json.Marshal(toSerialize)
}

type NullableDateRangeFilter struct {
	value *DateRangeFilter
	isSet bool
}

func (v NullableDateRangeFilter) Get() *DateRangeFilter {
	return v.value
}

func (v *NullableDateRangeFilter) Set(val *DateRangeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableDateRangeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableDateRangeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateRangeFilter(val *DateRangeFilter) *NullableDateRangeFilter {
	return &NullableDateRangeFilter{value: val, isSet: true}
}

func (v NullableDateRangeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateRangeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
