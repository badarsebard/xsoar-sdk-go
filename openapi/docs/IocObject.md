# IocObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFields** | Pointer to **map[string]map[string]interface{}** | The keys should be the field&#39;s display name all lower and without spaces. For example: Scan IP -&gt; scanip To get the actual key name you can also go to Cortex XSOAR CLI and run /incident_add and look for the key that you would like to update | [optional] 
**Account** | Pointer to **string** |  | [optional] 
**AggregatedReliability** | Pointer to **string** |  | [optional] 
**CalculatedTime** | Pointer to **time.Time** | Do not set the fields bellow this line | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to [**[]Comment**](Comment.md) |  | [optional] 
**DeletedFeedFetchTime** | Pointer to **time.Time** |  | [optional] 
**Expiration** | Pointer to **time.Time** |  | [optional] 
**ExpirationSource** | Pointer to [**ExpirationSource**](ExpirationSource.md) |  | [optional] 
**ExpirationStatus** | Pointer to **string** |  | [optional] 
**FirstSeen** | Pointer to **time.Time** |  | [optional] 
**FirstSeenEntryID** | Pointer to **string** |  | [optional] 
**Highlight** | Pointer to **map[string][]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IndicatorType** | Pointer to **string** |  | [optional] 
**InsightCache** | Pointer to [**InsightCache**](InsightCache.md) |  | [optional] 
**InvestigationIDs** | Pointer to **[]string** |  | [optional] 
**IsShared** | Pointer to **bool** |  | [optional] 
**LastReputationRun** | Pointer to **time.Time** |  | [optional] 
**LastSeen** | Pointer to **time.Time** |  | [optional] 
**LastSeenEntryID** | Pointer to **string** |  | [optional] 
**ManualExpirationTime** | Pointer to **time.Time** |  | [optional] 
**ManualScore** | Pointer to **bool** |  | [optional] 
**ManualSetTime** | Pointer to **time.Time** |  | [optional] 
**ManuallyEditedFields** | Pointer to **[]string** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**ModifiedTime** | Pointer to **time.Time** |  | [optional] 
**ModuleToFeedMap** | Pointer to [**map[string]FeedIndicator**](FeedIndicator.md) |  | [optional] 
**NumericId** | Pointer to **int64** |  | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**RelatedIncCount** | Pointer to **int64** |  | [optional] 
**Score** | Pointer to **int64** |  | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**SetBy** | Pointer to **string** |  | [optional] 
**SortValues** | Pointer to **[]string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceBrands** | Pointer to **[]string** |  | [optional] 
**SourceInstances** | Pointer to **[]string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewIocObject

`func NewIocObject() *IocObject`

NewIocObject instantiates a new IocObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIocObjectWithDefaults

`func NewIocObjectWithDefaults() *IocObject`

NewIocObjectWithDefaults instantiates a new IocObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomFields

`func (o *IocObject) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *IocObject) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *IocObject) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *IocObject) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetAccount

`func (o *IocObject) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IocObject) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IocObject) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IocObject) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAggregatedReliability

`func (o *IocObject) GetAggregatedReliability() string`

GetAggregatedReliability returns the AggregatedReliability field if non-nil, zero value otherwise.

### GetAggregatedReliabilityOk

`func (o *IocObject) GetAggregatedReliabilityOk() (*string, bool)`

GetAggregatedReliabilityOk returns a tuple with the AggregatedReliability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedReliability

`func (o *IocObject) SetAggregatedReliability(v string)`

SetAggregatedReliability sets AggregatedReliability field to given value.

### HasAggregatedReliability

`func (o *IocObject) HasAggregatedReliability() bool`

HasAggregatedReliability returns a boolean if a field has been set.

### GetCalculatedTime

`func (o *IocObject) GetCalculatedTime() time.Time`

GetCalculatedTime returns the CalculatedTime field if non-nil, zero value otherwise.

### GetCalculatedTimeOk

`func (o *IocObject) GetCalculatedTimeOk() (*time.Time, bool)`

GetCalculatedTimeOk returns a tuple with the CalculatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedTime

`func (o *IocObject) SetCalculatedTime(v time.Time)`

SetCalculatedTime sets CalculatedTime field to given value.

### HasCalculatedTime

`func (o *IocObject) HasCalculatedTime() bool`

HasCalculatedTime returns a boolean if a field has been set.

### GetComment

`func (o *IocObject) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IocObject) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IocObject) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IocObject) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetComments

`func (o *IocObject) GetComments() []Comment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *IocObject) GetCommentsOk() (*[]Comment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *IocObject) SetComments(v []Comment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *IocObject) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetDeletedFeedFetchTime

