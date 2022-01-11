# RelationshipAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** |  | [optional] 
**EntityA** | Pointer to **string** |  | [optional] 
**EntityAFamily** | Pointer to **string** |  | [optional] 
**EntityAType** | Pointer to **string** |  | [optional] 
**EntityB** | Pointer to **string** |  | [optional] 
**EntityBFamily** | Pointer to **string** |  | [optional] 
**EntityBType** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to [**CustomFields**](CustomFields.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Instance** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Reliability** | Pointer to **string** |  | [optional] 
**ReverseName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewRelationshipAPI

`func NewRelationshipAPI() *RelationshipAPI`

NewRelationshipAPI instantiates a new RelationshipAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipAPIWithDefaults

`func NewRelationshipAPIWithDefaults() *RelationshipAPI`

NewRelationshipAPIWithDefaults instantiates a new RelationshipAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *RelationshipAPI) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *RelationshipAPI) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *RelationshipAPI) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *RelationshipAPI) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetEntityA

`func (o *RelationshipAPI) GetEntityA() string`

GetEntityA returns the EntityA field if non-nil, zero value otherwise.

### GetEntityAOk

`func (o *RelationshipAPI) GetEntityAOk() (*string, bool)`

GetEntityAOk returns a tuple with the EntityA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityA

`func (o *RelationshipAPI) SetEntityA(v string)`

SetEntityA sets EntityA field to given value.

### HasEntityA

`func (o *RelationshipAPI) HasEntityA() bool`

HasEntityA returns a boolean if a field has been set.

### GetEntityAFamily

`func (o *RelationshipAPI) GetEntityAFamily() string`

GetEntityAFamily returns the EntityAFamily field if non-nil, zero value otherwise.

### GetEntityAFamilyOk

`func (o *RelationshipAPI) GetEntityAFamilyOk() (*string, bool)`

GetEntityAFamilyOk returns a tuple with the EntityAFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityAFamily

`func (o *RelationshipAPI) SetEntityAFamily(v string)`

SetEntityAFamily sets EntityAFamily field to given value.

### HasEntityAFamily

`func (o *RelationshipAPI) HasEntityAFamily() bool`

HasEntityAFamily returns a boolean if a field has been set.

### GetEntityAType

`func (o *RelationshipAPI) GetEntityAType() string`

GetEntityAType returns the EntityAType field if non-nil, zero value otherwise.

### GetEntityATypeOk

`func (o *RelationshipAPI) GetEntityATypeOk() (*string, bool)`

GetEntityATypeOk returns a tuple with the EntityAType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityAType

`func (o *RelationshipAPI) SetEntityAType(v string)`

SetEntityAType sets EntityAType field to given value.

### HasEntityAType

`func (o *RelationshipAPI) HasEntityAType() bool`

HasEntityAType returns a boolean if a field has been set.

### GetEntityB

`func (o *RelationshipAPI) GetEntityB() string`

GetEntityB returns the EntityB field if non-nil, zero value otherwise.

### GetEntityBOk

`func (o *RelationshipAPI) GetEntityBOk() (*string, bool)`

GetEntityBOk returns a tuple with the EntityB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityB

`func (o *RelationshipAPI) SetEntityB(v string)`

SetEntityB sets EntityB field to given value.

### HasEntityB

`func (o *RelationshipAPI) HasEntityB() bool`

HasEntityB returns a boolean if a field has been set.

### GetEntityBFamily

`func (o *RelationshipAPI) GetEntityBFamily() string`

GetEntityBFamily returns the EntityBFamily field if non-nil, zero value otherwise.

### GetEntityBFamilyOk

`func (o *RelationshipAPI) GetEntityBFamilyOk() (*string, bool)`

GetEntityBFamilyOk returns a tuple with the EntityBFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityBFamily

`func (o *RelationshipAPI) SetEntityBFamily(v string)`

SetEntityBFamily sets EntityBFamily field to given value.

### HasEntityBFamily

`func (o *RelationshipAPI) HasEntityBFamily() bool`

HasEntityBFamily returns a boolean if a field has been set.

### GetEntityBType

`func (o *RelationshipAPI) GetEntityBType() string`

GetEntityBType returns the EntityBType field if non-nil, zero value otherwise.

### GetEntityBTypeOk

`func (o *RelationshipAPI) GetEntityBTypeOk() (*string, bool)`

GetEntityBTypeOk returns a tuple with the EntityBType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityBType

`func (o *RelationshipAPI) SetEntityBType(v string)`

SetEntityBType sets EntityBType field to given value.

### HasEntityBType

`func (o *RelationshipAPI) HasEntityBType() bool`

HasEntityBType returns a boolean if a field has been set.

### GetFields

`func (o *RelationshipAPI) GetFields() CustomFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RelationshipAPI) GetFieldsOk() (*CustomFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RelationshipAPI) SetFields(v CustomFields)`

SetFields sets Fields field to given value.

### HasFields

`func (o *RelationshipAPI) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *RelationshipAPI) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationshipAPI) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationshipAPI) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelationshipAPI) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstance

`func (o *RelationshipAPI) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *RelationshipAPI) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *RelationshipAPI) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *RelationshipAPI) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetName

`func (o *RelationshipAPI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RelationshipAPI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RelationshipAPI) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RelationshipAPI) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReliability

`func (o *RelationshipAPI) GetReliability() string`

GetReliability returns the Reliability field if non-nil, zero value otherwise.

### GetReliabilityOk

`func (o *RelationshipAPI) GetReliabilityOk() (*string, bool)`

GetReliabilityOk returns a tuple with the Reliability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliability

`func (o *RelationshipAPI) SetReliability(v string)`

SetReliability sets Reliability field to given value.

### HasReliability

`func (o *RelationshipAPI) HasReliability() bool`

HasReliability returns a boolean if a field has been set.

### GetReverseName

`func (o *RelationshipAPI) GetReverseName() string`

GetReverseName returns the ReverseName field if non-nil, zero value otherwise.

### GetReverseNameOk

`func (o *RelationshipAPI) GetReverseNameOk() (*string, bool)`

GetReverseNameOk returns a tuple with the ReverseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseName

`func (o *RelationshipAPI) SetReverseName(v string)`

SetReverseName sets ReverseName field to given value.

### HasReverseName

`func (o *RelationshipAPI) HasReverseName() bool`

HasReverseName returns a boolean if a field has been set.

### GetType

`func (o *RelationshipAPI) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipAPI) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipAPI) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipAPI) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


