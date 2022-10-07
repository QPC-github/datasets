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

// V1AssemblyDatasetDescriptors struct for V1AssemblyDatasetDescriptors
type V1AssemblyDatasetDescriptors struct {
	Datasets *[]V1AssemblyDatasetDescriptor `json:"datasets,omitempty"`
	TotalCount *int32 `json:"total_count,omitempty"`
}

// NewV1AssemblyDatasetDescriptors instantiates a new V1AssemblyDatasetDescriptors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AssemblyDatasetDescriptors() *V1AssemblyDatasetDescriptors {
	this := V1AssemblyDatasetDescriptors{}
	return &this
}

// NewV1AssemblyDatasetDescriptorsWithDefaults instantiates a new V1AssemblyDatasetDescriptors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AssemblyDatasetDescriptorsWithDefaults() *V1AssemblyDatasetDescriptors {
	this := V1AssemblyDatasetDescriptors{}
	return &this
}

// GetDatasets returns the Datasets field value if set, zero value otherwise.
func (o *V1AssemblyDatasetDescriptors) GetDatasets() []V1AssemblyDatasetDescriptor {
	if o == nil || o.Datasets == nil {
		var ret []V1AssemblyDatasetDescriptor
		return ret
	}
	return *o.Datasets
}

// GetDatasetsOk returns a tuple with the Datasets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetDescriptors) GetDatasetsOk() (*[]V1AssemblyDatasetDescriptor, bool) {
	if o == nil || o.Datasets == nil {
		return nil, false
	}
	return o.Datasets, true
}

// HasDatasets returns a boolean if a field has been set.
func (o *V1AssemblyDatasetDescriptors) HasDatasets() bool {
	if o != nil && o.Datasets != nil {
		return true
	}

	return false
}

// SetDatasets gets a reference to the given []V1AssemblyDatasetDescriptor and assigns it to the Datasets field.
func (o *V1AssemblyDatasetDescriptors) SetDatasets(v []V1AssemblyDatasetDescriptor) {
	o.Datasets = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *V1AssemblyDatasetDescriptors) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetDescriptors) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *V1AssemblyDatasetDescriptors) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *V1AssemblyDatasetDescriptors) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o V1AssemblyDatasetDescriptors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Datasets != nil && len(o.GetDatasets()) > 0  {
		toSerialize["datasets"] = o.Datasets
	}
	if o.TotalCount != nil  {
		toSerialize["total_count"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullableV1AssemblyDatasetDescriptors struct {
	value *V1AssemblyDatasetDescriptors
	isSet bool
}

func (v NullableV1AssemblyDatasetDescriptors) Get() *V1AssemblyDatasetDescriptors {
	return v.value
}

func (v *NullableV1AssemblyDatasetDescriptors) Set(val *V1AssemblyDatasetDescriptors) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AssemblyDatasetDescriptors) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AssemblyDatasetDescriptors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AssemblyDatasetDescriptors(val *V1AssemblyDatasetDescriptors) *NullableV1AssemblyDatasetDescriptors {
	return &NullableV1AssemblyDatasetDescriptors{value: val, isSet: true}
}

func (v NullableV1AssemblyDatasetDescriptors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AssemblyDatasetDescriptors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


