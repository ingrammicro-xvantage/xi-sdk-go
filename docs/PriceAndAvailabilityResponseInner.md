# PriceAndAvailabilityResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductStatusCode** | Pointer to **string** | Codes signifying whether the sku is active or not. | [optional] 
**ProductStatusMessage** | Pointer to **string** | Message returned saying whether sku is active. | [optional] 
**IngramPartNumber** | Pointer to **string** | Ingram Micro unique part number for the product. | [optional] 
**VendorPartNumber** | Pointer to **string** | Vendor’s part number for the product. | [optional] 
**ExtendedVendorPartNumber** | Pointer to **string** | Extended Vendor Part Number. *Currently, this feature is not available in these countries (Mexico, Turkey, New Zealand, Colombia, Chile, Brazil, Peru, Western Sahara). | [optional] 
**CustomerPartNumber** | Pointer to **string** | Reseller / end-user’s part number for the product. | [optional] 
**Upc** | Pointer to **string** | The UPC code for the product. Consists of 12 numeric digits that are uniquely assigned to each trade item. | [optional] 
**PartNumberType** | Pointer to **string** | Number type of the part. | [optional] 
**VendorNumber** | Pointer to **string** | Vendor number that identifies the product. | [optional] 
**VendorName** | Pointer to **string** | Vendor name for the order. | [optional] 
**Description** | Pointer to **string** | The description given for the product. | [optional] 
**ProductClass** | Pointer to **string** | Indicates whether the product is directly shipped from the vendor’s warehouse or if the product ships from Ingram Micro’s warehouse. Class Codes are Ingram classifications on how skus are stocked A &#x3D; Product that is stocked usually in all IM warehouses and replenished on a regular basis. B &#x3D; Product that is stocked in limited IM warehouses and replenished on a regular basis C &#x3D; Product that is stocked in fewer IM warehouses and replenished on a regular basis. D &#x3D; Product that Ingram Micro has elected to discontinue. E &#x3D; Product that will be phased out later, according to the vendor. You may not want to replenish this product, but instead sell down what is in stock. F &#x3D; Product that we carry for a specific customer or supplier under a contractual agreement. N &#x3D; New Sku. Classification before first receipt O &#x3D; Discontinued product to be liquidated S&#x3D; Order for Specialized Demand (Order to backorder) X&#x3D; direct ship from Vendor V &#x3D; product that vendor has elected to discontinue. | [optional] 
**Uom** | Pointer to **string** | The description given for the product. | [optional] 
**ProductStatus** | Pointer to **string** | Status that gives whether the product is Active. | [optional] 
**AcceptBackOrder** | Pointer to **bool** | Boolean that indicates if the product accepts backorder. | [optional] 
**ProductAuthorized** | Pointer to **bool** | Boolean that indicates whether a product is authorized. | [optional] 
**ReturnableProduct** | Pointer to **bool** | Boolean that indicates if the product can be returned. | [optional] 
**EndUserInfoRequired** | Pointer to **bool** | Boolean that indicates  if end user information is required. | [optional] 
**GovtSpecialPriceAvailable** | Pointer to **bool** | Boolean that indicates if special pricing is available for the product. | [optional] 
**GovtProgramType** | Pointer to **string** | Program type, “PA” for government orders, “ED” for education order. | [optional] 
**GovtEndUserType** | Pointer to **string** | Type of end user of the program. F &#x3D; Federal, S &#x3D; State, E &#x3D; Local, K &#x3D; K-12 school, and H &#x3D; Higher Education. | [optional] 
**Availability** | Pointer to [**PriceAndAvailabilityResponseInnerAvailability**](PriceAndAvailabilityResponseInnerAvailability.md) |  | [optional] 
**ReserveInventoryDetails** | Pointer to [**[]PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner**](PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner.md) |  | [optional] 
**Pricing** | Pointer to [**PriceAndAvailabilityResponseInnerPricing**](PriceAndAvailabilityResponseInnerPricing.md) |  | [optional] 
**Discounts** | Pointer to [**[]PriceAndAvailabilityResponseInnerDiscountsInner**](PriceAndAvailabilityResponseInnerDiscountsInner.md) |  | [optional] 
**BundlePartIndicator** | Pointer to **bool** | True of false value to indicate whether it’s bundle part. *Currently, this feature is not available in these countries (Mexico, Turkey, New Zealand, Colombia, Chile, Brazil, Peru, Western Sahara). | [optional] 
**ServiceFees** | Pointer to [**[]PriceAndAvailabilityResponseInnerServiceFeesInner**](PriceAndAvailabilityResponseInnerServiceFeesInner.md) | *Currently, this feature is not available in these countries (Mexico, Turkey, New Zealand, Colombia, Chile, Brazil, Peru, Western Sahara). | [optional] 

