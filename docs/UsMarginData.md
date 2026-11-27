# UsMarginData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Margins** | [**[]UsOrderMargin**](UsOrderMargin.md) | One entry per order, in request order | [optional] [default to null]
**Shortfall** | **float64** | required_margin minus available_to_invest, never below 0 | [optional] [default to null]
**Sufficient** | **bool** | Whether available_to_invest covers required_margin | [optional] [default to null]
**RequiredMargin** | **float64** | Sum of required margin across the orders | [optional] [default to null]
**AvailableToInvest** | **float64** | Buying power in the US wallet | [optional] [default to null]

[Back to Model list] [Back to API list] [Back to README]

