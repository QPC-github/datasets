/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets version 1 API is considred stable and will not be subject to breaking changes. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
)

// V1AssemblyDatasetAvailability struct for V1AssemblyDatasetAvailability
type V1AssemblyDatasetAvailability struct {
	ValidAssemblies *[]string `json:"valid_assemblies,omitempty"`
	InvalidAssemblies *[]string `json:"invalid_assemblies,omitempty"`
	Reason *string `json:"reason,omitempty"`
}

// NewV1AssemblyDatasetAvailability instantiates a new V1AssemblyDatasetAvailability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AssemblyDatasetAvailability() *V1AssemblyDatasetAvailability {
	this := V1AssemblyDatasetAvailability{}
	return &this
}

// NewV1AssemblyDatasetAvailabilityWithDefaults instantiates a new V1AssemblyDatasetAvailability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AssemblyDatasetAvailabilityWithDefaults() *V1AssemblyDatasetAvailability {
	this := V1AssemblyDatasetAvailability{}
	return &this
}

// GetValidAssemblies returns the ValidAssemblies field value if set, zero value otherwise.
func (o *V1AssemblyDatasetAvailability) GetValidAssemblies() []string {
	if o == nil || o.ValidAssemblies == nil {
		var ret []string
		return ret
	}
	return *o.ValidAssemblies
}

// GetValidAssembliesOk returns a tuple with the ValidAssemblies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetAvailability) GetValidAssembliesOk() (*[]string, bool) {
	if o == nil || o.ValidAssemblies == nil {
		return nil, false
	}
	return o.ValidAssemblies, true
}

// HasValidAssemblies returns a boolean if a field has been set.
func (o *V1AssemblyDatasetAvailability) HasValidAssemblies() bool {
	if o != nil && o.ValidAssemblies != nil {
		return true
	}

	return false
}

// SetValidAssemblies gets a reference to the given []string and assigns it to the ValidAssemblies field.
func (o *V1AssemblyDatasetAvailability) SetValidAssemblies(v []string) {
	o.ValidAssemblies = &v
}

// GetInvalidAssemblies returns the InvalidAssemblies field value if set, zero value otherwise.
func (o *V1AssemblyDatasetAvailability) GetInvalidAssemblies() []string {
	if o == nil || o.InvalidAssemblies == nil {
		var ret []string
		return ret
	}
	return *o.InvalidAssemblies
}

// GetInvalidAssembliesOk returns a tuple with the InvalidAssemblies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetAvailability) GetInvalidAssembliesOk() (*[]string, bool) {
	if o == nil || o.InvalidAssemblies == nil {
		return nil, false
	}
	return o.InvalidAssemblies, true
}

// HasInvalidAssemblies returns a boolean if a field has been set.
func (o *V1AssemblyDatasetAvailability) HasInvalidAssemblies() bool {
	if o != nil && o.InvalidAssemblies != nil {
		return true
	}

	return false
}

// SetInvalidAssemblies gets a reference to the given []string and assigns it to the InvalidAssemblies field.
func (o *V1AssemblyDatasetAvailability) SetInvalidAssemblies(v []string) {
	o.InvalidAssemblies = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *V1AssemblyDatasetAvailability) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetAvailability) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *V1AssemblyDatasetAvailability) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *V1AssemblyDatasetAvailability) SetReason(v string) {
	o.Reason = &v
}

func (o V1AssemblyDatasetAvailability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ValidAssemblies != nil && len(o.GetValidAssemblies()) > 0  {
		toSerialize["valid_assemblies"] = o.ValidAssemblies
	}
	if o.InvalidAssemblies != nil && len(o.GetInvalidAssemblies()) > 0  {
		toSerialize["invalid_assemblies"] = o.InvalidAssemblies
	}
	if o.Reason != nil  {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableV1AssemblyDatasetAvailability struct {
	value *V1AssemblyDatasetAvailability
	isSet bool
}

func (v NullableV1AssemblyDatasetAvailability) Get() *V1AssemblyDatasetAvailability {
	return v.value
}

func (v *NullableV1AssemblyDatasetAvailability) Set(val *V1AssemblyDatasetAvailability) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AssemblyDatasetAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AssemblyDatasetAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AssemblyDatasetAvailability(val *V1AssemblyDatasetAvailability) *NullableV1AssemblyDatasetAvailability {
	return &NullableV1AssemblyDatasetAvailability{value: val, isSet: true}
}

func (v NullableV1AssemblyDatasetAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AssemblyDatasetAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


