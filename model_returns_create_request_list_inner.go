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

// checks if the ReturnsCreateRequestListInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnsCreateRequestListInner{}

// ReturnsCreateRequestListInner struct for ReturnsCreateRequestListInner
type ReturnsCreateRequestListInner struct {
	// The Invoice number of the order.
	InvoiceNumber string `json:"invoiceNumber"`
	// Date of an Invoice.
	InvoiceDate string `json:"invoiceDate"`
	// The reseller's order number for reference in their system.
	CustomerOrderNumber *string `json:"customerOrderNumber,omitempty"`
	// Unique line number from Ingram.
	IngramPartNumber *string `json:"ingramPartNumber,omitempty"`
	// Vendor Part Number.
	VendorPartNumber *string `json:"vendorPartNumber,omitempty"`
	// Serial number of the product.
	SerialNumber *string `json:"serialNumber,omitempty"`
	// Return quantity of the product.
	Quantity int32 `json:"quantity"`
	// Primary reason to return the product.
	PrimaryReason string `json:"primaryReason"`
	// Secondary reason to return the product.
	SecondaryReason string `json:"secondaryReason"`
	// Return notes.
	Notes *string `json:"notes,omitempty"`
	// Reference number to return the product.
	ReferenceNumber *string `json:"referenceNumber,omitempty"`
	// Suffix used to identify billing address.
	BillToAddressId *string `json:"billToAddressId,omitempty"`
	ShipFromInfo []ReturnsCreateRequestListInnerShipFromInfoInner `json:"shipFromInfo"`
	// Number of boxes to return.
	NumberOfBoxes int32 `json:"numberOfBoxes"`
}

type _ReturnsCreateRequestListInner ReturnsCreateRequestListInner

// NewReturnsCreateRequestListInner instantiates a new ReturnsCreateRequestListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnsCreateRequestListInner(invoiceNumber string, invoiceDate string, quantity int32, primaryReason string, secondaryReason string, shipFromInfo []ReturnsCreateRequestListInnerShipFromInfoInner, numberOfBoxes int32) *ReturnsCreateRequestListInner {
	this := ReturnsCreateRequestListInner{}
	this.InvoiceNumber = invoiceNumber
	this.InvoiceDate = invoiceDate
	this.Quantity = quantity
	this.PrimaryReason = primaryReason
	this.SecondaryReason = secondaryReason
	this.ShipFromInfo = shipFromInfo
	this.NumberOfBoxes = numberOfBoxes
	return &this
}

// NewReturnsCreateRequestListInnerWithDefaults instantiates a new ReturnsCreateRequestListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnsCreateRequestListInnerWithDefaults() *ReturnsCreateRequestListInner {
	this := ReturnsCreateRequestListInner{}
	return &this
}

// GetInvoiceNumber returns the InvoiceNumber field value
func (o *ReturnsCreateRequestListInner) GetInvoiceNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceNumber
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInner) GetInvoiceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceNumber, true
}

// SetInvoiceNumber sets field value
func (o *ReturnsCreateRequestListInner) SetInvoiceNumber(v string) {
	o.InvoiceNumber = v
}

// GetInvoiceDate returns the InvoiceDate field value
func (o *ReturnsCreateRequestListInner) GetInvoiceDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceDate
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInner) GetInvoiceDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceDate, true
}

// SetInvoiceDate sets field value
func (o *ReturnsCreateRequestListInner) SetInvoiceDate(v string) {
	o.InvoiceDate = v
}

// GetCustomerOrderNumber returns the CustomerOrderNumber field value if set, zero value otherwise.
func (o *ReturnsCreateRequestListInner) GetCustomerOrderNumber() string {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		var ret string
		return ret
	}
	return *o.CustomerOrderNumber
}

// GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInner) GetCustomerOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		return nil, false
	}
	return o.CustomerOrderNumber, true
}

// HasCustomerOrderNumber returns a boolean if a field has been set.
func (o *ReturnsCreateRequestListInner) HasCustomerOrderNumber() bool {
	if o != nil && !IsNil(o.CustomerOrderNumber) {
		return true
	}

	return false
}

// SetCustomerOrderNumber gets a reference to the given string and assigns it to the CustomerOrderNumber field.
func (o *ReturnsCreateRequestListInner) SetCustomerOrderNumber(v string) {
	o.CustomerOrderNumber = &v
}

