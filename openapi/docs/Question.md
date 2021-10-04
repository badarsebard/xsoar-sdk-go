# Question

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to [**[]GridColumn**](GridColumn.md) |  | [optional] 
**DefaultRows** | Pointer to **[]map[string]map[string]interface{}** |  | [optional] 
**FieldAssociated** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**LabelArg** | Pointer to [**AdvanceArg**](AdvanceArg.md) |  | [optional] 
**Options** | Pointer to **[]string** |  | [optional] 
**OptionsArg** | Pointer to [**[]AdvanceArg**](AdvanceArg.md) |  | [optional] 
**Placeholder** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Tooltip** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewQuestion

`func NewQuestion() *Question`

NewQuestion instantiates a new Question object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionWithDefaults

`func NewQuestionWithDefaults() *Question`

NewQuestionWithDefaults instantiates a new Question object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *Question) GetColumns() []GridColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *Question) GetColumnsOk() (*[]GridColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *Question) SetColumns(v []GridColumn)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *Question) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetDefaultRows

`func (o *Question) GetDefaultRows() []map[string]map[string]interface{}`

GetDefaultRows returns the DefaultRows field if non-nil, zero value otherwise.

### GetDefaultRowsOk

`func (o *Question) GetDefaultRowsOk() (*[]map[string]map[string]interface{}, bool)`

GetDefaultRowsOk returns a tuple with the DefaultRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRows

`func (o *Question) SetDefaultRows(v []map[string]map[string]interface{})`

SetDefaultRows sets DefaultRows field to given value.

### HasDefaultRows

`func (o *Question) HasDefaultRows() bool`

HasDefaultRows returns a boolean if a field has been set.

### GetFieldAssociated

`func (o *Question) GetFieldAssociated() string`

GetFieldAssociated returns the FieldAssociated field if non-nil, zero value otherwise.

### GetFieldAssociatedOk

`func (o *Question) GetFieldAssociatedOk() (*string, bool)`

GetFieldAssociatedOk returns a tuple with the FieldAssociated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldAssociated

`func (o *Question) SetFieldAssociated(v string)`

SetFieldAssociated sets FieldAssociated field to given value.

### HasFieldAssociated

`func (o *Question) HasFieldAssociated() bool`

HasFieldAssociated returns a boolean if a field has been set.

### GetId

`func (o *Question) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Question) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Question) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Question) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *Question) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Question) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Question) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Question) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLabelArg

`func (o *Question) GetLabelArg() AdvanceArg`

GetLabelArg returns the LabelArg field if non-nil, zero value otherwise.

### GetLabelArgOk

`func (o *Question) GetLabelArgOk() (*AdvanceArg, bool)`

GetLabelArgOk returns a tuple with the LabelArg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelArg

`func (o *Question) SetLabelArg(v AdvanceArg)`

SetLabelArg sets LabelArg field to given value.

### HasLabelArg

`func (o *Question) HasLabelArg() bool`

HasLabelArg returns a boolean if a field has been set.

### GetOptions

`func (o *Question) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Question) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Question) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Question) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOptionsArg

`func (o *Question) GetOptionsArg() []AdvanceArg`

GetOptionsArg returns the OptionsArg field if non-nil, zero value otherwise.

### GetOptionsArgOk

`func (o *Question) GetOptionsArgOk() (*[]AdvanceArg, bool)`

GetOptionsArgOk returns a tuple with the OptionsArg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsArg

`func (o *Question) SetOptionsArg(v []AdvanceArg)`

SetOptionsArg sets OptionsArg field to given value.

### HasOptionsArg

`func (o *Question) HasOptionsArg() bool`

HasOptionsArg returns a boolean if a field has been set.

### GetPlaceholder

`func (o *Question) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *Question) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *Question) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *Question) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetReadOnly

`func (o *Question) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *Question) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *Question) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *Question) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRequired

`func (o *Question) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *Question) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *Question) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *Question) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetTooltip

`func (o *Question) GetTooltip() string`

GetTooltip returns the Tooltip field if non-nil, zero value otherwise.

### GetTooltipOk

`func (o *Question) GetTooltipOk() (*string, bool)`

GetTooltipOk returns a tuple with the Tooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltip

`func (o *Question) SetTooltip(v string)`

SetTooltip sets Tooltip field to given value.

### HasTooltip

`func (o *Question) HasTooltip() bool`

HasTooltip returns a boolean if a field has been set.

### GetType

`func (o *Question) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Question) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Question) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Question) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


