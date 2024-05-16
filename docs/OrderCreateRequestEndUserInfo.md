# OrderCreateRequestEndUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndUserId** | Pointer to **string** | ID for the end user/customer in Ingram Micro&#39;s system. | [optional] 
**Contact** | Pointer to **string** | The contact name for the end user/customer. | [optional] 
**CompanyName** | Pointer to **string** | The company name for the end user/customer. Required for Impulse countries. | [optional] 
**Name1** | Pointer to **string** | name1 | [optional] 
**Name2** | Pointer to **string** | name2 | [optional] 
**AddressLine1** | Pointer to **string** | The end user/customer&#39;s street address and building or house number. Required for Impulse countries. | [optional] 
**AddressLine2** | Pointer to **string** | The end user/customer&#39;s apartment number. | [optional] 
**AddressLine3** | Pointer to **string** | Line 3 of the address for the end user/customer. | [optional] 
**AddressLine4** | Pointer to **string** | Street address4 | [optional] 
**City** | Pointer to **string** | The end user/customer&#39;s city. Required for Impulse countries. | [optional] 
**State** | Pointer to **string** | The end user/customer&#39;s state. Required for Impulse countries but optional for EMEA countries. | [optional] 
**PostalCode** | Pointer to **string** | The end user/customer&#39;s zip or postal code. Required for Impulse countries. | [optional] 
**CountryCode** | Pointer to **string** | The end user/customer&#39;s two-character ISO country code. | [optional] 
**PhoneNumber** | Pointer to **string** | The end user/customer&#39;s phone number. | [optional] 
**Email** | Pointer to **string** | The end user/customer&#39;s email. | [optional] 

## Methods

### NewOrderCreateRequestEndUserInfo

`func NewOrderCreateRequestEndUserInfo() *OrderCreateRequestEndUserInfo`

NewOrderCreateRequestEndUserInfo instantiates a new OrderCreateRequestEndUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateRequestEndUserInfoWithDefaults

`func NewOrderCreateRequestEndUserInfoWithDefaults() *OrderCreateRequestEndUserInfo`

NewOrderCreateRequestEndUserInfoWithDefaults instantiates a new OrderCreateRequestEndUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndUserId

`func (o *OrderCreateRequestEndUserInfo) GetEndUserId() string`

GetEndUserId returns the EndUserId field if non-nil, zero value otherwise.

### GetEndUserIdOk

`func (o *OrderCreateRequestEndUserInfo) GetEndUserIdOk() (*string, bool)`

GetEndUserIdOk returns a tuple with the EndUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserId

`func (o *OrderCreateRequestEndUserInfo) SetEndUserId(v string)`

SetEndUserId sets EndUserId field to given value.

### HasEndUserId

`func (o *OrderCreateRequestEndUserInfo) HasEndUserId() bool`

HasEndUserId returns a boolean if a field has been set.

### GetContact

`func (o *OrderCreateRequestEndUserInfo) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *OrderCreateRequestEndUserInfo) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *OrderCreateRequestEndUserInfo) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *OrderCreateRequestEndUserInfo) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetCompanyName

`func (o *OrderCreateRequestEndUserInfo) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *OrderCreateRequestEndUserInfo) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *OrderCreateRequestEndUserInfo) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *OrderCreateRequestEndUserInfo) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetName1

`func (o *OrderCreateRequestEndUserInfo) GetName1() string`

GetName1 returns the Name1 field if non-nil, zero value otherwise.

### GetName1Ok

`func (o *OrderCreateRequestEndUserInfo) GetName1Ok() (*string, bool)`

GetName1Ok returns a tuple with the Name1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName1

`func (o *OrderCreateRequestEndUserInfo) SetName1(v string)`

SetName1 sets Name1 field to given value.

### HasName1

`func (o *OrderCreateRequestEndUserInfo) HasName1() bool`

HasName1 returns a boolean if a field has been set.

### GetName2

`func (o *OrderCreateRequestEndUserInfo) GetName2() string`

GetName2 returns the Name2 field if non-nil, zero value otherwise.

### GetName2Ok

`func (o *OrderCreateRequestEndUserInfo) GetName2Ok() (*string, bool)`

GetName2Ok returns a tuple with the Name2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName2

`func (o *OrderCreateRequestEndUserInfo) SetName2(v string)`

SetName2 sets Name2 field to given value.

### HasName2

`func (o *OrderCreateRequestEndUserInfo) HasName2() bool`

HasName2 returns a boolean if a field has been set.

### GetAddressLine1

`func (o *OrderCreateRequestEndUserInfo) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *OrderCreateRequestEndUserInfo) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *OrderCreateRequestEndUserInfo) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *OrderCreateRequestEndUserInfo) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *OrderCreateRequestEndUserInfo) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *OrderCreateRequestEndUserInfo) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *OrderCreateRequestEndUserInfo) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *OrderCreateRequestEndUserInfo) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *OrderCreateRequestEndUserInfo) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *OrderCreateRequestEndUserInfo) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *OrderCreateRequestEndUserInfo) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *OrderCreateRequestEndUserInfo) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetAddressLine4

`func (o *OrderCreateRequestEndUserInfo) GetAddressLine4() string`

GetAddressLine4 returns the AddressLine4 field if non-nil, zero value otherwise.

### GetAddressLine4Ok

`func (o *OrderCreateRequestEndUserInfo) GetAddressLine4Ok() (*string, bool)`

GetAddressLine4Ok returns a tuple with the AddressLine4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine4

`func (o *OrderCreateRequestEndUserInfo) SetAddressLine4(v string)`

SetAddressLine4 sets AddressLine4 field to given value.

### HasAddressLine4

`func (o *OrderCreateRequestEndUserInfo) HasAddressLine4() bool`

HasAddressLine4 returns a boolean if a field has been set.

### GetCity

`func (o *OrderCreateRequestEndUserInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrderCreateRequestEndUserInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrderCreateRequestEndUserInfo) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrderCreateRequestEndUserInfo) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *OrderCreateRequestEndUserInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrderCreateRequestEndUserInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrderCreateRequestEndUserInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OrderCreateRequestEndUserInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *OrderCreateRequestEndUserInfo) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *OrderCreateRequestEndUserInfo) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *OrderCreateRequestEndUserInfo) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *OrderCreateRequestEndUserInfo) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *OrderCreateRequestEndUserInfo) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *OrderCreateRequestEndUserInfo) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *OrderCreateRequestEndUserInfo) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *OrderCreateRequestEndUserInfo) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *OrderCreateRequestEndUserInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *OrderCreateRequestEndUserInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *OrderCreateRequestEndUserInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *OrderCreateRequestEndUserInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetEmail

`func (o *OrderCreateRequestEndUserInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrderCreateRequestEndUserInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrderCreateRequestEndUserInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrderCreateRequestEndUserInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


