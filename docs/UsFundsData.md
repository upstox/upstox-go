# UsFundsData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cash** | **float64** | Cash in the wallet | [optional] [default to null]
**AvailableToTrade** | **float64** | Buying power; includes unsettled instant-funding credit | [optional] [default to null]
**AvailableToWithdraw** | **float64** | Settled cash that can be withdrawn | [optional] [default to null]
**UnsettledCash** | **float64** | Cash not yet settled: cash minus available_to_withdraw | [optional] [default to null]
**ExchangeRate** | **float64** | Indicative USD to INR rate, for display only | [optional] [default to null]
**AvailableToTradeInr** | **float64** | available_to_trade x exchange_rate | [optional] [default to null]
**AvailableToWithdrawInr** | **float64** | available_to_withdraw x exchange_rate | [optional] [default to null]
**CashInr** | **float64** | cash x exchange_rate | [optional] [default to null]
**UnsettledCashInr** | **float64** | unsettled_cash x exchange_rate | [optional] [default to null]

[Back to Model list] [Back to API list] [Back to README]

