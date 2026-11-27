# UsOrderEstimateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | **float64** | Shares, up to 9 decimals; fractions only on fractionable stocks | [default to null]
**Price** | **float64** | Limit price in USD; required for a limit order. Optional for a market order, which is estimated at the last traded price and falls back to this price | [optional] [default to null]
**InstrumentKey** | **string** | US instrument key | [default to null]
**TransactionType** | **string** |  | [default to null]
**OrderType** | **string** |  | [default to null]

[Back to Model list] [Back to API list] [Back to README]

