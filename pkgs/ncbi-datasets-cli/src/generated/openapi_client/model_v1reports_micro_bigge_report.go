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

// V1reportsMicroBiggeReport struct for V1reportsMicroBiggeReport
type V1reportsMicroBiggeReport struct {
	TargetAcc *string `json:"target_acc,omitempty"`
	Element *V1reportsElement `json:"element,omitempty"`
	Location *V1reportsSeqRangeSet `json:"location,omitempty"`
	Type *string `json:"type,omitempty"`
	Subtype *string `json:"subtype,omitempty"`
	Class *string `json:"class,omitempty"`
	Subclass *string `json:"subclass,omitempty"`
	AmrMethod *string `json:"amr_method,omitempty"`
	IsPlus *bool `json:"is_plus,omitempty"`
	ClosestReferenceSequenceComparison *V1reportsClosestReference `json:"closest_reference_sequence_comparison,omitempty"`
	Taxonomy *V1reportsTaxonomy `json:"taxonomy,omitempty"`
	Biosample *V1reportsBiosample `json:"biosample,omitempty"`
	ReadToAssemblyCoverage *V1reportsReadToAssemblyCoverage `json:"read_to_assembly_coverage,omitempty"`
	AmrFinderPlus *V1reportsAmrFinderPlus `json:"amr_finder_plus,omitempty"`
	GenesOnContig *[]string `json:"genes_on_contig,omitempty"`
	GenesOnIsolate *[]string `json:"genes_on_isolate,omitempty"`
}

// NewV1reportsMicroBiggeReport instantiates a new V1reportsMicroBiggeReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsMicroBiggeReport() *V1reportsMicroBiggeReport {
	this := V1reportsMicroBiggeReport{}
	return &this
}

// NewV1reportsMicroBiggeReportWithDefaults instantiates a new V1reportsMicroBiggeReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsMicroBiggeReportWithDefaults() *V1reportsMicroBiggeReport {
	this := V1reportsMicroBiggeReport{}
	return &this
}

// GetTargetAcc returns the TargetAcc field value if set, zero value otherwise.
func (o *V1reportsMicroBiggeReport) GetTargetAcc() string {
	if o == nil || o.TargetAcc == nil {
		var ret string
		return ret
	}
	return *o.TargetAcc
}

// GetTargetAccOk returns a tuple with the TargetAcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsMicroBiggeReport) GetTargetAccOk() (*string, bool) {
	if o == nil || o.TargetAcc == nil {
		return nil, false
	}
	return o.TargetAcc, true
}

// HasTargetAcc returns a boolean if a field has been set.
func (o *V1reportsMicroBiggeReport) HasTargetAcc() bool {
	if o != nil && o.TargetAcc != nil {
		return true
	}

	return false
}

// SetTargetAcc gets a reference to the given string and assigns it to the TargetAcc field.
func (o *V1reportsMicroBiggeReport) SetTargetAcc(v string) {
	o.TargetAcc = &v
}

// GetElement returns the Element field value if set, zero value otherwise.
func (o *V1reportsMicroBiggeReport) GetElement() V1reportsElement {
	if o == nil || o.Element == nil {
		var ret V1reportsElement
		return ret
	}
	return *o.Element
}

// GetElementOk returns a tuple with the Element field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsMicroBiggeReport) GetElementOk() (*V1reportsElement, bool) {
	if o == nil || o.Element == nil {
		return nil, false
	}
	return o.Element, true
}

// HasElement returns a boolean if a field has been set.
func (o *V1reportsMicroBiggeReport) HasElement() bool {
	if o != nil && o.Element != nil {
		return true
	}

	return false
}

