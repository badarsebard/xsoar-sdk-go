# Info

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **float64** |  | [optional] 
**BufferSpace** | Pointer to **int64** |  | [optional] 
**Busy** | Pointer to **int64** |  | [optional] 
**HighPriorityAvailable** | Pointer to **float64** |  | [optional] 
**HighPriorityBufferSpace** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProcessInfo** | Pointer to [**[]ProcessInfo**](ProcessInfo.md) |  | [optional] 
**ShouldStop** | Pointer to **bool** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewInfo

`func NewInfo() *Info`

NewInfo instantiates a new Info object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoWithDefaults

`func NewInfoWithDefaults() *Info`

NewInfoWithDefaults instantiates a new Info object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *Info) GetAvailable() float64`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *Info) GetAvailableOk() (*float64, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *Info) SetAvailable(v float64)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *Info) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetBufferSpace

`func (o *Info) GetBufferSpace() int64`

GetBufferSpace returns the BufferSpace field if non-nil, zero value otherwise.

### GetBufferSpaceOk

`func (o *Info) GetBufferSpaceOk() (*int64, bool)`

GetBufferSpaceOk returns a tuple with the BufferSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferSpace

`func (o *Info) SetBufferSpace(v int64)`

SetBufferSpace sets BufferSpace field to given value.

### HasBufferSpace

`func (o *Info) HasBufferSpace() bool`

HasBufferSpace returns a boolean if a field has been set.

### GetBusy

`func (o *Info) GetBusy() int64`

GetBusy returns the Busy field if non-nil, zero value otherwise.

### GetBusyOk

`func (o *Info) GetBusyOk() (*int64, bool)`

GetBusyOk returns a tuple with the Busy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusy

`func (o *Info) SetBusy(v int64)`

SetBusy sets Busy field to given value.

### HasBusy

`func (o *Info) HasBusy() bool`

HasBusy returns a boolean if a field has been set.

### GetHighPriorityAvailable

`func (o *Info) GetHighPriorityAvailable() float64`

GetHighPriorityAvailable returns the HighPriorityAvailable field if non-nil, zero value otherwise.

### GetHighPriorityAvailableOk

`func (o *Info) GetHighPriorityAvailableOk() (*float64, bool)`

GetHighPriorityAvailableOk returns a tuple with the HighPriorityAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPriorityAvailable

`func (o *Info) SetHighPriorityAvailable(v float64)`

SetHighPriorityAvailable sets HighPriorityAvailable field to given value.

### HasHighPriorityAvailable

`func (o *Info) HasHighPriorityAvailable() bool`

HasHighPriorityAvailable returns a boolean if a field has been set.

### GetHighPriorityBufferSpace

`func (o *Info) GetHighPriorityBufferSpace() int64`

GetHighPriorityBufferSpace returns the HighPriorityBufferSpace field if non-nil, zero value otherwise.

### GetHighPriorityBufferSpaceOk

`func (o *Info) GetHighPriorityBufferSpaceOk() (*int64, bool)`

GetHighPriorityBufferSpaceOk returns a tuple with the HighPriorityBufferSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPriorityBufferSpace

`func (o *Info) SetHighPriorityBufferSpace(v int64)`

SetHighPriorityBufferSpace sets HighPriorityBufferSpace field to given value.

### HasHighPriorityBufferSpace

`func (o *Info) HasHighPriorityBufferSpace() bool`

HasHighPriorityBufferSpace returns a boolean if a field has been set.

### GetName

`func (o *Info) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Info) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Info) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Info) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProcessInfo

`func (o *Info) GetProcessInfo() []ProcessInfo`

GetProcessInfo returns the ProcessInfo field if non-nil, zero value otherwise.

### GetProcessInfoOk

`func (o *Info) GetProcessInfoOk() (*[]ProcessInfo, bool)`

GetProcessInfoOk returns a tuple with the ProcessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInfo

`func (o *Info) SetProcessInfo(v []ProcessInfo)`

SetProcessInfo sets ProcessInfo field to given value.

### HasProcessInfo

`func (o *Info) HasProcessInfo() bool`

HasProcessInfo returns a boolean if a field has been set.

### GetShouldStop

`func (o *Info) GetShouldStop() bool`

GetShouldStop returns the ShouldStop field if non-nil, zero value otherwise.

### GetShouldStopOk

`func (o *Info) GetShouldStopOk() (*bool, bool)`

GetShouldStopOk returns a tuple with the ShouldStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldStop

`func (o *Info) SetShouldStop(v bool)`

SetShouldStop sets ShouldStop field to given value.

### HasShouldStop

`func (o *Info) HasShouldStop() bool`

HasShouldStop returns a boolean if a field has been set.

### GetTotal

`func (o *Info) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Info) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Info) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Info) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


