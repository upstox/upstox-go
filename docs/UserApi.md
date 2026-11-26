# UserApi

All URIs are relative to *https://api.upstox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPayout**](UserApi.md#CancelPayout) | **Delete** /v2/user/payments/payout/{transaction_id} | 
[**GetKillSwitch**](UserApi.md#GetKillSwitch) | **Get** /v2/user/kill-switch | Get kill switch status
[**GetPayinHistory**](UserApi.md#GetPayinHistory) | **Get** /v2/user/payments/payin | 
[**GetPayoutHistory**](UserApi.md#GetPayoutHistory) | **Get** /v2/user/payments/payout | 
[**GetPayoutModes**](UserApi.md#GetPayoutModes) | **Get** /v2/user/payments/payout/modes | 
[**GetProfile**](UserApi.md#GetProfile) | **Get** /v2/user/profile | Get profile
[**GetUserFundMargin**](UserApi.md#GetUserFundMargin) | **Get** /v2/user/get-funds-and-margin | Get User Fund And Margin
[**GetUserFundMarginV3**](UserApi.md#GetUserFundMarginV3) | **Get** /v3/user/get-funds-and-margin | Get User Fund And Margin
[**GetUserIps**](UserApi.md#GetUserIps) | **Get** /v2/user/ip | 
[**InitiatePayout**](UserApi.md#InitiatePayout) | **Post** /v2/user/payments/payout | 
[**ModifyPayout**](UserApi.md#ModifyPayout) | **Put** /v2/user/payments/payout/{transaction_id} | 
[**UpdateKillSwitch**](UserApi.md#UpdateKillSwitch) | **Post** /v2/user/kill-switch | Update kill switch
[**UpdateUserIp**](UserApi.md#UpdateUserIp) | **Put** /v2/user/ip | 

# **CancelPayout**
> PayoutDetailsResponse CancelPayout(ctx, transactionId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transactionId** | **string**| Payout transaction id | 

### Return type

[**PayoutDetailsResponse**](PayoutDetailsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetKillSwitch**
> KillSwitchResponse GetKillSwitch(ctx, )
Get kill switch status

Returns the disable/enable status of all trading segments for the user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**KillSwitchResponse**](KillSwitchResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetPayinHistory**
> PaymentHistoryResponse GetPayinHistory(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PaymentHistoryResponse**](PaymentHistoryResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetPayoutHistory**
> PaymentHistoryResponse GetPayoutHistory(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PaymentHistoryResponse**](PaymentHistoryResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetPayoutModes**
> PayoutModesResponse GetPayoutModes(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PayoutModesResponse**](PayoutModesResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetProfile**
> GetProfileResponse GetProfile(ctx, )
Get profile

This API allows to fetch the complete information of the user who is logged in including the products, order types and exchanges enabled for the user

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetProfileResponse**](GetProfileResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetUserFundMargin**
> GetUserFundMarginResponse GetUserFundMargin(ctx, optional)
Get User Fund And Margin

Shows the balance of the user in equity and commodity market.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiGetUserFundMarginOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetUserFundMarginOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **segment** | **optional.String**|  | 

### Return type

[**GetUserFundMarginResponse**](GetUserFundMarginResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetUserFundMarginV3**
> GetUserFundMarginV3Response GetUserFundMarginV3(ctx, )
Get User Fund And Margin

Shows the detailed balance breakdown of the user including cash, pledged margin, available-to-trade, and unavailable-to-trade components.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetUserFundMarginV3Response**](GetUserFundMarginV3Response.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetUserIps**
> UserIpResponse GetUserIps(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UserIpResponse**](UserIpResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **InitiatePayout**
> PayoutDetailsResponse InitiatePayout(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InitiatePayoutRequest**](InitiatePayoutRequest.md)|  | 

### Return type

[**PayoutDetailsResponse**](PayoutDetailsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **ModifyPayout**
> PayoutDetailsResponse ModifyPayout(ctx, body, transactionId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModifyPayoutRequest**](ModifyPayoutRequest.md)|  | 
  **transactionId** | **string**| Payout transaction id | 

### Return type

[**PayoutDetailsResponse**](PayoutDetailsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **UpdateKillSwitch**
> KillSwitchResponse UpdateKillSwitch(ctx, body)
Update kill switch

Enable or disable specific trading segments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]KillSwitchSegmentUpdateRequest**](KillSwitchSegmentUpdateRequest.md)|  | 

### Return type

[**KillSwitchResponse**](KillSwitchResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*, application/json

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **UpdateUserIp**
> UserIpResponse UpdateUserIp(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateUserIpRequest**](UpdateUserIpRequest.md)|  | 

### Return type

[**UserIpResponse**](UserIpResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

