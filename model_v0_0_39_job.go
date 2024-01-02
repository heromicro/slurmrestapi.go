/*
Slurm Rest API

API to access and control Slurm.

API version: Slurm-23.11.1&openapi/v0.0.39&openapi/slurmctld&openapi/slurmdbd&openapi&openapi/dbv0.0.39
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the V0039Job type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039Job{}

// V0039Job struct for V0039Job
type V0039Job struct {
	Account *string `json:"account,omitempty"`
	Comment *V0040JobComment `json:"comment,omitempty"`
	AllocationNodes *int32 `json:"allocation_nodes,omitempty"`
	Array *V0039JobArray `json:"array,omitempty"`
	Association *V0039AssocShort `json:"association,omitempty"`
	Block *string `json:"block,omitempty"`
	Cluster *string `json:"cluster,omitempty"`
	Constraints *string `json:"constraints,omitempty"`
	Container *string `json:"container,omitempty"`
	DerivedExitCode *V0039JobExitCode `json:"derived_exit_code,omitempty"`
	Time *V0039JobTime `json:"time,omitempty"`
	ExitCode *V0039JobExitCode `json:"exit_code,omitempty"`
	Extra *string `json:"extra,omitempty"`
	FailedNode *string `json:"failed_node,omitempty"`
	Flags []string `json:"flags,omitempty"`
	Group *string `json:"group,omitempty"`
	Het *V0039JobHet `json:"het,omitempty"`
	JobId *int32 `json:"job_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Licenses *string `json:"licenses,omitempty"`
	Mcs *V0040JobMcs `json:"mcs,omitempty"`
	Nodes *string `json:"nodes,omitempty"`
	Partition *string `json:"partition,omitempty"`
	// Hold (true) or release (false) job
	Hold *bool `json:"hold,omitempty"`
	Priority *V0039Uint32NoVal `json:"priority,omitempty"`
	Qos *string `json:"qos,omitempty"`
	Required *V0039JobRequired `json:"required,omitempty"`
	KillRequestUser *string `json:"kill_request_user,omitempty"`
	Reservation *V0040JobReservation `json:"reservation,omitempty"`
	Script *string `json:"script,omitempty"`
	State *V0039JobState `json:"state,omitempty"`
	Steps []V0039Step `json:"steps,omitempty"`
	SubmitLine *string `json:"submit_line,omitempty"`
	Tres *V0039JobTres `json:"tres,omitempty"`
	UsedGres *string `json:"used_gres,omitempty"`
	User *string `json:"user,omitempty"`
	Wckey *V0039WckeyTag `json:"wckey,omitempty"`
	WorkingDirectory *string `json:"working_directory,omitempty"`
}

// NewV0039Job instantiates a new V0039Job object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039Job() *V0039Job {
	this := V0039Job{}
	return &this
}

// NewV0039JobWithDefaults instantiates a new V0039Job object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039JobWithDefaults() *V0039Job {
	this := V0039Job{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *V0039Job) GetAccount() string {
	if o == nil || IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *V0039Job) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *V0039Job) SetAccount(v string) {
	o.Account = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *V0039Job) GetComment() V0040JobComment {
	if o == nil || IsNil(o.Comment) {
		var ret V0040JobComment
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetCommentOk() (*V0040JobComment, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *V0039Job) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given V0040JobComment and assigns it to the Comment field.
func (o *V0039Job) SetComment(v V0040JobComment) {
	o.Comment = &v
}

// GetAllocationNodes returns the AllocationNodes field value if set, zero value otherwise.
func (o *V0039Job) GetAllocationNodes() int32 {
	if o == nil || IsNil(o.AllocationNodes) {
		var ret int32
		return ret
	}
	return *o.AllocationNodes
}

// GetAllocationNodesOk returns a tuple with the AllocationNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetAllocationNodesOk() (*int32, bool) {
	if o == nil || IsNil(o.AllocationNodes) {
		return nil, false
	}
	return o.AllocationNodes, true
}

// HasAllocationNodes returns a boolean if a field has been set.
func (o *V0039Job) HasAllocationNodes() bool {
	if o != nil && !IsNil(o.AllocationNodes) {
		return true
	}

	return false
}

// SetAllocationNodes gets a reference to the given int32 and assigns it to the AllocationNodes field.
func (o *V0039Job) SetAllocationNodes(v int32) {
	o.AllocationNodes = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *V0039Job) GetArray() V0039JobArray {
	if o == nil || IsNil(o.Array) {
		var ret V0039JobArray
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetArrayOk() (*V0039JobArray, bool) {
	if o == nil || IsNil(o.Array) {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *V0039Job) HasArray() bool {
	if o != nil && !IsNil(o.Array) {
		return true
	}

	return false
}

// SetArray gets a reference to the given V0039JobArray and assigns it to the Array field.
func (o *V0039Job) SetArray(v V0039JobArray) {
	o.Array = &v
}

// GetAssociation returns the Association field value if set, zero value otherwise.
func (o *V0039Job) GetAssociation() V0039AssocShort {
	if o == nil || IsNil(o.Association) {
		var ret V0039AssocShort
		return ret
	}
	return *o.Association
}

// GetAssociationOk returns a tuple with the Association field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetAssociationOk() (*V0039AssocShort, bool) {
	if o == nil || IsNil(o.Association) {
		return nil, false
	}
	return o.Association, true
}

// HasAssociation returns a boolean if a field has been set.
func (o *V0039Job) HasAssociation() bool {
	if o != nil && !IsNil(o.Association) {
		return true
	}

	return false
}

// SetAssociation gets a reference to the given V0039AssocShort and assigns it to the Association field.
func (o *V0039Job) SetAssociation(v V0039AssocShort) {
	o.Association = &v
}

// GetBlock returns the Block field value if set, zero value otherwise.
func (o *V0039Job) GetBlock() string {
	if o == nil || IsNil(o.Block) {
		var ret string
		return ret
	}
	return *o.Block
}

// GetBlockOk returns a tuple with the Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetBlockOk() (*string, bool) {
	if o == nil || IsNil(o.Block) {
		return nil, false
	}
	return o.Block, true
}

// HasBlock returns a boolean if a field has been set.
func (o *V0039Job) HasBlock() bool {
	if o != nil && !IsNil(o.Block) {
		return true
	}

	return false
}

// SetBlock gets a reference to the given string and assigns it to the Block field.
func (o *V0039Job) SetBlock(v string) {
	o.Block = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *V0039Job) GetCluster() string {
	if o == nil || IsNil(o.Cluster) {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetClusterOk() (*string, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *V0039Job) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *V0039Job) SetCluster(v string) {
	o.Cluster = &v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *V0039Job) GetConstraints() string {
	if o == nil || IsNil(o.Constraints) {
		var ret string
		return ret
	}
	return *o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetConstraintsOk() (*string, bool) {
	if o == nil || IsNil(o.Constraints) {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *V0039Job) HasConstraints() bool {
	if o != nil && !IsNil(o.Constraints) {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given string and assigns it to the Constraints field.
func (o *V0039Job) SetConstraints(v string) {
	o.Constraints = &v
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *V0039Job) GetContainer() string {
	if o == nil || IsNil(o.Container) {
		var ret string
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetContainerOk() (*string, bool) {
	if o == nil || IsNil(o.Container) {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *V0039Job) HasContainer() bool {
	if o != nil && !IsNil(o.Container) {
		return true
	}

	return false
}

// SetContainer gets a reference to the given string and assigns it to the Container field.
func (o *V0039Job) SetContainer(v string) {
	o.Container = &v
}

// GetDerivedExitCode returns the DerivedExitCode field value if set, zero value otherwise.
func (o *V0039Job) GetDerivedExitCode() V0039JobExitCode {
	if o == nil || IsNil(o.DerivedExitCode) {
		var ret V0039JobExitCode
		return ret
	}
	return *o.DerivedExitCode
}

// GetDerivedExitCodeOk returns a tuple with the DerivedExitCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetDerivedExitCodeOk() (*V0039JobExitCode, bool) {
	if o == nil || IsNil(o.DerivedExitCode) {
		return nil, false
	}
	return o.DerivedExitCode, true
}

// HasDerivedExitCode returns a boolean if a field has been set.
func (o *V0039Job) HasDerivedExitCode() bool {
	if o != nil && !IsNil(o.DerivedExitCode) {
		return true
	}

	return false
}

// SetDerivedExitCode gets a reference to the given V0039JobExitCode and assigns it to the DerivedExitCode field.
func (o *V0039Job) SetDerivedExitCode(v V0039JobExitCode) {
	o.DerivedExitCode = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *V0039Job) GetTime() V0039JobTime {
	if o == nil || IsNil(o.Time) {
		var ret V0039JobTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetTimeOk() (*V0039JobTime, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *V0039Job) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given V0039JobTime and assigns it to the Time field.
func (o *V0039Job) SetTime(v V0039JobTime) {
	o.Time = &v
}

// GetExitCode returns the ExitCode field value if set, zero value otherwise.
func (o *V0039Job) GetExitCode() V0039JobExitCode {
	if o == nil || IsNil(o.ExitCode) {
		var ret V0039JobExitCode
		return ret
	}
	return *o.ExitCode
}

// GetExitCodeOk returns a tuple with the ExitCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetExitCodeOk() (*V0039JobExitCode, bool) {
	if o == nil || IsNil(o.ExitCode) {
		return nil, false
	}
	return o.ExitCode, true
}

// HasExitCode returns a boolean if a field has been set.
func (o *V0039Job) HasExitCode() bool {
	if o != nil && !IsNil(o.ExitCode) {
		return true
	}

	return false
}

// SetExitCode gets a reference to the given V0039JobExitCode and assigns it to the ExitCode field.
func (o *V0039Job) SetExitCode(v V0039JobExitCode) {
	o.ExitCode = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *V0039Job) GetExtra() string {
	if o == nil || IsNil(o.Extra) {
		var ret string
		return ret
	}
	return *o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetExtraOk() (*string, bool) {
	if o == nil || IsNil(o.Extra) {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *V0039Job) HasExtra() bool {
	if o != nil && !IsNil(o.Extra) {
		return true
	}

	return false
}

// SetExtra gets a reference to the given string and assigns it to the Extra field.
func (o *V0039Job) SetExtra(v string) {
	o.Extra = &v
}

// GetFailedNode returns the FailedNode field value if set, zero value otherwise.
func (o *V0039Job) GetFailedNode() string {
	if o == nil || IsNil(o.FailedNode) {
		var ret string
		return ret
	}
	return *o.FailedNode
}

// GetFailedNodeOk returns a tuple with the FailedNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetFailedNodeOk() (*string, bool) {
	if o == nil || IsNil(o.FailedNode) {
		return nil, false
	}
	return o.FailedNode, true
}

// HasFailedNode returns a boolean if a field has been set.
func (o *V0039Job) HasFailedNode() bool {
	if o != nil && !IsNil(o.FailedNode) {
		return true
	}

	return false
}

// SetFailedNode gets a reference to the given string and assigns it to the FailedNode field.
func (o *V0039Job) SetFailedNode(v string) {
	o.FailedNode = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0039Job) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0039Job) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0039Job) SetFlags(v []string) {
	o.Flags = v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *V0039Job) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *V0039Job) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *V0039Job) SetGroup(v string) {
	o.Group = &v
}

// GetHet returns the Het field value if set, zero value otherwise.
func (o *V0039Job) GetHet() V0039JobHet {
	if o == nil || IsNil(o.Het) {
		var ret V0039JobHet
		return ret
	}
	return *o.Het
}

// GetHetOk returns a tuple with the Het field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetHetOk() (*V0039JobHet, bool) {
	if o == nil || IsNil(o.Het) {
		return nil, false
	}
	return o.Het, true
}

// HasHet returns a boolean if a field has been set.
func (o *V0039Job) HasHet() bool {
	if o != nil && !IsNil(o.Het) {
		return true
	}

	return false
}

// SetHet gets a reference to the given V0039JobHet and assigns it to the Het field.
func (o *V0039Job) SetHet(v V0039JobHet) {
	o.Het = &v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *V0039Job) GetJobId() int32 {
	if o == nil || IsNil(o.JobId) {
		var ret int32
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetJobIdOk() (*int32, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *V0039Job) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given int32 and assigns it to the JobId field.
func (o *V0039Job) SetJobId(v int32) {
	o.JobId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0039Job) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0039Job) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0039Job) SetName(v string) {
	o.Name = &v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *V0039Job) GetLicenses() string {
	if o == nil || IsNil(o.Licenses) {
		var ret string
		return ret
	}
	return *o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetLicensesOk() (*string, bool) {
	if o == nil || IsNil(o.Licenses) {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *V0039Job) HasLicenses() bool {
	if o != nil && !IsNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given string and assigns it to the Licenses field.
func (o *V0039Job) SetLicenses(v string) {
	o.Licenses = &v
}

// GetMcs returns the Mcs field value if set, zero value otherwise.
func (o *V0039Job) GetMcs() V0040JobMcs {
	if o == nil || IsNil(o.Mcs) {
		var ret V0040JobMcs
		return ret
	}
	return *o.Mcs
}

// GetMcsOk returns a tuple with the Mcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetMcsOk() (*V0040JobMcs, bool) {
	if o == nil || IsNil(o.Mcs) {
		return nil, false
	}
	return o.Mcs, true
}

// HasMcs returns a boolean if a field has been set.
func (o *V0039Job) HasMcs() bool {
	if o != nil && !IsNil(o.Mcs) {
		return true
	}

	return false
}

// SetMcs gets a reference to the given V0040JobMcs and assigns it to the Mcs field.
func (o *V0039Job) SetMcs(v V0040JobMcs) {
	o.Mcs = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *V0039Job) GetNodes() string {
	if o == nil || IsNil(o.Nodes) {
		var ret string
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetNodesOk() (*string, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *V0039Job) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given string and assigns it to the Nodes field.
func (o *V0039Job) SetNodes(v string) {
	o.Nodes = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *V0039Job) GetPartition() string {
	if o == nil || IsNil(o.Partition) {
		var ret string
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetPartitionOk() (*string, bool) {
	if o == nil || IsNil(o.Partition) {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *V0039Job) HasPartition() bool {
	if o != nil && !IsNil(o.Partition) {
		return true
	}

	return false
}

// SetPartition gets a reference to the given string and assigns it to the Partition field.
func (o *V0039Job) SetPartition(v string) {
	o.Partition = &v
}

// GetHold returns the Hold field value if set, zero value otherwise.
func (o *V0039Job) GetHold() bool {
	if o == nil || IsNil(o.Hold) {
		var ret bool
		return ret
	}
	return *o.Hold
}

// GetHoldOk returns a tuple with the Hold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetHoldOk() (*bool, bool) {
	if o == nil || IsNil(o.Hold) {
		return nil, false
	}
	return o.Hold, true
}

// HasHold returns a boolean if a field has been set.
func (o *V0039Job) HasHold() bool {
	if o != nil && !IsNil(o.Hold) {
		return true
	}

	return false
}

// SetHold gets a reference to the given bool and assigns it to the Hold field.
func (o *V0039Job) SetHold(v bool) {
	o.Hold = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *V0039Job) GetPriority() V0039Uint32NoVal {
	if o == nil || IsNil(o.Priority) {
		var ret V0039Uint32NoVal
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetPriorityOk() (*V0039Uint32NoVal, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *V0039Job) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given V0039Uint32NoVal and assigns it to the Priority field.
func (o *V0039Job) SetPriority(v V0039Uint32NoVal) {
	o.Priority = &v
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *V0039Job) GetQos() string {
	if o == nil || IsNil(o.Qos) {
		var ret string
		return ret
	}
	return *o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetQosOk() (*string, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *V0039Job) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given string and assigns it to the Qos field.
func (o *V0039Job) SetQos(v string) {
	o.Qos = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *V0039Job) GetRequired() V0039JobRequired {
	if o == nil || IsNil(o.Required) {
		var ret V0039JobRequired
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetRequiredOk() (*V0039JobRequired, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *V0039Job) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given V0039JobRequired and assigns it to the Required field.
func (o *V0039Job) SetRequired(v V0039JobRequired) {
	o.Required = &v
}

// GetKillRequestUser returns the KillRequestUser field value if set, zero value otherwise.
func (o *V0039Job) GetKillRequestUser() string {
	if o == nil || IsNil(o.KillRequestUser) {
		var ret string
		return ret
	}
	return *o.KillRequestUser
}

// GetKillRequestUserOk returns a tuple with the KillRequestUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetKillRequestUserOk() (*string, bool) {
	if o == nil || IsNil(o.KillRequestUser) {
		return nil, false
	}
	return o.KillRequestUser, true
}

// HasKillRequestUser returns a boolean if a field has been set.
func (o *V0039Job) HasKillRequestUser() bool {
	if o != nil && !IsNil(o.KillRequestUser) {
		return true
	}

	return false
}

// SetKillRequestUser gets a reference to the given string and assigns it to the KillRequestUser field.
func (o *V0039Job) SetKillRequestUser(v string) {
	o.KillRequestUser = &v
}

// GetReservation returns the Reservation field value if set, zero value otherwise.
func (o *V0039Job) GetReservation() V0040JobReservation {
	if o == nil || IsNil(o.Reservation) {
		var ret V0040JobReservation
		return ret
	}
	return *o.Reservation
}

// GetReservationOk returns a tuple with the Reservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetReservationOk() (*V0040JobReservation, bool) {
	if o == nil || IsNil(o.Reservation) {
		return nil, false
	}
	return o.Reservation, true
}

// HasReservation returns a boolean if a field has been set.
func (o *V0039Job) HasReservation() bool {
	if o != nil && !IsNil(o.Reservation) {
		return true
	}

	return false
}

// SetReservation gets a reference to the given V0040JobReservation and assigns it to the Reservation field.
func (o *V0039Job) SetReservation(v V0040JobReservation) {
	o.Reservation = &v
}

// GetScript returns the Script field value if set, zero value otherwise.
func (o *V0039Job) GetScript() string {
	if o == nil || IsNil(o.Script) {
		var ret string
		return ret
	}
	return *o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetScriptOk() (*string, bool) {
	if o == nil || IsNil(o.Script) {
		return nil, false
	}
	return o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *V0039Job) HasScript() bool {
	if o != nil && !IsNil(o.Script) {
		return true
	}

	return false
}

// SetScript gets a reference to the given string and assigns it to the Script field.
func (o *V0039Job) SetScript(v string) {
	o.Script = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *V0039Job) GetState() V0039JobState {
	if o == nil || IsNil(o.State) {
		var ret V0039JobState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetStateOk() (*V0039JobState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *V0039Job) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given V0039JobState and assigns it to the State field.
func (o *V0039Job) SetState(v V0039JobState) {
	o.State = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *V0039Job) GetSteps() []V0039Step {
	if o == nil || IsNil(o.Steps) {
		var ret []V0039Step
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetStepsOk() ([]V0039Step, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *V0039Job) HasSteps() bool {
	if o != nil && !IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []V0039Step and assigns it to the Steps field.
func (o *V0039Job) SetSteps(v []V0039Step) {
	o.Steps = v
}

// GetSubmitLine returns the SubmitLine field value if set, zero value otherwise.
func (o *V0039Job) GetSubmitLine() string {
	if o == nil || IsNil(o.SubmitLine) {
		var ret string
		return ret
	}
	return *o.SubmitLine
}

// GetSubmitLineOk returns a tuple with the SubmitLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetSubmitLineOk() (*string, bool) {
	if o == nil || IsNil(o.SubmitLine) {
		return nil, false
	}
	return o.SubmitLine, true
}

// HasSubmitLine returns a boolean if a field has been set.
func (o *V0039Job) HasSubmitLine() bool {
	if o != nil && !IsNil(o.SubmitLine) {
		return true
	}

	return false
}

// SetSubmitLine gets a reference to the given string and assigns it to the SubmitLine field.
func (o *V0039Job) SetSubmitLine(v string) {
	o.SubmitLine = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *V0039Job) GetTres() V0039JobTres {
	if o == nil || IsNil(o.Tres) {
		var ret V0039JobTres
		return ret
	}
	return *o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetTresOk() (*V0039JobTres, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *V0039Job) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given V0039JobTres and assigns it to the Tres field.
func (o *V0039Job) SetTres(v V0039JobTres) {
	o.Tres = &v
}

// GetUsedGres returns the UsedGres field value if set, zero value otherwise.
func (o *V0039Job) GetUsedGres() string {
	if o == nil || IsNil(o.UsedGres) {
		var ret string
		return ret
	}
	return *o.UsedGres
}

// GetUsedGresOk returns a tuple with the UsedGres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetUsedGresOk() (*string, bool) {
	if o == nil || IsNil(o.UsedGres) {
		return nil, false
	}
	return o.UsedGres, true
}

// HasUsedGres returns a boolean if a field has been set.
func (o *V0039Job) HasUsedGres() bool {
	if o != nil && !IsNil(o.UsedGres) {
		return true
	}

	return false
}

// SetUsedGres gets a reference to the given string and assigns it to the UsedGres field.
func (o *V0039Job) SetUsedGres(v string) {
	o.UsedGres = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *V0039Job) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *V0039Job) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *V0039Job) SetUser(v string) {
	o.User = &v
}

// GetWckey returns the Wckey field value if set, zero value otherwise.
func (o *V0039Job) GetWckey() V0039WckeyTag {
	if o == nil || IsNil(o.Wckey) {
		var ret V0039WckeyTag
		return ret
	}
	return *o.Wckey
}

// GetWckeyOk returns a tuple with the Wckey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetWckeyOk() (*V0039WckeyTag, bool) {
	if o == nil || IsNil(o.Wckey) {
		return nil, false
	}
	return o.Wckey, true
}

// HasWckey returns a boolean if a field has been set.
func (o *V0039Job) HasWckey() bool {
	if o != nil && !IsNil(o.Wckey) {
		return true
	}

	return false
}

// SetWckey gets a reference to the given V0039WckeyTag and assigns it to the Wckey field.
func (o *V0039Job) SetWckey(v V0039WckeyTag) {
	o.Wckey = &v
}

// GetWorkingDirectory returns the WorkingDirectory field value if set, zero value otherwise.
func (o *V0039Job) GetWorkingDirectory() string {
	if o == nil || IsNil(o.WorkingDirectory) {
		var ret string
		return ret
	}
	return *o.WorkingDirectory
}

// GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Job) GetWorkingDirectoryOk() (*string, bool) {
	if o == nil || IsNil(o.WorkingDirectory) {
		return nil, false
	}
	return o.WorkingDirectory, true
}

// HasWorkingDirectory returns a boolean if a field has been set.
func (o *V0039Job) HasWorkingDirectory() bool {
	if o != nil && !IsNil(o.WorkingDirectory) {
		return true
	}

	return false
}

// SetWorkingDirectory gets a reference to the given string and assigns it to the WorkingDirectory field.
func (o *V0039Job) SetWorkingDirectory(v string) {
	o.WorkingDirectory = &v
}

func (o V0039Job) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039Job) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.AllocationNodes) {
		toSerialize["allocation_nodes"] = o.AllocationNodes
	}
	if !IsNil(o.Array) {
		toSerialize["array"] = o.Array
	}
	if !IsNil(o.Association) {
		toSerialize["association"] = o.Association
	}
	if !IsNil(o.Block) {
		toSerialize["block"] = o.Block
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Constraints) {
		toSerialize["constraints"] = o.Constraints
	}
	if !IsNil(o.Container) {
		toSerialize["container"] = o.Container
	}
	if !IsNil(o.DerivedExitCode) {
		toSerialize["derived_exit_code"] = o.DerivedExitCode
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.ExitCode) {
		toSerialize["exit_code"] = o.ExitCode
	}
	if !IsNil(o.Extra) {
		toSerialize["extra"] = o.Extra
	}
	if !IsNil(o.FailedNode) {
		toSerialize["failed_node"] = o.FailedNode
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Het) {
		toSerialize["het"] = o.Het
	}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Licenses) {
		toSerialize["licenses"] = o.Licenses
	}
	if !IsNil(o.Mcs) {
		toSerialize["mcs"] = o.Mcs
	}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	if !IsNil(o.Partition) {
		toSerialize["partition"] = o.Partition
	}
	if !IsNil(o.Hold) {
		toSerialize["hold"] = o.Hold
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Qos) {
		toSerialize["qos"] = o.Qos
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.KillRequestUser) {
		toSerialize["kill_request_user"] = o.KillRequestUser
	}
	if !IsNil(o.Reservation) {
		toSerialize["reservation"] = o.Reservation
	}
	if !IsNil(o.Script) {
		toSerialize["script"] = o.Script
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}
	if !IsNil(o.SubmitLine) {
		toSerialize["submit_line"] = o.SubmitLine
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	if !IsNil(o.UsedGres) {
		toSerialize["used_gres"] = o.UsedGres
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Wckey) {
		toSerialize["wckey"] = o.Wckey
	}
	if !IsNil(o.WorkingDirectory) {
		toSerialize["working_directory"] = o.WorkingDirectory
	}
	return toSerialize, nil
}

type NullableV0039Job struct {
	value *V0039Job
	isSet bool
}

func (v NullableV0039Job) Get() *V0039Job {
	return v.value
}

func (v *NullableV0039Job) Set(val *V0039Job) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039Job) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039Job) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039Job(val *V0039Job) *NullableV0039Job {
	return &NullableV0039Job{value: val, isSet: true}
}

func (v NullableV0039Job) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039Job) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


