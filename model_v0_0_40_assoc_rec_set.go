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

// checks if the V0040AssocRecSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040AssocRecSet{}

// V0040AssocRecSet struct for V0040AssocRecSet
type V0040AssocRecSet struct {
	// Comment for the association
	Comment *string `json:"comment,omitempty"`
	// Which QOS id is this association default
	Defaultqos *string `json:"defaultqos,omitempty"`
	Grpjobs *V0040Uint32NoVal `json:"grpjobs,omitempty"`
	Grpjobsaccrue *V0040Uint32NoVal `json:"grpjobsaccrue,omitempty"`
	Grpsubmitjobs *V0040Uint32NoVal `json:"grpsubmitjobs,omitempty"`
	Grptres []V0040Tres `json:"grptres,omitempty"`
	Grptresmins []V0040Tres `json:"grptresmins,omitempty"`
	Grptresrunmins []V0040Tres `json:"grptresrunmins,omitempty"`
	Grpwall *V0040Uint32NoVal `json:"grpwall,omitempty"`
	Maxjobs *V0040Uint32NoVal `json:"maxjobs,omitempty"`
	Maxjobsaccrue *V0040Uint32NoVal `json:"maxjobsaccrue,omitempty"`
	Maxsubmitjobs *V0040Uint32NoVal `json:"maxsubmitjobs,omitempty"`
	Maxtresminsperjob []V0040Tres `json:"maxtresminsperjob,omitempty"`
	Maxtresrunmins []V0040Tres `json:"maxtresrunmins,omitempty"`
	Maxtresperjob []V0040Tres `json:"maxtresperjob,omitempty"`
	Maxtrespernode []V0040Tres `json:"maxtrespernode,omitempty"`
	Maxwalldurationperjob *V0040Uint32NoVal `json:"maxwalldurationperjob,omitempty"`
	Minpriothresh *V0040Uint32NoVal `json:"minpriothresh,omitempty"`
	// Name of parent account
	Parent *string `json:"parent,omitempty"`
	Priority *V0040Uint32NoVal `json:"priority,omitempty"`
	// List of QOS names
	Qoslevel []string `json:"qoslevel,omitempty"`
	// Number of shares allocated to this association
	Fairshare *int32 `json:"fairshare,omitempty"`
}

// NewV0040AssocRecSet instantiates a new V0040AssocRecSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040AssocRecSet() *V0040AssocRecSet {
	this := V0040AssocRecSet{}
	return &this
}

// NewV0040AssocRecSetWithDefaults instantiates a new V0040AssocRecSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040AssocRecSetWithDefaults() *V0040AssocRecSet {
	this := V0040AssocRecSet{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *V0040AssocRecSet) SetComment(v string) {
	o.Comment = &v
}

// GetDefaultqos returns the Defaultqos field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetDefaultqos() string {
	if o == nil || IsNil(o.Defaultqos) {
		var ret string
		return ret
	}
	return *o.Defaultqos
}

// GetDefaultqosOk returns a tuple with the Defaultqos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetDefaultqosOk() (*string, bool) {
	if o == nil || IsNil(o.Defaultqos) {
		return nil, false
	}
	return o.Defaultqos, true
}

// HasDefaultqos returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasDefaultqos() bool {
	if o != nil && !IsNil(o.Defaultqos) {
		return true
	}

	return false
}

// SetDefaultqos gets a reference to the given string and assigns it to the Defaultqos field.
func (o *V0040AssocRecSet) SetDefaultqos(v string) {
	o.Defaultqos = &v
}

// GetGrpjobs returns the Grpjobs field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetGrpjobs() V0040Uint32NoVal {
	if o == nil || IsNil(o.Grpjobs) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Grpjobs
}

// GetGrpjobsOk returns a tuple with the Grpjobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetGrpjobsOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Grpjobs) {
		return nil, false
	}
	return o.Grpjobs, true
}

// HasGrpjobs returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasGrpjobs() bool {
	if o != nil && !IsNil(o.Grpjobs) {
		return true
	}

	return false
}

// SetGrpjobs gets a reference to the given V0040Uint32NoVal and assigns it to the Grpjobs field.
func (o *V0040AssocRecSet) SetGrpjobs(v V0040Uint32NoVal) {
	o.Grpjobs = &v
}

// GetGrpjobsaccrue returns the Grpjobsaccrue field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetGrpjobsaccrue() V0040Uint32NoVal {
	if o == nil || IsNil(o.Grpjobsaccrue) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Grpjobsaccrue
}

