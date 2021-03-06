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

// Scheduler struct for Scheduler
type Scheduler struct {
	Cron       *string    `json:"cron,omitempty"`
	CronView   *bool      `json:"cronView,omitempty"`
	EndingDate *time.Time `json:"endingDate,omitempty"`
	// EndingType holds the type of schedule Ending
	EndingType     *string    `json:"endingType,omitempty"`
	HumanCron      *HumanCron `json:"humanCron,omitempty"`
	Recurrent      *bool      `json:"recurrent,omitempty"`
	StartDate      *time.Time `json:"startDate,omitempty"`
	Times          *int64     `json:"times,omitempty"`
	Timezone       *string    `json:"timezone,omitempty"`
	TimezoneOffset *int64     `json:"timezoneOffset,omitempty"`
}

// NewScheduler instantiates a new Scheduler object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduler() *Scheduler {
	this := Scheduler{}
	return &this
}

// NewSchedulerWithDefaults instantiates a new Scheduler object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulerWithDefaults() *Scheduler {
	this := Scheduler{}
	return &this
}

// GetCron returns the Cron field value if set, zero value otherwise.
func (o *Scheduler) GetCron() string {
	if o == nil || o.Cron == nil {
		var ret string
		return ret
	}
	return *o.Cron
}

// GetCronOk returns a tuple with the Cron field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scheduler) GetCronOk() (*string, bool) {
	if o == nil || o.Cron == nil {
		return nil, false
	}
	return o.Cron, true
}

// HasCron returns a boolean if a field has been set.
func (o *Scheduler) HasCron() bool {
	if o != nil && o.Cron != nil {
		return true
	}

	return false
}

// SetCron gets a reference to the given string and assigns it to the Cron field.
func (o *Scheduler) SetCron(v string) {
	o.Cron = &v
}

// GetCronView returns the CronView field value if set, zero value otherwise.
func (o *Scheduler) GetCronView() bool {
	if o == nil || o.CronView == nil {
		var ret bool
		return ret
	}
	return *o.CronView
}

// GetCronViewOk returns a tuple with the CronView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scheduler) GetCronViewOk() (*bool, bool) {
	if o == nil || o.CronView == nil {
		return nil, false
	}
	return o.CronView, true
}

// HasCronView returns a boolean if a field has been set.
func (o *Scheduler) HasCronView() bool {
	if o != nil && o.CronView != nil {
		return true
	}

	return false
}

// SetCronView gets a reference to the given bool and assigns it to the CronView field.
func (o *Scheduler) SetCronView(v bool) {
	o.CronView = &v
}

// GetEndingDate returns the EndingDate field value if set, zero value otherwise.
func (o *Scheduler) GetEndingDate() time.Time {
	if o == nil || o.EndingDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndingDate
}

// GetEndingDateOk returns a tuple with the EndingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scheduler) GetEndingDateOk() (*time.Time, bool) {
	if o == nil || o.EndingDate == nil {
		return nil, false
	}
	return o.EndingDate, true
}

// HasEndingDate returns a boolean if a field has been set.
func (o *Scheduler) HasEndingDate() bool {
	if o != nil && o.EndingDate != nil {
		return true
	}

	return false
}

// SetEndingDate gets a reference to the given time.Time and assigns it to the EndingDate field.
func (o *Scheduler) SetEndingDate(v time.Time) {
	o.EndingDate = &v
}

// GetEndingType returns the EndingType field value if set, zero value otherwise.
func (o *Scheduler) GetEndingType() string {
	if o == nil || o.EndingType == nil {
		var ret string
		return ret
	}
	return *o.EndingType
}

// GetEndingTypeOk returns a tuple with the EndingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scheduler) GetEndingTypeOk() (*string, bool) {
	if o == nil || o.EndingType == nil {
		return nil, false
	}
	return o.EndingType, true
}

// HasEndingType returns a boolean if a field has been set.
func (o *Scheduler) HasEndingType() bool {
	if o != nil && o.EndingType != nil {
		return true
	}

	return false
}

// SetEndingType gets a reference to the given string and assigns it to the EndingType field.
func (o *Scheduler) SetEndingType(v string) {
	o.EndingType = &v
}

// GetHumanCron returns the HumanCron field value if set, zero value otherwise.
func (o *Scheduler) GetHumanCron() HumanCron {
	if o == nil || o.HumanCron == nil {
		var ret HumanCron
		return ret
	}
	return *o.HumanCron
}

