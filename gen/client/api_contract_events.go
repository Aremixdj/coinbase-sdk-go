/*
Coinbase Platform API

This is the OpenAPI 3.0 specification for the Coinbase Platform APIs, used in conjunction with the Coinbase Platform SDKs.

API version: 0.0.1-alpha
Contact: yuga.cohler@coinbase.com
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


// ContractEventsAPIService ContractEventsAPI service
type ContractEventsAPIService service

type ApiListContractEventsRequest struct {
	ctx context.Context
	ApiService *ContractEventsAPIService
	networkId string
	protocolName *string
	contractAddress string
	contractName *string
	eventName *string
	fromBlockHeight *int32
	toBlockHeight *int32
	nextPage *string
}

// Case-sensitive name of the blockchain protocol
func (r ApiListContractEventsRequest) ProtocolName(protocolName string) ApiListContractEventsRequest {
	r.protocolName = &protocolName
	return r
}

// Case-sensitive name of the specific contract within the project
func (r ApiListContractEventsRequest) ContractName(contractName string) ApiListContractEventsRequest {
	r.contractName = &contractName
	return r
}

// Case-sensitive name of the event to filter for in the contract&#39;s logs
func (r ApiListContractEventsRequest) EventName(eventName string) ApiListContractEventsRequest {
	r.eventName = &eventName
	return r
}

// Lower bound of the block range to query (inclusive)
func (r ApiListContractEventsRequest) FromBlockHeight(fromBlockHeight int32) ApiListContractEventsRequest {
	r.fromBlockHeight = &fromBlockHeight
	return r
}

// Upper bound of the block range to query (inclusive)
func (r ApiListContractEventsRequest) ToBlockHeight(toBlockHeight int32) ApiListContractEventsRequest {
	r.toBlockHeight = &toBlockHeight
	return r
}

// Pagination token for retrieving the next set of results
func (r ApiListContractEventsRequest) NextPage(nextPage string) ApiListContractEventsRequest {
	r.nextPage = &nextPage
	return r
}

func (r ApiListContractEventsRequest) Execute() (*ContractEventList, *http.Response, error) {
	return r.ApiService.ListContractEventsExecute(r)
}

/*
ListContractEvents Get contract events

Retrieve events for a specific contract

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Unique identifier for the blockchain network
 @param contractAddress EVM address of the smart contract (42 characters, including '0x', in lowercase)
 @return ApiListContractEventsRequest
*/
func (a *ContractEventsAPIService) ListContractEvents(ctx context.Context, networkId string, contractAddress string) ApiListContractEventsRequest {
	return ApiListContractEventsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		contractAddress: contractAddress,
	}
}

// Execute executes the request
//  @return ContractEventList
func (a *ContractEventsAPIService) ListContractEventsExecute(r ApiListContractEventsRequest) (*ContractEventList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContractEventList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContractEventsAPIService.ListContractEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/networks/{network_id}/smart_contracts/{contract_address}/events"
	localVarPath = strings.Replace(localVarPath, "{"+"network_id"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"contract_address"+"}", url.PathEscape(parameterValueToString(r.contractAddress, "contractAddress")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.protocolName == nil {
		return localVarReturnValue, nil, reportError("protocolName is required and must be specified")
	}
	if r.contractName == nil {
		return localVarReturnValue, nil, reportError("contractName is required and must be specified")
	}
	if r.eventName == nil {
		return localVarReturnValue, nil, reportError("eventName is required and must be specified")
	}
	if r.fromBlockHeight == nil {
		return localVarReturnValue, nil, reportError("fromBlockHeight is required and must be specified")
	}
	if r.toBlockHeight == nil {
		return localVarReturnValue, nil, reportError("toBlockHeight is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "protocol_name", r.protocolName, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "contract_name", r.contractName, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "event_name", r.eventName, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "from_block_height", r.fromBlockHeight, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "to_block_height", r.toBlockHeight, "")
	if r.nextPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "next_page", r.nextPage, "")
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
