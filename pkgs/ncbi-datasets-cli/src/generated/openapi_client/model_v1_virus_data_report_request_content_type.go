/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets version 1 API is considred stable and will not be subject to breaking changes. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
	"fmt"
)

// V1VirusDataReportRequestContentType the model 'V1VirusDataReportRequestContentType'
type V1VirusDataReportRequestContentType string

// List of v1VirusDataReportRequestContentType
const (
	V1VIRUSDATAREPORTREQUESTCONTENTTYPE_COMPLETE V1VirusDataReportRequestContentType = "COMPLETE"
	V1VIRUSDATAREPORTREQUESTCONTENTTYPE_ACCESSIONS_ONLY V1VirusDataReportRequestContentType = "ACCESSIONS_ONLY"
)

// All allowed values of V1VirusDataReportRequestContentType enum
var AllowedV1VirusDataReportRequestContentTypeEnumValues = []V1VirusDataReportRequestContentType{
	"COMPLETE",
	"ACCESSIONS_ONLY",
}

func (v *V1VirusDataReportRequestContentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1VirusDataReportRequestContentType(value)
	for _, existing := range AllowedV1VirusDataReportRequestContentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1VirusDataReportRequestContentType", value)
}

// NewV1VirusDataReportRequestContentTypeFromValue returns a pointer to a valid V1VirusDataReportRequestContentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1VirusDataReportRequestContentTypeFromValue(v string) (*V1VirusDataReportRequestContentType, error) {
	ev := V1VirusDataReportRequestContentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1VirusDataReportRequestContentType: valid values are %v", v, AllowedV1VirusDataReportRequestContentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1VirusDataReportRequestContentType) IsValid() bool {
	for _, existing := range AllowedV1VirusDataReportRequestContentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1VirusDataReportRequestContentType value
func (v V1VirusDataReportRequestContentType) Ptr() *V1VirusDataReportRequestContentType {
	return &v
}

type NullableV1VirusDataReportRequestContentType struct {
	value *V1VirusDataReportRequestContentType
	isSet bool
}

func (v NullableV1VirusDataReportRequestContentType) Get() *V1VirusDataReportRequestContentType {
	return v.value
}

func (v *NullableV1VirusDataReportRequestContentType) Set(val *V1VirusDataReportRequestContentType) {
	v.value = val
	v.isSet = true
}

func (v NullableV1VirusDataReportRequestContentType) IsSet() bool {
	return v.isSet
}

func (v *NullableV1VirusDataReportRequestContentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1VirusDataReportRequestContentType(val *V1VirusDataReportRequestContentType) *NullableV1VirusDataReportRequestContentType {
	return &NullableV1VirusDataReportRequestContentType{value: val, isSet: true}
}

func (v NullableV1VirusDataReportRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1VirusDataReportRequestContentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

