# V0042AssocRecSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Arbitrary comment | [optional] 
**Defaultqos** | Pointer to **string** | Default QOS | [optional] 
**Grpjobs** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**Grpjobsaccrue** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**Grpsubmitjobs** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**Grptres** | Pointer to [**[]V0042Tres**](V0042Tres.md) |  | [optional] 
**Grptresmins** | Pointer to [**[]V0042Tres**](V0042Tres.md) |  | [optional] 
**Grptresrunmins** | Pointer to [**[]V0042Tres**](V0042Tres.md) |  | [optional] 
**Grpwall** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**Maxjobs** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**Maxjobsaccrue** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**Maxsubmitjobs** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**Maxtresminsperjob** | Pointer to [**[]V0042Tres**](V0042Tres.md) |  | [optional] 
**Maxtresrunmins** | Pointer to [**[]V0042Tres**](V0042Tres.md) |  | [optional] 
**Maxtresperjob** | Pointer to [**[]V0042Tres**](V0042Tres.md) |  | [optional] 
**Maxtrespernode** | Pointer to [**[]V0042Tres**](V0042Tres.md) |  | [optional] 
**Maxwalldurationperjob** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**Minpriothresh** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**Parent** | Pointer to **string** | Name of parent account | [optional] 
**Priority** | Pointer to [**V0042Uint32NoValStruct**](V0042Uint32NoValStruct.md) |  | [optional] 
**Qoslevel** | Pointer to **[]string** | List of QOS names | [optional] 
**Fairshare** | Pointer to **int32** | Allocated shares used for fairshare calculation | [optional] 

## Methods

### NewV0042AssocRecSet

`func NewV0042AssocRecSet() *V0042AssocRecSet`

NewV0042AssocRecSet instantiates a new V0042AssocRecSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0042AssocRecSetWithDefaults

`func NewV0042AssocRecSetWithDefaults() *V0042AssocRecSet`

NewV0042AssocRecSetWithDefaults instantiates a new V0042AssocRecSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *V0042AssocRecSet) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0042AssocRecSet) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0042AssocRecSet) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0042AssocRecSet) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDefaultqos

`func (o *V0042AssocRecSet) GetDefaultqos() string`

GetDefaultqos returns the Defaultqos field if non-nil, zero value otherwise.

### GetDefaultqosOk

`func (o *V0042AssocRecSet) GetDefaultqosOk() (*string, bool)`

GetDefaultqosOk returns a tuple with the Defaultqos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultqos

`func (o *V0042AssocRecSet) SetDefaultqos(v string)`

SetDefaultqos sets Defaultqos field to given value.

### HasDefaultqos

`func (o *V0042AssocRecSet) HasDefaultqos() bool`

HasDefaultqos returns a boolean if a field has been set.

### GetGrpjobs

`func (o *V0042AssocRecSet) GetGrpjobs() V0042Uint32NoValStruct`

GetGrpjobs returns the Grpjobs field if non-nil, zero value otherwise.

### GetGrpjobsOk

`func (o *V0042AssocRecSet) GetGrpjobsOk() (*V0042Uint32NoValStruct, bool)`

GetGrpjobsOk returns a tuple with the Grpjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpjobs

`func (o *V0042AssocRecSet) SetGrpjobs(v V0042Uint32NoValStruct)`

SetGrpjobs sets Grpjobs field to given value.

### HasGrpjobs

`func (o *V0042AssocRecSet) HasGrpjobs() bool`

HasGrpjobs returns a boolean if a field has been set.

### GetGrpjobsaccrue

`func (o *V0042AssocRecSet) GetGrpjobsaccrue() V0042Uint32NoValStruct`

GetGrpjobsaccrue returns the Grpjobsaccrue field if non-nil, zero value otherwise.

### GetGrpjobsaccrueOk

`func (o *V0042AssocRecSet) GetGrpjobsaccrueOk() (*V0042Uint32NoValStruct, bool)`

GetGrpjobsaccrueOk returns a tuple with the Grpjobsaccrue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpjobsaccrue

`func (o *V0042AssocRecSet) SetGrpjobsaccrue(v V0042Uint32NoValStruct)`

SetGrpjobsaccrue sets Grpjobsaccrue field to given value.

### HasGrpjobsaccrue

`func (o *V0042AssocRecSet) HasGrpjobsaccrue() bool`

