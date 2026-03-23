# V0042JobAllocReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hetjob** | Pointer to [**[]V0042JobDescMsg**](V0042JobDescMsg.md) |  | [optional] 
**Job** | Pointer to [**V0042JobDescMsg**](V0042JobDescMsg.md) |  | [optional] 

## Methods

### NewV0042JobAllocReq

`func NewV0042JobAllocReq() *V0042JobAllocReq`

NewV0042JobAllocReq instantiates a new V0042JobAllocReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042JobAllocReqWithDefaults

`func NewV0042JobAllocReqWithDefaults() *V0042JobAllocReq`

NewV0042JobAllocReqWithDefaults instantiates a new V0042JobAllocReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHetjob

`func (o *V0042JobAllocReq) GetHetjob() []V0042JobDescMsg`

GetHetjob returns the Hetjob field if non-nil, zero value otherwise.

### GetHetjobOk

`func (o *V0042JobAllocReq) GetHetjobOk() (*[]V0042JobDescMsg, bool)`

GetHetjobOk returns a tuple with the Hetjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetjob

`func (o *V0042JobAllocReq) SetHetjob(v []V0042JobDescMsg)`

SetHetjob sets Hetjob field to given value.

### HasHetjob

`func (o *V0042JobAllocReq) HasHetjob() bool`

HasHetjob returns a boolean if a field has been set.

### GetJob

`func (o *V0042JobAllocReq) GetJob() V0042JobDescMsg`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *V0042JobAllocReq) GetJobOk() (*V0042JobDescMsg, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *V0042JobAllocReq) SetJob(v V0042JobDescMsg)`

SetJob sets Job field to given value.

### HasJob

`func (o *V0042JobAllocReq) HasJob() bool`

HasJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