`func (o *IocObject) GetDeletedFeedFetchTime() time.Time`

GetDeletedFeedFetchTime returns the DeletedFeedFetchTime field if non-nil, zero value otherwise.

### GetDeletedFeedFetchTimeOk

`func (o *IocObject) GetDeletedFeedFetchTimeOk() (*time.Time, bool)`

GetDeletedFeedFetchTimeOk returns a tuple with the DeletedFeedFetchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedFeedFetchTime

`func (o *IocObject) SetDeletedFeedFetchTime(v time.Time)`

SetDeletedFeedFetchTime sets DeletedFeedFetchTime field to given value.

### HasDeletedFeedFetchTime

`func (o *IocObject) HasDeletedFeedFetchTime() bool`

HasDeletedFeedFetchTime returns a boolean if a field has been set.

### GetExpiration

`func (o *IocObject) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *IocObject) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *IocObject) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *IocObject) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetExpirationSource

`func (o *IocObject) GetExpirationSource() ExpirationSource`

GetExpirationSource returns the ExpirationSource field if non-nil, zero value otherwise.

### GetExpirationSourceOk

`func (o *IocObject) GetExpirationSourceOk() (*ExpirationSource, bool)`

GetExpirationSourceOk returns a tuple with the ExpirationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationSource

`func (o *IocObject) SetExpirationSource(v ExpirationSource)`

SetExpirationSource sets ExpirationSource field to given value.

### HasExpirationSource

`func (o *IocObject) HasExpirationSource() bool`

HasExpirationSource returns a boolean if a field has been set.

### GetExpirationStatus

`func (o *IocObject) GetExpirationStatus() string`

GetExpirationStatus returns the ExpirationStatus field if non-nil, zero value otherwise.

### GetExpirationStatusOk

`func (o *IocObject) GetExpirationStatusOk() (*string, bool)`

GetExpirationStatusOk returns a tuple with the ExpirationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationStatus

`func (o *IocObject) SetExpirationStatus(v string)`

SetExpirationStatus sets ExpirationStatus field to given value.

### HasExpirationStatus

`func (o *IocObject) HasExpirationStatus() bool`

HasExpirationStatus returns a boolean if a field has been set.

### GetFirstSeen

`func (o *IocObject) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *IocObject) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *IocObject) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *IocObject) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetFirstSeenEntryID

`func (o *IocObject) GetFirstSeenEntryID() string`

GetFirstSeenEntryID returns the FirstSeenEntryID field if non-nil, zero value otherwise.

### GetFirstSeenEntryIDOk

`func (o *IocObject) GetFirstSeenEntryIDOk() (*string, bool)`

GetFirstSeenEntryIDOk returns a tuple with the FirstSeenEntryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenEntryID

`func (o *IocObject) SetFirstSeenEntryID(v string)`

SetFirstSeenEntryID sets FirstSeenEntryID field to given value.

### HasFirstSeenEntryID

`func (o *IocObject) HasFirstSeenEntryID() bool`

HasFirstSeenEntryID returns a boolean if a field has been set.

### GetHighlight

`func (o *IocObject) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *IocObject) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *IocObject) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *IocObject) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetId

`func (o *IocObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IocObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IocObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IocObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndicatorType

`func (o *IocObject) GetIndicatorType() string`

GetIndicatorType returns the IndicatorType field if non-nil, zero value otherwise.

### GetIndicatorTypeOk

`func (o *IocObject) GetIndicatorTypeOk() (*string, bool)`

GetIndicatorTypeOk returns a tuple with the IndicatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatorType

`func (o *IocObject) SetIndicatorType(v string)`

SetIndicatorType sets IndicatorType field to given value.

### HasIndicatorType

`func (o *IocObject) HasIndicatorType() bool`

HasIndicatorType returns a boolean if a field has been set.

### GetInsightCache

`func (o *IocObject) GetInsightCache() InsightCache`

GetInsightCache returns the InsightCache field if non-nil, zero value otherwise.

### GetInsightCacheOk

`func (o *IocObject) GetInsightCacheOk() (*InsightCache, bool)`

GetInsightCacheOk returns a tuple with the InsightCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightCache

`func (o *IocObject) SetInsightCache(v InsightCache)`

SetInsightCache sets InsightCache field to given value.

### HasInsightCache

`func (o *IocObject) HasInsightCache() bool`

HasInsightCache returns a boolean if a field has been set.

### GetInvestigationIDs

`func (o *IocObject) GetInvestigationIDs() []string`

GetInvestigationIDs returns the InvestigationIDs field if non-nil, zero value otherwise.

### GetInvestigationIDsOk

`func (o *IocObject) GetInvestigationIDsOk() (*[]string, bool)`

GetInvestigationIDsOk returns a tuple with the InvestigationIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigationIDs

`func (o *IocObject) SetInvestigationIDs(v []string)`

SetInvestigationIDs sets InvestigationIDs field to given value.

### HasInvestigationIDs

`func (o *IocObject) HasInvestigationIDs() bool`

HasInvestigationIDs returns a boolean if a field has been set.

### GetIsShared

`func (o *IocObject) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *IocObject) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *IocObject) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *IocObject) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetLastReputationRun

`func (o *IocObject) GetLastReputationRun() time.Time`

GetLastReputationRun returns the LastReputationRun field if non-nil, zero value otherwise.

### GetLastReputationRunOk

`func (o *IocObject) GetLastReputationRunOk() (*time.Time, bool)`

GetLastReputationRunOk returns a tuple with the LastReputationRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReputationRun

`func (o *IocObject) SetLastReputationRun(v time.Time)`

SetLastReputationRun sets LastReputationRun field to given value.

### HasLastReputationRun

`func (o *IocObject) HasLastReputationRun() bool`

HasLastReputationRun returns a boolean if a field has been set.

### GetLastSeen

`func (o *IocObject) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *IocObject) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *IocObject) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *IocObject) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLastSeenEntryID

`func (o *IocObject) GetLastSeenEntryID() string`

GetLastSeenEntryID returns the LastSeenEntryID field if non-nil, zero value otherwise.

### GetLastSeenEntryIDOk

`func (o *IocObject) GetLastSeenEntryIDOk() (*string, bool)`

GetLastSeenEntryIDOk returns a tuple with the LastSeenEntryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenEntryID

`func (o *IocObject) SetLastSeenEntryID(v string)`

SetLastSeenEntryID sets LastSeenEntryID field to given value.

### HasLastSeenEntryID

`func (o *IocObject) HasLastSeenEntryID() bool`

HasLastSeenEntryID returns a boolean if a field has been set.

### GetManualExpirationTime

`func (o *IocObject) GetManualExpirationTime() time.Time`

GetManualExpirationTime returns the ManualExpirationTime field if non-nil, zero value otherwise.

### GetManualExpirationTimeOk

`func (o *IocObject) GetManualExpirationTimeOk() (*time.Time, bool)`

GetManualExpirationTimeOk returns a tuple with the ManualExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualExpirationTime

`func (o *IocObject) SetManualExpirationTime(v time.Time)`

SetManualExpirationTime sets ManualExpirationTime field to given value.

### HasManualExpirationTime

`func (o *IocObject) HasManualExpirationTime() bool`

HasManualExpirationTime returns a boolean if a field has been set.

### GetManualScore

`func (o *IocObject) GetManualScore() bool`

GetManualScore returns the ManualScore field if non-nil, zero value otherwise.

### GetManualScoreOk

`func (o *IocObject) GetManualScoreOk() (*bool, bool)`

GetManualScoreOk returns a tuple with the ManualScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualScore

`func (o *IocObject) SetManualScore(v bool)`

SetManualScore sets ManualScore field to given value.

### HasManualScore

`func (o *IocObject) HasManualScore() bool`

HasManualScore returns a boolean if a field has been set.

### GetManualSetTime

`func (o *IocObject) GetManualSetTime() time.Time`

GetManualSetTime returns the ManualSetTime field if non-nil, zero value otherwise.

### GetManualSetTimeOk

`func (o *IocObject) GetManualSetTimeOk() (*time.Time, bool)`

GetManualSetTimeOk returns a tuple with the ManualSetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualSetTime

`func (o *IocObject) SetManualSetTime(v time.Time)`

SetManualSetTime sets ManualSetTime field to given value.

### HasManualSetTime

`func (o *IocObject) HasManualSetTime() bool`

HasManualSetTime returns a boolean if a field has been set.

### GetManuallyEditedFields

`func (o *IocObject) GetManuallyEditedFields() []string`

GetManuallyEditedFields returns the ManuallyEditedFields field if non-nil, zero value otherwise.

### GetManuallyEditedFieldsOk

`func (o *IocObject) GetManuallyEditedFieldsOk() (*[]string, bool)`

GetManuallyEditedFieldsOk returns a tuple with the ManuallyEditedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyEditedFields

`func (o *IocObject) SetManuallyEditedFields(v []string)`

SetManuallyEditedFields sets ManuallyEditedFields field to given value.