HasGrpjobsaccrue returns a boolean if a field has been set.

### GetGrpsubmitjobs

`func (o *V0042AssocRecSet) GetGrpsubmitjobs() V0042Uint32NoValStruct`

GetGrpsubmitjobs returns the Grpsubmitjobs field if non-nil, zero value otherwise.

### GetGrpsubmitjobsOk

`func (o *V0042AssocRecSet) GetGrpsubmitjobsOk() (*V0042Uint32NoValStruct, bool)`

GetGrpsubmitjobsOk returns a tuple with the Grpsubmitjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpsubmitjobs

`func (o *V0042AssocRecSet) SetGrpsubmitjobs(v V0042Uint32NoValStruct)`

SetGrpsubmitjobs sets Grpsubmitjobs field to given value.

### HasGrpsubmitjobs

`func (o *V0042AssocRecSet) HasGrpsubmitjobs() bool`

HasGrpsubmitjobs returns a boolean if a field has been set.

### GetGrptres

`func (o *V0042AssocRecSet) GetGrptres() []V0042Tres`

GetGrptres returns the Grptres field if non-nil, zero value otherwise.

### GetGrptresOk

`func (o *V0042AssocRecSet) GetGrptresOk() (*[]V0042Tres, bool)`

GetGrptresOk returns a tuple with the Grptres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrptres

`func (o *V0042AssocRecSet) SetGrptres(v []V0042Tres)`

SetGrptres sets Grptres field to given value.

### HasGrptres

`func (o *V0042AssocRecSet) HasGrptres() bool`

HasGrptres returns a boolean if a field has been set.

### GetGrptresmins

`func (o *V0042AssocRecSet) GetGrptresmins() []V0042Tres`

GetGrptresmins returns the Grptresmins field if non-nil, zero value otherwise.

### GetGrptresminsOk

`func (o *V0042AssocRecSet) GetGrptresminsOk() (*[]V0042Tres, bool)`

GetGrptresminsOk returns a tuple with the Grptresmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrptresmins

`func (o *V0042AssocRecSet) SetGrptresmins(v []V0042Tres)`

SetGrptresmins sets Grptresmins field to given value.

### HasGrptresmins

`func (o *V0042AssocRecSet) HasGrptresmins() bool`

HasGrptresmins returns a boolean if a field has been set.

### GetGrptresrunmins

`func (o *V0042AssocRecSet) GetGrptresrunmins() []V0042Tres`

GetGrptresrunmins returns the Grptresrunmins field if non-nil, zero value otherwise.

### GetGrptresrunminsOk

`func (o *V0042AssocRecSet) GetGrptresrunminsOk() (*[]V0042Tres, bool)`

GetGrptresrunminsOk returns a tuple with the Grptresrunmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrptresrunmins

`func (o *V0042AssocRecSet) SetGrptresrunmins(v []V0042Tres)`

SetGrptresrunmins sets Grptresrunmins field to given value.

### HasGrptresrunmins

`func (o *V0042AssocRecSet) HasGrptresrunmins() bool`

HasGrptresrunmins returns a boolean if a field has been set.

### GetGrpwall

`func (o *V0042AssocRecSet) GetGrpwall() V0042Uint32NoValStruct`

GetGrpwall returns the Grpwall field if non-nil, zero value otherwise.

### GetGrpwallOk

`func (o *V0042AssocRecSet) GetGrpwallOk() (*V0042Uint32NoValStruct, bool)`

GetGrpwallOk returns a tuple with the Grpwall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpwall

`func (o *V0042AssocRecSet) SetGrpwall(v V0042Uint32NoValStruct)`

SetGrpwall sets Grpwall field to given value.

### HasGrpwall

`func (o *V0042AssocRecSet) HasGrpwall() bool`

HasGrpwall returns a boolean if a field has been set.

### GetMaxjobs

`func (o *V0042AssocRecSet) GetMaxjobs() V0042Uint32NoValStruct`

GetMaxjobs returns the Maxjobs field if non-nil, zero value otherwise.

### GetMaxjobsOk

`func (o *V0042AssocRecSet) GetMaxjobsOk() (*V0042Uint32NoValStruct, bool)`

GetMaxjobsOk returns a tuple with the Maxjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxjobs

`func (o *V0042AssocRecSet) SetMaxjobs(v V0042Uint32NoValStruct)`

