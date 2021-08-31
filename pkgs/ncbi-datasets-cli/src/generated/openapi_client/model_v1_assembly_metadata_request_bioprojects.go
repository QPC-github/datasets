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

// V1AssemblyMetadataRequestBioprojects struct for V1AssemblyMetadataRequestBioprojects
type V1AssemblyMetadataRequestBioprojects struct {
	Accessions *[]string `json:"accessions,omitempty"`
}

// NewV1AssemblyMetadataRequestBioprojects instantiates a new V1AssemblyMetadataRequestBioprojects object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AssemblyMetadataRequestBioprojects() *V1AssemblyMetadataRequestBioprojects {
	this := V1AssemblyMetadataRequestBioprojects{}
	return &this
}

// NewV1AssemblyMetadataRequestBioprojectsWithDefaults instantiates a new V1AssemblyMetadataRequestBioprojects object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AssemblyMetadataRequestBioprojectsWithDefaults() *V1AssemblyMetadataRequestBioprojects {
	this := V1AssemblyMetadataRequestBioprojects{}
	return &this
}

// GetAccessions returns the Accessions field value if set, zero value otherwise.
func (o *V1AssemblyMetadataRequestBioprojects) GetAccessions() []string {
	if o == nil || o.Accessions == nil {
		var ret []string
		return ret
	}
	return *o.Accessions
}

// GetAccessionsOk returns a tuple with the Accessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyMetadataRequestBioprojects) GetAccessionsOk() (*[]string, bool) {
	if o == nil || o.Accessions == nil {
		return nil, false
	}
	return o.Accessions, true
}

// HasAccessions returns a boolean if a field has been set.
func (o *V1AssemblyMetadataRequestBioprojects) HasAccessions() bool {
	if o != nil && o.Accessions != nil {
		return true
	}

	return false
}

// SetAccessions gets a reference to the given []string and assigns it to the Accessions field.
func (o *V1AssemblyMetadataRequestBioprojects) SetAccessions(v []string) {
	o.Accessions = &v
}

func (o V1AssemblyMetadataRequestBioprojects) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accessions != nil && len(o.GetAccessions()) > 0  {
		toSerialize["accessions"] = o.Accessions
	}
	return json.Marshal(toSerialize)
}

type NullableV1AssemblyMetadataRequestBioprojects struct {
	value *V1AssemblyMetadataRequestBioprojects
	isSet bool
}

func (v NullableV1AssemblyMetadataRequestBioprojects) Get() *V1AssemblyMetadataRequestBioprojects {
	return v.value
}

func (v *NullableV1AssemblyMetadataRequestBioprojects) Set(val *V1AssemblyMetadataRequestBioprojects) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AssemblyMetadataRequestBioprojects) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AssemblyMetadataRequestBioprojects) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AssemblyMetadataRequestBioprojects(val *V1AssemblyMetadataRequestBioprojects) *NullableV1AssemblyMetadataRequestBioprojects {
	return &NullableV1AssemblyMetadataRequestBioprojects{value: val, isSet: true}
}

func (v NullableV1AssemblyMetadataRequestBioprojects) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AssemblyMetadataRequestBioprojects) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