// GetHumanCronOk returns a tuple with the HumanCron field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scheduler) GetHumanCronOk() (*HumanCron, bool) {
	if o == nil || o.HumanCron == nil {
		return nil, false
	}
	return o.HumanCron, true
}

// HasHumanCron returns a boolean if a field has been set.
func (o *Scheduler) HasHumanCron() bool {
	if o != nil && o.HumanCron != nil {
		return true
	}

	return false
}

// SetHumanCron gets a reference to the given HumanCron and assigns it to the HumanCron field.
func (o *Scheduler) SetHumanCron(v HumanCron) {
	o.HumanCron = &v
}

// GetRecurrent returns the Recurrent field value if set, zero value otherwise.
func (o *Scheduler) GetRecurrent() bool {
	if o == nil || o.Recurrent == nil {
		var ret bool
		return ret
	}
	return *o.Recurrent
}

// GetRecurrentOk returns a tuple with the Recurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scheduler) GetRecurrentOk() (*bool, bool) {
	if o == nil || o.Recurrent == nil {
		return nil, false
	}
	return o.Recurrent, true
}

// HasRecurrent returns a boolean if a field has been set.
func (o *Scheduler) HasRecurrent() bool {
	if o != nil && o.Recurrent != nil {
		return true
	}

	return false
}

// SetRecurrent gets a reference to the given bool and assigns it to the Recurrent field.
func (o *Scheduler) SetRecurrent(v bool) {
	o.Recurrent = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Scheduler) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scheduler) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Scheduler) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *Scheduler) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetTimes returns the Times field value if set, zero value otherwise.
func (o *Scheduler) GetTimes() int64 {
	if o == nil || o.Times == nil {
		var ret int64
		return ret
	}
	return *o.Times
}

// GetTimesOk returns a tuple with the Times field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scheduler) GetTimesOk() (*int64, bool) {
	if o == nil || o.Times == nil {
		return nil, false
	}
	return o.Times, true
}

// HasTimes returns a boolean if a field has been set.
func (o *Scheduler) HasTimes() bool {
	if o != nil && o.Times != nil {
		return true
	}

	return false
}

// SetTimes gets a reference to the given int64 and assigns it to the Times field.
func (o *Scheduler) SetTimes(v int64) {
	o.Times = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *Scheduler) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scheduler) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *Scheduler) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *Scheduler) SetTimezone(v string) {
	o.Timezone = &v
}

// GetTimezoneOffset returns the TimezoneOffset field value if set, zero value otherwise.
func (o *Scheduler) GetTimezoneOffset() int64 {
	if o == nil || o.TimezoneOffset == nil {
		var ret int64
		return ret
	}
	return *o.TimezoneOffset
}

// GetTimezoneOffsetOk returns a tuple with the TimezoneOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scheduler) GetTimezoneOffsetOk() (*int64, bool) {
	if o == nil || o.TimezoneOffset == nil {
		return nil, false
	}
	return o.TimezoneOffset, true
}

// HasTimezoneOffset returns a boolean if a field has been set.
func (o *Scheduler) HasTimezoneOffset() bool {
	if o != nil && o.TimezoneOffset != nil {
		return true
	}

	return false
}

// SetTimezoneOffset gets a reference to the given int64 and assigns it to the TimezoneOffset field.
func (o *Scheduler) SetTimezoneOffset(v int64) {
	o.TimezoneOffset = &v
}

func (o Scheduler) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cron != nil {
		toSerialize["cron"] = o.Cron
	}
	if o.CronView != nil {
		toSerialize["cronView"] = o.CronView
	}
	if o.EndingDate != nil {
		toSerialize["endingDate"] = o.EndingDate
	}
	if o.EndingType != nil {
		toSerialize["endingType"] = o.EndingType
	}
	if o.HumanCron != nil {
		toSerialize["humanCron"] = o.HumanCron
	}
	if o.Recurrent != nil {
		toSerialize["recurrent"] = o.Recurrent
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.Times != nil {
		toSerialize["times"] = o.Times
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if o.TimezoneOffset != nil {
		toSerialize["timezoneOffset"] = o.TimezoneOffset
	}
	return json.Marshal(toSerialize)
}

type NullableScheduler struct {
	value *Scheduler
	isSet bool
}

func (v NullableScheduler) Get() *Scheduler {
	return v.value
}

func (v *NullableScheduler) Set(val *Scheduler) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduler) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduler) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduler(val *Scheduler) *NullableScheduler {
	return &NullableScheduler{value: val, isSet: true}
}

func (v NullableScheduler) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduler) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
