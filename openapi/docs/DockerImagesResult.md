# DockerImagesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | Pointer to [**[]DockerImage**](DockerImage.md) |  | [optional] 

## Methods

### NewDockerImagesResult

`func NewDockerImagesResult() *DockerImagesResult`

NewDockerImagesResult instantiates a new DockerImagesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerImagesResultWithDefaults

`func NewDockerImagesResultWithDefaults() *DockerImagesResult`

NewDockerImagesResultWithDefaults instantiates a new DockerImagesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *DockerImagesResult) GetImages() []DockerImage`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *DockerImagesResult) GetImagesOk() (*[]DockerImage, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *DockerImagesResult) SetImages(v []DockerImage)`

SetImages sets Images field to given value.

### HasImages

`func (o *DockerImagesResult) HasImages() bool`

HasImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


