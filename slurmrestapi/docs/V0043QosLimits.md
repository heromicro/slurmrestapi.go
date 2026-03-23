# V0043QosLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GraceTime** | Pointer to **int32** | GraceTime - Preemption grace time in seconds to be extended to a job which has been selected for preemption | [optional] 
**Max** | Pointer to [**V0043QosLimitsMax**](V0043QosLimitsMax.md) |  | [optional] 
**Factor** | Pointer to [**V0043Float64NoValStruct**](V0043Float64NoValStruct.md) |  | [optional] 
**Min** | Pointer to [**V0043QosLimitsMin**](V0043QosLimitsMin.md) |  | [optional] 

## Methods

### NewV0043QosLimits

`func NewV0043QosLimits() *V0043QosLimits`

NewV0043QosLimits instantiates a new V0043QosLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043QosLimitsWithDefaults

`func NewV0043QosLimitsWithDefaults() *V0043QosLimits`

NewV0043QosLimitsWithDefaults instantiates a new V0043QosLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGraceTime

`func (o *V0043QosLimits) GetGraceTime() int32`

GetGraceTime returns the GraceTime field if non-nil, zero value otherwise.

### GetGraceTimeOk

`func (o *V0043QosLimits) GetGraceTimeOk() (*int32, bool)`

GetGraceTimeOk returns a tuple with the GraceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceTime

`func (o *V0043QosLimits) SetGraceTime(v int32)`

SetGraceTime sets GraceTime field to given value.

### HasGraceTime

`func (o *V0043QosLimits) HasGraceTime() bool`

HasGraceTime returns a boolean if a field has been set.

### GetMax

`func (o *V0043QosLimits) GetMax() V0043QosLimitsMax`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *V0043QosLimits) GetMaxOk() (*V0043QosLimitsMax, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *V0043QosLimits) SetMax(v V0043QosLimitsMax)`

SetMax sets Max field to given value.

### HasMax

`func (o *V0043QosLimits) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetFactor

`func (o *V0043QosLimits) GetFactor() V0043Float64NoValStruct`

GetFactor returns the Factor field if non-nil, zero value otherwise.

### GetFactorOk

`func (o *V0043QosLimits) GetFactorOk() (*V0043Float64NoValStruct, bool)`

GetFactorOk returns a tuple with the Factor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactor

`func (o *V0043QosLimits) SetFactor(v V0043Float64NoValStruct)`

SetFactor sets Factor field to given value.

### HasFactor

`func (o *V0043QosLimits) HasFactor() bool`

HasFactor returns a boolean if a field has been set.

### GetMin

`func (o *V0043QosLimits) GetMin() V0043QosLimitsMin`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *V0043QosLimits) GetMinOk() (*V0043QosLimitsMin, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *V0043QosLimits) SetMin(v V0043QosLimitsMin)`

SetMin sets Min field to given value.

### HasMin

`func (o *V0043QosLimits) HasMin() bool`

HasMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


