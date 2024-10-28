# ProductDetailResponseIndicatorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasWarranty** | Pointer to **bool** | Boolean that indicates whether product has a warranty. | [optional] 
**IsNewProduct** | Pointer to **bool** | Boolean that indicates whether it’s a new product.  | [optional] 
**HasReturnLimits** | Pointer to **bool** | Boolean that indicates whether there is any limit to return the product. | [optional] 
**IsBackOrderAllowed** | Pointer to **bool** | Boolean that indicates whether back order is allowed for the product. | [optional] 
**IsShippedFromPartner** | Pointer to **bool** | Boolean that indicates whether product is shipped from the partner. | [optional] 
**IsReplacementProduct** | Pointer to **bool** | Boolean that indicates whether product is a replacement product. | [optional] 
**IsDirectship** | Pointer to **bool** | Boolean that indicates whether it’s a direct ship product. | [optional] 
**IsDownloadable** | Pointer to **bool** | Boolean that indicates whether product is downloadable. | [optional] 
**IsDigitalType** | Pointer to **bool** | Boolean that indicates whether it’s a digital product.  | [optional] 
**SkuType** | Pointer to **string** | skutype | [optional] 
**HasStdSpecialPrice** | Pointer to **bool** | Boolean that indicates whether product has any standard special price. | [optional] 
**HasAcopSpecialPrice** | Pointer to **bool** | Boolean that indicates whether product has any ACOP special price. | [optional] 
**HasAcopQuantityBreak** | Pointer to **bool** | Boolean that indicates whether product has any ACOP quantity break. | [optional] 
**HasStdWebDiscount** | Pointer to **bool** | Boolean that indicates whether product has any standard web discount. | [optional] 
**HasSpecialBid** | Pointer to **bool** | Boolean that indicates whether product has any special bid. | [optional] 
**IsExportableToCountry** | Pointer to **bool** | Boolean that indicates whether product is exportable. | [optional] 
**IsDiscontinuedProduct** | Pointer to **bool** | Boolean that indicates whether it’s a discontinued product. | [optional] 
**IsRefurbishedProduct** | Pointer to **bool** | Boolean that indicates whether product is refurbished. | [optional] 
**IsReturnableProduct** | Pointer to **bool** | Boolean that indicates if the product can be returned. | [optional] 
**IsIngramShip** | Pointer to **bool** | Boolean that indicates whether it’s a Ingram shipped product. | [optional] 
**IsEnduserRequired** | Pointer to **bool** | Do vendor requires Enduser name required to create an order. | [optional] 
**IsHeavyWeight** | Pointer to **bool** | Boolean that indicates whether it’s  heavy weight product. | [optional] 
**HasLtl** | Pointer to **bool** | Boolean that indicates whether it hasLtl or not. | [optional] 
**IsClearanceProduct** | Pointer to **bool** | Boolean that indicates whether it’s clearnce product. | [optional] 
**HasBundle** | Pointer to **bool** | Boolean that indicates whether it’s a bundled product. | [optional] 
**IsOversizeProduct** | Pointer to **bool** | Boolean that indicates whether it’s oversized product. | [optional] 
**IsPreorderProduct** | Pointer to **bool** | Boolean that indicates whether it’s a preorder product. | [optional] 
**IsLicenseProduct** | Pointer to **bool** | Boolean that indicates whether it’s a licened product. | [optional] 
**IsDirectshipOrderable** | Pointer to **bool** | Boolean that indicates whether product is directship orderable. | [optional] 
**IsServiceSku** | Pointer to **bool** | Boolean that indicates whether product is service SKU. | [optional] 
**IsConfigurable** | Pointer to **bool** | Boolean that indicates whether product is configurable. | [optional] 

## Methods

### NewProductDetailResponseIndicatorsInner

`func NewProductDetailResponseIndicatorsInner() *ProductDetailResponseIndicatorsInner`

NewProductDetailResponseIndicatorsInner instantiates a new ProductDetailResponseIndicatorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductDetailResponseIndicatorsInnerWithDefaults

`func NewProductDetailResponseIndicatorsInnerWithDefaults() *ProductDetailResponseIndicatorsInner`