## Methods

### NewPriceAndAvailabilityResponseInner

`func NewPriceAndAvailabilityResponseInner() *PriceAndAvailabilityResponseInner`

NewPriceAndAvailabilityResponseInner instantiates a new PriceAndAvailabilityResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceAndAvailabilityResponseInnerWithDefaults

`func NewPriceAndAvailabilityResponseInnerWithDefaults() *PriceAndAvailabilityResponseInner`

NewPriceAndAvailabilityResponseInnerWithDefaults instantiates a new PriceAndAvailabilityResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductStatusCode

`func (o *PriceAndAvailabilityResponseInner) GetProductStatusCode() string`

GetProductStatusCode returns the ProductStatusCode field if non-nil, zero value otherwise.

### GetProductStatusCodeOk

`func (o *PriceAndAvailabilityResponseInner) GetProductStatusCodeOk() (*string, bool)`

GetProductStatusCodeOk returns a tuple with the ProductStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductStatusCode

`func (o *PriceAndAvailabilityResponseInner) SetProductStatusCode(v string)`

SetProductStatusCode sets ProductStatusCode field to given value.

### HasProductStatusCode

`func (o *PriceAndAvailabilityResponseInner) HasProductStatusCode() bool`

HasProductStatusCode returns a boolean if a field has been set.

### GetProductStatusMessage

`func (o *PriceAndAvailabilityResponseInner) GetProductStatusMessage() string`

GetProductStatusMessage returns the ProductStatusMessage field if non-nil, zero value otherwise.

### GetProductStatusMessageOk

`func (o *PriceAndAvailabilityResponseInner) GetProductStatusMessageOk() (*string, bool)`

GetProductStatusMessageOk returns a tuple with the ProductStatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductStatusMessage

`func (o *PriceAndAvailabilityResponseInner) SetProductStatusMessage(v string)`

SetProductStatusMessage sets ProductStatusMessage field to given value.

### HasProductStatusMessage

`func (o *PriceAndAvailabilityResponseInner) HasProductStatusMessage() bool`

HasProductStatusMessage returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *PriceAndAvailabilityResponseInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *PriceAndAvailabilityResponseInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *PriceAndAvailabilityResponseInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *PriceAndAvailabilityResponseInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *PriceAndAvailabilityResponseInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *PriceAndAvailabilityResponseInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *PriceAndAvailabilityResponseInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *PriceAndAvailabilityResponseInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetExtendedVendorPartNumber

`func (o *PriceAndAvailabilityResponseInner) GetExtendedVendorPartNumber() string`

GetExtendedVendorPartNumber returns the ExtendedVendorPartNumber field if non-nil, zero value otherwise.

### GetExtendedVendorPartNumberOk

`func (o *PriceAndAvailabilityResponseInner) GetExtendedVendorPartNumberOk() (*string, bool)`

GetExtendedVendorPartNumberOk returns a tuple with the ExtendedVendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedVendorPartNumber

`func (o *PriceAndAvailabilityResponseInner) SetExtendedVendorPartNumber(v string)`

SetExtendedVendorPartNumber sets ExtendedVendorPartNumber field to given value.

### HasExtendedVendorPartNumber

`func (o *PriceAndAvailabilityResponseInner) HasExtendedVendorPartNumber() bool`

HasExtendedVendorPartNumber returns a boolean if a field has been set.

### GetCustomerPartNumber

`func (o *PriceAndAvailabilityResponseInner) GetCustomerPartNumber() string`

GetCustomerPartNumber returns the CustomerPartNumber field if non-nil, zero value otherwise.

### GetCustomerPartNumberOk

`func (o *PriceAndAvailabilityResponseInner) GetCustomerPartNumberOk() (*string, bool)`

GetCustomerPartNumberOk returns a tuple with the CustomerPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPartNumber

`func (o *PriceAndAvailabilityResponseInner) SetCustomerPartNumber(v string)`

SetCustomerPartNumber sets CustomerPartNumber field to given value.

### HasCustomerPartNumber

`func (o *PriceAndAvailabilityResponseInner) HasCustomerPartNumber() bool`

HasCustomerPartNumber returns a boolean if a field has been set.

### GetUpc

