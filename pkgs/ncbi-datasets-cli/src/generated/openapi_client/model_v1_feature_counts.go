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

// V1FeatureCounts struct for V1FeatureCounts
type V1FeatureCounts struct {
	GeneCounts *V1GeneCounts `json:"gene_counts,omitempty"`
}

// NewV1FeatureCounts instantiates a new V1FeatureCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1FeatureCounts() *V1FeatureCounts {
	this := V1FeatureCounts{}
	return &this
}

// NewV1FeatureCountsWithDefaults instantiates a new V1FeatureCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1FeatureCountsWithDefaults() *V1FeatureCounts {
	this := V1FeatureCounts{}
	return &this
}

// GetGeneCounts returns the GeneCounts field value if set, zero value otherwise.
func (o *V1FeatureCounts) GetGeneCounts() V1GeneCounts {
	if o == nil || o.GeneCounts == nil {
		var ret V1GeneCounts
		return ret
	}
	return *o.GeneCounts
}

// GetGeneCountsOk returns a tuple with the GeneCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1FeatureCounts) GetGeneCountsOk() (*V1GeneCounts, bool) {
	if o == nil || o.GeneCounts == nil {
		return nil, false
	}
	return o.GeneCounts, true
}

// HasGeneCounts returns a boolean if a field has been set.
func (o *V1FeatureCounts) HasGeneCounts() bool {
	if o != nil && o.GeneCounts != nil {
		return true
	}

	return false
}

// SetGeneCounts gets a reference to the given V1GeneCounts and assigns it to the GeneCounts field.
func (o *V1FeatureCounts) SetGeneCounts(v V1GeneCounts) {
	o.GeneCounts = &v
}

func (o V1FeatureCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GeneCounts != nil  {
		toSerialize["gene_counts"] = o.GeneCounts
	}
	return json.Marshal(toSerialize)
}

type NullableV1FeatureCounts struct {
	value *V1FeatureCounts
	isSet bool
}

func (v NullableV1FeatureCounts) Get() *V1FeatureCounts {
	return v.value
}

func (v *NullableV1FeatureCounts) Set(val *V1FeatureCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableV1FeatureCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableV1FeatureCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1FeatureCounts(val *V1FeatureCounts) *NullableV1FeatureCounts {
	return &NullableV1FeatureCounts{value: val, isSet: true}
}

func (v NullableV1FeatureCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1FeatureCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


