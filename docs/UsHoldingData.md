# UsHoldingData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | **string** |  | [optional] [default to null]
**Quantity** | **float64** | Shares held | [optional] [default to null]
**Pnl** | **float64** | Unrealised P&amp;L in USD: current_value minus invested_value, negated for a short | [optional] [default to null]
**InstrumentKey** | **string** | US instrument key | [optional] [default to null]
**TradingSymbol** | **string** |  | [optional] [default to null]
**CompanyName** | **string** |  | [optional] [default to null]
**AveragePrice** | **float64** | Average buy price in USD | [optional] [default to null]
**AvailableQuantity** | **float64** | Shares free to sell | [optional] [default to null]
**TransactionType** | **string** | Position direction as sent by the holdings service: long or short | [optional] [default to null]
**InvestedValue** | **float64** | quantity x average_price, in USD | [optional] [default to null]
**LastPrice** | **float64** | Last traded price in USD; null when market data is unavailable | [optional] [default to null]
**PreviousClose** | **float64** | Previous close in USD; null when market data is unavailable | [optional] [default to null]
**CurrentValue** | **float64** | quantity x last_price, in USD | [optional] [default to null]
**DayPnl** | **float64** | Today&#x27;s P&amp;L in USD: quantity x (last_price minus previous_close), negated for a short | [optional] [default to null]

[Back to Model list] [Back to API list] [Back to README]

