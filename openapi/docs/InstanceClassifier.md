# InstanceClassifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brands** | Pointer to **[]string** |  | [optional] 
**CommitMessage** | Pointer to **string** |  | [optional] 
**DefaultIncidentType** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Feed** | Pointer to **bool** |  | [optional] 
**FromServerVersion** | Pointer to **string** |  | [optional] 
**Highlight** | Pointer to **map[string][]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IncidentSamples** | Pointer to **[]map[string]interface{}** | Incidents is a list of incident entities | [optional] 
**IndicatorSamples** | Pointer to [**[]FeedIndicator**](FeedIndicator.md) |  | [optional] 
**InstanceIds** | Pointer to **[]string** |  | [optional] 
**ItemVersion** | Pointer to **string** |  | [optional] 
**KeyTypeMap** | Pointer to **map[string]interface{}** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**LogicalVersion** | Pointer to **int64** |  | [optional] 
**Mapping** | Pointer to **map[string]interface{}** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NameRaw** | Pointer to **string** |  | [optional] 
**NumericId** | Pointer to **int64** |  | [optional] 
**PackID** | Pointer to **string** |  | [optional] 
**PackPropagationLabels** | Pointer to **[]string** |  | [optional] 
**PrevName** | Pointer to **string** |  | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**PropagationLabels** | Pointer to **[]string** |  | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**ShouldCommit** | Pointer to **bool** |  | [optional] 
**SortValues** | Pointer to **[]string** |  | [optional] 
**SourceClassifierId** | Pointer to **string** |  | [optional] 
**System** | Pointer to **bool** |  | [optional] 
**ToServerVersion** | Pointer to **string** |  | [optional] 
**Transformer** | Pointer to **map[string]interface{}** |  | [optional] 
**Type** | Pointer to **string** | MapperType for instance classifier | [optional] 
**UnclassifiedCases** | Pointer to **map[string]map[string]int64** |  | [optional] 
**VcShouldIgnore** | Pointer to **bool** |  | [optional] 
**VcShouldKeepItemLegacyProdMachine** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewInstanceClassifier

`func NewInstanceClassifier() *InstanceClassifier`

NewInstanceClassifier instantiates a new InstanceClassifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceClassifierWithDefaults

`func NewInstanceClassifierWithDefaults() *InstanceClassifier`

NewInstanceClassifierWithDefaults instantiates a new InstanceClassifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrands

`func (o *InstanceClassifier) GetBrands() []string`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *InstanceClassifier) GetBrandsOk() (*[]string, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *InstanceClassifier) SetBrands(v []string)`

SetBrands sets Brands field to given value.

### HasBrands

`func (o *InstanceClassifier) HasBrands() bool`

HasBrands returns a boolean if a field has been set.

### GetCommitMessage

`func (o *InstanceClassifier) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *InstanceClassifier) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *InstanceClassifier) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *InstanceClassifier) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetDefaultIncidentType

`func (o *InstanceClassifier) GetDefaultIncidentType() string`

GetDefaultIncidentType returns the DefaultIncidentType field if non-nil, zero value otherwise.

### GetDefaultIncidentTypeOk

`func (o *InstanceClassifier) GetDefaultIncidentTypeOk() (*string, bool)`

GetDefaultIncidentTypeOk returns a tuple with the DefaultIncidentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIncidentType

`func (o *InstanceClassifier) SetDefaultIncidentType(v string)`

SetDefaultIncidentType sets DefaultIncidentType field to given value.

### HasDefaultIncidentType

`func (o *InstanceClassifier) HasDefaultIncidentType() bool`

HasDefaultIncidentType returns a boolean if a field has been set.

### GetDescription

`func (o *InstanceClassifier) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstanceClassifier) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstanceClassifier) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstanceClassifier) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeed

`func (o *InstanceClassifier) GetFeed() bool`

GetFeed returns the Feed field if non-nil, zero value otherwise.

### GetFeedOk

`func (o *InstanceClassifier) GetFeedOk() (*bool, bool)`

GetFeedOk returns a tuple with the Feed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeed

`func (o *InstanceClassifier) SetFeed(v bool)`

SetFeed sets Feed field to given value.

### HasFeed

`func (o *InstanceClassifier) HasFeed() bool`

HasFeed returns a boolean if a field has been set.

### GetFromServerVersion

`func (o *InstanceClassifier) GetFromServerVersion() string`

GetFromServerVersion returns the FromServerVersion field if non-nil, zero value otherwise.

### GetFromServerVersionOk

`func (o *InstanceClassifier) GetFromServerVersionOk() (*string, bool)`

GetFromServerVersionOk returns a tuple with the FromServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromServerVersion

`func (o *InstanceClassifier) SetFromServerVersion(v string)`

SetFromServerVersion sets FromServerVersion field to given value.

### HasFromServerVersion

`func (o *InstanceClassifier) HasFromServerVersion() bool`

HasFromServerVersion returns a boolean if a field has been set.

### GetHighlight

`func (o *InstanceClassifier) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *InstanceClassifier) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *InstanceClassifier) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *InstanceClassifier) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetId

