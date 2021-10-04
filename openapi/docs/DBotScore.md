# DBotScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** |  | [optional] 
**ContentFormat** | Pointer to **string** |  | [optional] 
**Context** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**IsTypedIndicator** | Pointer to **bool** |  | [optional] 
**Reliability** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **int64** |  | [optional] 
**ScoreChangeTimestamp** | Pointer to **time.Time** | We need to track when the score changes to know if we need to re-calculate the overall score | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewDBotScore

`func NewDBotScore() *DBotScore`

NewDBotScore instantiates a new DBotScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDBotScoreWithDefaults

`func NewDBotScoreWithDefaults() *DBotScore`

NewDBotScoreWithDefaults instantiates a new DBotScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *DBotScore) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DBotScore) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DBotScore) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *DBotScore) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentFormat

`func (o *DBotScore) GetContentFormat() string`

GetContentFormat returns the ContentFormat field if non-nil, zero value otherwise.

### GetContentFormatOk

`func (o *DBotScore) GetContentFormatOk() (*string, bool)`

GetContentFormatOk returns a tuple with the ContentFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFormat

`func (o *DBotScore) SetContentFormat(v string)`

SetContentFormat sets ContentFormat field to given value.

### HasContentFormat

`func (o *DBotScore) HasContentFormat() bool`

HasContentFormat returns a boolean if a field has been set.

### GetContext

`func (o *DBotScore) GetContext() map[string]map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DBotScore) GetContextOk() (*map[string]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DBotScore) SetContext(v map[string]map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *DBotScore) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetIsTypedIndicator

`func (o *DBotScore) GetIsTypedIndicator() bool`

GetIsTypedIndicator returns the IsTypedIndicator field if non-nil, zero value otherwise.

### GetIsTypedIndicatorOk

`func (o *DBotScore) GetIsTypedIndicatorOk() (*bool, bool)`

GetIsTypedIndicatorOk returns a tuple with the IsTypedIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTypedIndicator

`func (o *DBotScore) SetIsTypedIndicator(v bool)`

SetIsTypedIndicator sets IsTypedIndicator field to given value.

### HasIsTypedIndicator

`func (o *DBotScore) HasIsTypedIndicator() bool`

HasIsTypedIndicator returns a boolean if a field has been set.

### GetReliability

`func (o *DBotScore) GetReliability() string`

GetReliability returns the Reliability field if non-nil, zero value otherwise.

### GetReliabilityOk

`func (o *DBotScore) GetReliabilityOk() (*string, bool)`

GetReliabilityOk returns a tuple with the Reliability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliability

`func (o *DBotScore) SetReliability(v string)`

SetReliability sets Reliability field to given value.

### HasReliability

`func (o *DBotScore) HasReliability() bool`

HasReliability returns a boolean if a field has been set.

### GetScore

`func (o *DBotScore) GetScore() int64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *DBotScore) GetScoreOk() (*int64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *DBotScore) SetScore(v int64)`

SetScore sets Score field to given value.

### HasScore

`func (o *DBotScore) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetScoreChangeTimestamp

`func (o *DBotScore) GetScoreChangeTimestamp() time.Time`

GetScoreChangeTimestamp returns the ScoreChangeTimestamp field if non-nil, zero value otherwise.

### GetScoreChangeTimestampOk

`func (o *DBotScore) GetScoreChangeTimestampOk() (*time.Time, bool)`

GetScoreChangeTimestampOk returns a tuple with the ScoreChangeTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreChangeTimestamp

`func (o *DBotScore) SetScoreChangeTimestamp(v time.Time)`

SetScoreChangeTimestamp sets ScoreChangeTimestamp field to given value.

### HasScoreChangeTimestamp

`func (o *DBotScore) HasScoreChangeTimestamp() bool`

HasScoreChangeTimestamp returns a boolean if a field has been set.

### GetTimestamp

`func (o *DBotScore) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DBotScore) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DBotScore) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DBotScore) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *DBotScore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DBotScore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DBotScore) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DBotScore) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


