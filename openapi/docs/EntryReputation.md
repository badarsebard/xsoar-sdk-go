# EntryReputation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Highlights** | Pointer to [**map[string]map[string][]Location**](map.md) |  | [optional] 
**ReputationsData** | Pointer to [**[]ReputationData**](ReputationData.md) |  | [optional] 

## Methods

### NewEntryReputation

`func NewEntryReputation() *EntryReputation`

NewEntryReputation instantiates a new EntryReputation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryReputationWithDefaults

`func NewEntryReputationWithDefaults() *EntryReputation`

NewEntryReputationWithDefaults instantiates a new EntryReputation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHighlights

`func (o *EntryReputation) GetHighlights() map[string]map[string][]Location`

GetHighlights returns the Highlights field if non-nil, zero value otherwise.

### GetHighlightsOk

`func (o *EntryReputation) GetHighlightsOk() (*map[string]map[string][]Location, bool)`

GetHighlightsOk returns a tuple with the Highlights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlights

`func (o *EntryReputation) SetHighlights(v map[string]map[string][]Location)`

SetHighlights sets Highlights field to given value.

### HasHighlights

`func (o *EntryReputation) HasHighlights() bool`

HasHighlights returns a boolean if a field has been set.

### GetReputationsData

`func (o *EntryReputation) GetReputationsData() []ReputationData`

GetReputationsData returns the ReputationsData field if non-nil, zero value otherwise.

### GetReputationsDataOk

`func (o *EntryReputation) GetReputationsDataOk() (*[]ReputationData, bool)`

GetReputationsDataOk returns a tuple with the ReputationsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReputationsData

`func (o *EntryReputation) SetReputationsData(v []ReputationData)`

SetReputationsData sets ReputationsData field to given value.

### HasReputationsData

`func (o *EntryReputation) HasReputationsData() bool`

HasReputationsData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


