# TokenRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | [default to null]
**ClientId** | **string** | OAuth API key that is a public identifier for app | [default to null]
**ClientSecret** | **string** | OAuth client secret that is a private secret known only to app and authorization server | [default to null]
**RedirectUri** | **string** | Authorization server will redirect the user back to the application via redirect url | [default to null]
**GrantType** | **string** | Type of grant used to get an access token | [default to null]
**RefreshExtendedToken** | **bool** | Whether to refresh (regenerate) the extended access token. Defaults to false when absent. Set to true to refresh; false or omit to reuse an existing extended token when available. | [optional] [default to null]

[Back to Model list] [Back to API list] [Back to README]

