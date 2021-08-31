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

// V1TaxonomyFilteredSubtreeResponseEdgesEntry A map.
type V1TaxonomyFilteredSubtreeResponseEdgesEntry struct {
	Default *V1TaxonomyFilteredSubtreeResponseEdge `json:"default,omitempty"`
}

// NewV1TaxonomyFilteredSubtreeResponseEdgesEntry instantiates a new V1TaxonomyFilteredSubtreeResponseEdgesEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1TaxonomyFilteredSubtreeResponseEdgesEntry() *V1TaxonomyFilteredSubtreeResponseEdgesEntry {
	this := V1TaxonomyFilteredSubtreeResponseEdgesEntry{}
	return &this
}

// NewV1TaxonomyFilteredSubtreeResponseEdgesEntryWithDefaults instantiates a new V1TaxonomyFilteredSubtreeResponseEdgesEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1TaxonomyFilteredSubtreeResponseEdgesEntryWithDefaults() *V1TaxonomyFilteredSubtreeResponseEdgesEntry {
	this := V1TaxonomyFilteredSubtreeResponseEdgesEntry{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *V1TaxonomyFilteredSubtreeResponseEdgesEntry) GetDefault() V1TaxonomyFilteredSubtreeResponseEdge {
	if o == nil || o.Default == nil {
		var ret V1TaxonomyFilteredSubtreeResponseEdge
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TaxonomyFilteredSubtreeResponseEdgesEntry) GetDefaultOk() (*V1TaxonomyFilteredSubtreeResponseEdge, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *V1TaxonomyFilteredSubtreeResponseEdgesEntry) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given V1TaxonomyFilteredSubtreeResponseEdge and assigns it to the Default field.
func (o *V1TaxonomyFilteredSubtreeResponseEdgesEntry) SetDefault(v V1TaxonomyFilteredSubtreeResponseEdge) {
	o.Default = &v
}

func (o V1TaxonomyFilteredSubtreeResponseEdgesEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Default != nil  {
		toSerialize["default"] = o.Default
	}
	return json.Marshal(toSerialize)
}

type NullableV1TaxonomyFilteredSubtreeResponseEdgesEntry struct {
	value *V1TaxonomyFilteredSubtreeResponseEdgesEntry
	isSet bool
}

func (v NullableV1TaxonomyFilteredSubtreeResponseEdgesEntry) Get() *V1TaxonomyFilteredSubtreeResponseEdgesEntry {
	return v.value
}

func (v *NullableV1TaxonomyFilteredSubtreeResponseEdgesEntry) Set(val *V1TaxonomyFilteredSubtreeResponseEdgesEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableV1TaxonomyFilteredSubtreeResponseEdgesEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableV1TaxonomyFilteredSubtreeResponseEdgesEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1TaxonomyFilteredSubtreeResponseEdgesEntry(val *V1TaxonomyFilteredSubtreeResponseEdgesEntry) *NullableV1TaxonomyFilteredSubtreeResponseEdgesEntry {
	return &NullableV1TaxonomyFilteredSubtreeResponseEdgesEntry{value: val, isSet: true}
}

func (v NullableV1TaxonomyFilteredSubtreeResponseEdgesEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1TaxonomyFilteredSubtreeResponseEdgesEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


