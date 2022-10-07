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

// V1AssemblyDatasetRequest struct for V1AssemblyDatasetRequest
type V1AssemblyDatasetRequest struct {
	AssemblyAccessions *[]string `json:"assembly_accessions,omitempty"`
	Accessions *[]string `json:"accessions,omitempty"`
	Chromosomes *[]string `json:"chromosomes,omitempty"`
	IncludeAnnotation *bool `json:"include_annotation,omitempty"`
	ExcludeSequence *bool `json:"exclude_sequence,omitempty"`
	IncludeAnnotationType *[]V1AnnotationForAssemblyType `json:"include_annotation_type,omitempty"`
	Hydrated *V1AssemblyDatasetRequestResolution `json:"hydrated,omitempty"`
	IncludeTsv *bool `json:"include_tsv,omitempty"`
}

// NewV1AssemblyDatasetRequest instantiates a new V1AssemblyDatasetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AssemblyDatasetRequest() *V1AssemblyDatasetRequest {
	this := V1AssemblyDatasetRequest{}
	var hydrated V1AssemblyDatasetRequestResolution = V1ASSEMBLYDATASETREQUESTRESOLUTION_FULLY_HYDRATED
	this.Hydrated = &hydrated
	return &this
}

// NewV1AssemblyDatasetRequestWithDefaults instantiates a new V1AssemblyDatasetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AssemblyDatasetRequestWithDefaults() *V1AssemblyDatasetRequest {
	this := V1AssemblyDatasetRequest{}
	var hydrated V1AssemblyDatasetRequestResolution = V1ASSEMBLYDATASETREQUESTRESOLUTION_FULLY_HYDRATED
	this.Hydrated = &hydrated
	return &this
}

// GetAssemblyAccessions returns the AssemblyAccessions field value if set, zero value otherwise.
func (o *V1AssemblyDatasetRequest) GetAssemblyAccessions() []string {
	if o == nil || o.AssemblyAccessions == nil {
		var ret []string
		return ret
	}
	return *o.AssemblyAccessions
}

// GetAssemblyAccessionsOk returns a tuple with the AssemblyAccessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetRequest) GetAssemblyAccessionsOk() (*[]string, bool) {
	if o == nil || o.AssemblyAccessions == nil {
		return nil, false
	}
	return o.AssemblyAccessions, true
}

// HasAssemblyAccessions returns a boolean if a field has been set.
func (o *V1AssemblyDatasetRequest) HasAssemblyAccessions() bool {
	if o != nil && o.AssemblyAccessions != nil {
		return true
	}

	return false
}

// SetAssemblyAccessions gets a reference to the given []string and assigns it to the AssemblyAccessions field.
func (o *V1AssemblyDatasetRequest) SetAssemblyAccessions(v []string) {
	o.AssemblyAccessions = &v
}

// GetAccessions returns the Accessions field value if set, zero value otherwise.
func (o *V1AssemblyDatasetRequest) GetAccessions() []string {
	if o == nil || o.Accessions == nil {
		var ret []string
		return ret
	}
	return *o.Accessions
}

// GetAccessionsOk returns a tuple with the Accessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetRequest) GetAccessionsOk() (*[]string, bool) {
	if o == nil || o.Accessions == nil {
		return nil, false
	}
	return o.Accessions, true
}

// HasAccessions returns a boolean if a field has been set.
func (o *V1AssemblyDatasetRequest) HasAccessions() bool {
	if o != nil && o.Accessions != nil {
		return true
	}

	return false
}

// SetAccessions gets a reference to the given []string and assigns it to the Accessions field.
func (o *V1AssemblyDatasetRequest) SetAccessions(v []string) {
	o.Accessions = &v
}

// GetChromosomes returns the Chromosomes field value if set, zero value otherwise.
func (o *V1AssemblyDatasetRequest) GetChromosomes() []string {
	if o == nil || o.Chromosomes == nil {
		var ret []string
		return ret
	}
	return *o.Chromosomes
}