// GetIngramPartNumber returns the IngramPartNumber field value if set, zero value otherwise.
func (o *ReturnsCreateRequestListInner) GetIngramPartNumber() string {
	if o == nil || IsNil(o.IngramPartNumber) {
		var ret string
		return ret
	}
	return *o.IngramPartNumber
}

// GetIngramPartNumberOk returns a tuple with the IngramPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInner) GetIngramPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramPartNumber) {
		return nil, false
	}
	return o.IngramPartNumber, true
}

// HasIngramPartNumber returns a boolean if a field has been set.
func (o *ReturnsCreateRequestListInner) HasIngramPartNumber() bool {
	if o != nil && !IsNil(o.IngramPartNumber) {
		return true
	}

	return false
}

// SetIngramPartNumber gets a reference to the given string and assigns it to the IngramPartNumber field.
func (o *ReturnsCreateRequestListInner) SetIngramPartNumber(v string) {
	o.IngramPartNumber = &v
}

// GetVendorPartNumber returns the VendorPartNumber field value if set, zero value otherwise.
func (o *ReturnsCreateRequestListInner) GetVendorPartNumber() string {
	if o == nil || IsNil(o.VendorPartNumber) {
		var ret string
		return ret
	}
	return *o.VendorPartNumber
}

// GetVendorPartNumberOk returns a tuple with the VendorPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInner) GetVendorPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VendorPartNumber) {
		return nil, false
	}
	return o.VendorPartNumber, true
}

// HasVendorPartNumber returns a boolean if a field has been set.
func (o *ReturnsCreateRequestListInner) HasVendorPartNumber() bool {
	if o != nil && !IsNil(o.VendorPartNumber) {
		return true
	}

	return false
}

// SetVendorPartNumber gets a reference to the given string and assigns it to the VendorPartNumber field.
func (o *ReturnsCreateRequestListInner) SetVendorPartNumber(v string) {
	o.VendorPartNumber = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *ReturnsCreateRequestListInner) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInner) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *ReturnsCreateRequestListInner) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *ReturnsCreateRequestListInner) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetQuantity returns the Quantity field value
func (o *ReturnsCreateRequestListInner) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInner) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *ReturnsCreateRequestListInner) SetQuantity(v int32) {
	o.Quantity = v
}

// GetPrimaryReason returns the PrimaryReason field value
func (o *ReturnsCreateRequestListInner) GetPrimaryReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryReason
}

// GetPrimaryReasonOk returns a tuple with the PrimaryReason field value
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInner) GetPrimaryReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryReason, true
}

// SetPrimaryReason sets field value
func (o *ReturnsCreateRequestListInner) SetPrimaryReason(v string) {
	o.PrimaryReason = v
}

// GetSecondaryReason returns the SecondaryReason field value
func (o *ReturnsCreateRequestListInner) GetSecondaryReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecondaryReason
}

// GetSecondaryReasonOk returns a tuple with the SecondaryReason field value
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInner) GetSecondaryReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecondaryReason, true
}

// SetSecondaryReason sets field value
func (o *ReturnsCreateRequestListInner) SetSecondaryReason(v string) {
	o.SecondaryReason = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ReturnsCreateRequestListInner) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInner) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ReturnsCreateRequestListInner) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *ReturnsCreateRequestListInner) SetNotes(v string) {
	o.Notes = &v
}

// GetReferenceNumber returns the ReferenceNumber field value if set, zero value otherwise.
func (o *ReturnsCreateRequestListInner) GetReferenceNumber() string {
	if o == nil || IsNil(o.ReferenceNumber) {
		var ret string
		return ret
	}
	return *o.ReferenceNumber
}

// GetReferenceNumberOk returns a tuple with the ReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInner) GetReferenceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceNumber) {
		return nil, false
	}
	return o.ReferenceNumber, true
}

// HasReferenceNumber returns a boolean if a field has been set.
func (o *ReturnsCreateRequestListInner) HasReferenceNumber() bool {
	if o != nil && !IsNil(o.ReferenceNumber) {
		return true
	}

	return false
}

// SetReferenceNumber gets a reference to the given string and assigns it to the ReferenceNumber field.
func (o *ReturnsCreateRequestListInner) SetReferenceNumber(v string) {
	o.ReferenceNumber = &v
}

// GetBillToAddressId returns the BillToAddressId field value if set, zero value otherwise.
func (o *ReturnsCreateRequestListInner) GetBillToAddressId() string {
	if o == nil || IsNil(o.BillToAddressId) {
		var ret string
		return ret
	}
	return *o.BillToAddressId
}

