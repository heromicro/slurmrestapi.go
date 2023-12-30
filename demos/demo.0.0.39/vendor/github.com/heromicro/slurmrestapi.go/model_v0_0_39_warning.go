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

// checks if the V0039Warning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039Warning{}

// V0039Warning struct for V0039Warning
type V0039Warning struct {
	// Earning message
	Warning *string `json:"warning,omitempty"`
	// Where error occurred in the source
	Source *string `json:"source,omitempty"`
	// Explaination of cause of error
	Description *string `json:"description,omitempty"`
}

// NewV0039Warning instantiates a new V0039Warning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039Warning() *V0039Warning {
	this := V0039Warning{}
	return &this
}

// NewV0039WarningWithDefaults instantiates a new V0039Warning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039WarningWithDefaults() *V0039Warning {
	this := V0039Warning{}
	return &this
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *V0039Warning) GetWarning() string {
	if o == nil || IsNil(o.Warning) {
		var ret string
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Warning) GetWarningOk() (*string, bool) {
	if o == nil || IsNil(o.Warning) {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *V0039Warning) HasWarning() bool {
	if o != nil && !IsNil(o.Warning) {
		return true
	}

	return false
}

// SetWarning gets a reference to the given string and assigns it to the Warning field.
func (o *V0039Warning) SetWarning(v string) {
	o.Warning = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *V0039Warning) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Warning) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *V0039Warning) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *V0039Warning) SetSource(v string) {
	o.Source = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V0039Warning) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Warning) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V0039Warning) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *V0039Warning) SetDescription(v string) {
	o.Description = &v
}

func (o V0039Warning) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039Warning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Warning) {
		toSerialize["warning"] = o.Warning
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableV0039Warning struct {
	value *V0039Warning
	isSet bool
}

func (v NullableV0039Warning) Get() *V0039Warning {
	return v.value
}

func (v *NullableV0039Warning) Set(val *V0039Warning) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039Warning) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039Warning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039Warning(val *V0039Warning) *NullableV0039Warning {
	return &NullableV0039Warning{value: val, isSet: true}
}

func (v NullableV0039Warning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039Warning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


