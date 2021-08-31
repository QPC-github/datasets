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

// V1reportsProtein struct for V1reportsProtein
type V1reportsProtein struct {
	AccessionVersion *string `json:"accession_version,omitempty"`
	Name *string `json:"name,omitempty"`
	Length *int32 `json:"length,omitempty"`
	IsoformName *string `json:"isoform_name,omitempty"`
	EnsemblProtein *string `json:"ensembl_protein,omitempty"`
	MaturePeptides *[]V1reportsMaturePeptide `json:"mature_peptides,omitempty"`
}

// NewV1reportsProtein instantiates a new V1reportsProtein object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsProtein() *V1reportsProtein {
	this := V1reportsProtein{}
	return &this
}

// NewV1reportsProteinWithDefaults instantiates a new V1reportsProtein object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsProteinWithDefaults() *V1reportsProtein {
	this := V1reportsProtein{}
	return &this
}

// GetAccessionVersion returns the AccessionVersion field value if set, zero value otherwise.
func (o *V1reportsProtein) GetAccessionVersion() string {
	if o == nil || o.AccessionVersion == nil {
		var ret string
		return ret
	}
	return *o.AccessionVersion
}

// GetAccessionVersionOk returns a tuple with the AccessionVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsProtein) GetAccessionVersionOk() (*string, bool) {
	if o == nil || o.AccessionVersion == nil {
		return nil, false
	}
	return o.AccessionVersion, true
}

// HasAccessionVersion returns a boolean if a field has been set.
func (o *V1reportsProtein) HasAccessionVersion() bool {
	if o != nil && o.AccessionVersion != nil {
		return true
	}

	return false
}

// SetAccessionVersion gets a reference to the given string and assigns it to the AccessionVersion field.
func (o *V1reportsProtein) SetAccessionVersion(v string) {
	o.AccessionVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1reportsProtein) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsProtein) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1reportsProtein) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1reportsProtein) SetName(v string) {
	o.Name = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *V1reportsProtein) GetLength() int32 {
	if o == nil || o.Length == nil {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsProtein) GetLengthOk() (*int32, bool) {
	if o == nil || o.Length == nil {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *V1reportsProtein) HasLength() bool {
	if o != nil && o.Length != nil {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *V1reportsProtein) SetLength(v int32) {
	o.Length = &v
}

// GetIsoformName returns the IsoformName field value if set, zero value otherwise.
func (o *V1reportsProtein) GetIsoformName() string {
	if o == nil || o.IsoformName == nil {
		var ret string
		return ret
	}
	return *o.IsoformName
}

// GetIsoformNameOk returns a tuple with the IsoformName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsProtein) GetIsoformNameOk() (*string, bool) {
	if o == nil || o.IsoformName == nil {
		return nil, false
	}
	return o.IsoformName, true
}

// HasIsoformName returns a boolean if a field has been set.
func (o *V1reportsProtein) HasIsoformName() bool {
	if o != nil && o.IsoformName != nil {
		return true
	}

	return false
}

// SetIsoformName gets a reference to the given string and assigns it to the IsoformName field.
func (o *V1reportsProtein) SetIsoformName(v string) {
	o.IsoformName = &v
}

// GetEnsemblProtein returns the EnsemblProtein field value if set, zero value otherwise.
func (o *V1reportsProtein) GetEnsemblProtein() string {
	if o == nil || o.EnsemblProtein == nil {
		var ret string
		return ret
	}
	return *o.EnsemblProtein
}

// GetEnsemblProteinOk returns a tuple with the EnsemblProtein field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsProtein) GetEnsemblProteinOk() (*string, bool) {
	if o == nil || o.EnsemblProtein == nil {
		return nil, false
	}
	return o.EnsemblProtein, true
}

// HasEnsemblProtein returns a boolean if a field has been set.
func (o *V1reportsProtein) HasEnsemblProtein() bool {
	if o != nil && o.EnsemblProtein != nil {
		return true
	}

	return false
}

// SetEnsemblProtein gets a reference to the given string and assigns it to the EnsemblProtein field.
func (o *V1reportsProtein) SetEnsemblProtein(v string) {
	o.EnsemblProtein = &v
}

// GetMaturePeptides returns the MaturePeptides field value if set, zero value otherwise.
func (o *V1reportsProtein) GetMaturePeptides() []V1reportsMaturePeptide {
	if o == nil || o.MaturePeptides == nil {
		var ret []V1reportsMaturePeptide
		return ret
	}
	return *o.MaturePeptides
}

// GetMaturePeptidesOk returns a tuple with the MaturePeptides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsProtein) GetMaturePeptidesOk() (*[]V1reportsMaturePeptide, bool) {
	if o == nil || o.MaturePeptides == nil {
		return nil, false
	}
	return o.MaturePeptides, true
}

// HasMaturePeptides returns a boolean if a field has been set.
func (o *V1reportsProtein) HasMaturePeptides() bool {
	if o != nil && o.MaturePeptides != nil {
		return true
	}

	return false
}

// SetMaturePeptides gets a reference to the given []V1reportsMaturePeptide and assigns it to the MaturePeptides field.
func (o *V1reportsProtein) SetMaturePeptides(v []V1reportsMaturePeptide) {
	o.MaturePeptides = &v
}

func (o V1reportsProtein) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessionVersion != nil  {
		toSerialize["accession_version"] = o.AccessionVersion
	}
	if o.Name != nil  {
		toSerialize["name"] = o.Name
	}
	if o.Length != nil  {
		toSerialize["length"] = o.Length
	}
	if o.IsoformName != nil  {
		toSerialize["isoform_name"] = o.IsoformName
	}
	if o.EnsemblProtein != nil  {
		toSerialize["ensembl_protein"] = o.EnsemblProtein
	}
	if o.MaturePeptides != nil && len(o.GetMaturePeptides()) > 0  {
		toSerialize["mature_peptides"] = o.MaturePeptides
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsProtein struct {
	value *V1reportsProtein
	isSet bool
}

func (v NullableV1reportsProtein) Get() *V1reportsProtein {
	return v.value
}

func (v *NullableV1reportsProtein) Set(val *V1reportsProtein) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsProtein) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsProtein) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsProtein(val *V1reportsProtein) *NullableV1reportsProtein {
	return &NullableV1reportsProtein{value: val, isSet: true}
}

func (v NullableV1reportsProtein) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsProtein) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


