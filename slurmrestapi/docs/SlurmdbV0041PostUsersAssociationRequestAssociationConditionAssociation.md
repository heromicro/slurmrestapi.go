# SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Arbitrary comment | [optional] 
**Defaultqos** | Pointer to **string** | Default QOS | [optional] 
**Grpjobs** | Pointer to [**SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrpjobs**](SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrpjobs.md) |  | [optional] 
**Grpjobsaccrue** | Pointer to [**SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrpjobsaccrue**](SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrpjobsaccrue.md) |  | [optional] 
**Grpsubmitjobs** | Pointer to [**SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrpsubmitjobs**](SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrpsubmitjobs.md) |  | [optional] 
**Grptres** | Pointer to [**[]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner**](SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner.md) | Maximum number of TRES able to be allocated by running jobs in this association and its children | [optional] 
**Grptresmins** | Pointer to [**[]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner**](SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner.md) | Total number of TRES minutes that can possibly be used by past, present and future jobs in this association and its children | [optional] 
**Grptresrunmins** | Pointer to [**[]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner**](SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner.md) | Maximum number of TRES minutes able to be allocated by running jobs in this association and its children | [optional] 
**Grpwall** | Pointer to [**SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrpwall**](SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrpwall.md) |  | [optional] 
**Maxjobs** | Pointer to [**SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxjobs**](SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxjobs.md) |  | [optional] 
**Maxjobsaccrue** | Pointer to [**SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxjobsaccrue**](SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxjobsaccrue.md) |  | [optional] 
**Maxsubmitjobs** | Pointer to [**SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs**](SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs.md) |  | [optional] 
**Maxtresminsperjob** | Pointer to [**[]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner**](SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner.md) | Maximum number of TRES minutes each job is able to use in this association | [optional] 
**Maxtresrunmins** | Pointer to [**[]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner**](SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner.md) | Maximum number of TRES minutes able to be allocated by running jobs in this association | [optional] 
**Maxtresperjob** | Pointer to [**[]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner**](SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner.md) | Maximum number of TRES each job is able to use in this association | [optional] 
**Maxtrespernode** | Pointer to [**[]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner**](SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner.md) | Maximum number of TRES each node is able to use | [optional] 
**Maxwalldurationperjob** | Pointer to [**SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxwalldurationperjob**](SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxwalldurationperjob.md) |  | [optional] 
**Minpriothresh** | Pointer to [**SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMinpriothresh**](SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMinpriothresh.md) |  | [optional] 
**Parent** | Pointer to **string** | Name of parent account | [optional] 
**Priority** | Pointer to [**SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationPriority**](SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationPriority.md) |  | [optional] 
**Qoslevel** | Pointer to **[]string** | List of available QOS names | [optional] 
**Fairshare** | Pointer to **int32** | Allocated shares used for fairshare calculation | [optional] 

## Methods

### NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation

`func NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation() *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation`

NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation instantiates a new SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationWithDefaults

`func NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationWithDefaults() *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation`

NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationWithDefaults instantiates a new SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDefaultqos

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetDefaultqos() string`

GetDefaultqos returns the Defaultqos field if non-nil, zero value otherwise.

### GetDefaultqosOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetDefaultqosOk() (*string, bool)`

GetDefaultqosOk returns a tuple with the Defaultqos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultqos

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetDefaultqos(v string)`

SetDefaultqos sets Defaultqos field to given value.

### HasDefaultqos

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasDefaultqos() bool`

HasDefaultqos returns a boolean if a field has been set.

### GetGrpjobs

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetGrpjobs() SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrpjobs`

GetGrpjobs returns the Grpjobs field if non-nil, zero value otherwise.

### GetGrpjobsOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetGrpjobsOk() (*SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrpjobs, bool)`

GetGrpjobsOk returns a tuple with the Grpjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpjobs

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetGrpjobs(v SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrpjobs)`

SetGrpjobs sets Grpjobs field to given value.

### HasGrpjobs

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasGrpjobs() bool`

HasGrpjobs returns a boolean if a field has been set.

### GetGrpjobsaccrue

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetGrpjobsaccrue() SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrpjobsaccrue`

GetGrpjobsaccrue returns the Grpjobsaccrue field if non-nil, zero value otherwise.

### GetGrpjobsaccrueOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetGrpjobsaccrueOk() (*SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrpjobsaccrue, bool)`

GetGrpjobsaccrueOk returns a tuple with the Grpjobsaccrue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpjobsaccrue

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetGrpjobsaccrue(v SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrpjobsaccrue)`

SetGrpjobsaccrue sets Grpjobsaccrue field to given value.

### HasGrpjobsaccrue

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasGrpjobsaccrue() bool`

HasGrpjobsaccrue returns a boolean if a field has been set.

