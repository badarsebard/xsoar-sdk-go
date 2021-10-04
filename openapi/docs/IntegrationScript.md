# IntegrationScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commands** | Pointer to [**[]Command**](Command.md) |  | [optional] 
**DockerImage** | Pointer to **string** |  | [optional] 
**Feed** | Pointer to **bool** |  | [optional] 
**IsFetch** | Pointer to **bool** |  | [optional] 
**IsFetchCredentials** | Pointer to **bool** |  | [optional] 
**IsFetchSamples** | Pointer to **bool** |  | [optional] 
**IsMappable** | Pointer to **bool** |  | [optional] 
**IsRemoteSyncIn** | Pointer to **bool** |  | [optional] 
**IsRemoteSyncOut** | Pointer to **bool** |  | [optional] 
**LongRunning** | Pointer to **bool** |  | [optional] 
**LongRunningPortMapping** | Pointer to **bool** |  | [optional] 
**ResetContext** | Pointer to **bool** |  | [optional] 
**RunOnce** | Pointer to **bool** |  | [optional] 
**Script** | Pointer to **string** |  | [optional] 
**Subtype** | Pointer to **string** | ScriptSubType holds the script type version | [optional] 
**Type** | Pointer to **string** | ScriptType holds the type of a script | [optional] 

## Methods

### NewIntegrationScript

`func NewIntegrationScript() *IntegrationScript`

NewIntegrationScript instantiates a new IntegrationScript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationScriptWithDefaults

`func NewIntegrationScriptWithDefaults() *IntegrationScript`

NewIntegrationScriptWithDefaults instantiates a new IntegrationScript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommands

`func (o *IntegrationScript) GetCommands() []Command`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *IntegrationScript) GetCommandsOk() (*[]Command, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *IntegrationScript) SetCommands(v []Command)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *IntegrationScript) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### GetDockerImage

`func (o *IntegrationScript) GetDockerImage() string`

GetDockerImage returns the DockerImage field if non-nil, zero value otherwise.

### GetDockerImageOk

`func (o *IntegrationScript) GetDockerImageOk() (*string, bool)`

GetDockerImageOk returns a tuple with the DockerImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImage

`func (o *IntegrationScript) SetDockerImage(v string)`

SetDockerImage sets DockerImage field to given value.

### HasDockerImage

`func (o *IntegrationScript) HasDockerImage() bool`

HasDockerImage returns a boolean if a field has been set.

### GetFeed

`func (o *IntegrationScript) GetFeed() bool`

GetFeed returns the Feed field if non-nil, zero value otherwise.

### GetFeedOk

`func (o *IntegrationScript) GetFeedOk() (*bool, bool)`

GetFeedOk returns a tuple with the Feed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeed

`func (o *IntegrationScript) SetFeed(v bool)`

SetFeed sets Feed field to given value.

### HasFeed

`func (o *IntegrationScript) HasFeed() bool`

HasFeed returns a boolean if a field has been set.

### GetIsFetch

`func (o *IntegrationScript) GetIsFetch() bool`

GetIsFetch returns the IsFetch field if non-nil, zero value otherwise.

### GetIsFetchOk

`func (o *IntegrationScript) GetIsFetchOk() (*bool, bool)`

GetIsFetchOk returns a tuple with the IsFetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFetch

`func (o *IntegrationScript) SetIsFetch(v bool)`

SetIsFetch sets IsFetch field to given value.

### HasIsFetch

`func (o *IntegrationScript) HasIsFetch() bool`

HasIsFetch returns a boolean if a field has been set.

### GetIsFetchCredentials

`func (o *IntegrationScript) GetIsFetchCredentials() bool`

GetIsFetchCredentials returns the IsFetchCredentials field if non-nil, zero value otherwise.

### GetIsFetchCredentialsOk

`func (o *IntegrationScript) GetIsFetchCredentialsOk() (*bool, bool)`

GetIsFetchCredentialsOk returns a tuple with the IsFetchCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFetchCredentials

`func (o *IntegrationScript) SetIsFetchCredentials(v bool)`

SetIsFetchCredentials sets IsFetchCredentials field to given value.

### HasIsFetchCredentials

`func (o *IntegrationScript) HasIsFetchCredentials() bool`

HasIsFetchCredentials returns a boolean if a field has been set.

### GetIsFetchSamples

`func (o *IntegrationScript) GetIsFetchSamples() bool`

GetIsFetchSamples returns the IsFetchSamples field if non-nil, zero value otherwise.

### GetIsFetchSamplesOk

`func (o *IntegrationScript) GetIsFetchSamplesOk() (*bool, bool)`

GetIsFetchSamplesOk returns a tuple with the IsFetchSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFetchSamples

`func (o *IntegrationScript) SetIsFetchSamples(v bool)`

SetIsFetchSamples sets IsFetchSamples field to given value.

### HasIsFetchSamples

`func (o *IntegrationScript) HasIsFetchSamples() bool`

HasIsFetchSamples returns a boolean if a field has been set.

### GetIsMappable

`func (o *IntegrationScript) GetIsMappable() bool`

