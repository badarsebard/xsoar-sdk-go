# ScriptAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to [**[]Argument**](Argument.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**ContextKeys** | Pointer to **[]string** |  | [optional] 
**DependsOn** | Pointer to **map[string][]string** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**Detached** | Pointer to **bool** |  | [optional] 
**DockerImage** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Outputs** | Pointer to [**[]Output**](Output.md) |  | [optional] 
**Permitted** | Pointer to **bool** |  | [optional] 
**Polling** | Pointer to **bool** |  | [optional] 
**PropagationLabels** | Pointer to **[]string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**RunAs** | Pointer to **string** |  | [optional] 
**ScriptTarget** | Pointer to **int64** | ScriptTarget represents the module where this script should run | [optional] 
**System** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** | ScriptType holds the type of a script | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewScriptAPI

`func NewScriptAPI() *ScriptAPI`

NewScriptAPI instantiates a new ScriptAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptAPIWithDefaults

`func NewScriptAPIWithDefaults() *ScriptAPI`

NewScriptAPIWithDefaults instantiates a new ScriptAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *ScriptAPI) GetArguments() []Argument`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ScriptAPI) GetArgumentsOk() (*[]Argument, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ScriptAPI) SetArguments(v []Argument)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *ScriptAPI) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetComment

`func (o *ScriptAPI) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ScriptAPI) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ScriptAPI) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ScriptAPI) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContextKeys

`func (o *ScriptAPI) GetContextKeys() []string`

GetContextKeys returns the ContextKeys field if non-nil, zero value otherwise.

### GetContextKeysOk

`func (o *ScriptAPI) GetContextKeysOk() (*[]string, bool)`

GetContextKeysOk returns a tuple with the ContextKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextKeys

`func (o *ScriptAPI) SetContextKeys(v []string)`

SetContextKeys sets ContextKeys field to given value.

### HasContextKeys

`func (o *ScriptAPI) HasContextKeys() bool`

HasContextKeys returns a boolean if a field has been set.

### GetDependsOn

`func (o *ScriptAPI) GetDependsOn() map[string][]string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *ScriptAPI) GetDependsOnOk() (*map[string][]string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *ScriptAPI) SetDependsOn(v map[string][]string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *ScriptAPI) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### GetDeprecated

`func (o *ScriptAPI) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *ScriptAPI) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *ScriptAPI) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *ScriptAPI) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDetached

`func (o *ScriptAPI) GetDetached() bool`

GetDetached returns the Detached field if non-nil, zero value otherwise.

### GetDetachedOk

`func (o *ScriptAPI) GetDetachedOk() (*bool, bool)`

GetDetachedOk returns a tuple with the Detached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetached

`func (o *ScriptAPI) SetDetached(v bool)`

SetDetached sets Detached field to given value.

### HasDetached

`func (o *ScriptAPI) HasDetached() bool`

HasDetached returns a boolean if a field has been set.

### GetDockerImage

`func (o *ScriptAPI) GetDockerImage() string`

GetDockerImage returns the DockerImage field if non-nil, zero value otherwise.

### GetDockerImageOk

`func (o *ScriptAPI) GetDockerImageOk() (*string, bool)`

GetDockerImageOk returns a tuple with the DockerImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImage

`func (o *ScriptAPI) SetDockerImage(v string)`

SetDockerImage sets DockerImage field to given value.

### HasDockerImage

`func (o *ScriptAPI) HasDockerImage() bool`

HasDockerImage returns a boolean if a field has been set.

### GetEnabled

`func (o *ScriptAPI) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ScriptAPI) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ScriptAPI) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ScriptAPI) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHidden

`func (o *ScriptAPI) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ScriptAPI) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ScriptAPI) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ScriptAPI) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetId

`func (o *ScriptAPI) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScriptAPI) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScriptAPI) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScriptAPI) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocked

`func (o *ScriptAPI) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ScriptAPI) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ScriptAPI) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ScriptAPI) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetModified

