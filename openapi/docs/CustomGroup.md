# CustomGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]Buckets**](Buckets.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomGroup

`func NewCustomGroup() *CustomGroup`

NewCustomGroup instantiates a new CustomGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomGroupWithDefaults

`func NewCustomGroupWithDefaults() *CustomGroup`

NewCustomGroupWithDefaults instantiates a new CustomGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *CustomGroup) GetConditions() []Buckets`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CustomGroup) GetConditionsOk() (*[]Buckets, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CustomGroup) SetConditions(v []Buckets)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *CustomGroup) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetName

`func (o *CustomGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomGroup) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


