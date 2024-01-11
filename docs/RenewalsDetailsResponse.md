# RenewalsDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RenewalId** | Pointer to **string** | Unique Ingram renewal ID. | [optional] 
**IngramOrderNumber** | Pointer to **string** | The IngramMicro sales order number. | [optional] 
**IngramOrderDate** | Pointer to **string** | The IngramMicro sales order date. | [optional] 
**ExpirationDate** | Pointer to **string** | Renewal expiration date. | [optional] 
**IngramPurchaseOrderNumber** | Pointer to **string** | Ingram purchase order number. | [optional] 
**CustomerOrderNumber** | Pointer to **string** | The reseller&#39;s order number for reference in their system. | [optional] 
**EndCustomerOrderNumber** | Pointer to **string** | The end customer&#39;s order number for reference in their system. | [optional] 
**RenewalValue** | Pointer to **float32** | The value of the renewal. | [optional] 
**EndUser** | Pointer to **string** | The company name for the end user/customer. | [optional] 
**Vendor** | Pointer to **string** | The name of the vendor. | [optional] 
**Status** | Pointer to **string** | The status of the renewal. | [optional] 
**EndUserInfo** | Pointer to [**[]RenewalsDetailsResponseEndUserInfoInner**](RenewalsDetailsResponseEndUserInfoInner.md) |  | [optional] 
**ReferenceNumber** | Pointer to [**[]RenewalsDetailsResponseReferenceNumberInner**](RenewalsDetailsResponseReferenceNumberInner.md) |  | [optional] 
**Products** | Pointer to [**[]RenewalsDetailsResponseProductsInner**](RenewalsDetailsResponseProductsInner.md) |  | [optional] 
**AdditionalAttributes** | Pointer to [**[]RenewalsDetailsResponseAdditionalAttributesInner**](RenewalsDetailsResponseAdditionalAttributesInner.md) |  | [optional] 

## Methods

### NewRenewalsDetailsResponse

`func NewRenewalsDetailsResponse() *RenewalsDetailsResponse`

NewRenewalsDetailsResponse instantiates a new RenewalsDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenewalsDetailsResponseWithDefaults

`func NewRenewalsDetailsResponseWithDefaults() *RenewalsDetailsResponse`

NewRenewalsDetailsResponseWithDefaults instantiates a new RenewalsDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRenewalId

`func (o *RenewalsDetailsResponse) GetRenewalId() string`

GetRenewalId returns the RenewalId field if non-nil, zero value otherwise.

### GetRenewalIdOk

`func (o *RenewalsDetailsResponse) GetRenewalIdOk() (*string, bool)`

GetRenewalIdOk returns a tuple with the RenewalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalId

`func (o *RenewalsDetailsResponse) SetRenewalId(v string)`

SetRenewalId sets RenewalId field to given value.

### HasRenewalId

`func (o *RenewalsDetailsResponse) HasRenewalId() bool`

HasRenewalId returns a boolean if a field has been set.

### GetIngramOrderNumber

`func (o *RenewalsDetailsResponse) GetIngramOrderNumber() string`

GetIngramOrderNumber returns the IngramOrderNumber field if non-nil, zero value otherwise.

### GetIngramOrderNumberOk

`func (o *RenewalsDetailsResponse) GetIngramOrderNumberOk() (*string, bool)`

GetIngramOrderNumberOk returns a tuple with the IngramOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramOrderNumber

`func (o *RenewalsDetailsResponse) SetIngramOrderNumber(v string)`

SetIngramOrderNumber sets IngramOrderNumber field to given value.

### HasIngramOrderNumber

`func (o *RenewalsDetailsResponse) HasIngramOrderNumber() bool`

HasIngramOrderNumber returns a boolean if a field has been set.

### GetIngramOrderDate

`func (o *RenewalsDetailsResponse) GetIngramOrderDate() string`

GetIngramOrderDate returns the IngramOrderDate field if non-nil, zero value otherwise.

### GetIngramOrderDateOk

`func (o *RenewalsDetailsResponse) GetIngramOrderDateOk() (*string, bool)`

GetIngramOrderDateOk returns a tuple with the IngramOrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramOrderDate

`func (o *RenewalsDetailsResponse) SetIngramOrderDate(v string)`

SetIngramOrderDate sets IngramOrderDate field to given value.

### HasIngramOrderDate

`func (o *RenewalsDetailsResponse) HasIngramOrderDate() bool`

HasIngramOrderDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *RenewalsDetailsResponse) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *RenewalsDetailsResponse) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *RenewalsDetailsResponse) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *RenewalsDetailsResponse) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetIngramPurchaseOrderNumber

`func (o *RenewalsDetailsResponse) GetIngramPurchaseOrderNumber() string`

GetIngramPurchaseOrderNumber returns the IngramPurchaseOrderNumber field if non-nil, zero value otherwise.

### GetIngramPurchaseOrderNumberOk

`func (o *RenewalsDetailsResponse) GetIngramPurchaseOrderNumberOk() (*string, bool)`

GetIngramPurchaseOrderNumberOk returns a tuple with the IngramPurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPurchaseOrderNumber

`func (o *RenewalsDetailsResponse) SetIngramPurchaseOrderNumber(v string)`

SetIngramPurchaseOrderNumber sets IngramPurchaseOrderNumber field to given value.

### HasIngramPurchaseOrderNumber

`func (o *RenewalsDetailsResponse) HasIngramPurchaseOrderNumber() bool`

