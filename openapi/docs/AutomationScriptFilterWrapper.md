# AutomationScriptFilterWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**GenericStringFilter**](GenericStringFilter.md) |  | [optional] 
**SavePassword** | Pointer to **bool** |  | [optional] 
**Script** | Pointer to [**AutomationScript**](AutomationScript.md) |  | [optional] 

## Methods

### NewAutomationScriptFilterWrapper

`func NewAutomationScriptFilterWrapper() *AutomationScriptFilterWrapper`

NewAutomationScriptFilterWrapper instantiates a new AutomationScriptFilterWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationScriptFilterWrapperWithDefaults

`func NewAutomationScriptFilterWrapperWithDefaults() *AutomationScriptFilterWrapper`

NewAutomationScriptFilterWrapperWithDefaults instantiates a new AutomationScriptFilterWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *AutomationScriptFilterWrapper) GetFilter() GenericStringFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AutomationScriptFilterWrapper) GetFilterOk() (*GenericStringFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AutomationScriptFilterWrapper) SetFilter(v GenericStringFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AutomationScriptFilterWrapper) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSavePassword

`func (o *AutomationScriptFilterWrapper) GetSavePassword() bool`

GetSavePassword returns the SavePassword field if non-nil, zero value otherwise.

### GetSavePasswordOk

`func (o *AutomationScriptFilterWrapper) GetSavePasswordOk() (*bool, bool)`

GetSavePasswordOk returns a tuple with the SavePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavePassword

`func (o *AutomationScriptFilterWrapper) SetSavePassword(v bool)`

SetSavePassword sets SavePassword field to given value.

### HasSavePassword

`func (o *AutomationScriptFilterWrapper) HasSavePassword() bool`

HasSavePassword returns a boolean if a field has been set.

### GetScript

`func (o *AutomationScriptFilterWrapper) GetScript() AutomationScript`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *AutomationScriptFilterWrapper) GetScriptOk() (*AutomationScript, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *AutomationScriptFilterWrapper) SetScript(v AutomationScript)`

SetScript sets Script field to given value.

### HasScript

`func (o *AutomationScriptFilterWrapper) HasScript() bool`

HasScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


