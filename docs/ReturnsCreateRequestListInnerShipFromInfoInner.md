# ReturnsCreateRequestListInnerShipFromInfoInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | **string** | Name of the company from where the product will be shipped. | 
**Contact** | **string** | Contact name of the person from where the product will be shipped. | 
**AddressLine1** | **string** | Ship from Address Line1. | 
**AddressLine2** | Pointer to **string** | Ship from Address Line2. | [optional] 
**AddressLine3** | Pointer to **string** | Ship from Address Line3. | [optional] 
**City** | **string** | Ship from City. | 
**State** | **string** | Ship from state. | 
**PostalCode** | **string** | Ship from postal code. | 
**CountryCode** | **string** | ship from country code. | 
**Email** | **string** | Ship from email. | 
**PhoneNumber** | Pointer to **string** | Ship from phone number. | [optional] 

## Methods

### NewReturnsCreateRequestListInnerShipFromInfoInner

`func NewReturnsCreateRequestListInnerShipFromInfoInner(companyName string, contact string, addressLine1 string, city string, state string, postalCode string, countryCode string, email string, ) *ReturnsCreateRequestListInnerShipFromInfoInner`

NewReturnsCreateRequestListInnerShipFromInfoInner instantiates a new ReturnsCreateRequestListInnerShipFromInfoInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnsCreateRequestListInnerShipFromInfoInnerWithDefaults

`func NewReturnsCreateRequestListInnerShipFromInfoInnerWithDefaults() *ReturnsCreateRequestListInnerShipFromInfoInner`

NewReturnsCreateRequestListInnerShipFromInfoInnerWithDefaults instantiates a new ReturnsCreateRequestListInnerShipFromInfoInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyName

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetContact

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetContact(v string)`

SetContact sets Contact field to given value.


### GetAddressLine1

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetAddressLine2

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressLine3

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) HasAddressLine3() bool`

HasAddressLine3 returns a boolean if a field has been set.

### GetCity

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetState(v string)`

SetState sets State field to given value.


### GetPostalCode

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetCountryCode

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetEmail

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhoneNumber

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ReturnsCreateRequestListInnerShipFromInfoInner) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