// GetGrpjobsaccrueOk returns a tuple with the Grpjobsaccrue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetGrpjobsaccrueOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Grpjobsaccrue) {
		return nil, false
	}
	return o.Grpjobsaccrue, true
}

// HasGrpjobsaccrue returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasGrpjobsaccrue() bool {
	if o != nil && !IsNil(o.Grpjobsaccrue) {
		return true
	}

	return false
}

// SetGrpjobsaccrue gets a reference to the given V0040Uint32NoVal and assigns it to the Grpjobsaccrue field.
func (o *V0040AssocRecSet) SetGrpjobsaccrue(v V0040Uint32NoVal) {
	o.Grpjobsaccrue = &v
}

// GetGrpsubmitjobs returns the Grpsubmitjobs field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetGrpsubmitjobs() V0040Uint32NoVal {
	if o == nil || IsNil(o.Grpsubmitjobs) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Grpsubmitjobs
}

// GetGrpsubmitjobsOk returns a tuple with the Grpsubmitjobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetGrpsubmitjobsOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Grpsubmitjobs) {
		return nil, false
	}
	return o.Grpsubmitjobs, true
}

// HasGrpsubmitjobs returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasGrpsubmitjobs() bool {
	if o != nil && !IsNil(o.Grpsubmitjobs) {
		return true
	}

	return false
}

// SetGrpsubmitjobs gets a reference to the given V0040Uint32NoVal and assigns it to the Grpsubmitjobs field.
func (o *V0040AssocRecSet) SetGrpsubmitjobs(v V0040Uint32NoVal) {
	o.Grpsubmitjobs = &v
}

// GetGrptres returns the Grptres field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetGrptres() []V0040Tres {
	if o == nil || IsNil(o.Grptres) {
		var ret []V0040Tres
		return ret
	}
	return o.Grptres
}

// GetGrptresOk returns a tuple with the Grptres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetGrptresOk() ([]V0040Tres, bool) {
	if o == nil || IsNil(o.Grptres) {
		return nil, false
	}
	return o.Grptres, true
}

// HasGrptres returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasGrptres() bool {
	if o != nil && !IsNil(o.Grptres) {
		return true
	}

	return false
}

// SetGrptres gets a reference to the given []V0040Tres and assigns it to the Grptres field.
func (o *V0040AssocRecSet) SetGrptres(v []V0040Tres) {
	o.Grptres = v
}

// GetGrptresmins returns the Grptresmins field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetGrptresmins() []V0040Tres {
	if o == nil || IsNil(o.Grptresmins) {
		var ret []V0040Tres
		return ret
	}
	return o.Grptresmins
}

// GetGrptresminsOk returns a tuple with the Grptresmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetGrptresminsOk() ([]V0040Tres, bool) {
	if o == nil || IsNil(o.Grptresmins) {
		return nil, false
	}
	return o.Grptresmins, true
}

// HasGrptresmins returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasGrptresmins() bool {
	if o != nil && !IsNil(o.Grptresmins) {
		return true
	}

	return false
}

// SetGrptresmins gets a reference to the given []V0040Tres and assigns it to the Grptresmins field.
func (o *V0040AssocRecSet) SetGrptresmins(v []V0040Tres) {
	o.Grptresmins = v
}

// GetGrptresrunmins returns the Grptresrunmins field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetGrptresrunmins() []V0040Tres {
	if o == nil || IsNil(o.Grptresrunmins) {
		var ret []V0040Tres
		return ret
	}
	return o.Grptresrunmins
}

// GetGrptresrunminsOk returns a tuple with the Grptresrunmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetGrptresrunminsOk() ([]V0040Tres, bool) {
	if o == nil || IsNil(o.Grptresrunmins) {
		return nil, false
	}
	return o.Grptresrunmins, true
}

// HasGrptresrunmins returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasGrptresrunmins() bool {
	if o != nil && !IsNil(o.Grptresrunmins) {
		return true
	}

	return false
}

// SetGrptresrunmins gets a reference to the given []V0040Tres and assigns it to the Grptresrunmins field.
func (o *V0040AssocRecSet) SetGrptresrunmins(v []V0040Tres) {
	o.Grptresrunmins = v
}

// GetGrpwall returns the Grpwall field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetGrpwall() V0040Uint32NoVal {
	if o == nil || IsNil(o.Grpwall) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Grpwall
}

// GetGrpwallOk returns a tuple with the Grpwall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetGrpwallOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Grpwall) {
		return nil, false
	}
	return o.Grpwall, true
}

// HasGrpwall returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasGrpwall() bool {
	if o != nil && !IsNil(o.Grpwall) {
		return true
	}

	return false
}

