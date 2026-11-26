# MutualFundOrderData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Current status of the order. | [optional] [default to null]
**Folio** | **string** | Folio number, when allotted. | [optional] [default to null]
**Fund** | **string** | Display name of the fund. | [optional] [default to null]
**Amount** | **float64** | Order amount. | [optional] [default to null]
**Variety** | **string** | Order variety or channel-specific classification, if available. | [optional] [default to null]
**Quantity** | **float64** | Units allotted or sold. | [optional] [default to null]
**Price** | **float64** | Price or NAV used for the transaction, when available. | [optional] [default to null]
**OrderId** | **string** | Unique order identifier. | [optional] [default to null]
**ExchangeOrderId** | **string** | Exchange-generated order id, if available. | [optional] [default to null]
**InstrumentKey** | **string** | ISIN of the fund, when available. | [optional] [default to null]
**StatusMessage** | **string** | Human-readable status or rejection reason. | [optional] [default to null]
**OrderTimestamp** | **string** | Timestamp when the order was registered. | [optional] [default to null]
**ExchangeTimestamp** | **string** | Exchange date for the order, if available. | [optional] [default to null]
**SettlementId** | **string** | Exchange settlement id, if available. | [optional] [default to null]
**TransactionType** | **string** | BUY or SELL. | [optional] [default to null]
**PurchaseType** | **string** | FRESH or ADDITIONAL for buys, when available. | [optional] [default to null]
**LastPrice** | **float64** | Last known NAV. | [optional] [default to null]
**AveragePrice** | **float64** | Allotted or sold NAV. | [optional] [default to null]
**LastPriceDate** | **string** | Date for which last NAV is available. | [optional] [default to null]

[Back to Model list] [Back to API list] [Back to README]

