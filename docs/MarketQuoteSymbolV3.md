# MarketQuoteSymbolV3

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ohlc** | [***OhlcV3**](OhlcV3.md) |  | [optional] [default to null]
**Depth** | [***DepthMap**](DepthMap.md) |  | [optional] [default to null]
**Timestamp** | **string** | Time in milliseconds at which the feeds was updated | [optional] [default to null]
**InstrumentToken** | **string** | Key issued by Upstox for the instrument | [optional] [default to null]
**Symbol** | **string** | Shows the trading symbol of the instrument | [optional] [default to null]
**LastPrice** | **float64** | The last traded price of symbol | [optional] [default to null]
**Volume** | **int64** | The volume traded today on symbol | [optional] [default to null]
**AveragePrice** | **float64** | Average price | [optional] [default to null]
**Oi** | **float64** | Total number of outstanding contracts held by market participants exchange-wide (only F&amp;O) | [optional] [default to null]
**NetChange** | **float64** | The absolute change from yesterday&#x27;s close to last traded price | [optional] [default to null]
**TotalBuyQuantity** | **float64** | The total number of bid quantity available for trading | [optional] [default to null]
**TotalSellQuantity** | **float64** | The total number of ask quantity available for trading | [optional] [default to null]
**LowerCircuitLimit** | **float64** | The lower circuit of symbol | [optional] [default to null]
**UpperCircuitLimit** | **float64** | The upper circuit of symbol | [optional] [default to null]
**LastTradeTime** | **string** | Time in milliseconds at which last trade happened | [optional] [default to null]
**OiDayHigh** | **float64** | The highest open interest recorded on symbol during the day | [optional] [default to null]
**OiDayLow** | **float64** | The lowest open interest recorded on symbol during the day | [optional] [default to null]
**PrevClosePrice** | **float64** | The close price of the symbol from the previous session of trading | [optional] [default to null]
**YearHigh** | **float64** | The highest price of the symbol over the trailing year | [optional] [default to null]
**YearLow** | **float64** | The lowest price of the symbol over the trailing year | [optional] [default to null]
**PreviousOi** | **float64** | The open interest of the symbol from the previous session (only F&amp;O) | [optional] [default to null]
**IndicativeEquilibriumPrice** | **float64** | The indicative equilibrium price discovered during the call auction session | [optional] [default to null]
**ReferencePrice** | **float64** | The reference price against which the call auction price band is applied | [optional] [default to null]
**IndicativeEquilibriumQuantity** | **int64** | The indicative equilibrium quantity matched at the indicative equilibrium price | [optional] [default to null]
**IndicativeImbalanceQuantityTotal** | **int64** | The total indicative imbalance quantity in the call auction order book | [optional] [default to null]
**IndicativeImbalanceQuantityMarket** | **int64** | The market-order component of the indicative imbalance quantity in the call auction order book | [optional] [default to null]
**CasEligible** | **bool** | Whether the symbol is eligible for the call auction session | [optional] [default to null]

[Back to Model list] [Back to API list] [Back to README]

