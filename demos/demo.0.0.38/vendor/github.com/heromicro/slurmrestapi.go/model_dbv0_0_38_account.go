/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.38
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrestapi

import (
	"encoding/json"
)

// checks if the Dbv0038Account type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0038Account{}

// Dbv0038Account Account description
type Dbv0038Account struct {
	// List of assigned associations
	Associations []Dbv0038AssociationShortInfo `json:"associations,omitempty"`
	// List of assigned coordinators
	Coordinators []Dbv0038CoordinatorInfo `json:"coordinators,omitempty"`
	// Description of account
	Description *string `json:"description,omitempty"`
	// Name of account
	Name *string `json:"name,omitempty"`
	// Assigned organization of account
	Organization *string `json:"organization,omitempty"`
	// List of properties of account
	Flags []string `json:"flags,omitempty"`
}

// NewDbv0038Account instantiates a new Dbv0038Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0038Account() *Dbv0038Account {
	this := Dbv0038Account{}
	return &this
}

// NewDbv0038AccountWithDefaults instantiates a new Dbv0038Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0038AccountWithDefaults() *Dbv0038Account {
	this := Dbv0038Account{}
	return &this
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *Dbv0038Account) GetAssociations() []Dbv0038AssociationShortInfo {
	if o == nil || IsNil(o.Associations) {
		var ret []Dbv0038AssociationShortInfo
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Account) GetAssociationsOk() ([]Dbv0038AssociationShortInfo, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *Dbv0038Account) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []Dbv0038AssociationShortInfo and assigns it to the Associations field.
func (o *Dbv0038Account) SetAssociations(v []Dbv0038AssociationShortInfo) {
	o.Associations = v
}

// GetCoordinators returns the Coordinators field value if set, zero value otherwise.
func (o *Dbv0038Account) GetCoordinators() []Dbv0038CoordinatorInfo {
	if o == nil || IsNil(o.Coordinators) {
		var ret []Dbv0038CoordinatorInfo
		return ret
	}
	return o.Coordinators
}

// GetCoordinatorsOk returns a tuple with the Coordinators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Account) GetCoordinatorsOk() ([]Dbv0038CoordinatorInfo, bool) {
	if o == nil || IsNil(o.Coordinators) {
		return nil, false
	}
	return o.Coordinators, true
}

// HasCoordinators returns a boolean if a field has been set.
func (o *Dbv0038Account) HasCoordinators() bool {
	if o != nil && !IsNil(o.Coordinators) {
		return true
	}

	return false
}

// SetCoordinators gets a reference to the given []Dbv0038CoordinatorInfo and assigns it to the Coordinators field.
func (o *Dbv0038Account) SetCoordinators(v []Dbv0038CoordinatorInfo) {
	o.Coordinators = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Dbv0038Account) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Account) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Dbv0038Account) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Dbv0038Account) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Dbv0038Account) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Account) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Dbv0038Account) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Dbv0038Account) SetName(v string) {
	o.Name = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *Dbv0038Account) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Account) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *Dbv0038Account) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *Dbv0038Account) SetOrganization(v string) {
	o.Organization = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *Dbv0038Account) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0038Account) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *Dbv0038Account) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *Dbv0038Account) SetFlags(v []string) {
	o.Flags = v
}

func (o Dbv0038Account) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0038Account) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	if !IsNil(o.Coordinators) {
		toSerialize["coordinators"] = o.Coordinators
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	return toSerialize, nil
}

type NullableDbv0038Account struct {
	value *Dbv0038Account
	isSet bool
}

func (v NullableDbv0038Account) Get() *Dbv0038Account {
	return v.value
}

func (v *NullableDbv0038Account) Set(val *Dbv0038Account) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0038Account) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0038Account) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0038Account(val *Dbv0038Account) *NullableDbv0038Account {
	return &NullableDbv0038Account{value: val, isSet: true}
}

func (v NullableDbv0038Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0038Account) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


