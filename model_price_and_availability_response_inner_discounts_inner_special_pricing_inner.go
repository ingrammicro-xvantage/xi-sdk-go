/*
Reseller API

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi.sdk.resellers

import (
	"encoding/json"
)

// checks if the PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner{}

// PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner struct for PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner
type PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner struct {
	// The type of discount being given to the customer.
	DiscountType *string `json:"discountType,omitempty"`
	// Pre-approved special pricing/bid number provided to the reseller by the vendor for special pricing and discounts. Used to track the bid number where different line items have different bid numbers. Line-level bid numbers take precedence over header-level bid numbers.
	SpecialBidNumer *string `json:"specialBidNumer,omitempty"`
	// Special pricing discount amount given to the customer.
	SpecialPricingDiscount *float32 `json:"specialPricingDiscount,omitempty"`
	// The effective date of the special pricing available to the customer.
	SpecialPricingEffectiveDate *string `json:"specialPricingEffectiveDate,omitempty"`
	// The expiration date of the special pricing available to the customer.
	SpecialPricingExpirationDate *string `json:"specialPricingExpirationDate,omitempty"`
	// The available quantity of products with discounts.
	SpecialPricingAvailableQuantity *int32 `json:"specialPricingAvailableQuantity,omitempty"`
	// The minimum quantity of products that have to be purchased to ensure the discount is applied.
	SpecialPricingMinQuantity *int32 `json:"specialPricingMinQuantity,omitempty"`
	// Type of Government Discount. *Currently, this discount is only available in the USA.
	GovernmentDiscountType *string `json:"governmentDiscountType,omitempty"`
	// Government Discounted Customer Price. *Currently, this discount is only available in the USA.
	GovernmentDiscountedCustomerPrice *float32 `json:"governmentDiscountedCustomerPrice,omitempty"`
}

// NewPriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner instantiates a new PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner() *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner {
	this := PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner{}
	return &this
}

// NewPriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInnerWithDefaults instantiates a new PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInnerWithDefaults() *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner {
	this := PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner{}
	return &this
}

// GetDiscountType returns the DiscountType field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetDiscountType() string {
	if o == nil || IsNil(o.DiscountType) {
		var ret string
		return ret
	}
	return *o.DiscountType
}

// GetDiscountTypeOk returns a tuple with the DiscountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetDiscountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountType) {
		return nil, false
	}
	return o.DiscountType, true
}

// HasDiscountType returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) HasDiscountType() bool {
	if o != nil && !IsNil(o.DiscountType) {
		return true
	}

	return false
}

// SetDiscountType gets a reference to the given string and assigns it to the DiscountType field.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetDiscountType(v string) {
	o.DiscountType = &v
}

// GetSpecialBidNumer returns the SpecialBidNumer field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialBidNumer() string {
	if o == nil || IsNil(o.SpecialBidNumer) {
		var ret string
		return ret
	}
	return *o.SpecialBidNumer
}

// GetSpecialBidNumerOk returns a tuple with the SpecialBidNumer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialBidNumerOk() (*string, bool) {
	if o == nil || IsNil(o.SpecialBidNumer) {
		return nil, false
	}
	return o.SpecialBidNumer, true
}

// HasSpecialBidNumer returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) HasSpecialBidNumer() bool {
	if o != nil && !IsNil(o.SpecialBidNumer) {
		return true
	}

	return false
}

// SetSpecialBidNumer gets a reference to the given string and assigns it to the SpecialBidNumer field.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetSpecialBidNumer(v string) {
	o.SpecialBidNumer = &v
}

// GetSpecialPricingDiscount returns the SpecialPricingDiscount field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialPricingDiscount() float32 {
	if o == nil || IsNil(o.SpecialPricingDiscount) {
		var ret float32
		return ret
	}
	return *o.SpecialPricingDiscount
}

// GetSpecialPricingDiscountOk returns a tuple with the SpecialPricingDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialPricingDiscountOk() (*float32, bool) {
	if o == nil || IsNil(o.SpecialPricingDiscount) {
		return nil, false
	}
	return o.SpecialPricingDiscount, true
}

// HasSpecialPricingDiscount returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) HasSpecialPricingDiscount() bool {
	if o != nil && !IsNil(o.SpecialPricingDiscount) {
		return true
	}

	return false
}

// SetSpecialPricingDiscount gets a reference to the given float32 and assigns it to the SpecialPricingDiscount field.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetSpecialPricingDiscount(v float32) {
	o.SpecialPricingDiscount = &v
}

// GetSpecialPricingEffectiveDate returns the SpecialPricingEffectiveDate field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialPricingEffectiveDate() string {
	if o == nil || IsNil(o.SpecialPricingEffectiveDate) {
		var ret string
		return ret
	}
	return *o.SpecialPricingEffectiveDate
}

// GetSpecialPricingEffectiveDateOk returns a tuple with the SpecialPricingEffectiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialPricingEffectiveDateOk() (*string, bool) {
	if o == nil || IsNil(o.SpecialPricingEffectiveDate) {
		return nil, false
	}
	return o.SpecialPricingEffectiveDate, true
}

// HasSpecialPricingEffectiveDate returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) HasSpecialPricingEffectiveDate() bool {
	if o != nil && !IsNil(o.SpecialPricingEffectiveDate) {
		return true
	}

	return false
}

// SetSpecialPricingEffectiveDate gets a reference to the given string and assigns it to the SpecialPricingEffectiveDate field.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetSpecialPricingEffectiveDate(v string) {
	o.SpecialPricingEffectiveDate = &v
}

// GetSpecialPricingExpirationDate returns the SpecialPricingExpirationDate field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialPricingExpirationDate() string {
	if o == nil || IsNil(o.SpecialPricingExpirationDate) {
		var ret string
		return ret
	}
	return *o.SpecialPricingExpirationDate
}

// GetSpecialPricingExpirationDateOk returns a tuple with the SpecialPricingExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialPricingExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.SpecialPricingExpirationDate) {
		return nil, false
	}
	return o.SpecialPricingExpirationDate, true
}

// HasSpecialPricingExpirationDate returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) HasSpecialPricingExpirationDate() bool {
	if o != nil && !IsNil(o.SpecialPricingExpirationDate) {
		return true
	}

	return false
}

// SetSpecialPricingExpirationDate gets a reference to the given string and assigns it to the SpecialPricingExpirationDate field.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetSpecialPricingExpirationDate(v string) {
	o.SpecialPricingExpirationDate = &v
}

// GetSpecialPricingAvailableQuantity returns the SpecialPricingAvailableQuantity field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialPricingAvailableQuantity() int32 {
	if o == nil || IsNil(o.SpecialPricingAvailableQuantity) {
		var ret int32
		return ret
	}
	return *o.SpecialPricingAvailableQuantity
}

// GetSpecialPricingAvailableQuantityOk returns a tuple with the SpecialPricingAvailableQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialPricingAvailableQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.SpecialPricingAvailableQuantity) {
		return nil, false
	}
	return o.SpecialPricingAvailableQuantity, true
}

// HasSpecialPricingAvailableQuantity returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) HasSpecialPricingAvailableQuantity() bool {
	if o != nil && !IsNil(o.SpecialPricingAvailableQuantity) {
		return true
	}

	return false
}

// SetSpecialPricingAvailableQuantity gets a reference to the given int32 and assigns it to the SpecialPricingAvailableQuantity field.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetSpecialPricingAvailableQuantity(v int32) {
	o.SpecialPricingAvailableQuantity = &v
}

// GetSpecialPricingMinQuantity returns the SpecialPricingMinQuantity field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialPricingMinQuantity() int32 {
	if o == nil || IsNil(o.SpecialPricingMinQuantity) {
		var ret int32
		return ret
	}
	return *o.SpecialPricingMinQuantity
}

// GetSpecialPricingMinQuantityOk returns a tuple with the SpecialPricingMinQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetSpecialPricingMinQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.SpecialPricingMinQuantity) {
		return nil, false
	}
	return o.SpecialPricingMinQuantity, true
}

// HasSpecialPricingMinQuantity returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) HasSpecialPricingMinQuantity() bool {
	if o != nil && !IsNil(o.SpecialPricingMinQuantity) {
		return true
	}

	return false
}

// SetSpecialPricingMinQuantity gets a reference to the given int32 and assigns it to the SpecialPricingMinQuantity field.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetSpecialPricingMinQuantity(v int32) {
	o.SpecialPricingMinQuantity = &v
}

// GetGovernmentDiscountType returns the GovernmentDiscountType field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetGovernmentDiscountType() string {
	if o == nil || IsNil(o.GovernmentDiscountType) {
		var ret string
		return ret
	}
	return *o.GovernmentDiscountType
}

// GetGovernmentDiscountTypeOk returns a tuple with the GovernmentDiscountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetGovernmentDiscountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GovernmentDiscountType) {
		return nil, false
	}
	return o.GovernmentDiscountType, true
}

// HasGovernmentDiscountType returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) HasGovernmentDiscountType() bool {
	if o != nil && !IsNil(o.GovernmentDiscountType) {
		return true
	}

	return false
}

// SetGovernmentDiscountType gets a reference to the given string and assigns it to the GovernmentDiscountType field.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetGovernmentDiscountType(v string) {
	o.GovernmentDiscountType = &v
}

// GetGovernmentDiscountedCustomerPrice returns the GovernmentDiscountedCustomerPrice field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetGovernmentDiscountedCustomerPrice() float32 {
	if o == nil || IsNil(o.GovernmentDiscountedCustomerPrice) {
		var ret float32
		return ret
	}
	return *o.GovernmentDiscountedCustomerPrice
}

// GetGovernmentDiscountedCustomerPriceOk returns a tuple with the GovernmentDiscountedCustomerPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) GetGovernmentDiscountedCustomerPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.GovernmentDiscountedCustomerPrice) {
		return nil, false
	}
	return o.GovernmentDiscountedCustomerPrice, true
}

// HasGovernmentDiscountedCustomerPrice returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) HasGovernmentDiscountedCustomerPrice() bool {
	if o != nil && !IsNil(o.GovernmentDiscountedCustomerPrice) {
		return true
	}

	return false
}

// SetGovernmentDiscountedCustomerPrice gets a reference to the given float32 and assigns it to the GovernmentDiscountedCustomerPrice field.
func (o *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) SetGovernmentDiscountedCustomerPrice(v float32) {
	o.GovernmentDiscountedCustomerPrice = &v
}

func (o PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DiscountType) {
		toSerialize["discountType"] = o.DiscountType
	}
	if !IsNil(o.SpecialBidNumer) {
		toSerialize["specialBidNumer"] = o.SpecialBidNumer
	}
	if !IsNil(o.SpecialPricingDiscount) {
		toSerialize["specialPricingDiscount"] = o.SpecialPricingDiscount
	}
	if !IsNil(o.SpecialPricingEffectiveDate) {
		toSerialize["specialPricingEffectiveDate"] = o.SpecialPricingEffectiveDate
	}
	if !IsNil(o.SpecialPricingExpirationDate) {
		toSerialize["specialPricingExpirationDate"] = o.SpecialPricingExpirationDate
	}
	if !IsNil(o.SpecialPricingAvailableQuantity) {
		toSerialize["specialPricingAvailableQuantity"] = o.SpecialPricingAvailableQuantity
	}
	if !IsNil(o.SpecialPricingMinQuantity) {
		toSerialize["specialPricingMinQuantity"] = o.SpecialPricingMinQuantity
	}
	if !IsNil(o.GovernmentDiscountType) {
		toSerialize["governmentDiscountType"] = o.GovernmentDiscountType
	}
	if !IsNil(o.GovernmentDiscountedCustomerPrice) {
		toSerialize["governmentDiscountedCustomerPrice"] = o.GovernmentDiscountedCustomerPrice
	}
	return toSerialize, nil
}

type NullablePriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner struct {
	value *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner
	isSet bool
}

func (v NullablePriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) Get() *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner {
	return v.value
}

func (v *NullablePriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) Set(val *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner(val *PriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) *NullablePriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner {
	return &NullablePriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner{value: val, isSet: true}
}

func (v NullablePriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceAndAvailabilityResponseInnerDiscountsInnerSpecialPricingInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


