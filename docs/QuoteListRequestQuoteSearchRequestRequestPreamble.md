# QuoteListRequestQuoteSearchRequestRequestPreamble

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerNumber** | **string** | Reseller Number (referred to as the account BCN) is the unique identifier for an Ingram Micro customer account. | 
**CustomerContact** | Pointer to **string** | Logged in User&#39;s email address. | [optional] 
**IsoCountryCode** | **string** | The ISO country codes are internationally recognized codes designated for each country represented by a two-letter combination (alpha-2). | 

## Methods

### NewQuoteListRequestQuoteSearchRequestRequestPreamble

`func NewQuoteListRequestQuoteSearchRequestRequestPreamble(customerNumber string, isoCountryCode string, ) *QuoteListRequestQuoteSearchRequestRequestPreamble`

NewQuoteListRequestQuoteSearchRequestRequestPreamble instantiates a new QuoteListRequestQuoteSearchRequestRequestPreamble object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteListRequestQuoteSearchRequestRequestPreambleWithDefaults

`func NewQuoteListRequestQuoteSearchRequestRequestPreambleWithDefaults() *QuoteListRequestQuoteSearchRequestRequestPreamble`

NewQuoteListRequestQuoteSearchRequestRequestPreambleWithDefaults instantiates a new QuoteListRequestQuoteSearchRequestRequestPreamble object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerNumber

`func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) GetCustomerNumber() string`

GetCustomerNumber returns the CustomerNumber field if non-nil, zero value otherwise.

### GetCustomerNumberOk

`func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) GetCustomerNumberOk() (*string, bool)`

GetCustomerNumberOk returns a tuple with the CustomerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNumber

`func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) SetCustomerNumber(v string)`

SetCustomerNumber sets CustomerNumber field to given value.


### GetCustomerContact

`func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) GetCustomerContact() string`

GetCustomerContact returns the CustomerContact field if non-nil, zero value otherwise.

### GetCustomerContactOk

`func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) GetCustomerContactOk() (*string, bool)`

GetCustomerContactOk returns a tuple with the CustomerContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerContact

`func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) SetCustomerContact(v string)`

SetCustomerContact sets CustomerContact field to given value.

### HasCustomerContact

`func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) HasCustomerContact() bool`

HasCustomerContact returns a boolean if a field has been set.

### GetIsoCountryCode

`func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) GetIsoCountryCode() string`

GetIsoCountryCode returns the IsoCountryCode field if non-nil, zero value otherwise.

### GetIsoCountryCodeOk

`func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) GetIsoCountryCodeOk() (*string, bool)`

GetIsoCountryCodeOk returns a tuple with the IsoCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCountryCode

`func (o *QuoteListRequestQuoteSearchRequestRequestPreamble) SetIsoCountryCode(v string)`

SetIsoCountryCode sets IsoCountryCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


