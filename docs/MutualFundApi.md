# MutualFundApi

All URIs are relative to *https://api.upstox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMutualFundHoldings**](MutualFundApi.md#GetMutualFundHoldings) | **Get** /v2/mf/holdings | Get mutual fund holdings
[**GetMutualFundOrder**](MutualFundApi.md#GetMutualFundOrder) | **Get** /v2/mf/orders/{order_id} | Get mutual fund order by id
[**GetMutualFundOrders**](MutualFundApi.md#GetMutualFundOrders) | **Get** /v2/mf/orders | Get mutual fund orders
[**GetMutualFundSips**](MutualFundApi.md#GetMutualFundSips) | **Get** /v2/mf/sips | Get mutual fund SIPs

# **GetMutualFundHoldings**
> GetMutualFundHoldingsResponse GetMutualFundHoldings(ctx, )
Get mutual fund holdings

Returns mutual fund holdings for the authenticated user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetMutualFundHoldingsResponse**](GetMutualFundHoldingsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetMutualFundOrder**
> GetMutualFundOrderDetailsResponse GetMutualFundOrder(ctx, orderId)
Get mutual fund order by id

Returns a single mutual fund order by order id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| Mutual fund order id | 

### Return type

[**GetMutualFundOrderDetailsResponse**](GetMutualFundOrderDetailsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetMutualFundOrders**
> GetMutualFundOrdersResponse GetMutualFundOrders(ctx, optional)
Get mutual fund orders

Returns mutual fund orders from the last 7 days for the authenticated user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MutualFundApiGetMutualFundOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MutualFundApiGetMutualFundOrdersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **optional.String**| Mutual fund order status | [default to ALL]
 **transactionType** | **optional.String**| Mutual fund order transaction type | [default to ALL]
 **pageNumber** | **optional.Int32**| Page number for pagination. | [default to 1]
 **records** | **optional.Int32**| no of records in a single page | [default to 10]

### Return type

[**GetMutualFundOrdersResponse**](GetMutualFundOrdersResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetMutualFundSips**
> GetMutualFundSipsResponse GetMutualFundSips(ctx, optional)
Get mutual fund SIPs

Returns active and paused SIP orders for the authenticated user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MutualFundApiGetMutualFundSipsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MutualFundApiGetMutualFundSipsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Page number for pagination. | [default to 1]
 **records** | **optional.Int32**| no of records in a single page | [default to 10]

### Return type

[**GetMutualFundSipsResponse**](GetMutualFundSipsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

