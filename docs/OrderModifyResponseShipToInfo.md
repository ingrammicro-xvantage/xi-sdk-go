# OrderModifyResponseShipToInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressId** | Pointer to **string** | Suffix used to identify shipping address. Created during onboarding. Resellers are provided with one or more address IDs depending on how many bill to addresses they need for various flooring companies they are using for credit. | [optional] 
**Contact** | Pointer to **string** | The company contact provided by the reseller. | [optional] 
**CompanyName** | Pointer to **string** | The name of the company the order will be shipped to. | [optional] 
**AddressLine1** | Pointer to **string** | The street address and building or house number the order will be shipped to. | [optional] 
**AddressLine2** | Pointer to **string** | The apartment number the order will be shipped to. | [optional] 
**AddressLine3** | Pointer to **string** | Line 3 of the address the order will be shipped to. | [optional] 
**City** | Pointer to **string** | The city the order will be shipped to. | [optional] 
**State** | Pointer to **string** | The state the order will be shipped to. | [optional] 
**PostalCode** | Pointer to **string** | The zip or postal code the order will be shipped to. | [optional] 
**CountryCode** | Pointer to **string** | The two-character ISO country code the order will be shipped to. | [optional] 
**PhoneNumber** | Pointer to **string** | The company contact phone number. | [optional] 
**Email** | Pointer to **string** | The company contact email address. | [optional] 

## Methods

### NewOrderModifyResponseShipToInfo

`func NewOrderModifyResponseShipToInfo() *OrderModifyResponseShipToInfo`

NewOrderModifyResponseShipToInfo instantiates a new OrderModifyResponseShipToInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderModifyResponseShipToInfoWithDefaults

`func NewOrderModifyResponseShipToInfoWithDefaults() *OrderModifyResponseShipToInfo`

NewOrderModifyResponseShipToInfoWithDefaults instantiates a new OrderModifyResponseShipToInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressId

`func (o *OrderModifyResponseShipToInfo) GetAddressId() string`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *OrderModifyResponseShipToInfo) GetAddressIdOk() (*string, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *OrderModifyResponseShipToInfo) SetAddressId(v string)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *OrderModifyResponseShipToInfo) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetContact

`func (o *OrderModifyResponseShipToInfo) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *OrderModifyResponseShipToInfo) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *OrderModifyResponseShipToInfo) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *OrderModifyResponseShipToInfo) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetCompanyName

`func (o *OrderModifyResponseShipToInfo) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *OrderModifyResponseShipToInfo) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *OrderModifyResponseShipToInfo) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *OrderModifyResponseShipToInfo) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetAddressLine1

`func (o *OrderModifyResponseShipToInfo) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *OrderModifyResponseShipToInfo) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *OrderModifyResponseShipToInfo) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *OrderModifyResponseShipToInfo) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *OrderModifyResponseShipToInfo) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *OrderModifyResponseShipToInfo) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *OrderModifyResponseShipToInfo) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *OrderModifyResponseShipToInfo) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *OrderModifyResponseShipToInfo) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *OrderModifyResponseShipToInfo) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *OrderModifyResponseShipToInfo) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *OrderModifyResponseShipToInfo) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetCity

`func (o *OrderModifyResponseShipToInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrderModifyResponseShipToInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrderModifyResponseShipToInfo) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrderModifyResponseShipToInfo) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *OrderModifyResponseShipToInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrderModifyResponseShipToInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrderModifyResponseShipToInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OrderModifyResponseShipToInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *OrderModifyResponseShipToInfo) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *OrderModifyResponseShipToInfo) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *OrderModifyResponseShipToInfo) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *OrderModifyResponseShipToInfo) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *OrderModifyResponseShipToInfo) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *OrderModifyResponseShipToInfo) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *OrderModifyResponseShipToInfo) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *OrderModifyResponseShipToInfo) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *OrderModifyResponseShipToInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *OrderModifyResponseShipToInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *OrderModifyResponseShipToInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *OrderModifyResponseShipToInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetEmail

`func (o *OrderModifyResponseShipToInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrderModifyResponseShipToInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrderModifyResponseShipToInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrderModifyResponseShipToInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


