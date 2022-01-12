# Widget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category the widget is related to. Used to display in widget library under category or dataType if empty. | [optional] 
**CommitMessage** | Pointer to **string** |  | [optional] 
**DataType** | Pointer to **string** | Data type of the widget. Describes what data does the widget query. supporting data types \&quot;incidents\&quot;,\&quot;messages\&quot;,\&quot;system\&quot;,\&quot;entries\&quot;,\&quot;tasks\&quot;, \&quot;audit\&quot;. | [optional] 
**DateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**Description** | Pointer to **string** | The description of the widget&#39;s usage and data representation. | [optional] 
**FromServerVersion** | Pointer to **string** |  | [optional] 
**Highlight** | Pointer to **map[string][]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsPredefined** | Pointer to **bool** | Is the widget a system widget. | [optional] 
**ItemVersion** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** | Is the widget locked for editing. | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** | Default name of the widget. | 
**NumericId** | Pointer to **int64** |  | [optional] 
**PackID** | Pointer to **string** |  | [optional] 
**PackPropagationLabels** | Pointer to **[]string** |  | [optional] 
**Params** | Pointer to **map[string]map[string]interface{}** | Additional parameters for this widget, depends on widget type and data. | [optional] 
**PrevName** | Pointer to **string** | The previous name of the widget. | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**PropagationLabels** | Pointer to **[]string** |  | [optional] 
**Query** | Pointer to **string** | Query to search on the dataType. | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**ShouldCommit** | Pointer to **bool** |  | [optional] 
**Size** | Pointer to **int64** | Maximum size for this widget data returned. | [optional] 
**Sort** | Pointer to [**[]Order**](Order.md) | Sorting array to sort the data received by the given Order parameters. | [optional] 
**SortValues** | Pointer to **[]string** |  | [optional] 
**ToServerVersion** | Pointer to **string** |  | [optional] 
**VcShouldIgnore** | Pointer to **bool** |  | [optional] 
**VcShouldKeepItemLegacyProdMachine** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**WidgetType** | **string** | Widget type describes how does the widget should recieve the data, and display it. Supporting types: \&quot;bar\&quot;, \&quot;column\&quot;, \&quot;pie\&quot;, \&quot;list\&quot;, \&quot;number\&quot;, \&quot;trend\&quot;, \&quot;text\&quot;, \&quot;duration\&quot;, \&quot;image\&quot;, \&quot;line\&quot;, and \&quot;table\&quot;. | 

## Methods

### NewWidget

`func NewWidget(name string, widgetType string, ) *Widget`

NewWidget instantiates a new Widget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetWithDefaults

`func NewWidgetWithDefaults() *Widget`

NewWidgetWithDefaults instantiates a new Widget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *Widget) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Widget) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Widget) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Widget) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCommitMessage

`func (o *Widget) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *Widget) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *Widget) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *Widget) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetDataType

`func (o *Widget) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *Widget) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *Widget) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *Widget) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetDateRange

