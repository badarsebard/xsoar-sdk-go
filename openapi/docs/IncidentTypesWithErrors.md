# IncidentTypesWithErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**IncidentTypes** | Pointer to [**[]IncidentType**](IncidentType.md) |  | [optional] 

## Methods

### NewIncidentTypesWithErrors

`func NewIncidentTypesWithErrors() *IncidentTypesWithErrors`

NewIncidentTypesWithErrors instantiates a new IncidentTypesWithErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentTypesWithErrorsWithDefaults

`func NewIncidentTypesWithErrorsWithDefaults() *IncidentTypesWithErrors`

NewIncidentTypesWithErrorsWithDefaults instantiates a new IncidentTypesWithErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *IncidentTypesWithErrors) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *IncidentTypesWithErrors) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *IncidentTypesWithErrors) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *IncidentTypesWithErrors) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIncidentTypes

`func (o *IncidentTypesWithErrors) GetIncidentTypes() []IncidentType`

GetIncidentTypes returns the IncidentTypes field if non-nil, zero value otherwise.

### GetIncidentTypesOk

`func (o *IncidentTypesWithErrors) GetIncidentTypesOk() (*[]IncidentType, bool)`

GetIncidentTypesOk returns a tuple with the IncidentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentTypes

`func (o *IncidentTypesWithErrors) SetIncidentTypes(v []IncidentType)`

SetIncidentTypes sets IncidentTypes field to given value.

### HasIncidentTypes

`func (o *IncidentTypesWithErrors) HasIncidentTypes() bool`

HasIncidentTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