`func (o *PriceAndAvailabilityResponseInner) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *PriceAndAvailabilityResponseInner) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *PriceAndAvailabilityResponseInner) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *PriceAndAvailabilityResponseInner) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### GetPartNumberType

`func (o *PriceAndAvailabilityResponseInner) GetPartNumberType() string`

GetPartNumberType returns the PartNumberType field if non-nil, zero value otherwise.

### GetPartNumberTypeOk

`func (o *PriceAndAvailabilityResponseInner) GetPartNumberTypeOk() (*string, bool)`

GetPartNumberTypeOk returns a tuple with the PartNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumberType

`func (o *PriceAndAvailabilityResponseInner) SetPartNumberType(v string)`

SetPartNumberType sets PartNumberType field to given value.

### HasPartNumberType

`func (o *PriceAndAvailabilityResponseInner) HasPartNumberType() bool`

HasPartNumberType returns a boolean if a field has been set.

### GetVendorNumber

`func (o *PriceAndAvailabilityResponseInner) GetVendorNumber() string`

GetVendorNumber returns the VendorNumber field if non-nil, zero value otherwise.

### GetVendorNumberOk

`func (o *PriceAndAvailabilityResponseInner) GetVendorNumberOk() (*string, bool)`

GetVendorNumberOk returns a tuple with the VendorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorNumber

`func (o *PriceAndAvailabilityResponseInner) SetVendorNumber(v string)`

SetVendorNumber sets VendorNumber field to given value.

### HasVendorNumber

`func (o *PriceAndAvailabilityResponseInner) HasVendorNumber() bool`

HasVendorNumber returns a boolean if a field has been set.

### GetVendorName

`func (o *PriceAndAvailabilityResponseInner) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *PriceAndAvailabilityResponseInner) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *PriceAndAvailabilityResponseInner) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *PriceAndAvailabilityResponseInner) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetDescription

