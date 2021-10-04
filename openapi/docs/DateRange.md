# DateRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromDate** | Pointer to **time.Time** |  | [optional] 
**FromDateLicense** | Pointer to **time.Time** |  | [optional] 
**Period** | Pointer to [**Period**](Period.md) |  | [optional] 
**ToDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDateRange

`func NewDateRange() *DateRange`

NewDateRange instantiates a new DateRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateRangeWithDefaults

`func NewDateRangeWithDefaults() *DateRange`

NewDateRangeWithDefaults instantiates a new DateRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromDate

`func (o *DateRange) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *DateRange) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *DateRange) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *DateRange) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetFromDateLicense

`func (o *DateRange) GetFromDateLicense() time.Time`

GetFromDateLicense returns the FromDateLicense field if non-nil, zero value otherwise.

### GetFromDateLicenseOk

`func (o *DateRange) GetFromDateLicenseOk() (*time.Time, bool)`

GetFromDateLicenseOk returns a tuple with the FromDateLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDateLicense

`func (o *DateRange) SetFromDateLicense(v time.Time)`

SetFromDateLicense sets FromDateLicense field to given value.

### HasFromDateLicense

`func (o *DateRange) HasFromDateLicense() bool`

HasFromDateLicense returns a boolean if a field has been set.

### GetPeriod

`func (o *DateRange) GetPeriod() Period`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DateRange) GetPeriodOk() (*Period, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DateRange) SetPeriod(v Period)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *DateRange) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetToDate

`func (o *DateRange) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *DateRange) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *DateRange) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *DateRange) HasToDate() bool`

HasToDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