`func (o *InstanceClassifier) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceClassifier) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceClassifier) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceClassifier) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncidentSamples

`func (o *InstanceClassifier) GetIncidentSamples() []map[string]interface{}`

GetIncidentSamples returns the IncidentSamples field if non-nil, zero value otherwise.

### GetIncidentSamplesOk

`func (o *InstanceClassifier) GetIncidentSamplesOk() (*[]map[string]interface{}, bool)`

GetIncidentSamplesOk returns a tuple with the IncidentSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentSamples

`func (o *InstanceClassifier) SetIncidentSamples(v []map[string]interface{})`

SetIncidentSamples sets IncidentSamples field to given value.

### HasIncidentSamples

`func (o *InstanceClassifier) HasIncidentSamples() bool`

HasIncidentSamples returns a boolean if a field has been set.

### GetIndicatorSamples

`func (o *InstanceClassifier) GetIndicatorSamples() []FeedIndicator`

GetIndicatorSamples returns the IndicatorSamples field if non-nil, zero value otherwise.

### GetIndicatorSamplesOk

`func (o *InstanceClassifier) GetIndicatorSamplesOk() (*[]FeedIndicator, bool)`

GetIndicatorSamplesOk returns a tuple with the IndicatorSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatorSamples

`func (o *InstanceClassifier) SetIndicatorSamples(v []FeedIndicator)`

SetIndicatorSamples sets IndicatorSamples field to given value.

### HasIndicatorSamples

`func (o *InstanceClassifier) HasIndicatorSamples() bool`

HasIndicatorSamples returns a boolean if a field has been set.

### GetInstanceIds

`func (o *InstanceClassifier) GetInstanceIds() []string`

GetInstanceIds returns the InstanceIds field if non-nil, zero value otherwise.

### GetInstanceIdsOk

`func (o *InstanceClassifier) GetInstanceIdsOk() (*[]string, bool)`

GetInstanceIdsOk returns a tuple with the InstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceIds

`func (o *InstanceClassifier) SetInstanceIds(v []string)`

SetInstanceIds sets InstanceIds field to given value.

### HasInstanceIds

`func (o *InstanceClassifier) HasInstanceIds() bool`

HasInstanceIds returns a boolean if a field has been set.

### GetItemVersion

`func (o *InstanceClassifier) GetItemVersion() string`

GetItemVersion returns the ItemVersion field if non-nil, zero value otherwise.

### GetItemVersionOk

`func (o *InstanceClassifier) GetItemVersionOk() (*string, bool)`

GetItemVersionOk returns a tuple with the ItemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVersion

`func (o *InstanceClassifier) SetItemVersion(v string)`

SetItemVersion sets ItemVersion field to given value.

### HasItemVersion

`func (o *InstanceClassifier) HasItemVersion() bool`

HasItemVersion returns a boolean if a field has been set.

### GetKeyTypeMap

`func (o *InstanceClassifier) GetKeyTypeMap() map[string]interface{}`

GetKeyTypeMap returns the KeyTypeMap field if non-nil, zero value otherwise.

### GetKeyTypeMapOk

`func (o *InstanceClassifier) GetKeyTypeMapOk() (*map[string]interface{}, bool)`

GetKeyTypeMapOk returns a tuple with the KeyTypeMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTypeMap

`func (o *InstanceClassifier) SetKeyTypeMap(v map[string]interface{})`

SetKeyTypeMap sets KeyTypeMap field to given value.

### HasKeyTypeMap

`func (o *InstanceClassifier) HasKeyTypeMap() bool`

HasKeyTypeMap returns a boolean if a field has been set.

### GetLocked

`func (o *InstanceClassifier) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *InstanceClassifier) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *InstanceClassifier) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *InstanceClassifier) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetLogicalVersion

`func (o *InstanceClassifier) GetLogicalVersion() int64`

GetLogicalVersion returns the LogicalVersion field if non-nil, zero value otherwise.

### GetLogicalVersionOk

`func (o *InstanceClassifier) GetLogicalVersionOk() (*int64, bool)`

GetLogicalVersionOk returns a tuple with the LogicalVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalVersion

`func (o *InstanceClassifier) SetLogicalVersion(v int64)`

SetLogicalVersion sets LogicalVersion field to given value.

### HasLogicalVersion

`func (o *InstanceClassifier) HasLogicalVersion() bool`

HasLogicalVersion returns a boolean if a field has been set.

### GetMapping