// GetBillToAddressIdOk returns a tuple with the BillToAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInner) GetBillToAddressIdOk() (*string, bool) {
	if o == nil || IsNil(o.BillToAddressId) {
		return nil, false
	}
	return o.BillToAddressId, true
}

// HasBillToAddressId returns a boolean if a field has been set.
func (o *ReturnsCreateRequestListInner) HasBillToAddressId() bool {
	if o != nil && !IsNil(o.BillToAddressId) {
		return true
	}

	return false
}

// SetBillToAddressId gets a reference to the given string and assigns it to the BillToAddressId field.
func (o *ReturnsCreateRequestListInner) SetBillToAddressId(v string) {
	o.BillToAddressId = &v
}

// GetShipFromInfo returns the ShipFromInfo field value
func (o *ReturnsCreateRequestListInner) GetShipFromInfo() []ReturnsCreateRequestListInnerShipFromInfoInner {
	if o == nil {
		var ret []ReturnsCreateRequestListInnerShipFromInfoInner
		return ret
	}

	return o.ShipFromInfo
}

// GetShipFromInfoOk returns a tuple with the ShipFromInfo field value
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInner) GetShipFromInfoOk() ([]ReturnsCreateRequestListInnerShipFromInfoInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShipFromInfo, true
}

// SetShipFromInfo sets field value
func (o *ReturnsCreateRequestListInner) SetShipFromInfo(v []ReturnsCreateRequestListInnerShipFromInfoInner) {
	o.ShipFromInfo = v
}

// GetNumberOfBoxes returns the NumberOfBoxes field value
func (o *ReturnsCreateRequestListInner) GetNumberOfBoxes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfBoxes
}

// GetNumberOfBoxesOk returns a tuple with the NumberOfBoxes field value
// and a boolean to check if the value has been set.
func (o *ReturnsCreateRequestListInner) GetNumberOfBoxesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfBoxes, true
}

// SetNumberOfBoxes sets field value
func (o *ReturnsCreateRequestListInner) SetNumberOfBoxes(v int32) {
	o.NumberOfBoxes = v
}

func (o ReturnsCreateRequestListInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnsCreateRequestListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoiceNumber"] = o.InvoiceNumber
	toSerialize["invoiceDate"] = o.InvoiceDate
	if !IsNil(o.CustomerOrderNumber) {
		toSerialize["customerOrderNumber"] = o.CustomerOrderNumber
	}
	if !IsNil(o.IngramPartNumber) {
		toSerialize["ingramPartNumber"] = o.IngramPartNumber
	}
	if !IsNil(o.VendorPartNumber) {
		toSerialize["vendorPartNumber"] = o.VendorPartNumber
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	toSerialize["quantity"] = o.Quantity
	toSerialize["primaryReason"] = o.PrimaryReason
	toSerialize["secondaryReason"] = o.SecondaryReason
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.ReferenceNumber) {
		toSerialize["referenceNumber"] = o.ReferenceNumber
	}
	if !IsNil(o.BillToAddressId) {
		toSerialize["billToAddressId"] = o.BillToAddressId
	}
	toSerialize["shipFromInfo"] = o.ShipFromInfo
	toSerialize["numberOfBoxes"] = o.NumberOfBoxes
	return toSerialize, nil
}

func (o *ReturnsCreateRequestListInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invoiceNumber",
		"invoiceDate",
		"quantity",
		"primaryReason",
		"secondaryReason",
		"shipFromInfo",
		"numberOfBoxes",
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

	varReturnsCreateRequestListInner := _ReturnsCreateRequestListInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReturnsCreateRequestListInner)

	if err != nil {
		return err
	}

	*o = ReturnsCreateRequestListInner(varReturnsCreateRequestListInner)

	return err
}

type NullableReturnsCreateRequestListInner struct {
	value *ReturnsCreateRequestListInner
	isSet bool
}

func (v NullableReturnsCreateRequestListInner) Get() *ReturnsCreateRequestListInner {
	return v.value
}

func (v *NullableReturnsCreateRequestListInner) Set(val *ReturnsCreateRequestListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnsCreateRequestListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnsCreateRequestListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnsCreateRequestListInner(val *ReturnsCreateRequestListInner) *NullableReturnsCreateRequestListInner {
	return &NullableReturnsCreateRequestListInner{value: val, isSet: true}
}

func (v NullableReturnsCreateRequestListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnsCreateRequestListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


