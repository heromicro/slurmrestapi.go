# V0041OpenapiSlurmdbdConfigRespQosInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Arbitrary description | [optional] 
**Flags** | Pointer to **[]string** | Flags, to avoid modifying current values specify NOT_SET | [optional] 
**Id** | Pointer to **int32** | Unique ID | [optional] 
**Limits** | Pointer to [**V0041OpenapiSlurmdbdConfigRespQosInnerLimits**](V0041OpenapiSlurmdbdConfigRespQosInnerLimits.md) |  | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Preempt** | Pointer to [**V0041OpenapiSlurmdbdConfigRespQosInnerPreempt**](V0041OpenapiSlurmdbdConfigRespQosInnerPreempt.md) |  | [optional] 
**Priority** | Pointer to [**V0041OpenapiSlurmdbdConfigRespQosInnerPriority**](V0041OpenapiSlurmdbdConfigRespQosInnerPriority.md) |  | [optional] 
**UsageFactor** | Pointer to [**V0041OpenapiSlurmdbdConfigRespQosInnerUsageFactor**](V0041OpenapiSlurmdbdConfigRespQosInnerUsageFactor.md) |  | [optional] 
**UsageThreshold** | Pointer to [**V0041OpenapiSlurmdbdConfigRespQosInnerUsageThreshold**](V0041OpenapiSlurmdbdConfigRespQosInnerUsageThreshold.md) |  | [optional] 

## Methods

### NewV0041OpenapiSlurmdbdConfigRespQosInner

`func NewV0041OpenapiSlurmdbdConfigRespQosInner() *V0041OpenapiSlurmdbdConfigRespQosInner`

NewV0041OpenapiSlurmdbdConfigRespQosInner instantiates a new V0041OpenapiSlurmdbdConfigRespQosInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSlurmdbdConfigRespQosInnerWithDefaults

`func NewV0041OpenapiSlurmdbdConfigRespQosInnerWithDefaults() *V0041OpenapiSlurmdbdConfigRespQosInner`

NewV0041OpenapiSlurmdbdConfigRespQosInnerWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespQosInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFlags

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetId

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLimits

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) GetLimits() V0041OpenapiSlurmdbdConfigRespQosInnerLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) GetLimitsOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) SetLimits(v V0041OpenapiSlurmdbdConfigRespQosInnerLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetName

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreempt

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) GetPreempt() V0041OpenapiSlurmdbdConfigRespQosInnerPreempt`

GetPreempt returns the Preempt field if non-nil, zero value otherwise.

### GetPreemptOk

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) GetPreemptOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerPreempt, bool)`

GetPreemptOk returns a tuple with the Preempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreempt

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) SetPreempt(v V0041OpenapiSlurmdbdConfigRespQosInnerPreempt)`

SetPreempt sets Preempt field to given value.

### HasPreempt

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) HasPreempt() bool`

HasPreempt returns a boolean if a field has been set.

### GetPriority

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) GetPriority() V0041OpenapiSlurmdbdConfigRespQosInnerPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) GetPriorityOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) SetPriority(v V0041OpenapiSlurmdbdConfigRespQosInnerPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetUsageFactor

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) GetUsageFactor() V0041OpenapiSlurmdbdConfigRespQosInnerUsageFactor`

GetUsageFactor returns the UsageFactor field if non-nil, zero value otherwise.

### GetUsageFactorOk

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) GetUsageFactorOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerUsageFactor, bool)`

GetUsageFactorOk returns a tuple with the UsageFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageFactor

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) SetUsageFactor(v V0041OpenapiSlurmdbdConfigRespQosInnerUsageFactor)`

SetUsageFactor sets UsageFactor field to given value.

### HasUsageFactor

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) HasUsageFactor() bool`

HasUsageFactor returns a boolean if a field has been set.

### GetUsageThreshold

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) GetUsageThreshold() V0041OpenapiSlurmdbdConfigRespQosInnerUsageThreshold`

GetUsageThreshold returns the UsageThreshold field if non-nil, zero value otherwise.

### GetUsageThresholdOk

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) GetUsageThresholdOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerUsageThreshold, bool)`

GetUsageThresholdOk returns a tuple with the UsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageThreshold

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) SetUsageThreshold(v V0041OpenapiSlurmdbdConfigRespQosInnerUsageThreshold)`

SetUsageThreshold sets UsageThreshold field to given value.

### HasUsageThreshold

`func (o *V0041OpenapiSlurmdbdConfigRespQosInner) HasUsageThreshold() bool`

HasUsageThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


