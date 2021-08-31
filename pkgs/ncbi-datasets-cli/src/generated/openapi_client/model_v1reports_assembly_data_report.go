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

// V1reportsAssemblyDataReport struct for V1reportsAssemblyDataReport
type V1reportsAssemblyDataReport struct {
	CommonName *string `json:"common_name,omitempty"`
	OrganismName *string `json:"organism_name,omitempty"`
	Breed *string `json:"breed,omitempty"`
	Cultivar *string `json:"cultivar,omitempty"`
	Ecotype *string `json:"ecotype,omitempty"`
	Isolate *string `json:"isolate,omitempty"`
	Sex *string `json:"sex,omitempty"`
	Strain *string `json:"strain,omitempty"`
	TaxId *int32 `json:"tax_id,omitempty"`
	AssemblyInfo *V1reportsAssemblyInfo `json:"assembly_info,omitempty"`
	AssemblyStats *V1reportsAssemblyStats `json:"assembly_stats,omitempty"`
	OrganelleInfo *[]V1reportsOrganelleInfo `json:"organelle_info,omitempty"`
	AnnotationInfo *V1reportsAnnotationInfo `json:"annotation_info,omitempty"`
	WgsInfo *V1reportsWGSInfo `json:"wgs_info,omitempty"`
}

// NewV1reportsAssemblyDataReport instantiates a new V1reportsAssemblyDataReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsAssemblyDataReport() *V1reportsAssemblyDataReport {
	this := V1reportsAssemblyDataReport{}
	return &this
}

// NewV1reportsAssemblyDataReportWithDefaults instantiates a new V1reportsAssemblyDataReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsAssemblyDataReportWithDefaults() *V1reportsAssemblyDataReport {
	this := V1reportsAssemblyDataReport{}
	return &this
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *V1reportsAssemblyDataReport) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsAssemblyDataReport) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *V1reportsAssemblyDataReport) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *V1reportsAssemblyDataReport) SetCommonName(v string) {
	o.CommonName = &v
}

// GetOrganismName returns the OrganismName field value if set, zero value otherwise.
func (o *V1reportsAssemblyDataReport) GetOrganismName() string {
	if o == nil || o.OrganismName == nil {
		var ret string
		return ret
	}
	return *o.OrganismName
}

// GetOrganismNameOk returns a tuple with the OrganismName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsAssemblyDataReport) GetOrganismNameOk() (*string, bool) {
	if o == nil || o.OrganismName == nil {
		return nil, false
	}
	return o.OrganismName, true
}

// HasOrganismName returns a boolean if a field has been set.
func (o *V1reportsAssemblyDataReport) HasOrganismName() bool {
	if o != nil && o.OrganismName != nil {
		return true
	}

	return false
}

// SetOrganismName gets a reference to the given string and assigns it to the OrganismName field.
func (o *V1reportsAssemblyDataReport) SetOrganismName(v string) {
	o.OrganismName = &v
}

// GetBreed returns the Breed field value if set, zero value otherwise.
func (o *V1reportsAssemblyDataReport) GetBreed() string {
	if o == nil || o.Breed == nil {
		var ret string
		return ret
	}
	return *o.Breed
}

// GetBreedOk returns a tuple with the Breed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsAssemblyDataReport) GetBreedOk() (*string, bool) {
	if o == nil || o.Breed == nil {
		return nil, false
	}
	return o.Breed, true
}

// HasBreed returns a boolean if a field has been set.
func (o *V1reportsAssemblyDataReport) HasBreed() bool {
	if o != nil && o.Breed != nil {
		return true
	}

	return false
}

// SetBreed gets a reference to the given string and assigns it to the Breed field.
func (o *V1reportsAssemblyDataReport) SetBreed(v string) {
	o.Breed = &v
}

// GetCultivar returns the Cultivar field value if set, zero value otherwise.
func (o *V1reportsAssemblyDataReport) GetCultivar() string {
	if o == nil || o.Cultivar == nil {
		var ret string
		return ret
	}
	return *o.Cultivar
}

// GetCultivarOk returns a tuple with the Cultivar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsAssemblyDataReport) GetCultivarOk() (*string, bool) {
	if o == nil || o.Cultivar == nil {
		return nil, false
	}
	return o.Cultivar, true
}

