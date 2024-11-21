/*
Coinbase Platform API

This is the OpenAPI 3.0 specification for the Coinbase Platform APIs, used in conjunction with the Coinbase Platform SDKs.

API version: 0.0.1-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type FundAPI interface {

	/*
	CreateFundOperation Create a new fund operation.

	Create a new fund operation with an address.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param walletId The ID of the wallet the address belongs to.
	@param addressId The onchain address to be funded.
	@return ApiCreateFundOperationRequest
	*/
	CreateFundOperation(ctx context.Context, walletId string, addressId string) ApiCreateFundOperationRequest

	// CreateFundOperationExecute executes the request
	//  @return FundOperation
	CreateFundOperationExecute(r ApiCreateFundOperationRequest) (*FundOperation, *http.Response, error)

	/*
	CreateFundQuote Create a Fund Operation quote.

	Create a new fund operation with an address.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param walletId The ID of the wallet the address belongs to.
	@param addressId The onchain address to be funded.
	@return ApiCreateFundQuoteRequest
	*/
	CreateFundQuote(ctx context.Context, walletId string, addressId string) ApiCreateFundQuoteRequest

	// CreateFundQuoteExecute executes the request
	//  @return FundQuote
	CreateFundQuoteExecute(r ApiCreateFundQuoteRequest) (*FundQuote, *http.Response, error)

	/*
	GetFundOperation Get fund operation.

	Get fund operation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param walletId The ID of the wallet the address belongs to.
	@param addressId The onchain address of the address that created the fund operation.
	@param fundOperationId The ID of the fund operation to fetch.
	@return ApiGetFundOperationRequest
	*/
	GetFundOperation(ctx context.Context, walletId string, addressId string, fundOperationId string) ApiGetFundOperationRequest

	// GetFundOperationExecute executes the request
	//  @return FundOperation
	GetFundOperationExecute(r ApiGetFundOperationRequest) (*FundOperation, *http.Response, error)

	/*
	ListFundOperations List fund operations for an address.

	List fund operations for an address.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param walletId The ID of the wallet the address belongs to.
	@param addressId The onchain address of the address to list fund operations for.
	@return ApiListFundOperationsRequest
	*/
	ListFundOperations(ctx context.Context, walletId string, addressId string) ApiListFundOperationsRequest

	// ListFundOperationsExecute executes the request
	//  @return FundOperationList
	ListFundOperationsExecute(r ApiListFundOperationsRequest) (*FundOperationList, *http.Response, error)
}

// FundAPIService FundAPI service
type FundAPIService service

type ApiCreateFundOperationRequest struct {
	ctx context.Context
	ApiService FundAPI
	walletId string
	addressId string
	createFundOperationRequest *CreateFundOperationRequest
}

func (r ApiCreateFundOperationRequest) CreateFundOperationRequest(createFundOperationRequest CreateFundOperationRequest) ApiCreateFundOperationRequest {
	r.createFundOperationRequest = &createFundOperationRequest
	return r
}

func (r ApiCreateFundOperationRequest) Execute() (*FundOperation, *http.Response, error) {
	return r.ApiService.CreateFundOperationExecute(r)
}

/*
CreateFundOperation Create a new fund operation.

Create a new fund operation with an address.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param walletId The ID of the wallet the address belongs to.
 @param addressId The onchain address to be funded.
 @return ApiCreateFundOperationRequest
*/
func (a *FundAPIService) CreateFundOperation(ctx context.Context, walletId string, addressId string) ApiCreateFundOperationRequest {
	return ApiCreateFundOperationRequest{
		ApiService: a,
		ctx: ctx,
		walletId: walletId,
		addressId: addressId,
	}
}