### GetGrpsubmitjobs

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetGrpsubmitjobs() SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrpsubmitjobs`

GetGrpsubmitjobs returns the Grpsubmitjobs field if non-nil, zero value otherwise.

### GetGrpsubmitjobsOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetGrpsubmitjobsOk() (*SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrpsubmitjobs, bool)`

GetGrpsubmitjobsOk returns a tuple with the Grpsubmitjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpsubmitjobs

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetGrpsubmitjobs(v SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrpsubmitjobs)`

SetGrpsubmitjobs sets Grpsubmitjobs field to given value.

### HasGrpsubmitjobs

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasGrpsubmitjobs() bool`

HasGrpsubmitjobs returns a boolean if a field has been set.

### GetGrptres

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetGrptres() []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner`

GetGrptres returns the Grptres field if non-nil, zero value otherwise.

### GetGrptresOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetGrptresOk() (*[]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner, bool)`

GetGrptresOk returns a tuple with the Grptres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrptres

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetGrptres(v []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner)`

SetGrptres sets Grptres field to given value.

### HasGrptres

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasGrptres() bool`

HasGrptres returns a boolean if a field has been set.

### GetGrptresmins

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetGrptresmins() []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner`

GetGrptresmins returns the Grptresmins field if non-nil, zero value otherwise.

### GetGrptresminsOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetGrptresminsOk() (*[]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner, bool)`

GetGrptresminsOk returns a tuple with the Grptresmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrptresmins

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetGrptresmins(v []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner)`

SetGrptresmins sets Grptresmins field to given value.

### HasGrptresmins

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasGrptresmins() bool`

HasGrptresmins returns a boolean if a field has been set.

### GetGrptresrunmins

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetGrptresrunmins() []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner`

GetGrptresrunmins returns the Grptresrunmins field if non-nil, zero value otherwise.

### GetGrptresrunminsOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetGrptresrunminsOk() (*[]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner, bool)`

GetGrptresrunminsOk returns a tuple with the Grptresrunmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrptresrunmins

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetGrptresrunmins(v []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner)`

SetGrptresrunmins sets Grptresrunmins field to given value.

### HasGrptresrunmins

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasGrptresrunmins() bool`

HasGrptresrunmins returns a boolean if a field has been set.

### GetGrpwall

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetGrpwall() SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrpwall`

GetGrpwall returns the Grpwall field if non-nil, zero value otherwise.

### GetGrpwallOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetGrpwallOk() (*SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrpwall, bool)`

GetGrpwallOk returns a tuple with the Grpwall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpwall

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetGrpwall(v SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrpwall)`

SetGrpwall sets Grpwall field to given value.

### HasGrpwall

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasGrpwall() bool`

HasGrpwall returns a boolean if a field has been set.

### GetMaxjobs

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetMaxjobs() SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxjobs`

GetMaxjobs returns the Maxjobs field if non-nil, zero value otherwise.

### GetMaxjobsOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetMaxjobsOk() (*SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxjobs, bool)`

GetMaxjobsOk returns a tuple with the Maxjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxjobs

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetMaxjobs(v SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxjobs)`

SetMaxjobs sets Maxjobs field to given value.

### HasMaxjobs

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasMaxjobs() bool`

HasMaxjobs returns a boolean if a field has been set.

### GetMaxjobsaccrue

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetMaxjobsaccrue() SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxjobsaccrue`

GetMaxjobsaccrue returns the Maxjobsaccrue field if non-nil, zero value otherwise.

### GetMaxjobsaccrueOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetMaxjobsaccrueOk() (*SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxjobsaccrue, bool)`

GetMaxjobsaccrueOk returns a tuple with the Maxjobsaccrue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxjobsaccrue

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetMaxjobsaccrue(v SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxjobsaccrue)`

SetMaxjobsaccrue sets Maxjobsaccrue field to given value.

### HasMaxjobsaccrue

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasMaxjobsaccrue() bool`

HasMaxjobsaccrue returns a boolean if a field has been set.

### GetMaxsubmitjobs

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetMaxsubmitjobs() SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs`

GetMaxsubmitjobs returns the Maxsubmitjobs field if non-nil, zero value otherwise.

### GetMaxsubmitjobsOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetMaxsubmitjobsOk() (*SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs, bool)`

GetMaxsubmitjobsOk returns a tuple with the Maxsubmitjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxsubmitjobs

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetMaxsubmitjobs(v SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxsubmitjobs)`

SetMaxsubmitjobs sets Maxsubmitjobs field to given value.

### HasMaxsubmitjobs

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasMaxsubmitjobs() bool`

HasMaxsubmitjobs returns a boolean if a field has been set.

