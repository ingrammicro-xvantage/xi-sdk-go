# OrderCreateRequestShipToInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressId** | Pointer to **string** | The ID references the resellers address in Ingram Micro&#39;s system for shipping. Provided to resellers during the onboarding process. | [optional] 
**Contact** | Pointer to **string** | The company contact provided by the reseller. | [optional] 
**CompanyName** | Pointer to **string** | The name of the company the order will be shipped to. | [optional] 
**Name1** | Pointer to **string** | name1 | [optional] 
**Name2** | Pointer to **string** | name2 | [optional] 
**AddressLine1** | Pointer to **string** | The street address and building or house number the order will be shipped to. | [optional] 
**AddressLine2** | Pointer to **string** | The apartment number the order will be shipped to. | [optional] 
**AddressLine3** | Pointer to **string** | Line 3 of the address the order will be shipped to. | [optional] 
**AddressLine4** | Pointer to **string** | Street address4 | [optional] 
**City** | Pointer to **string** | The city the order will be shipped to. | [optional] 
**State** | Pointer to **string** | The state the order will be shipped to. | [optional] 
**PostalCode** | Pointer to **string** | The zip or postal code the order will be shipped to. | [optional] 
**CountryCode** | Pointer to **string** | The two-character ISO country code the order will be shipped to. | [optional] 
**PhoneNumber** | Pointer to **string** | The company contact phone number. | [optional] 
**Email** | Pointer to **string** | The company contact email address. | [optional] 

## Methods

### NewOrderCreateRequestShipToInfo

`func NewOrderCreateRequestShipToInfo() *OrderCreateRequestShipToInfo`

NewOrderCreateRequestShipToInfo instantiates a new OrderCreateRequestShipToInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateRequestShipToInfoWithDefaults

`func NewOrderCreateRequestShipToInfoWithDefaults() *OrderCreateRequestShipToInfo`

NewOrderCreateRequestShipToInfoWithDefaults instantiates a new OrderCreateRequestShipToInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressId

`func (o *OrderCreateRequestShipToInfo) GetAddressId() string`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *OrderCreateRequestShipToInfo) GetAddressIdOk() (*string, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *OrderCreateRequestShipToInfo) SetAddressId(v string)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *OrderCreateRequestShipToInfo) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetContact

`func (o *OrderCreateRequestShipToInfo) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *OrderCreateRequestShipToInfo) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *OrderCreateRequestShipToInfo) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *OrderCreateRequestShipToInfo) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetCompanyName

`func (o *OrderCreateRequestShipToInfo) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *OrderCreateRequestShipToInfo) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *OrderCreateRequestShipToInfo) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *OrderCreateRequestShipToInfo) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetName1

`func (o *OrderCreateRequestShipToInfo) GetName1() string`

GetName1 returns the Name1 field if non-nil, zero value otherwise.

### GetName1Ok

`func (o *OrderCreateRequestShipToInfo) GetName1Ok() (*string, bool)`

GetName1Ok returns a tuple with the Name1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName1

`func (o *OrderCreateRequestShipToInfo) SetName1(v string)`

SetName1 sets Name1 field to given value.

### HasName1

`func (o *OrderCreateRequestShipToInfo) HasName1() bool`

HasName1 returns a boolean if a field has been set.

### GetName2

`func (o *OrderCreateRequestShipToInfo) GetName2() string`

GetName2 returns the Name2 field if non-nil, zero value otherwise.

### GetName2Ok

`func (o *OrderCreateRequestShipToInfo) GetName2Ok() (*string, bool)`

GetName2Ok returns a tuple with the Name2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName2

`func (o *OrderCreateRequestShipToInfo) SetName2(v string)`

SetName2 sets Name2 field to given value.

### HasName2

`func (o *OrderCreateRequestShipToInfo) HasName2() bool`

HasName2 returns a boolean if a field has been set.

### GetAddressLine1

`func (o *OrderCreateRequestShipToInfo) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *OrderCreateRequestShipToInfo) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *OrderCreateRequestShipToInfo) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *OrderCreateRequestShipToInfo) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *OrderCreateRequestShipToInfo) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *OrderCreateRequestShipToInfo) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *OrderCreateRequestShipToInfo) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *OrderCreateRequestShipToInfo) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *OrderCreateRequestShipToInfo) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *OrderCreateRequestShipToInfo) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *OrderCreateRequestShipToInfo) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *OrderCreateRequestShipToInfo) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetAddressLine4

`func (o *OrderCreateRequestShipToInfo) GetAddressLine4() string`

GetAddressLine4 returns the AddressLine4 field if non-nil, zero value otherwise.

### GetAddressLine4Ok

`func (o *OrderCreateRequestShipToInfo) GetAddressLine4Ok() (*string, bool)`

GetAddressLine4Ok returns a tuple with the AddressLine4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine4

`func (o *OrderCreateRequestShipToInfo) SetAddressLine4(v string)`

SetAddressLine4 sets AddressLine4 field to given value.

### HasAddressLine4

`func (o *OrderCreateRequestShipToInfo) HasAddressLine4() bool`

HasAddressLine4 returns a boolean if a field has been set.

### GetCity

`func (o *OrderCreateRequestShipToInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrderCreateRequestShipToInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrderCreateRequestShipToInfo) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrderCreateRequestShipToInfo) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *OrderCreateRequestShipToInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrderCreateRequestShipToInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrderCreateRequestShipToInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OrderCreateRequestShipToInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *OrderCreateRequestShipToInfo) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *OrderCreateRequestShipToInfo) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *OrderCreateRequestShipToInfo) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *OrderCreateRequestShipToInfo) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *OrderCreateRequestShipToInfo) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *OrderCreateRequestShipToInfo) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *OrderCreateRequestShipToInfo) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *OrderCreateRequestShipToInfo) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *OrderCreateRequestShipToInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *OrderCreateRequestShipToInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *OrderCreateRequestShipToInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *OrderCreateRequestShipToInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetEmail

`func (o *OrderCreateRequestShipToInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrderCreateRequestShipToInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrderCreateRequestShipToInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrderCreateRequestShipToInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