// Execute executes the request
//  @return FundOperation
func (a *FundAPIService) CreateFundOperationExecute(r ApiCreateFundOperationRequest) (*FundOperation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FundOperation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundAPIService.CreateFundOperation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/wallets/{wallet_id}/addresses/{address_id}/fund_operations"
	localVarPath = strings.Replace(localVarPath, "{"+"wallet_id"+"}", url.PathEscape(parameterValueToString(r.walletId, "walletId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address_id"+"}", url.PathEscape(parameterValueToString(r.addressId, "addressId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createFundOperationRequest == nil {
		return localVarReturnValue, nil, reportError("createFundOperationRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.createFundOperationRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

type ApiCreateFundQuoteRequest struct {
	ctx context.Context
	ApiService FundAPI
	walletId string
	addressId string
	createFundQuoteRequest *CreateFundQuoteRequest
}

func (r ApiCreateFundQuoteRequest) CreateFundQuoteRequest(createFundQuoteRequest CreateFundQuoteRequest) ApiCreateFundQuoteRequest {
	r.createFundQuoteRequest = &createFundQuoteRequest
	return r
}

func (r ApiCreateFundQuoteRequest) Execute() (*FundQuote, *http.Response, error) {
	return r.ApiService.CreateFundQuoteExecute(r)
}

/*
CreateFundQuote Create a Fund Operation quote.

Create a new fund operation with an address.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param walletId The ID of the wallet the address belongs to.
 @param addressId The onchain address to be funded.
 @return ApiCreateFundQuoteRequest
*/
func (a *FundAPIService) CreateFundQuote(ctx context.Context, walletId string, addressId string) ApiCreateFundQuoteRequest {
	return ApiCreateFundQuoteRequest{
		ApiService: a,
		ctx: ctx,
		walletId: walletId,
		addressId: addressId,
	}
}

// Execute executes the request
//  @return FundQuote
func (a *FundAPIService) CreateFundQuoteExecute(r ApiCreateFundQuoteRequest) (*FundQuote, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FundQuote
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundAPIService.CreateFundQuote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/wallets/{wallet_id}/addresses/{address_id}/fund_operations/quote"
	localVarPath = strings.Replace(localVarPath, "{"+"wallet_id"+"}", url.PathEscape(parameterValueToString(r.walletId, "walletId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address_id"+"}", url.PathEscape(parameterValueToString(r.addressId, "addressId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createFundQuoteRequest == nil {
		return localVarReturnValue, nil, reportError("createFundQuoteRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.createFundQuoteRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

type ApiGetFundOperationRequest struct {
	ctx context.Context
	ApiService FundAPI
	walletId string
	addressId string
	fundOperationId string
}

func (r ApiGetFundOperationRequest) Execute() (*FundOperation, *http.Response, error) {
	return r.ApiService.GetFundOperationExecute(r)
}

/*
GetFundOperation Get fund operation.

Get fund operation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param walletId The ID of the wallet the address belongs to.
 @param addressId The onchain address of the address that created the fund operation.
 @param fundOperationId The ID of the fund operation to fetch.
 @return ApiGetFundOperationRequest
*/
func (a *FundAPIService) GetFundOperation(ctx context.Context, walletId string, addressId string, fundOperationId string) ApiGetFundOperationRequest {
	return ApiGetFundOperationRequest{
		ApiService: a,
		ctx: ctx,
		walletId: walletId,
		addressId: addressId,
		fundOperationId: fundOperationId,
	}
}

// Execute executes the request
//  @return FundOperation
func (a *FundAPIService) GetFundOperationExecute(r ApiGetFundOperationRequest) (*FundOperation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FundOperation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundAPIService.GetFundOperation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/wallets/{wallet_id}/addresses/{address_id}/fund_operations/{fund_operation_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"wallet_id"+"}", url.PathEscape(parameterValueToString(r.walletId, "walletId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address_id"+"}", url.PathEscape(parameterValueToString(r.addressId, "addressId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fund_operation_id"+"}", url.PathEscape(parameterValueToString(r.fundOperationId, "fundOperationId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

type ApiListFundOperationsRequest struct {
	ctx context.Context
	ApiService FundAPI
	walletId string
	addressId string
	limit *int32
	page *string
}

// A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 10.
func (r ApiListFundOperationsRequest) Limit(limit int32) ApiListFundOperationsRequest {
	r.limit = &limit
	return r
}

// A cursor for pagination across multiple pages of results. Don&#39;t include this parameter on the first call. Use the next_page value returned in a previous response to request subsequent results.
func (r ApiListFundOperationsRequest) Page(page string) ApiListFundOperationsRequest {
	r.page = &page
	return r
}

func (r ApiListFundOperationsRequest) Execute() (*FundOperationList, *http.Response, error) {
	return r.ApiService.ListFundOperationsExecute(r)
}

/*
ListFundOperations List fund operations for an address.

List fund operations for an address.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param walletId The ID of the wallet the address belongs to.
 @param addressId The onchain address of the address to list fund operations for.
 @return ApiListFundOperationsRequest
*/
func (a *FundAPIService) ListFundOperations(ctx context.Context, walletId string, addressId string) ApiListFundOperationsRequest {
	return ApiListFundOperationsRequest{
		ApiService: a,
		ctx: ctx,
		walletId: walletId,
		addressId: addressId,
	}
}

// Execute executes the request
//  @return FundOperationList
func (a *FundAPIService) ListFundOperationsExecute(r ApiListFundOperationsRequest) (*FundOperationList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FundOperationList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundAPIService.ListFundOperations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/wallets/{wallet_id}/addresses/{address_id}/fund_operations"
	localVarPath = strings.Replace(localVarPath, "{"+"wallet_id"+"}", url.PathEscape(parameterValueToString(r.walletId, "walletId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address_id"+"}", url.PathEscape(parameterValueToString(r.addressId, "addressId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
