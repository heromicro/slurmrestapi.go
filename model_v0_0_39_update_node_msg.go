/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.39
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the V0039UpdateNodeMsg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039UpdateNodeMsg{}

// V0039UpdateNodeMsg struct for V0039UpdateNodeMsg
type V0039UpdateNodeMsg struct {
	Comment *string `json:"comment,omitempty"`
	CpuBind *int32 `json:"cpu_bind,omitempty"`
	Extra *string `json:"extra,omitempty"`
	Features []string `json:"features,omitempty"`
	FeaturesAct []string `json:"features_act,omitempty"`
	Gres *string `json:"gres,omitempty"`
	Address []string `json:"address,omitempty"`
	Hostname []string `json:"hostname,omitempty"`
	Name []string `json:"name,omitempty"`
	State []string `json:"state,omitempty"`
	Reason *string `json:"reason,omitempty"`
	ReasonUid *string `json:"reason_uid,omitempty"`
	ResumeAfter *V0039Uint32NoVal `json:"resume_after,omitempty"`
	Weight *V0039Uint32NoVal `json:"weight,omitempty"`
}

// NewV0039UpdateNodeMsg instantiates a new V0039UpdateNodeMsg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039UpdateNodeMsg() *V0039UpdateNodeMsg {
	this := V0039UpdateNodeMsg{}
	return &this
}

// NewV0039UpdateNodeMsgWithDefaults instantiates a new V0039UpdateNodeMsg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039UpdateNodeMsgWithDefaults() *V0039UpdateNodeMsg {
	this := V0039UpdateNodeMsg{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *V0039UpdateNodeMsg) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039UpdateNodeMsg) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *V0039UpdateNodeMsg) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *V0039UpdateNodeMsg) SetComment(v string) {
	o.Comment = &v
}

// GetCpuBind returns the CpuBind field value if set, zero value otherwise.
func (o *V0039UpdateNodeMsg) GetCpuBind() int32 {
	if o == nil || IsNil(o.CpuBind) {
		var ret int32
		return ret
	}
	return *o.CpuBind
}

// GetCpuBindOk returns a tuple with the CpuBind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039UpdateNodeMsg) GetCpuBindOk() (*int32, bool) {
	if o == nil || IsNil(o.CpuBind) {
		return nil, false
	}
	return o.CpuBind, true
}

// HasCpuBind returns a boolean if a field has been set.
func (o *V0039UpdateNodeMsg) HasCpuBind() bool {
	if o != nil && !IsNil(o.CpuBind) {
		return true
	}

	return false
}

// SetCpuBind gets a reference to the given int32 and assigns it to the CpuBind field.
func (o *V0039UpdateNodeMsg) SetCpuBind(v int32) {
	o.CpuBind = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *V0039UpdateNodeMsg) GetExtra() string {
	if o == nil || IsNil(o.Extra) {
		var ret string
		return ret
	}
	return *o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039UpdateNodeMsg) GetExtraOk() (*string, bool) {
	if o == nil || IsNil(o.Extra) {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *V0039UpdateNodeMsg) HasExtra() bool {
	if o != nil && !IsNil(o.Extra) {
		return true
	}

	return false
}

// SetExtra gets a reference to the given string and assigns it to the Extra field.
func (o *V0039UpdateNodeMsg) SetExtra(v string) {
	o.Extra = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *V0039UpdateNodeMsg) GetFeatures() []string {
	if o == nil || IsNil(o.Features) {
		var ret []string
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039UpdateNodeMsg) GetFeaturesOk() ([]string, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *V0039UpdateNodeMsg) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *V0039UpdateNodeMsg) SetFeatures(v []string) {
	o.Features = v
}

// GetFeaturesAct returns the FeaturesAct field value if set, zero value otherwise.
func (o *V0039UpdateNodeMsg) GetFeaturesAct() []string {
	if o == nil || IsNil(o.FeaturesAct) {
		var ret []string
		return ret
	}
	return o.FeaturesAct
}

// GetFeaturesActOk returns a tuple with the FeaturesAct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039UpdateNodeMsg) GetFeaturesActOk() ([]string, bool) {
	if o == nil || IsNil(o.FeaturesAct) {
		return nil, false
	}
	return o.FeaturesAct, true
}

// HasFeaturesAct returns a boolean if a field has been set.
func (o *V0039UpdateNodeMsg) HasFeaturesAct() bool {
	if o != nil && !IsNil(o.FeaturesAct) {
		return true
	}

	return false
}

// SetFeaturesAct gets a reference to the given []string and assigns it to the FeaturesAct field.
func (o *V0039UpdateNodeMsg) SetFeaturesAct(v []string) {
	o.FeaturesAct = v
}

// GetGres returns the Gres field value if set, zero value otherwise.
func (o *V0039UpdateNodeMsg) GetGres() string {
	if o == nil || IsNil(o.Gres) {
		var ret string
		return ret
	}
	return *o.Gres
}

// GetGresOk returns a tuple with the Gres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039UpdateNodeMsg) GetGresOk() (*string, bool) {
	if o == nil || IsNil(o.Gres) {
		return nil, false
	}
	return o.Gres, true
}

// HasGres returns a boolean if a field has been set.
func (o *V0039UpdateNodeMsg) HasGres() bool {
	if o != nil && !IsNil(o.Gres) {
		return true
	}

	return false
}