HasIngramPurchaseOrderNumber returns a boolean if a field has been set.

### GetCustomerOrderNumber

`func (o *RenewalsDetailsResponse) GetCustomerOrderNumber() string`

GetCustomerOrderNumber returns the CustomerOrderNumber field if non-nil, zero value otherwise.

### GetCustomerOrderNumberOk

`func (o *RenewalsDetailsResponse) GetCustomerOrderNumberOk() (*string, bool)`

GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOrderNumber

`func (o *RenewalsDetailsResponse) SetCustomerOrderNumber(v string)`

SetCustomerOrderNumber sets CustomerOrderNumber field to given value.

### HasCustomerOrderNumber

`func (o *RenewalsDetailsResponse) HasCustomerOrderNumber() bool`

HasCustomerOrderNumber returns a boolean if a field has been set.

### GetEndCustomerOrderNumber

`func (o *RenewalsDetailsResponse) GetEndCustomerOrderNumber() string`

GetEndCustomerOrderNumber returns the EndCustomerOrderNumber field if non-nil, zero value otherwise.

### GetEndCustomerOrderNumberOk

`func (o *RenewalsDetailsResponse) GetEndCustomerOrderNumberOk() (*string, bool)`

GetEndCustomerOrderNumberOk returns a tuple with the EndCustomerOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCustomerOrderNumber

`func (o *RenewalsDetailsResponse) SetEndCustomerOrderNumber(v string)`

SetEndCustomerOrderNumber sets EndCustomerOrderNumber field to given value.

### HasEndCustomerOrderNumber

`func (o *RenewalsDetailsResponse) HasEndCustomerOrderNumber() bool`

HasEndCustomerOrderNumber returns a boolean if a field has been set.

### GetRenewalValue

`func (o *RenewalsDetailsResponse) GetRenewalValue() float32`

GetRenewalValue returns the RenewalValue field if non-nil, zero value otherwise.

### GetRenewalValueOk

`func (o *RenewalsDetailsResponse) GetRenewalValueOk() (*float32, bool)`

GetRenewalValueOk returns a tuple with the RenewalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalValue

`func (o *RenewalsDetailsResponse) SetRenewalValue(v float32)`

SetRenewalValue sets RenewalValue field to given value.

### HasRenewalValue

`func (o *RenewalsDetailsResponse) HasRenewalValue() bool`

HasRenewalValue returns a boolean if a field has been set.

### GetEndUser

`func (o *RenewalsDetailsResponse) GetEndUser() string`

GetEndUser returns the EndUser field if non-nil, zero value otherwise.

### GetEndUserOk

`func (o *RenewalsDetailsResponse) GetEndUserOk() (*string, bool)`

GetEndUserOk returns a tuple with the EndUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUser

`func (o *RenewalsDetailsResponse) SetEndUser(v string)`

SetEndUser sets EndUser field to given value.

### HasEndUser

`func (o *RenewalsDetailsResponse) HasEndUser() bool`

HasEndUser returns a boolean if a field has been set.

### GetVendor

`func (o *RenewalsDetailsResponse) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *RenewalsDetailsResponse) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *RenewalsDetailsResponse) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *RenewalsDetailsResponse) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetStatus

`func (o *RenewalsDetailsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RenewalsDetailsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RenewalsDetailsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RenewalsDetailsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEndUserInfo

`func (o *RenewalsDetailsResponse) GetEndUserInfo() []RenewalsDetailsResponseEndUserInfoInner`

GetEndUserInfo returns the EndUserInfo field if non-nil, zero value otherwise.

### GetEndUserInfoOk

`func (o *RenewalsDetailsResponse) GetEndUserInfoOk() (*[]RenewalsDetailsResponseEndUserInfoInner, bool)`

GetEndUserInfoOk returns a tuple with the EndUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserInfo

`func (o *RenewalsDetailsResponse) SetEndUserInfo(v []RenewalsDetailsResponseEndUserInfoInner)`

SetEndUserInfo sets EndUserInfo field to given value.

### HasEndUserInfo

`func (o *RenewalsDetailsResponse) HasEndUserInfo() bool`

HasEndUserInfo returns a boolean if a field has been set.

### GetReferenceNumber

`func (o *RenewalsDetailsResponse) GetReferenceNumber() []RenewalsDetailsResponseReferenceNumberInner`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *RenewalsDetailsResponse) GetReferenceNumberOk() (*[]RenewalsDetailsResponseReferenceNumberInner, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *RenewalsDetailsResponse) SetReferenceNumber(v []RenewalsDetailsResponseReferenceNumberInner)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *RenewalsDetailsResponse) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### GetProducts

`func (o *RenewalsDetailsResponse) GetProducts() []RenewalsDetailsResponseProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *RenewalsDetailsResponse) GetProductsOk() (*[]RenewalsDetailsResponseProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *RenewalsDetailsResponse) SetProducts(v []RenewalsDetailsResponseProductsInner)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *RenewalsDetailsResponse) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetAdditionalAttributes

`func (o *RenewalsDetailsResponse) GetAdditionalAttributes() []RenewalsDetailsResponseAdditionalAttributesInner`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *RenewalsDetailsResponse) GetAdditionalAttributesOk() (*[]RenewalsDetailsResponseAdditionalAttributesInner, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *RenewalsDetailsResponse) SetAdditionalAttributes(v []RenewalsDetailsResponseAdditionalAttributesInner)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *RenewalsDetailsResponse) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


