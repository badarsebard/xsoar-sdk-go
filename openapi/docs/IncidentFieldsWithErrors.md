# IncidentFieldsWithErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**IncidentFields** | Pointer to [**[]IncidentField**](IncidentField.md) |  | [optional] 

## Methods

### NewIncidentFieldsWithErrors

`func NewIncidentFieldsWithErrors() *IncidentFieldsWithErrors`

NewIncidentFieldsWithErrors instantiates a new IncidentFieldsWithErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentFieldsWithErrorsWithDefaults

`func NewIncidentFieldsWithErrorsWithDefaults() *IncidentFieldsWithErrors`

NewIncidentFieldsWithErrorsWithDefaults instantiates a new IncidentFieldsWithErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *IncidentFieldsWithErrors) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *IncidentFieldsWithErrors) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *IncidentFieldsWithErrors) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *IncidentFieldsWithErrors) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIncidentFields

`func (o *IncidentFieldsWithErrors) GetIncidentFields() []IncidentField`

GetIncidentFields returns the IncidentFields field if non-nil, zero value otherwise.

### GetIncidentFieldsOk

`func (o *IncidentFieldsWithErrors) GetIncidentFieldsOk() (*[]IncidentField, bool)`

GetIncidentFieldsOk returns a tuple with the IncidentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentFields

`func (o *IncidentFieldsWithErrors) SetIncidentFields(v []IncidentField)`

SetIncidentFields sets IncidentFields field to given value.

### HasIncidentFields

`func (o *IncidentFieldsWithErrors) HasIncidentFields() bool`

HasIncidentFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


