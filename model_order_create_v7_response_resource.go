/*
XI Sdk Resellers

For Resellers seeking to innovate with Ingram Micro's API solutions, automate your eCommerce experience with our array of API's and webhooks to craft a seamless journey for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the OrderCreateV7ResponseResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateV7ResponseResource{}

// OrderCreateV7ResponseResource struct for OrderCreateV7ResponseResource
type OrderCreateV7ResponseResource struct {
	// The reseller's unique PO/Order number.
	CustomerOrderNumber *string `json:"customerOrderNumber,omitempty"`
	// Suffix used to identify billing address. Created during onboarding. Resellers are provided with one or more address IDs depending on how many bill to addresses they need for various flooring companies they are using for credit
	BillToAddressId *string `json:"billToAddressId,omitempty"`
	// true for multiple orders
	OrderSplit *bool `json:"orderSplit,omitempty"`
	// true for partial order succesfully placed
	ProcessedPartially *bool `json:"processedPartially,omitempty"`
	// Total of all the orders including taxes and fees.
	PurchaseOrderTotal *float32 `json:"purchaseOrderTotal,omitempty"`
	ShipToInfo *OrderCreateV7ResponseResourceShipToInfo `json:"shipToInfo,omitempty"`
	// Order-level details.
	Orders []OrderCreateV7ResponseResourceOrdersInner `json:"orders,omitempty"`
}

// NewOrderCreateV7ResponseResource instantiates a new OrderCreateV7ResponseResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateV7ResponseResource() *OrderCreateV7ResponseResource {
	this := OrderCreateV7ResponseResource{}
	return &this
}

// NewOrderCreateV7ResponseResourceWithDefaults instantiates a new OrderCreateV7ResponseResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateV7ResponseResourceWithDefaults() *OrderCreateV7ResponseResource {
	this := OrderCreateV7ResponseResource{}
	return &this
}

// GetCustomerOrderNumber returns the CustomerOrderNumber field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResource) GetCustomerOrderNumber() string {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		var ret string
		return ret
	}
	return *o.CustomerOrderNumber
}

// GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResource) GetCustomerOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		return nil, false
	}
	return o.CustomerOrderNumber, true
}

// HasCustomerOrderNumber returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResource) HasCustomerOrderNumber() bool {
	if o != nil && !IsNil(o.CustomerOrderNumber) {
		return true
	}

	return false
}

// SetCustomerOrderNumber gets a reference to the given string and assigns it to the CustomerOrderNumber field.
func (o *OrderCreateV7ResponseResource) SetCustomerOrderNumber(v string) {
	o.CustomerOrderNumber = &v
}

// GetBillToAddressId returns the BillToAddressId field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResource) GetBillToAddressId() string {
	if o == nil || IsNil(o.BillToAddressId) {
		var ret string
		return ret
	}
	return *o.BillToAddressId
}

// GetBillToAddressIdOk returns a tuple with the BillToAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResource) GetBillToAddressIdOk() (*string, bool) {
	if o == nil || IsNil(o.BillToAddressId) {
		return nil, false
	}
	return o.BillToAddressId, true
}

// HasBillToAddressId returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResource) HasBillToAddressId() bool {
	if o != nil && !IsNil(o.BillToAddressId) {
		return true
	}

	return false
}

// SetBillToAddressId gets a reference to the given string and assigns it to the BillToAddressId field.
func (o *OrderCreateV7ResponseResource) SetBillToAddressId(v string) {
	o.BillToAddressId = &v
}

// GetOrderSplit returns the OrderSplit field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResource) GetOrderSplit() bool {
	if o == nil || IsNil(o.OrderSplit) {
		var ret bool
		return ret
	}
	return *o.OrderSplit
}

// GetOrderSplitOk returns a tuple with the OrderSplit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResource) GetOrderSplitOk() (*bool, bool) {
	if o == nil || IsNil(o.OrderSplit) {
		return nil, false
	}
	return o.OrderSplit, true
}

// HasOrderSplit returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResource) HasOrderSplit() bool {
	if o != nil && !IsNil(o.OrderSplit) {
		return true
	}

	return false
}

// SetOrderSplit gets a reference to the given bool and assigns it to the OrderSplit field.
func (o *OrderCreateV7ResponseResource) SetOrderSplit(v bool) {
	o.OrderSplit = &v
}

// GetProcessedPartially returns the ProcessedPartially field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResource) GetProcessedPartially() bool {
	if o == nil || IsNil(o.ProcessedPartially) {
		var ret bool
		return ret
	}
	return *o.ProcessedPartially
}

// GetProcessedPartiallyOk returns a tuple with the ProcessedPartially field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResource) GetProcessedPartiallyOk() (*bool, bool) {
	if o == nil || IsNil(o.ProcessedPartially) {
		return nil, false
	}
	return o.ProcessedPartially, true
}

// HasProcessedPartially returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResource) HasProcessedPartially() bool {
	if o != nil && !IsNil(o.ProcessedPartially) {
		return true
	}

	return false
}

// SetProcessedPartially gets a reference to the given bool and assigns it to the ProcessedPartially field.
func (o *OrderCreateV7ResponseResource) SetProcessedPartially(v bool) {
	o.ProcessedPartially = &v
}

// GetPurchaseOrderTotal returns the PurchaseOrderTotal field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResource) GetPurchaseOrderTotal() float32 {
	if o == nil || IsNil(o.PurchaseOrderTotal) {
		var ret float32
		return ret
	}
	return *o.PurchaseOrderTotal
}

// GetPurchaseOrderTotalOk returns a tuple with the PurchaseOrderTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResource) GetPurchaseOrderTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.PurchaseOrderTotal) {
		return nil, false
	}
	return o.PurchaseOrderTotal, true
}

// HasPurchaseOrderTotal returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResource) HasPurchaseOrderTotal() bool {
	if o != nil && !IsNil(o.PurchaseOrderTotal) {
		return true
	}

	return false
}

// SetPurchaseOrderTotal gets a reference to the given float32 and assigns it to the PurchaseOrderTotal field.
func (o *OrderCreateV7ResponseResource) SetPurchaseOrderTotal(v float32) {
	o.PurchaseOrderTotal = &v
}

// GetShipToInfo returns the ShipToInfo field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResource) GetShipToInfo() OrderCreateV7ResponseResourceShipToInfo {
	if o == nil || IsNil(o.ShipToInfo) {
		var ret OrderCreateV7ResponseResourceShipToInfo
		return ret
	}
	return *o.ShipToInfo
}

// GetShipToInfoOk returns a tuple with the ShipToInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResource) GetShipToInfoOk() (*OrderCreateV7ResponseResourceShipToInfo, bool) {
	if o == nil || IsNil(o.ShipToInfo) {
		return nil, false
	}
	return o.ShipToInfo, true
}

// HasShipToInfo returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResource) HasShipToInfo() bool {
	if o != nil && !IsNil(o.ShipToInfo) {
		return true
	}

	return false
}

// SetShipToInfo gets a reference to the given OrderCreateV7ResponseResourceShipToInfo and assigns it to the ShipToInfo field.
func (o *OrderCreateV7ResponseResource) SetShipToInfo(v OrderCreateV7ResponseResourceShipToInfo) {
	o.ShipToInfo = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResource) GetOrders() []OrderCreateV7ResponseResourceOrdersInner {
	if o == nil || IsNil(o.Orders) {
		var ret []OrderCreateV7ResponseResourceOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResource) GetOrdersOk() ([]OrderCreateV7ResponseResourceOrdersInner, bool) {
	if o == nil || IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResource) HasOrders() bool {
	if o != nil && !IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []OrderCreateV7ResponseResourceOrdersInner and assigns it to the Orders field.
func (o *OrderCreateV7ResponseResource) SetOrders(v []OrderCreateV7ResponseResourceOrdersInner) {
	o.Orders = v
}

func (o OrderCreateV7ResponseResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateV7ResponseResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerOrderNumber) {
		toSerialize["customerOrderNumber"] = o.CustomerOrderNumber
	}
	if !IsNil(o.BillToAddressId) {
		toSerialize["billToAddressId"] = o.BillToAddressId
	}
	if !IsNil(o.OrderSplit) {
		toSerialize["orderSplit"] = o.OrderSplit
	}
	if !IsNil(o.ProcessedPartially) {
		toSerialize["processedPartially"] = o.ProcessedPartially
	}
	if !IsNil(o.PurchaseOrderTotal) {
		toSerialize["purchaseOrderTotal"] = o.PurchaseOrderTotal
	}
	if !IsNil(o.ShipToInfo) {
		toSerialize["shipToInfo"] = o.ShipToInfo
	}
	if !IsNil(o.Orders) {
		toSerialize["orders"] = o.Orders
	}
	return toSerialize, nil
}

type NullableOrderCreateV7ResponseResource struct {
	value *OrderCreateV7ResponseResource
	isSet bool
}

func (v NullableOrderCreateV7ResponseResource) Get() *OrderCreateV7ResponseResource {
	return v.value
}

func (v *NullableOrderCreateV7ResponseResource) Set(val *OrderCreateV7ResponseResource) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateV7ResponseResource) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateV7ResponseResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateV7ResponseResource(val *OrderCreateV7ResponseResource) *NullableOrderCreateV7ResponseResource {
	return &NullableOrderCreateV7ResponseResource{value: val, isSet: true}
}

func (v NullableOrderCreateV7ResponseResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateV7ResponseResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


