# InstrumentsApi

All URIs are relative to *https://api.upstox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchInstrument**](InstrumentsApi.md#SearchInstrument) | **Get** /v2/instruments/search | Search instruments

# **SearchInstrument**
> SearchInstrumentResponse SearchInstrument(ctx, query, optional)
Search instruments

This API provides the functionality to retrieve the instrument details for searched query

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| natural language query | 
 **optional** | ***InstrumentsApiSearchInstrumentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstrumentsApiSearchInstrumentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exchanges** | **optional.String**| Comma separated exchanges | 
 **segments** | **optional.String**| Comma separated segments | 
 **instrumentTypes** | **optional.String**| Comma separated instrument types | 
 **expiry** | **optional.String**| expiry in format: yyyy-MM-dd | 
 **atmOffset** | **optional.Int32**| page number greater than or equal to 1 | 
 **pageNumber** | **optional.Int32**| ATM offset number for queries like ATM, ATM+1, ATM-1, etc. | [default to 1]
 **records** | **optional.Int32**| no of records in a single page | [default to 10]

### Return type

[**SearchInstrumentResponse**](SearchInstrumentResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

