# V0042QosLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GraceTime** | Pointer to **int32** | GraceTime | [optional] 
**Max** | Pointer to [**V0042QosLimitsMax**](V0042QosLimitsMax.md) |  | [optional] 
**Factor** | Pointer to [**V0042Float64NoValStruct**](V0042Float64NoValStruct.md) |  | [optional] 
**Min** | Pointer to [**V0042QosLimitsMin**](V0042QosLimitsMin.md) |  | [optional] 

## Methods

### NewV0042QosLimits

`func NewV0042QosLimits() *V0042QosLimits`

NewV0042QosLimits instantiates a new V0042QosLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042QosLimitsWithDefaults

`func NewV0042QosLimitsWithDefaults() *V0042QosLimits`

NewV0042QosLimitsWithDefaults instantiates a new V0042QosLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGraceTime

`func (o *V0042QosLimits) GetGraceTime() int32`

GetGraceTime returns the GraceTime field if non-nil, zero value otherwise.

### GetGraceTimeOk

`func (o *V0042QosLimits) GetGraceTimeOk() (*int32, bool)`

GetGraceTimeOk returns a tuple with the GraceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceTime

`func (o *V0042QosLimits) SetGraceTime(v int32)`

SetGraceTime sets GraceTime field to given value.

### HasGraceTime

`func (o *V0042QosLimits) HasGraceTime() bool`

HasGraceTime returns a boolean if a field has been set.

### GetMax

`func (o *V0042QosLimits) GetMax() V0042QosLimitsMax`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *V0042QosLimits) GetMaxOk() (*V0042QosLimitsMax, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *V0042QosLimits) SetMax(v V0042QosLimitsMax)`

SetMax sets Max field to given value.

### HasMax

`func (o *V0042QosLimits) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetFactor

`func (o *V0042QosLimits) GetFactor() V0042Float64NoValStruct`

GetFactor returns the Factor field if non-nil, zero value otherwise.

### GetFactorOk

`func (o *V0042QosLimits) GetFactorOk() (*V0042Float64NoValStruct, bool)`

GetFactorOk returns a tuple with the Factor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactor

`func (o *V0042QosLimits) SetFactor(v V0042Float64NoValStruct)`

SetFactor sets Factor field to given value.

### HasFactor

`func (o *V0042QosLimits) HasFactor() bool`

HasFactor returns a boolean if a field has been set.

### GetMin

`func (o *V0042QosLimits) GetMin() V0042QosLimitsMin`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *V0042QosLimits) GetMinOk() (*V0042QosLimitsMin, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *V0042QosLimits) SetMin(v V0042QosLimitsMin)`

SetMin sets Min field to given value.

### HasMin

`func (o *V0042QosLimits) HasMin() bool`

HasMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