// HasCultivar returns a boolean if a field has been set.
func (o *V1reportsAssemblyDataReport) HasCultivar() bool {
	if o != nil && o.Cultivar != nil {
		return true
	}

	return false
}

// SetCultivar gets a reference to the given string and assigns it to the Cultivar field.
func (o *V1reportsAssemblyDataReport) SetCultivar(v string) {
	o.Cultivar = &v
}

// GetEcotype returns the Ecotype field value if set, zero value otherwise.
func (o *V1reportsAssemblyDataReport) GetEcotype() string {
	if o == nil || o.Ecotype == nil {
		var ret string
		return ret
	}
	return *o.Ecotype
}

// GetEcotypeOk returns a tuple with the Ecotype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsAssemblyDataReport) GetEcotypeOk() (*string, bool) {
	if o == nil || o.Ecotype == nil {
		return nil, false
	}
	return o.Ecotype, true
}

// HasEcotype returns a boolean if a field has been set.
func (o *V1reportsAssemblyDataReport) HasEcotype() bool {
	if o != nil && o.Ecotype != nil {
		return true
	}

	return false
}

// SetEcotype gets a reference to the given string and assigns it to the Ecotype field.
func (o *V1reportsAssemblyDataReport) SetEcotype(v string) {
	o.Ecotype = &v
}

// GetIsolate returns the Isolate field value if set, zero value otherwise.
func (o *V1reportsAssemblyDataReport) GetIsolate() string {
	if o == nil || o.Isolate == nil {
		var ret string
		return ret
	}
	return *o.Isolate
}

// GetIsolateOk returns a tuple with the Isolate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsAssemblyDataReport) GetIsolateOk() (*string, bool) {
	if o == nil || o.Isolate == nil {
		return nil, false
	}
	return o.Isolate, true
}

// HasIsolate returns a boolean if a field has been set.
func (o *V1reportsAssemblyDataReport) HasIsolate() bool {
	if o != nil && o.Isolate != nil {
		return true
	}

	return false
}

// SetIsolate gets a reference to the given string and assigns it to the Isolate field.
func (o *V1reportsAssemblyDataReport) SetIsolate(v string) {
	o.Isolate = &v
}

// GetSex returns the Sex field value if set, zero value otherwise.
func (o *V1reportsAssemblyDataReport) GetSex() string {
	if o == nil || o.Sex == nil {
		var ret string
		return ret
	}
	return *o.Sex
}

// GetSexOk returns a tuple with the Sex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsAssemblyDataReport) GetSexOk() (*string, bool) {
	if o == nil || o.Sex == nil {
		return nil, false
	}
	return o.Sex, true
}

// HasSex returns a boolean if a field has been set.
func (o *V1reportsAssemblyDataReport) HasSex() bool {
	if o != nil && o.Sex != nil {
		return true
	}

	return false
}

// SetSex gets a reference to the given string and assigns it to the Sex field.
func (o *V1reportsAssemblyDataReport) SetSex(v string) {
	o.Sex = &v
}

// GetStrain returns the Strain field value if set, zero value otherwise.
func (o *V1reportsAssemblyDataReport) GetStrain() string {
	if o == nil || o.Strain == nil {
		var ret string
		return ret
	}
	return *o.Strain
}

// GetStrainOk returns a tuple with the Strain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsAssemblyDataReport) GetStrainOk() (*string, bool) {
	if o == nil || o.Strain == nil {
		return nil, false
	}
	return o.Strain, true
}

// HasStrain returns a boolean if a field has been set.
func (o *V1reportsAssemblyDataReport) HasStrain() bool {
	if o != nil && o.Strain != nil {
		return true
	}

	return false
}

