# ProcessInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int64** | A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years. | [optional] 
**HighPriority** | Pointer to **bool** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProcessInfo

`func NewProcessInfo() *ProcessInfo`

NewProcessInfo instantiates a new ProcessInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessInfoWithDefaults

`func NewProcessInfoWithDefaults() *ProcessInfo`

NewProcessInfoWithDefaults instantiates a new ProcessInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *ProcessInfo) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ProcessInfo) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ProcessInfo) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ProcessInfo) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetDuration

`func (o *ProcessInfo) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ProcessInfo) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ProcessInfo) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ProcessInfo) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetHighPriority

`func (o *ProcessInfo) GetHighPriority() bool`

GetHighPriority returns the HighPriority field if non-nil, zero value otherwise.

### GetHighPriorityOk

`func (o *ProcessInfo) GetHighPriorityOk() (*bool, bool)`

GetHighPriorityOk returns a tuple with the HighPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPriority

`func (o *ProcessInfo) SetHighPriority(v bool)`

SetHighPriority sets HighPriority field to given value.

### HasHighPriority

`func (o *ProcessInfo) HasHighPriority() bool`

HasHighPriority returns a boolean if a field has been set.

### GetStartedAt

`func (o *ProcessInfo) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ProcessInfo) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ProcessInfo) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ProcessInfo) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


