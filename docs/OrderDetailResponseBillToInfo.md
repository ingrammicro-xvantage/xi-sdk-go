# OrderDetailResponseBillToInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to **string** | The company contact provided by the reseller. | [optional] 
**CompanyName** | Pointer to **string** | The name of the company the order will be billed to. | [optional] 
**Name1** | Pointer to **string** | First name. | [optional] 
**Name2** | Pointer to **string** | Last name. | [optional] 
**AddressLine1** | Pointer to **string** | The street address and building or house number the order will be billed to. | [optional] 
**AddressLine2** | Pointer to **string** | The apartment number the order will be billed to. | [optional] 
**AddressLine3** | Pointer to **string** | Address line 3. | [optional] 
**City** | Pointer to **string** | The city the order will be billed to. | [optional] 
**State** | Pointer to **string** | The state the order will be billed to. | [optional] 
**PostalCode** | Pointer to **string** | The zip or postal code the order will be billed to. | [optional] 
**CountryCode** | Pointer to **string** | The two-character ISO country code the order will be billed to. | [optional] 
**PhoneNumber** | Pointer to **string** | The company contact phone number. | [optional] 
**Email** | Pointer to **string** | The company contact email address. | [optional] 

## Methods

### NewOrderDetailResponseBillToInfo

`func NewOrderDetailResponseBillToInfo() *OrderDetailResponseBillToInfo`

NewOrderDetailResponseBillToInfo instantiates a new OrderDetailResponseBillToInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDetailResponseBillToInfoWithDefaults

`func NewOrderDetailResponseBillToInfoWithDefaults() *OrderDetailResponseBillToInfo`

NewOrderDetailResponseBillToInfoWithDefaults instantiates a new OrderDetailResponseBillToInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *OrderDetailResponseBillToInfo) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *OrderDetailResponseBillToInfo) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *OrderDetailResponseBillToInfo) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *OrderDetailResponseBillToInfo) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetCompanyName

`func (o *OrderDetailResponseBillToInfo) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *OrderDetailResponseBillToInfo) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *OrderDetailResponseBillToInfo) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *OrderDetailResponseBillToInfo) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetName1

`func (o *OrderDetailResponseBillToInfo) GetName1() string`

GetName1 returns the Name1 field if non-nil, zero value otherwise.

### GetName1Ok

`func (o *OrderDetailResponseBillToInfo) GetName1Ok() (*string, bool)`

GetName1Ok returns a tuple with the Name1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName1

`func (o *OrderDetailResponseBillToInfo) SetName1(v string)`

SetName1 sets Name1 field to given value.

### HasName1

`func (o *OrderDetailResponseBillToInfo) HasName1() bool`

HasName1 returns a boolean if a field has been set.

### GetName2

`func (o *OrderDetailResponseBillToInfo) GetName2() string`

GetName2 returns the Name2 field if non-nil, zero value otherwise.

### GetName2Ok

`func (o *OrderDetailResponseBillToInfo) GetName2Ok() (*string, bool)`

GetName2Ok returns a tuple with the Name2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName2

`func (o *OrderDetailResponseBillToInfo) SetName2(v string)`

SetName2 sets Name2 field to given value.

### HasName2

`func (o *OrderDetailResponseBillToInfo) HasName2() bool`

HasName2 returns a boolean if a field has been set.

### GetAddressLine1

`func (o *OrderDetailResponseBillToInfo) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *OrderDetailResponseBillToInfo) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *OrderDetailResponseBillToInfo) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *OrderDetailResponseBillToInfo) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *OrderDetailResponseBillToInfo) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *OrderDetailResponseBillToInfo) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *OrderDetailResponseBillToInfo) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *OrderDetailResponseBillToInfo) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *OrderDetailResponseBillToInfo) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *OrderDetailResponseBillToInfo) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *OrderDetailResponseBillToInfo) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *OrderDetailResponseBillToInfo) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetCity

`func (o *OrderDetailResponseBillToInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrderDetailResponseBillToInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrderDetailResponseBillToInfo) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrderDetailResponseBillToInfo) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *OrderDetailResponseBillToInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrderDetailResponseBillToInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrderDetailResponseBillToInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OrderDetailResponseBillToInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *OrderDetailResponseBillToInfo) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *OrderDetailResponseBillToInfo) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *OrderDetailResponseBillToInfo) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *OrderDetailResponseBillToInfo) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *OrderDetailResponseBillToInfo) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *OrderDetailResponseBillToInfo) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *OrderDetailResponseBillToInfo) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *OrderDetailResponseBillToInfo) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *OrderDetailResponseBillToInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *OrderDetailResponseBillToInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *OrderDetailResponseBillToInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *OrderDetailResponseBillToInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetEmail

`func (o *OrderDetailResponseBillToInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrderDetailResponseBillToInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrderDetailResponseBillToInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrderDetailResponseBillToInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


