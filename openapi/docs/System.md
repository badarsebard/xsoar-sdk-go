# System

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to [**SystemAgent**](SystemAgent.md) |  | [optional] 
**Arch** | Pointer to **string** |  | [optional] 
**Ciphers** | Pointer to **[]string** |  | [optional] 
**Credentials** | Pointer to **string** |  | [optional] 
**EngineId** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Integrationinstanceid** | Pointer to **string** |  | [optional] 
**Issharedagent** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**ServicesID** | Pointer to **string** |  | [optional] 
**Smb** | Pointer to **int64** |  | [optional] 
**Smbport** | Pointer to **int32** |  | [optional] 
**Sshkey** | Pointer to **string** |  | [optional] 
**Sshport** | Pointer to **int32** |  | [optional] 
**TerminalOptions** | Pointer to [**TerminalOptions**](TerminalOptions.md) |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Workgroup** | Pointer to **string** |  | [optional] 

## Methods

### NewSystem

`func NewSystem() *System`

NewSystem instantiates a new System object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemWithDefaults

`func NewSystemWithDefaults() *System`

NewSystemWithDefaults instantiates a new System object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *System) GetAgent() SystemAgent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *System) GetAgentOk() (*SystemAgent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *System) SetAgent(v SystemAgent)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *System) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetArch

`func (o *System) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *System) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *System) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *System) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetCiphers

`func (o *System) GetCiphers() []string`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *System) GetCiphersOk() (*[]string, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *System) SetCiphers(v []string)`

SetCiphers sets Ciphers field to given value.

### HasCiphers

`func (o *System) HasCiphers() bool`

HasCiphers returns a boolean if a field has been set.

### GetCredentials

`func (o *System) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *System) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *System) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *System) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEngineId

`func (o *System) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *System) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *System) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *System) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetHost

`func (o *System) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *System) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *System) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *System) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetIntegrationinstanceid

`func (o *System) GetIntegrationinstanceid() string`

GetIntegrationinstanceid returns the Integrationinstanceid field if non-nil, zero value otherwise.

### GetIntegrationinstanceidOk

`func (o *System) GetIntegrationinstanceidOk() (*string, bool)`

GetIntegrationinstanceidOk returns a tuple with the Integrationinstanceid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationinstanceid

`func (o *System) SetIntegrationinstanceid(v string)`

SetIntegrationinstanceid sets Integrationinstanceid field to given value.

### HasIntegrationinstanceid

`func (o *System) HasIntegrationinstanceid() bool`

HasIntegrationinstanceid returns a boolean if a field has been set.

### GetIssharedagent

`func (o *System) GetIssharedagent() bool`

GetIssharedagent returns the Issharedagent field if non-nil, zero value otherwise.

### GetIssharedagentOk

`func (o *System) GetIssharedagentOk() (*bool, bool)`

GetIssharedagentOk returns a tuple with the Issharedagent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssharedagent

`func (o *System) SetIssharedagent(v bool)`

SetIssharedagent sets Issharedagent field to given value.

### HasIssharedagent

`func (o *System) HasIssharedagent() bool`

HasIssharedagent returns a boolean if a field has been set.

### GetName

`func (o *System) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *System) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *System) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *System) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOs

`func (o *System) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *System) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *System) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *System) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPassword

`func (o *System) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *System) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *System) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *System) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetServicesID

`func (o *System) GetServicesID() string`

GetServicesID returns the ServicesID field if non-nil, zero value otherwise.

### GetServicesIDOk

`func (o *System) GetServicesIDOk() (*string, bool)`

GetServicesIDOk returns a tuple with the ServicesID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesID

`func (o *System) SetServicesID(v string)`

SetServicesID sets ServicesID field to given value.

### HasServicesID

`func (o *System) HasServicesID() bool`

HasServicesID returns a boolean if a field has been set.

### GetSmb

`func (o *System) GetSmb() int64`

GetSmb returns the Smb field if non-nil, zero value otherwise.

### GetSmbOk

`func (o *System) GetSmbOk() (*int64, bool)`

GetSmbOk returns a tuple with the Smb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmb

`func (o *System) SetSmb(v int64)`

SetSmb sets Smb field to given value.

### HasSmb

`func (o *System) HasSmb() bool`

HasSmb returns a boolean if a field has been set.

### GetSmbport

`func (o *System) GetSmbport() int32`

GetSmbport returns the Smbport field if non-nil, zero value otherwise.

### GetSmbportOk

`func (o *System) GetSmbportOk() (*int32, bool)`

GetSmbportOk returns a tuple with the Smbport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbport

`func (o *System) SetSmbport(v int32)`

SetSmbport sets Smbport field to given value.

### HasSmbport

`func (o *System) HasSmbport() bool`

HasSmbport returns a boolean if a field has been set.

### GetSshkey

`func (o *System) GetSshkey() string`

GetSshkey returns the Sshkey field if non-nil, zero value otherwise.

### GetSshkeyOk

`func (o *System) GetSshkeyOk() (*string, bool)`

GetSshkeyOk returns a tuple with the Sshkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshkey

`func (o *System) SetSshkey(v string)`

SetSshkey sets Sshkey field to given value.

### HasSshkey

`func (o *System) HasSshkey() bool`

HasSshkey returns a boolean if a field has been set.

### GetSshport

`func (o *System) GetSshport() int32`

GetSshport returns the Sshport field if non-nil, zero value otherwise.

### GetSshportOk

`func (o *System) GetSshportOk() (*int32, bool)`

GetSshportOk returns a tuple with the Sshport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshport

`func (o *System) SetSshport(v int32)`

SetSshport sets Sshport field to given value.

### HasSshport

`func (o *System) HasSshport() bool`

HasSshport returns a boolean if a field has been set.

### GetTerminalOptions

`func (o *System) GetTerminalOptions() TerminalOptions`

GetTerminalOptions returns the TerminalOptions field if non-nil, zero value otherwise.

### GetTerminalOptionsOk

`func (o *System) GetTerminalOptionsOk() (*TerminalOptions, bool)`

GetTerminalOptionsOk returns a tuple with the TerminalOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalOptions

`func (o *System) SetTerminalOptions(v TerminalOptions)`

SetTerminalOptions sets TerminalOptions field to given value.

### HasTerminalOptions

`func (o *System) HasTerminalOptions() bool`

HasTerminalOptions returns a boolean if a field has been set.

### GetUser

`func (o *System) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *System) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *System) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *System) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetWorkgroup

`func (o *System) GetWorkgroup() string`

GetWorkgroup returns the Workgroup field if non-nil, zero value otherwise.

### GetWorkgroupOk

`func (o *System) GetWorkgroupOk() (*string, bool)`

GetWorkgroupOk returns a tuple with the Workgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkgroup

`func (o *System) SetWorkgroup(v string)`

SetWorkgroup sets Workgroup field to given value.

### HasWorkgroup

`func (o *System) HasWorkgroup() bool`

HasWorkgroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


