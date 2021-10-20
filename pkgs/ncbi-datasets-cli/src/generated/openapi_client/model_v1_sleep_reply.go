/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/rehydrate/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
)

// V1SleepReply struct for V1SleepReply
type V1SleepReply struct {
	Reply *string `json:"reply,omitempty"`
}

// NewV1SleepReply instantiates a new V1SleepReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1SleepReply() *V1SleepReply {
	this := V1SleepReply{}
	return &this
}

// NewV1SleepReplyWithDefaults instantiates a new V1SleepReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1SleepReplyWithDefaults() *V1SleepReply {
	this := V1SleepReply{}
	return &this
}

// GetReply returns the Reply field value if set, zero value otherwise.
func (o *V1SleepReply) GetReply() string {
	if o == nil || o.Reply == nil {
		var ret string
		return ret
	}
	return *o.Reply
}

// GetReplyOk returns a tuple with the Reply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SleepReply) GetReplyOk() (*string, bool) {
	if o == nil || o.Reply == nil {
		return nil, false
	}
	return o.Reply, true
}

// HasReply returns a boolean if a field has been set.
func (o *V1SleepReply) HasReply() bool {
	if o != nil && o.Reply != nil {
		return true
	}

	return false
}

// SetReply gets a reference to the given string and assigns it to the Reply field.
func (o *V1SleepReply) SetReply(v string) {
	o.Reply = &v
}

func (o V1SleepReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Reply != nil  {
		toSerialize["reply"] = o.Reply
	}
	return json.Marshal(toSerialize)
}

type NullableV1SleepReply struct {
	value *V1SleepReply
	isSet bool
}

func (v NullableV1SleepReply) Get() *V1SleepReply {
	return v.value
}

func (v *NullableV1SleepReply) Set(val *V1SleepReply) {
	v.value = val
	v.isSet = true
}

func (v NullableV1SleepReply) IsSet() bool {
	return v.isSet
}

func (v *NullableV1SleepReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1SleepReply(val *V1SleepReply) *NullableV1SleepReply {
	return &NullableV1SleepReply{value: val, isSet: true}
}

func (v NullableV1SleepReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1SleepReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

