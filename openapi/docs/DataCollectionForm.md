# DataCollectionForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Expired** | Pointer to **bool** |  | [optional] 
**Questions** | Pointer to [**[]Question**](Question.md) |  | [optional] 
**Sender** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**TotalAnswers** | Pointer to **int32** |  | [optional] 

## Methods

### NewDataCollectionForm

`func NewDataCollectionForm() *DataCollectionForm`

NewDataCollectionForm instantiates a new DataCollectionForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataCollectionFormWithDefaults

`func NewDataCollectionFormWithDefaults() *DataCollectionForm`

NewDataCollectionFormWithDefaults instantiates a new DataCollectionForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DataCollectionForm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataCollectionForm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataCollectionForm) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataCollectionForm) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpired

`func (o *DataCollectionForm) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *DataCollectionForm) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *DataCollectionForm) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *DataCollectionForm) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetQuestions

`func (o *DataCollectionForm) GetQuestions() []Question`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *DataCollectionForm) GetQuestionsOk() (*[]Question, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *DataCollectionForm) SetQuestions(v []Question)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *DataCollectionForm) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.

### GetSender

`func (o *DataCollectionForm) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *DataCollectionForm) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *DataCollectionForm) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *DataCollectionForm) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetTitle

`func (o *DataCollectionForm) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DataCollectionForm) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DataCollectionForm) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DataCollectionForm) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTotalAnswers

`func (o *DataCollectionForm) GetTotalAnswers() int32`

GetTotalAnswers returns the TotalAnswers field if non-nil, zero value otherwise.

### GetTotalAnswersOk

`func (o *DataCollectionForm) GetTotalAnswersOk() (*int32, bool)`

GetTotalAnswersOk returns a tuple with the TotalAnswers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAnswers

`func (o *DataCollectionForm) SetTotalAnswers(v int32)`

SetTotalAnswers sets TotalAnswers field to given value.

### HasTotalAnswers

`func (o *DataCollectionForm) HasTotalAnswers() bool`

HasTotalAnswers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


