# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cron** | Pointer to **string** |  | [optional] 
**CronView** | Pointer to **bool** |  | [optional] 
**EndingDate** | Pointer to **time.Time** |  | [optional] 
**EndingType** | Pointer to **string** | EndingType holds the type of schedule Ending | [optional] 
**HumanCron** | Pointer to [**HumanCron**](HumanCron.md) |  | [optional] 
**Recurrent** | Pointer to **bool** |  | [optional] 
**Scheduled** | Pointer to **bool** | is it scheduled | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**Times** | Pointer to **int64** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**TimezoneOffset** | Pointer to **int64** |  | [optional] 

## Methods

### NewSchedule

`func NewSchedule() *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCron

`func (o *Schedule) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *Schedule) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *Schedule) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *Schedule) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetCronView

`func (o *Schedule) GetCronView() bool`

GetCronView returns the CronView field if non-nil, zero value otherwise.

### GetCronViewOk

`func (o *Schedule) GetCronViewOk() (*bool, bool)`

GetCronViewOk returns a tuple with the CronView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronView

`func (o *Schedule) SetCronView(v bool)`

SetCronView sets CronView field to given value.

### HasCronView

`func (o *Schedule) HasCronView() bool`

HasCronView returns a boolean if a field has been set.

### GetEndingDate

`func (o *Schedule) GetEndingDate() time.Time`

GetEndingDate returns the EndingDate field if non-nil, zero value otherwise.

### GetEndingDateOk

`func (o *Schedule) GetEndingDateOk() (*time.Time, bool)`

GetEndingDateOk returns a tuple with the EndingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingDate

`func (o *Schedule) SetEndingDate(v time.Time)`

SetEndingDate sets EndingDate field to given value.

### HasEndingDate

`func (o *Schedule) HasEndingDate() bool`

HasEndingDate returns a boolean if a field has been set.

### GetEndingType

`func (o *Schedule) GetEndingType() string`

GetEndingType returns the EndingType field if non-nil, zero value otherwise.

### GetEndingTypeOk

`func (o *Schedule) GetEndingTypeOk() (*string, bool)`

GetEndingTypeOk returns a tuple with the EndingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingType

`func (o *Schedule) SetEndingType(v string)`

SetEndingType sets EndingType field to given value.

### HasEndingType

`func (o *Schedule) HasEndingType() bool`

HasEndingType returns a boolean if a field has been set.

### GetHumanCron

`func (o *Schedule) GetHumanCron() HumanCron`

GetHumanCron returns the HumanCron field if non-nil, zero value otherwise.

### GetHumanCronOk

`func (o *Schedule) GetHumanCronOk() (*HumanCron, bool)`

GetHumanCronOk returns a tuple with the HumanCron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumanCron

`func (o *Schedule) SetHumanCron(v HumanCron)`

SetHumanCron sets HumanCron field to given value.

### HasHumanCron

`func (o *Schedule) HasHumanCron() bool`

HasHumanCron returns a boolean if a field has been set.

### GetRecurrent

`func (o *Schedule) GetRecurrent() bool`

GetRecurrent returns the Recurrent field if non-nil, zero value otherwise.

### GetRecurrentOk

`func (o *Schedule) GetRecurrentOk() (*bool, bool)`

GetRecurrentOk returns a tuple with the Recurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrent

`func (o *Schedule) SetRecurrent(v bool)`

SetRecurrent sets Recurrent field to given value.

### HasRecurrent

`func (o *Schedule) HasRecurrent() bool`

HasRecurrent returns a boolean if a field has been set.

### GetScheduled

`func (o *Schedule) GetScheduled() bool`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *Schedule) GetScheduledOk() (*bool, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *Schedule) SetScheduled(v bool)`

SetScheduled sets Scheduled field to given value.

### HasScheduled

`func (o *Schedule) HasScheduled() bool`

HasScheduled returns a boolean if a field has been set.

### GetStartDate

`func (o *Schedule) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Schedule) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Schedule) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Schedule) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTimes

`func (o *Schedule) GetTimes() int64`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *Schedule) GetTimesOk() (*int64, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *Schedule) SetTimes(v int64)`

SetTimes sets Times field to given value.

### HasTimes

`func (o *Schedule) HasTimes() bool`

HasTimes returns a boolean if a field has been set.

### GetTimezone

`func (o *Schedule) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Schedule) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Schedule) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Schedule) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTimezoneOffset

`func (o *Schedule) GetTimezoneOffset() int64`

GetTimezoneOffset returns the TimezoneOffset field if non-nil, zero value otherwise.

### GetTimezoneOffsetOk

`func (o *Schedule) GetTimezoneOffsetOk() (*int64, bool)`

GetTimezoneOffsetOk returns a tuple with the TimezoneOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneOffset

`func (o *Schedule) SetTimezoneOffset(v int64)`

SetTimezoneOffset sets TimezoneOffset field to given value.

### HasTimezoneOffset

`func (o *Schedule) HasTimezoneOffset() bool`

HasTimezoneOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


