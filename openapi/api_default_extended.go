package openapi

import (
	_context "context"
	_nethttp "net/http"
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

func (r ApiGetAccountRequest) SetAccountName(accountname string) ApiGetAccountRequest {
	r.accountname = accountname
	return r
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DefaultApiService) GetAccountExecute(r ApiGetAccountRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		accountsWrapper      AccountsWrapper
		localVarHTTPResponse *_nethttp.Response
		localVarReturnValue  map[string]interface{}
	)

	accountsWrapper, localVarHTTPResponse, err := a.ListAccounts(r.ctx).Execute()
	if err != nil {
		return nil, localVarHTTPResponse, err
	}

	for _, account := range accountsWrapper.Items {
		if r.accountname == account["name"].(string) {
			localVarReturnValue = account
			break
		}
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

func (r ApiGetAccountRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetAccountExecute(r)
}

type ApiGetHostRequest struct {
	hostname   string
	ctx        _context.Context
	ApiService *DefaultApiService
}

func (r ApiGetHostRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetHostExecute(r)
}

/*
GetHost Method for GetHost

Get host

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetHostRequest
*/
func (a *DefaultApiService) GetHost(ctx _context.Context, hostname string) ApiGetHostRequest {
	return ApiGetHostRequest{
		ApiService: a,
		ctx:        ctx,
		hostname:   hostname,
	}
}

func (r ApiGetHostRequest) SetHostName(hostname string) ApiGetHostRequest {
	r.hostname = hostname
	return r
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *DefaultApiService) GetHostExecute(r ApiGetHostRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		hosts                Hosts
		localVarHTTPResponse *_nethttp.Response
		localVarReturnValue  map[string]interface{}
	)

	hosts, localVarHTTPResponse, err := a.ListHosts(r.ctx).Execute()
	if err != nil {
		return nil, localVarHTTPResponse, err
	}

	for _, host := range hosts.Items {
		if r.hostname == host["host"].(string) {
			localVarReturnValue = host
			break
		}
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetIntegrationInstanceRequest struct {
	identifier string
	ctx        _context.Context
	ApiService *DefaultApiService
}

func (r ApiGetIntegrationInstanceRequest) SetIdentifier(identifier string) ApiGetIntegrationInstanceRequest {
	r.identifier = identifier
	return r
}

func (r ApiGetIntegrationInstanceRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetIntegrationInstanceExecute(r)
}

/*
GetIntegrationInstance Get integration

Get integration

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
		localVarHTTPResponse *_nethttp.Response
	)
	size := NewInlineObject2()
	size.SetSize(500)
	integrations, localVarHTTPResponse, err := a.ListIntegrations(r.ctx).Size(*size).Execute()
	if err != nil {
		return nil, localVarHTTPResponse, err
	}

	for _, instance := range integrations["instances"].([]interface{}) {
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
}

func (r ApiGetIntegrationInstanceAccountRequest) SetIdentifier(identifier string) ApiGetIntegrationInstanceAccountRequest {
	r.identifier = identifier
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
		localVarHTTPResponse *_nethttp.Response
	)
	size := NewInlineObject3()
	size.SetSize(500)
	integrations, localVarHTTPResponse, err := a.ListIntegrationsAccount(r.ctx, "acc_"+r.acc).Size(*size).Execute()
	if err != nil {
		return nil, localVarHTTPResponse, err
	}

	for _, instance := range integrations["instances"].([]interface{}) {
		if inst := instance.(map[string]interface{}); inst["name"].(string) == r.identifier || inst["id"].(string) == r.identifier {
			return inst, localVarHTTPResponse, nil
		}
	}

	return nil, localVarHTTPResponse, nil
}

type ApiGetClassifierRequest struct {
	ctx        _context.Context
	ApiService *DefaultApiService
	identifier string
}

func (r ApiGetClassifierRequest) SetIdentifier(identifier string) ApiGetClassifierRequest {
	r.identifier = identifier
	return r
}

func (r ApiGetClassifierRequest) Execute() (InstanceClassifier, *_nethttp.Response, error) {
	return r.ApiService.GetClassifierExecute(r)
}

/*
GetClassifier Get classifier

Get classifier

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetClassifierRequest
*/
func (a *DefaultApiService) GetClassifier(ctx _context.Context) ApiGetClassifierRequest {
	return ApiGetClassifierRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DefaultApiService) GetClassifierExecute(r ApiGetClassifierRequest) (InstanceClassifier, *_nethttp.Response, error) {
	var (
		localVarHTTPResponse *_nethttp.Response
	)
	classifiers, localVarHTTPResponse, err := a.ListClassifiers(r.ctx).Execute()
	if err != nil {
		return InstanceClassifier{}, localVarHTTPResponse, err
	}

	for _, classifier := range classifiers.GetClassifiers() {
		if *classifier.Name == r.identifier || *classifier.Id == r.identifier {
			return classifier, localVarHTTPResponse, nil
		}
	}
	notFoundErr := GenericOpenAPIError{error: "Could not find classifier with identifier: " + r.identifier}
	return InstanceClassifier{}, localVarHTTPResponse, notFoundErr
}

type ApiGetClassifierAccountRequest struct {
	ctx        _context.Context
	ApiService *DefaultApiService
	acc        string
	identifier string
}

func (r ApiGetClassifierAccountRequest) SetIdentifier(identifier string) ApiGetClassifierAccountRequest {
	r.identifier = identifier
	return r
}

func (r ApiGetClassifierAccountRequest) Execute() (InstanceClassifier, *_nethttp.Response, error) {
	return r.ApiService.GetClassifierAccountExecute(r)
}

/*
GetClassifierAccount Get classifier

Get classifier

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param acc
 @return ApiGetClassifierAccountRequest
*/
func (a *DefaultApiService) GetClassifierAccount(ctx _context.Context, acc string) ApiGetClassifierAccountRequest {
	return ApiGetClassifierAccountRequest{
		ApiService: a,
		ctx:        ctx,
		acc:        acc,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DefaultApiService) GetClassifierAccountExecute(r ApiGetClassifierAccountRequest) (InstanceClassifier, *_nethttp.Response, error) {
	var (
		localVarHTTPResponse *_nethttp.Response
	)
	classifiers, localVarHTTPResponse, err := a.ListClassifiersAccount(r.ctx, "acc_"+r.acc).Execute()
	if err != nil {
		return InstanceClassifier{}, localVarHTTPResponse, err
	}

	for _, classifier := range classifiers.GetClassifiers() {
		if *classifier.Name == r.identifier || *classifier.Id == r.identifier {
			return classifier, localVarHTTPResponse, nil
		}
	}
	notFoundErr := GenericOpenAPIError{error: "Could not find classifier with identifier: " + r.identifier}
	return InstanceClassifier{}, localVarHTTPResponse, notFoundErr
}
