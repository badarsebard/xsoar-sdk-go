# ContentItemFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromServerVersion** | Pointer to [**Version**](Version.md) |  | [optional] 
**ItemVersion** | Pointer to [**Version**](Version.md) |  | [optional] 
**PackID** | Pointer to **string** |  | [optional] 
**PackPropagationLabels** | Pointer to **[]string** |  | [optional] 
**PropagationLabels** | Pointer to **[]string** |  | [optional] 
**ToServerVersion** | Pointer to [**Version**](Version.md) |  | [optional] 

## Methods

### NewContentItemFields

`func NewContentItemFields() *ContentItemFields`

NewContentItemFields instantiates a new ContentItemFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentItemFieldsWithDefaults

`func NewContentItemFieldsWithDefaults() *ContentItemFields`

NewContentItemFieldsWithDefaults instantiates a new ContentItemFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromServerVersion

`func (o *ContentItemFields) GetFromServerVersion() Version`

GetFromServerVersion returns the FromServerVersion field if non-nil, zero value otherwise.

### GetFromServerVersionOk

`func (o *ContentItemFields) GetFromServerVersionOk() (*Version, bool)`

GetFromServerVersionOk returns a tuple with the FromServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromServerVersion

`func (o *ContentItemFields) SetFromServerVersion(v Version)`

SetFromServerVersion sets FromServerVersion field to given value.

### HasFromServerVersion

`func (o *ContentItemFields) HasFromServerVersion() bool`

HasFromServerVersion returns a boolean if a field has been set.

### GetItemVersion

`func (o *ContentItemFields) GetItemVersion() Version`

GetItemVersion returns the ItemVersion field if non-nil, zero value otherwise.

### GetItemVersionOk

`func (o *ContentItemFields) GetItemVersionOk() (*Version, bool)`

GetItemVersionOk returns a tuple with the ItemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVersion

`func (o *ContentItemFields) SetItemVersion(v Version)`

SetItemVersion sets ItemVersion field to given value.

### HasItemVersion

`func (o *ContentItemFields) HasItemVersion() bool`

HasItemVersion returns a boolean if a field has been set.

### GetPackID

`func (o *ContentItemFields) GetPackID() string`

GetPackID returns the PackID field if non-nil, zero value otherwise.

### GetPackIDOk

`func (o *ContentItemFields) GetPackIDOk() (*string, bool)`

GetPackIDOk returns a tuple with the PackID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackID

`func (o *ContentItemFields) SetPackID(v string)`

SetPackID sets PackID field to given value.

### HasPackID

`func (o *ContentItemFields) HasPackID() bool`

HasPackID returns a boolean if a field has been set.

### GetPackPropagationLabels

`func (o *ContentItemFields) GetPackPropagationLabels() []string`

GetPackPropagationLabels returns the PackPropagationLabels field if non-nil, zero value otherwise.

### GetPackPropagationLabelsOk

`func (o *ContentItemFields) GetPackPropagationLabelsOk() (*[]string, bool)`

GetPackPropagationLabelsOk returns a tuple with the PackPropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackPropagationLabels

`func (o *ContentItemFields) SetPackPropagationLabels(v []string)`

SetPackPropagationLabels sets PackPropagationLabels field to given value.

### HasPackPropagationLabels

`func (o *ContentItemFields) HasPackPropagationLabels() bool`

HasPackPropagationLabels returns a boolean if a field has been set.

### GetPropagationLabels

`func (o *ContentItemFields) GetPropagationLabels() []string`

GetPropagationLabels returns the PropagationLabels field if non-nil, zero value otherwise.

### GetPropagationLabelsOk

`func (o *ContentItemFields) GetPropagationLabelsOk() (*[]string, bool)`

GetPropagationLabelsOk returns a tuple with the PropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationLabels

`func (o *ContentItemFields) SetPropagationLabels(v []string)`

SetPropagationLabels sets PropagationLabels field to given value.

### HasPropagationLabels

`func (o *ContentItemFields) HasPropagationLabels() bool`

HasPropagationLabels returns a boolean if a field has been set.

### GetToServerVersion

`func (o *ContentItemFields) GetToServerVersion() Version`

GetToServerVersion returns the ToServerVersion field if non-nil, zero value otherwise.

### GetToServerVersionOk

`func (o *ContentItemFields) GetToServerVersionOk() (*Version, bool)`

GetToServerVersionOk returns a tuple with the ToServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToServerVersion

`func (o *ContentItemFields) SetToServerVersion(v Version)`

SetToServerVersion sets ToServerVersion field to given value.

### HasToServerVersion

`func (o *ContentItemFields) HasToServerVersion() bool`

HasToServerVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