`func (o *ScriptAPI) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ScriptAPI) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ScriptAPI) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ScriptAPI) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *ScriptAPI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScriptAPI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScriptAPI) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScriptAPI) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputs

`func (o *ScriptAPI) GetOutputs() []Output`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *ScriptAPI) GetOutputsOk() (*[]Output, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *ScriptAPI) SetOutputs(v []Output)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *ScriptAPI) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetPermitted

`func (o *ScriptAPI) GetPermitted() bool`

GetPermitted returns the Permitted field if non-nil, zero value otherwise.

### GetPermittedOk

`func (o *ScriptAPI) GetPermittedOk() (*bool, bool)`

GetPermittedOk returns a tuple with the Permitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitted

`func (o *ScriptAPI) SetPermitted(v bool)`

SetPermitted sets Permitted field to given value.

### HasPermitted

`func (o *ScriptAPI) HasPermitted() bool`

HasPermitted returns a boolean if a field has been set.

### GetPolling

`func (o *ScriptAPI) GetPolling() bool`

GetPolling returns the Polling field if non-nil, zero value otherwise.

### GetPollingOk

`func (o *ScriptAPI) GetPollingOk() (*bool, bool)`

GetPollingOk returns a tuple with the Polling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolling

`func (o *ScriptAPI) SetPolling(v bool)`

SetPolling sets Polling field to given value.

### HasPolling

`func (o *ScriptAPI) HasPolling() bool`

HasPolling returns a boolean if a field has been set.

### GetPropagationLabels

`func (o *ScriptAPI) GetPropagationLabels() []string`

GetPropagationLabels returns the PropagationLabels field if non-nil, zero value otherwise.

### GetPropagationLabelsOk

`func (o *ScriptAPI) GetPropagationLabelsOk() (*[]string, bool)`

GetPropagationLabelsOk returns a tuple with the PropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationLabels

`func (o *ScriptAPI) SetPropagationLabels(v []string)`

SetPropagationLabels sets PropagationLabels field to given value.

### HasPropagationLabels

`func (o *ScriptAPI) HasPropagationLabels() bool`

HasPropagationLabels returns a boolean if a field has been set.

### GetRoles

`func (o *ScriptAPI) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ScriptAPI) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ScriptAPI) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ScriptAPI) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRunAs

`func (o *ScriptAPI) GetRunAs() string`

GetRunAs returns the RunAs field if non-nil, zero value otherwise.

### GetRunAsOk

`func (o *ScriptAPI) GetRunAsOk() (*string, bool)`

GetRunAsOk returns a tuple with the RunAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAs

`func (o *ScriptAPI) SetRunAs(v string)`

SetRunAs sets RunAs field to given value.

### HasRunAs

`func (o *ScriptAPI) HasRunAs() bool`

HasRunAs returns a boolean if a field has been set.

### GetScriptTarget

`func (o *ScriptAPI) GetScriptTarget() int64`

GetScriptTarget returns the ScriptTarget field if non-nil, zero value otherwise.

### GetScriptTargetOk

`func (o *ScriptAPI) GetScriptTargetOk() (*int64, bool)`

GetScriptTargetOk returns a tuple with the ScriptTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptTarget

`func (o *ScriptAPI) SetScriptTarget(v int64)`

SetScriptTarget sets ScriptTarget field to given value.

### HasScriptTarget

`func (o *ScriptAPI) HasScriptTarget() bool`

HasScriptTarget returns a boolean if a field has been set.

### GetSystem

`func (o *ScriptAPI) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *ScriptAPI) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *ScriptAPI) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *ScriptAPI) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetTags

`func (o *ScriptAPI) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ScriptAPI) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ScriptAPI) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ScriptAPI) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *ScriptAPI) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScriptAPI) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScriptAPI) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ScriptAPI) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *ScriptAPI) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ScriptAPI) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ScriptAPI) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ScriptAPI) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVersion

`func (o *ScriptAPI) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ScriptAPI) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ScriptAPI) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ScriptAPI) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


