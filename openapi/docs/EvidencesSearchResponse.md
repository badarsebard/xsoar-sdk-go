# EvidencesSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Evidences** | Pointer to [**Evidences**](Evidences.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewEvidencesSearchResponse

`func NewEvidencesSearchResponse() *EvidencesSearchResponse`

NewEvidencesSearchResponse instantiates a new EvidencesSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvidencesSearchResponseWithDefaults

`func NewEvidencesSearchResponseWithDefaults() *EvidencesSearchResponse`

NewEvidencesSearchResponseWithDefaults instantiates a new EvidencesSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvidences

`func (o *EvidencesSearchResponse) GetEvidences() Evidences`

GetEvidences returns the Evidences field if non-nil, zero value otherwise.

### GetEvidencesOk

`func (o *EvidencesSearchResponse) GetEvidencesOk() (*Evidences, bool)`

GetEvidencesOk returns a tuple with the Evidences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidences

`func (o *EvidencesSearchResponse) SetEvidences(v Evidences)`

SetEvidences sets Evidences field to given value.

### HasEvidences

`func (o *EvidencesSearchResponse) HasEvidences() bool`

HasEvidences returns a boolean if a field has been set.

### GetTotal

`func (o *EvidencesSearchResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *EvidencesSearchResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *EvidencesSearchResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *EvidencesSearchResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


