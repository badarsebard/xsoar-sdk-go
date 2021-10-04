# Version

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digits** | Pointer to **[]int64** | WARNING: when adding new attributes or changing the names of the existing ones, remember to add support in UnmarshalJSON for items that were exported by msgpack. | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewVersion

`func NewVersion() *Version`

NewVersion instantiates a new Version object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionWithDefaults

`func NewVersionWithDefaults() *Version`

NewVersionWithDefaults instantiates a new Version object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigits

`func (o *Version) GetDigits() []int64`

GetDigits returns the Digits field if non-nil, zero value otherwise.

### GetDigitsOk

`func (o *Version) GetDigitsOk() (*[]int64, bool)`

GetDigitsOk returns a tuple with the Digits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigits

`func (o *Version) SetDigits(v []int64)`

SetDigits sets Digits field to given value.

### HasDigits

`func (o *Version) HasDigits() bool`

HasDigits returns a boolean if a field has been set.

### GetLabel

`func (o *Version) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Version) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Version) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Version) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


