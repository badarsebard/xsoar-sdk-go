# IndicatorContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | Pointer to **string** |  | [optional] 
**Indicator** | Pointer to [**IocObject**](IocObject.md) |  | [optional] 
**InvestigationId** | Pointer to **string** |  | [optional] 
**Manually** | Pointer to **bool** |  | [optional] 
**SeenNow** | Pointer to **bool** |  | [optional] 

## Methods

### NewIndicatorContext

`func NewIndicatorContext() *IndicatorContext`

NewIndicatorContext instantiates a new IndicatorContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndicatorContextWithDefaults

`func NewIndicatorContextWithDefaults() *IndicatorContext`

NewIndicatorContextWithDefaults instantiates a new IndicatorContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *IndicatorContext) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *IndicatorContext) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *IndicatorContext) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *IndicatorContext) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetIndicator

`func (o *IndicatorContext) GetIndicator() IocObject`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *IndicatorContext) GetIndicatorOk() (*IocObject, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *IndicatorContext) SetIndicator(v IocObject)`

SetIndicator sets Indicator field to given value.

### HasIndicator

`func (o *IndicatorContext) HasIndicator() bool`

HasIndicator returns a boolean if a field has been set.

### GetInvestigationId

`func (o *IndicatorContext) GetInvestigationId() string`

GetInvestigationId returns the InvestigationId field if non-nil, zero value otherwise.

### GetInvestigationIdOk

`func (o *IndicatorContext) GetInvestigationIdOk() (*string, bool)`

GetInvestigationIdOk returns a tuple with the InvestigationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationId

`func (o *IndicatorContext) SetInvestigationId(v string)`

SetInvestigationId sets InvestigationId field to given value.

### HasInvestigationId

`func (o *IndicatorContext) HasInvestigationId() bool`

HasInvestigationId returns a boolean if a field has been set.

### GetManually

`func (o *IndicatorContext) GetManually() bool`

GetManually returns the Manually field if non-nil, zero value otherwise.

### GetManuallyOk

`func (o *IndicatorContext) GetManuallyOk() (*bool, bool)`

GetManuallyOk returns a tuple with the Manually field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManually

`func (o *IndicatorContext) SetManually(v bool)`

SetManually sets Manually field to given value.

### HasManually

`func (o *IndicatorContext) HasManually() bool`

HasManually returns a boolean if a field has been set.

### GetSeenNow

`func (o *IndicatorContext) GetSeenNow() bool`

GetSeenNow returns the SeenNow field if non-nil, zero value otherwise.

### GetSeenNowOk

`func (o *IndicatorContext) GetSeenNowOk() (*bool, bool)`

GetSeenNowOk returns a tuple with the SeenNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenNow

`func (o *IndicatorContext) SetSeenNow(v bool)`

SetSeenNow sets SeenNow field to given value.

### HasSeenNow

`func (o *IndicatorContext) HasSeenNow() bool`

HasSeenNow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


