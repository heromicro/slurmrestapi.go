/*
Slurm Rest API

API to access and control Slurm.

API version: Slurm-23.11.1&openapi/v0.0.39&openapi/slurmctld&openapi/slurmdbd&openapi/v0.0.38&openapi/dbv0.0.38&openapi/dbv0.0.39
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the V0039Step type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039Step{}

// V0039Step struct for V0039Step
type V0039Step struct {
	Time *V0039StepTime `json:"time,omitempty"`
	ExitCode *V0039JobExitCode `json:"exit_code,omitempty"`
	Nodes *V0039StepNodes `json:"nodes,omitempty"`
	Tasks *V0040StepTasks `json:"tasks,omitempty"`
	Pid *string `json:"pid,omitempty"`
	CPU *V0039StepCPU `json:"CPU,omitempty"`
	KillRequestUser *string `json:"kill_request_user,omitempty"`
	State *string `json:"state,omitempty"`
	Statistics *V0039StepStatistics `json:"statistics,omitempty"`
	Step *V0039StepStep `json:"step,omitempty"`
	Task *V0040StepTask `json:"task,omitempty"`
	Tres *V0039StepTres `json:"tres,omitempty"`
}

// NewV0039Step instantiates a new V0039Step object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039Step() *V0039Step {
	this := V0039Step{}
	return &this
}

// NewV0039StepWithDefaults instantiates a new V0039Step object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039StepWithDefaults() *V0039Step {
	this := V0039Step{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *V0039Step) GetTime() V0039StepTime {
	if o == nil || IsNil(o.Time) {
		var ret V0039StepTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Step) GetTimeOk() (*V0039StepTime, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *V0039Step) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given V0039StepTime and assigns it to the Time field.
func (o *V0039Step) SetTime(v V0039StepTime) {
	o.Time = &v
}

// GetExitCode returns the ExitCode field value if set, zero value otherwise.
func (o *V0039Step) GetExitCode() V0039JobExitCode {
	if o == nil || IsNil(o.ExitCode) {
		var ret V0039JobExitCode
		return ret
	}
	return *o.ExitCode
}

// GetExitCodeOk returns a tuple with the ExitCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Step) GetExitCodeOk() (*V0039JobExitCode, bool) {
	if o == nil || IsNil(o.ExitCode) {
		return nil, false
	}
	return o.ExitCode, true
}

// HasExitCode returns a boolean if a field has been set.
func (o *V0039Step) HasExitCode() bool {
	if o != nil && !IsNil(o.ExitCode) {
		return true
	}

	return false
}

// SetExitCode gets a reference to the given V0039JobExitCode and assigns it to the ExitCode field.
func (o *V0039Step) SetExitCode(v V0039JobExitCode) {
	o.ExitCode = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *V0039Step) GetNodes() V0039StepNodes {
	if o == nil || IsNil(o.Nodes) {
		var ret V0039StepNodes
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Step) GetNodesOk() (*V0039StepNodes, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *V0039Step) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given V0039StepNodes and assigns it to the Nodes field.
func (o *V0039Step) SetNodes(v V0039StepNodes) {
	o.Nodes = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *V0039Step) GetTasks() V0040StepTasks {
	if o == nil || IsNil(o.Tasks) {
		var ret V0040StepTasks
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Step) GetTasksOk() (*V0040StepTasks, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *V0039Step) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given V0040StepTasks and assigns it to the Tasks field.
func (o *V0039Step) SetTasks(v V0040StepTasks) {
	o.Tasks = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *V0039Step) GetPid() string {
	if o == nil || IsNil(o.Pid) {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Step) GetPidOk() (*string, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *V0039Step) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *V0039Step) SetPid(v string) {
	o.Pid = &v
}

// GetCPU returns the CPU field value if set, zero value otherwise.
func (o *V0039Step) GetCPU() V0039StepCPU {
	if o == nil || IsNil(o.CPU) {
		var ret V0039StepCPU
		return ret
	}
	return *o.CPU
}

// GetCPUOk returns a tuple with the CPU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Step) GetCPUOk() (*V0039StepCPU, bool) {
	if o == nil || IsNil(o.CPU) {
		return nil, false
	}
	return o.CPU, true
}

// HasCPU returns a boolean if a field has been set.
func (o *V0039Step) HasCPU() bool {
	if o != nil && !IsNil(o.CPU) {
		return true
	}

	return false
}

// SetCPU gets a reference to the given V0039StepCPU and assigns it to the CPU field.
func (o *V0039Step) SetCPU(v V0039StepCPU) {
	o.CPU = &v
}

// GetKillRequestUser returns the KillRequestUser field value if set, zero value otherwise.
func (o *V0039Step) GetKillRequestUser() string {
	if o == nil || IsNil(o.KillRequestUser) {
		var ret string
		return ret
	}
	return *o.KillRequestUser
}

// GetKillRequestUserOk returns a tuple with the KillRequestUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Step) GetKillRequestUserOk() (*string, bool) {
	if o == nil || IsNil(o.KillRequestUser) {
		return nil, false
	}
	return o.KillRequestUser, true
}

// HasKillRequestUser returns a boolean if a field has been set.
func (o *V0039Step) HasKillRequestUser() bool {
	if o != nil && !IsNil(o.KillRequestUser) {
		return true
	}

	return false
}

// SetKillRequestUser gets a reference to the given string and assigns it to the KillRequestUser field.
func (o *V0039Step) SetKillRequestUser(v string) {
	o.KillRequestUser = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *V0039Step) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Step) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *V0039Step) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *V0039Step) SetState(v string) {
	o.State = &v
}

// GetStatistics returns the Statistics field value if set, zero value otherwise.
func (o *V0039Step) GetStatistics() V0039StepStatistics {
	if o == nil || IsNil(o.Statistics) {
		var ret V0039StepStatistics
		return ret
	}
	return *o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Step) GetStatisticsOk() (*V0039StepStatistics, bool) {
	if o == nil || IsNil(o.Statistics) {
		return nil, false
	}
	return o.Statistics, true
}

// HasStatistics returns a boolean if a field has been set.
func (o *V0039Step) HasStatistics() bool {
	if o != nil && !IsNil(o.Statistics) {
		return true
	}

	return false
}

// SetStatistics gets a reference to the given V0039StepStatistics and assigns it to the Statistics field.
func (o *V0039Step) SetStatistics(v V0039StepStatistics) {
	o.Statistics = &v
}

// GetStep returns the Step field value if set, zero value otherwise.
func (o *V0039Step) GetStep() V0039StepStep {
	if o == nil || IsNil(o.Step) {
		var ret V0039StepStep
		return ret
	}
	return *o.Step
}

// GetStepOk returns a tuple with the Step field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Step) GetStepOk() (*V0039StepStep, bool) {
	if o == nil || IsNil(o.Step) {
		return nil, false
	}
	return o.Step, true
}

// HasStep returns a boolean if a field has been set.
func (o *V0039Step) HasStep() bool {
	if o != nil && !IsNil(o.Step) {
		return true
	}

	return false
}

// SetStep gets a reference to the given V0039StepStep and assigns it to the Step field.
func (o *V0039Step) SetStep(v V0039StepStep) {
	o.Step = &v
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *V0039Step) GetTask() V0040StepTask {
	if o == nil || IsNil(o.Task) {
		var ret V0040StepTask
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Step) GetTaskOk() (*V0040StepTask, bool) {
	if o == nil || IsNil(o.Task) {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *V0039Step) HasTask() bool {
	if o != nil && !IsNil(o.Task) {
		return true
	}

	return false
}

// SetTask gets a reference to the given V0040StepTask and assigns it to the Task field.
func (o *V0039Step) SetTask(v V0040StepTask) {
	o.Task = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *V0039Step) GetTres() V0039StepTres {
	if o == nil || IsNil(o.Tres) {
		var ret V0039StepTres
		return ret
	}
	return *o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Step) GetTresOk() (*V0039StepTres, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *V0039Step) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given V0039StepTres and assigns it to the Tres field.
func (o *V0039Step) SetTres(v V0039StepTres) {
	o.Tres = &v
}

func (o V0039Step) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039Step) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.ExitCode) {
		toSerialize["exit_code"] = o.ExitCode
	}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	if !IsNil(o.Tasks) {
		toSerialize["tasks"] = o.Tasks
	}
	if !IsNil(o.Pid) {
		toSerialize["pid"] = o.Pid
	}
	if !IsNil(o.CPU) {
		toSerialize["CPU"] = o.CPU
	}
	if !IsNil(o.KillRequestUser) {
		toSerialize["kill_request_user"] = o.KillRequestUser
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Statistics) {
		toSerialize["statistics"] = o.Statistics
	}
	if !IsNil(o.Step) {
		toSerialize["step"] = o.Step
	}
	if !IsNil(o.Task) {
		toSerialize["task"] = o.Task
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	return toSerialize, nil
}

type NullableV0039Step struct {
	value *V0039Step
	isSet bool
}

func (v NullableV0039Step) Get() *V0039Step {
	return v.value
}

func (v *NullableV0039Step) Set(val *V0039Step) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039Step) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039Step) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039Step(val *V0039Step) *NullableV0039Step {
	return &NullableV0039Step{value: val, isSet: true}
}

func (v NullableV0039Step) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039Step) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