// SetStrain gets a reference to the given string and assigns it to the Strain field.
func (o *V1reportsAssemblyDataReport) SetStrain(v string) {
	o.Strain = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *V1reportsAssemblyDataReport) GetTaxId() int32 {
	if o == nil || o.TaxId == nil {
		var ret int32
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsAssemblyDataReport) GetTaxIdOk() (*int32, bool) {
	if o == nil || o.TaxId == nil {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *V1reportsAssemblyDataReport) HasTaxId() bool {
	if o != nil && o.TaxId != nil {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given int32 and assigns it to the TaxId field.
func (o *V1reportsAssemblyDataReport) SetTaxId(v int32) {
	o.TaxId = &v
}

// GetAssemblyInfo returns the AssemblyInfo field value if set, zero value otherwise.
func (o *V1reportsAssemblyDataReport) GetAssemblyInfo() V1reportsAssemblyInfo {
	if o == nil || o.AssemblyInfo == nil {
		var ret V1reportsAssemblyInfo
		return ret
	}
	return *o.AssemblyInfo
}

// GetAssemblyInfoOk returns a tuple with the AssemblyInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsAssemblyDataReport) GetAssemblyInfoOk() (*V1reportsAssemblyInfo, bool) {
	if o == nil || o.AssemblyInfo == nil {
		return nil, false
	}
	return o.AssemblyInfo, true
}

// HasAssemblyInfo returns a boolean if a field has been set.
func (o *V1reportsAssemblyDataReport) HasAssemblyInfo() bool {
	if o != nil && o.AssemblyInfo != nil {
		return true
	}

	return false
}

// SetAssemblyInfo gets a reference to the given V1reportsAssemblyInfo and assigns it to the AssemblyInfo field.
func (o *V1reportsAssemblyDataReport) SetAssemblyInfo(v V1reportsAssemblyInfo) {
	o.AssemblyInfo = &v
}

// GetAssemblyStats returns the AssemblyStats field value if set, zero value otherwise.
func (o *V1reportsAssemblyDataReport) GetAssemblyStats() V1reportsAssemblyStats {
	if o == nil || o.AssemblyStats == nil {
		var ret V1reportsAssemblyStats
		return ret
	}
	return *o.AssemblyStats
}

// GetAssemblyStatsOk returns a tuple with the AssemblyStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsAssemblyDataReport) GetAssemblyStatsOk() (*V1reportsAssemblyStats, bool) {
	if o == nil || o.AssemblyStats == nil {
		return nil, false
	}
	return o.AssemblyStats, true
}

// HasAssemblyStats returns a boolean if a field has been set.
func (o *V1reportsAssemblyDataReport) HasAssemblyStats() bool {
	if o != nil && o.AssemblyStats != nil {
		return true
	}

	return false
}

// SetAssemblyStats gets a reference to the given V1reportsAssemblyStats and assigns it to the AssemblyStats field.
func (o *V1reportsAssemblyDataReport) SetAssemblyStats(v V1reportsAssemblyStats) {
	o.AssemblyStats = &v
}

// GetOrganelleInfo returns the OrganelleInfo field value if set, zero value otherwise.
func (o *V1reportsAssemblyDataReport) GetOrganelleInfo() []V1reportsOrganelleInfo {
	if o == nil || o.OrganelleInfo == nil {
		var ret []V1reportsOrganelleInfo
		return ret
	}
	return *o.OrganelleInfo
}

// GetOrganelleInfoOk returns a tuple with the OrganelleInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsAssemblyDataReport) GetOrganelleInfoOk() (*[]V1reportsOrganelleInfo, bool) {
	if o == nil || o.OrganelleInfo == nil {
		return nil, false
	}
	return o.OrganelleInfo, true
}

// HasOrganelleInfo returns a boolean if a field has been set.
func (o *V1reportsAssemblyDataReport) HasOrganelleInfo() bool {
	if o != nil && o.OrganelleInfo != nil {
		return true
	}

	return false
}

// SetOrganelleInfo gets a reference to the given []V1reportsOrganelleInfo and assigns it to the OrganelleInfo field.
func (o *V1reportsAssemblyDataReport) SetOrganelleInfo(v []V1reportsOrganelleInfo) {
	o.OrganelleInfo = &v
}

// GetAnnotationInfo returns the AnnotationInfo field value if set, zero value otherwise.
func (o *V1reportsAssemblyDataReport) GetAnnotationInfo() V1reportsAnnotationInfo {
	if o == nil || o.AnnotationInfo == nil {
		var ret V1reportsAnnotationInfo
		return ret
	}
	return *o.AnnotationInfo
}

// GetAnnotationInfoOk returns a tuple with the AnnotationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsAssemblyDataReport) GetAnnotationInfoOk() (*V1reportsAnnotationInfo, bool) {
	if o == nil || o.AnnotationInfo == nil {
		return nil, false
	}
	return o.AnnotationInfo, true
}

