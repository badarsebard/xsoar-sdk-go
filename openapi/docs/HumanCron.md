# HumanCron

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtTimeHour** | Pointer to **string** |  | [optional] 
**AtTimeMinute** | Pointer to **string** |  | [optional] 
**Days** | Pointer to **[]string** |  | [optional] 
**SchedulingType** | Pointer to **string** | the following fields are deprecated. do not use them. | [optional] 
**TimePeriod** | Pointer to **int64** |  | [optional] 
**TimePeriodType** | Pointer to **string** |  | [optional] 

## Methods

### NewHumanCron

`func NewHumanCron() *HumanCron`

NewHumanCron instantiates a new HumanCron object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHumanCronWithDefaults

`func NewHumanCronWithDefaults() *HumanCron`

NewHumanCronWithDefaults instantiates a new HumanCron object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtTimeHour

`func (o *HumanCron) GetAtTimeHour() string`

GetAtTimeHour returns the AtTimeHour field if non-nil, zero value otherwise.

### GetAtTimeHourOk

`func (o *HumanCron) GetAtTimeHourOk() (*string, bool)`

GetAtTimeHourOk returns a tuple with the AtTimeHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtTimeHour

`func (o *HumanCron) SetAtTimeHour(v string)`

SetAtTimeHour sets AtTimeHour field to given value.

### HasAtTimeHour

`func (o *HumanCron) HasAtTimeHour() bool`

HasAtTimeHour returns a boolean if a field has been set.

### GetAtTimeMinute

`func (o *HumanCron) GetAtTimeMinute() string`

GetAtTimeMinute returns the AtTimeMinute field if non-nil, zero value otherwise.

### GetAtTimeMinuteOk

`func (o *HumanCron) GetAtTimeMinuteOk() (*string, bool)`

GetAtTimeMinuteOk returns a tuple with the AtTimeMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtTimeMinute

`func (o *HumanCron) SetAtTimeMinute(v string)`

SetAtTimeMinute sets AtTimeMinute field to given value.

### HasAtTimeMinute

`func (o *HumanCron) HasAtTimeMinute() bool`

HasAtTimeMinute returns a boolean if a field has been set.

### GetDays

`func (o *HumanCron) GetDays() []string`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *HumanCron) GetDaysOk() (*[]string, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *HumanCron) SetDays(v []string)`

SetDays sets Days field to given value.

### HasDays

`func (o *HumanCron) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetSchedulingType

`func (o *HumanCron) GetSchedulingType() string`

GetSchedulingType returns the SchedulingType field if non-nil, zero value otherwise.

### GetSchedulingTypeOk

`func (o *HumanCron) GetSchedulingTypeOk() (*string, bool)`

GetSchedulingTypeOk returns a tuple with the SchedulingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingType

`func (o *HumanCron) SetSchedulingType(v string)`

SetSchedulingType sets SchedulingType field to given value.

### HasSchedulingType

`func (o *HumanCron) HasSchedulingType() bool`

HasSchedulingType returns a boolean if a field has been set.

### GetTimePeriod

`func (o *HumanCron) GetTimePeriod() int64`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *HumanCron) GetTimePeriodOk() (*int64, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *HumanCron) SetTimePeriod(v int64)`

SetTimePeriod sets TimePeriod field to given value.

### HasTimePeriod

`func (o *HumanCron) HasTimePeriod() bool`

HasTimePeriod returns a boolean if a field has been set.

### GetTimePeriodType

`func (o *HumanCron) GetTimePeriodType() string`

GetTimePeriodType returns the TimePeriodType field if non-nil, zero value otherwise.

### GetTimePeriodTypeOk

`func (o *HumanCron) GetTimePeriodTypeOk() (*string, bool)`

GetTimePeriodTypeOk returns a tuple with the TimePeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriodType

`func (o *HumanCron) SetTimePeriodType(v string)`

SetTimePeriodType sets TimePeriodType field to given value.

### HasTimePeriodType

`func (o *HumanCron) HasTimePeriodType() bool`

HasTimePeriodType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


