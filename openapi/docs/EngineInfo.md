# EngineInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engine** | Pointer to **string** | Engine that will run the script | [optional] 
**EngineGroup** | Pointer to **string** | EngineGroup that will run the script | [optional] 

## Methods

### NewEngineInfo

`func NewEngineInfo() *EngineInfo`

NewEngineInfo instantiates a new EngineInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineInfoWithDefaults

`func NewEngineInfoWithDefaults() *EngineInfo`

NewEngineInfoWithDefaults instantiates a new EngineInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngine

`func (o *EngineInfo) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *EngineInfo) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *EngineInfo) SetEngine(v string)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *EngineInfo) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetEngineGroup

`func (o *EngineInfo) GetEngineGroup() string`

GetEngineGroup returns the EngineGroup field if non-nil, zero value otherwise.

### GetEngineGroupOk

`func (o *EngineInfo) GetEngineGroupOk() (*string, bool)`

GetEngineGroupOk returns a tuple with the EngineGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineGroup

`func (o *EngineInfo) SetEngineGroup(v string)`

SetEngineGroup sets EngineGroup field to given value.

### HasEngineGroup

`func (o *EngineInfo) HasEngineGroup() bool`

HasEngineGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


