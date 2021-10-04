# AuditResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audits** | Pointer to [**[]Audit**](Audit.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewAuditResult

`func NewAuditResult() *AuditResult`

NewAuditResult instantiates a new AuditResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditResultWithDefaults

`func NewAuditResultWithDefaults() *AuditResult`

NewAuditResultWithDefaults instantiates a new AuditResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudits

`func (o *AuditResult) GetAudits() []Audit`

GetAudits returns the Audits field if non-nil, zero value otherwise.

### GetAuditsOk

`func (o *AuditResult) GetAuditsOk() (*[]Audit, bool)`

GetAuditsOk returns a tuple with the Audits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudits

`func (o *AuditResult) SetAudits(v []Audit)`

SetAudits sets Audits field to given value.

### HasAudits

`func (o *AuditResult) HasAudits() bool`

HasAudits returns a boolean if a field has been set.

### GetTotal

`func (o *AuditResult) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AuditResult) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AuditResult) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AuditResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


