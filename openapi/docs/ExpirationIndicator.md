# ExpirationIndicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedFeedFetchTime** | Pointer to **time.Time** |  | [optional] 
**Expiration** | Pointer to **time.Time** |  | [optional] 
**ExpirationSource** | Pointer to [**ExpirationSource**](ExpirationSource.md) |  | [optional] 
**ExpirationStatus** | Pointer to **string** |  | [optional] 
**ManualExpirationTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewExpirationIndicator

`func NewExpirationIndicator() *ExpirationIndicator`

NewExpirationIndicator instantiates a new ExpirationIndicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpirationIndicatorWithDefaults

`func NewExpirationIndicatorWithDefaults() *ExpirationIndicator`

NewExpirationIndicatorWithDefaults instantiates a new ExpirationIndicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedFeedFetchTime

`func (o *ExpirationIndicator) GetDeletedFeedFetchTime() time.Time`

GetDeletedFeedFetchTime returns the DeletedFeedFetchTime field if non-nil, zero value otherwise.

### GetDeletedFeedFetchTimeOk

`func (o *ExpirationIndicator) GetDeletedFeedFetchTimeOk() (*time.Time, bool)`

GetDeletedFeedFetchTimeOk returns a tuple with the DeletedFeedFetchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedFeedFetchTime

`func (o *ExpirationIndicator) SetDeletedFeedFetchTime(v time.Time)`

SetDeletedFeedFetchTime sets DeletedFeedFetchTime field to given value.

### HasDeletedFeedFetchTime

`func (o *ExpirationIndicator) HasDeletedFeedFetchTime() bool`

HasDeletedFeedFetchTime returns a boolean if a field has been set.

### GetExpiration

`func (o *ExpirationIndicator) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *ExpirationIndicator) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *ExpirationIndicator) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *ExpirationIndicator) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetExpirationSource

`func (o *ExpirationIndicator) GetExpirationSource() ExpirationSource`

GetExpirationSource returns the ExpirationSource field if non-nil, zero value otherwise.

### GetExpirationSourceOk

`func (o *ExpirationIndicator) GetExpirationSourceOk() (*ExpirationSource, bool)`

GetExpirationSourceOk returns a tuple with the ExpirationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationSource

`func (o *ExpirationIndicator) SetExpirationSource(v ExpirationSource)`

SetExpirationSource sets ExpirationSource field to given value.

### HasExpirationSource

`func (o *ExpirationIndicator) HasExpirationSource() bool`

HasExpirationSource returns a boolean if a field has been set.

### GetExpirationStatus

`func (o *ExpirationIndicator) GetExpirationStatus() string`

GetExpirationStatus returns the ExpirationStatus field if non-nil, zero value otherwise.

### GetExpirationStatusOk

`func (o *ExpirationIndicator) GetExpirationStatusOk() (*string, bool)`

GetExpirationStatusOk returns a tuple with the ExpirationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationStatus

`func (o *ExpirationIndicator) SetExpirationStatus(v string)`

SetExpirationStatus sets ExpirationStatus field to given value.

### HasExpirationStatus

`func (o *ExpirationIndicator) HasExpirationStatus() bool`

HasExpirationStatus returns a boolean if a field has been set.

### GetManualExpirationTime

`func (o *ExpirationIndicator) GetManualExpirationTime() time.Time`

GetManualExpirationTime returns the ManualExpirationTime field if non-nil, zero value otherwise.

### GetManualExpirationTimeOk

`func (o *ExpirationIndicator) GetManualExpirationTimeOk() (*time.Time, bool)`

GetManualExpirationTimeOk returns a tuple with the ManualExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualExpirationTime

`func (o *ExpirationIndicator) SetManualExpirationTime(v time.Time)`

SetManualExpirationTime sets ManualExpirationTime field to given value.

### HasManualExpirationTime

`func (o *ExpirationIndicator) HasManualExpirationTime() bool`

HasManualExpirationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


