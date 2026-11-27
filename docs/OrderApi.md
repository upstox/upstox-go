# OrderApi

All URIs are relative to *https://api.upstox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelMultiOrder**](OrderApi.md#CancelMultiOrder) | **Delete** /v2/order/multi/cancel | Cancel multi order
[**CancelOrder**](OrderApi.md#CancelOrder) | **Delete** /v2/order/cancel | Cancel order
[**ExitPositions**](OrderApi.md#ExitPositions) | **Post** /v2/order/positions/exit | Exit all positions
[**GetOrderBook**](OrderApi.md#GetOrderBook) | **Get** /v2/order/retrieve-all | Get order book
[**GetOrderDetails**](OrderApi.md#GetOrderDetails) | **Get** /v2/order/history | Get order history
[**GetOrderStatus**](OrderApi.md#GetOrderStatus) | **Get** /v2/order/details | Get order details
[**GetTradeHistory**](OrderApi.md#GetTradeHistory) | **Get** /v2/order/trades/get-trades-for-day | Get trades
[**GetTradesByOrder**](OrderApi.md#GetTradesByOrder) | **Get** /v2/order/trades | Get trades for order
[**ModifyOrder**](OrderApi.md#ModifyOrder) | **Put** /v2/order/modify | Modify order
[**PlaceMultiOrder**](OrderApi.md#PlaceMultiOrder) | **Post** /v2/order/multi/place | Place multi order
[**PlaceOrder**](OrderApi.md#PlaceOrder) | **Post** /v2/order/place | Place order

# **CancelMultiOrder**
> CancelOrExitMultiOrderResponse CancelMultiOrder(ctx, optional)
Cancel multi order

API to cancel all the open or pending orders which can be applied to both AMO and regular orders.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiCancelMultiOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiCancelMultiOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | **optional.String**| The tag associated with the orders for which the orders must be cancelled | 
 **segment** | **optional.String**| The segment for which the orders must be cancelled | 
 **xAlgoName** | **optional.String**| Algorithm name, used to track and manage orders | 

### Return type

[**CancelOrExitMultiOrderResponse**](CancelOrExitMultiOrderResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **CancelOrder**
> CancelOrderResponse CancelOrder(ctx, orderId, optional)
Cancel order

Cancel order API can be used for two purposes: Cancelling an open order which could be an AMO or a normal order It is also used to EXIT a CO or OCO(bracket order)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| The order ID for which the order must be cancelled | 
 **optional** | ***OrderApiCancelOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiCancelOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAlgoName** | **optional.String**| Algorithm name, used to track and manage orders | 

### Return type

[**CancelOrderResponse**](CancelOrderResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **ExitPositions**
> CancelOrExitMultiOrderResponse ExitPositions(ctx, optional)
Exit all positions

This API provides the functionality to exit all the positions 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiExitPositionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiExitPositionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | **optional.String**| The tag associated with the positions for which the positions must be exit | 
 **segment** | **optional.String**| The segment for which the positions must be exit | 
 **xAlgoName** | **optional.String**| Algorithm name, used to track and manage orders | 

### Return type

[**CancelOrExitMultiOrderResponse**](CancelOrExitMultiOrderResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetOrderBook**
> GetOrderBookResponse GetOrderBook(ctx, )
Get order book

This API provides the list of orders placed by the user. The orders placed by the user is transient for a day and is cleared by the end of the trading session. This API returns all states of the orders, namely, open, pending, and filled ones.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetOrderBookResponse**](GetOrderBookResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetOrderDetails**
> GetOrderResponse GetOrderDetails(ctx, optional)
Get order history

This API provides the details of the particular order the user has placed. The orders placed by the user is transient for a day and are cleared by the end of the trading session. This API returns all states of the orders, namely, open, pending, and filled ones.  The order history can be requested either using order_id or tag.  If both the options are passed, the history of the order which precisely matches both the order_id and tag will be returned in the response.  If only the tag is passed, the history of all the associated orders which matches the tag will be shared in the response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiGetOrderDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiGetOrderDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **optional.String**| The order reference ID for which the order history is required | 
 **tag** | **optional.String**| The unique tag of the order for which the order history is being requested | 

### Return type

[**GetOrderResponse**](GetOrderResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetOrderStatus**
> GetOrderDetailsResponse GetOrderStatus(ctx, optional)
Get order details

This API provides the recent detail of the particular order the user has placed. The orders placed by the user is transient for a day and are cleared by the end of the trading session.\\n\\nThe order details can be requested using order_id.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiGetOrderStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiGetOrderStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **optional.String**| The order reference ID for which the order details is required | 

### Return type

[**GetOrderDetailsResponse**](GetOrderDetailsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetTradeHistory**
> GetTradeResponse GetTradeHistory(ctx, )
Get trades

Retrieve the trades executed for the day

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetTradeResponse**](GetTradeResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetTradesByOrder**
> GetTradeResponse GetTradesByOrder(ctx, orderId)
Get trades for order

Retrieve the trades executed for an order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| The order ID for which the order to get order trades | 

### Return type

[**GetTradeResponse**](GetTradeResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **ModifyOrder**
> ModifyOrderResponse ModifyOrder(ctx, body, optional)
Modify order

This API allows you to modify an order. For modification orderId is mandatory. With orderId you need to send the optional parameter which needs to be modified. In case the optional parameters aren't sent, the default will be considered from the original order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModifyOrderRequest**](ModifyOrderRequest.md)|  | 
 **optional** | ***OrderApiModifyOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiModifyOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAlgoName** | **optional.**| Algorithm name, used to track and manage orders | 

### Return type

[**ModifyOrderResponse**](ModifyOrderResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **PlaceMultiOrder**
> MultiOrderResponse PlaceMultiOrder(ctx, body, optional)
Place multi order

This API allows you to place multiple orders to the exchange via Upstox.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]MultiOrderRequest**](MultiOrderRequest.md)|  | 
 **optional** | ***OrderApiPlaceMultiOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiPlaceMultiOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **origin** | **optional.**|  | 
 **xAlgoName** | **optional.**| Algorithm name, used to track and manage orders | 

### Return type

[**MultiOrderResponse**](MultiOrderResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **PlaceOrder**
> PlaceOrderResponse PlaceOrder(ctx, body, optional)
Place order

This API allows you to place an order to the exchange via Upstox.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PlaceOrderRequest**](PlaceOrderRequest.md)|  | 
 **optional** | ***OrderApiPlaceOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiPlaceOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAlgoName** | **optional.**| Algorithm name, used to track and manage orders | 

### Return type

[**PlaceOrderResponse**](PlaceOrderResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

