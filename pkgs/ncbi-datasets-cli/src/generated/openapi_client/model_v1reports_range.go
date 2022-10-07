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

// V1reportsRange struct for V1reportsRange
type V1reportsRange struct {
	Begin *string `json:"begin,omitempty"`
	End *string `json:"end,omitempty"`
	Orientation *V1reportsOrientation `json:"orientation,omitempty"`
	Order *int32 `json:"order,omitempty"`
}

// NewV1reportsRange instantiates a new V1reportsRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsRange() *V1reportsRange {
	this := V1reportsRange{}
	var orientation V1reportsOrientation = V1REPORTSORIENTATION_NONE
	this.Orientation = &orientation
	return &this
}

// NewV1reportsRangeWithDefaults instantiates a new V1reportsRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsRangeWithDefaults() *V1reportsRange {
	this := V1reportsRange{}
	var orientation V1reportsOrientation = V1REPORTSORIENTATION_NONE
	this.Orientation = &orientation
	return &this
}

// GetBegin returns the Begin field value if set, zero value otherwise.
func (o *V1reportsRange) GetBegin() string {
	if o == nil || o.Begin == nil {
		var ret string
		return ret
	}
	return *o.Begin
}

// GetBeginOk returns a tuple with the Begin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsRange) GetBeginOk() (*string, bool) {
	if o == nil || o.Begin == nil {
		return nil, false
	}
	return o.Begin, true
}

// HasBegin returns a boolean if a field has been set.
func (o *V1reportsRange) HasBegin() bool {
	if o != nil && o.Begin != nil {
		return true
	}

	return false
}

// SetBegin gets a reference to the given string and assigns it to the Begin field.
func (o *V1reportsRange) SetBegin(v string) {
	o.Begin = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *V1reportsRange) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsRange) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *V1reportsRange) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *V1reportsRange) SetEnd(v string) {
	o.End = &v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *V1reportsRange) GetOrientation() V1reportsOrientation {
	if o == nil || o.Orientation == nil {
		var ret V1reportsOrientation
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsRange) GetOrientationOk() (*V1reportsOrientation, bool) {
	if o == nil || o.Orientation == nil {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *V1reportsRange) HasOrientation() bool {
	if o != nil && o.Orientation != nil {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given V1reportsOrientation and assigns it to the Orientation field.
func (o *V1reportsRange) SetOrientation(v V1reportsOrientation) {
	o.Orientation = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *V1reportsRange) GetOrder() int32 {
	if o == nil || o.Order == nil {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsRange) GetOrderOk() (*int32, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *V1reportsRange) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *V1reportsRange) SetOrder(v int32) {
	o.Order = &v
}

func (o V1reportsRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Begin != nil  {
		toSerialize["begin"] = o.Begin
	}
	if o.End != nil  {
		toSerialize["end"] = o.End
	}
	if o.Orientation != nil  {
		toSerialize["orientation"] = o.Orientation
	}
	if o.Order != nil  {
		toSerialize["order"] = o.Order
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsRange struct {
	value *V1reportsRange
	isSet bool
}

func (v NullableV1reportsRange) Get() *V1reportsRange {
	return v.value
}

func (v *NullableV1reportsRange) Set(val *V1reportsRange) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsRange) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsRange(val *V1reportsRange) *NullableV1reportsRange {
	return &NullableV1reportsRange{value: val, isSet: true}
}

func (v NullableV1reportsRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


