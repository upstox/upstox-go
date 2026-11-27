# IncomeStatementData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | Report type: consolidated or standalone | [optional] [default to null]
**TimePeriod** | **string** | Time period: yearly or quarterly | [optional] [default to null]
**UnitsIn** | **string** | Unit of monetary values | [optional] [default to null]
**IncomeStatement** | [**[]IncomeStatementEntry**](IncomeStatementEntry.md) | Income statement entries grouped by type (Revenue, Operating profit, Net profit) | [optional] [default to null]
**FullStatement** | [**[]FinancialStatementEntry**](FinancialStatementEntry.md) | List of financial statement income statement items with historical values | [optional] [default to null]

[Back to Model list] [Back to API list] [Back to README]