// SetElement gets a reference to the given V1reportsElement and assigns it to the Element field.
func (o *V1reportsMicroBiggeReport) SetElement(v V1reportsElement) {
	o.Element = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *V1reportsMicroBiggeReport) GetLocation() V1reportsSeqRangeSet {
	if o == nil || o.Location == nil {
		var ret V1reportsSeqRangeSet
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsMicroBiggeReport) GetLocationOk() (*V1reportsSeqRangeSet, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *V1reportsMicroBiggeReport) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given V1reportsSeqRangeSet and assigns it to the Location field.
func (o *V1reportsMicroBiggeReport) SetLocation(v V1reportsSeqRangeSet) {
	o.Location = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1reportsMicroBiggeReport) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsMicroBiggeReport) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1reportsMicroBiggeReport) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V1reportsMicroBiggeReport) SetType(v string) {
	o.Type = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *V1reportsMicroBiggeReport) GetSubtype() string {
	if o == nil || o.Subtype == nil {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsMicroBiggeReport) GetSubtypeOk() (*string, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *V1reportsMicroBiggeReport) HasSubtype() bool {
	if o != nil && o.Subtype != nil {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *V1reportsMicroBiggeReport) SetSubtype(v string) {
	o.Subtype = &v
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *V1reportsMicroBiggeReport) GetClass() string {
	if o == nil || o.Class == nil {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsMicroBiggeReport) GetClassOk() (*string, bool) {
	if o == nil || o.Class == nil {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *V1reportsMicroBiggeReport) HasClass() bool {
	if o != nil && o.Class != nil {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *V1reportsMicroBiggeReport) SetClass(v string) {
	o.Class = &v
}

// GetSubclass returns the Subclass field value if set, zero value otherwise.
func (o *V1reportsMicroBiggeReport) GetSubclass() string {
	if o == nil || o.Subclass == nil {
		var ret string
		return ret
	}
	return *o.Subclass
}

// GetSubclassOk returns a tuple with the Subclass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsMicroBiggeReport) GetSubclassOk() (*string, bool) {
	if o == nil || o.Subclass == nil {
		return nil, false
	}
	return o.Subclass, true
}

// HasSubclass returns a boolean if a field has been set.
func (o *V1reportsMicroBiggeReport) HasSubclass() bool {
	if o != nil && o.Subclass != nil {
		return true
	}

	return false
}

// SetSubclass gets a reference to the given string and assigns it to the Subclass field.
func (o *V1reportsMicroBiggeReport) SetSubclass(v string) {
	o.Subclass = &v
}

// GetAmrMethod returns the AmrMethod field value if set, zero value otherwise.
func (o *V1reportsMicroBiggeReport) GetAmrMethod() string {
	if o == nil || o.AmrMethod == nil {
		var ret string
		return ret
	}
	return *o.AmrMethod
}

// GetAmrMethodOk returns a tuple with the AmrMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsMicroBiggeReport) GetAmrMethodOk() (*string, bool) {
	if o == nil || o.AmrMethod == nil {
		return nil, false
	}
	return o.AmrMethod, true
}

// HasAmrMethod returns a boolean if a field has been set.
func (o *V1reportsMicroBiggeReport) HasAmrMethod() bool {
	if o != nil && o.AmrMethod != nil {
		return true
	}

	return false
}

// SetAmrMethod gets a reference to the given string and assigns it to the AmrMethod field.
func (o *V1reportsMicroBiggeReport) SetAmrMethod(v string) {
	o.AmrMethod = &v
}

// GetIsPlus returns the IsPlus field value if set, zero value otherwise.
func (o *V1reportsMicroBiggeReport) GetIsPlus() bool {
	if o == nil || o.IsPlus == nil {
		var ret bool
		return ret
	}
	return *o.IsPlus
}

// GetIsPlusOk returns a tuple with the IsPlus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsMicroBiggeReport) GetIsPlusOk() (*bool, bool) {
	if o == nil || o.IsPlus == nil {
		return nil, false
	}
	return o.IsPlus, true
}

// HasIsPlus returns a boolean if a field has been set.
func (o *V1reportsMicroBiggeReport) HasIsPlus() bool {
	if o != nil && o.IsPlus != nil {
		return true
	}

	return false
}

// SetIsPlus gets a reference to the given bool and assigns it to the IsPlus field.
func (o *V1reportsMicroBiggeReport) SetIsPlus(v bool) {
	o.IsPlus = &v
}

// GetClosestReferenceSequenceComparison returns the ClosestReferenceSequenceComparison field value if set, zero value otherwise.
func (o *V1reportsMicroBiggeReport) GetClosestReferenceSequenceComparison() V1reportsClosestReference {
	if o == nil || o.ClosestReferenceSequenceComparison == nil {
		var ret V1reportsClosestReference
		return ret
	}
	return *o.ClosestReferenceSequenceComparison
}

// GetClosestReferenceSequenceComparisonOk returns a tuple with the ClosestReferenceSequenceComparison field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsMicroBiggeReport) GetClosestReferenceSequenceComparisonOk() (*V1reportsClosestReference, bool) {
	if o == nil || o.ClosestReferenceSequenceComparison == nil {
		return nil, false
	}
	return o.ClosestReferenceSequenceComparison, true
}

// HasClosestReferenceSequenceComparison returns a boolean if a field has been set.
func (o *V1reportsMicroBiggeReport) HasClosestReferenceSequenceComparison() bool {
	if o != nil && o.ClosestReferenceSequenceComparison != nil {
		return true
	}

	return false
}

// SetClosestReferenceSequenceComparison gets a reference to the given V1reportsClosestReference and assigns it to the ClosestReferenceSequenceComparison field.
func (o *V1reportsMicroBiggeReport) SetClosestReferenceSequenceComparison(v V1reportsClosestReference) {
	o.ClosestReferenceSequenceComparison = &v
}

// GetTaxonomy returns the Taxonomy field value if set, zero value otherwise.
func (o *V1reportsMicroBiggeReport) GetTaxonomy() V1reportsTaxonomy {
	if o == nil || o.Taxonomy == nil {
		var ret V1reportsTaxonomy
		return ret
	}
	return *o.Taxonomy
}

// GetTaxonomyOk returns a tuple with the Taxonomy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsMicroBiggeReport) GetTaxonomyOk() (*V1reportsTaxonomy, bool) {
	if o == nil || o.Taxonomy == nil {
		return nil, false
	}
	return o.Taxonomy, true
}