`func (o *InstanceClassifier) GetMapping() map[string]interface{}`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *InstanceClassifier) GetMappingOk() (*map[string]interface{}, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *InstanceClassifier) SetMapping(v map[string]interface{})`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *InstanceClassifier) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### GetModified

`func (o *InstanceClassifier) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *InstanceClassifier) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *InstanceClassifier) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *InstanceClassifier) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *InstanceClassifier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceClassifier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceClassifier) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceClassifier) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameRaw

`func (o *InstanceClassifier) GetNameRaw() string`

GetNameRaw returns the NameRaw field if non-nil, zero value otherwise.

### GetNameRawOk

`func (o *InstanceClassifier) GetNameRawOk() (*string, bool)`

GetNameRawOk returns a tuple with the NameRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameRaw

`func (o *InstanceClassifier) SetNameRaw(v string)`

SetNameRaw sets NameRaw field to given value.

### HasNameRaw

`func (o *InstanceClassifier) HasNameRaw() bool`

HasNameRaw returns a boolean if a field has been set.

### GetNumericId

`func (o *InstanceClassifier) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *InstanceClassifier) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *InstanceClassifier) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *InstanceClassifier) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetPackID

`func (o *InstanceClassifier) GetPackID() string`

GetPackID returns the PackID field if non-nil, zero value otherwise.

### GetPackIDOk

`func (o *InstanceClassifier) GetPackIDOk() (*string, bool)`

GetPackIDOk returns a tuple with the PackID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackID

`func (o *InstanceClassifier) SetPackID(v string)`

SetPackID sets PackID field to given value.

### HasPackID

`func (o *InstanceClassifier) HasPackID() bool`

HasPackID returns a boolean if a field has been set.

### GetPackPropagationLabels

`func (o *InstanceClassifier) GetPackPropagationLabels() []string`

GetPackPropagationLabels returns the PackPropagationLabels field if non-nil, zero value otherwise.

### GetPackPropagationLabelsOk

`func (o *InstanceClassifier) GetPackPropagationLabelsOk() (*[]string, bool)`

GetPackPropagationLabelsOk returns a tuple with the PackPropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackPropagationLabels

`func (o *InstanceClassifier) SetPackPropagationLabels(v []string)`

SetPackPropagationLabels sets PackPropagationLabels field to given value.

### HasPackPropagationLabels

`func (o *InstanceClassifier) HasPackPropagationLabels() bool`

HasPackPropagationLabels returns a boolean if a field has been set.

### GetPrevName

`func (o *InstanceClassifier) GetPrevName() string`

GetPrevName returns the PrevName field if non-nil, zero value otherwise.

### GetPrevNameOk

`func (o *InstanceClassifier) GetPrevNameOk() (*string, bool)`

GetPrevNameOk returns a tuple with the PrevName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevName

`func (o *InstanceClassifier) SetPrevName(v string)`

SetPrevName sets PrevName field to given value.

### HasPrevName

`func (o *InstanceClassifier) HasPrevName() bool`

HasPrevName returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *InstanceClassifier) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *InstanceClassifier) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *InstanceClassifier) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *InstanceClassifier) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetPropagationLabels

`func (o *InstanceClassifier) GetPropagationLabels() []string`

GetPropagationLabels returns the PropagationLabels field if non-nil, zero value otherwise.

### GetPropagationLabelsOk

`func (o *InstanceClassifier) GetPropagationLabelsOk() (*[]string, bool)`

GetPropagationLabelsOk returns a tuple with the PropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationLabels

`func (o *InstanceClassifier) SetPropagationLabels(v []string)`

SetPropagationLabels sets PropagationLabels field to given value.

### HasPropagationLabels

`func (o *InstanceClassifier) HasPropagationLabels() bool`

HasPropagationLabels returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *InstanceClassifier) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *InstanceClassifier) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *InstanceClassifier) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *InstanceClassifier) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetShouldCommit

`func (o *InstanceClassifier) GetShouldCommit() bool`

GetShouldCommit returns the ShouldCommit field if non-nil, zero value otherwise.

### GetShouldCommitOk

`func (o *InstanceClassifier) GetShouldCommitOk() (*bool, bool)`

GetShouldCommitOk returns a tuple with the ShouldCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCommit

`func (o *InstanceClassifier) SetShouldCommit(v bool)`

SetShouldCommit sets ShouldCommit field to given value.

### HasShouldCommit

`func (o *InstanceClassifier) HasShouldCommit() bool`

HasShouldCommit returns a boolean if a field has been set.

### GetSortValues

`func (o *InstanceClassifier) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *InstanceClassifier) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *InstanceClassifier) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *InstanceClassifier) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetSourceClassifierId

`func (o *InstanceClassifier) GetSourceClassifierId() string`

GetSourceClassifierId returns the SourceClassifierId field if non-nil, zero value otherwise.

### GetSourceClassifierIdOk

