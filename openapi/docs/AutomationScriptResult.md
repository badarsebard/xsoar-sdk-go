# AutomationScriptResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PythonEnabled** | Pointer to **bool** |  | [optional] 
**Scripts** | Pointer to [**[]ScriptAPI**](ScriptAPI.md) |  | [optional] 
**SelectedScript** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Suggestions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAutomationScriptResult

`func NewAutomationScriptResult() *AutomationScriptResult`

NewAutomationScriptResult instantiates a new AutomationScriptResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationScriptResultWithDefaults

`func NewAutomationScriptResultWithDefaults() *AutomationScriptResult`

NewAutomationScriptResultWithDefaults instantiates a new AutomationScriptResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPythonEnabled

`func (o *AutomationScriptResult) GetPythonEnabled() bool`

GetPythonEnabled returns the PythonEnabled field if non-nil, zero value otherwise.

### GetPythonEnabledOk

`func (o *AutomationScriptResult) GetPythonEnabledOk() (*bool, bool)`

GetPythonEnabledOk returns a tuple with the PythonEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonEnabled

`func (o *AutomationScriptResult) SetPythonEnabled(v bool)`

SetPythonEnabled sets PythonEnabled field to given value.

### HasPythonEnabled

`func (o *AutomationScriptResult) HasPythonEnabled() bool`

HasPythonEnabled returns a boolean if a field has been set.

### GetScripts

`func (o *AutomationScriptResult) GetScripts() []ScriptAPI`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *AutomationScriptResult) GetScriptsOk() (*[]ScriptAPI, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *AutomationScriptResult) SetScripts(v []ScriptAPI)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *AutomationScriptResult) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetSelectedScript

`func (o *AutomationScriptResult) GetSelectedScript() map[string]map[string]interface{}`

GetSelectedScript returns the SelectedScript field if non-nil, zero value otherwise.

### GetSelectedScriptOk

`func (o *AutomationScriptResult) GetSelectedScriptOk() (*map[string]map[string]interface{}, bool)`

GetSelectedScriptOk returns a tuple with the SelectedScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedScript

`func (o *AutomationScriptResult) SetSelectedScript(v map[string]map[string]interface{})`

SetSelectedScript sets SelectedScript field to given value.

### HasSelectedScript

`func (o *AutomationScriptResult) HasSelectedScript() bool`

HasSelectedScript returns a boolean if a field has been set.

### GetSuggestions

`func (o *AutomationScriptResult) GetSuggestions() []string`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *AutomationScriptResult) GetSuggestionsOk() (*[]string, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *AutomationScriptResult) SetSuggestions(v []string)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *AutomationScriptResult) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


