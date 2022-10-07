/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets version 1 API is considred stable and will not be subject to breaking changes. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"os"
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// ProkaryoteApiService ProkaryoteApi service
type ProkaryoteApiService service

type ApiDownloadProkaryoteGenePackageRequest struct {
	ctx _context.Context
	ApiService *ProkaryoteApiService
	accessions []string	
	includeAnnotationType *[]V1Fasta	
	geneFlankConfigLength *int32	
	taxon *string	
	filename *string	
    Headers map[string]string
}

// Select additional types of annotation to include in the data package.  If unset, no annotation is provided.
func (r *ApiDownloadProkaryoteGenePackageRequest) IncludeAnnotationType(includeAnnotationType []V1Fasta) *ApiDownloadProkaryoteGenePackageRequest {
	r.includeAnnotationType = &includeAnnotationType
	return r
}
func (r *ApiDownloadProkaryoteGenePackageRequest) GeneFlankConfigLength(geneFlankConfigLength int32) *ApiDownloadProkaryoteGenePackageRequest {
	r.geneFlankConfigLength = &geneFlankConfigLength
	return r
}
// NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank When specified, return data from this taxon and its subtree
func (r *ApiDownloadProkaryoteGenePackageRequest) Taxon(taxon string) *ApiDownloadProkaryoteGenePackageRequest {
	r.taxon = &taxon
	return r
}
// Output file name.
func (r *ApiDownloadProkaryoteGenePackageRequest) Filename(filename string) *ApiDownloadProkaryoteGenePackageRequest {
	r.filename = &filename
	return r
}

func (r ApiDownloadProkaryoteGenePackageRequest) Execute() (*os.File, *_nethttp.Response, error) {
	return r.ApiService.DownloadProkaryoteGenePackageExecute(r)
}

/*
DownloadProkaryoteGenePackage Get a prokaryote gene dataset by RefSeq protein accession

Get a prokaryote gene dataset including gene and protein fasta sequence, annotation and metadata by prokaryote protein accession.
 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accessions WP prokaryote protein accession
 @return ApiDownloadProkaryoteGenePackageRequest
*/
func (a *ProkaryoteApiService) DownloadProkaryoteGenePackage(ctx _context.Context, accessions []string) ApiDownloadProkaryoteGenePackageRequest {
	return ApiDownloadProkaryoteGenePackageRequest{
		ApiService: a,
		ctx: ctx,
		accessions: accessions,
	}
}

// Execute executes the request
//  @return *os.File
func (a *ProkaryoteApiService) DownloadProkaryoteGenePackageExecute(r ApiDownloadProkaryoteGenePackageRequest) (*os.File, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProkaryoteApiService.DownloadProkaryoteGenePackage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/protein/accession/{accessions}/download"
	localVarPath = strings.Replace(localVarPath, "{"+"accessions"+"}", _neturl.PathEscape(parameterToString(r.accessions, "csv")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.includeAnnotationType != nil {
		t := *r.includeAnnotationType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("include_annotation_type", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("include_annotation_type", parameterToString(t, "multi"))
		}
	}
	if r.geneFlankConfigLength != nil {
		localVarQueryParams.Add("gene_flank_config.length", parameterToString(*r.geneFlankConfigLength, ""))
	}
	if r.taxon != nil {
		localVarQueryParams.Add("taxon", parameterToString(*r.taxon, ""))
	}
	if r.filename != nil {
		localVarQueryParams.Add("filename", parameterToString(*r.filename, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/zip", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// override localVarHeaderParams with the headers passed into the function
	if len(r.Headers) > 0 {
		for k, v := range r.Headers { 
			localVarHeaderParams[k] = v
		}
	}

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuthHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	if localVarHTTPResponse.Header.Get("Content-Type") != "application/json" {
		return localVarReturnValue, localVarHTTPResponse, nil
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDownloadProkaryoteGenePackagePostRequest struct {
	ctx _context.Context
	ApiService *ProkaryoteApiService
	v1ProkaryoteGeneRequest *V1ProkaryoteGeneRequest	
	filename *string	
    Headers map[string]string
}

func (r *ApiDownloadProkaryoteGenePackagePostRequest) V1ProkaryoteGeneRequest(v1ProkaryoteGeneRequest V1ProkaryoteGeneRequest) *ApiDownloadProkaryoteGenePackagePostRequest {
	r.v1ProkaryoteGeneRequest = &v1ProkaryoteGeneRequest
	return r
}
// Output file name.
func (r *ApiDownloadProkaryoteGenePackagePostRequest) Filename(filename string) *ApiDownloadProkaryoteGenePackagePostRequest {
	r.filename = &filename
	return r
}

func (r ApiDownloadProkaryoteGenePackagePostRequest) Execute() (*os.File, *_nethttp.Response, error) {
	return r.ApiService.DownloadProkaryoteGenePackagePostExecute(r)
}

/*
DownloadProkaryoteGenePackagePost Get a prokaryote gene dataset by RefSeq protein accession by POST

Get a prokaryote gene dataset including gene and protein fasta sequence, annotation and metadata by prokaryote protein accession by POST.
 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDownloadProkaryoteGenePackagePostRequest
*/
func (a *ProkaryoteApiService) DownloadProkaryoteGenePackagePost(ctx _context.Context, v1ProkaryoteGeneRequest *V1ProkaryoteGeneRequest) ApiDownloadProkaryoteGenePackagePostRequest {
	return ApiDownloadProkaryoteGenePackagePostRequest{
		ApiService: a,
		ctx: ctx,
		v1ProkaryoteGeneRequest: v1ProkaryoteGeneRequest,
	}
}

// Execute executes the request
//  @return *os.File
func (a *ProkaryoteApiService) DownloadProkaryoteGenePackagePostExecute(r ApiDownloadProkaryoteGenePackagePostRequest) (*os.File, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProkaryoteApiService.DownloadProkaryoteGenePackagePost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/protein/accession/download"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.v1ProkaryoteGeneRequest == nil {
		return localVarReturnValue, nil, reportError("v1ProkaryoteGeneRequest is required and must be specified")
	}

	if r.filename != nil {
		localVarQueryParams.Add("filename", parameterToString(*r.filename, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/zip", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// override localVarHeaderParams with the headers passed into the function
	if len(r.Headers) > 0 {
		for k, v := range r.Headers { 
			localVarHeaderParams[k] = v
		}
	}

	// body params
	localVarPostBody = r.v1ProkaryoteGeneRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuthHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	if localVarHTTPResponse.Header.Get("Content-Type") != "application/json" {
		return localVarReturnValue, localVarHTTPResponse, nil
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
