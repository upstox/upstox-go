# MarketHolidaysAndTimingsApi

All URIs are relative to *https://api.upstox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChangeOiData**](MarketHolidaysAndTimingsApi.md#GetChangeOiData) | **Get** /v2/market/change-oi | Get Change in OI Data
[**GetDiiData**](MarketHolidaysAndTimingsApi.md#GetDiiData) | **Get** /v2/market/dii | Get DII Data
[**GetExchangeTimings**](MarketHolidaysAndTimingsApi.md#GetExchangeTimings) | **Get** /v2/market/timings/{date} | Get Exchange Timings on particular date
[**GetFiiData**](MarketHolidaysAndTimingsApi.md#GetFiiData) | **Get** /v2/market/fii | Get FII Data
[**GetHoliday**](MarketHolidaysAndTimingsApi.md#GetHoliday) | **Get** /v2/market/holidays/{date} | Get Holiday on particular date
[**GetHolidays**](MarketHolidaysAndTimingsApi.md#GetHolidays) | **Get** /v2/market/holidays | Get Holiday list of current year
[**GetMarketStatus**](MarketHolidaysAndTimingsApi.md#GetMarketStatus) | **Get** /v2/market/status/{exchange} | Get Market status for particular exchange
[**GetMaxPainData**](MarketHolidaysAndTimingsApi.md#GetMaxPainData) | **Get** /v2/market/max-pain | Get Max Pain Data
[**GetOiData**](MarketHolidaysAndTimingsApi.md#GetOiData) | **Get** /v2/market/oi | Get OI Data
[**GetPcrData**](MarketHolidaysAndTimingsApi.md#GetPcrData) | **Get** /v2/market/pcr | Get PCR Data
[**GetSmartlistFutures**](MarketHolidaysAndTimingsApi.md#GetSmartlistFutures) | **Get** /v2/market/smartlist/futures | Get Smartlist Futures
[**GetSmartlistMtf**](MarketHolidaysAndTimingsApi.md#GetSmartlistMtf) | **Get** /v2/market/smartlist/mtf | Get Smartlist MTF
[**GetSmartlistOptions**](MarketHolidaysAndTimingsApi.md#GetSmartlistOptions) | **Get** /v2/market/smartlist/options | Get Smartlist Options

# **GetChangeOiData**
> AnalyticsResponse GetChangeOiData(ctx, instrumentKey, expiry, date, interval)
Get Change in OI Data

Fetches change in OI data for the given instrument key, expiry, date and interval.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instrumentKey** | **string**| Underlying asset instrument key | 
  **expiry** | **string**| Expiry date in YYYY-MM-DD format | 
  **date** | **string**| Date in YYYY-MM-DD format for which change OI data is required | 
  **interval** | **int32**| Number of days for which difference in OI is required | 

### Return type

[**AnalyticsResponse**](AnalyticsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetDiiData**
> AnalyticsResponse GetDiiData(ctx, dataType, interval, optional)
Get DII Data

Fetches DII activity data for the given interval.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataType** | **string**| Data type. Allowed value: NSE_EQ|CASH | 
  **interval** | **string**| Interval. Allowed values: 1D, 1M | 
 **optional** | ***MarketHolidaysAndTimingsApiGetDiiDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarketHolidaysAndTimingsApiGetDiiDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.String**| Start date in YYYY-MM-DD format (optional) | 

### Return type

[**AnalyticsResponse**](AnalyticsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetExchangeTimings**
> GetExchangeTimingResponse GetExchangeTimings(ctx, date)
Get Exchange Timings on particular date

This API provides the functionality to retrieve the exchange timings on particular date

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **date** | **string**|  | 

### Return type

[**GetExchangeTimingResponse**](GetExchangeTimingResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetFiiData**
> AnalyticsResponse GetFiiData(ctx, dataType, interval, optional)
Get FII Data

Fetches FII activity data for the given data type and interval.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataType** | **string**| Data type. Allowed values: NSE_FO|INDEX_FUTURES, NSE_FO|STOCK_FUTURES, NSE_FO|INDEX_OPTIONS, NSE_FO|STOCK_OPTIONS, NSE_EQ|CASH | 
  **interval** | **string**| Interval. Allowed values: 1D, 1M | 
 **optional** | ***MarketHolidaysAndTimingsApiGetFiiDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarketHolidaysAndTimingsApiGetFiiDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.String**| Start date in YYYY-MM-DD format (optional) | 

### Return type

[**AnalyticsResponse**](AnalyticsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetHoliday**
> GetHolidayResponse GetHoliday(ctx, date)
Get Holiday on particular date

This API provides the functionality to retrieve the holiday on particular date

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **date** | **string**|  | 

### Return type

[**GetHolidayResponse**](GetHolidayResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetHolidays**
> GetHolidayResponse GetHolidays(ctx, )
Get Holiday list of current year

This API provides the functionality to retrieve the holiday list of current year

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetHolidayResponse**](GetHolidayResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetMarketStatus**
> GetMarketStatusResponse GetMarketStatus(ctx, exchange)
Get Market status for particular exchange

This API provides the functionality to retrieve the market status for particular exchange

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **exchange** | **string**|  | 

### Return type

[**GetMarketStatusResponse**](GetMarketStatusResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetMaxPainData**
> AnalyticsResponse GetMaxPainData(ctx, instrumentKey, expiry, date, bucketInterval)
Get Max Pain Data

Fetches Max Pain data for the given instrument key, expiry, date and bucket interval.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instrumentKey** | **string**| Underlying asset instrument key | 
  **expiry** | **string**| Expiry date in YYYY-MM-DD format | 
  **date** | **string**| Date in YYYY-MM-DD format for which Max Pain data is required | 
  **bucketInterval** | **int32**| Bucket interval in minutes for the insights list | 

### Return type

[**AnalyticsResponse**](AnalyticsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetOiData**
> AnalyticsResponse GetOiData(ctx, instrumentKey, expiry, date)
Get OI Data

Fetches OI data for the given instrument key, expiry and date.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instrumentKey** | **string**| Underlying asset instrument key | 
  **expiry** | **string**| Expiry date in YYYY-MM-DD format | 
  **date** | **string**| Date in YYYY-MM-DD format for which OI data is required | 

### Return type

[**AnalyticsResponse**](AnalyticsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetPcrData**
> AnalyticsResponse GetPcrData(ctx, instrumentKey, expiry, date, bucketInterval)
Get PCR Data

Fetches PCR (Put-Call Ratio) data for the given instrument key, expiry, date and bucket interval.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instrumentKey** | **string**| Underlying asset instrument key | 
  **expiry** | **string**| Expiry date in YYYY-MM-DD format | 
  **date** | **string**| Date in YYYY-MM-DD format for which PCR data is required | 
  **bucketInterval** | **int32**| Bucket interval in minutes for the insights list | 

### Return type

[**AnalyticsResponse**](AnalyticsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetSmartlistFutures**
> AnalyticsResponse GetSmartlistFutures(ctx, optional)
Get Smartlist Futures

Fetches the smartlist of futures for the given asset type and category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MarketHolidaysAndTimingsApiGetSmartlistFuturesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarketHolidaysAndTimingsApiGetSmartlistFuturesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetType** | **optional.String**| Asset type. Allowed values: INDEX, STOCK, COMMODITY | 
 **category** | **optional.String**| Category. Allowed values: TOP_TRADED, MOST_ACTIVE, OI_GAINERS, OI_LOSERS, PRICE_GAINERS, PRICE_LOSERS, PREMIUM, DISCOUNT (COMMODITY supports only: TOP_TRADED, MOST_ACTIVE, OI_GAINERS, OI_LOSERS) | 
 **pageNumber** | **optional.Int32**| Page number (optional, 1-indexed) | 
 **pageSize** | **optional.Int32**| Page size (optional, max 50) | 

### Return type

[**AnalyticsResponse**](AnalyticsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetSmartlistMtf**
> AnalyticsResponse GetSmartlistMtf(ctx, optional)
Get Smartlist MTF

Fetches the smartlist of MTF (Margin Trade Funding) stocks. Prices are enriched with live LTP data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MarketHolidaysAndTimingsApiGetSmartlistMtfOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarketHolidaysAndTimingsApiGetSmartlistMtfOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Page number (optional, 1-indexed) | 
 **pageSize** | **optional.Int32**| Page size (optional, max 50) | 

### Return type

[**AnalyticsResponse**](AnalyticsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetSmartlistOptions**
> AnalyticsResponse GetSmartlistOptions(ctx, optional)
Get Smartlist Options

Fetches the smartlist of options for the given asset type and category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MarketHolidaysAndTimingsApiGetSmartlistOptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarketHolidaysAndTimingsApiGetSmartlistOptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetType** | **optional.String**| Asset type. Allowed values: INDEX, STOCK, COMMODITY | 
 **category** | **optional.String**| Category. Allowed values: TOP_TRADED, MOST_ACTIVE, OI_GAINERS, OI_LOSERS, PRICE_GAINERS, PRICE_LOSERS, IV_GAINERS, IV_LOSERS, UNDER_5000, UNDER_10000 (COMMODITY supports only: TOP_TRADED, MOST_ACTIVE, OI_GAINERS, OI_LOSERS) | 
 **pageNumber** | **optional.Int32**| Page number (optional, 1-indexed) | 
 **pageSize** | **optional.Int32**| Page size (optional, max 50) | 

### Return type

[**AnalyticsResponse**](AnalyticsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