// HasTaxonomy returns a boolean if a field has been set.
func (o *V1reportsMicroBiggeReport) HasTaxonomy() bool {
	if o != nil && o.Taxonomy != nil {
		return true
	}

	return false
}

// SetTaxonomy gets a reference to the given V1reportsTaxonomy and assigns it to the Taxonomy field.
func (o *V1reportsMicroBiggeReport) SetTaxonomy(v V1reportsTaxonomy) {
	o.Taxonomy = &v
}

// GetBiosample returns the Biosample field value if set, zero value otherwise.
func (o *V1reportsMicroBiggeReport) GetBiosample() V1reportsBiosample {
	if o == nil || o.Biosample == nil {
		var ret V1reportsBiosample
		return ret
	}
	return *o.Biosample
}

// GetBiosampleOk returns a tuple with the Biosample field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsMicroBiggeReport) GetBiosampleOk() (*V1reportsBiosample, bool) {
	if o == nil || o.Biosample == nil {
		return nil, false
	}
	return o.Biosample, true
}

// HasBiosample returns a boolean if a field has been set.
func (o *V1reportsMicroBiggeReport) HasBiosample() bool {
	if o != nil && o.Biosample != nil {
		return true
	}

	return false
}

// SetBiosample gets a reference to the given V1reportsBiosample and assigns it to the Biosample field.
func (o *V1reportsMicroBiggeReport) SetBiosample(v V1reportsBiosample) {
	o.Biosample = &v
}

// GetReadToAssemblyCoverage returns the ReadToAssemblyCoverage field value if set, zero value otherwise.
func (o *V1reportsMicroBiggeReport) GetReadToAssemblyCoverage() V1reportsReadToAssemblyCoverage {
	if o == nil || o.ReadToAssemblyCoverage == nil {
		var ret V1reportsReadToAssemblyCoverage
		return ret
	}
	return *o.ReadToAssemblyCoverage
}

// GetReadToAssemblyCoverageOk returns a tuple with the ReadToAssemblyCoverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsMicroBiggeReport) GetReadToAssemblyCoverageOk() (*V1reportsReadToAssemblyCoverage, bool) {
	if o == nil || o.ReadToAssemblyCoverage == nil {
		return nil, false
	}
	return o.ReadToAssemblyCoverage, true
}

// HasReadToAssemblyCoverage returns a boolean if a field has been set.
func (o *V1reportsMicroBiggeReport) HasReadToAssemblyCoverage() bool {
	if o != nil && o.ReadToAssemblyCoverage != nil {
		return true
	}

	return false
}

// SetReadToAssemblyCoverage gets a reference to the given V1reportsReadToAssemblyCoverage and assigns it to the ReadToAssemblyCoverage field.
func (o *V1reportsMicroBiggeReport) SetReadToAssemblyCoverage(v V1reportsReadToAssemblyCoverage) {
	o.ReadToAssemblyCoverage = &v
}

// GetAmrFinderPlus returns the AmrFinderPlus field value if set, zero value otherwise.
func (o *V1reportsMicroBiggeReport) GetAmrFinderPlus() V1reportsAmrFinderPlus {
	if o == nil || o.AmrFinderPlus == nil {
		var ret V1reportsAmrFinderPlus
		return ret
	}
	return *o.AmrFinderPlus
}

// GetAmrFinderPlusOk returns a tuple with the AmrFinderPlus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsMicroBiggeReport) GetAmrFinderPlusOk() (*V1reportsAmrFinderPlus, bool) {
	if o == nil || o.AmrFinderPlus == nil {
		return nil, false
	}
	return o.AmrFinderPlus, true
}

// HasAmrFinderPlus returns a boolean if a field has been set.
func (o *V1reportsMicroBiggeReport) HasAmrFinderPlus() bool {
	if o != nil && o.AmrFinderPlus != nil {
		return true
	}

	return false
}

// SetAmrFinderPlus gets a reference to the given V1reportsAmrFinderPlus and assigns it to the AmrFinderPlus field.
func (o *V1reportsMicroBiggeReport) SetAmrFinderPlus(v V1reportsAmrFinderPlus) {
	o.AmrFinderPlus = &v
}

