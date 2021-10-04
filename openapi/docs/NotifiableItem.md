# NotifiableItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bcc** | Pointer to [**AdvanceArg**](AdvanceArg.md) |  | [optional] 
**Body** | Pointer to [**AdvanceArg**](AdvanceArg.md) |  | [optional] 
**Cc** | Pointer to [**AdvanceArg**](AdvanceArg.md) |  | [optional] 
**DefaultOption** | Pointer to **string** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**Methods** | Pointer to **[]string** |  | [optional] 
**ReplyOptions** | Pointer to **[]string** |  | [optional] 
**Subject** | Pointer to [**AdvanceArg**](AdvanceArg.md) |  | [optional] 
**Timings** | Pointer to [**NotifyTimings**](NotifyTimings.md) |  | [optional] 
**To** | Pointer to [**AdvanceArg**](AdvanceArg.md) |  | [optional] 

## Methods

### NewNotifiableItem

`func NewNotifiableItem() *NotifiableItem`

NewNotifiableItem instantiates a new NotifiableItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifiableItemWithDefaults

`func NewNotifiableItemWithDefaults() *NotifiableItem`

NewNotifiableItemWithDefaults instantiates a new NotifiableItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBcc

`func (o *NotifiableItem) GetBcc() AdvanceArg`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *NotifiableItem) GetBccOk() (*AdvanceArg, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *NotifiableItem) SetBcc(v AdvanceArg)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *NotifiableItem) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetBody

`func (o *NotifiableItem) GetBody() AdvanceArg`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *NotifiableItem) GetBodyOk() (*AdvanceArg, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *NotifiableItem) SetBody(v AdvanceArg)`

SetBody sets Body field to given value.

### HasBody

`func (o *NotifiableItem) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCc

`func (o *NotifiableItem) GetCc() AdvanceArg`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *NotifiableItem) GetCcOk() (*AdvanceArg, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *NotifiableItem) SetCc(v AdvanceArg)`

SetCc sets Cc field to given value.

### HasCc

`func (o *NotifiableItem) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetDefaultOption

`func (o *NotifiableItem) GetDefaultOption() string`

GetDefaultOption returns the DefaultOption field if non-nil, zero value otherwise.

### GetDefaultOptionOk

`func (o *NotifiableItem) GetDefaultOptionOk() (*string, bool)`

GetDefaultOptionOk returns a tuple with the DefaultOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOption

`func (o *NotifiableItem) SetDefaultOption(v string)`

SetDefaultOption sets DefaultOption field to given value.

### HasDefaultOption

`func (o *NotifiableItem) HasDefaultOption() bool`

HasDefaultOption returns a boolean if a field has been set.

### GetFormat

`func (o *NotifiableItem) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *NotifiableItem) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *NotifiableItem) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *NotifiableItem) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetMethods

`func (o *NotifiableItem) GetMethods() []string`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *NotifiableItem) GetMethodsOk() (*[]string, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *NotifiableItem) SetMethods(v []string)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *NotifiableItem) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### GetReplyOptions

`func (o *NotifiableItem) GetReplyOptions() []string`

GetReplyOptions returns the ReplyOptions field if non-nil, zero value otherwise.

### GetReplyOptionsOk

`func (o *NotifiableItem) GetReplyOptionsOk() (*[]string, bool)`

GetReplyOptionsOk returns a tuple with the ReplyOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyOptions

`func (o *NotifiableItem) SetReplyOptions(v []string)`

SetReplyOptions sets ReplyOptions field to given value.

### HasReplyOptions

`func (o *NotifiableItem) HasReplyOptions() bool`

HasReplyOptions returns a boolean if a field has been set.

### GetSubject

`func (o *NotifiableItem) GetSubject() AdvanceArg`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *NotifiableItem) GetSubjectOk() (*AdvanceArg, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *NotifiableItem) SetSubject(v AdvanceArg)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *NotifiableItem) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTimings

`func (o *NotifiableItem) GetTimings() NotifyTimings`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *NotifiableItem) GetTimingsOk() (*NotifyTimings, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimings

`func (o *NotifiableItem) SetTimings(v NotifyTimings)`

SetTimings sets Timings field to given value.

### HasTimings

`func (o *NotifiableItem) HasTimings() bool`

HasTimings returns a boolean if a field has been set.

### GetTo

`func (o *NotifiableItem) GetTo() AdvanceArg`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *NotifiableItem) GetToOk() (*AdvanceArg, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *NotifiableItem) SetTo(v AdvanceArg)`

SetTo sets To field to given value.

### HasTo

`func (o *NotifiableItem) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


