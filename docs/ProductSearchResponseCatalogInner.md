# ProductSearchResponseCatalogInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the product. | [optional] 
**Category** | Pointer to **string** | The category of the product. Example: Displays. | [optional] 
**SubCategory** | Pointer to **string** | The sub category for the product. Example: ComputernMonitors. | [optional] 
**ProductType** | Pointer to **string** | The product type of the product. Example: LCD Monitors. | [optional] 
**IngramPartNumber** | Pointer to **string** | The Unique IngramMicro part number for the product. | [optional] 
**VendorPartNumber** | Pointer to **string** | The vendor part number for the product. | [optional] 
**UpcCode** | Pointer to **string** | The UPC code for the product. Consists of 12 numeric digits that are uniquly assigned to each trade item. | [optional] 
**VendorName** | Pointer to **string** | The name of the vendor/manufacturer of the product. | [optional] 
**EndUserRequired** | Pointer to **string** | Indicates whether the contact information for the end user/customer is required, which determines pricing and discounts. | [optional] 
**HasDiscounts** | Pointer to **string** | Specifies if there are discounts available for the product. | [optional] 
**Type** | Pointer to **string** | The SKU type of product. One of Physical, Digital, or Any. | [optional] 
**Discontinued** | Pointer to **string** | Indicates if the product has been discontinued. | [optional] 
**NewProduct** | Pointer to **string** | Indicates if the product is new. For digital products, newer than 10 days. For physical products, newer than 150 days. | [optional] 
**DirectShip** | Pointer to **string** | Indicates if the product will be shipped directly to the reseller or end user from the vendor/manufacturer. | [optional] 
**HasWarranty** | Pointer to **string** | Indicates if the product has a warranty. | [optional] 
**Links** | Pointer to [**[]ProductSearchResponseCatalogInnerLinksInner**](ProductSearchResponseCatalogInnerLinksInner.md) |  | [optional] 
**ExtraDescription** | Pointer to **string** | The extended description of the product. | [optional] 
**ReplacementSku** | Pointer to **string** | Identifies a SKU that is a comparable subsititution of the current SKU if available. | [optional] 
**AuthorizedToPurchase** | Pointer to **string** | It is true when it exists in matched queries field of ealstic search API. | [optional] 

## Methods

### NewProductSearchResponseCatalogInner

`func NewProductSearchResponseCatalogInner() *ProductSearchResponseCatalogInner`

NewProductSearchResponseCatalogInner instantiates a new ProductSearchResponseCatalogInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductSearchResponseCatalogInnerWithDefaults

`func NewProductSearchResponseCatalogInnerWithDefaults() *ProductSearchResponseCatalogInner`

NewProductSearchResponseCatalogInnerWithDefaults instantiates a new ProductSearchResponseCatalogInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ProductSearchResponseCatalogInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductSearchResponseCatalogInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductSearchResponseCatalogInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductSearchResponseCatalogInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *ProductSearchResponseCatalogInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ProductSearchResponseCatalogInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ProductSearchResponseCatalogInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ProductSearchResponseCatalogInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetSubCategory

`func (o *ProductSearchResponseCatalogInner) GetSubCategory() string`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *ProductSearchResponseCatalogInner) GetSubCategoryOk() (*string, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *ProductSearchResponseCatalogInner) SetSubCategory(v string)`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *ProductSearchResponseCatalogInner) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### GetProductType

`func (o *ProductSearchResponseCatalogInner) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ProductSearchResponseCatalogInner) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ProductSearchResponseCatalogInner) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *ProductSearchResponseCatalogInner) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *ProductSearchResponseCatalogInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *ProductSearchResponseCatalogInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *ProductSearchResponseCatalogInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *ProductSearchResponseCatalogInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *ProductSearchResponseCatalogInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *ProductSearchResponseCatalogInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *ProductSearchResponseCatalogInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *ProductSearchResponseCatalogInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetUpcCode

