# UpdateDataBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**All** | Pointer to **bool** |  | [optional] 
**CloseNotes** | Pointer to **string** |  | [optional] 
**CloseReason** | Pointer to **string** |  | [optional] 
**Columns** | Pointer to **[]string** |  | [optional] 
**Data** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Filter** | Pointer to [**IncidentFilter**](IncidentFilter.md) |  | [optional] 
**Force** | Pointer to **bool** |  | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 
**Line** | Pointer to **string** |  | [optional] 
**OriginalIncidentId** | Pointer to **string** |  | [optional] 
**OverrideInvestigation** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateDataBatch

`func NewUpdateDataBatch() *UpdateDataBatch`

NewUpdateDataBatch instantiates a new UpdateDataBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDataBatchWithDefaults

`func NewUpdateDataBatchWithDefaults() *UpdateDataBatch`

NewUpdateDataBatchWithDefaults instantiates a new UpdateDataBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomFields

`func (o *UpdateDataBatch) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *UpdateDataBatch) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *UpdateDataBatch) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *UpdateDataBatch) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetAll

`func (o *UpdateDataBatch) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *UpdateDataBatch) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *UpdateDataBatch) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *UpdateDataBatch) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetCloseNotes

`func (o *UpdateDataBatch) GetCloseNotes() string`

GetCloseNotes returns the CloseNotes field if non-nil, zero value otherwise.

### GetCloseNotesOk

`func (o *UpdateDataBatch) GetCloseNotesOk() (*string, bool)`

GetCloseNotesOk returns a tuple with the CloseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseNotes

`func (o *UpdateDataBatch) SetCloseNotes(v string)`

SetCloseNotes sets CloseNotes field to given value.

### HasCloseNotes

`func (o *UpdateDataBatch) HasCloseNotes() bool`

HasCloseNotes returns a boolean if a field has been set.

### GetCloseReason

`func (o *UpdateDataBatch) GetCloseReason() string`

GetCloseReason returns the CloseReason field if non-nil, zero value otherwise.

### GetCloseReasonOk

`func (o *UpdateDataBatch) GetCloseReasonOk() (*string, bool)`

GetCloseReasonOk returns a tuple with the CloseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseReason

`func (o *UpdateDataBatch) SetCloseReason(v string)`

SetCloseReason sets CloseReason field to given value.

### HasCloseReason

`func (o *UpdateDataBatch) HasCloseReason() bool`

HasCloseReason returns a boolean if a field has been set.

### GetColumns

`func (o *UpdateDataBatch) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *UpdateDataBatch) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *UpdateDataBatch) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *UpdateDataBatch) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetData

`func (o *UpdateDataBatch) GetData() map[string]map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateDataBatch) GetDataOk() (*map[string]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateDataBatch) SetData(v map[string]map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *UpdateDataBatch) HasData() bool`

HasData returns a boolean if a field has been set.

### GetFilter

`func (o *UpdateDataBatch) GetFilter() IncidentFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *UpdateDataBatch) GetFilterOk() (*IncidentFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *UpdateDataBatch) SetFilter(v IncidentFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *UpdateDataBatch) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetForce

`func (o *UpdateDataBatch) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *UpdateDataBatch) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *UpdateDataBatch) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *UpdateDataBatch) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetIds

`func (o *UpdateDataBatch) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *UpdateDataBatch) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *UpdateDataBatch) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *UpdateDataBatch) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetLine

`func (o *UpdateDataBatch) GetLine() string`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *UpdateDataBatch) GetLineOk() (*string, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *UpdateDataBatch) SetLine(v string)`

SetLine sets Line field to given value.

### HasLine

`func (o *UpdateDataBatch) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetOriginalIncidentId

`func (o *UpdateDataBatch) GetOriginalIncidentId() string`

GetOriginalIncidentId returns the OriginalIncidentId field if non-nil, zero value otherwise.

### GetOriginalIncidentIdOk

`func (o *UpdateDataBatch) GetOriginalIncidentIdOk() (*string, bool)`

GetOriginalIncidentIdOk returns a tuple with the OriginalIncidentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalIncidentId

`func (o *UpdateDataBatch) SetOriginalIncidentId(v string)`

SetOriginalIncidentId sets OriginalIncidentId field to given value.

### HasOriginalIncidentId

`func (o *UpdateDataBatch) HasOriginalIncidentId() bool`

HasOriginalIncidentId returns a boolean if a field has been set.

### GetOverrideInvestigation

`func (o *UpdateDataBatch) GetOverrideInvestigation() bool`

GetOverrideInvestigation returns the OverrideInvestigation field if non-nil, zero value otherwise.

### GetOverrideInvestigationOk

`func (o *UpdateDataBatch) GetOverrideInvestigationOk() (*bool, bool)`

GetOverrideInvestigationOk returns a tuple with the OverrideInvestigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideInvestigation

`func (o *UpdateDataBatch) SetOverrideInvestigation(v bool)`

SetOverrideInvestigation sets OverrideInvestigation field to given value.

### HasOverrideInvestigation

`func (o *UpdateDataBatch) HasOverrideInvestigation() bool`

HasOverrideInvestigation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


