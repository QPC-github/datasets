/*
 * NCBI Datasets API
 *
 * ### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/rehydrate/), and retrieve the individual data files at a later time. 
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
)

// V1reportsFeatureCounts struct for V1reportsFeatureCounts
type V1reportsFeatureCounts struct {
	GeneCounts *V1reportsGeneCounts `json:"gene_counts,omitempty"`
}

// NewV1reportsFeatureCounts instantiates a new V1reportsFeatureCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsFeatureCounts() *V1reportsFeatureCounts {
	this := V1reportsFeatureCounts{}
	return &this
}

// NewV1reportsFeatureCountsWithDefaults instantiates a new V1reportsFeatureCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsFeatureCountsWithDefaults() *V1reportsFeatureCounts {
	this := V1reportsFeatureCounts{}
	return &this
}

// GetGeneCounts returns the GeneCounts field value if set, zero value otherwise.
func (o *V1reportsFeatureCounts) GetGeneCounts() V1reportsGeneCounts {
	if o == nil || o.GeneCounts == nil {
		var ret V1reportsGeneCounts
		return ret
	}
	return *o.GeneCounts
}

// GetGeneCountsOk returns a tuple with the GeneCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsFeatureCounts) GetGeneCountsOk() (*V1reportsGeneCounts, bool) {
	if o == nil || o.GeneCounts == nil {
		return nil, false
	}
	return o.GeneCounts, true
}

// HasGeneCounts returns a boolean if a field has been set.
func (o *V1reportsFeatureCounts) HasGeneCounts() bool {
	if o != nil && o.GeneCounts != nil {
		return true
	}

	return false
}

// SetGeneCounts gets a reference to the given V1reportsGeneCounts and assigns it to the GeneCounts field.
func (o *V1reportsFeatureCounts) SetGeneCounts(v V1reportsGeneCounts) {
	o.GeneCounts = &v
}

func (o V1reportsFeatureCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GeneCounts != nil  {
		toSerialize["gene_counts"] = o.GeneCounts
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsFeatureCounts struct {
	value *V1reportsFeatureCounts
	isSet bool
}

func (v NullableV1reportsFeatureCounts) Get() *V1reportsFeatureCounts {
	return v.value
}

func (v *NullableV1reportsFeatureCounts) Set(val *V1reportsFeatureCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsFeatureCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsFeatureCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsFeatureCounts(val *V1reportsFeatureCounts) *NullableV1reportsFeatureCounts {
	return &NullableV1reportsFeatureCounts{value: val, isSet: true}
}

func (v NullableV1reportsFeatureCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsFeatureCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