`func (o *Widget) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *Widget) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *Widget) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *Widget) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetDescription

`func (o *Widget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Widget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Widget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Widget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFromServerVersion

`func (o *Widget) GetFromServerVersion() string`

GetFromServerVersion returns the FromServerVersion field if non-nil, zero value otherwise.

### GetFromServerVersionOk

`func (o *Widget) GetFromServerVersionOk() (*string, bool)`

GetFromServerVersionOk returns a tuple with the FromServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromServerVersion

`func (o *Widget) SetFromServerVersion(v string)`

SetFromServerVersion sets FromServerVersion field to given value.

### HasFromServerVersion

`func (o *Widget) HasFromServerVersion() bool`

HasFromServerVersion returns a boolean if a field has been set.

### GetHighlight

`func (o *Widget) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *Widget) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *Widget) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *Widget) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetId

`func (o *Widget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Widget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Widget) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Widget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPredefined

`func (o *Widget) GetIsPredefined() bool`

GetIsPredefined returns the IsPredefined field if non-nil, zero value otherwise.

### GetIsPredefinedOk

`func (o *Widget) GetIsPredefinedOk() (*bool, bool)`

GetIsPredefinedOk returns a tuple with the IsPredefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPredefined

`func (o *Widget) SetIsPredefined(v bool)`

SetIsPredefined sets IsPredefined field to given value.

### HasIsPredefined

`func (o *Widget) HasIsPredefined() bool`

HasIsPredefined returns a boolean if a field has been set.

### GetItemVersion

`func (o *Widget) GetItemVersion() string`

GetItemVersion returns the ItemVersion field if non-nil, zero value otherwise.

### GetItemVersionOk

`func (o *Widget) GetItemVersionOk() (*string, bool)`

GetItemVersionOk returns a tuple with the ItemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVersion

`func (o *Widget) SetItemVersion(v string)`

SetItemVersion sets ItemVersion field to given value.

### HasItemVersion

`func (o *Widget) HasItemVersion() bool`

HasItemVersion returns a boolean if a field has been set.

### GetLocked

`func (o *Widget) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Widget) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Widget) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *Widget) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetModified

`func (o *Widget) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Widget) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Widget) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Widget) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *Widget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Widget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Widget) SetName(v string)`

SetName sets Name field to given value.


### GetNumericId

`func (o *Widget) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *Widget) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *Widget) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *Widget) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetPackID

`func (o *Widget) GetPackID() string`

GetPackID returns the PackID field if non-nil, zero value otherwise.

### GetPackIDOk

`func (o *Widget) GetPackIDOk() (*string, bool)`

GetPackIDOk returns a tuple with the PackID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackID

`func (o *Widget) SetPackID(v string)`

SetPackID sets PackID field to given value.

### HasPackID

`func (o *Widget) HasPackID() bool`

HasPackID returns a boolean if a field has been set.

### GetPackPropagationLabels

`func (o *Widget) GetPackPropagationLabels() []string`

GetPackPropagationLabels returns the PackPropagationLabels field if non-nil, zero value otherwise.

### GetPackPropagationLabelsOk

`func (o *Widget) GetPackPropagationLabelsOk() (*[]string, bool)`

GetPackPropagationLabelsOk returns a tuple with the PackPropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackPropagationLabels

`func (o *Widget) SetPackPropagationLabels(v []string)`

SetPackPropagationLabels sets PackPropagationLabels field to given value.

### HasPackPropagationLabels

`func (o *Widget) HasPackPropagationLabels() bool`

HasPackPropagationLabels returns a boolean if a field has been set.

### GetParams

`func (o *Widget) GetParams() map[string]map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Widget) GetParamsOk() (*map[string]map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Widget) SetParams(v map[string]map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *Widget) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetPrevName

`func (o *Widget) GetPrevName() string`

GetPrevName returns the PrevName field if non-nil, zero value otherwise.

### GetPrevNameOk

`func (o *Widget) GetPrevNameOk() (*string, bool)`

GetPrevNameOk returns a tuple with the PrevName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevName

`func (o *Widget) SetPrevName(v string)`

SetPrevName sets PrevName field to given value.

### HasPrevName

`func (o *Widget) HasPrevName() bool`

HasPrevName returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *Widget) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *Widget) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *Widget) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *Widget) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetPropagationLabels

`func (o *Widget) GetPropagationLabels() []string`

GetPropagationLabels returns the PropagationLabels field if non-nil, zero value otherwise.

### GetPropagationLabelsOk

`func (o *Widget) GetPropagationLabelsOk() (*[]string, bool)`

GetPropagationLabelsOk returns a tuple with the PropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationLabels

`func (o *Widget) SetPropagationLabels(v []string)`

SetPropagationLabels sets PropagationLabels field to given value.

### HasPropagationLabels

`func (o *Widget) HasPropagationLabels() bool`

HasPropagationLabels returns a boolean if a field has been set.

### GetQuery

`func (o *Widget) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Widget) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Widget) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *Widget) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *Widget) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *Widget) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *Widget) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *Widget) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetShouldCommit

