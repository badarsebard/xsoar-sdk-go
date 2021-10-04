# NewDockerImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | Pointer to **string** |  | [optional] 
**Dependencies** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Packages** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNewDockerImageRequest

`func NewNewDockerImageRequest() *NewDockerImageRequest`

NewNewDockerImageRequest instantiates a new NewDockerImageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewDockerImageRequestWithDefaults

`func NewNewDockerImageRequestWithDefaults() *NewDockerImageRequest`

NewNewDockerImageRequestWithDefaults instantiates a new NewDockerImageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *NewDockerImageRequest) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *NewDockerImageRequest) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *NewDockerImageRequest) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *NewDockerImageRequest) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetDependencies

`func (o *NewDockerImageRequest) GetDependencies() []string`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *NewDockerImageRequest) GetDependenciesOk() (*[]string, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *NewDockerImageRequest) SetDependencies(v []string)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *NewDockerImageRequest) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetName

`func (o *NewDockerImageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewDockerImageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewDockerImageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NewDockerImageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPackages

`func (o *NewDockerImageRequest) GetPackages() []string`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *NewDockerImageRequest) GetPackagesOk() (*[]string, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *NewDockerImageRequest) SetPackages(v []string)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *NewDockerImageRequest) HasPackages() bool`

HasPackages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


