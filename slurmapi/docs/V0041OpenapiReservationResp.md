# V0041OpenapiReservationResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reservations** | [**[]V0041OpenapiReservationRespReservationsInner**](V0041OpenapiReservationRespReservationsInner.md) | List of reservations | 
**LastUpdate** | [**V0041OpenapiReservationRespLastUpdate**](V0041OpenapiReservationRespLastUpdate.md) |  | 
**Meta** | Pointer to [**SlurmV0041GetShares200ResponseMeta**](SlurmV0041GetShares200ResponseMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]SlurmV0041GetShares200ResponseErrorsInner**](SlurmV0041GetShares200ResponseErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]SlurmV0041GetShares200ResponseWarningsInner**](SlurmV0041GetShares200ResponseWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewV0041OpenapiReservationResp

`func NewV0041OpenapiReservationResp(reservations []V0041OpenapiReservationRespReservationsInner, lastUpdate V0041OpenapiReservationRespLastUpdate, ) *V0041OpenapiReservationResp`

NewV0041OpenapiReservationResp instantiates a new V0041OpenapiReservationResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiReservationRespWithDefaults

`func NewV0041OpenapiReservationRespWithDefaults() *V0041OpenapiReservationResp`

NewV0041OpenapiReservationRespWithDefaults instantiates a new V0041OpenapiReservationResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReservations

`func (o *V0041OpenapiReservationResp) GetReservations() []V0041OpenapiReservationRespReservationsInner`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *V0041OpenapiReservationResp) GetReservationsOk() (*[]V0041OpenapiReservationRespReservationsInner, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *V0041OpenapiReservationResp) SetReservations(v []V0041OpenapiReservationRespReservationsInner)`

SetReservations sets Reservations field to given value.


### GetLastUpdate

`func (o *V0041OpenapiReservationResp) GetLastUpdate() V0041OpenapiReservationRespLastUpdate`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *V0041OpenapiReservationResp) GetLastUpdateOk() (*V0041OpenapiReservationRespLastUpdate, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *V0041OpenapiReservationResp) SetLastUpdate(v V0041OpenapiReservationRespLastUpdate)`

SetLastUpdate sets LastUpdate field to given value.


### GetMeta

`func (o *V0041OpenapiReservationResp) GetMeta() SlurmV0041GetShares200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiReservationResp) GetMetaOk() (*SlurmV0041GetShares200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiReservationResp) SetMeta(v SlurmV0041GetShares200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiReservationResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiReservationResp) GetErrors() []SlurmV0041GetShares200ResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiReservationResp) GetErrorsOk() (*[]SlurmV0041GetShares200ResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiReservationResp) SetErrors(v []SlurmV0041GetShares200ResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiReservationResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiReservationResp) GetWarnings() []SlurmV0041GetShares200ResponseWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiReservationResp) GetWarningsOk() (*[]SlurmV0041GetShares200ResponseWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiReservationResp) SetWarnings(v []SlurmV0041GetShares200ResponseWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiReservationResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


