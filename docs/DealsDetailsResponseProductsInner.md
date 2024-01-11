# DealsDetailsResponseProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngramPartNumber** | Pointer to **string** | Unique Ingram part number. | [optional] 
**VendorPartNumber** | Pointer to **string** | Vendor Part number for the product. | [optional] 
**Upc** | Pointer to **string** | The UPC code for the product. Consists of 12 numeric digits that are uniquely assigned to each trade item. | [optional] 
**ProductDescription** | Pointer to **string** | Description of the product. | [optional] 
**Msrp** | Pointer to **float32** | Manufacturer Suggested Retail Price. | [optional] 
**ExtendedMSRP** | Pointer to **float32** | Extended MSRP - Manufacturer Suggested Retail Price X Quantity. | [optional] 
**StandardPrice** | Pointer to **float32** | Standard price of the line item. | [optional] 
**ApprovedQuantity** | Pointer to **int32** | Total quantity approved for the deal. | [optional] 
**RemainingQuantity** | Pointer to **int32** | The quantity remaining as part of the deal for the customer to order. | [optional] 
**Comments** | Pointer to **string** | Comments of the deal. | [optional] 
**SpecialConditions** | Pointer to **string** | Special conditions of the deal. | [optional] 
**StartDate** | Pointer to **string** | Start Date. | [optional] 
**ExpirationDate** | Pointer to **string** | Expiration date. | [optional] 
**DaysRemaining** | Pointer to **int32** | Number of days remaining before the deal expires. | [optional] 

## Methods

### NewDealsDetailsResponseProductsInner

`func NewDealsDetailsResponseProductsInner() *DealsDetailsResponseProductsInner`

NewDealsDetailsResponseProductsInner instantiates a new DealsDetailsResponseProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealsDetailsResponseProductsInnerWithDefaults

`func NewDealsDetailsResponseProductsInnerWithDefaults() *DealsDetailsResponseProductsInner`

NewDealsDetailsResponseProductsInnerWithDefaults instantiates a new DealsDetailsResponseProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngramPartNumber

`func (o *DealsDetailsResponseProductsInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *DealsDetailsResponseProductsInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *DealsDetailsResponseProductsInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *DealsDetailsResponseProductsInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *DealsDetailsResponseProductsInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *DealsDetailsResponseProductsInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *DealsDetailsResponseProductsInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *DealsDetailsResponseProductsInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetUpc

`func (o *DealsDetailsResponseProductsInner) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *DealsDetailsResponseProductsInner) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *DealsDetailsResponseProductsInner) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *DealsDetailsResponseProductsInner) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### GetProductDescription

`func (o *DealsDetailsResponseProductsInner) GetProductDescription() string`

GetProductDescription returns the ProductDescription field if non-nil, zero value otherwise.

### GetProductDescriptionOk

`func (o *DealsDetailsResponseProductsInner) GetProductDescriptionOk() (*string, bool)`

GetProductDescriptionOk returns a tuple with the ProductDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescription

`func (o *DealsDetailsResponseProductsInner) SetProductDescription(v string)`

SetProductDescription sets ProductDescription field to given value.

### HasProductDescription

`func (o *DealsDetailsResponseProductsInner) HasProductDescription() bool`

HasProductDescription returns a boolean if a field has been set.

### GetMsrp

`func (o *DealsDetailsResponseProductsInner) GetMsrp() float32`

GetMsrp returns the Msrp field if non-nil, zero value otherwise.

### GetMsrpOk

`func (o *DealsDetailsResponseProductsInner) GetMsrpOk() (*float32, bool)`

GetMsrpOk returns a tuple with the Msrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsrp

`func (o *DealsDetailsResponseProductsInner) SetMsrp(v float32)`

SetMsrp sets Msrp field to given value.

### HasMsrp

`func (o *DealsDetailsResponseProductsInner) HasMsrp() bool`

HasMsrp returns a boolean if a field has been set.

### GetExtendedMSRP

`func (o *DealsDetailsResponseProductsInner) GetExtendedMSRP() float32`

GetExtendedMSRP returns the ExtendedMSRP field if non-nil, zero value otherwise.

### GetExtendedMSRPOk

`func (o *DealsDetailsResponseProductsInner) GetExtendedMSRPOk() (*float32, bool)`

GetExtendedMSRPOk returns a tuple with the ExtendedMSRP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedMSRP

`func (o *DealsDetailsResponseProductsInner) SetExtendedMSRP(v float32)`

SetExtendedMSRP sets ExtendedMSRP field to given value.