`func (o *ProductSearchResponseCatalogInner) GetUpcCode() string`

GetUpcCode returns the UpcCode field if non-nil, zero value otherwise.

### GetUpcCodeOk

`func (o *ProductSearchResponseCatalogInner) GetUpcCodeOk() (*string, bool)`

GetUpcCodeOk returns a tuple with the UpcCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcCode

`func (o *ProductSearchResponseCatalogInner) SetUpcCode(v string)`

SetUpcCode sets UpcCode field to given value.

### HasUpcCode

`func (o *ProductSearchResponseCatalogInner) HasUpcCode() bool`

HasUpcCode returns a boolean if a field has been set.

### GetVendorName

`func (o *ProductSearchResponseCatalogInner) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *ProductSearchResponseCatalogInner) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *ProductSearchResponseCatalogInner) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *ProductSearchResponseCatalogInner) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetEndUserRequired

`func (o *ProductSearchResponseCatalogInner) GetEndUserRequired() string`

GetEndUserRequired returns the EndUserRequired field if non-nil, zero value otherwise.

### GetEndUserRequiredOk

`func (o *ProductSearchResponseCatalogInner) GetEndUserRequiredOk() (*string, bool)`

GetEndUserRequiredOk returns a tuple with the EndUserRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserRequired

`func (o *ProductSearchResponseCatalogInner) SetEndUserRequired(v string)`

SetEndUserRequired sets EndUserRequired field to given value.

### HasEndUserRequired

`func (o *ProductSearchResponseCatalogInner) HasEndUserRequired() bool`

HasEndUserRequired returns a boolean if a field has been set.

### GetHasDiscounts

`func (o *ProductSearchResponseCatalogInner) GetHasDiscounts() string`

GetHasDiscounts returns the HasDiscounts field if non-nil, zero value otherwise.

### GetHasDiscountsOk

`func (o *ProductSearchResponseCatalogInner) GetHasDiscountsOk() (*string, bool)`

GetHasDiscountsOk returns a tuple with the HasDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDiscounts

`func (o *ProductSearchResponseCatalogInner) SetHasDiscounts(v string)`

SetHasDiscounts sets HasDiscounts field to given value.

### HasHasDiscounts

`func (o *ProductSearchResponseCatalogInner) HasHasDiscounts() bool`

HasHasDiscounts returns a boolean if a field has been set.

### GetType

`func (o *ProductSearchResponseCatalogInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductSearchResponseCatalogInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductSearchResponseCatalogInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProductSearchResponseCatalogInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDiscontinued

`func (o *ProductSearchResponseCatalogInner) GetDiscontinued() string`

GetDiscontinued returns the Discontinued field if non-nil, zero value otherwise.

### GetDiscontinuedOk

`func (o *ProductSearchResponseCatalogInner) GetDiscontinuedOk() (*string, bool)`

GetDiscontinuedOk returns a tuple with the Discontinued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscontinued

`func (o *ProductSearchResponseCatalogInner) SetDiscontinued(v string)`

SetDiscontinued sets Discontinued field to given value.

### HasDiscontinued

`func (o *ProductSearchResponseCatalogInner) HasDiscontinued() bool`

HasDiscontinued returns a boolean if a field has been set.

### GetNewProduct

`func (o *ProductSearchResponseCatalogInner) GetNewProduct() string`

GetNewProduct returns the NewProduct field if non-nil, zero value otherwise.

### GetNewProductOk

`func (o *ProductSearchResponseCatalogInner) GetNewProductOk() (*string, bool)`

GetNewProductOk returns a tuple with the NewProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewProduct

`func (o *ProductSearchResponseCatalogInner) SetNewProduct(v string)`

SetNewProduct sets NewProduct field to given value.

### HasNewProduct

`func (o *ProductSearchResponseCatalogInner) HasNewProduct() bool`

HasNewProduct returns a boolean if a field has been set.

### GetDirectShip

`func (o *ProductSearchResponseCatalogInner) GetDirectShip() string`

GetDirectShip returns the DirectShip field if non-nil, zero value otherwise.

### GetDirectShipOk

`func (o *ProductSearchResponseCatalogInner) GetDirectShipOk() (*string, bool)`

GetDirectShipOk returns a tuple with the DirectShip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectShip

`func (o *ProductSearchResponseCatalogInner) SetDirectShip(v string)`

SetDirectShip sets DirectShip field to given value.

### HasDirectShip

`func (o *ProductSearchResponseCatalogInner) HasDirectShip() bool`

HasDirectShip returns a boolean if a field has been set.

### GetHasWarranty

`func (o *ProductSearchResponseCatalogInner) GetHasWarranty() string`

GetHasWarranty returns the HasWarranty field if non-nil, zero value otherwise.

### GetHasWarrantyOk

`func (o *ProductSearchResponseCatalogInner) GetHasWarrantyOk() (*string, bool)`

GetHasWarrantyOk returns a tuple with the HasWarranty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWarranty

`func (o *ProductSearchResponseCatalogInner) SetHasWarranty(v string)`

SetHasWarranty sets HasWarranty field to given value.

### HasHasWarranty

`func (o *ProductSearchResponseCatalogInner) HasHasWarranty() bool`

HasHasWarranty returns a boolean if a field has been set.

### GetLinks

`func (o *ProductSearchResponseCatalogInner) GetLinks() []ProductSearchResponseCatalogInnerLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProductSearchResponseCatalogInner) GetLinksOk() (*[]ProductSearchResponseCatalogInnerLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProductSearchResponseCatalogInner) SetLinks(v []ProductSearchResponseCatalogInnerLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProductSearchResponseCatalogInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetExtraDescription

`func (o *ProductSearchResponseCatalogInner) GetExtraDescription() string`

GetExtraDescription returns the ExtraDescription field if non-nil, zero value otherwise.

### GetExtraDescriptionOk

`func (o *ProductSearchResponseCatalogInner) GetExtraDescriptionOk() (*string, bool)`

GetExtraDescriptionOk returns a tuple with the ExtraDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDescription

`func (o *ProductSearchResponseCatalogInner) SetExtraDescription(v string)`

SetExtraDescription sets ExtraDescription field to given value.

### HasExtraDescription

`func (o *ProductSearchResponseCatalogInner) HasExtraDescription() bool`

HasExtraDescription returns a boolean if a field has been set.

### GetReplacementSku

`func (o *ProductSearchResponseCatalogInner) GetReplacementSku() string`

GetReplacementSku returns the ReplacementSku field if non-nil, zero value otherwise.

### GetReplacementSkuOk

`func (o *ProductSearchResponseCatalogInner) GetReplacementSkuOk() (*string, bool)`

GetReplacementSkuOk returns a tuple with the ReplacementSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementSku

`func (o *ProductSearchResponseCatalogInner) SetReplacementSku(v string)`

SetReplacementSku sets ReplacementSku field to given value.

### HasReplacementSku

`func (o *ProductSearchResponseCatalogInner) HasReplacementSku() bool`

HasReplacementSku returns a boolean if a field has been set.

### GetAuthorizedToPurchase

`func (o *ProductSearchResponseCatalogInner) GetAuthorizedToPurchase() string`

GetAuthorizedToPurchase returns the AuthorizedToPurchase field if non-nil, zero value otherwise.

### GetAuthorizedToPurchaseOk

`func (o *ProductSearchResponseCatalogInner) GetAuthorizedToPurchaseOk() (*string, bool)`

GetAuthorizedToPurchaseOk returns a tuple with the AuthorizedToPurchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedToPurchase

`func (o *ProductSearchResponseCatalogInner) SetAuthorizedToPurchase(v string)`

SetAuthorizedToPurchase sets AuthorizedToPurchase field to given value.

### HasAuthorizedToPurchase

`func (o *ProductSearchResponseCatalogInner) HasAuthorizedToPurchase() bool`

HasAuthorizedToPurchase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


