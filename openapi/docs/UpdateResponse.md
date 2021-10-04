# UpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotUpdated** | Pointer to **int64** |  | [optional] 
**UpdatedIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateResponse

`func NewUpdateResponse() *UpdateResponse`

NewUpdateResponse instantiates a new UpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResponseWithDefaults

`func NewUpdateResponseWithDefaults() *UpdateResponse`

NewUpdateResponseWithDefaults instantiates a new UpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotUpdated

`func (o *UpdateResponse) GetNotUpdated() int64`

GetNotUpdated returns the NotUpdated field if non-nil, zero value otherwise.

### GetNotUpdatedOk

`func (o *UpdateResponse) GetNotUpdatedOk() (*int64, bool)`

GetNotUpdatedOk returns a tuple with the NotUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotUpdated

`func (o *UpdateResponse) SetNotUpdated(v int64)`

SetNotUpdated sets NotUpdated field to given value.

### HasNotUpdated

`func (o *UpdateResponse) HasNotUpdated() bool`

HasNotUpdated returns a boolean if a field has been set.

### GetUpdatedIds

`func (o *UpdateResponse) GetUpdatedIds() []string`

GetUpdatedIds returns the UpdatedIds field if non-nil, zero value otherwise.

### GetUpdatedIdsOk

`func (o *UpdateResponse) GetUpdatedIdsOk() (*[]string, bool)`

GetUpdatedIdsOk returns a tuple with the UpdatedIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedIds

`func (o *UpdateResponse) SetUpdatedIds(v []string)`

SetUpdatedIds sets UpdatedIds field to given value.

### HasUpdatedIds

`func (o *UpdateResponse) HasUpdatedIds() bool`

HasUpdatedIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


