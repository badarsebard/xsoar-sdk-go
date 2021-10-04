/*
Cortex XSOAR API

This is the public REST API to integrate with the Cortex XSOAR server. HTTP request can be sent using any HTTP-client.  For an example dedicated client take a look at: https://github.com/demisto/demisto-py.  Requests must include API-key that can be generated in the Cortex XSOAR web client under 'Settings' -> 'Integrations' -> 'API keys'   Optimistic Locking and Versioning\\:  When using Cortex XSOAR REST API, you will need to make sure to work on the latest version of the item (incident, entry, etc.), otherwise, you will get a DB version error (which not allow you to override a newer item). In addition, you can pass 'version\\: -1' to force data override (make sure that other users data might be lost).  Assume that Alice and Bob both read the same data from Cortex XSOAR server, then they both changed the data, and then both tried to write the new versions back to the server. Whose changes should be saved? Alice’s? Bob’s? To solve this, each data item in Cortex XSOAR has a numeric incremental version. If Alice saved an item with version 4 and Bob trying to save the same item with version 3, Cortex XSOAR will rollback Bob request and returns a DB version conflict error. Bob will need to get the latest item and work on it so Alice work will not get lost.  Example request using 'curl'\\:  ``` curl 'https://hostname:443/incidents/search' -H 'content-type: application/json' -H 'accept: application/json' -H 'Authorization: <API Key goes here>' --data-binary '{\"filter\":{\"query\":\"-status:closed -category:job\",\"period\":{\"by\":\"day\",\"fromValue\":7}}}' --compressed ```

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// DataCollectionForm struct for DataCollectionForm
type DataCollectionForm struct {
	Description  *string     `json:"description,omitempty"`
	Expired      *bool       `json:"expired,omitempty"`
	Questions    *[]Question `json:"questions,omitempty"`
	Sender       *string     `json:"sender,omitempty"`
	Title        *string     `json:"title,omitempty"`
	TotalAnswers *int32      `json:"totalAnswers,omitempty"`
}

// NewDataCollectionForm instantiates a new DataCollectionForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataCollectionForm() *DataCollectionForm {
	this := DataCollectionForm{}
	return &this
}

// NewDataCollectionFormWithDefaults instantiates a new DataCollectionForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataCollectionFormWithDefaults() *DataCollectionForm {
	this := DataCollectionForm{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DataCollectionForm) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCollectionForm) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DataCollectionForm) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DataCollectionForm) SetDescription(v string) {
	o.Description = &v
}

// GetExpired returns the Expired field value if set, zero value otherwise.
func (o *DataCollectionForm) GetExpired() bool {
	if o == nil || o.Expired == nil {
		var ret bool
		return ret
	}
	return *o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCollectionForm) GetExpiredOk() (*bool, bool) {
	if o == nil || o.Expired == nil {
		return nil, false
	}
	return o.Expired, true
}

// HasExpired returns a boolean if a field has been set.
func (o *DataCollectionForm) HasExpired() bool {
	if o != nil && o.Expired != nil {
		return true
	}

	return false
}

// SetExpired gets a reference to the given bool and assigns it to the Expired field.
func (o *DataCollectionForm) SetExpired(v bool) {
	o.Expired = &v
}

// GetQuestions returns the Questions field value if set, zero value otherwise.
func (o *DataCollectionForm) GetQuestions() []Question {
	if o == nil || o.Questions == nil {
		var ret []Question
		return ret
	}
	return *o.Questions
}

// GetQuestionsOk returns a tuple with the Questions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCollectionForm) GetQuestionsOk() (*[]Question, bool) {
	if o == nil || o.Questions == nil {
		return nil, false
	}
	return o.Questions, true
}

// HasQuestions returns a boolean if a field has been set.
func (o *DataCollectionForm) HasQuestions() bool {
	if o != nil && o.Questions != nil {
		return true
	}

	return false
}

// SetQuestions gets a reference to the given []Question and assigns it to the Questions field.
func (o *DataCollectionForm) SetQuestions(v []Question) {
	o.Questions = &v
}

// GetSender returns the Sender field value if set, zero value otherwise.
func (o *DataCollectionForm) GetSender() string {
	if o == nil || o.Sender == nil {
		var ret string
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCollectionForm) GetSenderOk() (*string, bool) {
	if o == nil || o.Sender == nil {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *DataCollectionForm) HasSender() bool {
	if o != nil && o.Sender != nil {
		return true
	}

	return false
}

// SetSender gets a reference to the given string and assigns it to the Sender field.
func (o *DataCollectionForm) SetSender(v string) {
	o.Sender = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *DataCollectionForm) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCollectionForm) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *DataCollectionForm) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *DataCollectionForm) SetTitle(v string) {
	o.Title = &v
}

// GetTotalAnswers returns the TotalAnswers field value if set, zero value otherwise.
func (o *DataCollectionForm) GetTotalAnswers() int32 {
	if o == nil || o.TotalAnswers == nil {
		var ret int32
		return ret
	}
	return *o.TotalAnswers
}

// GetTotalAnswersOk returns a tuple with the TotalAnswers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCollectionForm) GetTotalAnswersOk() (*int32, bool) {
	if o == nil || o.TotalAnswers == nil {
		return nil, false
	}
	return o.TotalAnswers, true
}

// HasTotalAnswers returns a boolean if a field has been set.
func (o *DataCollectionForm) HasTotalAnswers() bool {
	if o != nil && o.TotalAnswers != nil {
		return true
	}

	return false
}

// SetTotalAnswers gets a reference to the given int32 and assigns it to the TotalAnswers field.
func (o *DataCollectionForm) SetTotalAnswers(v int32) {
	o.TotalAnswers = &v
}

func (o DataCollectionForm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Expired != nil {
		toSerialize["expired"] = o.Expired
	}
	if o.Questions != nil {
		toSerialize["questions"] = o.Questions
	}
	if o.Sender != nil {
		toSerialize["sender"] = o.Sender
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.TotalAnswers != nil {
		toSerialize["totalAnswers"] = o.TotalAnswers
	}
	return json.Marshal(toSerialize)
}

type NullableDataCollectionForm struct {
	value *DataCollectionForm
	isSet bool
}

func (v NullableDataCollectionForm) Get() *DataCollectionForm {
	return v.value
}

func (v *NullableDataCollectionForm) Set(val *DataCollectionForm) {
	v.value = val
	v.isSet = true
}

func (v NullableDataCollectionForm) IsSet() bool {
	return v.isSet
}

func (v *NullableDataCollectionForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataCollectionForm(val *DataCollectionForm) *NullableDataCollectionForm {
	return &NullableDataCollectionForm{value: val, isSet: true}
}

func (v NullableDataCollectionForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataCollectionForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
