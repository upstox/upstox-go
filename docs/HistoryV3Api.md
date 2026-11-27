# HistoryV3Api

All URIs are relative to *https://api.upstox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistoricalCandleData**](HistoryV3Api.md#GetHistoricalCandleData) | **Get** /v3/historical-candle/{instrumentKey}/{unit}/{interval}/{to_date} | Historical candle data
[**GetHistoricalCandleDataRange**](HistoryV3Api.md#GetHistoricalCandleDataRange) | **Get** /v3/historical-candle/{instrumentKey}/{unit}/{interval}/{to_date}/{from_date} | Historical candle data
[**GetIntraDayCandleData**](HistoryV3Api.md#GetIntraDayCandleData) | **Get** /v3/historical-candle/intraday/{instrumentKey}/{unit}/{interval} | Intra day candle data

# **GetHistoricalCandleData**
> GetHistoricalCandleResponse GetHistoricalCandleData(ctx, instrumentKey, unit, interval, toDate)
Historical candle data

Get OHLC values for all instruments for the present trading day with expanded interval options.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instrumentKey** | **string**|  | 
  **unit** | **string**|  | 
  **interval** | **int32**|  | 
  **toDate** | **string**|  | 

### Return type

[**GetHistoricalCandleResponse**](GetHistoricalCandleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetHistoricalCandleDataRange**
> GetHistoricalCandleResponse GetHistoricalCandleDataRange(ctx, instrumentKey, unit, interval, toDate, fromDate)
Historical candle data

Get OHLC values for all instruments for the present trading day with expanded interval options

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instrumentKey** | **string**|  | 
  **unit** | **string**|  | 
  **interval** | **int32**|  | 
  **toDate** | **string**|  | 
  **fromDate** | **string**|  | 

### Return type

[**GetHistoricalCandleResponse**](GetHistoricalCandleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetIntraDayCandleData**
> GetIntraDayCandleResponse GetIntraDayCandleData(ctx, instrumentKey, unit, interval)
Intra day candle data

Get OHLC values for all instruments for the present trading day

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instrumentKey** | **string**|  | 
  **unit** | **string**|  | 
  **interval** | **int32**|  | 

### Return type

[**GetIntraDayCandleResponse**](GetIntraDayCandleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