NewProductDetailResponseIndicatorsInnerWithDefaults instantiates a new ProductDetailResponseIndicatorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasWarranty

`func (o *ProductDetailResponseIndicatorsInner) GetHasWarranty() bool`

GetHasWarranty returns the HasWarranty field if non-nil, zero value otherwise.

### GetHasWarrantyOk

`func (o *ProductDetailResponseIndicatorsInner) GetHasWarrantyOk() (*bool, bool)`

GetHasWarrantyOk returns a tuple with the HasWarranty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWarranty

`func (o *ProductDetailResponseIndicatorsInner) SetHasWarranty(v bool)`

SetHasWarranty sets HasWarranty field to given value.

### HasHasWarranty

`func (o *ProductDetailResponseIndicatorsInner) HasHasWarranty() bool`

HasHasWarranty returns a boolean if a field has been set.

### GetIsNewProduct

`func (o *ProductDetailResponseIndicatorsInner) GetIsNewProduct() bool`

GetIsNewProduct returns the IsNewProduct field if non-nil, zero value otherwise.

### GetIsNewProductOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsNewProductOk() (*bool, bool)`

GetIsNewProductOk returns a tuple with the IsNewProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNewProduct

`func (o *ProductDetailResponseIndicatorsInner) SetIsNewProduct(v bool)`

SetIsNewProduct sets IsNewProduct field to given value.

### HasIsNewProduct

`func (o *ProductDetailResponseIndicatorsInner) HasIsNewProduct() bool`

HasIsNewProduct returns a boolean if a field has been set.

### GetHasReturnLimits

`func (o *ProductDetailResponseIndicatorsInner) GetHasReturnLimits() bool`

GetHasReturnLimits returns the HasReturnLimits field if non-nil, zero value otherwise.

### GetHasReturnLimitsOk

`func (o *ProductDetailResponseIndicatorsInner) GetHasReturnLimitsOk() (*bool, bool)`

GetHasReturnLimitsOk returns a tuple with the HasReturnLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReturnLimits

`func (o *ProductDetailResponseIndicatorsInner) SetHasReturnLimits(v bool)`

SetHasReturnLimits sets HasReturnLimits field to given value.

### HasHasReturnLimits

`func (o *ProductDetailResponseIndicatorsInner) HasHasReturnLimits() bool`

HasHasReturnLimits returns a boolean if a field has been set.

### GetIsBackOrderAllowed

`func (o *ProductDetailResponseIndicatorsInner) GetIsBackOrderAllowed() bool`

GetIsBackOrderAllowed returns the IsBackOrderAllowed field if non-nil, zero value otherwise.

### GetIsBackOrderAllowedOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsBackOrderAllowedOk() (*bool, bool)`

GetIsBackOrderAllowedOk returns a tuple with the IsBackOrderAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackOrderAllowed

`func (o *ProductDetailResponseIndicatorsInner) SetIsBackOrderAllowed(v bool)`

SetIsBackOrderAllowed sets IsBackOrderAllowed field to given value.

### HasIsBackOrderAllowed

`func (o *ProductDetailResponseIndicatorsInner) HasIsBackOrderAllowed() bool`

HasIsBackOrderAllowed returns a boolean if a field has been set.

### GetIsShippedFromPartner

`func (o *ProductDetailResponseIndicatorsInner) GetIsShippedFromPartner() bool`

GetIsShippedFromPartner returns the IsShippedFromPartner field if non-nil, zero value otherwise.

### GetIsShippedFromPartnerOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsShippedFromPartnerOk() (*bool, bool)`

GetIsShippedFromPartnerOk returns a tuple with the IsShippedFromPartner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShippedFromPartner

`func (o *ProductDetailResponseIndicatorsInner) SetIsShippedFromPartner(v bool)`

SetIsShippedFromPartner sets IsShippedFromPartner field to given value.

### HasIsShippedFromPartner

`func (o *ProductDetailResponseIndicatorsInner) HasIsShippedFromPartner() bool`

HasIsShippedFromPartner returns a boolean if a field has been set.