### HasManuallyEditedFields

`func (o *IocObject) HasManuallyEditedFields() bool`

HasManuallyEditedFields returns a boolean if a field has been set.

### GetModified

`func (o *IocObject) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *IocObject) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *IocObject) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *IocObject) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetModifiedTime

`func (o *IocObject) GetModifiedTime() time.Time`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *IocObject) GetModifiedTimeOk() (*time.Time, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *IocObject) SetModifiedTime(v time.Time)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *IocObject) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetModuleToFeedMap

`func (o *IocObject) GetModuleToFeedMap() map[string]FeedIndicator`

GetModuleToFeedMap returns the ModuleToFeedMap field if non-nil, zero value otherwise.

### GetModuleToFeedMapOk

`func (o *IocObject) GetModuleToFeedMapOk() (*map[string]FeedIndicator, bool)`

GetModuleToFeedMapOk returns a tuple with the ModuleToFeedMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleToFeedMap

`func (o *IocObject) SetModuleToFeedMap(v map[string]FeedIndicator)`

SetModuleToFeedMap sets ModuleToFeedMap field to given value.

### HasModuleToFeedMap

`func (o *IocObject) HasModuleToFeedMap() bool`

HasModuleToFeedMap returns a boolean if a field has been set.

### GetNumericId

`func (o *IocObject) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *IocObject) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *IocObject) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *IocObject) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *IocObject) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *IocObject) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *IocObject) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *IocObject) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetRelatedIncCount

`func (o *IocObject) GetRelatedIncCount() int64`

GetRelatedIncCount returns the RelatedIncCount field if non-nil, zero value otherwise.

### GetRelatedIncCountOk

`func (o *IocObject) GetRelatedIncCountOk() (*int64, bool)`

GetRelatedIncCountOk returns a tuple with the RelatedIncCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedIncCount

`func (o *IocObject) SetRelatedIncCount(v int64)`

SetRelatedIncCount sets RelatedIncCount field to given value.

### HasRelatedIncCount

`func (o *IocObject) HasRelatedIncCount() bool`

HasRelatedIncCount returns a boolean if a field has been set.

### GetScore

`func (o *IocObject) GetScore() int64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *IocObject) GetScoreOk() (*int64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *IocObject) SetScore(v int64)`

SetScore sets Score field to given value.

### HasScore

`func (o *IocObject) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *IocObject) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *IocObject) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *IocObject) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *IocObject) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetSetBy

`func (o *IocObject) GetSetBy() string`

GetSetBy returns the SetBy field if non-nil, zero value otherwise.

### GetSetByOk

`func (o *IocObject) GetSetByOk() (*string, bool)`

GetSetByOk returns a tuple with the SetBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetBy

`func (o *IocObject) SetSetBy(v string)`

SetSetBy sets SetBy field to given value.

### HasSetBy

`func (o *IocObject) HasSetBy() bool`

HasSetBy returns a boolean if a field has been set.

### GetSortValues

`func (o *IocObject) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *IocObject) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *IocObject) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *IocObject) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetSource

`func (o *IocObject) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IocObject) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IocObject) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *IocObject) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceBrands

`func (o *IocObject) GetSourceBrands() []string`

GetSourceBrands returns the SourceBrands field if non-nil, zero value otherwise.

### GetSourceBrandsOk

`func (o *IocObject) GetSourceBrandsOk() (*[]string, bool)`

GetSourceBrandsOk returns a tuple with the SourceBrands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBrands

`func (o *IocObject) SetSourceBrands(v []string)`

SetSourceBrands sets SourceBrands field to given value.

### HasSourceBrands

`func (o *IocObject) HasSourceBrands() bool`

HasSourceBrands returns a boolean if a field has been set.

### GetSourceInstances

`func (o *IocObject) GetSourceInstances() []string`

GetSourceInstances returns the SourceInstances field if non-nil, zero value otherwise.

### GetSourceInstancesOk

`func (o *IocObject) GetSourceInstancesOk() (*[]string, bool)`

GetSourceInstancesOk returns a tuple with the SourceInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInstances

`func (o *IocObject) SetSourceInstances(v []string)`

SetSourceInstances sets SourceInstances field to given value.

### HasSourceInstances

`func (o *IocObject) HasSourceInstances() bool`

HasSourceInstances returns a boolean if a field has been set.

### GetTimestamp

`func (o *IocObject) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *IocObject) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *IocObject) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *IocObject) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetValue

`func (o *IocObject) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IocObject) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IocObject) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IocObject) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVersion

`func (o *IocObject) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IocObject) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IocObject) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IocObject) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


