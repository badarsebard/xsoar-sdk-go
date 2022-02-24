# InvTaskDebug

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BreakpointCondition** | Pointer to [**[]ArgAtomicFilter**](ArgAtomicFilter.md) | ArgFilter - represent a slice of atomic filters with OR condition between them (e.i. - atomic1 OR atomic2 OR ...) | [optional] 
**InputOverrides** | Pointer to **map[string]string** |  | [optional] 
**IsMarkedWithBreakpoint** | Pointer to **bool** |  | [optional] 
**IsMarkedWithSkip** | Pointer to **bool** |  | [optional] 
**OutputOverrides** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**ShouldOverrideDebugInfo** | Pointer to **bool** |  | [optional] 
**SkipConditionOverrideTo** | Pointer to **string** |  | [optional] 
**TaskOriginalId** | Pointer to **string** |  | [optional] 

## Methods

### NewInvTaskDebug

`func NewInvTaskDebug() *InvTaskDebug`

NewInvTaskDebug instantiates a new InvTaskDebug object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvTaskDebugWithDefaults

`func NewInvTaskDebugWithDefaults() *InvTaskDebug`

NewInvTaskDebugWithDefaults instantiates a new InvTaskDebug object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakpointCondition

`func (o *InvTaskDebug) GetBreakpointCondition() []ArgAtomicFilter`

GetBreakpointCondition returns the BreakpointCondition field if non-nil, zero value otherwise.

### GetBreakpointConditionOk

`func (o *InvTaskDebug) GetBreakpointConditionOk() (*[]ArgAtomicFilter, bool)`

GetBreakpointConditionOk returns a tuple with the BreakpointCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakpointCondition

`func (o *InvTaskDebug) SetBreakpointCondition(v []ArgAtomicFilter)`

SetBreakpointCondition sets BreakpointCondition field to given value.

### HasBreakpointCondition

`func (o *InvTaskDebug) HasBreakpointCondition() bool`

HasBreakpointCondition returns a boolean if a field has been set.

### GetInputOverrides

`func (o *InvTaskDebug) GetInputOverrides() map[string]string`

GetInputOverrides returns the InputOverrides field if non-nil, zero value otherwise.

### GetInputOverridesOk

`func (o *InvTaskDebug) GetInputOverridesOk() (*map[string]string, bool)`

GetInputOverridesOk returns a tuple with the InputOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputOverrides

`func (o *InvTaskDebug) SetInputOverrides(v map[string]string)`

SetInputOverrides sets InputOverrides field to given value.

### HasInputOverrides

`func (o *InvTaskDebug) HasInputOverrides() bool`

HasInputOverrides returns a boolean if a field has been set.

### GetIsMarkedWithBreakpoint

`func (o *InvTaskDebug) GetIsMarkedWithBreakpoint() bool`

GetIsMarkedWithBreakpoint returns the IsMarkedWithBreakpoint field if non-nil, zero value otherwise.

### GetIsMarkedWithBreakpointOk

`func (o *InvTaskDebug) GetIsMarkedWithBreakpointOk() (*bool, bool)`

GetIsMarkedWithBreakpointOk returns a tuple with the IsMarkedWithBreakpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarkedWithBreakpoint

`func (o *InvTaskDebug) SetIsMarkedWithBreakpoint(v bool)`

SetIsMarkedWithBreakpoint sets IsMarkedWithBreakpoint field to given value.

### HasIsMarkedWithBreakpoint

`func (o *InvTaskDebug) HasIsMarkedWithBreakpoint() bool`

HasIsMarkedWithBreakpoint returns a boolean if a field has been set.

### GetIsMarkedWithSkip

`func (o *InvTaskDebug) GetIsMarkedWithSkip() bool`

GetIsMarkedWithSkip returns the IsMarkedWithSkip field if non-nil, zero value otherwise.

### GetIsMarkedWithSkipOk

`func (o *InvTaskDebug) GetIsMarkedWithSkipOk() (*bool, bool)`

GetIsMarkedWithSkipOk returns a tuple with the IsMarkedWithSkip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarkedWithSkip

`func (o *InvTaskDebug) SetIsMarkedWithSkip(v bool)`

SetIsMarkedWithSkip sets IsMarkedWithSkip field to given value.

### HasIsMarkedWithSkip

`func (o *InvTaskDebug) HasIsMarkedWithSkip() bool`

HasIsMarkedWithSkip returns a boolean if a field has been set.

### GetOutputOverrides

`func (o *InvTaskDebug) GetOutputOverrides() map[string]map[string]interface{}`

GetOutputOverrides returns the OutputOverrides field if non-nil, zero value otherwise.

### GetOutputOverridesOk

`func (o *InvTaskDebug) GetOutputOverridesOk() (*map[string]map[string]interface{}, bool)`

GetOutputOverridesOk returns a tuple with the OutputOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputOverrides

`func (o *InvTaskDebug) SetOutputOverrides(v map[string]map[string]interface{})`

SetOutputOverrides sets OutputOverrides field to given value.

### HasOutputOverrides

`func (o *InvTaskDebug) HasOutputOverrides() bool`

HasOutputOverrides returns a boolean if a field has been set.

### GetShouldOverrideDebugInfo

`func (o *InvTaskDebug) GetShouldOverrideDebugInfo() bool`

GetShouldOverrideDebugInfo returns the ShouldOverrideDebugInfo field if non-nil, zero value otherwise.

### GetShouldOverrideDebugInfoOk

`func (o *InvTaskDebug) GetShouldOverrideDebugInfoOk() (*bool, bool)`

GetShouldOverrideDebugInfoOk returns a tuple with the ShouldOverrideDebugInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldOverrideDebugInfo

`func (o *InvTaskDebug) SetShouldOverrideDebugInfo(v bool)`

SetShouldOverrideDebugInfo sets ShouldOverrideDebugInfo field to given value.

### HasShouldOverrideDebugInfo

`func (o *InvTaskDebug) HasShouldOverrideDebugInfo() bool`

HasShouldOverrideDebugInfo returns a boolean if a field has been set.

### GetSkipConditionOverrideTo

`func (o *InvTaskDebug) GetSkipConditionOverrideTo() string`

GetSkipConditionOverrideTo returns the SkipConditionOverrideTo field if non-nil, zero value otherwise.

### GetSkipConditionOverrideToOk

`func (o *InvTaskDebug) GetSkipConditionOverrideToOk() (*string, bool)`

GetSkipConditionOverrideToOk returns a tuple with the SkipConditionOverrideTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipConditionOverrideTo

`func (o *InvTaskDebug) SetSkipConditionOverrideTo(v string)`

SetSkipConditionOverrideTo sets SkipConditionOverrideTo field to given value.

### HasSkipConditionOverrideTo

`func (o *InvTaskDebug) HasSkipConditionOverrideTo() bool`

HasSkipConditionOverrideTo returns a boolean if a field has been set.

### GetTaskOriginalId

`func (o *InvTaskDebug) GetTaskOriginalId() string`

GetTaskOriginalId returns the TaskOriginalId field if non-nil, zero value otherwise.

### GetTaskOriginalIdOk

`func (o *InvTaskDebug) GetTaskOriginalIdOk() (*string, bool)`

GetTaskOriginalIdOk returns a tuple with the TaskOriginalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOriginalId

`func (o *InvTaskDebug) SetTaskOriginalId(v string)`

SetTaskOriginalId sets TaskOriginalId field to given value.

### HasTaskOriginalId

`func (o *InvTaskDebug) HasTaskOriginalId() bool`

HasTaskOriginalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