`func (o *Widget) GetShouldCommit() bool`

GetShouldCommit returns the ShouldCommit field if non-nil, zero value otherwise.

### GetShouldCommitOk

`func (o *Widget) GetShouldCommitOk() (*bool, bool)`

GetShouldCommitOk returns a tuple with the ShouldCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCommit

`func (o *Widget) SetShouldCommit(v bool)`

SetShouldCommit sets ShouldCommit field to given value.

### HasShouldCommit

`func (o *Widget) HasShouldCommit() bool`

HasShouldCommit returns a boolean if a field has been set.

### GetSize

`func (o *Widget) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Widget) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Widget) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *Widget) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSort

`func (o *Widget) GetSort() []Order`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *Widget) GetSortOk() (*[]Order, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *Widget) SetSort(v []Order)`

SetSort sets Sort field to given value.

### HasSort

`func (o *Widget) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetSortValues

`func (o *Widget) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *Widget) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *Widget) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *Widget) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetToServerVersion

`func (o *Widget) GetToServerVersion() string`

GetToServerVersion returns the ToServerVersion field if non-nil, zero value otherwise.

### GetToServerVersionOk

`func (o *Widget) GetToServerVersionOk() (*string, bool)`

GetToServerVersionOk returns a tuple with the ToServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToServerVersion

`func (o *Widget) SetToServerVersion(v string)`

SetToServerVersion sets ToServerVersion field to given value.

### HasToServerVersion

`func (o *Widget) HasToServerVersion() bool`

HasToServerVersion returns a boolean if a field has been set.

### GetVcShouldIgnore

`func (o *Widget) GetVcShouldIgnore() bool`

GetVcShouldIgnore returns the VcShouldIgnore field if non-nil, zero value otherwise.

### GetVcShouldIgnoreOk

`func (o *Widget) GetVcShouldIgnoreOk() (*bool, bool)`

GetVcShouldIgnoreOk returns a tuple with the VcShouldIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldIgnore

`func (o *Widget) SetVcShouldIgnore(v bool)`

SetVcShouldIgnore sets VcShouldIgnore field to given value.

### HasVcShouldIgnore

`func (o *Widget) HasVcShouldIgnore() bool`

HasVcShouldIgnore returns a boolean if a field has been set.

### GetVcShouldKeepItemLegacyProdMachine

`func (o *Widget) GetVcShouldKeepItemLegacyProdMachine() bool`

GetVcShouldKeepItemLegacyProdMachine returns the VcShouldKeepItemLegacyProdMachine field if non-nil, zero value otherwise.

### GetVcShouldKeepItemLegacyProdMachineOk

`func (o *Widget) GetVcShouldKeepItemLegacyProdMachineOk() (*bool, bool)`

GetVcShouldKeepItemLegacyProdMachineOk returns a tuple with the VcShouldKeepItemLegacyProdMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldKeepItemLegacyProdMachine

`func (o *Widget) SetVcShouldKeepItemLegacyProdMachine(v bool)`

SetVcShouldKeepItemLegacyProdMachine sets VcShouldKeepItemLegacyProdMachine field to given value.

### HasVcShouldKeepItemLegacyProdMachine

`func (o *Widget) HasVcShouldKeepItemLegacyProdMachine() bool`

HasVcShouldKeepItemLegacyProdMachine returns a boolean if a field has been set.

### GetVersion

`func (o *Widget) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Widget) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Widget) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Widget) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWidgetType

`func (o *Widget) GetWidgetType() string`

GetWidgetType returns the WidgetType field if non-nil, zero value otherwise.

### GetWidgetTypeOk

`func (o *Widget) GetWidgetTypeOk() (*string, bool)`

GetWidgetTypeOk returns a tuple with the WidgetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetType

`func (o *Widget) SetWidgetType(v string)`

SetWidgetType sets WidgetType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


