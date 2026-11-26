# IpoDetailsData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | IPO id (slug). Pass this as the id to the apply API | [optional] [default to null]
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
**DailyStartTime** | **string** | Time bidding opens each day, in HH:mm IST | [optional] [default to null]
**DailyEndTime** | **string** | Time bidding closes each day, in HH:mm IST | [optional] [default to null]
**FaceValue** | **float64** | Face value of one share | [optional] [default to null]
**TickSize** | **float64** | Minimum price increment for a bid | [optional] [default to null]
**LotSize** | **int32** | Number of shares in one lot | [optional] [default to null]
**MinimumQuantity** | **int32** | Smallest quantity that can be bid for; bids must be in multiples of the lot size | [optional] [default to null]
**CutOffPrice** | **float64** | Cut-off price for retail applications | [optional] [default to null]
**ListingPrice** | **float64** | Price the share listed at; null until listing | [optional] [default to null]
**ListingExchange** | **string** | Exchange(s) the issue lists on. Comma-separated when the issue lists on more than one | [optional] [default to null]
**RhpUrl** | **string** | Link to the Red Herring Prospectus | [optional] [default to null]
**DrhpUrl** | **string** | Link to the Draft Red Herring Prospectus | [optional] [default to null]
**Timeline** | [***IpoTimeline**](IpoTimeline.md) |  | [optional] [default to null]
**RegistrarInfo** | [***IpoRegistrarInfo**](IpoRegistrarInfo.md) |  | [optional] [default to null]
**TotalSubscription** | **string** | Overall subscription so far; null until the issue opens | [optional] [default to null]
**Investors** | [**[]IpoInvestorType**](IpoInvestorType.md) | Investor categories the issue accepts | [optional] [default to null]

[Back to Model list] [Back to API list] [Back to README]

