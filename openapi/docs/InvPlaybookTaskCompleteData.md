# InvPlaybookTaskCompleteData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**CalculatedDescription** | Pointer to **string** |  | [optional] 
**CalculatedTaskName** | Pointer to **string** |  | [optional] 
**CompletedBy** | Pointer to **string** |  | [optional] 
**CompletedCount** | Pointer to **int64** |  | [optional] 
**CompletedDate** | Pointer to **time.Time** |  | [optional] 
**Entries** | Pointer to **[]string** |  | [optional] 
**Input** | Pointer to **string** |  | [optional] 
**Outputs** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**PlaybookInputs** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**State** | Pointer to **string** | TaskState indicates the state of the task during the incident/investigation execution | [optional] 
**WillNotExecuteCount** | Pointer to **int64** |  | [optional] 
**WillNotExecuteReason** | Pointer to **string** |  | [optional] 

## Methods

### NewInvPlaybookTaskCompleteData

`func NewInvPlaybookTaskCompleteData() *InvPlaybookTaskCompleteData`

NewInvPlaybookTaskCompleteData instantiates a new InvPlaybookTaskCompleteData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvPlaybookTaskCompleteDataWithDefaults

`func NewInvPlaybookTaskCompleteDataWithDefaults() *InvPlaybookTaskCompleteData`

NewInvPlaybookTaskCompleteDataWithDefaults instantiates a new InvPlaybookTaskCompleteData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *InvPlaybookTaskCompleteData) GetArguments() map[string]map[string]interface{}`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *InvPlaybookTaskCompleteData) GetArgumentsOk() (*map[string]map[string]interface{}, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *InvPlaybookTaskCompleteData) SetArguments(v map[string]map[string]interface{})`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *InvPlaybookTaskCompleteData) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetCalculatedDescription

`func (o *InvPlaybookTaskCompleteData) GetCalculatedDescription() string`

GetCalculatedDescription returns the CalculatedDescription field if non-nil, zero value otherwise.

### GetCalculatedDescriptionOk

`func (o *InvPlaybookTaskCompleteData) GetCalculatedDescriptionOk() (*string, bool)`

GetCalculatedDescriptionOk returns a tuple with the CalculatedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedDescription

`func (o *InvPlaybookTaskCompleteData) SetCalculatedDescription(v string)`

SetCalculatedDescription sets CalculatedDescription field to given value.

### HasCalculatedDescription

`func (o *InvPlaybookTaskCompleteData) HasCalculatedDescription() bool`

HasCalculatedDescription returns a boolean if a field has been set.

### GetCalculatedTaskName

`func (o *InvPlaybookTaskCompleteData) GetCalculatedTaskName() string`

GetCalculatedTaskName returns the CalculatedTaskName field if non-nil, zero value otherwise.

### GetCalculatedTaskNameOk

`func (o *InvPlaybookTaskCompleteData) GetCalculatedTaskNameOk() (*string, bool)`

GetCalculatedTaskNameOk returns a tuple with the CalculatedTaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedTaskName

`func (o *InvPlaybookTaskCompleteData) SetCalculatedTaskName(v string)`

SetCalculatedTaskName sets CalculatedTaskName field to given value.

### HasCalculatedTaskName

`func (o *InvPlaybookTaskCompleteData) HasCalculatedTaskName() bool`

HasCalculatedTaskName returns a boolean if a field has been set.

### GetCompletedBy

`func (o *InvPlaybookTaskCompleteData) GetCompletedBy() string`

GetCompletedBy returns the CompletedBy field if non-nil, zero value otherwise.

### GetCompletedByOk

`func (o *InvPlaybookTaskCompleteData) GetCompletedByOk() (*string, bool)`

GetCompletedByOk returns a tuple with the CompletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedBy

`func (o *InvPlaybookTaskCompleteData) SetCompletedBy(v string)`

SetCompletedBy sets CompletedBy field to given value.

### HasCompletedBy

`func (o *InvPlaybookTaskCompleteData) HasCompletedBy() bool`

HasCompletedBy returns a boolean if a field has been set.

### GetCompletedCount

`func (o *InvPlaybookTaskCompleteData) GetCompletedCount() int64`

GetCompletedCount returns the CompletedCount field if non-nil, zero value otherwise.

### GetCompletedCountOk

`func (o *InvPlaybookTaskCompleteData) GetCompletedCountOk() (*int64, bool)`

GetCompletedCountOk returns a tuple with the CompletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedCount

`func (o *InvPlaybookTaskCompleteData) SetCompletedCount(v int64)`

SetCompletedCount sets CompletedCount field to given value.

### HasCompletedCount