// SetGrpwall gets a reference to the given V0040Uint32NoVal and assigns it to the Grpwall field.
func (o *V0040AssocRecSet) SetGrpwall(v V0040Uint32NoVal) {
	o.Grpwall = &v
}

// GetMaxjobs returns the Maxjobs field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetMaxjobs() V0040Uint32NoVal {
	if o == nil || IsNil(o.Maxjobs) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Maxjobs
}

// GetMaxjobsOk returns a tuple with the Maxjobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetMaxjobsOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Maxjobs) {
		return nil, false
	}
	return o.Maxjobs, true
}

// HasMaxjobs returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasMaxjobs() bool {
	if o != nil && !IsNil(o.Maxjobs) {
		return true
	}

	return false
}

// SetMaxjobs gets a reference to the given V0040Uint32NoVal and assigns it to the Maxjobs field.
func (o *V0040AssocRecSet) SetMaxjobs(v V0040Uint32NoVal) {
	o.Maxjobs = &v
}

// GetMaxjobsaccrue returns the Maxjobsaccrue field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetMaxjobsaccrue() V0040Uint32NoVal {
	if o == nil || IsNil(o.Maxjobsaccrue) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Maxjobsaccrue
}

// GetMaxjobsaccrueOk returns a tuple with the Maxjobsaccrue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetMaxjobsaccrueOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Maxjobsaccrue) {
		return nil, false
	}
	return o.Maxjobsaccrue, true
}

// HasMaxjobsaccrue returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasMaxjobsaccrue() bool {
	if o != nil && !IsNil(o.Maxjobsaccrue) {
		return true
	}

	return false
}

// SetMaxjobsaccrue gets a reference to the given V0040Uint32NoVal and assigns it to the Maxjobsaccrue field.
func (o *V0040AssocRecSet) SetMaxjobsaccrue(v V0040Uint32NoVal) {
	o.Maxjobsaccrue = &v
}

// GetMaxsubmitjobs returns the Maxsubmitjobs field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetMaxsubmitjobs() V0040Uint32NoVal {
	if o == nil || IsNil(o.Maxsubmitjobs) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Maxsubmitjobs
}

// GetMaxsubmitjobsOk returns a tuple with the Maxsubmitjobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetMaxsubmitjobsOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Maxsubmitjobs) {
		return nil, false
	}
	return o.Maxsubmitjobs, true
}

// HasMaxsubmitjobs returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasMaxsubmitjobs() bool {
	if o != nil && !IsNil(o.Maxsubmitjobs) {
		return true
	}

	return false
}

// SetMaxsubmitjobs gets a reference to the given V0040Uint32NoVal and assigns it to the Maxsubmitjobs field.
func (o *V0040AssocRecSet) SetMaxsubmitjobs(v V0040Uint32NoVal) {
	o.Maxsubmitjobs = &v
}

// GetMaxtresminsperjob returns the Maxtresminsperjob field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetMaxtresminsperjob() []V0040Tres {
	if o == nil || IsNil(o.Maxtresminsperjob) {
		var ret []V0040Tres
		return ret
	}
	return o.Maxtresminsperjob
}

// GetMaxtresminsperjobOk returns a tuple with the Maxtresminsperjob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetMaxtresminsperjobOk() ([]V0040Tres, bool) {
	if o == nil || IsNil(o.Maxtresminsperjob) {
		return nil, false
	}
	return o.Maxtresminsperjob, true
}

// HasMaxtresminsperjob returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasMaxtresminsperjob() bool {
	if o != nil && !IsNil(o.Maxtresminsperjob) {
		return true
	}

	return false
}

// SetMaxtresminsperjob gets a reference to the given []V0040Tres and assigns it to the Maxtresminsperjob field.
func (o *V0040AssocRecSet) SetMaxtresminsperjob(v []V0040Tres) {
	o.Maxtresminsperjob = v
}

// GetMaxtresrunmins returns the Maxtresrunmins field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetMaxtresrunmins() []V0040Tres {
	if o == nil || IsNil(o.Maxtresrunmins) {
		var ret []V0040Tres
		return ret
	}
	return o.Maxtresrunmins
}

// GetMaxtresrunminsOk returns a tuple with the Maxtresrunmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetMaxtresrunminsOk() ([]V0040Tres, bool) {
	if o == nil || IsNil(o.Maxtresrunmins) {
		return nil, false
	}
	return o.Maxtresrunmins, true
}

