# USStocksOrderApi

All URIs are relative to *https://api.upstox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelUsOrder**](USStocksOrderApi.md#CancelUsOrder) | **Delete** /v2/us/order/cancel | Cancel US stocks order
[**EstimateUsOrder**](USStocksOrderApi.md#EstimateUsOrder) | **Post** /v2/us/order/estimate | Estimate US stocks order charges
[**GetUsOrderBook**](USStocksOrderApi.md#GetUsOrderBook) | **Get** /v2/us/order/retrieve-all | Get US stocks order book
[**GetUsOrderDetails**](USStocksOrderApi.md#GetUsOrderDetails) | **Get** /v2/us/order/details | Get US stocks order details
[**PlaceUsOrder**](USStocksOrderApi.md#PlaceUsOrder) | **Post** /v2/us/order/place | Place US stocks order

# **CancelUsOrder**
> UsOrderResponse CancelUsOrder(ctx, orderId)
Cancel US stocks order

Cancels an open US stocks order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| The order id to cancel | 

### Return type

[**UsOrderResponse**](UsOrderResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **EstimateUsOrder**
> UsOrderEstimateResponse EstimateUsOrder(ctx, body)
Estimate US stocks order charges

Returns the brokerage, GST and US fees for an order before it is placed; nothing is placed. A market order is estimated at the price sent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UsOrderEstimateRequest**](UsOrderEstimateRequest.md)|  | 

### Return type

[**UsOrderEstimateResponse**](UsOrderEstimateResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetUsOrderBook**
> UsOrderBookResponse GetUsOrderBook(ctx, optional)
Get US stocks order book

Returns the user's US stocks orders for the day, in every state; optionally only those with a tag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***USStocksOrderApiGetUsOrderBookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a USStocksOrderApiGetUsOrderBookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | **optional.String**| Return only orders placed with this tag; exact, case-sensitive match | 

### Return type

[**UsOrderBookResponse**](UsOrderBookResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetUsOrderDetails**
> UsOrderDetailsResponse GetUsOrderDetails(ctx, orderId)
Get US stocks order details

Returns the latest state of one US stocks order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| The order id to fetch | 

### Return type

[**UsOrderDetailsResponse**](UsOrderDetailsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **PlaceUsOrder**
> UsOrderResponse PlaceUsOrder(ctx, body)
Place US stocks order

Places a market or limit order for a US stock, sized by quantity or by USD amount.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UsPlaceOrderRequest**](UsPlaceOrderRequest.md)|  | 

### Return type

[**UsOrderResponse**](UsOrderResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