`func (o *PriceAndAvailabilityResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PriceAndAvailabilityResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PriceAndAvailabilityResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PriceAndAvailabilityResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProductClass

`func (o *PriceAndAvailabilityResponseInner) GetProductClass() string`

GetProductClass returns the ProductClass field if non-nil, zero value otherwise.

### GetProductClassOk

`func (o *PriceAndAvailabilityResponseInner) GetProductClassOk() (*string, bool)`

GetProductClassOk returns a tuple with the ProductClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductClass

`func (o *PriceAndAvailabilityResponseInner) SetProductClass(v string)`

SetProductClass sets ProductClass field to given value.

### HasProductClass

`func (o *PriceAndAvailabilityResponseInner) HasProductClass() bool`

HasProductClass returns a boolean if a field has been set.

### GetUom

`func (o *PriceAndAvailabilityResponseInner) GetUom() string`

GetUom returns the Uom field if non-nil, zero value otherwise.

### GetUomOk

`func (o *PriceAndAvailabilityResponseInner) GetUomOk() (*string, bool)`

GetUomOk returns a tuple with the Uom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUom

`func (o *PriceAndAvailabilityResponseInner) SetUom(v string)`

SetUom sets Uom field to given value.

### HasUom

`func (o *PriceAndAvailabilityResponseInner) HasUom() bool`

HasUom returns a boolean if a field has been set.

### GetProductStatus

`func (o *PriceAndAvailabilityResponseInner) GetProductStatus() string`

GetProductStatus returns the ProductStatus field if non-nil, zero value otherwise.

### GetProductStatusOk

`func (o *PriceAndAvailabilityResponseInner) GetProductStatusOk() (*string, bool)`

GetProductStatusOk returns a tuple with the ProductStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductStatus

`func (o *PriceAndAvailabilityResponseInner) SetProductStatus(v string)`

SetProductStatus sets ProductStatus field to given value.

### HasProductStatus

`func (o *PriceAndAvailabilityResponseInner) HasProductStatus() bool`

HasProductStatus returns a boolean if a field has been set.

### GetAcceptBackOrder

`func (o *PriceAndAvailabilityResponseInner) GetAcceptBackOrder() bool`

GetAcceptBackOrder returns the AcceptBackOrder field if non-nil, zero value otherwise.

### GetAcceptBackOrderOk

`func (o *PriceAndAvailabilityResponseInner) GetAcceptBackOrderOk() (*bool, bool)`

GetAcceptBackOrderOk returns a tuple with the AcceptBackOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptBackOrder

`func (o *PriceAndAvailabilityResponseInner) SetAcceptBackOrder(v bool)`

SetAcceptBackOrder sets AcceptBackOrder field to given value.

### HasAcceptBackOrder

`func (o *PriceAndAvailabilityResponseInner) HasAcceptBackOrder() bool`

HasAcceptBackOrder returns a boolean if a field has been set.

### GetProductAuthorized

`func (o *PriceAndAvailabilityResponseInner) GetProductAuthorized() bool`

GetProductAuthorized returns the ProductAuthorized field if non-nil, zero value otherwise.

### GetProductAuthorizedOk

`func (o *PriceAndAvailabilityResponseInner) GetProductAuthorizedOk() (*bool, bool)`

GetProductAuthorizedOk returns a tuple with the ProductAuthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAuthorized

`func (o *PriceAndAvailabilityResponseInner) SetProductAuthorized(v bool)`

SetProductAuthorized sets ProductAuthorized field to given value.

### HasProductAuthorized

`func (o *PriceAndAvailabilityResponseInner) HasProductAuthorized() bool`

HasProductAuthorized returns a boolean if a field has been set.

### GetReturnableProduct

`func (o *PriceAndAvailabilityResponseInner) GetReturnableProduct() bool`

GetReturnableProduct returns the ReturnableProduct field if non-nil, zero value otherwise.

### GetReturnableProductOk

`func (o *PriceAndAvailabilityResponseInner) GetReturnableProductOk() (*bool, bool)`

GetReturnableProductOk returns a tuple with the ReturnableProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnableProduct

`func (o *PriceAndAvailabilityResponseInner) SetReturnableProduct(v bool)`

SetReturnableProduct sets ReturnableProduct field to given value.

### HasReturnableProduct

`func (o *PriceAndAvailabilityResponseInner) HasReturnableProduct() bool`

HasReturnableProduct returns a boolean if a field has been set.

### GetEndUserInfoRequired

`func (o *PriceAndAvailabilityResponseInner) GetEndUserInfoRequired() bool`

GetEndUserInfoRequired returns the EndUserInfoRequired field if non-nil, zero value otherwise.

### GetEndUserInfoRequiredOk

`func (o *PriceAndAvailabilityResponseInner) GetEndUserInfoRequiredOk() (*bool, bool)`

GetEndUserInfoRequiredOk returns a tuple with the EndUserInfoRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserInfoRequired

`func (o *PriceAndAvailabilityResponseInner) SetEndUserInfoRequired(v bool)`

SetEndUserInfoRequired sets EndUserInfoRequired field to given value.

### HasEndUserInfoRequired

`func (o *PriceAndAvailabilityResponseInner) HasEndUserInfoRequired() bool`

HasEndUserInfoRequired returns a boolean if a field has been set.

### GetGovtSpecialPriceAvailable

`func (o *PriceAndAvailabilityResponseInner) GetGovtSpecialPriceAvailable() bool`

GetGovtSpecialPriceAvailable returns the GovtSpecialPriceAvailable field if non-nil, zero value otherwise.

### GetGovtSpecialPriceAvailableOk

`func (o *PriceAndAvailabilityResponseInner) GetGovtSpecialPriceAvailableOk() (*bool, bool)`

GetGovtSpecialPriceAvailableOk returns a tuple with the GovtSpecialPriceAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovtSpecialPriceAvailable

`func (o *PriceAndAvailabilityResponseInner) SetGovtSpecialPriceAvailable(v bool)`

SetGovtSpecialPriceAvailable sets GovtSpecialPriceAvailable field to given value.

### HasGovtSpecialPriceAvailable

`func (o *PriceAndAvailabilityResponseInner) HasGovtSpecialPriceAvailable() bool`

HasGovtSpecialPriceAvailable returns a boolean if a field has been set.

### GetGovtProgramType

`func (o *PriceAndAvailabilityResponseInner) GetGovtProgramType() string`

GetGovtProgramType returns the GovtProgramType field if non-nil, zero value otherwise.

### GetGovtProgramTypeOk

`func (o *PriceAndAvailabilityResponseInner) GetGovtProgramTypeOk() (*string, bool)`

GetGovtProgramTypeOk returns a tuple with the GovtProgramType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovtProgramType

`func (o *PriceAndAvailabilityResponseInner) SetGovtProgramType(v string)`

SetGovtProgramType sets GovtProgramType field to given value.

### HasGovtProgramType

`func (o *PriceAndAvailabilityResponseInner) HasGovtProgramType() bool`

HasGovtProgramType returns a boolean if a field has been set.

### GetGovtEndUserType

`func (o *PriceAndAvailabilityResponseInner) GetGovtEndUserType() string`

GetGovtEndUserType returns the GovtEndUserType field if non-nil, zero value otherwise.

### GetGovtEndUserTypeOk

`func (o *PriceAndAvailabilityResponseInner) GetGovtEndUserTypeOk() (*string, bool)`

GetGovtEndUserTypeOk returns a tuple with the GovtEndUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovtEndUserType

`func (o *PriceAndAvailabilityResponseInner) SetGovtEndUserType(v string)`

SetGovtEndUserType sets GovtEndUserType field to given value.

### HasGovtEndUserType

`func (o *PriceAndAvailabilityResponseInner) HasGovtEndUserType() bool`

HasGovtEndUserType returns a boolean if a field has been set.

### GetAvailability

`func (o *PriceAndAvailabilityResponseInner) GetAvailability() PriceAndAvailabilityResponseInnerAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *PriceAndAvailabilityResponseInner) GetAvailabilityOk() (*PriceAndAvailabilityResponseInnerAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *PriceAndAvailabilityResponseInner) SetAvailability(v PriceAndAvailabilityResponseInnerAvailability)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *PriceAndAvailabilityResponseInner) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetReserveInventoryDetails

