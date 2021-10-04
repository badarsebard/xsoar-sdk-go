# LayoutSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to [**[]LayoutField**](LayoutField.md) |  | [optional] 
**Highlight** | Pointer to **map[string][]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsVisible** | Pointer to **bool** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumericId** | Pointer to **int64** |  | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**Query** | Pointer to **map[string]interface{}** |  | [optional] 
**QueryType** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**SortValues** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewLayoutSection

`func NewLayoutSection() *LayoutSection`

NewLayoutSection instantiates a new LayoutSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayoutSectionWithDefaults

`func NewLayoutSectionWithDefaults() *LayoutSection`

NewLayoutSectionWithDefaults instantiates a new LayoutSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *LayoutSection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LayoutSection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LayoutSection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LayoutSection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFields

`func (o *LayoutSection) GetFields() []LayoutField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *LayoutSection) GetFieldsOk() (*[]LayoutField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *LayoutSection) SetFields(v []LayoutField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *LayoutSection) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetHighlight

`func (o *LayoutSection) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *LayoutSection) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *LayoutSection) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *LayoutSection) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetId

`func (o *LayoutSection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LayoutSection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LayoutSection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LayoutSection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsVisible

`func (o *LayoutSection) GetIsVisible() bool`

GetIsVisible returns the IsVisible field if non-nil, zero value otherwise.

### GetIsVisibleOk

`func (o *LayoutSection) GetIsVisibleOk() (*bool, bool)`

GetIsVisibleOk returns a tuple with the IsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisible

`func (o *LayoutSection) SetIsVisible(v bool)`

SetIsVisible sets IsVisible field to given value.

### HasIsVisible

`func (o *LayoutSection) HasIsVisible() bool`

HasIsVisible returns a boolean if a field has been set.

### GetModified

`func (o *LayoutSection) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *LayoutSection) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *LayoutSection) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *LayoutSection) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *LayoutSection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LayoutSection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LayoutSection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LayoutSection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumericId

`func (o *LayoutSection) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *LayoutSection) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *LayoutSection) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *LayoutSection) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *LayoutSection) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *LayoutSection) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *LayoutSection) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *LayoutSection) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetQuery

`func (o *LayoutSection) GetQuery() map[string]interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *LayoutSection) GetQueryOk() (*map[string]interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *LayoutSection) SetQuery(v map[string]interface{})`

SetQuery sets Query field to given value.

### HasQuery

`func (o *LayoutSection) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetQueryType

`func (o *LayoutSection) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *LayoutSection) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *LayoutSection) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.

### HasQueryType

`func (o *LayoutSection) HasQueryType() bool`

HasQueryType returns a boolean if a field has been set.

### GetReadOnly

`func (o *LayoutSection) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *LayoutSection) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *LayoutSection) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *LayoutSection) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *LayoutSection) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *LayoutSection) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *LayoutSection) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *LayoutSection) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetSortValues

`func (o *LayoutSection) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *LayoutSection) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *LayoutSection) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *LayoutSection) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetType

`func (o *LayoutSection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LayoutSection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LayoutSection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LayoutSection) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *LayoutSection) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LayoutSection) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LayoutSection) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LayoutSection) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


