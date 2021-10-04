# WhitelistedIndicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Highlight** | Pointer to **map[string][]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**NumericId** | Pointer to **int64** |  | [optional] 
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Reputations** | Pointer to **[]string** |  | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**SortValues** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**WhitelistTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewWhitelistedIndicator

`func NewWhitelistedIndicator() *WhitelistedIndicator`

NewWhitelistedIndicator instantiates a new WhitelistedIndicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhitelistedIndicatorWithDefaults

`func NewWhitelistedIndicatorWithDefaults() *WhitelistedIndicator`

NewWhitelistedIndicatorWithDefaults instantiates a new WhitelistedIndicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHighlight

`func (o *WhitelistedIndicator) GetHighlight() map[string][]string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *WhitelistedIndicator) GetHighlightOk() (*map[string][]string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *WhitelistedIndicator) SetHighlight(v map[string][]string)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *WhitelistedIndicator) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetId

`func (o *WhitelistedIndicator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhitelistedIndicator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhitelistedIndicator) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WhitelistedIndicator) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocked

`func (o *WhitelistedIndicator) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *WhitelistedIndicator) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *WhitelistedIndicator) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *WhitelistedIndicator) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetModified

`func (o *WhitelistedIndicator) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *WhitelistedIndicator) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *WhitelistedIndicator) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *WhitelistedIndicator) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetNumericId

`func (o *WhitelistedIndicator) GetNumericId() int64`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *WhitelistedIndicator) GetNumericIdOk() (*int64, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *WhitelistedIndicator) SetNumericId(v int64)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *WhitelistedIndicator) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetPrimaryTerm

`func (o *WhitelistedIndicator) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *WhitelistedIndicator) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *WhitelistedIndicator) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *WhitelistedIndicator) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetReason

`func (o *WhitelistedIndicator) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WhitelistedIndicator) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WhitelistedIndicator) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *WhitelistedIndicator) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReputations

`func (o *WhitelistedIndicator) GetReputations() []string`

GetReputations returns the Reputations field if non-nil, zero value otherwise.

### GetReputationsOk

`func (o *WhitelistedIndicator) GetReputationsOk() (*[]string, bool)`

GetReputationsOk returns a tuple with the Reputations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReputations

`func (o *WhitelistedIndicator) SetReputations(v []string)`

SetReputations sets Reputations field to given value.

### HasReputations

`func (o *WhitelistedIndicator) HasReputations() bool`

HasReputations returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *WhitelistedIndicator) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *WhitelistedIndicator) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *WhitelistedIndicator) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *WhitelistedIndicator) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetSortValues

`func (o *WhitelistedIndicator) GetSortValues() []string`

GetSortValues returns the SortValues field if non-nil, zero value otherwise.

### GetSortValuesOk

`func (o *WhitelistedIndicator) GetSortValuesOk() (*[]string, bool)`

GetSortValuesOk returns a tuple with the SortValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValues

`func (o *WhitelistedIndicator) SetSortValues(v []string)`

SetSortValues sets SortValues field to given value.

### HasSortValues

`func (o *WhitelistedIndicator) HasSortValues() bool`

HasSortValues returns a boolean if a field has been set.

### GetType

`func (o *WhitelistedIndicator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhitelistedIndicator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhitelistedIndicator) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WhitelistedIndicator) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *WhitelistedIndicator) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *WhitelistedIndicator) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *WhitelistedIndicator) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *WhitelistedIndicator) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetValue

`func (o *WhitelistedIndicator) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WhitelistedIndicator) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WhitelistedIndicator) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WhitelistedIndicator) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVersion

`func (o *WhitelistedIndicator) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WhitelistedIndicator) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WhitelistedIndicator) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WhitelistedIndicator) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWhitelistTime

`func (o *WhitelistedIndicator) GetWhitelistTime() time.Time`

GetWhitelistTime returns the WhitelistTime field if non-nil, zero value otherwise.

### GetWhitelistTimeOk

`func (o *WhitelistedIndicator) GetWhitelistTimeOk() (*time.Time, bool)`

GetWhitelistTimeOk returns a tuple with the WhitelistTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistTime

`func (o *WhitelistedIndicator) SetWhitelistTime(v time.Time)`

SetWhitelistTime sets WhitelistTime field to given value.

### HasWhitelistTime

`func (o *WhitelistedIndicator) HasWhitelistTime() bool`

HasWhitelistTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


