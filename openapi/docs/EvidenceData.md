# EvidenceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFields** | Pointer to [**map[string]AdvanceArg**](AdvanceArg.md) | This field must have empty json key | [optional] 
**Description** | Pointer to [**AdvanceArg**](AdvanceArg.md) |  | [optional] 
**Occurred** | Pointer to [**AdvanceArg**](AdvanceArg.md) |  | [optional] 
**Tags** | Pointer to [**AdvanceArg**](AdvanceArg.md) |  | [optional] 

## Methods

### NewEvidenceData

`func NewEvidenceData() *EvidenceData`

NewEvidenceData instantiates a new EvidenceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvidenceDataWithDefaults

`func NewEvidenceDataWithDefaults() *EvidenceData`

NewEvidenceDataWithDefaults instantiates a new EvidenceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomFields

`func (o *EvidenceData) GetCustomFields() map[string]AdvanceArg`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *EvidenceData) GetCustomFieldsOk() (*map[string]AdvanceArg, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *EvidenceData) SetCustomFields(v map[string]AdvanceArg)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *EvidenceData) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDescription

`func (o *EvidenceData) GetDescription() AdvanceArg`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EvidenceData) GetDescriptionOk() (*AdvanceArg, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EvidenceData) SetDescription(v AdvanceArg)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EvidenceData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOccurred

`func (o *EvidenceData) GetOccurred() AdvanceArg`

GetOccurred returns the Occurred field if non-nil, zero value otherwise.

### GetOccurredOk

`func (o *EvidenceData) GetOccurredOk() (*AdvanceArg, bool)`

GetOccurredOk returns a tuple with the Occurred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurred

`func (o *EvidenceData) SetOccurred(v AdvanceArg)`

SetOccurred sets Occurred field to given value.

### HasOccurred

`func (o *EvidenceData) HasOccurred() bool`

HasOccurred returns a boolean if a field has been set.

### GetTags

`func (o *EvidenceData) GetTags() AdvanceArg`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EvidenceData) GetTagsOk() (*AdvanceArg, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EvidenceData) SetTags(v AdvanceArg)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EvidenceData) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


