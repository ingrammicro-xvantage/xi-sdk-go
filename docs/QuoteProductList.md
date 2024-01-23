# QuoteProductList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteProductGuid** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**BidStartDate** | Pointer to **string** |  | [optional] 
**BidExpiryDate** | Pointer to **string** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**LineNumber** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**VendorPartNumber** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **string** |  | [optional] 
**IsSuggestionProduct** | Pointer to **string** |  | [optional] 
**VpnCategory** | Pointer to **string** |  | [optional] 
**QuoteProductsSupplierPartAuxiliaryId** | Pointer to **string** |  | [optional] 
**QuoteProductsVendor** | Pointer to **string** |  | [optional] 
**Price** | Pointer to [**QuoteProductListPrice**](QuoteProductListPrice.md) |  | [optional] 

## Methods

### NewQuoteProductList

`func NewQuoteProductList() *QuoteProductList`

NewQuoteProductList instantiates a new QuoteProductList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteProductListWithDefaults

`func NewQuoteProductListWithDefaults() *QuoteProductList`

NewQuoteProductListWithDefaults instantiates a new QuoteProductList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteProductGuid

`func (o *QuoteProductList) GetQuoteProductGuid() string`

GetQuoteProductGuid returns the QuoteProductGuid field if non-nil, zero value otherwise.

### GetQuoteProductGuidOk

`func (o *QuoteProductList) GetQuoteProductGuidOk() (*string, bool)`

GetQuoteProductGuidOk returns a tuple with the QuoteProductGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteProductGuid

`func (o *QuoteProductList) SetQuoteProductGuid(v string)`

SetQuoteProductGuid sets QuoteProductGuid field to given value.

### HasQuoteProductGuid

`func (o *QuoteProductList) HasQuoteProductGuid() bool`

HasQuoteProductGuid returns a boolean if a field has been set.

### GetQuantity

`func (o *QuoteProductList) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuoteProductList) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuoteProductList) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuoteProductList) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetComments

`func (o *QuoteProductList) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *QuoteProductList) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *QuoteProductList) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *QuoteProductList) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetBidStartDate

`func (o *QuoteProductList) GetBidStartDate() string`

GetBidStartDate returns the BidStartDate field if non-nil, zero value otherwise.

### GetBidStartDateOk

`func (o *QuoteProductList) GetBidStartDateOk() (*string, bool)`

GetBidStartDateOk returns a tuple with the BidStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidStartDate

`func (o *QuoteProductList) SetBidStartDate(v string)`

SetBidStartDate sets BidStartDate field to given value.

### HasBidStartDate

`func (o *QuoteProductList) HasBidStartDate() bool`

HasBidStartDate returns a boolean if a field has been set.

### GetBidExpiryDate

`func (o *QuoteProductList) GetBidExpiryDate() string`

GetBidExpiryDate returns the BidExpiryDate field if non-nil, zero value otherwise.

### GetBidExpiryDateOk

`func (o *QuoteProductList) GetBidExpiryDateOk() (*string, bool)`

GetBidExpiryDateOk returns a tuple with the BidExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidExpiryDate

`func (o *QuoteProductList) SetBidExpiryDate(v string)`

SetBidExpiryDate sets BidExpiryDate field to given value.

### HasBidExpiryDate

`func (o *QuoteProductList) HasBidExpiryDate() bool`

HasBidExpiryDate returns a boolean if a field has been set.

### GetSku

`func (o *QuoteProductList) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *QuoteProductList) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *QuoteProductList) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *QuoteProductList) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetLineNumber

`func (o *QuoteProductList) GetLineNumber() string`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *QuoteProductList) GetLineNumberOk() (*string, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *QuoteProductList) SetLineNumber(v string)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *QuoteProductList) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### GetDescription

`func (o *QuoteProductList) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuoteProductList) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuoteProductList) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuoteProductList) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *QuoteProductList) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *QuoteProductList) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *QuoteProductList) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *QuoteProductList) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetWeight

`func (o *QuoteProductList) GetWeight() string`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *QuoteProductList) GetWeightOk() (*string, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *QuoteProductList) SetWeight(v string)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *QuoteProductList) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetIsSuggestionProduct

`func (o *QuoteProductList) GetIsSuggestionProduct() string`

GetIsSuggestionProduct returns the IsSuggestionProduct field if non-nil, zero value otherwise.

### GetIsSuggestionProductOk

`func (o *QuoteProductList) GetIsSuggestionProductOk() (*string, bool)`

GetIsSuggestionProductOk returns a tuple with the IsSuggestionProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuggestionProduct

`func (o *QuoteProductList) SetIsSuggestionProduct(v string)`

SetIsSuggestionProduct sets IsSuggestionProduct field to given value.

### HasIsSuggestionProduct

`func (o *QuoteProductList) HasIsSuggestionProduct() bool`

HasIsSuggestionProduct returns a boolean if a field has been set.

### GetVpnCategory

`func (o *QuoteProductList) GetVpnCategory() string`

GetVpnCategory returns the VpnCategory field if non-nil, zero value otherwise.

### GetVpnCategoryOk

`func (o *QuoteProductList) GetVpnCategoryOk() (*string, bool)`

GetVpnCategoryOk returns a tuple with the VpnCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnCategory

`func (o *QuoteProductList) SetVpnCategory(v string)`

SetVpnCategory sets VpnCategory field to given value.

### HasVpnCategory

`func (o *QuoteProductList) HasVpnCategory() bool`

HasVpnCategory returns a boolean if a field has been set.

### GetQuoteProductsSupplierPartAuxiliaryId

`func (o *QuoteProductList) GetQuoteProductsSupplierPartAuxiliaryId() string`

GetQuoteProductsSupplierPartAuxiliaryId returns the QuoteProductsSupplierPartAuxiliaryId field if non-nil, zero value otherwise.

### GetQuoteProductsSupplierPartAuxiliaryIdOk

`func (o *QuoteProductList) GetQuoteProductsSupplierPartAuxiliaryIdOk() (*string, bool)`

GetQuoteProductsSupplierPartAuxiliaryIdOk returns a tuple with the QuoteProductsSupplierPartAuxiliaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteProductsSupplierPartAuxiliaryId

`func (o *QuoteProductList) SetQuoteProductsSupplierPartAuxiliaryId(v string)`

SetQuoteProductsSupplierPartAuxiliaryId sets QuoteProductsSupplierPartAuxiliaryId field to given value.

### HasQuoteProductsSupplierPartAuxiliaryId

`func (o *QuoteProductList) HasQuoteProductsSupplierPartAuxiliaryId() bool`

HasQuoteProductsSupplierPartAuxiliaryId returns a boolean if a field has been set.

### GetQuoteProductsVendor

`func (o *QuoteProductList) GetQuoteProductsVendor() string`

GetQuoteProductsVendor returns the QuoteProductsVendor field if non-nil, zero value otherwise.

### GetQuoteProductsVendorOk

`func (o *QuoteProductList) GetQuoteProductsVendorOk() (*string, bool)`

GetQuoteProductsVendorOk returns a tuple with the QuoteProductsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteProductsVendor

`func (o *QuoteProductList) SetQuoteProductsVendor(v string)`

SetQuoteProductsVendor sets QuoteProductsVendor field to given value.

### HasQuoteProductsVendor

`func (o *QuoteProductList) HasQuoteProductsVendor() bool`

HasQuoteProductsVendor returns a boolean if a field has been set.

### GetPrice

`func (o *QuoteProductList) GetPrice() QuoteProductListPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *QuoteProductList) GetPriceOk() (*QuoteProductListPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *QuoteProductList) SetPrice(v QuoteProductListPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *QuoteProductList) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


