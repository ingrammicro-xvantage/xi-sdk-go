# OrderCreateV7RequestEndUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndUserId** | Pointer to **string** | ID for the end user/customer in Ingram Micro&#39;s system. | [optional] 
**Contact** | Pointer to **string** | The contact name for the end user/customer. | [optional] 
**CompanyName** | Pointer to **string** | The company name for the end user/customer. Required for Impulse countries. | [optional] 
**AddressLine1** | Pointer to **string** | The end user/customer&#39;s street address and building or house number. Required for Impulse countries. | [optional] 
**AddressLine2** | Pointer to **string** | The end user/customer&#39;s apartment number. | [optional] 
**City** | Pointer to **string** | The end user/customer&#39;s city. Required for Impulse countries. | [optional] 
**State** | Pointer to **string** | The end user/customer&#39;s state. Required for Impulse countries but optional for EMEA countries. | [optional] 
**PostalCode** | Pointer to **string** | The end user/customer&#39;s zip or postal code. Required for Impulse countries. | [optional] 
**CountryCode** | Pointer to **string** | The end user/customer&#39;s two-character ISO country code. | [optional] 
**PhoneNumber** | Pointer to **int32** | The end user/customer&#39;s phone number. | [optional] 
**Email** | Pointer to **string** | The end user/customer&#39;s email. | [optional] 

## Methods

### NewOrderCreateV7RequestEndUserInfo

`func NewOrderCreateV7RequestEndUserInfo() *OrderCreateV7RequestEndUserInfo`

NewOrderCreateV7RequestEndUserInfo instantiates a new OrderCreateV7RequestEndUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateV7RequestEndUserInfoWithDefaults

`func NewOrderCreateV7RequestEndUserInfoWithDefaults() *OrderCreateV7RequestEndUserInfo`

NewOrderCreateV7RequestEndUserInfoWithDefaults instantiates a new OrderCreateV7RequestEndUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndUserId

`func (o *OrderCreateV7RequestEndUserInfo) GetEndUserId() string`

GetEndUserId returns the EndUserId field if non-nil, zero value otherwise.

### GetEndUserIdOk

`func (o *OrderCreateV7RequestEndUserInfo) GetEndUserIdOk() (*string, bool)`

GetEndUserIdOk returns a tuple with the EndUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserId

`func (o *OrderCreateV7RequestEndUserInfo) SetEndUserId(v string)`

SetEndUserId sets EndUserId field to given value.

### HasEndUserId

`func (o *OrderCreateV7RequestEndUserInfo) HasEndUserId() bool`

HasEndUserId returns a boolean if a field has been set.

### GetContact

`func (o *OrderCreateV7RequestEndUserInfo) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *OrderCreateV7RequestEndUserInfo) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *OrderCreateV7RequestEndUserInfo) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *OrderCreateV7RequestEndUserInfo) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetCompanyName

`func (o *OrderCreateV7RequestEndUserInfo) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *OrderCreateV7RequestEndUserInfo) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *OrderCreateV7RequestEndUserInfo) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *OrderCreateV7RequestEndUserInfo) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetAddressLine1

`func (o *OrderCreateV7RequestEndUserInfo) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *OrderCreateV7RequestEndUserInfo) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *OrderCreateV7RequestEndUserInfo) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *OrderCreateV7RequestEndUserInfo) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *OrderCreateV7RequestEndUserInfo) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *OrderCreateV7RequestEndUserInfo) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *OrderCreateV7RequestEndUserInfo) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *OrderCreateV7RequestEndUserInfo) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *OrderCreateV7RequestEndUserInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrderCreateV7RequestEndUserInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrderCreateV7RequestEndUserInfo) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrderCreateV7RequestEndUserInfo) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *OrderCreateV7RequestEndUserInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrderCreateV7RequestEndUserInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrderCreateV7RequestEndUserInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OrderCreateV7RequestEndUserInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *OrderCreateV7RequestEndUserInfo) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *OrderCreateV7RequestEndUserInfo) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *OrderCreateV7RequestEndUserInfo) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *OrderCreateV7RequestEndUserInfo) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *OrderCreateV7RequestEndUserInfo) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *OrderCreateV7RequestEndUserInfo) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *OrderCreateV7RequestEndUserInfo) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *OrderCreateV7RequestEndUserInfo) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *OrderCreateV7RequestEndUserInfo) GetPhoneNumber() int32`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *OrderCreateV7RequestEndUserInfo) GetPhoneNumberOk() (*int32, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *OrderCreateV7RequestEndUserInfo) SetPhoneNumber(v int32)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *OrderCreateV7RequestEndUserInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetEmail

`func (o *OrderCreateV7RequestEndUserInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrderCreateV7RequestEndUserInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrderCreateV7RequestEndUserInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrderCreateV7RequestEndUserInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


