# TimerTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**FieldName** | Pointer to **string** |  | [optional] 

## Methods

### NewTimerTrigger

`func NewTimerTrigger() *TimerTrigger`

NewTimerTrigger instantiates a new TimerTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimerTriggerWithDefaults

`func NewTimerTriggerWithDefaults() *TimerTrigger`

NewTimerTriggerWithDefaults instantiates a new TimerTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *TimerTrigger) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TimerTrigger) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TimerTrigger) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *TimerTrigger) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetFieldName

`func (o *TimerTrigger) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *TimerTrigger) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *TimerTrigger) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *TimerTrigger) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


