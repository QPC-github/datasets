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

// V1MicroBiggeDatasetRequest struct for V1MicroBiggeDatasetRequest
type V1MicroBiggeDatasetRequest struct {
	OpaqueSolrQuery *string `json:"opaque_solr_query,omitempty"`
	Files *[]V1MicroBiggeDatasetRequestFileType `json:"files,omitempty"`
	ElementFlankConfig *V1MicroBiggeDatasetRequestElementFlankConfig `json:"element_flank_config,omitempty"`
}

// NewV1MicroBiggeDatasetRequest instantiates a new V1MicroBiggeDatasetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1MicroBiggeDatasetRequest() *V1MicroBiggeDatasetRequest {
	this := V1MicroBiggeDatasetRequest{}
	return &this
}

// NewV1MicroBiggeDatasetRequestWithDefaults instantiates a new V1MicroBiggeDatasetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1MicroBiggeDatasetRequestWithDefaults() *V1MicroBiggeDatasetRequest {
	this := V1MicroBiggeDatasetRequest{}
	return &this
}

// GetOpaqueSolrQuery returns the OpaqueSolrQuery field value if set, zero value otherwise.
func (o *V1MicroBiggeDatasetRequest) GetOpaqueSolrQuery() string {
	if o == nil || o.OpaqueSolrQuery == nil {
		var ret string
		return ret
	}
	return *o.OpaqueSolrQuery
}

// GetOpaqueSolrQueryOk returns a tuple with the OpaqueSolrQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1MicroBiggeDatasetRequest) GetOpaqueSolrQueryOk() (*string, bool) {
	if o == nil || o.OpaqueSolrQuery == nil {
		return nil, false
	}
	return o.OpaqueSolrQuery, true
}

// HasOpaqueSolrQuery returns a boolean if a field has been set.
func (o *V1MicroBiggeDatasetRequest) HasOpaqueSolrQuery() bool {
	if o != nil && o.OpaqueSolrQuery != nil {
		return true
	}

	return false
}

// SetOpaqueSolrQuery gets a reference to the given string and assigns it to the OpaqueSolrQuery field.
func (o *V1MicroBiggeDatasetRequest) SetOpaqueSolrQuery(v string) {
	o.OpaqueSolrQuery = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *V1MicroBiggeDatasetRequest) GetFiles() []V1MicroBiggeDatasetRequestFileType {
	if o == nil || o.Files == nil {
		var ret []V1MicroBiggeDatasetRequestFileType
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1MicroBiggeDatasetRequest) GetFilesOk() (*[]V1MicroBiggeDatasetRequestFileType, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *V1MicroBiggeDatasetRequest) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []V1MicroBiggeDatasetRequestFileType and assigns it to the Files field.
func (o *V1MicroBiggeDatasetRequest) SetFiles(v []V1MicroBiggeDatasetRequestFileType) {
	o.Files = &v
}

// GetElementFlankConfig returns the ElementFlankConfig field value if set, zero value otherwise.
func (o *V1MicroBiggeDatasetRequest) GetElementFlankConfig() V1MicroBiggeDatasetRequestElementFlankConfig {
	if o == nil || o.ElementFlankConfig == nil {
		var ret V1MicroBiggeDatasetRequestElementFlankConfig
		return ret
	}
	return *o.ElementFlankConfig
}

// GetElementFlankConfigOk returns a tuple with the ElementFlankConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1MicroBiggeDatasetRequest) GetElementFlankConfigOk() (*V1MicroBiggeDatasetRequestElementFlankConfig, bool) {
	if o == nil || o.ElementFlankConfig == nil {
		return nil, false
	}
	return o.ElementFlankConfig, true
}

// HasElementFlankConfig returns a boolean if a field has been set.
func (o *V1MicroBiggeDatasetRequest) HasElementFlankConfig() bool {
	if o != nil && o.ElementFlankConfig != nil {
		return true
	}

	return false
}

// SetElementFlankConfig gets a reference to the given V1MicroBiggeDatasetRequestElementFlankConfig and assigns it to the ElementFlankConfig field.
func (o *V1MicroBiggeDatasetRequest) SetElementFlankConfig(v V1MicroBiggeDatasetRequestElementFlankConfig) {
	o.ElementFlankConfig = &v
}

func (o V1MicroBiggeDatasetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OpaqueSolrQuery != nil  {
		toSerialize["opaque_solr_query"] = o.OpaqueSolrQuery
	}
	if o.Files != nil && len(o.GetFiles()) > 0  {
		toSerialize["files"] = o.Files
	}
	if o.ElementFlankConfig != nil  {
		toSerialize["element_flank_config"] = o.ElementFlankConfig
	}
	return json.Marshal(toSerialize)
}

type NullableV1MicroBiggeDatasetRequest struct {
	value *V1MicroBiggeDatasetRequest
	isSet bool
}

func (v NullableV1MicroBiggeDatasetRequest) Get() *V1MicroBiggeDatasetRequest {
	return v.value
}

func (v *NullableV1MicroBiggeDatasetRequest) Set(val *V1MicroBiggeDatasetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1MicroBiggeDatasetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1MicroBiggeDatasetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1MicroBiggeDatasetRequest(val *V1MicroBiggeDatasetRequest) *NullableV1MicroBiggeDatasetRequest {
	return &NullableV1MicroBiggeDatasetRequest{value: val, isSet: true}
}

func (v NullableV1MicroBiggeDatasetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1MicroBiggeDatasetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

