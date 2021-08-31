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
	"fmt"
)

// V1GenomicRegionGenomicRegionType the model 'V1GenomicRegionGenomicRegionType'
type V1GenomicRegionGenomicRegionType string

// List of v1GenomicRegionGenomicRegionType
const (
	V1GENOMICREGIONGENOMICREGIONTYPE_UNKNOWN V1GenomicRegionGenomicRegionType = "UNKNOWN"
	V1GENOMICREGIONGENOMICREGIONTYPE_REFSEQ_GENE V1GenomicRegionGenomicRegionType = "REFSEQ_GENE"
	V1GENOMICREGIONGENOMICREGIONTYPE_PSEUDOGENE V1GenomicRegionGenomicRegionType = "PSEUDOGENE"
	V1GENOMICREGIONGENOMICREGIONTYPE_BIOLOGICAL_REGION V1GenomicRegionGenomicRegionType = "BIOLOGICAL_REGION"
	V1GENOMICREGIONGENOMICREGIONTYPE_OTHER V1GenomicRegionGenomicRegionType = "OTHER"
)

var allowedV1GenomicRegionGenomicRegionTypeEnumValues = []V1GenomicRegionGenomicRegionType{
	"UNKNOWN",
	"REFSEQ_GENE",
	"PSEUDOGENE",
	"BIOLOGICAL_REGION",
	"OTHER",
}

func (v *V1GenomicRegionGenomicRegionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1GenomicRegionGenomicRegionType(value)
	for _, existing := range allowedV1GenomicRegionGenomicRegionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1GenomicRegionGenomicRegionType", value)
}

// NewV1GenomicRegionGenomicRegionTypeFromValue returns a pointer to a valid V1GenomicRegionGenomicRegionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1GenomicRegionGenomicRegionTypeFromValue(v string) (*V1GenomicRegionGenomicRegionType, error) {
	ev := V1GenomicRegionGenomicRegionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1GenomicRegionGenomicRegionType: valid values are %v", v, allowedV1GenomicRegionGenomicRegionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1GenomicRegionGenomicRegionType) IsValid() bool {
	for _, existing := range allowedV1GenomicRegionGenomicRegionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1GenomicRegionGenomicRegionType value
func (v V1GenomicRegionGenomicRegionType) Ptr() *V1GenomicRegionGenomicRegionType {
	return &v
}

type NullableV1GenomicRegionGenomicRegionType struct {
	value *V1GenomicRegionGenomicRegionType
	isSet bool
}

func (v NullableV1GenomicRegionGenomicRegionType) Get() *V1GenomicRegionGenomicRegionType {
	return v.value
}

func (v *NullableV1GenomicRegionGenomicRegionType) Set(val *V1GenomicRegionGenomicRegionType) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GenomicRegionGenomicRegionType) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GenomicRegionGenomicRegionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GenomicRegionGenomicRegionType(val *V1GenomicRegionGenomicRegionType) *NullableV1GenomicRegionGenomicRegionType {
	return &NullableV1GenomicRegionGenomicRegionType{value: val, isSet: true}
}

func (v NullableV1GenomicRegionGenomicRegionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GenomicRegionGenomicRegionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

