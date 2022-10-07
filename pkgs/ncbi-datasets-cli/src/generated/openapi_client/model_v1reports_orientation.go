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

// V1reportsOrientation the model 'V1reportsOrientation'
type V1reportsOrientation string

// List of v1reportsOrientation
const (
	V1REPORTSORIENTATION_NONE V1reportsOrientation = "none"
	V1REPORTSORIENTATION_PLUS V1reportsOrientation = "plus"
	V1REPORTSORIENTATION_MINUS V1reportsOrientation = "minus"
)

// All allowed values of V1reportsOrientation enum
var AllowedV1reportsOrientationEnumValues = []V1reportsOrientation{
	"none",
	"plus",
	"minus",
}

func (v *V1reportsOrientation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1reportsOrientation(value)
	for _, existing := range AllowedV1reportsOrientationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1reportsOrientation", value)
}

// NewV1reportsOrientationFromValue returns a pointer to a valid V1reportsOrientation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1reportsOrientationFromValue(v string) (*V1reportsOrientation, error) {
	ev := V1reportsOrientation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1reportsOrientation: valid values are %v", v, AllowedV1reportsOrientationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1reportsOrientation) IsValid() bool {
	for _, existing := range AllowedV1reportsOrientationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1reportsOrientation value
func (v V1reportsOrientation) Ptr() *V1reportsOrientation {
	return &v
}

type NullableV1reportsOrientation struct {
	value *V1reportsOrientation
	isSet bool
}

func (v NullableV1reportsOrientation) Get() *V1reportsOrientation {
	return v.value
}

func (v *NullableV1reportsOrientation) Set(val *V1reportsOrientation) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsOrientation) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsOrientation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsOrientation(val *V1reportsOrientation) *NullableV1reportsOrientation {
	return &NullableV1reportsOrientation{value: val, isSet: true}
}

func (v NullableV1reportsOrientation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsOrientation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

