# IpoListingData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | IPO id (slug). Pass this as the id to the IPO details and apply APIs | [optional] [default to null]
**Symbol** | **string** | Trading symbol of the issuer | [optional] [default to null]
**Name** | **string** | Name of the issuer | [optional] [default to null]
**Status** | **string** | Current stage of the IPO | [optional] [default to null]
**Isin** | **string** | ISIN of the issue | [optional] [default to null]
**IssueType** | **string** | Issue type of the IPO. &#x60;regular&#x60; is a mainboard issue | [optional] [default to null]
**IssueSize** | **float64** | Total issue size in crores | [optional] [default to null]
**Industry** | **string** | Industry the issuer operates in | [optional] [default to null]
**MinimumPrice** | **float64** | Lower end of the price band, per share | [optional] [default to null]
**MaximumPrice** | **float64** | Upper end of the price band, per share | [optional] [default to null]
**BiddingStartDate** | **string** | Date bidding opens, in yyyy-MM-dd | [optional] [default to null]
**BiddingEndDate** | **string** | Date bidding closes, in yyyy-MM-dd | [optional] [default to null]
**TotalSubscription** | **string** | Overall subscription so far; null until the issue opens | [optional] [default to null]
**Investors** | [**[]IpoInvestorType**](IpoInvestorType.md) | Investor categories the issue accepts. Empty when the listing data carries none — the IPO details API is the authoritative source | [optional] [default to null]

[Back to Model list] [Back to API list] [Back to README]