### HasExtendedMSRP

`func (o *DealsDetailsResponseProductsInner) HasExtendedMSRP() bool`

HasExtendedMSRP returns a boolean if a field has been set.

### GetStandardPrice

`func (o *DealsDetailsResponseProductsInner) GetStandardPrice() float32`

GetStandardPrice returns the StandardPrice field if non-nil, zero value otherwise.

### GetStandardPriceOk

`func (o *DealsDetailsResponseProductsInner) GetStandardPriceOk() (*float32, bool)`

GetStandardPriceOk returns a tuple with the StandardPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPrice

`func (o *DealsDetailsResponseProductsInner) SetStandardPrice(v float32)`

SetStandardPrice sets StandardPrice field to given value.

### HasStandardPrice

`func (o *DealsDetailsResponseProductsInner) HasStandardPrice() bool`

HasStandardPrice returns a boolean if a field has been set.

### GetApprovedQuantity

`func (o *DealsDetailsResponseProductsInner) GetApprovedQuantity() int32`

GetApprovedQuantity returns the ApprovedQuantity field if non-nil, zero value otherwise.

### GetApprovedQuantityOk

`func (o *DealsDetailsResponseProductsInner) GetApprovedQuantityOk() (*int32, bool)`

GetApprovedQuantityOk returns a tuple with the ApprovedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedQuantity

`func (o *DealsDetailsResponseProductsInner) SetApprovedQuantity(v int32)`

SetApprovedQuantity sets ApprovedQuantity field to given value.

### HasApprovedQuantity

`func (o *DealsDetailsResponseProductsInner) HasApprovedQuantity() bool`

HasApprovedQuantity returns a boolean if a field has been set.

### GetRemainingQuantity

`func (o *DealsDetailsResponseProductsInner) GetRemainingQuantity() int32`

GetRemainingQuantity returns the RemainingQuantity field if non-nil, zero value otherwise.

### GetRemainingQuantityOk

`func (o *DealsDetailsResponseProductsInner) GetRemainingQuantityOk() (*int32, bool)`

GetRemainingQuantityOk returns a tuple with the RemainingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingQuantity

`func (o *DealsDetailsResponseProductsInner) SetRemainingQuantity(v int32)`

SetRemainingQuantity sets RemainingQuantity field to given value.

### HasRemainingQuantity

`func (o *DealsDetailsResponseProductsInner) HasRemainingQuantity() bool`

HasRemainingQuantity returns a boolean if a field has been set.

### GetComments

`func (o *DealsDetailsResponseProductsInner) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *DealsDetailsResponseProductsInner) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *DealsDetailsResponseProductsInner) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *DealsDetailsResponseProductsInner) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetSpecialConditions

`func (o *DealsDetailsResponseProductsInner) GetSpecialConditions() string`

GetSpecialConditions returns the SpecialConditions field if non-nil, zero value otherwise.

### GetSpecialConditionsOk

`func (o *DealsDetailsResponseProductsInner) GetSpecialConditionsOk() (*string, bool)`

GetSpecialConditionsOk returns a tuple with the SpecialConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialConditions

`func (o *DealsDetailsResponseProductsInner) SetSpecialConditions(v string)`

SetSpecialConditions sets SpecialConditions field to given value.

### HasSpecialConditions

`func (o *DealsDetailsResponseProductsInner) HasSpecialConditions() bool`

HasSpecialConditions returns a boolean if a field has been set.

### GetStartDate

`func (o *DealsDetailsResponseProductsInner) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DealsDetailsResponseProductsInner) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DealsDetailsResponseProductsInner) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DealsDetailsResponseProductsInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *DealsDetailsResponseProductsInner) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *DealsDetailsResponseProductsInner) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *DealsDetailsResponseProductsInner) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *DealsDetailsResponseProductsInner) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetDaysRemaining

`func (o *DealsDetailsResponseProductsInner) GetDaysRemaining() int32`

GetDaysRemaining returns the DaysRemaining field if non-nil, zero value otherwise.

### GetDaysRemainingOk

`func (o *DealsDetailsResponseProductsInner) GetDaysRemainingOk() (*int32, bool)`

GetDaysRemainingOk returns a tuple with the DaysRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysRemaining

`func (o *DealsDetailsResponseProductsInner) SetDaysRemaining(v int32)`

SetDaysRemaining sets DaysRemaining field to given value.

### HasDaysRemaining

`func (o *DealsDetailsResponseProductsInner) HasDaysRemaining() bool`

HasDaysRemaining returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


