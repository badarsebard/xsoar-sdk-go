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

// GenericStringDateFilter GenericStringDateFilter is a general filter that will fetch entities using the Query value and a date filter
type GenericStringDateFilter struct {
	// Cache of join functions
	Cache           *map[string][]string `json:"Cache,omitempty"`
	Fields          *[]string            `json:"fields,omitempty"`
	FromDate        *time.Time           `json:"fromDate,omitempty"`
	FromDateLicense *time.Time           `json:"fromDateLicense,omitempty"`
	// Do not use workers mechanism while searching bleve
	IgnoreWorkers *bool `json:"ignoreWorkers,omitempty"`
	// 0-based page
	Page   *int64  `json:"page,omitempty"`
	Period *Period `json:"period,omitempty"`
	Query  *string `json:"query,omitempty"`
	// Efficient next page, pass max sort value from previous page
	SearchAfter *[]string `json:"searchAfter,omitempty"`
	// Efficient next page, pass max ES sort value from previous page
	SearchAfterElastic *[]string `json:"searchAfterElastic,omitempty"`
	// Efficient prev page, pass min sort value from next page
	SearchBefore *[]string `json:"searchBefore,omitempty"`
	// Efficient prev page, pass min ES sort value from next page
	SearchBeforeElastic *[]string `json:"searchBeforeElastic,omitempty"`
	// Size is limited to 1000, if not passed it defaults to 0, and no results will return
	Size *int64 `json:"size,omitempty"`
	// The sort order
	Sort *[]Order `json:"sort,omitempty"`
	// A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years.
	TimeFrame *int64     `json:"timeFrame,omitempty"`
	ToDate    *time.Time `json:"toDate,omitempty"`
}

// NewGenericStringDateFilter instantiates a new GenericStringDateFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericStringDateFilter() *GenericStringDateFilter {
	this := GenericStringDateFilter{}
	return &this
}

// NewGenericStringDateFilterWithDefaults instantiates a new GenericStringDateFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericStringDateFilterWithDefaults() *GenericStringDateFilter {
	this := GenericStringDateFilter{}
	return &this
}

// GetCache returns the Cache field value if set, zero value otherwise.
func (o *GenericStringDateFilter) GetCache() map[string][]string {
	if o == nil || o.Cache == nil {
		var ret map[string][]string
		return ret
	}
	return *o.Cache
}

// GetCacheOk returns a tuple with the Cache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericStringDateFilter) GetCacheOk() (*map[string][]string, bool) {
	if o == nil || o.Cache == nil {
		return nil, false
	}
	return o.Cache, true
}

// HasCache returns a boolean if a field has been set.
func (o *GenericStringDateFilter) HasCache() bool {
	if o != nil && o.Cache != nil {
		return true
	}

	return false
}

// SetCache gets a reference to the given map[string][]string and assigns it to the Cache field.
func (o *GenericStringDateFilter) SetCache(v map[string][]string) {
	o.Cache = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *GenericStringDateFilter) GetFields() []string {
	if o == nil || o.Fields == nil {
		var ret []string
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericStringDateFilter) GetFieldsOk() (*[]string, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *GenericStringDateFilter) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []string and assigns it to the Fields field.
func (o *GenericStringDateFilter) SetFields(v []string) {
	o.Fields = &v
}

// GetFromDate returns the FromDate field value if set, zero value otherwise.
func (o *GenericStringDateFilter) GetFromDate() time.Time {
	if o == nil || o.FromDate == nil {
		var ret time.Time
		return ret
	}
	return *o.FromDate
}

// GetFromDateOk returns a tuple with the FromDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericStringDateFilter) GetFromDateOk() (*time.Time, bool) {
	if o == nil || o.FromDate == nil {
		return nil, false
	}
	return o.FromDate, true
}

// HasFromDate returns a boolean if a field has been set.
func (o *GenericStringDateFilter) HasFromDate() bool {
	if o != nil && o.FromDate != nil {
		return true
	}

	return false
}

// SetFromDate gets a reference to the given time.Time and assigns it to the FromDate field.
func (o *GenericStringDateFilter) SetFromDate(v time.Time) {
	o.FromDate = &v
}

// GetFromDateLicense returns the FromDateLicense field value if set, zero value otherwise.
func (o *GenericStringDateFilter) GetFromDateLicense() time.Time {
	if o == nil || o.FromDateLicense == nil {
		var ret time.Time
		return ret
	}
	return *o.FromDateLicense
}

// GetFromDateLicenseOk returns a tuple with the FromDateLicense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericStringDateFilter) GetFromDateLicenseOk() (*time.Time, bool) {
	if o == nil || o.FromDateLicense == nil {
		return nil, false
	}
	return o.FromDateLicense, true
}

// HasFromDateLicense returns a boolean if a field has been set.
func (o *GenericStringDateFilter) HasFromDateLicense() bool {
	if o != nil && o.FromDateLicense != nil {
		return true
	}

	return false
}

