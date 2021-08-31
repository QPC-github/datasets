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

// V1reportsProkaryoteGeneLocation struct for V1reportsProkaryoteGeneLocation
type V1reportsProkaryoteGeneLocation struct {
	ProteinAccession *string `json:"protein_accession,omitempty"`
	RefseqGenomicLocation *V1reportsSeqRangeWithAssembly `json:"refseq_genomic_location,omitempty"`
	GenbankGenomicLocation *V1reportsSeqRangeWithAssembly `json:"genbank_genomic_location,omitempty"`
	Organism *V1reportsOrganism `json:"organism,omitempty"`
	Completeness *V1reportsProkaryoteGeneLocationCompleteness `json:"completeness,omitempty"`
	ChromosomeName *string `json:"chromosome_name,omitempty"`
}

// NewV1reportsProkaryoteGeneLocation instantiates a new V1reportsProkaryoteGeneLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsProkaryoteGeneLocation() *V1reportsProkaryoteGeneLocation {
	this := V1reportsProkaryoteGeneLocation{}
	var completeness V1reportsProkaryoteGeneLocationCompleteness = V1REPORTSPROKARYOTEGENELOCATIONCOMPLETENESS_COMPLETE
	this.Completeness = &completeness
	return &this
}

// NewV1reportsProkaryoteGeneLocationWithDefaults instantiates a new V1reportsProkaryoteGeneLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsProkaryoteGeneLocationWithDefaults() *V1reportsProkaryoteGeneLocation {
	this := V1reportsProkaryoteGeneLocation{}
	var completeness V1reportsProkaryoteGeneLocationCompleteness = V1REPORTSPROKARYOTEGENELOCATIONCOMPLETENESS_COMPLETE
	this.Completeness = &completeness
	return &this
}

// GetProteinAccession returns the ProteinAccession field value if set, zero value otherwise.
func (o *V1reportsProkaryoteGeneLocation) GetProteinAccession() string {
	if o == nil || o.ProteinAccession == nil {
		var ret string
		return ret
	}
	return *o.ProteinAccession
}

// GetProteinAccessionOk returns a tuple with the ProteinAccession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsProkaryoteGeneLocation) GetProteinAccessionOk() (*string, bool) {
	if o == nil || o.ProteinAccession == nil {
		return nil, false
	}
	return o.ProteinAccession, true
}

// HasProteinAccession returns a boolean if a field has been set.
func (o *V1reportsProkaryoteGeneLocation) HasProteinAccession() bool {
	if o != nil && o.ProteinAccession != nil {
		return true
	}

	return false
}

// SetProteinAccession gets a reference to the given string and assigns it to the ProteinAccession field.
func (o *V1reportsProkaryoteGeneLocation) SetProteinAccession(v string) {
	o.ProteinAccession = &v
}

// GetRefseqGenomicLocation returns the RefseqGenomicLocation field value if set, zero value otherwise.
func (o *V1reportsProkaryoteGeneLocation) GetRefseqGenomicLocation() V1reportsSeqRangeWithAssembly {
	if o == nil || o.RefseqGenomicLocation == nil {
		var ret V1reportsSeqRangeWithAssembly
		return ret
	}
	return *o.RefseqGenomicLocation
}

// GetRefseqGenomicLocationOk returns a tuple with the RefseqGenomicLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsProkaryoteGeneLocation) GetRefseqGenomicLocationOk() (*V1reportsSeqRangeWithAssembly, bool) {
	if o == nil || o.RefseqGenomicLocation == nil {
		return nil, false
	}
	return o.RefseqGenomicLocation, true
}

// HasRefseqGenomicLocation returns a boolean if a field has been set.
func (o *V1reportsProkaryoteGeneLocation) HasRefseqGenomicLocation() bool {
	if o != nil && o.RefseqGenomicLocation != nil {
		return true
	}

	return false
}

// SetRefseqGenomicLocation gets a reference to the given V1reportsSeqRangeWithAssembly and assigns it to the RefseqGenomicLocation field.
func (o *V1reportsProkaryoteGeneLocation) SetRefseqGenomicLocation(v V1reportsSeqRangeWithAssembly) {
	o.RefseqGenomicLocation = &v
}

// GetGenbankGenomicLocation returns the GenbankGenomicLocation field value if set, zero value otherwise.
func (o *V1reportsProkaryoteGeneLocation) GetGenbankGenomicLocation() V1reportsSeqRangeWithAssembly {
	if o == nil || o.GenbankGenomicLocation == nil {
		var ret V1reportsSeqRangeWithAssembly
		return ret
	}
	return *o.GenbankGenomicLocation
}

// GetGenbankGenomicLocationOk returns a tuple with the GenbankGenomicLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsProkaryoteGeneLocation) GetGenbankGenomicLocationOk() (*V1reportsSeqRangeWithAssembly, bool) {
	if o == nil || o.GenbankGenomicLocation == nil {
		return nil, false
	}
	return o.GenbankGenomicLocation, true
}

// HasGenbankGenomicLocation returns a boolean if a field has been set.
func (o *V1reportsProkaryoteGeneLocation) HasGenbankGenomicLocation() bool {
	if o != nil && o.GenbankGenomicLocation != nil {
		return true
	}

	return false
}

