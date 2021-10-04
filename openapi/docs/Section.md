# Section

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoPageBreak** | Pointer to **bool** |  | [optional] 
**Automation** | Pointer to [**ReportAutomation**](ReportAutomation.md) |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayType** | Pointer to **string** |  | [optional] 
**EmptyNotification** | Pointer to **string** |  | [optional] 
**FromDate** | Pointer to **string** |  | [optional] 
**Layout** | Pointer to **map[string]interface{}** |  | [optional] 
**Query** | Pointer to [**ReportQuery**](ReportQuery.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**TitleStyle** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**ToDate** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewSection

`func NewSection() *Section`

NewSection instantiates a new Section object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionWithDefaults

`func NewSectionWithDefaults() *Section`

NewSectionWithDefaults instantiates a new Section object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoPageBreak

`func (o *Section) GetAutoPageBreak() bool`

GetAutoPageBreak returns the AutoPageBreak field if non-nil, zero value otherwise.

### GetAutoPageBreakOk

`func (o *Section) GetAutoPageBreakOk() (*bool, bool)`

GetAutoPageBreakOk returns a tuple with the AutoPageBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPageBreak

`func (o *Section) SetAutoPageBreak(v bool)`

SetAutoPageBreak sets AutoPageBreak field to given value.

### HasAutoPageBreak

`func (o *Section) HasAutoPageBreak() bool`

HasAutoPageBreak returns a boolean if a field has been set.

### GetAutomation

`func (o *Section) GetAutomation() ReportAutomation`

GetAutomation returns the Automation field if non-nil, zero value otherwise.

### GetAutomationOk

`func (o *Section) GetAutomationOk() (*ReportAutomation, bool)`

GetAutomationOk returns a tuple with the Automation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomation

`func (o *Section) SetAutomation(v ReportAutomation)`

SetAutomation sets Automation field to given value.

### HasAutomation

`func (o *Section) HasAutomation() bool`

HasAutomation returns a boolean if a field has been set.

### GetData

`func (o *Section) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Section) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Section) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *Section) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *Section) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Section) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Section) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Section) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayType

`func (o *Section) GetDisplayType() string`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *Section) GetDisplayTypeOk() (*string, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *Section) SetDisplayType(v string)`

SetDisplayType sets DisplayType field to given value.

### HasDisplayType

`func (o *Section) HasDisplayType() bool`

HasDisplayType returns a boolean if a field has been set.

### GetEmptyNotification

`func (o *Section) GetEmptyNotification() string`

GetEmptyNotification returns the EmptyNotification field if non-nil, zero value otherwise.

### GetEmptyNotificationOk

`func (o *Section) GetEmptyNotificationOk() (*string, bool)`

GetEmptyNotificationOk returns a tuple with the EmptyNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyNotification

`func (o *Section) SetEmptyNotification(v string)`

SetEmptyNotification sets EmptyNotification field to given value.

### HasEmptyNotification

`func (o *Section) HasEmptyNotification() bool`

HasEmptyNotification returns a boolean if a field has been set.

### GetFromDate

`func (o *Section) GetFromDate() string`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *Section) GetFromDateOk() (*string, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *Section) SetFromDate(v string)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *Section) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetLayout

`func (o *Section) GetLayout() map[string]interface{}`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *Section) GetLayoutOk() (*map[string]interface{}, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *Section) SetLayout(v map[string]interface{})`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *Section) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetQuery

`func (o *Section) GetQuery() ReportQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Section) GetQueryOk() (*ReportQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Section) SetQuery(v ReportQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *Section) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTitle

`func (o *Section) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Section) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Section) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Section) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleStyle

`func (o *Section) GetTitleStyle() map[string]map[string]interface{}`

GetTitleStyle returns the TitleStyle field if non-nil, zero value otherwise.

### GetTitleStyleOk

`func (o *Section) GetTitleStyleOk() (*map[string]map[string]interface{}, bool)`

GetTitleStyleOk returns a tuple with the TitleStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleStyle

`func (o *Section) SetTitleStyle(v map[string]map[string]interface{})`

SetTitleStyle sets TitleStyle field to given value.

### HasTitleStyle

`func (o *Section) HasTitleStyle() bool`

HasTitleStyle returns a boolean if a field has been set.

### GetToDate

`func (o *Section) GetToDate() string`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *Section) GetToDateOk() (*string, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *Section) SetToDate(v string)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *Section) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### GetType

`func (o *Section) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Section) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Section) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Section) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