// HasMaxtresrunmins returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasMaxtresrunmins() bool {
	if o != nil && !IsNil(o.Maxtresrunmins) {
		return true
	}

	return false
}

// SetMaxtresrunmins gets a reference to the given []V0040Tres and assigns it to the Maxtresrunmins field.
func (o *V0040AssocRecSet) SetMaxtresrunmins(v []V0040Tres) {
	o.Maxtresrunmins = v
}

// GetMaxtresperjob returns the Maxtresperjob field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetMaxtresperjob() []V0040Tres {
	if o == nil || IsNil(o.Maxtresperjob) {
		var ret []V0040Tres
		return ret
	}
	return o.Maxtresperjob
}

// GetMaxtresperjobOk returns a tuple with the Maxtresperjob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetMaxtresperjobOk() ([]V0040Tres, bool) {
	if o == nil || IsNil(o.Maxtresperjob) {
		return nil, false
	}
	return o.Maxtresperjob, true
}

// HasMaxtresperjob returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasMaxtresperjob() bool {
	if o != nil && !IsNil(o.Maxtresperjob) {
		return true
	}

	return false
}

// SetMaxtresperjob gets a reference to the given []V0040Tres and assigns it to the Maxtresperjob field.
func (o *V0040AssocRecSet) SetMaxtresperjob(v []V0040Tres) {
	o.Maxtresperjob = v
}

// GetMaxtrespernode returns the Maxtrespernode field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetMaxtrespernode() []V0040Tres {
	if o == nil || IsNil(o.Maxtrespernode) {
		var ret []V0040Tres
		return ret
	}
	return o.Maxtrespernode
}

// GetMaxtrespernodeOk returns a tuple with the Maxtrespernode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetMaxtrespernodeOk() ([]V0040Tres, bool) {
	if o == nil || IsNil(o.Maxtrespernode) {
		return nil, false
	}
	return o.Maxtrespernode, true
}

// HasMaxtrespernode returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasMaxtrespernode() bool {
	if o != nil && !IsNil(o.Maxtrespernode) {
		return true
	}

	return false
}

// SetMaxtrespernode gets a reference to the given []V0040Tres and assigns it to the Maxtrespernode field.
func (o *V0040AssocRecSet) SetMaxtrespernode(v []V0040Tres) {
	o.Maxtrespernode = v
}

// GetMaxwalldurationperjob returns the Maxwalldurationperjob field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetMaxwalldurationperjob() V0040Uint32NoVal {
	if o == nil || IsNil(o.Maxwalldurationperjob) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Maxwalldurationperjob
}

// GetMaxwalldurationperjobOk returns a tuple with the Maxwalldurationperjob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetMaxwalldurationperjobOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Maxwalldurationperjob) {
		return nil, false
	}
	return o.Maxwalldurationperjob, true
}

// HasMaxwalldurationperjob returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasMaxwalldurationperjob() bool {
	if o != nil && !IsNil(o.Maxwalldurationperjob) {
		return true
	}

	return false
}

// SetMaxwalldurationperjob gets a reference to the given V0040Uint32NoVal and assigns it to the Maxwalldurationperjob field.
func (o *V0040AssocRecSet) SetMaxwalldurationperjob(v V0040Uint32NoVal) {
	o.Maxwalldurationperjob = &v
}

// GetMinpriothresh returns the Minpriothresh field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetMinpriothresh() V0040Uint32NoVal {
	if o == nil || IsNil(o.Minpriothresh) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Minpriothresh
}

// GetMinpriothreshOk returns a tuple with the Minpriothresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetMinpriothreshOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Minpriothresh) {
		return nil, false
	}
	return o.Minpriothresh, true
}

// HasMinpriothresh returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasMinpriothresh() bool {
	if o != nil && !IsNil(o.Minpriothresh) {
		return true
	}

	return false
}

// SetMinpriothresh gets a reference to the given V0040Uint32NoVal and assigns it to the Minpriothresh field.
func (o *V0040AssocRecSet) SetMinpriothresh(v V0040Uint32NoVal) {
	o.Minpriothresh = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *V0040AssocRecSet) SetParent(v string) {
	o.Parent = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetPriority() V0040Uint32NoVal {
	if o == nil || IsNil(o.Priority) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetPriorityOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given V0040Uint32NoVal and assigns it to the Priority field.
func (o *V0040AssocRecSet) SetPriority(v V0040Uint32NoVal) {
	o.Priority = &v
}

// GetQoslevel returns the Qoslevel field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetQoslevel() []string {
	if o == nil || IsNil(o.Qoslevel) {
		var ret []string
		return ret
	}
	return o.Qoslevel
}

// GetQoslevelOk returns a tuple with the Qoslevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetQoslevelOk() ([]string, bool) {
	if o == nil || IsNil(o.Qoslevel) {
		return nil, false
	}
	return o.Qoslevel, true
}

