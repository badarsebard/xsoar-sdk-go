# VersionedCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitMessage** | Pointer to **string** |  | [optional] 
**ShouldCommit** | Pointer to **bool** |  | [optional] 

## Methods

### NewVersionedCommit

`func NewVersionedCommit() *VersionedCommit`

NewVersionedCommit instantiates a new VersionedCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionedCommitWithDefaults

`func NewVersionedCommitWithDefaults() *VersionedCommit`

NewVersionedCommitWithDefaults instantiates a new VersionedCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitMessage

`func (o *VersionedCommit) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *VersionedCommit) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *VersionedCommit) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *VersionedCommit) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetShouldCommit

`func (o *VersionedCommit) GetShouldCommit() bool`

GetShouldCommit returns the ShouldCommit field if non-nil, zero value otherwise.

### GetShouldCommitOk

`func (o *VersionedCommit) GetShouldCommitOk() (*bool, bool)`

GetShouldCommitOk returns a tuple with the ShouldCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCommit

`func (o *VersionedCommit) SetShouldCommit(v bool)`

SetShouldCommit sets ShouldCommit field to given value.

### HasShouldCommit

`func (o *VersionedCommit) HasShouldCommit() bool`

HasShouldCommit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


