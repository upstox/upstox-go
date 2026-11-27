# FundamentalsApi

All URIs are relative to *https://api.upstox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBalanceSheet**](FundamentalsApi.md#GetBalanceSheet) | **Get** /v2/fundamentals/{isin}/balance-sheet | Get Balance Sheet
[**GetCashFlow**](FundamentalsApi.md#GetCashFlow) | **Get** /v2/fundamentals/{isin}/cash-flow | Get Cash Flow
[**GetCompanyProfile**](FundamentalsApi.md#GetCompanyProfile) | **Get** /v2/fundamentals/{isin}/profile | Get Company Profile
[**GetCompetitors**](FundamentalsApi.md#GetCompetitors) | **Get** /v2/fundamentals/{instrument_key}/competitors | Get Competitors
[**GetCorporateActions**](FundamentalsApi.md#GetCorporateActions) | **Get** /v2/fundamentals/{isin}/corporate-actions | Get Corporate Actions
[**GetIncomeStatement**](FundamentalsApi.md#GetIncomeStatement) | **Get** /v2/fundamentals/{isin}/income-statement | Get Income Statement
[**GetKeyRatios**](FundamentalsApi.md#GetKeyRatios) | **Get** /v2/fundamentals/{isin}/key-ratios | Get Key Ratios
[**GetShareHoldings**](FundamentalsApi.md#GetShareHoldings) | **Get** /v2/fundamentals/{isin}/share-holdings | Get Share Holdings

# **GetBalanceSheet**
> BalanceSheetResponse GetBalanceSheet(ctx, isin, optional)
Get Balance Sheet

Fetches balance sheet data for the given ISIN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isin** | **string**| ISIN of the company | 
 **optional** | ***FundamentalsApiGetBalanceSheetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundamentalsApiGetBalanceSheetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| Report type. Allowed values: consolidated or standalone | [default to consolidated]
 **fs** | **optional.Bool**| Whether to include the full financial statement | [default to false]

### Return type

[**BalanceSheetResponse**](BalanceSheetResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetCashFlow**
> CashFlowResponse GetCashFlow(ctx, isin, optional)
Get Cash Flow

Fetches cash flow data for the given ISIN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isin** | **string**| ISIN of the company | 
 **optional** | ***FundamentalsApiGetCashFlowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundamentalsApiGetCashFlowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| Report type. Allowed values: consolidated, standalone | [default to consolidated]
 **fs** | **optional.Bool**| Whether to include the full financial statement | [default to false]

### Return type

[**CashFlowResponse**](CashFlowResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetCompanyProfile**
> CompanyProfileResponse GetCompanyProfile(ctx, isin)
Get Company Profile

Fetches company profile information for the given ISIN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isin** | **string**| ISIN of the company | 

### Return type

[**CompanyProfileResponse**](CompanyProfileResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetCompetitors**
> CompetitorsResponse GetCompetitors(ctx, instrumentKey)
Get Competitors

Fetches the list of competitors for the given instrument_key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instrumentKey** | **string**| Instrument key of the company | 

### Return type

[**CompetitorsResponse**](CompetitorsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetCorporateActions**
> CorporateActionsResponse GetCorporateActions(ctx, isin)
Get Corporate Actions

Fetches corporate actions for the given ISIN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isin** | **string**| ISIN of the company | 

### Return type

[**CorporateActionsResponse**](CorporateActionsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetIncomeStatement**
> IncomeStatementResponse GetIncomeStatement(ctx, isin, optional)
Get Income Statement

Fetches income statement data for the given ISIN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isin** | **string**| ISIN of the company | 
 **optional** | ***FundamentalsApiGetIncomeStatementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundamentalsApiGetIncomeStatementOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| Report type. Allowed values: consolidated or standalone | [default to consolidated]
 **timePeriod** | **optional.String**| Time period. Allowed values: yearly, quarterly | [default to yearly]
 **fs** | **optional.Bool**| Whether to include the full financial statement | [default to false]

### Return type

[**IncomeStatementResponse**](IncomeStatementResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetKeyRatios**
> KeyRatiosResponse GetKeyRatios(ctx, isin)
Get Key Ratios

Fetches key financial ratios for the given ISIN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isin** | **string**| ISIN of the company | 

### Return type

[**KeyRatiosResponse**](KeyRatiosResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

# **GetShareHoldings**
> ShareHoldingsResponse GetShareHoldings(ctx, isin)
Get Share Holdings

Fetches shareholding pattern for the given ISIN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isin** | **string**| ISIN of the company | 

### Return type

[**ShareHoldingsResponse**](ShareHoldingsResponse.md)

### Authorization

OAUTH2

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [Back to API list] [Back to Model list] [Back to README]

