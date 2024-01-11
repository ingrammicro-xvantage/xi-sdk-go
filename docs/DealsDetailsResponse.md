# DealsDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DealId** | Pointer to **string** | Deal/Special bid number. | [optional] 
**Version** | Pointer to **string** | Most recent version number of the deal. | [optional] 
**EndUser** | Pointer to **string** | The end user/customer&#39;s name. | [optional] 
**ExtendedMsrp** | Pointer to **float32** | Extended MSRP - Manufacturer Suggested Retail Price X Quantity. | [optional] 
**Vendor** | Pointer to **string** | The vendor&#39;s name. | [optional] 
**DealReceivedOn** | Pointer to **string** | The date on which the deal starts. | [optional] 
**DealExpiryDate** | Pointer to **string** | Expiration date of the deal/Special bid. | [optional] 
**PriceProtectionEndDate** | Pointer to **string** | The date on which the price protection will end. | [optional] 
**CurrencyCode** | Pointer to **string** | Country specific currency code. | [optional] 
**EndUserInfo** | Pointer to [**[]RenewalsDetailsResponseEndUserInfoInner**](RenewalsDetailsResponseEndUserInfoInner.md) |  | [optional] 
**Products** | Pointer to [**[]DealsDetailsResponseProductsInner**](DealsDetailsResponseProductsInner.md) |  | [optional] 

## Methods

### NewDealsDetailsResponse

`func NewDealsDetailsResponse() *DealsDetailsResponse`

NewDealsDetailsResponse instantiates a new DealsDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealsDetailsResponseWithDefaults

`func NewDealsDetailsResponseWithDefaults() *DealsDetailsResponse`

NewDealsDetailsResponseWithDefaults instantiates a new DealsDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDealId

`func (o *DealsDetailsResponse) GetDealId() string`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *DealsDetailsResponse) GetDealIdOk() (*string, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *DealsDetailsResponse) SetDealId(v string)`

SetDealId sets DealId field to given value.

### HasDealId

`func (o *DealsDetailsResponse) HasDealId() bool`

HasDealId returns a boolean if a field has been set.

### GetVersion

`func (o *DealsDetailsResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DealsDetailsResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DealsDetailsResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DealsDetailsResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEndUser

`func (o *DealsDetailsResponse) GetEndUser() string`

GetEndUser returns the EndUser field if non-nil, zero value otherwise.

### GetEndUserOk

`func (o *DealsDetailsResponse) GetEndUserOk() (*string, bool)`

GetEndUserOk returns a tuple with the EndUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUser

`func (o *DealsDetailsResponse) SetEndUser(v string)`

SetEndUser sets EndUser field to given value.

### HasEndUser

`func (o *DealsDetailsResponse) HasEndUser() bool`

HasEndUser returns a boolean if a field has been set.

### GetExtendedMsrp

`func (o *DealsDetailsResponse) GetExtendedMsrp() float32`

GetExtendedMsrp returns the ExtendedMsrp field if non-nil, zero value otherwise.

### GetExtendedMsrpOk

`func (o *DealsDetailsResponse) GetExtendedMsrpOk() (*float32, bool)`

GetExtendedMsrpOk returns a tuple with the ExtendedMsrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedMsrp

`func (o *DealsDetailsResponse) SetExtendedMsrp(v float32)`

SetExtendedMsrp sets ExtendedMsrp field to given value.

### HasExtendedMsrp

`func (o *DealsDetailsResponse) HasExtendedMsrp() bool`

HasExtendedMsrp returns a boolean if a field has been set.

### GetVendor

`func (o *DealsDetailsResponse) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *DealsDetailsResponse) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *DealsDetailsResponse) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *DealsDetailsResponse) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetDealReceivedOn

`func (o *DealsDetailsResponse) GetDealReceivedOn() string`

GetDealReceivedOn returns the DealReceivedOn field if non-nil, zero value otherwise.

### GetDealReceivedOnOk

`func (o *DealsDetailsResponse) GetDealReceivedOnOk() (*string, bool)`

GetDealReceivedOnOk returns a tuple with the DealReceivedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealReceivedOn

`func (o *DealsDetailsResponse) SetDealReceivedOn(v string)`

SetDealReceivedOn sets DealReceivedOn field to given value.

### HasDealReceivedOn

`func (o *DealsDetailsResponse) HasDealReceivedOn() bool`

HasDealReceivedOn returns a boolean if a field has been set.

### GetDealExpiryDate

`func (o *DealsDetailsResponse) GetDealExpiryDate() string`

GetDealExpiryDate returns the DealExpiryDate field if non-nil, zero value otherwise.

### GetDealExpiryDateOk

`func (o *DealsDetailsResponse) GetDealExpiryDateOk() (*string, bool)`

GetDealExpiryDateOk returns a tuple with the DealExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealExpiryDate

`func (o *DealsDetailsResponse) SetDealExpiryDate(v string)`

SetDealExpiryDate sets DealExpiryDate field to given value.

### HasDealExpiryDate

`func (o *DealsDetailsResponse) HasDealExpiryDate() bool`

HasDealExpiryDate returns a boolean if a field has been set.

### GetPriceProtectionEndDate

`func (o *DealsDetailsResponse) GetPriceProtectionEndDate() string`

GetPriceProtectionEndDate returns the PriceProtectionEndDate field if non-nil, zero value otherwise.

### GetPriceProtectionEndDateOk

`func (o *DealsDetailsResponse) GetPriceProtectionEndDateOk() (*string, bool)`

GetPriceProtectionEndDateOk returns a tuple with the PriceProtectionEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtectionEndDate

`func (o *DealsDetailsResponse) SetPriceProtectionEndDate(v string)`

SetPriceProtectionEndDate sets PriceProtectionEndDate field to given value.

### HasPriceProtectionEndDate

`func (o *DealsDetailsResponse) HasPriceProtectionEndDate() bool`

HasPriceProtectionEndDate returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *DealsDetailsResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *DealsDetailsResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *DealsDetailsResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *DealsDetailsResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetEndUserInfo

`func (o *DealsDetailsResponse) GetEndUserInfo() []RenewalsDetailsResponseEndUserInfoInner`

GetEndUserInfo returns the EndUserInfo field if non-nil, zero value otherwise.

### GetEndUserInfoOk

`func (o *DealsDetailsResponse) GetEndUserInfoOk() (*[]RenewalsDetailsResponseEndUserInfoInner, bool)`

GetEndUserInfoOk returns a tuple with the EndUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserInfo

`func (o *DealsDetailsResponse) SetEndUserInfo(v []RenewalsDetailsResponseEndUserInfoInner)`

SetEndUserInfo sets EndUserInfo field to given value.

### HasEndUserInfo

`func (o *DealsDetailsResponse) HasEndUserInfo() bool`

HasEndUserInfo returns a boolean if a field has been set.

### GetProducts

`func (o *DealsDetailsResponse) GetProducts() []DealsDetailsResponseProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *DealsDetailsResponse) GetProductsOk() (*[]DealsDetailsResponseProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *DealsDetailsResponse) SetProducts(v []DealsDetailsResponseProductsInner)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *DealsDetailsResponse) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