### GetMaxtresminsperjob

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetMaxtresminsperjob() []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner`

GetMaxtresminsperjob returns the Maxtresminsperjob field if non-nil, zero value otherwise.

### GetMaxtresminsperjobOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetMaxtresminsperjobOk() (*[]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner, bool)`

GetMaxtresminsperjobOk returns a tuple with the Maxtresminsperjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxtresminsperjob

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetMaxtresminsperjob(v []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner)`

SetMaxtresminsperjob sets Maxtresminsperjob field to given value.

### HasMaxtresminsperjob

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasMaxtresminsperjob() bool`

HasMaxtresminsperjob returns a boolean if a field has been set.

### GetMaxtresrunmins

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetMaxtresrunmins() []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner`

GetMaxtresrunmins returns the Maxtresrunmins field if non-nil, zero value otherwise.

### GetMaxtresrunminsOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetMaxtresrunminsOk() (*[]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner, bool)`

GetMaxtresrunminsOk returns a tuple with the Maxtresrunmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxtresrunmins

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetMaxtresrunmins(v []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner)`

SetMaxtresrunmins sets Maxtresrunmins field to given value.

### HasMaxtresrunmins

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasMaxtresrunmins() bool`

HasMaxtresrunmins returns a boolean if a field has been set.

### GetMaxtresperjob

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetMaxtresperjob() []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner`

GetMaxtresperjob returns the Maxtresperjob field if non-nil, zero value otherwise.

### GetMaxtresperjobOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetMaxtresperjobOk() (*[]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner, bool)`

GetMaxtresperjobOk returns a tuple with the Maxtresperjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxtresperjob

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetMaxtresperjob(v []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner)`

SetMaxtresperjob sets Maxtresperjob field to given value.

### HasMaxtresperjob

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasMaxtresperjob() bool`

HasMaxtresperjob returns a boolean if a field has been set.

### GetMaxtrespernode

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetMaxtrespernode() []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner`

GetMaxtrespernode returns the Maxtrespernode field if non-nil, zero value otherwise.

### GetMaxtrespernodeOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetMaxtrespernodeOk() (*[]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner, bool)`

GetMaxtrespernodeOk returns a tuple with the Maxtrespernode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxtrespernode

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetMaxtrespernode(v []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner)`

SetMaxtrespernode sets Maxtrespernode field to given value.

### HasMaxtrespernode

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasMaxtrespernode() bool`

HasMaxtrespernode returns a boolean if a field has been set.

### GetMaxwalldurationperjob

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetMaxwalldurationperjob() SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxwalldurationperjob`

GetMaxwalldurationperjob returns the Maxwalldurationperjob field if non-nil, zero value otherwise.

### GetMaxwalldurationperjobOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetMaxwalldurationperjobOk() (*SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxwalldurationperjob, bool)`

GetMaxwalldurationperjobOk returns a tuple with the Maxwalldurationperjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxwalldurationperjob

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetMaxwalldurationperjob(v SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMaxwalldurationperjob)`

SetMaxwalldurationperjob sets Maxwalldurationperjob field to given value.

### HasMaxwalldurationperjob

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasMaxwalldurationperjob() bool`

HasMaxwalldurationperjob returns a boolean if a field has been set.

### GetMinpriothresh

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetMinpriothresh() SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMinpriothresh`

GetMinpriothresh returns the Minpriothresh field if non-nil, zero value otherwise.

### GetMinpriothreshOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetMinpriothreshOk() (*SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMinpriothresh, bool)`

GetMinpriothreshOk returns a tuple with the Minpriothresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinpriothresh

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetMinpriothresh(v SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationMinpriothresh)`

SetMinpriothresh sets Minpriothresh field to given value.

### HasMinpriothresh

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasMinpriothresh() bool`

HasMinpriothresh returns a boolean if a field has been set.

### GetParent

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPriority

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetPriority() SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetPriorityOk() (*SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetPriority(v SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQoslevel

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetQoslevel() []string`

GetQoslevel returns the Qoslevel field if non-nil, zero value otherwise.

### GetQoslevelOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetQoslevelOk() (*[]string, bool)`

GetQoslevelOk returns a tuple with the Qoslevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoslevel

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetQoslevel(v []string)`

SetQoslevel sets Qoslevel field to given value.

### HasQoslevel

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasQoslevel() bool`

HasQoslevel returns a boolean if a field has been set.

### GetFairshare

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetFairshare() int32`

GetFairshare returns the Fairshare field if non-nil, zero value otherwise.

### GetFairshareOk

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) GetFairshareOk() (*int32, bool)`

GetFairshareOk returns a tuple with the Fairshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFairshare

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) SetFairshare(v int32)`

SetFairshare sets Fairshare field to given value.

### HasFairshare

`func (o *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) HasFairshare() bool`

HasFairshare returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


