# OrderV3Api

All URIs are relative to *https://api.upstox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelGTTOrder**](OrderV3Api.md#CancelGTTOrder) | **Delete** /v3/order/gtt/cancel | Cancel GTT order
[**CancelOrderV3**](OrderV3Api.md#CancelOrderV3) | **Delete** /v3/order/cancel | 
[**GetGttOrderDetails**](OrderV3Api.md#GetGttOrderDetails) | **Get** /v3/order/gtt | Get GTT order details
[**ModifyGTTOrder**](OrderV3Api.md#ModifyGTTOrder) | **Put** /v3/order/gtt/modify | Modify GTT order
[**ModifyOrderV3**](OrderV3Api.md#ModifyOrderV3) | **Put** /v3/order/modify | 
[**PlaceGTTOrder**](OrderV3Api.md#PlaceGTTOrder) | **Post** /v3/order/gtt/place | Place GTT order
[**PlaceOrderV3**](OrderV3Api.md#PlaceOrderV3) | **Post** /v3/order/place | 

# **CancelGTTOrder**
> GttTriggerOrderResponse CancelGTTOrder(ctx, body, optional)
Cancel GTT order

This API allows you to cancel GTT orders.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GttCancelOrderRequest**](GttCancelOrderRequest.md)|  | 
 **optional** | ***OrderV3ApiCancelGTTOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderV3ApiCancelGTTOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **origin** | **optional.**|  | 
 **xAlgoName** | **optional.**| Algorithm name, used to track and manage orders | 

### Return type

[**GttTriggerOrderResponse**](GttTriggerOrderResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **CancelOrderV3**
> CancelOrderV3Response CancelOrderV3(ctx, orderId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**|  | 
 **optional** | ***OrderV3ApiCancelOrderV3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderV3ApiCancelOrderV3Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAlgoName** | **optional.String**| Algorithm name, used to track and manage orders | 

### Return type

[**CancelOrderV3Response**](CancelOrderV3Response.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetGttOrderDetails**
> GetGttOrderResponse GetGttOrderDetails(ctx, optional)
Get GTT order details

GTT_ORDER_DESCRIPTION

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderV3ApiGetGttOrderDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderV3ApiGetGttOrderDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gttOrderId** | **optional.String**| Unique identifier of the GTT order for which the order history is required | 

### Return type

[**GetGttOrderResponse**](GetGttOrderResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **ModifyGTTOrder**
> GttTriggerOrderResponse ModifyGTTOrder(ctx, body, optional)
Modify GTT order

This API allows you to modify GTT orders.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GttModifyOrderRequest**](GttModifyOrderRequest.md)|  | 
 **optional** | ***OrderV3ApiModifyGTTOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderV3ApiModifyGTTOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **origin** | **optional.**|  | 
 **xAlgoName** | **optional.**| Algorithm name, used to track and manage orders | 

### Return type

[**GttTriggerOrderResponse**](GttTriggerOrderResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **ModifyOrderV3**
> ModifyOrderV3Response ModifyOrderV3(ctx, body, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModifyOrderRequest**](ModifyOrderRequest.md)|  | 
 **optional** | ***OrderV3ApiModifyOrderV3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderV3ApiModifyOrderV3Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAlgoName** | **optional.**| Algorithm name, used to track and manage orders | 

### Return type

[**ModifyOrderV3Response**](ModifyOrderV3Response.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **PlaceGTTOrder**
> GttTriggerOrderResponse PlaceGTTOrder(ctx, body, optional)
Place GTT order

This API allows you to place GTT orders.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GttPlaceOrderRequest**](GttPlaceOrderRequest.md)|  | 
 **optional** | ***OrderV3ApiPlaceGTTOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderV3ApiPlaceGTTOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **origin** | **optional.**|  | 
 **xAlgoName** | **optional.**| Algorithm name, used to track and manage orders | 

### Return type

[**GttTriggerOrderResponse**](GttTriggerOrderResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **PlaceOrderV3**
> PlaceOrderV3Response PlaceOrderV3(ctx, body, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PlaceOrderV3Request**](PlaceOrderV3Request.md)|  | 
 **optional** | ***OrderV3ApiPlaceOrderV3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderV3ApiPlaceOrderV3Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAlgoName** | **optional.**| Algorithm name, used to track and manage orders | 

### Return type

[**PlaceOrderV3Response**](PlaceOrderV3Response.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