// HasQoslevel returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasQoslevel() bool {
	if o != nil && !IsNil(o.Qoslevel) {
		return true
	}

	return false
}

// SetQoslevel gets a reference to the given []string and assigns it to the Qoslevel field.
func (o *V0040AssocRecSet) SetQoslevel(v []string) {
	o.Qoslevel = v
}

// GetFairshare returns the Fairshare field value if set, zero value otherwise.
func (o *V0040AssocRecSet) GetFairshare() int32 {
	if o == nil || IsNil(o.Fairshare) {
		var ret int32
		return ret
	}
	return *o.Fairshare
}

// GetFairshareOk returns a tuple with the Fairshare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040AssocRecSet) GetFairshareOk() (*int32, bool) {
	if o == nil || IsNil(o.Fairshare) {
		return nil, false
	}
	return o.Fairshare, true
}

// HasFairshare returns a boolean if a field has been set.
func (o *V0040AssocRecSet) HasFairshare() bool {
	if o != nil && !IsNil(o.Fairshare) {
		return true
	}

	return false
}

// SetFairshare gets a reference to the given int32 and assigns it to the Fairshare field.
func (o *V0040AssocRecSet) SetFairshare(v int32) {
	o.Fairshare = &v
}

func (o V0040AssocRecSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040AssocRecSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Defaultqos) {
		toSerialize["defaultqos"] = o.Defaultqos
	}
	if !IsNil(o.Grpjobs) {
		toSerialize["grpjobs"] = o.Grpjobs
	}
	if !IsNil(o.Grpjobsaccrue) {
		toSerialize["grpjobsaccrue"] = o.Grpjobsaccrue
	}
	if !IsNil(o.Grpsubmitjobs) {
		toSerialize["grpsubmitjobs"] = o.Grpsubmitjobs
	}
	if !IsNil(o.Grptres) {
		toSerialize["grptres"] = o.Grptres
	}
	if !IsNil(o.Grptresmins) {
		toSerialize["grptresmins"] = o.Grptresmins
	}
	if !IsNil(o.Grptresrunmins) {
		toSerialize["grptresrunmins"] = o.Grptresrunmins
	}
	if !IsNil(o.Grpwall) {
		toSerialize["grpwall"] = o.Grpwall
	}
	if !IsNil(o.Maxjobs) {
		toSerialize["maxjobs"] = o.Maxjobs
	}
	if !IsNil(o.Maxjobsaccrue) {
		toSerialize["maxjobsaccrue"] = o.Maxjobsaccrue
	}
	if !IsNil(o.Maxsubmitjobs) {
		toSerialize["maxsubmitjobs"] = o.Maxsubmitjobs
	}
	if !IsNil(o.Maxtresminsperjob) {
		toSerialize["maxtresminsperjob"] = o.Maxtresminsperjob
	}
	if !IsNil(o.Maxtresrunmins) {
		toSerialize["maxtresrunmins"] = o.Maxtresrunmins
	}
	if !IsNil(o.Maxtresperjob) {
		toSerialize["maxtresperjob"] = o.Maxtresperjob
	}
	if !IsNil(o.Maxtrespernode) {
		toSerialize["maxtrespernode"] = o.Maxtrespernode
	}
	if !IsNil(o.Maxwalldurationperjob) {
		toSerialize["maxwalldurationperjob"] = o.Maxwalldurationperjob
	}
	if !IsNil(o.Minpriothresh) {
		toSerialize["minpriothresh"] = o.Minpriothresh
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Qoslevel) {
		toSerialize["qoslevel"] = o.Qoslevel
	}
	if !IsNil(o.Fairshare) {
		toSerialize["fairshare"] = o.Fairshare
	}
	return toSerialize, nil
}

type NullableV0040AssocRecSet struct {
	value *V0040AssocRecSet
	isSet bool
}

func (v NullableV0040AssocRecSet) Get() *V0040AssocRecSet {
	return v.value
}

func (v *NullableV0040AssocRecSet) Set(val *V0040AssocRecSet) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040AssocRecSet) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040AssocRecSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040AssocRecSet(val *V0040AssocRecSet) *NullableV0040AssocRecSet {
	return &NullableV0040AssocRecSet{value: val, isSet: true}
}

func (v NullableV0040AssocRecSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040AssocRecSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