// GetChromosomesOk returns a tuple with the Chromosomes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetRequest) GetChromosomesOk() (*[]string, bool) {
	if o == nil || o.Chromosomes == nil {
		return nil, false
	}
	return o.Chromosomes, true
}

// HasChromosomes returns a boolean if a field has been set.
func (o *V1AssemblyDatasetRequest) HasChromosomes() bool {
	if o != nil && o.Chromosomes != nil {
		return true
	}

	return false
}

// SetChromosomes gets a reference to the given []string and assigns it to the Chromosomes field.
func (o *V1AssemblyDatasetRequest) SetChromosomes(v []string) {
	o.Chromosomes = &v
}

// GetIncludeAnnotation returns the IncludeAnnotation field value if set, zero value otherwise.
func (o *V1AssemblyDatasetRequest) GetIncludeAnnotation() bool {
	if o == nil || o.IncludeAnnotation == nil {
		var ret bool
		return ret
	}
	return *o.IncludeAnnotation
}

// GetIncludeAnnotationOk returns a tuple with the IncludeAnnotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetRequest) GetIncludeAnnotationOk() (*bool, bool) {
	if o == nil || o.IncludeAnnotation == nil {
		return nil, false
	}
	return o.IncludeAnnotation, true
}

// HasIncludeAnnotation returns a boolean if a field has been set.
func (o *V1AssemblyDatasetRequest) HasIncludeAnnotation() bool {
	if o != nil && o.IncludeAnnotation != nil {
		return true
	}

	return false
}

// SetIncludeAnnotation gets a reference to the given bool and assigns it to the IncludeAnnotation field.
func (o *V1AssemblyDatasetRequest) SetIncludeAnnotation(v bool) {
	o.IncludeAnnotation = &v
}

// GetExcludeSequence returns the ExcludeSequence field value if set, zero value otherwise.
func (o *V1AssemblyDatasetRequest) GetExcludeSequence() bool {
	if o == nil || o.ExcludeSequence == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeSequence
}

// GetExcludeSequenceOk returns a tuple with the ExcludeSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetRequest) GetExcludeSequenceOk() (*bool, bool) {
	if o == nil || o.ExcludeSequence == nil {
		return nil, false
	}
	return o.ExcludeSequence, true
}

// HasExcludeSequence returns a boolean if a field has been set.
func (o *V1AssemblyDatasetRequest) HasExcludeSequence() bool {
	if o != nil && o.ExcludeSequence != nil {
		return true
	}

	return false
}

// SetExcludeSequence gets a reference to the given bool and assigns it to the ExcludeSequence field.
func (o *V1AssemblyDatasetRequest) SetExcludeSequence(v bool) {
	o.ExcludeSequence = &v
}

// GetIncludeAnnotationType returns the IncludeAnnotationType field value if set, zero value otherwise.
func (o *V1AssemblyDatasetRequest) GetIncludeAnnotationType() []V1AnnotationForAssemblyType {
	if o == nil || o.IncludeAnnotationType == nil {
		var ret []V1AnnotationForAssemblyType
		return ret
	}
	return *o.IncludeAnnotationType
}

// GetIncludeAnnotationTypeOk returns a tuple with the IncludeAnnotationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetRequest) GetIncludeAnnotationTypeOk() (*[]V1AnnotationForAssemblyType, bool) {
	if o == nil || o.IncludeAnnotationType == nil {
		return nil, false
	}
	return o.IncludeAnnotationType, true
}

// HasIncludeAnnotationType returns a boolean if a field has been set.
func (o *V1AssemblyDatasetRequest) HasIncludeAnnotationType() bool {
	if o != nil && o.IncludeAnnotationType != nil {
		return true
	}

	return false
}

// SetIncludeAnnotationType gets a reference to the given []V1AnnotationForAssemblyType and assigns it to the IncludeAnnotationType field.
func (o *V1AssemblyDatasetRequest) SetIncludeAnnotationType(v []V1AnnotationForAssemblyType) {
	o.IncludeAnnotationType = &v
}