// GetGenesOnContig returns the GenesOnContig field value if set, zero value otherwise.
func (o *V1reportsMicroBiggeReport) GetGenesOnContig() []string {
	if o == nil || o.GenesOnContig == nil {
		var ret []string
		return ret
	}
	return *o.GenesOnContig
}

// GetGenesOnContigOk returns a tuple with the GenesOnContig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsMicroBiggeReport) GetGenesOnContigOk() (*[]string, bool) {
	if o == nil || o.GenesOnContig == nil {
		return nil, false
	}
	return o.GenesOnContig, true
}

// HasGenesOnContig returns a boolean if a field has been set.
func (o *V1reportsMicroBiggeReport) HasGenesOnContig() bool {
	if o != nil && o.GenesOnContig != nil {
		return true
	}

	return false
}

// SetGenesOnContig gets a reference to the given []string and assigns it to the GenesOnContig field.
func (o *V1reportsMicroBiggeReport) SetGenesOnContig(v []string) {
	o.GenesOnContig = &v
}

// GetGenesOnIsolate returns the GenesOnIsolate field value if set, zero value otherwise.
func (o *V1reportsMicroBiggeReport) GetGenesOnIsolate() []string {
	if o == nil || o.GenesOnIsolate == nil {
		var ret []string
		return ret
	}
	return *o.GenesOnIsolate
}

// GetGenesOnIsolateOk returns a tuple with the GenesOnIsolate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsMicroBiggeReport) GetGenesOnIsolateOk() (*[]string, bool) {
	if o == nil || o.GenesOnIsolate == nil {
		return nil, false
	}
	return o.GenesOnIsolate, true
}

// HasGenesOnIsolate returns a boolean if a field has been set.
func (o *V1reportsMicroBiggeReport) HasGenesOnIsolate() bool {
	if o != nil && o.GenesOnIsolate != nil {
		return true
	}

	return false
}

// SetGenesOnIsolate gets a reference to the given []string and assigns it to the GenesOnIsolate field.
func (o *V1reportsMicroBiggeReport) SetGenesOnIsolate(v []string) {
	o.GenesOnIsolate = &v
}

func (o V1reportsMicroBiggeReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TargetAcc != nil  {
		toSerialize["target_acc"] = o.TargetAcc
	}
	if o.Element != nil  {
		toSerialize["element"] = o.Element
	}
	if o.Location != nil  {
		toSerialize["location"] = o.Location
	}
	if o.Type != nil  {
		toSerialize["type"] = o.Type
	}
	if o.Subtype != nil  {
		toSerialize["subtype"] = o.Subtype
	}
	if o.Class != nil  {
		toSerialize["class"] = o.Class
	}
	if o.Subclass != nil  {
		toSerialize["subclass"] = o.Subclass
	}
	if o.AmrMethod != nil  {
		toSerialize["amr_method"] = o.AmrMethod
	}
	if o.IsPlus != nil  {
		toSerialize["is_plus"] = o.IsPlus
	}
	if o.ClosestReferenceSequenceComparison != nil  {
		toSerialize["closest_reference_sequence_comparison"] = o.ClosestReferenceSequenceComparison
	}
	if o.Taxonomy != nil  {
		toSerialize["taxonomy"] = o.Taxonomy
	}
	if o.Biosample != nil  {
		toSerialize["biosample"] = o.Biosample
	}
	if o.ReadToAssemblyCoverage != nil  {
		toSerialize["read_to_assembly_coverage"] = o.ReadToAssemblyCoverage
	}
	if o.AmrFinderPlus != nil  {
		toSerialize["amr_finder_plus"] = o.AmrFinderPlus
	}
	if o.GenesOnContig != nil && len(o.GetGenesOnContig()) > 0  {
		toSerialize["genes_on_contig"] = o.GenesOnContig
	}
	if o.GenesOnIsolate != nil && len(o.GetGenesOnIsolate()) > 0  {
		toSerialize["genes_on_isolate"] = o.GenesOnIsolate
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsMicroBiggeReport struct {
	value *V1reportsMicroBiggeReport
	isSet bool
}

func (v NullableV1reportsMicroBiggeReport) Get() *V1reportsMicroBiggeReport {
	return v.value
}

func (v *NullableV1reportsMicroBiggeReport) Set(val *V1reportsMicroBiggeReport) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsMicroBiggeReport) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsMicroBiggeReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsMicroBiggeReport(val *V1reportsMicroBiggeReport) *NullableV1reportsMicroBiggeReport {
	return &NullableV1reportsMicroBiggeReport{value: val, isSet: true}
}

func (v NullableV1reportsMicroBiggeReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsMicroBiggeReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