// SetGres gets a reference to the given string and assigns it to the Gres field.
func (o *V0039UpdateNodeMsg) SetGres(v string) {
	o.Gres = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *V0039UpdateNodeMsg) GetAddress() []string {
	if o == nil || IsNil(o.Address) {
		var ret []string
		return ret
	}
	return o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039UpdateNodeMsg) GetAddressOk() ([]string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *V0039UpdateNodeMsg) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given []string and assigns it to the Address field.
func (o *V0039UpdateNodeMsg) SetAddress(v []string) {
	o.Address = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *V0039UpdateNodeMsg) GetHostname() []string {
	if o == nil || IsNil(o.Hostname) {
		var ret []string
		return ret
	}
	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039UpdateNodeMsg) GetHostnameOk() ([]string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *V0039UpdateNodeMsg) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given []string and assigns it to the Hostname field.
func (o *V0039UpdateNodeMsg) SetHostname(v []string) {
	o.Hostname = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0039UpdateNodeMsg) GetName() []string {
	if o == nil || IsNil(o.Name) {
		var ret []string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039UpdateNodeMsg) GetNameOk() ([]string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0039UpdateNodeMsg) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given []string and assigns it to the Name field.
func (o *V0039UpdateNodeMsg) SetName(v []string) {
	o.Name = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *V0039UpdateNodeMsg) GetState() []string {
	if o == nil || IsNil(o.State) {
		var ret []string
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039UpdateNodeMsg) GetStateOk() ([]string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *V0039UpdateNodeMsg) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given []string and assigns it to the State field.
func (o *V0039UpdateNodeMsg) SetState(v []string) {
	o.State = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *V0039UpdateNodeMsg) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039UpdateNodeMsg) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *V0039UpdateNodeMsg) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *V0039UpdateNodeMsg) SetReason(v string) {
	o.Reason = &v
}

// GetReasonUid returns the ReasonUid field value if set, zero value otherwise.
func (o *V0039UpdateNodeMsg) GetReasonUid() string {
	if o == nil || IsNil(o.ReasonUid) {
		var ret string
		return ret
	}
	return *o.ReasonUid
}

// GetReasonUidOk returns a tuple with the ReasonUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039UpdateNodeMsg) GetReasonUidOk() (*string, bool) {
	if o == nil || IsNil(o.ReasonUid) {
		return nil, false
	}
	return o.ReasonUid, true
}

// HasReasonUid returns a boolean if a field has been set.
func (o *V0039UpdateNodeMsg) HasReasonUid() bool {
	if o != nil && !IsNil(o.ReasonUid) {
		return true
	}

	return false
}

// SetReasonUid gets a reference to the given string and assigns it to the ReasonUid field.
func (o *V0039UpdateNodeMsg) SetReasonUid(v string) {
	o.ReasonUid = &v
}

// GetResumeAfter returns the ResumeAfter field value if set, zero value otherwise.
func (o *V0039UpdateNodeMsg) GetResumeAfter() V0039Uint32NoVal {
	if o == nil || IsNil(o.ResumeAfter) {
		var ret V0039Uint32NoVal
		return ret
	}
	return *o.ResumeAfter
}

// GetResumeAfterOk returns a tuple with the ResumeAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039UpdateNodeMsg) GetResumeAfterOk() (*V0039Uint32NoVal, bool) {
	if o == nil || IsNil(o.ResumeAfter) {
		return nil, false
	}
	return o.ResumeAfter, true
}

// HasResumeAfter returns a boolean if a field has been set.
func (o *V0039UpdateNodeMsg) HasResumeAfter() bool {
	if o != nil && !IsNil(o.ResumeAfter) {
		return true
	}

	return false
}

// SetResumeAfter gets a reference to the given V0039Uint32NoVal and assigns it to the ResumeAfter field.
func (o *V0039UpdateNodeMsg) SetResumeAfter(v V0039Uint32NoVal) {
	o.ResumeAfter = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *V0039UpdateNodeMsg) GetWeight() V0039Uint32NoVal {
	if o == nil || IsNil(o.Weight) {
		var ret V0039Uint32NoVal
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039UpdateNodeMsg) GetWeightOk() (*V0039Uint32NoVal, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *V0039UpdateNodeMsg) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given V0039Uint32NoVal and assigns it to the Weight field.
func (o *V0039UpdateNodeMsg) SetWeight(v V0039Uint32NoVal) {
	o.Weight = &v
}

func (o V0039UpdateNodeMsg) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039UpdateNodeMsg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CpuBind) {
		toSerialize["cpu_bind"] = o.CpuBind
	}
	if !IsNil(o.Extra) {
		toSerialize["extra"] = o.Extra
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.FeaturesAct) {
		toSerialize["features_act"] = o.FeaturesAct
	}
	if !IsNil(o.Gres) {
		toSerialize["gres"] = o.Gres
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.ReasonUid) {
		toSerialize["reason_uid"] = o.ReasonUid
	}
	if !IsNil(o.ResumeAfter) {
		toSerialize["resume_after"] = o.ResumeAfter
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullableV0039UpdateNodeMsg struct {
	value *V0039UpdateNodeMsg
	isSet bool
}

func (v NullableV0039UpdateNodeMsg) Get() *V0039UpdateNodeMsg {
	return v.value
}

func (v *NullableV0039UpdateNodeMsg) Set(val *V0039UpdateNodeMsg) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039UpdateNodeMsg) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039UpdateNodeMsg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039UpdateNodeMsg(val *V0039UpdateNodeMsg) *NullableV0039UpdateNodeMsg {
	return &NullableV0039UpdateNodeMsg{value: val, isSet: true}
}

func (v NullableV0039UpdateNodeMsg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039UpdateNodeMsg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


