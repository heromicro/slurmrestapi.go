# SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | TRES type (CPU, MEM, etc) | 
**Name** | Pointer to **string** | TRES name (if applicable) | [optional] 
**Id** | Pointer to **int32** | ID used in database | [optional] 
**Count** | Pointer to **int64** | TRES count (0 if listed generically) | [optional] 

## Methods

### NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner

`func NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner(type_ string, ) *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner`

NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner instantiates a new SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInnerWithDefaults

`func NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInnerWithDefaults() *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner`

NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInnerWithDefaults instantiates a new SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCount

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


