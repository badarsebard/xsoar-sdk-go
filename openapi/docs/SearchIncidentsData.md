# SearchIncidentsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**IncidentFilter**](IncidentFilter.md) |  | [optional] 
**UserFilter** | Pointer to **bool** |  | [optional] 

## Methods

### NewSearchIncidentsData

`func NewSearchIncidentsData() *SearchIncidentsData`

NewSearchIncidentsData instantiates a new SearchIncidentsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIncidentsDataWithDefaults

`func NewSearchIncidentsDataWithDefaults() *SearchIncidentsData`

NewSearchIncidentsDataWithDefaults instantiates a new SearchIncidentsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *SearchIncidentsData) GetFilter() IncidentFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SearchIncidentsData) GetFilterOk() (*IncidentFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SearchIncidentsData) SetFilter(v IncidentFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *SearchIncidentsData) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetUserFilter

`func (o *SearchIncidentsData) GetUserFilter() bool`

GetUserFilter returns the UserFilter field if non-nil, zero value otherwise.

### GetUserFilterOk

`func (o *SearchIncidentsData) GetUserFilterOk() (*bool, bool)`

GetUserFilterOk returns a tuple with the UserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFilter

`func (o *SearchIncidentsData) SetUserFilter(v bool)`

SetUserFilter sets UserFilter field to given value.

### HasUserFilter

`func (o *SearchIncidentsData) HasUserFilter() bool`

HasUserFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


