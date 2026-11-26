# IpoApplyRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | IPO id (slug) to apply for, as returned in &#x60;id&#x60; by the IPO listing and details APIs | [default to null]
**Upi** | **string** | UPI id used to block the application amount | [default to null]
**Category** | **string** | Investor category to apply under. Must be one the issue accepts — see &#x60;investors[].category&#x60; in the IPO details API | [default to null]
**Bids** | [**[]IpoBidRequest**](IpoBidRequest.md) | List of bids for the application (1 to 3 bids) | [default to null]

[Back to Model list] [Back to API list] [Back to README]