### GetIsReplacementProduct

`func (o *ProductDetailResponseIndicatorsInner) GetIsReplacementProduct() bool`

GetIsReplacementProduct returns the IsReplacementProduct field if non-nil, zero value otherwise.

### GetIsReplacementProductOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsReplacementProductOk() (*bool, bool)`

GetIsReplacementProductOk returns a tuple with the IsReplacementProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReplacementProduct

`func (o *ProductDetailResponseIndicatorsInner) SetIsReplacementProduct(v bool)`

SetIsReplacementProduct sets IsReplacementProduct field to given value.

### HasIsReplacementProduct

`func (o *ProductDetailResponseIndicatorsInner) HasIsReplacementProduct() bool`

HasIsReplacementProduct returns a boolean if a field has been set.

### GetIsDirectship

`func (o *ProductDetailResponseIndicatorsInner) GetIsDirectship() bool`

GetIsDirectship returns the IsDirectship field if non-nil, zero value otherwise.

### GetIsDirectshipOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsDirectshipOk() (*bool, bool)`

GetIsDirectshipOk returns a tuple with the IsDirectship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectship

`func (o *ProductDetailResponseIndicatorsInner) SetIsDirectship(v bool)`

SetIsDirectship sets IsDirectship field to given value.

### HasIsDirectship

`func (o *ProductDetailResponseIndicatorsInner) HasIsDirectship() bool`

HasIsDirectship returns a boolean if a field has been set.

### GetIsDownloadable

`func (o *ProductDetailResponseIndicatorsInner) GetIsDownloadable() bool`

GetIsDownloadable returns the IsDownloadable field if non-nil, zero value otherwise.

### GetIsDownloadableOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsDownloadableOk() (*bool, bool)`

GetIsDownloadableOk returns a tuple with the IsDownloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDownloadable

`func (o *ProductDetailResponseIndicatorsInner) SetIsDownloadable(v bool)`

SetIsDownloadable sets IsDownloadable field to given value.

### HasIsDownloadable

`func (o *ProductDetailResponseIndicatorsInner) HasIsDownloadable() bool`

HasIsDownloadable returns a boolean if a field has been set.

### GetIsDigitalType

`func (o *ProductDetailResponseIndicatorsInner) GetIsDigitalType() bool`

GetIsDigitalType returns the IsDigitalType field if non-nil, zero value otherwise.

### GetIsDigitalTypeOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsDigitalTypeOk() (*bool, bool)`

GetIsDigitalTypeOk returns a tuple with the IsDigitalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDigitalType

`func (o *ProductDetailResponseIndicatorsInner) SetIsDigitalType(v bool)`

SetIsDigitalType sets IsDigitalType field to given value.

### HasIsDigitalType

`func (o *ProductDetailResponseIndicatorsInner) HasIsDigitalType() bool`

HasIsDigitalType returns a boolean if a field has been set.

### GetSkuType

`func (o *ProductDetailResponseIndicatorsInner) GetSkuType() string`

GetSkuType returns the SkuType field if non-nil, zero value otherwise.

### GetSkuTypeOk

`func (o *ProductDetailResponseIndicatorsInner) GetSkuTypeOk() (*string, bool)`

GetSkuTypeOk returns a tuple with the SkuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuType

`func (o *ProductDetailResponseIndicatorsInner) SetSkuType(v string)`

SetSkuType sets SkuType field to given value.

### HasSkuType

`func (o *ProductDetailResponseIndicatorsInner) HasSkuType() bool`

HasSkuType returns a boolean if a field has been set.

### GetHasStdSpecialPrice

`func (o *ProductDetailResponseIndicatorsInner) GetHasStdSpecialPrice() bool`

GetHasStdSpecialPrice returns the HasStdSpecialPrice field if non-nil, zero value otherwise.

### GetHasStdSpecialPriceOk

`func (o *ProductDetailResponseIndicatorsInner) GetHasStdSpecialPriceOk() (*bool, bool)`

GetHasStdSpecialPriceOk returns a tuple with the HasStdSpecialPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasStdSpecialPrice

