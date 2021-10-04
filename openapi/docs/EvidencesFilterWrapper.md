# EvidencesFilterWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**GenericStringDateFilter**](GenericStringDateFilter.md) |  | [optional] 
**IncidentID** | Pointer to **string** |  | [optional] 

## Methods

### NewEvidencesFilterWrapper

`func NewEvidencesFilterWrapper() *EvidencesFilterWrapper`

NewEvidencesFilterWrapper instantiates a new EvidencesFilterWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvidencesFilterWrapperWithDefaults

`func NewEvidencesFilterWrapperWithDefaults() *EvidencesFilterWrapper`

NewEvidencesFilterWrapperWithDefaults instantiates a new EvidencesFilterWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *EvidencesFilterWrapper) GetFilter() GenericStringDateFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *EvidencesFilterWrapper) GetFilterOk() (*GenericStringDateFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *EvidencesFilterWrapper) SetFilter(v GenericStringDateFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *EvidencesFilterWrapper) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetIncidentID

`func (o *EvidencesFilterWrapper) GetIncidentID() string`

GetIncidentID returns the IncidentID field if non-nil, zero value otherwise.

### GetIncidentIDOk

`func (o *EvidencesFilterWrapper) GetIncidentIDOk() (*string, bool)`

GetIncidentIDOk returns a tuple with the IncidentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentID

`func (o *EvidencesFilterWrapper) SetIncidentID(v string)`

SetIncidentID sets IncidentID field to given value.

### HasIncidentID

`func (o *EvidencesFilterWrapper) HasIncidentID() bool`

HasIncidentID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