// SetFromDateLicense gets a reference to the given time.Time and assigns it to the FromDateLicense field.
func (o *GenericStringDateFilter) SetFromDateLicense(v time.Time) {
	o.FromDateLicense = &v
}

// GetIgnoreWorkers returns the IgnoreWorkers field value if set, zero value otherwise.
func (o *GenericStringDateFilter) GetIgnoreWorkers() bool {
	if o == nil || o.IgnoreWorkers == nil {
		var ret bool
		return ret
	}
	return *o.IgnoreWorkers
}

// GetIgnoreWorkersOk returns a tuple with the IgnoreWorkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericStringDateFilter) GetIgnoreWorkersOk() (*bool, bool) {
	if o == nil || o.IgnoreWorkers == nil {
		return nil, false
	}
	return o.IgnoreWorkers, true
}

// HasIgnoreWorkers returns a boolean if a field has been set.
func (o *GenericStringDateFilter) HasIgnoreWorkers() bool {
	if o != nil && o.IgnoreWorkers != nil {
		return true
	}

	return false
}

// SetIgnoreWorkers gets a reference to the given bool and assigns it to the IgnoreWorkers field.
func (o *GenericStringDateFilter) SetIgnoreWorkers(v bool) {
	o.IgnoreWorkers = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GenericStringDateFilter) GetPage() int64 {
	if o == nil || o.Page == nil {
		var ret int64
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericStringDateFilter) GetPageOk() (*int64, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GenericStringDateFilter) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int64 and assigns it to the Page field.
func (o *GenericStringDateFilter) SetPage(v int64) {
	o.Page = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *GenericStringDateFilter) GetPeriod() Period {
	if o == nil || o.Period == nil {
		var ret Period
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericStringDateFilter) GetPeriodOk() (*Period, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *GenericStringDateFilter) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given Period and assigns it to the Period field.
func (o *GenericStringDateFilter) SetPeriod(v Period) {
	o.Period = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *GenericStringDateFilter) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericStringDateFilter) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *GenericStringDateFilter) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *GenericStringDateFilter) SetQuery(v string) {
	o.Query = &v
}

// GetSearchAfter returns the SearchAfter field value if set, zero value otherwise.
func (o *GenericStringDateFilter) GetSearchAfter() []string {
	if o == nil || o.SearchAfter == nil {
		var ret []string
		return ret
	}
	return *o.SearchAfter
}

// GetSearchAfterOk returns a tuple with the SearchAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericStringDateFilter) GetSearchAfterOk() (*[]string, bool) {
	if o == nil || o.SearchAfter == nil {
		return nil, false
	}
	return o.SearchAfter, true
}

// HasSearchAfter returns a boolean if a field has been set.
func (o *GenericStringDateFilter) HasSearchAfter() bool {
	if o != nil && o.SearchAfter != nil {
		return true
	}

	return false
}

// SetSearchAfter gets a reference to the given []string and assigns it to the SearchAfter field.
func (o *GenericStringDateFilter) SetSearchAfter(v []string) {
	o.SearchAfter = &v
}

// GetSearchAfterElastic returns the SearchAfterElastic field value if set, zero value otherwise.
func (o *GenericStringDateFilter) GetSearchAfterElastic() []string {
	if o == nil || o.SearchAfterElastic == nil {
		var ret []string
		return ret
	}
	return *o.SearchAfterElastic
}

// GetSearchAfterElasticOk returns a tuple with the SearchAfterElastic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericStringDateFilter) GetSearchAfterElasticOk() (*[]string, bool) {
	if o == nil || o.SearchAfterElastic == nil {
		return nil, false
	}
	return o.SearchAfterElastic, true
}

// HasSearchAfterElastic returns a boolean if a field has been set.
func (o *GenericStringDateFilter) HasSearchAfterElastic() bool {
	if o != nil && o.SearchAfterElastic != nil {
		return true
	}

	return false
}

// SetSearchAfterElastic gets a reference to the given []string and assigns it to the SearchAfterElastic field.
func (o *GenericStringDateFilter) SetSearchAfterElastic(v []string) {
	o.SearchAfterElastic = &v
}

// GetSearchBefore returns the SearchBefore field value if set, zero value otherwise.
func (o *GenericStringDateFilter) GetSearchBefore() []string {
	if o == nil || o.SearchBefore == nil {
		var ret []string
		return ret
	}
	return *o.SearchBefore
}

// GetSearchBeforeOk returns a tuple with the SearchBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericStringDateFilter) GetSearchBeforeOk() (*[]string, bool) {
	if o == nil || o.SearchBefore == nil {
		return nil, false
	}
	return o.SearchBefore, true
}

// HasSearchBefore returns a boolean if a field has been set.
func (o *GenericStringDateFilter) HasSearchBefore() bool {
	if o != nil && o.SearchBefore != nil {
		return true
	}

	return false
}

// SetSearchBefore gets a reference to the given []string and assigns it to the SearchBefore field.
func (o *GenericStringDateFilter) SetSearchBefore(v []string) {
	o.SearchBefore = &v
}