`func (o *ProductDetailResponseIndicatorsInner) SetHasStdSpecialPrice(v bool)`

SetHasStdSpecialPrice sets HasStdSpecialPrice field to given value.

### HasHasStdSpecialPrice

`func (o *ProductDetailResponseIndicatorsInner) HasHasStdSpecialPrice() bool`

HasHasStdSpecialPrice returns a boolean if a field has been set.

### GetHasAcopSpecialPrice

`func (o *ProductDetailResponseIndicatorsInner) GetHasAcopSpecialPrice() bool`

GetHasAcopSpecialPrice returns the HasAcopSpecialPrice field if non-nil, zero value otherwise.

### GetHasAcopSpecialPriceOk

`func (o *ProductDetailResponseIndicatorsInner) GetHasAcopSpecialPriceOk() (*bool, bool)`

GetHasAcopSpecialPriceOk returns a tuple with the HasAcopSpecialPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAcopSpecialPrice

`func (o *ProductDetailResponseIndicatorsInner) SetHasAcopSpecialPrice(v bool)`

SetHasAcopSpecialPrice sets HasAcopSpecialPrice field to given value.

### HasHasAcopSpecialPrice

`func (o *ProductDetailResponseIndicatorsInner) HasHasAcopSpecialPrice() bool`

HasHasAcopSpecialPrice returns a boolean if a field has been set.

### GetHasAcopQuantityBreak

`func (o *ProductDetailResponseIndicatorsInner) GetHasAcopQuantityBreak() bool`

GetHasAcopQuantityBreak returns the HasAcopQuantityBreak field if non-nil, zero value otherwise.

### GetHasAcopQuantityBreakOk

`func (o *ProductDetailResponseIndicatorsInner) GetHasAcopQuantityBreakOk() (*bool, bool)`

GetHasAcopQuantityBreakOk returns a tuple with the HasAcopQuantityBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAcopQuantityBreak

`func (o *ProductDetailResponseIndicatorsInner) SetHasAcopQuantityBreak(v bool)`

SetHasAcopQuantityBreak sets HasAcopQuantityBreak field to given value.

### HasHasAcopQuantityBreak

`func (o *ProductDetailResponseIndicatorsInner) HasHasAcopQuantityBreak() bool`

HasHasAcopQuantityBreak returns a boolean if a field has been set.

### GetHasStdWebDiscount

`func (o *ProductDetailResponseIndicatorsInner) GetHasStdWebDiscount() bool`

GetHasStdWebDiscount returns the HasStdWebDiscount field if non-nil, zero value otherwise.

### GetHasStdWebDiscountOk

`func (o *ProductDetailResponseIndicatorsInner) GetHasStdWebDiscountOk() (*bool, bool)`

GetHasStdWebDiscountOk returns a tuple with the HasStdWebDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasStdWebDiscount

`func (o *ProductDetailResponseIndicatorsInner) SetHasStdWebDiscount(v bool)`

SetHasStdWebDiscount sets HasStdWebDiscount field to given value.

### HasHasStdWebDiscount

`func (o *ProductDetailResponseIndicatorsInner) HasHasStdWebDiscount() bool`

HasHasStdWebDiscount returns a boolean if a field has been set.

### GetHasSpecialBid

`func (o *ProductDetailResponseIndicatorsInner) GetHasSpecialBid() bool`

GetHasSpecialBid returns the HasSpecialBid field if non-nil, zero value otherwise.

### GetHasSpecialBidOk

`func (o *ProductDetailResponseIndicatorsInner) GetHasSpecialBidOk() (*bool, bool)`

GetHasSpecialBidOk returns a tuple with the HasSpecialBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSpecialBid

`func (o *ProductDetailResponseIndicatorsInner) SetHasSpecialBid(v bool)`

SetHasSpecialBid sets HasSpecialBid field to given value.

### HasHasSpecialBid

`func (o *ProductDetailResponseIndicatorsInner) HasHasSpecialBid() bool`

HasHasSpecialBid returns a boolean if a field has been set.

### GetIsExportableToCountry

`func (o *ProductDetailResponseIndicatorsInner) GetIsExportableToCountry() bool`

