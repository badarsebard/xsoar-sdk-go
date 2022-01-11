# TaskCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to [**[]ArgFilter**](ArgFilter.md) |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskCondition

`func NewTaskCondition() *TaskCondition`

NewTaskCondition instantiates a new TaskCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskConditionWithDefaults

`func NewTaskConditionWithDefaults() *TaskCondition`

NewTaskConditionWithDefaults instantiates a new TaskCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *TaskCondition) GetCondition() []ArgFilter`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *TaskCondition) GetConditionOk() (*[]ArgFilter, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *TaskCondition) SetCondition(v []ArgFilter)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *TaskCondition) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetLabel

`func (o *TaskCondition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TaskCondition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TaskCondition) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *TaskCondition) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