`func (o *InstanceClassifier) GetSourceClassifierIdOk() (*string, bool)`

GetSourceClassifierIdOk returns a tuple with the SourceClassifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClassifierId

`func (o *InstanceClassifier) SetSourceClassifierId(v string)`

SetSourceClassifierId sets SourceClassifierId field to given value.

### HasSourceClassifierId

`func (o *InstanceClassifier) HasSourceClassifierId() bool`

HasSourceClassifierId returns a boolean if a field has been set.

### GetSystem

`func (o *InstanceClassifier) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *InstanceClassifier) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *InstanceClassifier) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *InstanceClassifier) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetToServerVersion

`func (o *InstanceClassifier) GetToServerVersion() string`

GetToServerVersion returns the ToServerVersion field if non-nil, zero value otherwise.

### GetToServerVersionOk

`func (o *InstanceClassifier) GetToServerVersionOk() (*string, bool)`

GetToServerVersionOk returns a tuple with the ToServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToServerVersion

`func (o *InstanceClassifier) SetToServerVersion(v string)`

SetToServerVersion sets ToServerVersion field to given value.

### HasToServerVersion

`func (o *InstanceClassifier) HasToServerVersion() bool`

HasToServerVersion returns a boolean if a field has been set.

### GetTransformer

`func (o *InstanceClassifier) GetTransformer() map[string]interface{}`

GetTransformer returns the Transformer field if non-nil, zero value otherwise.

### GetTransformerOk

`func (o *InstanceClassifier) GetTransformerOk() (*map[string]interface{}, bool)`

GetTransformerOk returns a tuple with the Transformer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformer

`func (o *InstanceClassifier) SetTransformer(v map[string]interface{})`

SetTransformer sets Transformer field to given value.

### HasTransformer

`func (o *InstanceClassifier) HasTransformer() bool`

HasTransformer returns a boolean if a field has been set.

### GetType

`func (o *InstanceClassifier) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceClassifier) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceClassifier) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InstanceClassifier) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnclassifiedCases

`func (o *InstanceClassifier) GetUnclassifiedCases() map[string]map[string]int64`

GetUnclassifiedCases returns the UnclassifiedCases field if non-nil, zero value otherwise.

### GetUnclassifiedCasesOk

`func (o *InstanceClassifier) GetUnclassifiedCasesOk() (*map[string]map[string]int64, bool)`

GetUnclassifiedCasesOk returns a tuple with the UnclassifiedCases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnclassifiedCases

`func (o *InstanceClassifier) SetUnclassifiedCases(v map[string]map[string]int64)`

SetUnclassifiedCases sets UnclassifiedCases field to given value.

### HasUnclassifiedCases

`func (o *InstanceClassifier) HasUnclassifiedCases() bool`

HasUnclassifiedCases returns a boolean if a field has been set.

### GetVcShouldIgnore

`func (o *InstanceClassifier) GetVcShouldIgnore() bool`

GetVcShouldIgnore returns the VcShouldIgnore field if non-nil, zero value otherwise.

### GetVcShouldIgnoreOk

`func (o *InstanceClassifier) GetVcShouldIgnoreOk() (*bool, bool)`

GetVcShouldIgnoreOk returns a tuple with the VcShouldIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldIgnore

`func (o *InstanceClassifier) SetVcShouldIgnore(v bool)`

SetVcShouldIgnore sets VcShouldIgnore field to given value.

### HasVcShouldIgnore

`func (o *InstanceClassifier) HasVcShouldIgnore() bool`

HasVcShouldIgnore returns a boolean if a field has been set.

### GetVcShouldKeepItemLegacyProdMachine

`func (o *InstanceClassifier) GetVcShouldKeepItemLegacyProdMachine() bool`

GetVcShouldKeepItemLegacyProdMachine returns the VcShouldKeepItemLegacyProdMachine field if non-nil, zero value otherwise.

### GetVcShouldKeepItemLegacyProdMachineOk

`func (o *InstanceClassifier) GetVcShouldKeepItemLegacyProdMachineOk() (*bool, bool)`

GetVcShouldKeepItemLegacyProdMachineOk returns a tuple with the VcShouldKeepItemLegacyProdMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcShouldKeepItemLegacyProdMachine

`func (o *InstanceClassifier) SetVcShouldKeepItemLegacyProdMachine(v bool)`

SetVcShouldKeepItemLegacyProdMachine sets VcShouldKeepItemLegacyProdMachine field to given value.

### HasVcShouldKeepItemLegacyProdMachine

`func (o *InstanceClassifier) HasVcShouldKeepItemLegacyProdMachine() bool`

HasVcShouldKeepItemLegacyProdMachine returns a boolean if a field has been set.

### GetVersion

`func (o *InstanceClassifier) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InstanceClassifier) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InstanceClassifier) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InstanceClassifier) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


