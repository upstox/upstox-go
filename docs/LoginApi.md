# LoginApi

All URIs are relative to *https://api.upstox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Authorize**](LoginApi.md#Authorize) | **Get** /v2/login/authorization/dialog | Authorize API
[**GenerateTokenViaTotpAndPin**](LoginApi.md#GenerateTokenViaTotpAndPin) | **Post** /v2/totp-login/authorization/token | Generate token using TOTP and PIN
[**InitTokenRequestForIndieUser**](LoginApi.md#InitTokenRequestForIndieUser) | **Post** /v3/login/auth/token/request/{client_id} | Init token API
[**Logout**](LoginApi.md#Logout) | **Delete** /v2/logout | Logout
[**Token**](LoginApi.md#Token) | **Post** /v2/login/authorization/token | Get token API

# **Authorize**
> Authorize(ctx, clientId, redirectUri, optional)
Authorize API

This provides details on the login endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**|  | 
  **redirectUri** | **string**|  | 
 **optional** | ***LoginApiAuthorizeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoginApiAuthorizeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **state** | **optional.String**|  | 
 **scope** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GenerateTokenViaTotpAndPin**
> TokenResponse GenerateTokenViaTotpAndPin(ctx, body, xApiKey, optional)
Generate token using TOTP and PIN

This API generates an access token for an automation client using the app's API key together with the user's 6-digit PIN and 6-digit TOTP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TotpPinTokenRequest**](TotpPinTokenRequest.md)|  | 
  **xApiKey** | **string**| OAuth API key that identifies the app | 
 **optional** | ***LoginApiGenerateTokenViaTotpAndPinOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoginApiGenerateTokenViaTotpAndPinOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xUserId** | **optional.**| User id to log in. Required for multi-client apps (6 alphanumeric characters); optional for individual apps, but must match the app&#x27;s registered user when supplied | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*, application/json

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **InitTokenRequestForIndieUser**
> IndieUserInitTokenResponse InitTokenRequestForIndieUser(ctx, body, clientId)
Init token API

This API provides the initialize the generate token and it's expiry for an indie user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IndieUserTokenRequest**](IndieUserTokenRequest.md)|  | 
  **clientId** | **string**|  | 

### Return type

[**IndieUserInitTokenResponse**](IndieUserInitTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*, application/json

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **Logout**
> LogoutResponse Logout(ctx, )
Logout

Logout

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LogoutResponse**](LogoutResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **Token**
> TokenResponse Token(ctx, code, clientId, clientSecret, redirectUri, grantType, refreshExtendedToken)
Get token API

This API provides the functionality to obtain opaque token from authorization_code exchange and also provides the user’s profile in the same response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**|  | 
  **clientId** | **string**|  | 
  **clientSecret** | **string**|  | 
  **redirectUri** | **string**|  | 
  **grantType** | **string**|  | 
  **refreshExtendedToken** | **bool**|  | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: */*, application/json

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

