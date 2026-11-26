# IpoOrderData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Reference id of the IPO the application was placed against | [optional] [default to null]
**Symbol** | **string** | Trading symbol of the issuer | [optional] [default to null]
**Exchange** | **string** | Exchange the application was submitted to. Casing follows the exchange feed and may vary | [optional] [default to null]
**RequestId** | **string** | Id of the request that created the application | [optional] [default to null]
**OrderId** | **string** | Application id. Pass this as order_id to the get-order and cancel-order APIs | [optional] [default to null]
**Status** | **string** | Lifecycle stage of the application. Known values: awaiting_mandate, ipo_allotted, ipo_not_allotted, application_deleted | [optional] [default to null]
**OrderStatus** | **string** | Outcome of the application. Known values: success, allotted, not_allotted | [optional] [default to null]
**PaymentStatus** | **string** | State of the UPI mandate backing the application. Known values: pending, mandate_accepted | [optional] [default to null]
**Category** | **string** | Investor category the application was placed under. IND (individual) and HNI can be applied for; EMP appears on employee-quota applications | [optional] [default to null]
**IssueType** | **string** | Issue type of the IPO. &#x60;regular&#x60; is a mainboard issue | [optional] [default to null]
**Reason** | **string** | Free-text reason from the exchange or registrar explaining the current status | [optional] [default to null]
**Upi** | **string** | UPI id the mandate was raised against | [optional] [default to null]
**UpiAmountBlocked** | **string** | Amount blocked in the applicant&#x27;s bank account for this application, as a decimal string. Absent until the mandate is accepted | [optional] [default to null]
**NseSubmittedDate** | **string** | When the application was submitted to NSE, in yyyy-MM-dd&#x27;T&#x27;HH:mm:ss; null if not submitted there | [optional] [default to null]
**BseSubmittedDate** | **string** | When the application was submitted to BSE, in yyyy-MM-dd&#x27;T&#x27;HH:mm:ss; null if not submitted there | [optional] [default to null]
**MandateApprovedDate** | **string** | When the UPI mandate was approved; null until approved | [optional] [default to null]
**RejectionDate** | **string** | When the application was rejected; null unless rejected | [optional] [default to null]
**MandateRejectionDate** | **string** | When the UPI mandate was rejected; null unless rejected | [optional] [default to null]
**CancelRequestedDate** | **string** | When cancellation was requested; null unless a cancellation was raised | [optional] [default to null]
**CancelAcceptedDate** | **string** | When cancellation was accepted; null unless the cancellation completed | [optional] [default to null]
**UnitsAllotted** | **int32** | Shares allotted. 0 until allotment completes, and on applications that were not allotted | [optional] [default to null]
**Bids** | [**[]IpoOrderBid**](IpoOrderBid.md) | Bids placed in this application | [optional] [default to null]
**CreatedAt** | **string** | When the application was created, in yyyy-MM-dd&#x27;T&#x27;HH:mm:ss | [optional] [default to null]
**LastUpdatedAt** | **string** | When the application was last updated, in yyyy-MM-dd&#x27;T&#x27;HH:mm:ss | [optional] [default to null]

[Back to Model list] [Back to API list] [Back to README]