GetIsMappable returns the IsMappable field if non-nil, zero value otherwise.

### GetIsMappableOk

`func (o *IntegrationScript) GetIsMappableOk() (*bool, bool)`

GetIsMappableOk returns a tuple with the IsMappable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMappable

`func (o *IntegrationScript) SetIsMappable(v bool)`

SetIsMappable sets IsMappable field to given value.

### HasIsMappable

`func (o *IntegrationScript) HasIsMappable() bool`

HasIsMappable returns a boolean if a field has been set.

### GetIsRemoteSyncIn

`func (o *IntegrationScript) GetIsRemoteSyncIn() bool`

GetIsRemoteSyncIn returns the IsRemoteSyncIn field if non-nil, zero value otherwise.

### GetIsRemoteSyncInOk

`func (o *IntegrationScript) GetIsRemoteSyncInOk() (*bool, bool)`

GetIsRemoteSyncInOk returns a tuple with the IsRemoteSyncIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteSyncIn

`func (o *IntegrationScript) SetIsRemoteSyncIn(v bool)`

SetIsRemoteSyncIn sets IsRemoteSyncIn field to given value.

### HasIsRemoteSyncIn

`func (o *IntegrationScript) HasIsRemoteSyncIn() bool`

HasIsRemoteSyncIn returns a boolean if a field has been set.

### GetIsRemoteSyncOut

`func (o *IntegrationScript) GetIsRemoteSyncOut() bool`

GetIsRemoteSyncOut returns the IsRemoteSyncOut field if non-nil, zero value otherwise.

### GetIsRemoteSyncOutOk

`func (o *IntegrationScript) GetIsRemoteSyncOutOk() (*bool, bool)`

GetIsRemoteSyncOutOk returns a tuple with the IsRemoteSyncOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteSyncOut

`func (o *IntegrationScript) SetIsRemoteSyncOut(v bool)`

SetIsRemoteSyncOut sets IsRemoteSyncOut field to given value.

### HasIsRemoteSyncOut

`func (o *IntegrationScript) HasIsRemoteSyncOut() bool`

HasIsRemoteSyncOut returns a boolean if a field has been set.

### GetLongRunning

`func (o *IntegrationScript) GetLongRunning() bool`

GetLongRunning returns the LongRunning field if non-nil, zero value otherwise.

### GetLongRunningOk

`func (o *IntegrationScript) GetLongRunningOk() (*bool, bool)`

GetLongRunningOk returns a tuple with the LongRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongRunning

`func (o *IntegrationScript) SetLongRunning(v bool)`

SetLongRunning sets LongRunning field to given value.

### HasLongRunning

`func (o *IntegrationScript) HasLongRunning() bool`

HasLongRunning returns a boolean if a field has been set.

### GetLongRunningPortMapping

`func (o *IntegrationScript) GetLongRunningPortMapping() bool`

GetLongRunningPortMapping returns the LongRunningPortMapping field if non-nil, zero value otherwise.

### GetLongRunningPortMappingOk

`func (o *IntegrationScript) GetLongRunningPortMappingOk() (*bool, bool)`

GetLongRunningPortMappingOk returns a tuple with the LongRunningPortMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongRunningPortMapping

`func (o *IntegrationScript) SetLongRunningPortMapping(v bool)`

SetLongRunningPortMapping sets LongRunningPortMapping field to given value.

### HasLongRunningPortMapping

`func (o *IntegrationScript) HasLongRunningPortMapping() bool`

HasLongRunningPortMapping returns a boolean if a field has been set.

### GetResetContext

`func (o *IntegrationScript) GetResetContext() bool`

GetResetContext returns the ResetContext field if non-nil, zero value otherwise.

### GetResetContextOk

`func (o *IntegrationScript) GetResetContextOk() (*bool, bool)`

GetResetContextOk returns a tuple with the ResetContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetContext

`func (o *IntegrationScript) SetResetContext(v bool)`

SetResetContext sets ResetContext field to given value.

### HasResetContext

`func (o *IntegrationScript) HasResetContext() bool`

HasResetContext returns a boolean if a field has been set.

### GetRunOnce

`func (o *IntegrationScript) GetRunOnce() bool`

GetRunOnce returns the RunOnce field if non-nil, zero value otherwise.

### GetRunOnceOk

`func (o *IntegrationScript) GetRunOnceOk() (*bool, bool)`

GetRunOnceOk returns a tuple with the RunOnce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOnce

`func (o *IntegrationScript) SetRunOnce(v bool)`

SetRunOnce sets RunOnce field to given value.

### HasRunOnce

`func (o *IntegrationScript) HasRunOnce() bool`

HasRunOnce returns a boolean if a field has been set.

### GetScript

`func (o *IntegrationScript) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *IntegrationScript) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *IntegrationScript) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *IntegrationScript) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetSubtype

`func (o *IntegrationScript) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *IntegrationScript) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *IntegrationScript) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *IntegrationScript) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetType

`func (o *IntegrationScript) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationScript) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationScript) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IntegrationScript) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


