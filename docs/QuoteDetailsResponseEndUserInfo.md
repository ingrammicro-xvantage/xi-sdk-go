# QuoteDetailsResponseEndUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to **string** | End User Name | [optional] 
**CompanyName** | Pointer to **string** | Contact name  of end user associated with the quote. | [optional] 
**AddressLine1** | Pointer to **string** | Address line 1 for end user associated with the quote | [optional] 
**AddressLine2** | Pointer to **string** | Address line 2 for end user associated with the quote. | [optional] 
**AddressLine3** | Pointer to **string** | Address line 3 for end user associated with the quote. | [optional] 
**City** | Pointer to **string** | City for end user associated with the quote | [optional] 
**State** | Pointer to **string** | Two letter state abreviation for end user associated with the quote | [optional] 
**Email** | Pointer to **string** | Email of end user the quote associated with the quote. | [optional] 
**PhoneNumber** | Pointer to **string** | Phone number of end user associated with the quote. | [optional] 
**PostalCode** | Pointer to **string** | Zip code of end user associated with the quote. | [optional] 
**MarketSegment** | Pointer to **string** | Market Segment of end user associated with the quote. End user market segment is included when end user is included in specific market segments like Educational, Government, Military, Medical - that may receive special pricing due to their segmentation. | [optional] 

## Methods

### NewQuoteDetailsResponseEndUserInfo

`func NewQuoteDetailsResponseEndUserInfo() *QuoteDetailsResponseEndUserInfo`

NewQuoteDetailsResponseEndUserInfo instantiates a new QuoteDetailsResponseEndUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteDetailsResponseEndUserInfoWithDefaults

`func NewQuoteDetailsResponseEndUserInfoWithDefaults() *QuoteDetailsResponseEndUserInfo`

NewQuoteDetailsResponseEndUserInfoWithDefaults instantiates a new QuoteDetailsResponseEndUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *QuoteDetailsResponseEndUserInfo) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *QuoteDetailsResponseEndUserInfo) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *QuoteDetailsResponseEndUserInfo) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *QuoteDetailsResponseEndUserInfo) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetCompanyName

`func (o *QuoteDetailsResponseEndUserInfo) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *QuoteDetailsResponseEndUserInfo) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *QuoteDetailsResponseEndUserInfo) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *QuoteDetailsResponseEndUserInfo) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetAddressLine1

`func (o *QuoteDetailsResponseEndUserInfo) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *QuoteDetailsResponseEndUserInfo) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *QuoteDetailsResponseEndUserInfo) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *QuoteDetailsResponseEndUserInfo) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *QuoteDetailsResponseEndUserInfo) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *QuoteDetailsResponseEndUserInfo) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *QuoteDetailsResponseEndUserInfo) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *QuoteDetailsResponseEndUserInfo) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *QuoteDetailsResponseEndUserInfo) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *QuoteDetailsResponseEndUserInfo) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *QuoteDetailsResponseEndUserInfo) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *QuoteDetailsResponseEndUserInfo) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetCity

`func (o *QuoteDetailsResponseEndUserInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *QuoteDetailsResponseEndUserInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *QuoteDetailsResponseEndUserInfo) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *QuoteDetailsResponseEndUserInfo) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *QuoteDetailsResponseEndUserInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *QuoteDetailsResponseEndUserInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *QuoteDetailsResponseEndUserInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *QuoteDetailsResponseEndUserInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetEmail

`func (o *QuoteDetailsResponseEndUserInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *QuoteDetailsResponseEndUserInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *QuoteDetailsResponseEndUserInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *QuoteDetailsResponseEndUserInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *QuoteDetailsResponseEndUserInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *QuoteDetailsResponseEndUserInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *QuoteDetailsResponseEndUserInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *QuoteDetailsResponseEndUserInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPostalCode

`func (o *QuoteDetailsResponseEndUserInfo) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *QuoteDetailsResponseEndUserInfo) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *QuoteDetailsResponseEndUserInfo) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *QuoteDetailsResponseEndUserInfo) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetMarketSegment

`func (o *QuoteDetailsResponseEndUserInfo) GetMarketSegment() string`

GetMarketSegment returns the MarketSegment field if non-nil, zero value otherwise.

### GetMarketSegmentOk

`func (o *QuoteDetailsResponseEndUserInfo) GetMarketSegmentOk() (*string, bool)`

GetMarketSegmentOk returns a tuple with the MarketSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketSegment

`func (o *QuoteDetailsResponseEndUserInfo) SetMarketSegment(v string)`

SetMarketSegment sets MarketSegment field to given value.

### HasMarketSegment

`func (o *QuoteDetailsResponseEndUserInfo) HasMarketSegment() bool`

HasMarketSegment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