GetIsExportableToCountry returns the IsExportableToCountry field if non-nil, zero value otherwise.

### GetIsExportableToCountryOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsExportableToCountryOk() (*bool, bool)`

GetIsExportableToCountryOk returns a tuple with the IsExportableToCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExportableToCountry

`func (o *ProductDetailResponseIndicatorsInner) SetIsExportableToCountry(v bool)`

SetIsExportableToCountry sets IsExportableToCountry field to given value.

### HasIsExportableToCountry

`func (o *ProductDetailResponseIndicatorsInner) HasIsExportableToCountry() bool`

HasIsExportableToCountry returns a boolean if a field has been set.

### GetIsDiscontinuedProduct

`func (o *ProductDetailResponseIndicatorsInner) GetIsDiscontinuedProduct() bool`

GetIsDiscontinuedProduct returns the IsDiscontinuedProduct field if non-nil, zero value otherwise.

### GetIsDiscontinuedProductOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsDiscontinuedProductOk() (*bool, bool)`

GetIsDiscontinuedProductOk returns a tuple with the IsDiscontinuedProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDiscontinuedProduct

`func (o *ProductDetailResponseIndicatorsInner) SetIsDiscontinuedProduct(v bool)`

SetIsDiscontinuedProduct sets IsDiscontinuedProduct field to given value.

### HasIsDiscontinuedProduct

`func (o *ProductDetailResponseIndicatorsInner) HasIsDiscontinuedProduct() bool`

HasIsDiscontinuedProduct returns a boolean if a field has been set.

### GetIsRefurbishedProduct

`func (o *ProductDetailResponseIndicatorsInner) GetIsRefurbishedProduct() bool`

GetIsRefurbishedProduct returns the IsRefurbishedProduct field if non-nil, zero value otherwise.

### GetIsRefurbishedProductOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsRefurbishedProductOk() (*bool, bool)`

GetIsRefurbishedProductOk returns a tuple with the IsRefurbishedProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRefurbishedProduct

`func (o *ProductDetailResponseIndicatorsInner) SetIsRefurbishedProduct(v bool)`

SetIsRefurbishedProduct sets IsRefurbishedProduct field to given value.

### HasIsRefurbishedProduct

`func (o *ProductDetailResponseIndicatorsInner) HasIsRefurbishedProduct() bool`

HasIsRefurbishedProduct returns a boolean if a field has been set.

### GetIsReturnableProduct

`func (o *ProductDetailResponseIndicatorsInner) GetIsReturnableProduct() bool`

GetIsReturnableProduct returns the IsReturnableProduct field if non-nil, zero value otherwise.

### GetIsReturnableProductOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsReturnableProductOk() (*bool, bool)`

GetIsReturnableProductOk returns a tuple with the IsReturnableProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReturnableProduct

`func (o *ProductDetailResponseIndicatorsInner) SetIsReturnableProduct(v bool)`

SetIsReturnableProduct sets IsReturnableProduct field to given value.

### HasIsReturnableProduct

`func (o *ProductDetailResponseIndicatorsInner) HasIsReturnableProduct() bool`

HasIsReturnableProduct returns a boolean if a field has been set.

### GetIsIngramShip

`func (o *ProductDetailResponseIndicatorsInner) GetIsIngramShip() bool`

GetIsIngramShip returns the IsIngramShip field if non-nil, zero value otherwise.

### GetIsIngramShipOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsIngramShipOk() (*bool, bool)`

GetIsIngramShipOk returns a tuple with the IsIngramShip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIngramShip

`func (o *ProductDetailResponseIndicatorsInner) SetIsIngramShip(v bool)`

SetIsIngramShip sets IsIngramShip field to given value.

### HasIsIngramShip

`func (o *ProductDetailResponseIndicatorsInner) HasIsIngramShip() bool`

HasIsIngramShip returns a boolean if a field has been set.

### GetIsEnduserRequired

`func (o *ProductDetailResponseIndicatorsInner) GetIsEnduserRequired() bool`

