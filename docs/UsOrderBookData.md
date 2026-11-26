# UsOrderBookData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [optional] [default to null]
**Exchange** | **string** |  | [optional] [default to null]
**Validity** | **string** |  | [optional] [default to null]
**Price** | **float64** | Limit price in USD | [optional] [default to null]
**Quantity** | **float64** | Ordered shares | [optional] [default to null]
**Amount** | **float64** | USD amount for an amount order | [optional] [default to null]
**Status** | **string** | Order status, normalised from the US order book; an open order is scheduled while the market is closed | [optional] [default to null]
**Fractionable** | **bool** | Whether the stock allows fractional quantities; null when the stock is not in the US instrument list | [optional] [default to null]
**Variety** | **string** |  | [optional] [default to null]
**Tag** | **string** |  | [optional] [default to null]
**OrderId** | **string** | Order id | [optional] [default to null]
**InstrumentKey** | **string** | US instrument key | [optional] [default to null]
**TradingSymbol** | **string** |  | [optional] [default to null]
**OrderType** | **string** |  | [optional] [default to null]
**TransactionType** | **string** |  | [optional] [default to null]
**AveragePrice** | **float64** | Average traded price in USD | [optional] [default to null]
**FilledQuantity** | **float64** | Traded shares | [optional] [default to null]
**PendingQuantity** | **float64** | Shares still to be traded | [optional] [default to null]
**StatusMessage** | **string** |  | [optional] [default to null]
**StatusMessageRaw** | **string** |  | [optional] [default to null]
**OrderTimestamp** | **string** | When the order was placed, IST, yyyy-MM-dd HH:mm:ss | [optional] [default to null]
**ExchangeTimestamp** | **string** | When the order was last updated, IST, yyyy-MM-dd HH:mm:ss | [optional] [default to null]
**IsAmo** | **bool** |  | [optional] [default to null]

[Back to Model list] [Back to API list] [Back to README]

