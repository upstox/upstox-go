# UsOrderEstimateData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brokerage** | **float64** | Upstox brokerage | [optional] [default to null]
**Gst** | **float64** | GST on brokerage | [optional] [default to null]
**TradeValue** | **float64** | quantity x limit price, or quantity x last traded price for a market order (the request price when there is no last traded price) | [optional] [default to null]
**RegulatoryFee** | **float64** | Regulatory fee as charged by the order service | [optional] [default to null]
**TafFee** | **float64** | FINRA trading activity fee | [optional] [default to null]
**CatFee** | **float64** | Consolidated audit trail fee | [optional] [default to null]
**TotalCharges** | **float64** | Sum of the charges above | [optional] [default to null]
**EstimatedTotal** | **float64** | Buy: trade value plus charges. Sell: trade value minus charges. Null when the estimate carries no charges | [optional] [default to null]

[Back to Model list] [Back to API list] [Back to README]