`func (o *PriceAndAvailabilityResponseInner) GetReserveInventoryDetails() []PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner`

GetReserveInventoryDetails returns the ReserveInventoryDetails field if non-nil, zero value otherwise.

### GetReserveInventoryDetailsOk

`func (o *PriceAndAvailabilityResponseInner) GetReserveInventoryDetailsOk() (*[]PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner, bool)`

GetReserveInventoryDetailsOk returns a tuple with the ReserveInventoryDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveInventoryDetails

`func (o *PriceAndAvailabilityResponseInner) SetReserveInventoryDetails(v []PriceAndAvailabilityResponseInnerReserveInventoryDetailsInner)`

SetReserveInventoryDetails sets ReserveInventoryDetails field to given value.

### HasReserveInventoryDetails

`func (o *PriceAndAvailabilityResponseInner) HasReserveInventoryDetails() bool`

HasReserveInventoryDetails returns a boolean if a field has been set.

### GetPricing

`func (o *PriceAndAvailabilityResponseInner) GetPricing() PriceAndAvailabilityResponseInnerPricing`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *PriceAndAvailabilityResponseInner) GetPricingOk() (*PriceAndAvailabilityResponseInnerPricing, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *PriceAndAvailabilityResponseInner) SetPricing(v PriceAndAvailabilityResponseInnerPricing)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *PriceAndAvailabilityResponseInner) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetDiscounts

`func (o *PriceAndAvailabilityResponseInner) GetDiscounts() []PriceAndAvailabilityResponseInnerDiscountsInner`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *PriceAndAvailabilityResponseInner) GetDiscountsOk() (*[]PriceAndAvailabilityResponseInnerDiscountsInner, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *PriceAndAvailabilityResponseInner) SetDiscounts(v []PriceAndAvailabilityResponseInnerDiscountsInner)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *PriceAndAvailabilityResponseInner) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetBundlePartIndicator

`func (o *PriceAndAvailabilityResponseInner) GetBundlePartIndicator() bool`

GetBundlePartIndicator returns the BundlePartIndicator field if non-nil, zero value otherwise.

### GetBundlePartIndicatorOk

`func (o *PriceAndAvailabilityResponseInner) GetBundlePartIndicatorOk() (*bool, bool)`

GetBundlePartIndicatorOk returns a tuple with the BundlePartIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundlePartIndicator

`func (o *PriceAndAvailabilityResponseInner) SetBundlePartIndicator(v bool)`

SetBundlePartIndicator sets BundlePartIndicator field to given value.

### HasBundlePartIndicator

`func (o *PriceAndAvailabilityResponseInner) HasBundlePartIndicator() bool`

HasBundlePartIndicator returns a boolean if a field has been set.

### GetServiceFees

`func (o *PriceAndAvailabilityResponseInner) GetServiceFees() []PriceAndAvailabilityResponseInnerServiceFeesInner`

GetServiceFees returns the ServiceFees field if non-nil, zero value otherwise.

### GetServiceFeesOk

`func (o *PriceAndAvailabilityResponseInner) GetServiceFeesOk() (*[]PriceAndAvailabilityResponseInnerServiceFeesInner, bool)`

GetServiceFeesOk returns a tuple with the ServiceFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFees

`func (o *PriceAndAvailabilityResponseInner) SetServiceFees(v []PriceAndAvailabilityResponseInnerServiceFeesInner)`

SetServiceFees sets ServiceFees field to given value.

### HasServiceFees

`func (o *PriceAndAvailabilityResponseInner) HasServiceFees() bool`

HasServiceFees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