SetMaxjobs sets Maxjobs field to given value.

### HasMaxjobs

`func (o *V0042AssocRecSet) HasMaxjobs() bool`

HasMaxjobs returns a boolean if a field has been set.

### GetMaxjobsaccrue

`func (o *V0042AssocRecSet) GetMaxjobsaccrue() V0042Uint32NoValStruct`

GetMaxjobsaccrue returns the Maxjobsaccrue field if non-nil, zero value otherwise.

### GetMaxjobsaccrueOk

`func (o *V0042AssocRecSet) GetMaxjobsaccrueOk() (*V0042Uint32NoValStruct, bool)`

GetMaxjobsaccrueOk returns a tuple with the Maxjobsaccrue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxjobsaccrue

`func (o *V0042AssocRecSet) SetMaxjobsaccrue(v V0042Uint32NoValStruct)`

SetMaxjobsaccrue sets Maxjobsaccrue field to given value.

### HasMaxjobsaccrue

`func (o *V0042AssocRecSet) HasMaxjobsaccrue() bool`

HasMaxjobsaccrue returns a boolean if a field has been set.

### GetMaxsubmitjobs

`func (o *V0042AssocRecSet) GetMaxsubmitjobs() V0042Uint32NoValStruct`

GetMaxsubmitjobs returns the Maxsubmitjobs field if non-nil, zero value otherwise.

### GetMaxsubmitjobsOk

`func (o *V0042AssocRecSet) GetMaxsubmitjobsOk() (*V0042Uint32NoValStruct, bool)`

GetMaxsubmitjobsOk returns a tuple with the Maxsubmitjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxsubmitjobs

`func (o *V0042AssocRecSet) SetMaxsubmitjobs(v V0042Uint32NoValStruct)`

SetMaxsubmitjobs sets Maxsubmitjobs field to given value.

### HasMaxsubmitjobs

`func (o *V0042AssocRecSet) HasMaxsubmitjobs() bool`

HasMaxsubmitjobs returns a boolean if a field has been set.

### GetMaxtresminsperjob

`func (o *V0042AssocRecSet) GetMaxtresminsperjob() []V0042Tres`

GetMaxtresminsperjob returns the Maxtresminsperjob field if non-nil, zero value otherwise.

### GetMaxtresminsperjobOk

`func (o *V0042AssocRecSet) GetMaxtresminsperjobOk() (*[]V0042Tres, bool)`

GetMaxtresminsperjobOk returns a tuple with the Maxtresminsperjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxtresminsperjob

`func (o *V0042AssocRecSet) SetMaxtresminsperjob(v []V0042Tres)`

SetMaxtresminsperjob sets Maxtresminsperjob field to given value.

### HasMaxtresminsperjob

`func (o *V0042AssocRecSet) HasMaxtresminsperjob() bool`

HasMaxtresminsperjob returns a boolean if a field has been set.

### GetMaxtresrunmins

`func (o *V0042AssocRecSet) GetMaxtresrunmins() []V0042Tres`

GetMaxtresrunmins returns the Maxtresrunmins field if non-nil, zero value otherwise.

### GetMaxtresrunminsOk

`func (o *V0042AssocRecSet) GetMaxtresrunminsOk() (*[]V0042Tres, bool)`

GetMaxtresrunminsOk returns a tuple with the Maxtresrunmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxtresrunmins

`func (o *V0042AssocRecSet) SetMaxtresrunmins(v []V0042Tres)`

SetMaxtresrunmins sets Maxtresrunmins field to given value.

### HasMaxtresrunmins

`func (o *V0042AssocRecSet) HasMaxtresrunmins() bool`

HasMaxtresrunmins returns a boolean if a field has been set.

### GetMaxtresperjob

`func (o *V0042AssocRecSet) GetMaxtresperjob() []V0042Tres`

GetMaxtresperjob returns the Maxtresperjob field if non-nil, zero value otherwise.

### GetMaxtresperjobOk

`func (o *V0042AssocRecSet) GetMaxtresperjobOk() (*[]V0042Tres, bool)`

GetMaxtresperjobOk returns a tuple with the Maxtresperjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxtresperjob

`func (o *V0042AssocRecSet) SetMaxtresperjob(v []V0042Tres)`

SetMaxtresperjob sets Maxtresperjob field to given value.

### HasMaxtresperjob

