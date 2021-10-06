package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

type ApiGetAccountRequest struct {
	accountname string
	ctx         _context.Context
	ApiService  *DefaultApiService
}

/*
GetAccount Get account

Get account

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAccountRequest
*/
func (a *DefaultApiService) GetAccount(ctx _context.Context, accountname string) ApiGetAccountRequest {
	return ApiGetAccountRequest{
		ApiService:  a,
		ctx:         ctx,
		accountname: accountname,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *DefaultApiService) GetAccountExecute(r ApiGetAccountRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod         = _nethttp.MethodGet
		localVarPostBody           interface{}
		localVarFormFileName       string
		localVarFileName           string
		localVarFileBytes          []byte
		localVarReturnValue        []map[string]interface{}
		localVarReturnValueAccount map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ListAccounts")
	if err != nil {
		return localVarReturnValueAccount, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValueAccount, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValueAccount, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValueAccount, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValueAccount, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValueAccount, localVarHTTPResponse, newErr
	}

	for _, value := range localVarReturnValue {
		if value["name"] == r.accountname {
			localVarReturnValueAccount = value
			break
		}
	}

	return localVarReturnValueAccount, localVarHTTPResponse, nil
}

func (r ApiGetAccountRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetAccountExecute(r)
}

type ApiGetHostRequest struct {
	hostName   string
	ctx        _context.Context
	ApiService *DefaultApiService
}

func (r ApiGetHostRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetHostExecute(r)
}

/*
GetHost Method for GetHost

List hosts

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetHostRequest
*/
func (a *DefaultApiService) GetHost(ctx _context.Context, hostName string) ApiGetHostRequest {
	return ApiGetHostRequest{
		ApiService: a,
		ctx:        ctx,
		hostName:   hostName,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *DefaultApiService) GetHostExecute(r ApiGetHostRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod      = _nethttp.MethodGet
		localVarPostBody        interface{}
		localVarFormFileName    string
		localVarFileName        string
		localVarFileBytes       []byte
		localVarReturnValue     []map[string]interface{}
		localVarReturnValueHost map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ListHosts")
	if err != nil {
		return localVarReturnValueHost, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/hosts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValueHost, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValueHost, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValueHost, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValueHost, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValueHost, localVarHTTPResponse, newErr
	}

	for _, host := range localVarReturnValue {
		if host["host"].(string) == r.hostName {
			localVarReturnValueHost = host
		}
	}

	return localVarReturnValueHost, localVarHTTPResponse, nil
}

type ApiGetIntegrationInstanceRequest struct {
	identifier string
	ctx        _context.Context
	ApiService *DefaultApiService
	size       *InlineObject
}

func (r ApiGetIntegrationInstanceRequest) Size(size InlineObject) ApiGetIntegrationInstanceRequest {
	r.size = &size
	return r
}

func (r ApiGetIntegrationInstanceRequest) Identifier(identifier string) ApiGetIntegrationInstanceRequest {
	r.identifier = identifier
	return r
}

func (r ApiGetIntegrationInstanceRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetIntegrationInstanceExecute(r)
}

/*
GetIntegrationInstance List integrations

List all integrations

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetIntegrationInstanceRequest
*/
func (a *DefaultApiService) GetIntegrationInstance(ctx _context.Context) ApiGetIntegrationInstanceRequest {
	return ApiGetIntegrationInstanceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DefaultApiService) GetIntegrationInstanceExecute(r ApiGetIntegrationInstanceRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ListIntegrations")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/settings/integration/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.size
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
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

	for _, instance := range localVarReturnValue["instances"].([]interface{}) {
		if inst := instance.(map[string]interface{}); inst["name"].(string) == r.identifier || inst["id"].(string) == r.identifier {
			return inst, localVarHTTPResponse, nil
		}
	}

	return nil, localVarHTTPResponse, nil
}

type ApiGetIntegrationInstanceAccountRequest struct {
	ctx        _context.Context
	ApiService *DefaultApiService
	acc        string
	identifier string
	size       *InlineObject
}

func (r ApiGetIntegrationInstanceAccountRequest) Size(size InlineObject) ApiGetIntegrationInstanceAccountRequest {
	r.size = &size
	return r
}

func (r ApiGetIntegrationInstanceAccountRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetIntegrationInstanceAccountExecute(r)
}

/*
GetIntegrationInstanceAccount List integrations

List all integrations

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param acc
 @return ApiGetIntegrationInstanceAccountRequest
*/
func (a *DefaultApiService) GetIntegrationInstanceAccount(ctx _context.Context, acc string) ApiGetIntegrationInstanceAccountRequest {
	return ApiGetIntegrationInstanceAccountRequest{
		ApiService: a,
		ctx:        ctx,
		acc:        acc,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DefaultApiService) GetIntegrationInstanceAccountExecute(r ApiGetIntegrationInstanceAccountRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ListIntegrationsAccount")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{acc}/settings/integration/search"
	localVarPath = strings.Replace(localVarPath, "{"+"acc"+"}", _neturl.PathEscape(parameterToString(r.acc, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.size
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
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

	for _, instance := range localVarReturnValue["instances"].([]interface{}) {
		if inst := instance.(map[string]interface{}); inst["name"].(string) == r.identifier || inst["id"].(string) == r.identifier {
			return inst, localVarHTTPResponse, nil
		}
	}

	return nil, localVarHTTPResponse, nil
}
