# OrderCreateResponseEndUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndUserId** | Pointer to **string** | The unique ID provided by the reseller for the end user/customer. | [optional] 
**Contact** | Pointer to **string** | The contact name for the end user/customer. | [optional] 
**CompanyName** | Pointer to **string** | The company name for the end user/customer. | [optional] 
**Name1** | Pointer to **string** | name1 | [optional] 
**Name2** | Pointer to **string** | name2 | [optional] 
**AddressLine1** | Pointer to **string** | The street adress and building or house number for the end user/customer. | [optional] 
**AddressLine2** | Pointer to **string** | The apartment number for the end user/customer. | [optional] 
**AddressLine3** | Pointer to **string** | Line 3 of the address for the end user/customer. | [optional] 
**AddressLine4** | Pointer to **string** | Street address4 | [optional] 
**City** | Pointer to **string** | The end user/customer&#39;s city. | [optional] 
**State** | Pointer to **string** | The end user/customer&#39;s state. | [optional] 
**PostalCode** | Pointer to **string** | The end user/customer&#39;s zip or postal code. | [optional] 
**CountryCode** | Pointer to **string** | The end user/customer&#39;s two-character ISO country code. | [optional] 
**PhoneNumber** | Pointer to **string** | The end user/customer&#39;s phone number. | [optional] 
**Email** | Pointer to **string** | The end user/customer&#39;s email. | [optional] 

## Methods

### NewOrderCreateResponseEndUserInfo

`func NewOrderCreateResponseEndUserInfo() *OrderCreateResponseEndUserInfo`

NewOrderCreateResponseEndUserInfo instantiates a new OrderCreateResponseEndUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateResponseEndUserInfoWithDefaults

`func NewOrderCreateResponseEndUserInfoWithDefaults() *OrderCreateResponseEndUserInfo`

NewOrderCreateResponseEndUserInfoWithDefaults instantiates a new OrderCreateResponseEndUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndUserId

`func (o *OrderCreateResponseEndUserInfo) GetEndUserId() string`

GetEndUserId returns the EndUserId field if non-nil, zero value otherwise.

### GetEndUserIdOk

`func (o *OrderCreateResponseEndUserInfo) GetEndUserIdOk() (*string, bool)`

GetEndUserIdOk returns a tuple with the EndUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserId

`func (o *OrderCreateResponseEndUserInfo) SetEndUserId(v string)`

SetEndUserId sets EndUserId field to given value.

### HasEndUserId

`func (o *OrderCreateResponseEndUserInfo) HasEndUserId() bool`

HasEndUserId returns a boolean if a field has been set.

### GetContact

`func (o *OrderCreateResponseEndUserInfo) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *OrderCreateResponseEndUserInfo) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *OrderCreateResponseEndUserInfo) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *OrderCreateResponseEndUserInfo) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetCompanyName

`func (o *OrderCreateResponseEndUserInfo) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *OrderCreateResponseEndUserInfo) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *OrderCreateResponseEndUserInfo) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *OrderCreateResponseEndUserInfo) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetName1

`func (o *OrderCreateResponseEndUserInfo) GetName1() string`

GetName1 returns the Name1 field if non-nil, zero value otherwise.

### GetName1Ok

`func (o *OrderCreateResponseEndUserInfo) GetName1Ok() (*string, bool)`

GetName1Ok returns a tuple with the Name1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName1

`func (o *OrderCreateResponseEndUserInfo) SetName1(v string)`

SetName1 sets Name1 field to given value.

### HasName1

`func (o *OrderCreateResponseEndUserInfo) HasName1() bool`

HasName1 returns a boolean if a field has been set.

### GetName2

`func (o *OrderCreateResponseEndUserInfo) GetName2() string`

GetName2 returns the Name2 field if non-nil, zero value otherwise.

### GetName2Ok

`func (o *OrderCreateResponseEndUserInfo) GetName2Ok() (*string, bool)`

GetName2Ok returns a tuple with the Name2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName2

`func (o *OrderCreateResponseEndUserInfo) SetName2(v string)`

SetName2 sets Name2 field to given value.

### HasName2

`func (o *OrderCreateResponseEndUserInfo) HasName2() bool`

HasName2 returns a boolean if a field has been set.

### GetAddressLine1

`func (o *OrderCreateResponseEndUserInfo) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *OrderCreateResponseEndUserInfo) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *OrderCreateResponseEndUserInfo) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *OrderCreateResponseEndUserInfo) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *OrderCreateResponseEndUserInfo) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *OrderCreateResponseEndUserInfo) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *OrderCreateResponseEndUserInfo) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *OrderCreateResponseEndUserInfo) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *OrderCreateResponseEndUserInfo) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *OrderCreateResponseEndUserInfo) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *OrderCreateResponseEndUserInfo) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *OrderCreateResponseEndUserInfo) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetAddressLine4

`func (o *OrderCreateResponseEndUserInfo) GetAddressLine4() string`

GetAddressLine4 returns the AddressLine4 field if non-nil, zero value otherwise.

### GetAddressLine4Ok

`func (o *OrderCreateResponseEndUserInfo) GetAddressLine4Ok() (*string, bool)`

GetAddressLine4Ok returns a tuple with the AddressLine4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine4

`func (o *OrderCreateResponseEndUserInfo) SetAddressLine4(v string)`

SetAddressLine4 sets AddressLine4 field to given value.

### HasAddressLine4

`func (o *OrderCreateResponseEndUserInfo) HasAddressLine4() bool`

HasAddressLine4 returns a boolean if a field has been set.

### GetCity

`func (o *OrderCreateResponseEndUserInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrderCreateResponseEndUserInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrderCreateResponseEndUserInfo) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrderCreateResponseEndUserInfo) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *OrderCreateResponseEndUserInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrderCreateResponseEndUserInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrderCreateResponseEndUserInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OrderCreateResponseEndUserInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *OrderCreateResponseEndUserInfo) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *OrderCreateResponseEndUserInfo) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *OrderCreateResponseEndUserInfo) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *OrderCreateResponseEndUserInfo) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *OrderCreateResponseEndUserInfo) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *OrderCreateResponseEndUserInfo) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *OrderCreateResponseEndUserInfo) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *OrderCreateResponseEndUserInfo) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *OrderCreateResponseEndUserInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *OrderCreateResponseEndUserInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *OrderCreateResponseEndUserInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *OrderCreateResponseEndUserInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetEmail

`func (o *OrderCreateResponseEndUserInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrderCreateResponseEndUserInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrderCreateResponseEndUserInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrderCreateResponseEndUserInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