`func (o *V0042AssocRecSet) HasMaxtresperjob() bool`

HasMaxtresperjob returns a boolean if a field has been set.

### GetMaxtrespernode

`func (o *V0042AssocRecSet) GetMaxtrespernode() []V0042Tres`

GetMaxtrespernode returns the Maxtrespernode field if non-nil, zero value otherwise.

### GetMaxtrespernodeOk

`func (o *V0042AssocRecSet) GetMaxtrespernodeOk() (*[]V0042Tres, bool)`

GetMaxtrespernodeOk returns a tuple with the Maxtrespernode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxtrespernode

`func (o *V0042AssocRecSet) SetMaxtrespernode(v []V0042Tres)`

SetMaxtrespernode sets Maxtrespernode field to given value.

### HasMaxtrespernode

`func (o *V0042AssocRecSet) HasMaxtrespernode() bool`

HasMaxtrespernode returns a boolean if a field has been set.

### GetMaxwalldurationperjob

`func (o *V0042AssocRecSet) GetMaxwalldurationperjob() V0042Uint32NoValStruct`

GetMaxwalldurationperjob returns the Maxwalldurationperjob field if non-nil, zero value otherwise.

### GetMaxwalldurationperjobOk

`func (o *V0042AssocRecSet) GetMaxwalldurationperjobOk() (*V0042Uint32NoValStruct, bool)`

GetMaxwalldurationperjobOk returns a tuple with the Maxwalldurationperjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxwalldurationperjob

`func (o *V0042AssocRecSet) SetMaxwalldurationperjob(v V0042Uint32NoValStruct)`

SetMaxwalldurationperjob sets Maxwalldurationperjob field to given value.

### HasMaxwalldurationperjob

`func (o *V0042AssocRecSet) HasMaxwalldurationperjob() bool`

HasMaxwalldurationperjob returns a boolean if a field has been set.

### GetMinpriothresh

`func (o *V0042AssocRecSet) GetMinpriothresh() V0042Uint32NoValStruct`

GetMinpriothresh returns the Minpriothresh field if non-nil, zero value otherwise.

### GetMinpriothreshOk

`func (o *V0042AssocRecSet) GetMinpriothreshOk() (*V0042Uint32NoValStruct, bool)`

GetMinpriothreshOk returns a tuple with the Minpriothresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinpriothresh

`func (o *V0042AssocRecSet) SetMinpriothresh(v V0042Uint32NoValStruct)`

SetMinpriothresh sets Minpriothresh field to given value.

### HasMinpriothresh

`func (o *V0042AssocRecSet) HasMinpriothresh() bool`

HasMinpriothresh returns a boolean if a field has been set.

### GetParent

`func (o *V0042AssocRecSet) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *V0042AssocRecSet) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *V0042AssocRecSet) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *V0042AssocRecSet) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPriority

`func (o *V0042AssocRecSet) GetPriority() V0042Uint32NoValStruct`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0042AssocRecSet) GetPriorityOk() (*V0042Uint32NoValStruct, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0042AssocRecSet) SetPriority(v V0042Uint32NoValStruct)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0042AssocRecSet) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQoslevel

`func (o *V0042AssocRecSet) GetQoslevel() []string`

GetQoslevel returns the Qoslevel field if non-nil, zero value otherwise.

### GetQoslevelOk

`func (o *V0042AssocRecSet) GetQoslevelOk() (*[]string, bool)`

GetQoslevelOk returns a tuple with the Qoslevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoslevel

`func (o *V0042AssocRecSet) SetQoslevel(v []string)`

SetQoslevel sets Qoslevel field to given value.

### HasQoslevel

`func (o *V0042AssocRecSet) HasQoslevel() bool`

HasQoslevel returns a boolean if a field has been set.

### GetFairshare

`func (o *V0042AssocRecSet) GetFairshare() int32`

GetFairshare returns the Fairshare field if non-nil, zero value otherwise.

### GetFairshareOk

`func (o *V0042AssocRecSet) GetFairshareOk() (*int32, bool)`

GetFairshareOk returns a tuple with the Fairshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFairshare

`func (o *V0042AssocRecSet) SetFairshare(v int32)`

SetFairshare sets Fairshare field to given value.

### HasFairshare

`func (o *V0042AssocRecSet) HasFairshare() bool`

HasFairshare returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


