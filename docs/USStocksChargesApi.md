# USStocksChargesApi

All URIs are relative to *https://api.upstox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculateUsMargin**](USStocksChargesApi.md#CalculateUsMargin) | **Post** /v2/us/charges/margin | Calculate US stocks margin

# **CalculateUsMargin**
> UsMarginResponse CalculateUsMargin(ctx, body)
Calculate US stocks margin

Returns the margin up to 20 orders require (trade value + brokerage + GST for a buy), the buying power available, any shortfall, and whether it is sufficient. Takes order estimate bodies; nothing is placed. A market order is valued at the last traded price. No leverage is offered on US stocks.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UsMarginRequest**](UsMarginRequest.md)|  | 

### Return type

[**UsMarginResponse**](UsMarginResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

