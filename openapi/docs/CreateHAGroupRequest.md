# CreateHAGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ElasticsearchAddress** | Pointer to **string** |  | [optional] 
**ElasticIndexPrefix** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateHAGroupRequest

`func NewCreateHAGroupRequest() *CreateHAGroupRequest`

NewCreateHAGroupRequest instantiates a new CreateHAGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHAGroupRequestWithDefaults

`func NewCreateHAGroupRequestWithDefaults() *CreateHAGroupRequest`

NewCreateHAGroupRequestWithDefaults instantiates a new CreateHAGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateHAGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateHAGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateHAGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateHAGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetElasticsearchAddress

`func (o *CreateHAGroupRequest) GetElasticsearchAddress() string`

GetElasticsearchAddress returns the ElasticsearchAddress field if non-nil, zero value otherwise.

### GetElasticsearchAddressOk

`func (o *CreateHAGroupRequest) GetElasticsearchAddressOk() (*string, bool)`

GetElasticsearchAddressOk returns a tuple with the ElasticsearchAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticsearchAddress

`func (o *CreateHAGroupRequest) SetElasticsearchAddress(v string)`

SetElasticsearchAddress sets ElasticsearchAddress field to given value.

### HasElasticsearchAddress

`func (o *CreateHAGroupRequest) HasElasticsearchAddress() bool`

HasElasticsearchAddress returns a boolean if a field has been set.

### GetElasticIndexPrefix

`func (o *CreateHAGroupRequest) GetElasticIndexPrefix() string`

GetElasticIndexPrefix returns the ElasticIndexPrefix field if non-nil, zero value otherwise.

### GetElasticIndexPrefixOk

`func (o *CreateHAGroupRequest) GetElasticIndexPrefixOk() (*string, bool)`

GetElasticIndexPrefixOk returns a tuple with the ElasticIndexPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticIndexPrefix

`func (o *CreateHAGroupRequest) SetElasticIndexPrefix(v string)`

SetElasticIndexPrefix sets ElasticIndexPrefix field to given value.

### HasElasticIndexPrefix

`func (o *CreateHAGroupRequest) HasElasticIndexPrefix() bool`

HasElasticIndexPrefix returns a boolean if a field has been set.

### GetId

`func (o *CreateHAGroupRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateHAGroupRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateHAGroupRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateHAGroupRequest) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


