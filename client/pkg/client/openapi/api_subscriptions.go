/*
Relay API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v20200615
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// SubscriptionsApiService SubscriptionsApi service
type SubscriptionsApiService service

type SubscriptionsApiGetWorkflowSubscriptionsRequest struct {
	ctx          context.Context
	ApiService   *SubscriptionsApiService
	workflowPath string
}

func (r SubscriptionsApiGetWorkflowSubscriptionsRequest) Execute() (*UserWorkflowSubscriptions, *http.Response, error) {
	return r.ApiService.GetWorkflowSubscriptionsExecute(r)
}

/*
GetWorkflowSubscriptions Retrieve the current user's workflow subscriptions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workflowPath Folder + Workflow name
 @return SubscriptionsApiGetWorkflowSubscriptionsRequest
*/
func (a *SubscriptionsApiService) GetWorkflowSubscriptions(ctx context.Context, workflowPath string) SubscriptionsApiGetWorkflowSubscriptionsRequest {
	return SubscriptionsApiGetWorkflowSubscriptionsRequest{
		ApiService:   a,
		ctx:          ctx,
		workflowPath: workflowPath,
	}
}

// Execute executes the request
//  @return UserWorkflowSubscriptions
func (a *SubscriptionsApiService) GetWorkflowSubscriptionsExecute(r SubscriptionsApiGetWorkflowSubscriptionsRequest) (*UserWorkflowSubscriptions, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UserWorkflowSubscriptions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsApiService.GetWorkflowSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/subscriptions/workflows/{workflowPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowPath"+"}", url.PathEscape(parameterToString(r.workflowPath, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.puppet.relay.v20200615+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v GetAccessDefaultResponse
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SubscriptionsApiGetWorkflowsSubscriptionsRequest struct {
	ctx        context.Context
	ApiService *SubscriptionsApiService
}

func (r SubscriptionsApiGetWorkflowsSubscriptionsRequest) Execute() (*UserWorkflowsSubscriptions, *http.Response, error) {
	return r.ApiService.GetWorkflowsSubscriptionsExecute(r)
}

/*
GetWorkflowsSubscriptions List subscriptions for all workflows for the current user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SubscriptionsApiGetWorkflowsSubscriptionsRequest
*/
func (a *SubscriptionsApiService) GetWorkflowsSubscriptions(ctx context.Context) SubscriptionsApiGetWorkflowsSubscriptionsRequest {
	return SubscriptionsApiGetWorkflowsSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return UserWorkflowsSubscriptions
func (a *SubscriptionsApiService) GetWorkflowsSubscriptionsExecute(r SubscriptionsApiGetWorkflowsSubscriptionsRequest) (*UserWorkflowsSubscriptions, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UserWorkflowsSubscriptions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsApiService.GetWorkflowsSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/subscriptions/workflows"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.puppet.relay.v20200615+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v GetAccessDefaultResponse
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SubscriptionsApiPutWorkflowSubscriptionsRequest struct {
	ctx                       context.Context
	ApiService                *SubscriptionsApiService
	workflowPath              string
	userWorkflowSubscriptions *UserWorkflowSubscriptions
}

// Update workflow subscriptions
func (r SubscriptionsApiPutWorkflowSubscriptionsRequest) UserWorkflowSubscriptions(userWorkflowSubscriptions UserWorkflowSubscriptions) SubscriptionsApiPutWorkflowSubscriptionsRequest {
	r.userWorkflowSubscriptions = &userWorkflowSubscriptions
	return r
}

func (r SubscriptionsApiPutWorkflowSubscriptionsRequest) Execute() (*UserWorkflowSubscriptions, *http.Response, error) {
	return r.ApiService.PutWorkflowSubscriptionsExecute(r)
}

/*
PutWorkflowSubscriptions Update the current user's workflow subscriptions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workflowPath Folder + Workflow name
 @return SubscriptionsApiPutWorkflowSubscriptionsRequest
*/
func (a *SubscriptionsApiService) PutWorkflowSubscriptions(ctx context.Context, workflowPath string) SubscriptionsApiPutWorkflowSubscriptionsRequest {
	return SubscriptionsApiPutWorkflowSubscriptionsRequest{
		ApiService:   a,
		ctx:          ctx,
		workflowPath: workflowPath,
	}
}

// Execute executes the request
//  @return UserWorkflowSubscriptions
func (a *SubscriptionsApiService) PutWorkflowSubscriptionsExecute(r SubscriptionsApiPutWorkflowSubscriptionsRequest) (*UserWorkflowSubscriptions, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UserWorkflowSubscriptions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsApiService.PutWorkflowSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/subscriptions/workflows/{workflowPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowPath"+"}", url.PathEscape(parameterToString(r.workflowPath, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userWorkflowSubscriptions == nil {
		return localVarReturnValue, nil, reportError("userWorkflowSubscriptions is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.puppet.relay.v20200615+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.puppet.relay.v20200615+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.userWorkflowSubscriptions
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v GetAccessDefaultResponse
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
