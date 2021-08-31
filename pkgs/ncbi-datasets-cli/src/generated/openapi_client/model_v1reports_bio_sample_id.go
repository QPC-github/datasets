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

// V1reportsBioSampleId struct for V1reportsBioSampleId
type V1reportsBioSampleId struct {
	Db *string `json:"db,omitempty"`
	Label *string `json:"label,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewV1reportsBioSampleId instantiates a new V1reportsBioSampleId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsBioSampleId() *V1reportsBioSampleId {
	this := V1reportsBioSampleId{}
	return &this
}

// NewV1reportsBioSampleIdWithDefaults instantiates a new V1reportsBioSampleId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsBioSampleIdWithDefaults() *V1reportsBioSampleId {
	this := V1reportsBioSampleId{}
	return &this
}

// GetDb returns the Db field value if set, zero value otherwise.
func (o *V1reportsBioSampleId) GetDb() string {
	if o == nil || o.Db == nil {
		var ret string
		return ret
	}
	return *o.Db
}

// GetDbOk returns a tuple with the Db field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleId) GetDbOk() (*string, bool) {
	if o == nil || o.Db == nil {
		return nil, false
	}
	return o.Db, true
}

// HasDb returns a boolean if a field has been set.
func (o *V1reportsBioSampleId) HasDb() bool {
	if o != nil && o.Db != nil {
		return true
	}

	return false
}

// SetDb gets a reference to the given string and assigns it to the Db field.
func (o *V1reportsBioSampleId) SetDb(v string) {
	o.Db = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *V1reportsBioSampleId) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleId) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *V1reportsBioSampleId) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *V1reportsBioSampleId) SetLabel(v string) {
	o.Label = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *V1reportsBioSampleId) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleId) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *V1reportsBioSampleId) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *V1reportsBioSampleId) SetValue(v string) {
	o.Value = &v
}

func (o V1reportsBioSampleId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Db != nil  {
		toSerialize["db"] = o.Db
	}
	if o.Label != nil  {
		toSerialize["label"] = o.Label
	}
	if o.Value != nil  {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsBioSampleId struct {
	value *V1reportsBioSampleId
	isSet bool
}

func (v NullableV1reportsBioSampleId) Get() *V1reportsBioSampleId {
	return v.value
}

func (v *NullableV1reportsBioSampleId) Set(val *V1reportsBioSampleId) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsBioSampleId) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsBioSampleId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsBioSampleId(val *V1reportsBioSampleId) *NullableV1reportsBioSampleId {
	return &NullableV1reportsBioSampleId{value: val, isSet: true}
}

func (v NullableV1reportsBioSampleId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsBioSampleId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