GetIsEnduserRequired returns the IsEnduserRequired field if non-nil, zero value otherwise.

### GetIsEnduserRequiredOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsEnduserRequiredOk() (*bool, bool)`

GetIsEnduserRequiredOk returns a tuple with the IsEnduserRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnduserRequired

`func (o *ProductDetailResponseIndicatorsInner) SetIsEnduserRequired(v bool)`

SetIsEnduserRequired sets IsEnduserRequired field to given value.

### HasIsEnduserRequired

`func (o *ProductDetailResponseIndicatorsInner) HasIsEnduserRequired() bool`

HasIsEnduserRequired returns a boolean if a field has been set.

### GetIsHeavyWeight

`func (o *ProductDetailResponseIndicatorsInner) GetIsHeavyWeight() bool`

GetIsHeavyWeight returns the IsHeavyWeight field if non-nil, zero value otherwise.

### GetIsHeavyWeightOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsHeavyWeightOk() (*bool, bool)`

GetIsHeavyWeightOk returns a tuple with the IsHeavyWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHeavyWeight

`func (o *ProductDetailResponseIndicatorsInner) SetIsHeavyWeight(v bool)`

SetIsHeavyWeight sets IsHeavyWeight field to given value.

### HasIsHeavyWeight

`func (o *ProductDetailResponseIndicatorsInner) HasIsHeavyWeight() bool`

HasIsHeavyWeight returns a boolean if a field has been set.

### GetHasLtl

`func (o *ProductDetailResponseIndicatorsInner) GetHasLtl() bool`

GetHasLtl returns the HasLtl field if non-nil, zero value otherwise.

### GetHasLtlOk

`func (o *ProductDetailResponseIndicatorsInner) GetHasLtlOk() (*bool, bool)`

GetHasLtlOk returns a tuple with the HasLtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLtl

`func (o *ProductDetailResponseIndicatorsInner) SetHasLtl(v bool)`

SetHasLtl sets HasLtl field to given value.

### HasHasLtl

`func (o *ProductDetailResponseIndicatorsInner) HasHasLtl() bool`

HasHasLtl returns a boolean if a field has been set.

### GetIsClearanceProduct

`func (o *ProductDetailResponseIndicatorsInner) GetIsClearanceProduct() bool`

GetIsClearanceProduct returns the IsClearanceProduct field if non-nil, zero value otherwise.

### GetIsClearanceProductOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsClearanceProductOk() (*bool, bool)`

GetIsClearanceProductOk returns a tuple with the IsClearanceProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClearanceProduct

`func (o *ProductDetailResponseIndicatorsInner) SetIsClearanceProduct(v bool)`

SetIsClearanceProduct sets IsClearanceProduct field to given value.

### HasIsClearanceProduct

`func (o *ProductDetailResponseIndicatorsInner) HasIsClearanceProduct() bool`

HasIsClearanceProduct returns a boolean if a field has been set.

### GetHasBundle

`func (o *ProductDetailResponseIndicatorsInner) GetHasBundle() bool`

GetHasBundle returns the HasBundle field if non-nil, zero value otherwise.

### GetHasBundleOk

`func (o *ProductDetailResponseIndicatorsInner) GetHasBundleOk() (*bool, bool)`

GetHasBundleOk returns a tuple with the HasBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBundle

`func (o *ProductDetailResponseIndicatorsInner) SetHasBundle(v bool)`

SetHasBundle sets HasBundle field to given value.

### HasHasBundle

`func (o *ProductDetailResponseIndicatorsInner) HasHasBundle() bool`

HasHasBundle returns a boolean if a field has been set.

### GetIsOversizeProduct

`func (o *ProductDetailResponseIndicatorsInner) GetIsOversizeProduct() bool`

GetIsOversizeProduct returns the IsOversizeProduct field if non-nil, zero value otherwise.

### GetIsOversizeProductOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsOversizeProductOk() (*bool, bool)`

GetIsOversizeProductOk returns a tuple with the IsOversizeProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOversizeProduct

`func (o *ProductDetailResponseIndicatorsInner) SetIsOversizeProduct(v bool)`

