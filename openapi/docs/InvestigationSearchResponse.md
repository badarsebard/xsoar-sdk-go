# InvestigationSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Investigation**](Investigation.md) | in: body | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewInvestigationSearchResponse

`func NewInvestigationSearchResponse() *InvestigationSearchResponse`

NewInvestigationSearchResponse instantiates a new InvestigationSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestigationSearchResponseWithDefaults

`func NewInvestigationSearchResponseWithDefaults() *InvestigationSearchResponse`

NewInvestigationSearchResponseWithDefaults instantiates a new InvestigationSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *InvestigationSearchResponse) GetData() []Investigation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InvestigationSearchResponse) GetDataOk() (*[]Investigation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InvestigationSearchResponse) SetData(v []Investigation)`

SetData sets Data field to given value.

### HasData

`func (o *InvestigationSearchResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *InvestigationSearchResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InvestigationSearchResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InvestigationSearchResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InvestigationSearchResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


