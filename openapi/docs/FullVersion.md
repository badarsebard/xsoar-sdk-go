# FullVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryTerm** | Pointer to **int64** |  | [optional] 
**SequenceNumber** | Pointer to **int64** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewFullVersion

`func NewFullVersion() *FullVersion`

NewFullVersion instantiates a new FullVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullVersionWithDefaults

`func NewFullVersionWithDefaults() *FullVersion`

NewFullVersionWithDefaults instantiates a new FullVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryTerm

`func (o *FullVersion) GetPrimaryTerm() int64`

GetPrimaryTerm returns the PrimaryTerm field if non-nil, zero value otherwise.

### GetPrimaryTermOk

`func (o *FullVersion) GetPrimaryTermOk() (*int64, bool)`

GetPrimaryTermOk returns a tuple with the PrimaryTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTerm

`func (o *FullVersion) SetPrimaryTerm(v int64)`

SetPrimaryTerm sets PrimaryTerm field to given value.

### HasPrimaryTerm

`func (o *FullVersion) HasPrimaryTerm() bool`

HasPrimaryTerm returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *FullVersion) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *FullVersion) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *FullVersion) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *FullVersion) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetVersion

`func (o *FullVersion) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FullVersion) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FullVersion) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FullVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


