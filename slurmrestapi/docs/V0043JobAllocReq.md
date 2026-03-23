# V0043JobAllocReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hetjob** | Pointer to [**[]V0043JobDescMsg**](V0043JobDescMsg.md) |  | [optional] 
**Job** | Pointer to [**V0043JobDescMsg**](V0043JobDescMsg.md) |  | [optional] 

## Methods

### NewV0043JobAllocReq

`func NewV0043JobAllocReq() *V0043JobAllocReq`

NewV0043JobAllocReq instantiates a new V0043JobAllocReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0043JobAllocReqWithDefaults

`func NewV0043JobAllocReqWithDefaults() *V0043JobAllocReq`

NewV0043JobAllocReqWithDefaults instantiates a new V0043JobAllocReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHetjob

`func (o *V0043JobAllocReq) GetHetjob() []V0043JobDescMsg`

GetHetjob returns the Hetjob field if non-nil, zero value otherwise.

### GetHetjobOk

`func (o *V0043JobAllocReq) GetHetjobOk() (*[]V0043JobDescMsg, bool)`

GetHetjobOk returns a tuple with the Hetjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetjob

`func (o *V0043JobAllocReq) SetHetjob(v []V0043JobDescMsg)`

SetHetjob sets Hetjob field to given value.

### HasHetjob

`func (o *V0043JobAllocReq) HasHetjob() bool`

HasHetjob returns a boolean if a field has been set.

### GetJob

`func (o *V0043JobAllocReq) GetJob() V0043JobDescMsg`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *V0043JobAllocReq) GetJobOk() (*V0043JobDescMsg, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *V0043JobAllocReq) SetJob(v V0043JobDescMsg)`

SetJob sets Job field to given value.

### HasJob

`func (o *V0043JobAllocReq) HasJob() bool`

HasJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


