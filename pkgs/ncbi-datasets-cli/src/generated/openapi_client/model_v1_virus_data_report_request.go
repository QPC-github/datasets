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

// V1VirusDataReportRequest struct for V1VirusDataReportRequest
type V1VirusDataReportRequest struct {
	Filter *V1VirusDatasetFilter `json:"filter,omitempty"`
	ReturnedContent *V1VirusDataReportRequestContentType `json:"returned_content,omitempty"`
	TableFields *[]string `json:"table_fields,omitempty"`
	PageSize *int32 `json:"page_size,omitempty"`
	PageToken *string `json:"page_token,omitempty"`
}

// NewV1VirusDataReportRequest instantiates a new V1VirusDataReportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1VirusDataReportRequest() *V1VirusDataReportRequest {
	this := V1VirusDataReportRequest{}
	var returnedContent V1VirusDataReportRequestContentType = V1VIRUSDATAREPORTREQUESTCONTENTTYPE_COMPLETE
	this.ReturnedContent = &returnedContent
	return &this
}

// NewV1VirusDataReportRequestWithDefaults instantiates a new V1VirusDataReportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1VirusDataReportRequestWithDefaults() *V1VirusDataReportRequest {
	this := V1VirusDataReportRequest{}
	var returnedContent V1VirusDataReportRequestContentType = V1VIRUSDATAREPORTREQUESTCONTENTTYPE_COMPLETE
	this.ReturnedContent = &returnedContent
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *V1VirusDataReportRequest) GetFilter() V1VirusDatasetFilter {
	if o == nil || o.Filter == nil {
		var ret V1VirusDatasetFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1VirusDataReportRequest) GetFilterOk() (*V1VirusDatasetFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *V1VirusDataReportRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given V1VirusDatasetFilter and assigns it to the Filter field.
func (o *V1VirusDataReportRequest) SetFilter(v V1VirusDatasetFilter) {
	o.Filter = &v
}

// GetReturnedContent returns the ReturnedContent field value if set, zero value otherwise.
func (o *V1VirusDataReportRequest) GetReturnedContent() V1VirusDataReportRequestContentType {
	if o == nil || o.ReturnedContent == nil {
		var ret V1VirusDataReportRequestContentType
		return ret
	}
	return *o.ReturnedContent
}

// GetReturnedContentOk returns a tuple with the ReturnedContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1VirusDataReportRequest) GetReturnedContentOk() (*V1VirusDataReportRequestContentType, bool) {
	if o == nil || o.ReturnedContent == nil {
		return nil, false
	}
	return o.ReturnedContent, true
}

// HasReturnedContent returns a boolean if a field has been set.
func (o *V1VirusDataReportRequest) HasReturnedContent() bool {
	if o != nil && o.ReturnedContent != nil {
		return true
	}

	return false
}

// SetReturnedContent gets a reference to the given V1VirusDataReportRequestContentType and assigns it to the ReturnedContent field.
func (o *V1VirusDataReportRequest) SetReturnedContent(v V1VirusDataReportRequestContentType) {
	o.ReturnedContent = &v
}

// GetTableFields returns the TableFields field value if set, zero value otherwise.
func (o *V1VirusDataReportRequest) GetTableFields() []string {
	if o == nil || o.TableFields == nil {
		var ret []string
		return ret
	}
	return *o.TableFields
}

// GetTableFieldsOk returns a tuple with the TableFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1VirusDataReportRequest) GetTableFieldsOk() (*[]string, bool) {
	if o == nil || o.TableFields == nil {
		return nil, false
	}
	return o.TableFields, true
}

// HasTableFields returns a boolean if a field has been set.
func (o *V1VirusDataReportRequest) HasTableFields() bool {
	if o != nil && o.TableFields != nil {
		return true
	}

	return false
}

// SetTableFields gets a reference to the given []string and assigns it to the TableFields field.
func (o *V1VirusDataReportRequest) SetTableFields(v []string) {
	o.TableFields = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *V1VirusDataReportRequest) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1VirusDataReportRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *V1VirusDataReportRequest) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *V1VirusDataReportRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *V1VirusDataReportRequest) GetPageToken() string {
	if o == nil || o.PageToken == nil {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1VirusDataReportRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || o.PageToken == nil {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *V1VirusDataReportRequest) HasPageToken() bool {
	if o != nil && o.PageToken != nil {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *V1VirusDataReportRequest) SetPageToken(v string) {
	o.PageToken = &v
}

func (o V1VirusDataReportRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil  {
		toSerialize["filter"] = o.Filter
	}
	if o.ReturnedContent != nil  {
		toSerialize["returned_content"] = o.ReturnedContent
	}
	if o.TableFields != nil && len(o.GetTableFields()) > 0  {
		toSerialize["table_fields"] = o.TableFields
	}
	if o.PageSize != nil  {
		toSerialize["page_size"] = o.PageSize
	}
	if o.PageToken != nil  {
		toSerialize["page_token"] = o.PageToken
	}
	return json.Marshal(toSerialize)
}

type NullableV1VirusDataReportRequest struct {
	value *V1VirusDataReportRequest
	isSet bool
}

func (v NullableV1VirusDataReportRequest) Get() *V1VirusDataReportRequest {
	return v.value
}

func (v *NullableV1VirusDataReportRequest) Set(val *V1VirusDataReportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1VirusDataReportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1VirusDataReportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1VirusDataReportRequest(val *V1VirusDataReportRequest) *NullableV1VirusDataReportRequest {
	return &NullableV1VirusDataReportRequest{value: val, isSet: true}
}

func (v NullableV1VirusDataReportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1VirusDataReportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