// SetGenbankGenomicLocation gets a reference to the given V1reportsSeqRangeWithAssembly and assigns it to the GenbankGenomicLocation field.
func (o *V1reportsProkaryoteGeneLocation) SetGenbankGenomicLocation(v V1reportsSeqRangeWithAssembly) {
	o.GenbankGenomicLocation = &v
}

// GetOrganism returns the Organism field value if set, zero value otherwise.
func (o *V1reportsProkaryoteGeneLocation) GetOrganism() V1reportsOrganism {
	if o == nil || o.Organism == nil {
		var ret V1reportsOrganism
		return ret
	}
	return *o.Organism
}

// GetOrganismOk returns a tuple with the Organism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsProkaryoteGeneLocation) GetOrganismOk() (*V1reportsOrganism, bool) {
	if o == nil || o.Organism == nil {
		return nil, false
	}
	return o.Organism, true
}

// HasOrganism returns a boolean if a field has been set.
func (o *V1reportsProkaryoteGeneLocation) HasOrganism() bool {
	if o != nil && o.Organism != nil {
		return true
	}

	return false
}

// SetOrganism gets a reference to the given V1reportsOrganism and assigns it to the Organism field.
func (o *V1reportsProkaryoteGeneLocation) SetOrganism(v V1reportsOrganism) {
	o.Organism = &v
}

// GetCompleteness returns the Completeness field value if set, zero value otherwise.
func (o *V1reportsProkaryoteGeneLocation) GetCompleteness() V1reportsProkaryoteGeneLocationCompleteness {
	if o == nil || o.Completeness == nil {
		var ret V1reportsProkaryoteGeneLocationCompleteness
		return ret
	}
	return *o.Completeness
}

// GetCompletenessOk returns a tuple with the Completeness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsProkaryoteGeneLocation) GetCompletenessOk() (*V1reportsProkaryoteGeneLocationCompleteness, bool) {
	if o == nil || o.Completeness == nil {
		return nil, false
	}
	return o.Completeness, true
}

// HasCompleteness returns a boolean if a field has been set.
func (o *V1reportsProkaryoteGeneLocation) HasCompleteness() bool {
	if o != nil && o.Completeness != nil {
		return true
	}

	return false
}

// SetCompleteness gets a reference to the given V1reportsProkaryoteGeneLocationCompleteness and assigns it to the Completeness field.
func (o *V1reportsProkaryoteGeneLocation) SetCompleteness(v V1reportsProkaryoteGeneLocationCompleteness) {
	o.Completeness = &v
}

// GetChromosomeName returns the ChromosomeName field value if set, zero value otherwise.
func (o *V1reportsProkaryoteGeneLocation) GetChromosomeName() string {
	if o == nil || o.ChromosomeName == nil {
		var ret string
		return ret
	}
	return *o.ChromosomeName
}

// GetChromosomeNameOk returns a tuple with the ChromosomeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsProkaryoteGeneLocation) GetChromosomeNameOk() (*string, bool) {
	if o == nil || o.ChromosomeName == nil {
		return nil, false
	}
	return o.ChromosomeName, true
}

// HasChromosomeName returns a boolean if a field has been set.
func (o *V1reportsProkaryoteGeneLocation) HasChromosomeName() bool {
	if o != nil && o.ChromosomeName != nil {
		return true
	}

	return false
}

// SetChromosomeName gets a reference to the given string and assigns it to the ChromosomeName field.
func (o *V1reportsProkaryoteGeneLocation) SetChromosomeName(v string) {
	o.ChromosomeName = &v
}

func (o V1reportsProkaryoteGeneLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProteinAccession != nil  {
		toSerialize["protein_accession"] = o.ProteinAccession
	}
	if o.RefseqGenomicLocation != nil  {
		toSerialize["refseq_genomic_location"] = o.RefseqGenomicLocation
	}
	if o.GenbankGenomicLocation != nil  {
		toSerialize["genbank_genomic_location"] = o.GenbankGenomicLocation
	}
	if o.Organism != nil  {
		toSerialize["organism"] = o.Organism
	}
	if o.Completeness != nil  {
		toSerialize["completeness"] = o.Completeness
	}
	if o.ChromosomeName != nil  {
		toSerialize["chromosome_name"] = o.ChromosomeName
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsProkaryoteGeneLocation struct {
	value *V1reportsProkaryoteGeneLocation
	isSet bool
}

func (v NullableV1reportsProkaryoteGeneLocation) Get() *V1reportsProkaryoteGeneLocation {
	return v.value
}

func (v *NullableV1reportsProkaryoteGeneLocation) Set(val *V1reportsProkaryoteGeneLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsProkaryoteGeneLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsProkaryoteGeneLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsProkaryoteGeneLocation(val *V1reportsProkaryoteGeneLocation) *NullableV1reportsProkaryoteGeneLocation {
	return &NullableV1reportsProkaryoteGeneLocation{value: val, isSet: true}
}

func (v NullableV1reportsProkaryoteGeneLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsProkaryoteGeneLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


