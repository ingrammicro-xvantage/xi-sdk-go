# OrderCreateV7RequestResellerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResellerId** | Pointer to **string** | The reseller&#39;s Id. | [optional] 
**CompanyName** | Pointer to **string** | The reseller&#39;s company name. | [optional] 
**Contact** | Pointer to **string** | The reseller&#39;s contact name. | [optional] 
**AddressLine1** | Pointer to **string** | The reseller&#39;s address line 1 | [optional] 
**AddressLine2** | Pointer to **string** | The reseller&#39;s address line 2. | [optional] 
**City** | Pointer to **string** | The reseller&#39;s city. | [optional] 
**State** | Pointer to **string** | The reseller&#39;s state. | [optional] 
**PostalCode** | Pointer to **string** | The reseller&#39;s zip or postal code. | [optional] 
**CountryCode** | Pointer to **string** | The reseller&#39;s two-character ISO country code. | [optional] 
**PhoneNumber** | Pointer to **int32** | The reseller&#39;s phone number. | [optional] 
**Email** | Pointer to **string** | The reseller&#39;s email address. | [optional] 

## Methods

### NewOrderCreateV7RequestResellerInfo

`func NewOrderCreateV7RequestResellerInfo() *OrderCreateV7RequestResellerInfo`

NewOrderCreateV7RequestResellerInfo instantiates a new OrderCreateV7RequestResellerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateV7RequestResellerInfoWithDefaults

`func NewOrderCreateV7RequestResellerInfoWithDefaults() *OrderCreateV7RequestResellerInfo`

NewOrderCreateV7RequestResellerInfoWithDefaults instantiates a new OrderCreateV7RequestResellerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResellerId

`func (o *OrderCreateV7RequestResellerInfo) GetResellerId() string`

GetResellerId returns the ResellerId field if non-nil, zero value otherwise.

### GetResellerIdOk

`func (o *OrderCreateV7RequestResellerInfo) GetResellerIdOk() (*string, bool)`

GetResellerIdOk returns a tuple with the ResellerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerId

`func (o *OrderCreateV7RequestResellerInfo) SetResellerId(v string)`

SetResellerId sets ResellerId field to given value.

### HasResellerId

`func (o *OrderCreateV7RequestResellerInfo) HasResellerId() bool`

HasResellerId returns a boolean if a field has been set.

### GetCompanyName

`func (o *OrderCreateV7RequestResellerInfo) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *OrderCreateV7RequestResellerInfo) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *OrderCreateV7RequestResellerInfo) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *OrderCreateV7RequestResellerInfo) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetContact

`func (o *OrderCreateV7RequestResellerInfo) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *OrderCreateV7RequestResellerInfo) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *OrderCreateV7RequestResellerInfo) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *OrderCreateV7RequestResellerInfo) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetAddressLine1

`func (o *OrderCreateV7RequestResellerInfo) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *OrderCreateV7RequestResellerInfo) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *OrderCreateV7RequestResellerInfo) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *OrderCreateV7RequestResellerInfo) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *OrderCreateV7RequestResellerInfo) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *OrderCreateV7RequestResellerInfo) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *OrderCreateV7RequestResellerInfo) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *OrderCreateV7RequestResellerInfo) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *OrderCreateV7RequestResellerInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrderCreateV7RequestResellerInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrderCreateV7RequestResellerInfo) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrderCreateV7RequestResellerInfo) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *OrderCreateV7RequestResellerInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrderCreateV7RequestResellerInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrderCreateV7RequestResellerInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OrderCreateV7RequestResellerInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *OrderCreateV7RequestResellerInfo) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *OrderCreateV7RequestResellerInfo) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *OrderCreateV7RequestResellerInfo) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *OrderCreateV7RequestResellerInfo) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *OrderCreateV7RequestResellerInfo) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *OrderCreateV7RequestResellerInfo) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *OrderCreateV7RequestResellerInfo) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *OrderCreateV7RequestResellerInfo) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *OrderCreateV7RequestResellerInfo) GetPhoneNumber() int32`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *OrderCreateV7RequestResellerInfo) GetPhoneNumberOk() (*int32, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *OrderCreateV7RequestResellerInfo) SetPhoneNumber(v int32)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *OrderCreateV7RequestResellerInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetEmail

`func (o *OrderCreateV7RequestResellerInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrderCreateV7RequestResellerInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrderCreateV7RequestResellerInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrderCreateV7RequestResellerInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