// GetSearchBeforeElastic returns the SearchBeforeElastic field value if set, zero value otherwise.
func (o *GenericStringDateFilter) GetSearchBeforeElastic() []string {
	if o == nil || o.SearchBeforeElastic == nil {
		var ret []string
		return ret
	}
	return *o.SearchBeforeElastic
}

// GetSearchBeforeElasticOk returns a tuple with the SearchBeforeElastic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericStringDateFilter) GetSearchBeforeElasticOk() (*[]string, bool) {
	if o == nil || o.SearchBeforeElastic == nil {
		return nil, false
	}
	return o.SearchBeforeElastic, true
}

// HasSearchBeforeElastic returns a boolean if a field has been set.
func (o *GenericStringDateFilter) HasSearchBeforeElastic() bool {
	if o != nil && o.SearchBeforeElastic != nil {
		return true
	}

	return false
}

// SetSearchBeforeElastic gets a reference to the given []string and assigns it to the SearchBeforeElastic field.
func (o *GenericStringDateFilter) SetSearchBeforeElastic(v []string) {
	o.SearchBeforeElastic = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *GenericStringDateFilter) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericStringDateFilter) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *GenericStringDateFilter) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *GenericStringDateFilter) SetSize(v int64) {
	o.Size = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *GenericStringDateFilter) GetSort() []Order {
	if o == nil || o.Sort == nil {
		var ret []Order
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericStringDateFilter) GetSortOk() (*[]Order, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *GenericStringDateFilter) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given []Order and assigns it to the Sort field.
func (o *GenericStringDateFilter) SetSort(v []Order) {
	o.Sort = &v
}

// GetTimeFrame returns the TimeFrame field value if set, zero value otherwise.
func (o *GenericStringDateFilter) GetTimeFrame() int64 {
	if o == nil || o.TimeFrame == nil {
		var ret int64
		return ret
	}
	return *o.TimeFrame
}

// GetTimeFrameOk returns a tuple with the TimeFrame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericStringDateFilter) GetTimeFrameOk() (*int64, bool) {
	if o == nil || o.TimeFrame == nil {
		return nil, false
	}
	return o.TimeFrame, true
}

// HasTimeFrame returns a boolean if a field has been set.
func (o *GenericStringDateFilter) HasTimeFrame() bool {
	if o != nil && o.TimeFrame != nil {
		return true
	}

	return false
}

// SetTimeFrame gets a reference to the given int64 and assigns it to the TimeFrame field.
func (o *GenericStringDateFilter) SetTimeFrame(v int64) {
	o.TimeFrame = &v
}

// GetToDate returns the ToDate field value if set, zero value otherwise.
func (o *GenericStringDateFilter) GetToDate() time.Time {
	if o == nil || o.ToDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ToDate
}

// GetToDateOk returns a tuple with the ToDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericStringDateFilter) GetToDateOk() (*time.Time, bool) {
	if o == nil || o.ToDate == nil {
		return nil, false
	}
	return o.ToDate, true
}

// HasToDate returns a boolean if a field has been set.
func (o *GenericStringDateFilter) HasToDate() bool {
	if o != nil && o.ToDate != nil {
		return true
	}

	return false
}

// SetToDate gets a reference to the given time.Time and assigns it to the ToDate field.
func (o *GenericStringDateFilter) SetToDate(v time.Time) {
	o.ToDate = &v
}

func (o GenericStringDateFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cache != nil {
		toSerialize["Cache"] = o.Cache
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.FromDate != nil {
		toSerialize["fromDate"] = o.FromDate
	}
	if o.FromDateLicense != nil {
		toSerialize["fromDateLicense"] = o.FromDateLicense
	}
	if o.IgnoreWorkers != nil {
		toSerialize["ignoreWorkers"] = o.IgnoreWorkers
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.SearchAfter != nil {
		toSerialize["searchAfter"] = o.SearchAfter
	}
	if o.SearchAfterElastic != nil {
		toSerialize["searchAfterElastic"] = o.SearchAfterElastic
	}
	if o.SearchBefore != nil {
		toSerialize["searchBefore"] = o.SearchBefore
	}
	if o.SearchBeforeElastic != nil {
		toSerialize["searchBeforeElastic"] = o.SearchBeforeElastic
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.TimeFrame != nil {
		toSerialize["timeFrame"] = o.TimeFrame
	}
	if o.ToDate != nil {
		toSerialize["toDate"] = o.ToDate
	}
	return json.Marshal(toSerialize)
}

type NullableGenericStringDateFilter struct {
	value *GenericStringDateFilter
	isSet bool
}

func (v NullableGenericStringDateFilter) Get() *GenericStringDateFilter {
	return v.value
}

func (v *NullableGenericStringDateFilter) Set(val *GenericStringDateFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericStringDateFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericStringDateFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericStringDateFilter(val *GenericStringDateFilter) *NullableGenericStringDateFilter {
	return &NullableGenericStringDateFilter{value: val, isSet: true}
}

func (v NullableGenericStringDateFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericStringDateFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
