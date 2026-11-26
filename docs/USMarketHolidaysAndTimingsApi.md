# USMarketHolidaysAndTimingsApi

All URIs are relative to *https://api.upstox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsExchangeTimings**](USMarketHolidaysAndTimingsApi.md#GetUsExchangeTimings) | **Get** /v2/us/market/timings/{date} | Get US market timings on a date
[**GetUsHoliday**](USMarketHolidaysAndTimingsApi.md#GetUsHoliday) | **Get** /v2/us/market/holidays/{date} | Get US market holiday on a date
[**GetUsHolidays**](USMarketHolidaysAndTimingsApi.md#GetUsHolidays) | **Get** /v2/us/market/holidays | Get US market holidays
[**GetUsMarketStatus**](USMarketHolidaysAndTimingsApi.md#GetUsMarketStatus) | **Get** /v2/us/market/status | Get US market status

# **GetUsExchangeTimings**
> GetUsExchangeTimingResponse GetUsExchangeTimings(ctx, date)
Get US market timings on a date

Returns the US regular and overnight session times on the date, or an empty list when closed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **date** | **string**|  | 

### Return type

[**GetUsExchangeTimingResponse**](GetUsExchangeTimingResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetUsHoliday**
> GetUsHolidayResponse GetUsHoliday(ctx, date)
Get US market holiday on a date

Returns the US market holiday on the date, or an empty list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **date** | **string**|  | 

### Return type

[**GetUsHolidayResponse**](GetUsHolidayResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetUsHolidays**
> GetUsHolidayResponse GetUsHolidays(ctx, )
Get US market holidays

Returns the US market holidays of the current year.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetUsHolidayResponse**](GetUsHolidayResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetUsMarketStatus**
> GetMarketStatusResponse GetUsMarketStatus(ctx, )
Get US market status

Returns the live status of the US market.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetMarketStatusResponse**](GetMarketStatusResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

