# PlaybookInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**PlaybookInputQuery** | Pointer to [**PlaybookInputQuery**](PlaybookInputQuery.md) |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to [**AdvanceArg**](AdvanceArg.md) |  | [optional] 

## Methods

### NewPlaybookInput

`func NewPlaybookInput() *PlaybookInput`

NewPlaybookInput instantiates a new PlaybookInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaybookInputWithDefaults

`func NewPlaybookInputWithDefaults() *PlaybookInput`

NewPlaybookInputWithDefaults instantiates a new PlaybookInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PlaybookInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlaybookInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlaybookInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlaybookInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *PlaybookInput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PlaybookInput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PlaybookInput) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PlaybookInput) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPlaybookInputQuery

`func (o *PlaybookInput) GetPlaybookInputQuery() PlaybookInputQuery`

GetPlaybookInputQuery returns the PlaybookInputQuery field if non-nil, zero value otherwise.

### GetPlaybookInputQueryOk

`func (o *PlaybookInput) GetPlaybookInputQueryOk() (*PlaybookInputQuery, bool)`

GetPlaybookInputQueryOk returns a tuple with the PlaybookInputQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookInputQuery

`func (o *PlaybookInput) SetPlaybookInputQuery(v PlaybookInputQuery)`

SetPlaybookInputQuery sets PlaybookInputQuery field to given value.

### HasPlaybookInputQuery

`func (o *PlaybookInput) HasPlaybookInputQuery() bool`

HasPlaybookInputQuery returns a boolean if a field has been set.

### GetRequired

`func (o *PlaybookInput) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *PlaybookInput) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *PlaybookInput) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *PlaybookInput) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetValue

`func (o *PlaybookInput) GetValue() AdvanceArg`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PlaybookInput) GetValueOk() (*AdvanceArg, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PlaybookInput) SetValue(v AdvanceArg)`

SetValue sets Value field to given value.

### HasValue

`func (o *PlaybookInput) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