`func (o *InvPlaybookTaskCompleteData) HasCompletedCount() bool`

HasCompletedCount returns a boolean if a field has been set.

### GetCompletedDate

`func (o *InvPlaybookTaskCompleteData) GetCompletedDate() time.Time`

GetCompletedDate returns the CompletedDate field if non-nil, zero value otherwise.

### GetCompletedDateOk

`func (o *InvPlaybookTaskCompleteData) GetCompletedDateOk() (*time.Time, bool)`

GetCompletedDateOk returns a tuple with the CompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDate

`func (o *InvPlaybookTaskCompleteData) SetCompletedDate(v time.Time)`

SetCompletedDate sets CompletedDate field to given value.

### HasCompletedDate

`func (o *InvPlaybookTaskCompleteData) HasCompletedDate() bool`

HasCompletedDate returns a boolean if a field has been set.

### GetEntries

`func (o *InvPlaybookTaskCompleteData) GetEntries() []string`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *InvPlaybookTaskCompleteData) GetEntriesOk() (*[]string, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *InvPlaybookTaskCompleteData) SetEntries(v []string)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *InvPlaybookTaskCompleteData) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetInput

`func (o *InvPlaybookTaskCompleteData) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *InvPlaybookTaskCompleteData) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *InvPlaybookTaskCompleteData) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *InvPlaybookTaskCompleteData) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetOutputs

`func (o *InvPlaybookTaskCompleteData) GetOutputs() map[string]map[string]interface{}`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *InvPlaybookTaskCompleteData) GetOutputsOk() (*map[string]map[string]interface{}, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *InvPlaybookTaskCompleteData) SetOutputs(v map[string]map[string]interface{})`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *InvPlaybookTaskCompleteData) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetPlaybookInputs

`func (o *InvPlaybookTaskCompleteData) GetPlaybookInputs() map[string]map[string]interface{}`

GetPlaybookInputs returns the PlaybookInputs field if non-nil, zero value otherwise.

### GetPlaybookInputsOk

`func (o *InvPlaybookTaskCompleteData) GetPlaybookInputsOk() (*map[string]map[string]interface{}, bool)`

GetPlaybookInputsOk returns a tuple with the PlaybookInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookInputs

`func (o *InvPlaybookTaskCompleteData) SetPlaybookInputs(v map[string]map[string]interface{})`

SetPlaybookInputs sets PlaybookInputs field to given value.

### HasPlaybookInputs

`func (o *InvPlaybookTaskCompleteData) HasPlaybookInputs() bool`

HasPlaybookInputs returns a boolean if a field has been set.

### GetStartDate

`func (o *InvPlaybookTaskCompleteData) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InvPlaybookTaskCompleteData) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InvPlaybookTaskCompleteData) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InvPlaybookTaskCompleteData) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetState

`func (o *InvPlaybookTaskCompleteData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InvPlaybookTaskCompleteData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InvPlaybookTaskCompleteData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *InvPlaybookTaskCompleteData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWillNotExecuteCount

`func (o *InvPlaybookTaskCompleteData) GetWillNotExecuteCount() int64`

GetWillNotExecuteCount returns the WillNotExecuteCount field if non-nil, zero value otherwise.

### GetWillNotExecuteCountOk

`func (o *InvPlaybookTaskCompleteData) GetWillNotExecuteCountOk() (*int64, bool)`

GetWillNotExecuteCountOk returns a tuple with the WillNotExecuteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWillNotExecuteCount

`func (o *InvPlaybookTaskCompleteData) SetWillNotExecuteCount(v int64)`

SetWillNotExecuteCount sets WillNotExecuteCount field to given value.

### HasWillNotExecuteCount

`func (o *InvPlaybookTaskCompleteData) HasWillNotExecuteCount() bool`

HasWillNotExecuteCount returns a boolean if a field has been set.

### GetWillNotExecuteReason

`func (o *InvPlaybookTaskCompleteData) GetWillNotExecuteReason() string`

GetWillNotExecuteReason returns the WillNotExecuteReason field if non-nil, zero value otherwise.

### GetWillNotExecuteReasonOk

`func (o *InvPlaybookTaskCompleteData) GetWillNotExecuteReasonOk() (*string, bool)`

GetWillNotExecuteReasonOk returns a tuple with the WillNotExecuteReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWillNotExecuteReason

`func (o *InvPlaybookTaskCompleteData) SetWillNotExecuteReason(v string)`

SetWillNotExecuteReason sets WillNotExecuteReason field to given value.

### HasWillNotExecuteReason

`func (o *InvPlaybookTaskCompleteData) HasWillNotExecuteReason() bool`

HasWillNotExecuteReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


