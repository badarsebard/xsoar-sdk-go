# EntryHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentDate** | Pointer to **time.Time** |  | [optional] 
**Contents** | Pointer to **string** |  | [optional] 
**ContentsFormat** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewEntryHistory

`func NewEntryHistory() *EntryHistory`

NewEntryHistory instantiates a new EntryHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryHistoryWithDefaults

`func NewEntryHistoryWithDefaults() *EntryHistory`

NewEntryHistoryWithDefaults instantiates a new EntryHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentDate

`func (o *EntryHistory) GetContentDate() time.Time`

GetContentDate returns the ContentDate field if non-nil, zero value otherwise.

### GetContentDateOk

`func (o *EntryHistory) GetContentDateOk() (*time.Time, bool)`

GetContentDateOk returns a tuple with the ContentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentDate

`func (o *EntryHistory) SetContentDate(v time.Time)`

SetContentDate sets ContentDate field to given value.

### HasContentDate

`func (o *EntryHistory) HasContentDate() bool`

HasContentDate returns a boolean if a field has been set.

### GetContents

`func (o *EntryHistory) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *EntryHistory) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *EntryHistory) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *EntryHistory) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetContentsFormat

`func (o *EntryHistory) GetContentsFormat() string`

GetContentsFormat returns the ContentsFormat field if non-nil, zero value otherwise.

### GetContentsFormatOk

`func (o *EntryHistory) GetContentsFormatOk() (*string, bool)`

GetContentsFormatOk returns a tuple with the ContentsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsFormat

`func (o *EntryHistory) SetContentsFormat(v string)`

SetContentsFormat sets ContentsFormat field to given value.

### HasContentsFormat

`func (o *EntryHistory) HasContentsFormat() bool`

HasContentsFormat returns a boolean if a field has been set.

### GetUser

`func (o *EntryHistory) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EntryHistory) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EntryHistory) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *EntryHistory) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