// HasAnnotationInfo returns a boolean if a field has been set.
func (o *V1reportsAssemblyDataReport) HasAnnotationInfo() bool {
	if o != nil && o.AnnotationInfo != nil {
		return true
	}

	return false
}

// SetAnnotationInfo gets a reference to the given V1reportsAnnotationInfo and assigns it to the AnnotationInfo field.
func (o *V1reportsAssemblyDataReport) SetAnnotationInfo(v V1reportsAnnotationInfo) {
	o.AnnotationInfo = &v
}

// GetWgsInfo returns the WgsInfo field value if set, zero value otherwise.
func (o *V1reportsAssemblyDataReport) GetWgsInfo() V1reportsWGSInfo {
	if o == nil || o.WgsInfo == nil {
		var ret V1reportsWGSInfo
		return ret
	}
	return *o.WgsInfo
}

// GetWgsInfoOk returns a tuple with the WgsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsAssemblyDataReport) GetWgsInfoOk() (*V1reportsWGSInfo, bool) {
	if o == nil || o.WgsInfo == nil {
		return nil, false
	}
	return o.WgsInfo, true
}

// HasWgsInfo returns a boolean if a field has been set.
func (o *V1reportsAssemblyDataReport) HasWgsInfo() bool {
	if o != nil && o.WgsInfo != nil {
		return true
	}

	return false
}

// SetWgsInfo gets a reference to the given V1reportsWGSInfo and assigns it to the WgsInfo field.
func (o *V1reportsAssemblyDataReport) SetWgsInfo(v V1reportsWGSInfo) {
	o.WgsInfo = &v
}

func (o V1reportsAssemblyDataReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CommonName != nil  {
		toSerialize["common_name"] = o.CommonName
	}
	if o.OrganismName != nil  {
		toSerialize["organism_name"] = o.OrganismName
	}
	if o.Breed != nil  {
		toSerialize["breed"] = o.Breed
	}
	if o.Cultivar != nil  {
		toSerialize["cultivar"] = o.Cultivar
	}
	if o.Ecotype != nil  {
		toSerialize["ecotype"] = o.Ecotype
	}
	if o.Isolate != nil  {
		toSerialize["isolate"] = o.Isolate
	}
	if o.Sex != nil  {
		toSerialize["sex"] = o.Sex
	}
	if o.Strain != nil  {
		toSerialize["strain"] = o.Strain
	}
	if o.TaxId != nil  {
		toSerialize["tax_id"] = o.TaxId
	}
	if o.AssemblyInfo != nil  {
		toSerialize["assembly_info"] = o.AssemblyInfo
	}
	if o.AssemblyStats != nil  {
		toSerialize["assembly_stats"] = o.AssemblyStats
	}
	if o.OrganelleInfo != nil && len(o.GetOrganelleInfo()) > 0  {
		toSerialize["organelle_info"] = o.OrganelleInfo
	}
	if o.AnnotationInfo != nil  {
		toSerialize["annotation_info"] = o.AnnotationInfo
	}
	if o.WgsInfo != nil  {
		toSerialize["wgs_info"] = o.WgsInfo
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsAssemblyDataReport struct {
	value *V1reportsAssemblyDataReport
	isSet bool
}

func (v NullableV1reportsAssemblyDataReport) Get() *V1reportsAssemblyDataReport {
	return v.value
}

func (v *NullableV1reportsAssemblyDataReport) Set(val *V1reportsAssemblyDataReport) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsAssemblyDataReport) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsAssemblyDataReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsAssemblyDataReport(val *V1reportsAssemblyDataReport) *NullableV1reportsAssemblyDataReport {
	return &NullableV1reportsAssemblyDataReport{value: val, isSet: true}
}

func (v NullableV1reportsAssemblyDataReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsAssemblyDataReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


