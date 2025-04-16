# V0042QosPreempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to **[]string** |  | [optional] 
**Mode** | Pointer to **[]string** |  | [optional] 
**ExemptTime** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 

## Methods

### NewV0042QosPreempt

`func NewV0042QosPreempt() *V0042QosPreempt`

NewV0042QosPreempt instantiates a new V0042QosPreempt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042QosPreemptWithDefaults

`func NewV0042QosPreemptWithDefaults() *V0042QosPreempt`

NewV0042QosPreemptWithDefaults instantiates a new V0042QosPreempt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *V0042QosPreempt) GetList() []string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *V0042QosPreempt) GetListOk() (*[]string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *V0042QosPreempt) SetList(v []string)`

SetList sets List field to given value.

### HasList

`func (o *V0042QosPreempt) HasList() bool`

HasList returns a boolean if a field has been set.

### GetMode

`func (o *V0042QosPreempt) GetMode() []string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *V0042QosPreempt) GetModeOk() (*[]string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *V0042QosPreempt) SetMode(v []string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *V0042QosPreempt) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetExemptTime

`func (o *V0042QosPreempt) GetExemptTime() V0042Uint32NoValStruct`

GetExemptTime returns the ExemptTime field if non-nil, zero value otherwise.

### GetExemptTimeOk

`func (o *V0042QosPreempt) GetExemptTimeOk() (*V0042Uint32NoValStruct, bool)`

GetExemptTimeOk returns a tuple with the ExemptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptTime

`func (o *V0042QosPreempt) SetExemptTime(v V0042Uint32NoValStruct)`

SetExemptTime sets ExemptTime field to given value.

### HasExemptTime

`func (o *V0042QosPreempt) HasExemptTime() bool`

HasExemptTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


