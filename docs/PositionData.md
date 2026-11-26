# PositionData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | **string** | Exchange to which the order is associated | [optional] [default to null]
**Multiplier** | **float64** | The quantity/lot size multiplier used for calculating P&amp;Ls | [optional] [default to null]
**Value** | **float64** | Net value of the position | [optional] [default to null]
**Pnl** | **float64** | Profit and loss - net returns on the position | [optional] [default to null]
**Product** | **string** | Shows if the order was either Intraday, Delivery, CO or OCO | [optional] [default to null]
**InstrumentToken** | **string** | Key issued by Upstox for the instrument | [optional] [default to null]
**AveragePrice** | **float64** | Average price at which the net position quantity was acquired | [optional] [default to null]
**BuyValue** | **float64** | Net value of the bought quantities | [optional] [default to null]
**OvernightQuantity** | **int32** | Quantity held previously and carried forward over night | [optional] [default to null]
**DayBuyValue** | **float64** | Amount at which the quantity is bought during the day | [optional] [default to null]
**DayBuyPrice** | **float64** | Average price at which the day qty was bought. Default is empty string | [optional] [default to null]
**OvernightBuyAmount** | **float64** | Amount at which the quantity was bought in the previous session | [optional] [default to null]
**OvernightBuyQuantity** | **int32** | Quantity bought in the previous session | [optional] [default to null]
**DayBuyQuantity** | **int32** | Quantity bought during the day | [optional] [default to null]
**DaySellValue** | **float64** | Amount at which the quantity is sold during the day | [optional] [default to null]
**DaySellPrice** | **float64** | Average price at which the day quantity was sold | [optional] [default to null]
**OvernightSellAmount** | **float64** | Amount at which the quantity was sold in the previous session | [optional] [default to null]
**OvernightSellQuantity** | **int32** | Quantity sold short in the previous session | [optional] [default to null]
**DaySellQuantity** | **int32** | Quantity sold during the day | [optional] [default to null]
**Quantity** | **int32** | Quantity left after nullifying Day and CF buy quantity towards Day and CF sell quantity | [optional] [default to null]
**LastPrice** | **float64** | Last traded market price of the instrument | [optional] [default to null]
**Unrealised** | **float64** | Day PnL generated against open positions | [optional] [default to null]
**Realised** | **float64** | Day PnL generated against closed positions | [optional] [default to null]
**SellValue** | **float64** | Net value of the sold quantities | [optional] [default to null]
**TradingSymbol** | **string** | Shows the trading symbol of the instrument | [optional] [default to null]
**ClosePrice** | **float64** | Closing price of the instrument from the last trading day | [optional] [default to null]
**BuyPrice** | **float64** | Average price at which quantities were bought | [optional] [default to null]
**SellPrice** | **float64** | Average price at which quantities were sold | [optional] [default to null]

[Back to Model list] [Back to API list] [Back to README]