// GetHydrated returns the Hydrated field value if set, zero value otherwise.
func (o *V1AssemblyDatasetRequest) GetHydrated() V1AssemblyDatasetRequestResolution {
	if o == nil || o.Hydrated == nil {
		var ret V1AssemblyDatasetRequestResolution
		return ret
	}
	return *o.Hydrated
}

// GetHydratedOk returns a tuple with the Hydrated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetRequest) GetHydratedOk() (*V1AssemblyDatasetRequestResolution, bool) {
	if o == nil || o.Hydrated == nil {
		return nil, false
	}
	return o.Hydrated, true
}

// HasHydrated returns a boolean if a field has been set.
func (o *V1AssemblyDatasetRequest) HasHydrated() bool {
	if o != nil && o.Hydrated != nil {
		return true
	}

	return false
}

// SetHydrated gets a reference to the given V1AssemblyDatasetRequestResolution and assigns it to the Hydrated field.
func (o *V1AssemblyDatasetRequest) SetHydrated(v V1AssemblyDatasetRequestResolution) {
	o.Hydrated = &v
}

// GetIncludeTsv returns the IncludeTsv field value if set, zero value otherwise.
func (o *V1AssemblyDatasetRequest) GetIncludeTsv() bool {
	if o == nil || o.IncludeTsv == nil {
		var ret bool
		return ret
	}
	return *o.IncludeTsv
}

// GetIncludeTsvOk returns a tuple with the IncludeTsv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetRequest) GetIncludeTsvOk() (*bool, bool) {
	if o == nil || o.IncludeTsv == nil {
		return nil, false
	}
	return o.IncludeTsv, true
}

// HasIncludeTsv returns a boolean if a field has been set.
func (o *V1AssemblyDatasetRequest) HasIncludeTsv() bool {
	if o != nil && o.IncludeTsv != nil {
		return true
	}

	return false
}

// SetIncludeTsv gets a reference to the given bool and assigns it to the IncludeTsv field.
func (o *V1AssemblyDatasetRequest) SetIncludeTsv(v bool) {
	o.IncludeTsv = &v
}

func (o V1AssemblyDatasetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssemblyAccessions != nil && len(o.GetAssemblyAccessions()) > 0  {
		toSerialize["assembly_accessions"] = o.AssemblyAccessions
	}
	if o.Accessions != nil && len(o.GetAccessions()) > 0  {
		toSerialize["accessions"] = o.Accessions
	}
	if o.Chromosomes != nil && len(o.GetChromosomes()) > 0  {
		toSerialize["chromosomes"] = o.Chromosomes
	}
	if o.IncludeAnnotation != nil  {
		toSerialize["include_annotation"] = o.IncludeAnnotation
	}
	if o.ExcludeSequence != nil  {
		toSerialize["exclude_sequence"] = o.ExcludeSequence
	}
	if o.IncludeAnnotationType != nil && len(o.GetIncludeAnnotationType()) > 0  {
		toSerialize["include_annotation_type"] = o.IncludeAnnotationType
	}
	if o.Hydrated != nil  {
		toSerialize["hydrated"] = o.Hydrated
	}
	if o.IncludeTsv != nil  {
		toSerialize["include_tsv"] = o.IncludeTsv
	}
	return json.Marshal(toSerialize)
}

type NullableV1AssemblyDatasetRequest struct {
	value *V1AssemblyDatasetRequest
	isSet bool
}

func (v NullableV1AssemblyDatasetRequest) Get() *V1AssemblyDatasetRequest {
	return v.value
}

func (v *NullableV1AssemblyDatasetRequest) Set(val *V1AssemblyDatasetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AssemblyDatasetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AssemblyDatasetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AssemblyDatasetRequest(val *V1AssemblyDatasetRequest) *NullableV1AssemblyDatasetRequest {
	return &NullableV1AssemblyDatasetRequest{value: val, isSet: true}
}

func (v NullableV1AssemblyDatasetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AssemblyDatasetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


