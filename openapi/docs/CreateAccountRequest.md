# CreateAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountRoles** | Pointer to **[]string** |  | [optional] 
**HostGroupId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PropagationLabels** | Pointer to **[]string** |  | [optional] 
**SyncOnCreation** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateAccountRequest

`func NewCreateAccountRequest() *CreateAccountRequest`

NewCreateAccountRequest instantiates a new CreateAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountRequestWithDefaults

`func NewCreateAccountRequestWithDefaults() *CreateAccountRequest`

NewCreateAccountRequestWithDefaults instantiates a new CreateAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountRoles

`func (o *CreateAccountRequest) GetAccountRoles() []string`

GetAccountRoles returns the AccountRoles field if non-nil, zero value otherwise.

### GetAccountRolesOk

`func (o *CreateAccountRequest) GetAccountRolesOk() (*[]string, bool)`

GetAccountRolesOk returns a tuple with the AccountRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRoles

`func (o *CreateAccountRequest) SetAccountRoles(v []string)`

SetAccountRoles sets AccountRoles field to given value.

### HasAccountRoles

`func (o *CreateAccountRequest) HasAccountRoles() bool`

HasAccountRoles returns a boolean if a field has been set.

### GetHostGroupId

`func (o *CreateAccountRequest) GetHostGroupId() string`

GetHostGroupId returns the HostGroupId field if non-nil, zero value otherwise.

### GetHostGroupIdOk

`func (o *CreateAccountRequest) GetHostGroupIdOk() (*string, bool)`

GetHostGroupIdOk returns a tuple with the HostGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroupId

`func (o *CreateAccountRequest) SetHostGroupId(v string)`

SetHostGroupId sets HostGroupId field to given value.

### HasHostGroupId

`func (o *CreateAccountRequest) HasHostGroupId() bool`

HasHostGroupId returns a boolean if a field has been set.

### GetName

`func (o *CreateAccountRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAccountRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAccountRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateAccountRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPropagationLabels

`func (o *CreateAccountRequest) GetPropagationLabels() []string`

GetPropagationLabels returns the PropagationLabels field if non-nil, zero value otherwise.

### GetPropagationLabelsOk

`func (o *CreateAccountRequest) GetPropagationLabelsOk() (*[]string, bool)`

GetPropagationLabelsOk returns a tuple with the PropagationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationLabels

`func (o *CreateAccountRequest) SetPropagationLabels(v []string)`

SetPropagationLabels sets PropagationLabels field to given value.

### HasPropagationLabels

`func (o *CreateAccountRequest) HasPropagationLabels() bool`

HasPropagationLabels returns a boolean if a field has been set.

### GetSyncOnCreation

`func (o *CreateAccountRequest) GetSyncOnCreation() bool`

GetSyncOnCreation returns the SyncOnCreation field if non-nil, zero value otherwise.

### GetSyncOnCreationOk

`func (o *CreateAccountRequest) GetSyncOnCreationOk() (*bool, bool)`

GetSyncOnCreationOk returns a tuple with the SyncOnCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncOnCreation

`func (o *CreateAccountRequest) SetSyncOnCreation(v bool)`

SetSyncOnCreation sets SyncOnCreation field to given value.

### HasSyncOnCreation

`func (o *CreateAccountRequest) HasSyncOnCreation() bool`

HasSyncOnCreation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


