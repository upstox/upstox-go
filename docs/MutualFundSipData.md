# MutualFundSipData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fund** | **string** | Display name of the fund. | [optional] [default to null]
**Status** | **string** | ACTIVE, PAUSED, or CANCELLED. | [optional] [default to null]
**Created** | **string** | Timestamp when the SIP was registered. | [optional] [default to null]
**Frequency** | **string** | Frequency (e.g. monthly, weekly). | [optional] [default to null]
**Instalments** | **int64** | Total number of instalments when applicable; -1 if open-ended. | [optional] [default to null]
**SipId** | **string** | Unique SIP identifier. | [optional] [default to null]
**SipRegNumber** | **string** | SIP registration number when available. | [optional] [default to null]
**InstrumentKey** | **string** | ISIN of the fund, when available. | [optional] [default to null]
**DividendType** | **string** | Dividend option when available (e.g. growth, idcw). | [optional] [default to null]
**TransactionType** | **string** | BUY or SELL. | [optional] [default to null]
**NextInstalment** | **string** | Next instalment date. | [optional] [default to null]
**InstalmentAmount** | **float64** | Amount per instalment. | [optional] [default to null]
**LastInstalment** | **string** | Last triggered instalment timestamp. | [optional] [default to null]
**PendingInstalments** | **int64** | Pending instalments when applicable. | [optional] [default to null]
**InstalmentDay** | **int64** | Day of month for monthly SIPs when applicable. | [optional] [default to null]
**CompletedInstalments** | **int64** | Completed instalment count when available. | [optional] [default to null]
**TriggerPrice** | **float64** | Trigger price when applicable. | [optional] [default to null]
**SipType** | **string** | SIP channel/type (e.g. sip, amc_sip). | [optional] [default to null]

[Back to Model list] [Back to API list] [Back to README]