SetIsOversizeProduct sets IsOversizeProduct field to given value.

### HasIsOversizeProduct

`func (o *ProductDetailResponseIndicatorsInner) HasIsOversizeProduct() bool`

HasIsOversizeProduct returns a boolean if a field has been set.

### GetIsPreorderProduct

`func (o *ProductDetailResponseIndicatorsInner) GetIsPreorderProduct() bool`

GetIsPreorderProduct returns the IsPreorderProduct field if non-nil, zero value otherwise.

### GetIsPreorderProductOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsPreorderProductOk() (*bool, bool)`

GetIsPreorderProductOk returns a tuple with the IsPreorderProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreorderProduct

`func (o *ProductDetailResponseIndicatorsInner) SetIsPreorderProduct(v bool)`

SetIsPreorderProduct sets IsPreorderProduct field to given value.

### HasIsPreorderProduct

`func (o *ProductDetailResponseIndicatorsInner) HasIsPreorderProduct() bool`

HasIsPreorderProduct returns a boolean if a field has been set.

### GetIsLicenseProduct

`func (o *ProductDetailResponseIndicatorsInner) GetIsLicenseProduct() bool`

GetIsLicenseProduct returns the IsLicenseProduct field if non-nil, zero value otherwise.

### GetIsLicenseProductOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsLicenseProductOk() (*bool, bool)`

GetIsLicenseProductOk returns a tuple with the IsLicenseProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLicenseProduct

`func (o *ProductDetailResponseIndicatorsInner) SetIsLicenseProduct(v bool)`

SetIsLicenseProduct sets IsLicenseProduct field to given value.

### HasIsLicenseProduct

`func (o *ProductDetailResponseIndicatorsInner) HasIsLicenseProduct() bool`

HasIsLicenseProduct returns a boolean if a field has been set.

### GetIsDirectshipOrderable

`func (o *ProductDetailResponseIndicatorsInner) GetIsDirectshipOrderable() bool`

GetIsDirectshipOrderable returns the IsDirectshipOrderable field if non-nil, zero value otherwise.

### GetIsDirectshipOrderableOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsDirectshipOrderableOk() (*bool, bool)`

GetIsDirectshipOrderableOk returns a tuple with the IsDirectshipOrderable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectshipOrderable

`func (o *ProductDetailResponseIndicatorsInner) SetIsDirectshipOrderable(v bool)`

SetIsDirectshipOrderable sets IsDirectshipOrderable field to given value.

### HasIsDirectshipOrderable

`func (o *ProductDetailResponseIndicatorsInner) HasIsDirectshipOrderable() bool`

HasIsDirectshipOrderable returns a boolean if a field has been set.

### GetIsServiceSku

`func (o *ProductDetailResponseIndicatorsInner) GetIsServiceSku() bool`

GetIsServiceSku returns the IsServiceSku field if non-nil, zero value otherwise.

### GetIsServiceSkuOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsServiceSkuOk() (*bool, bool)`

GetIsServiceSkuOk returns a tuple with the IsServiceSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsServiceSku

`func (o *ProductDetailResponseIndicatorsInner) SetIsServiceSku(v bool)`

SetIsServiceSku sets IsServiceSku field to given value.

### HasIsServiceSku

`func (o *ProductDetailResponseIndicatorsInner) HasIsServiceSku() bool`

HasIsServiceSku returns a boolean if a field has been set.

### GetIsConfigurable

`func (o *ProductDetailResponseIndicatorsInner) GetIsConfigurable() bool`

GetIsConfigurable returns the IsConfigurable field if non-nil, zero value otherwise.

### GetIsConfigurableOk

`func (o *ProductDetailResponseIndicatorsInner) GetIsConfigurableOk() (*bool, bool)`

GetIsConfigurableOk returns a tuple with the IsConfigurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigurable

`func (o *ProductDetailResponseIndicatorsInner) SetIsConfigurable(v bool)`

SetIsConfigurable sets IsConfigurable field to given value.

### HasIsConfigurable

`func (o *ProductDetailResponseIndicatorsInner) HasIsConfigurable() bool`

HasIsConfigurable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


