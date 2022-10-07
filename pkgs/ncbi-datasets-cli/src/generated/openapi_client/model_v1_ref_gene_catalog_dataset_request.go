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

// V1RefGeneCatalogDatasetRequest struct for V1RefGeneCatalogDatasetRequest
type V1RefGeneCatalogDatasetRequest struct {
	OpaqueSolrQuery *string `json:"opaque_solr_query,omitempty"`
	Files *[]V1RefGeneCatalogDatasetRequestFileType `json:"files,omitempty"`
	ElementFlankConfig *V1ElementFlankConfig `json:"element_flank_config,omitempty"`
}

// NewV1RefGeneCatalogDatasetRequest instantiates a new V1RefGeneCatalogDatasetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1RefGeneCatalogDatasetRequest() *V1RefGeneCatalogDatasetRequest {
	this := V1RefGeneCatalogDatasetRequest{}
	return &this
}

// NewV1RefGeneCatalogDatasetRequestWithDefaults instantiates a new V1RefGeneCatalogDatasetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1RefGeneCatalogDatasetRequestWithDefaults() *V1RefGeneCatalogDatasetRequest {
	this := V1RefGeneCatalogDatasetRequest{}
	return &this
}

// GetOpaqueSolrQuery returns the OpaqueSolrQuery field value if set, zero value otherwise.
func (o *V1RefGeneCatalogDatasetRequest) GetOpaqueSolrQuery() string {
	if o == nil || o.OpaqueSolrQuery == nil {
		var ret string
		return ret
	}
	return *o.OpaqueSolrQuery
}

// GetOpaqueSolrQueryOk returns a tuple with the OpaqueSolrQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RefGeneCatalogDatasetRequest) GetOpaqueSolrQueryOk() (*string, bool) {
	if o == nil || o.OpaqueSolrQuery == nil {
		return nil, false
	}
	return o.OpaqueSolrQuery, true
}

// HasOpaqueSolrQuery returns a boolean if a field has been set.
func (o *V1RefGeneCatalogDatasetRequest) HasOpaqueSolrQuery() bool {
	if o != nil && o.OpaqueSolrQuery != nil {
		return true
	}

	return false
}

// SetOpaqueSolrQuery gets a reference to the given string and assigns it to the OpaqueSolrQuery field.
func (o *V1RefGeneCatalogDatasetRequest) SetOpaqueSolrQuery(v string) {
	o.OpaqueSolrQuery = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *V1RefGeneCatalogDatasetRequest) GetFiles() []V1RefGeneCatalogDatasetRequestFileType {
	if o == nil || o.Files == nil {
		var ret []V1RefGeneCatalogDatasetRequestFileType
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RefGeneCatalogDatasetRequest) GetFilesOk() (*[]V1RefGeneCatalogDatasetRequestFileType, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *V1RefGeneCatalogDatasetRequest) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []V1RefGeneCatalogDatasetRequestFileType and assigns it to the Files field.
func (o *V1RefGeneCatalogDatasetRequest) SetFiles(v []V1RefGeneCatalogDatasetRequestFileType) {
	o.Files = &v
}

// GetElementFlankConfig returns the ElementFlankConfig field value if set, zero value otherwise.
func (o *V1RefGeneCatalogDatasetRequest) GetElementFlankConfig() V1ElementFlankConfig {
	if o == nil || o.ElementFlankConfig == nil {
		var ret V1ElementFlankConfig
		return ret
	}
	return *o.ElementFlankConfig
}

// GetElementFlankConfigOk returns a tuple with the ElementFlankConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RefGeneCatalogDatasetRequest) GetElementFlankConfigOk() (*V1ElementFlankConfig, bool) {
	if o == nil || o.ElementFlankConfig == nil {
		return nil, false
	}
	return o.ElementFlankConfig, true
}

// HasElementFlankConfig returns a boolean if a field has been set.
func (o *V1RefGeneCatalogDatasetRequest) HasElementFlankConfig() bool {
	if o != nil && o.ElementFlankConfig != nil {
		return true
	}

	return false
}

// SetElementFlankConfig gets a reference to the given V1ElementFlankConfig and assigns it to the ElementFlankConfig field.
func (o *V1RefGeneCatalogDatasetRequest) SetElementFlankConfig(v V1ElementFlankConfig) {
	o.ElementFlankConfig = &v
}

func (o V1RefGeneCatalogDatasetRequest) MarshalJSON() ([]byte, error) {
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

type NullableV1RefGeneCatalogDatasetRequest struct {
	value *V1RefGeneCatalogDatasetRequest
	isSet bool
}

func (v NullableV1RefGeneCatalogDatasetRequest) Get() *V1RefGeneCatalogDatasetRequest {
	return v.value
}

func (v *NullableV1RefGeneCatalogDatasetRequest) Set(val *V1RefGeneCatalogDatasetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1RefGeneCatalogDatasetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1RefGeneCatalogDatasetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1RefGeneCatalogDatasetRequest(val *V1RefGeneCatalogDatasetRequest) *NullableV1RefGeneCatalogDatasetRequest {
	return &NullableV1RefGeneCatalogDatasetRequest{value: val, isSet: true}
}

func (v NullableV1RefGeneCatalogDatasetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1RefGeneCatalogDatasetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


