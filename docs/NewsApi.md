# NewsApi

All URIs are relative to *https://api.upstox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNews**](NewsApi.md#GetNews) | **Get** /v2/news | Get News

# **GetNews**
> GetNewsResponse GetNews(ctx, category, optional)
Get News

Fetches news articles for the given instrument keys, or the user's current positions or holdings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **category** | **string**| Category of news to fetch. Allowed values: instrument_keys, positions, holdings | 
 **optional** | ***NewsApiGetNewsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NewsApiGetNewsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instrumentKeys** | **optional.String**| Comma-separated list of instrument keys. Required when category is instrument_keys. | 
 **pageNumber** | **optional.Int32**| Page number for pagination | [default to 1]
 **pageSize** | **optional.Int32**| Number of records per page | [default to 100]

### Return type

[**GetNewsResponse**](GetNewsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

