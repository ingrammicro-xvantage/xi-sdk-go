/*
Reseller API

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi.sdk.resellers

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner{}

// OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner struct for OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner
type OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner struct {
	// Values are “P” for product or “C” for comments. This can be left blank when ordering product and a “P” will be assumed.  If you are adding a COMMENT, then this value must be “C”.  Extended spec for comments:   Attribute Name: “commenttext” Attribute Value: “thank you for the order”  To make the comment invisible to the packing slip place “///” in front of the comment in the Attribute Value field.  This will allow the Ingram sales rep to see the comment on the order but will not forward on to shipping documents.
	Linetype *string `json:"linetype,omitempty"`
	// This is used when a partner wants to use their own line number. Can be left blank.
	Linenumber *string `json:"linenumber,omitempty"`
	// This is the Ingram sku number to be used for placing an order.
	Ingrampartnumber *string `json:"ingrampartnumber,omitempty"`
	// The quantity that is to be ordered.
	Quantity string `json:"quantity"`
	// The Manufacturer part number. Can be used to place an order instead of the Ingram sku.  If there are multiple Ingram part numbers to one vendor part number.  The order will be rejected.
	Vendorpartnumber *string `json:"vendorpartnumber,omitempty"`
	// This is the Customers unique part numbers that must be crossed referenced to the Ingram Micro Sku before it can be used.  Please contact your sales rep for additional information on how to set this up.
	Customerpartnumber *string `json:"customerpartnumber,omitempty"`
	UPCCode *string `json:"UPCCode,omitempty"`
	Warehouseid *string `json:"warehouseid,omitempty"`
	// This is a requested price from the customer. Pre-approval is necessary before using this feature.  A methodology called price variance to manage requested pricing needs to be setup in advance by your sales rep.  If unit price is provided without this advanced setup the unit price will be ignored and standard Ingram Micro pricing will apply.
	Unitprice *string `json:"unitprice,omitempty"`
	Enduser *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerEnduser `json:"enduser,omitempty"`
	Productextendedspecs []OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner `json:"productextendedspecs,omitempty"`
}

type _OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner

// NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner instantiates a new OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner(quantity string) *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner {
	this := OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner{}
	this.Quantity = quantity
	return &this
}

// NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerWithDefaults instantiates a new OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerWithDefaults() *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner {
	this := OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner{}
	return &this
}

// GetLinetype returns the Linetype field value if set, zero value otherwise.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetLinetype() string {
	if o == nil || IsNil(o.Linetype) {
		var ret string
		return ret
	}
	return *o.Linetype
}

// GetLinetypeOk returns a tuple with the Linetype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetLinetypeOk() (*string, bool) {
	if o == nil || IsNil(o.Linetype) {
		return nil, false
	}
	return o.Linetype, true
}

// HasLinetype returns a boolean if a field has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) HasLinetype() bool {
	if o != nil && !IsNil(o.Linetype) {
		return true
	}

	return false
}

// SetLinetype gets a reference to the given string and assigns it to the Linetype field.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetLinetype(v string) {
	o.Linetype = &v
}

// GetLinenumber returns the Linenumber field value if set, zero value otherwise.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetLinenumber() string {
	if o == nil || IsNil(o.Linenumber) {
		var ret string
		return ret
	}
	return *o.Linenumber
}

// GetLinenumberOk returns a tuple with the Linenumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetLinenumberOk() (*string, bool) {
	if o == nil || IsNil(o.Linenumber) {
		return nil, false
	}
	return o.Linenumber, true
}

// HasLinenumber returns a boolean if a field has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) HasLinenumber() bool {
	if o != nil && !IsNil(o.Linenumber) {
		return true
	}

	return false
}

// SetLinenumber gets a reference to the given string and assigns it to the Linenumber field.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetLinenumber(v string) {
	o.Linenumber = &v
}

// GetIngrampartnumber returns the Ingrampartnumber field value if set, zero value otherwise.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetIngrampartnumber() string {
	if o == nil || IsNil(o.Ingrampartnumber) {
		var ret string
		return ret
	}
	return *o.Ingrampartnumber
}

// GetIngrampartnumberOk returns a tuple with the Ingrampartnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetIngrampartnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Ingrampartnumber) {
		return nil, false
	}
	return o.Ingrampartnumber, true
}

// HasIngrampartnumber returns a boolean if a field has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) HasIngrampartnumber() bool {
	if o != nil && !IsNil(o.Ingrampartnumber) {
		return true
	}

	return false
}

// SetIngrampartnumber gets a reference to the given string and assigns it to the Ingrampartnumber field.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetIngrampartnumber(v string) {
	o.Ingrampartnumber = &v
}

// GetQuantity returns the Quantity field value
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetQuantity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetQuantityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetQuantity(v string) {
	o.Quantity = v
}

// GetVendorpartnumber returns the Vendorpartnumber field value if set, zero value otherwise.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetVendorpartnumber() string {
	if o == nil || IsNil(o.Vendorpartnumber) {
		var ret string
		return ret
	}
	return *o.Vendorpartnumber
}

// GetVendorpartnumberOk returns a tuple with the Vendorpartnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetVendorpartnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Vendorpartnumber) {
		return nil, false
	}
	return o.Vendorpartnumber, true
}

// HasVendorpartnumber returns a boolean if a field has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) HasVendorpartnumber() bool {
	if o != nil && !IsNil(o.Vendorpartnumber) {
		return true
	}

	return false
}

// SetVendorpartnumber gets a reference to the given string and assigns it to the Vendorpartnumber field.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetVendorpartnumber(v string) {
	o.Vendorpartnumber = &v
}

// GetCustomerpartnumber returns the Customerpartnumber field value if set, zero value otherwise.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetCustomerpartnumber() string {
	if o == nil || IsNil(o.Customerpartnumber) {
		var ret string
		return ret
	}
	return *o.Customerpartnumber
}

// GetCustomerpartnumberOk returns a tuple with the Customerpartnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetCustomerpartnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Customerpartnumber) {
		return nil, false
	}
	return o.Customerpartnumber, true
}

// HasCustomerpartnumber returns a boolean if a field has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) HasCustomerpartnumber() bool {
	if o != nil && !IsNil(o.Customerpartnumber) {
		return true
	}

	return false
}

// SetCustomerpartnumber gets a reference to the given string and assigns it to the Customerpartnumber field.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetCustomerpartnumber(v string) {
	o.Customerpartnumber = &v
}

// GetUPCCode returns the UPCCode field value if set, zero value otherwise.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetUPCCode() string {
	if o == nil || IsNil(o.UPCCode) {
		var ret string
		return ret
	}
	return *o.UPCCode
}

// GetUPCCodeOk returns a tuple with the UPCCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetUPCCodeOk() (*string, bool) {
	if o == nil || IsNil(o.UPCCode) {
		return nil, false
	}
	return o.UPCCode, true
}

// HasUPCCode returns a boolean if a field has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) HasUPCCode() bool {
	if o != nil && !IsNil(o.UPCCode) {
		return true
	}

	return false
}

// SetUPCCode gets a reference to the given string and assigns it to the UPCCode field.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetUPCCode(v string) {
	o.UPCCode = &v
}

// GetWarehouseid returns the Warehouseid field value if set, zero value otherwise.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetWarehouseid() string {
	if o == nil || IsNil(o.Warehouseid) {
		var ret string
		return ret
	}
	return *o.Warehouseid
}

// GetWarehouseidOk returns a tuple with the Warehouseid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetWarehouseidOk() (*string, bool) {
	if o == nil || IsNil(o.Warehouseid) {
		return nil, false
	}
	return o.Warehouseid, true
}

// HasWarehouseid returns a boolean if a field has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) HasWarehouseid() bool {
	if o != nil && !IsNil(o.Warehouseid) {
		return true
	}

	return false
}

// SetWarehouseid gets a reference to the given string and assigns it to the Warehouseid field.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetWarehouseid(v string) {
	o.Warehouseid = &v
}

// GetUnitprice returns the Unitprice field value if set, zero value otherwise.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetUnitprice() string {
	if o == nil || IsNil(o.Unitprice) {
		var ret string
		return ret
	}
	return *o.Unitprice
}

// GetUnitpriceOk returns a tuple with the Unitprice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetUnitpriceOk() (*string, bool) {
	if o == nil || IsNil(o.Unitprice) {
		return nil, false
	}
	return o.Unitprice, true
}

// HasUnitprice returns a boolean if a field has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) HasUnitprice() bool {
	if o != nil && !IsNil(o.Unitprice) {
		return true
	}

	return false
}

// SetUnitprice gets a reference to the given string and assigns it to the Unitprice field.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetUnitprice(v string) {
	o.Unitprice = &v
}

// GetEnduser returns the Enduser field value if set, zero value otherwise.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetEnduser() OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerEnduser {
	if o == nil || IsNil(o.Enduser) {
		var ret OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerEnduser
		return ret
	}
	return *o.Enduser
}

// GetEnduserOk returns a tuple with the Enduser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetEnduserOk() (*OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerEnduser, bool) {
	if o == nil || IsNil(o.Enduser) {
		return nil, false
	}
	return o.Enduser, true
}

// HasEnduser returns a boolean if a field has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) HasEnduser() bool {
	if o != nil && !IsNil(o.Enduser) {
		return true
	}

	return false
}

// SetEnduser gets a reference to the given OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerEnduser and assigns it to the Enduser field.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetEnduser(v OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerEnduser) {
	o.Enduser = &v
}

// GetProductextendedspecs returns the Productextendedspecs field value if set, zero value otherwise.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetProductextendedspecs() []OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner {
	if o == nil || IsNil(o.Productextendedspecs) {
		var ret []OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner
		return ret
	}
	return o.Productextendedspecs
}

// GetProductextendedspecsOk returns a tuple with the Productextendedspecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetProductextendedspecsOk() ([]OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner, bool) {
	if o == nil || IsNil(o.Productextendedspecs) {
		return nil, false
	}
	return o.Productextendedspecs, true
}

// HasProductextendedspecs returns a boolean if a field has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) HasProductextendedspecs() bool {
	if o != nil && !IsNil(o.Productextendedspecs) {
		return true
	}

	return false
}

// SetProductextendedspecs gets a reference to the given []OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner and assigns it to the Productextendedspecs field.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetProductextendedspecs(v []OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner) {
	o.Productextendedspecs = v
}

func (o OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Linetype) {
		toSerialize["linetype"] = o.Linetype
	}
	if !IsNil(o.Linenumber) {
		toSerialize["linenumber"] = o.Linenumber
	}
	if !IsNil(o.Ingrampartnumber) {
		toSerialize["ingrampartnumber"] = o.Ingrampartnumber
	}
	toSerialize["quantity"] = o.Quantity
	if !IsNil(o.Vendorpartnumber) {
		toSerialize["vendorpartnumber"] = o.Vendorpartnumber
	}
	if !IsNil(o.Customerpartnumber) {
		toSerialize["customerpartnumber"] = o.Customerpartnumber
	}
	if !IsNil(o.UPCCode) {
		toSerialize["UPCCode"] = o.UPCCode
	}
	if !IsNil(o.Warehouseid) {
		toSerialize["warehouseid"] = o.Warehouseid
	}
	if !IsNil(o.Unitprice) {
		toSerialize["unitprice"] = o.Unitprice
	}
	if !IsNil(o.Enduser) {
		toSerialize["enduser"] = o.Enduser
	}
	if !IsNil(o.Productextendedspecs) {
		toSerialize["productextendedspecs"] = o.Productextendedspecs
	}
	return toSerialize, nil
}

func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"quantity",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner := _OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner)

	if err != nil {
		return err
	}

	*o = OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner(varOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner)

	return err
}

type NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner struct {
	value *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner
	isSet bool
}

func (v NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) Get() *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner {
	return v.value
}

func (v *NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) Set(val *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner(val *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) *NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner {
	return &NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner{value: val, isSet: true}
}

func (v NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

