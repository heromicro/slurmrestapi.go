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
	"bytes"
	"fmt"
)

// checks if the V0039User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039User{}

// V0039User struct for V0039User
type V0039User struct {
	AdministratorLevel []string `json:"administrator_level,omitempty"`
	Associations []V0039AssocShort `json:"associations,omitempty"`
	Coordinators []V0039Coord `json:"coordinators,omitempty"`
	Default *V0040UserDefault `json:"default,omitempty"`
	Flags []string `json:"flags,omitempty"`
	Name string `json:"name"`
	OldName *string `json:"old_name,omitempty"`
	Wckeys []V0039Wckey `json:"wckeys,omitempty"`
}

type _V0039User V0039User

// NewV0039User instantiates a new V0039User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039User(name string) *V0039User {
	this := V0039User{}
	this.Name = name
	return &this
}

// NewV0039UserWithDefaults instantiates a new V0039User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039UserWithDefaults() *V0039User {
	this := V0039User{}
	return &this
}

// GetAdministratorLevel returns the AdministratorLevel field value if set, zero value otherwise.
func (o *V0039User) GetAdministratorLevel() []string {
	if o == nil || IsNil(o.AdministratorLevel) {
		var ret []string
		return ret
	}
	return o.AdministratorLevel
}

// GetAdministratorLevelOk returns a tuple with the AdministratorLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039User) GetAdministratorLevelOk() ([]string, bool) {
	if o == nil || IsNil(o.AdministratorLevel) {
		return nil, false
	}
	return o.AdministratorLevel, true
}

// HasAdministratorLevel returns a boolean if a field has been set.
func (o *V0039User) HasAdministratorLevel() bool {
	if o != nil && !IsNil(o.AdministratorLevel) {
		return true
	}

	return false
}

// SetAdministratorLevel gets a reference to the given []string and assigns it to the AdministratorLevel field.
func (o *V0039User) SetAdministratorLevel(v []string) {
	o.AdministratorLevel = v
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *V0039User) GetAssociations() []V0039AssocShort {
	if o == nil || IsNil(o.Associations) {
		var ret []V0039AssocShort
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039User) GetAssociationsOk() ([]V0039AssocShort, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *V0039User) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []V0039AssocShort and assigns it to the Associations field.
func (o *V0039User) SetAssociations(v []V0039AssocShort) {
	o.Associations = v
}

// GetCoordinators returns the Coordinators field value if set, zero value otherwise.
func (o *V0039User) GetCoordinators() []V0039Coord {
	if o == nil || IsNil(o.Coordinators) {
		var ret []V0039Coord
		return ret
	}
	return o.Coordinators
}

// GetCoordinatorsOk returns a tuple with the Coordinators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039User) GetCoordinatorsOk() ([]V0039Coord, bool) {
	if o == nil || IsNil(o.Coordinators) {
		return nil, false
	}
	return o.Coordinators, true
}

// HasCoordinators returns a boolean if a field has been set.
func (o *V0039User) HasCoordinators() bool {
	if o != nil && !IsNil(o.Coordinators) {
		return true
	}

	return false
}

// SetCoordinators gets a reference to the given []V0039Coord and assigns it to the Coordinators field.
func (o *V0039User) SetCoordinators(v []V0039Coord) {
	o.Coordinators = v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *V0039User) GetDefault() V0040UserDefault {
	if o == nil || IsNil(o.Default) {
		var ret V0040UserDefault
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039User) GetDefaultOk() (*V0040UserDefault, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *V0039User) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given V0040UserDefault and assigns it to the Default field.
func (o *V0039User) SetDefault(v V0040UserDefault) {
	o.Default = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0039User) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039User) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0039User) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0039User) SetFlags(v []string) {
	o.Flags = v
}

// GetName returns the Name field value
func (o *V0039User) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V0039User) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V0039User) SetName(v string) {
	o.Name = v
}

// GetOldName returns the OldName field value if set, zero value otherwise.
func (o *V0039User) GetOldName() string {
	if o == nil || IsNil(o.OldName) {
		var ret string
		return ret
	}
	return *o.OldName
}

// GetOldNameOk returns a tuple with the OldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039User) GetOldNameOk() (*string, bool) {
	if o == nil || IsNil(o.OldName) {
		return nil, false
	}
	return o.OldName, true
}

// HasOldName returns a boolean if a field has been set.
func (o *V0039User) HasOldName() bool {
	if o != nil && !IsNil(o.OldName) {
		return true
	}

	return false
}

// SetOldName gets a reference to the given string and assigns it to the OldName field.
func (o *V0039User) SetOldName(v string) {
	o.OldName = &v
}

// GetWckeys returns the Wckeys field value if set, zero value otherwise.
func (o *V0039User) GetWckeys() []V0039Wckey {
	if o == nil || IsNil(o.Wckeys) {
		var ret []V0039Wckey
		return ret
	}
	return o.Wckeys
}

// GetWckeysOk returns a tuple with the Wckeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039User) GetWckeysOk() ([]V0039Wckey, bool) {
	if o == nil || IsNil(o.Wckeys) {
		return nil, false
	}
	return o.Wckeys, true
}

// HasWckeys returns a boolean if a field has been set.
func (o *V0039User) HasWckeys() bool {
	if o != nil && !IsNil(o.Wckeys) {
		return true
	}

	return false
}

// SetWckeys gets a reference to the given []V0039Wckey and assigns it to the Wckeys field.
func (o *V0039User) SetWckeys(v []V0039Wckey) {
	o.Wckeys = v
}

func (o V0039User) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdministratorLevel) {
		toSerialize["administrator_level"] = o.AdministratorLevel
	}
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	if !IsNil(o.Coordinators) {
		toSerialize["coordinators"] = o.Coordinators
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OldName) {
		toSerialize["old_name"] = o.OldName
	}
	if !IsNil(o.Wckeys) {
		toSerialize["wckeys"] = o.Wckeys
	}
	return toSerialize, nil
}

func (o *V0039User) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV0039User := _V0039User{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0039User)

	if err != nil {
		return err
	}

	*o = V0039User(varV0039User)

	return err
}

type NullableV0039User struct {
	value *V0039User
	isSet bool
}

func (v NullableV0039User) Get() *V0039User {
	return v.value
}

func (v *NullableV0039User) Set(val *V0039User) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039User) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039User) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039User(val *V0039User) *NullableV0039User {
	return &NullableV0039User{value: val, isSet: true}
}

func (v NullableV0039User) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039User) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


