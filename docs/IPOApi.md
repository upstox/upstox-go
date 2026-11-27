# IPOApi

All URIs are relative to *https://api.upstox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyForIpo**](IPOApi.md#ApplyForIpo) | **Post** /v2/ipos/orders | Apply for IPO
[**CancelIpoOrder**](IPOApi.md#CancelIpoOrder) | **Delete** /v2/ipos/orders/{order_id} | Cancel IPO Order
[**GetIpoDetails**](IPOApi.md#GetIpoDetails) | **Get** /v2/ipos/{id} | Get IPO Details
[**GetIpoListing**](IPOApi.md#GetIpoListing) | **Get** /v2/ipos | Get IPO Listing
[**GetIpoOrderById**](IPOApi.md#GetIpoOrderById) | **Get** /v2/ipos/orders/{order_id} | Get IPO Order
[**GetIpoOrders**](IPOApi.md#GetIpoOrders) | **Get** /v2/ipos/orders | Get IPO Orders

# **ApplyForIpo**
> IpoApplyResponse ApplyForIpo(ctx, body)
Apply for IPO

Places an IPO application for the authenticated user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpoApplyRequest**](IpoApplyRequest.md)|  | 

### Return type

[**IpoApplyResponse**](IpoApplyResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **CancelIpoOrder**
> IpoCancelResponse CancelIpoOrder(ctx, orderId)
Cancel IPO Order

Cancels/deletes an IPO order of the authenticated user by order id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| IPO application id, as returned in &#x60;order_id&#x60; by the apply and orders APIs | 

### Return type

[**IpoCancelResponse**](IpoCancelResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetIpoDetails**
> IpoDetailsResponse GetIpoDetails(ctx, id)
Get IPO Details

Fetches detailed information for an IPO by its slug ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| IPO id (slug), as returned in &#x60;id&#x60; by the IPO listing API | 

### Return type

[**IpoDetailsResponse**](IpoDetailsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetIpoListing**
> IpoListingResponse GetIpoListing(ctx, optional)
Get IPO Listing

Fetches list of IPOs filtered by status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPOApiGetIpoListingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPOApiGetIpoListingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **optional.String**| IPO status to filter by | 
 **issueType** | **optional.String**| Issue type filter. &#x60;regular&#x60; is a mainboard issue. Returns both when omitted | 
 **pageNumber** | **optional.String**| Page number, starting at 1 | 
 **records** | **optional.String**| Number of records per page | 

### Return type

[**IpoListingResponse**](IpoListingResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetIpoOrderById**
> IpoOrderDetailResponse GetIpoOrderById(ctx, orderId)
Get IPO Order

Fetches a single IPO order of the authenticated user by order id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| IPO application id, as returned in &#x60;order_id&#x60; by the apply and orders APIs | 

### Return type

[**IpoOrderDetailResponse**](IpoOrderDetailResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetIpoOrders**
> IpoOrderResponse GetIpoOrders(ctx, optional)
Get IPO Orders

Fetches the authenticated user's IPO orders/applications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPOApiGetIpoOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPOApiGetIpoOrdersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.String**| Page number, starting at 1 | 
 **records** | **optional.String**| Number of records per page | 

### Return type

[**IpoOrderResponse**](IpoOrderResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

