/*
XI SDK Resellers

For Resellers. Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OrderDeleteRequestServicerequestOrderDeleteRequestDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDeleteRequestServicerequestOrderDeleteRequestDetails{}

// OrderDeleteRequestServicerequestOrderDeleteRequestDetails struct for OrderDeleteRequestServicerequestOrderDeleteRequestDetails
type OrderDeleteRequestServicerequestOrderDeleteRequestDetails struct {
	// Date order entered
	EntryDate string `json:"entryDate"`
	// Ingram Micro's first 2 numbers of the order number
	OrderBranch string `json:"orderBranch"`
	// Ingram Micro's middle 6 numbers of the order#
	OrderNumber string `json:"orderNumber"`
	RejectionCode *string `json:"rejectionCode,omitempty"`
	// Ingram Micro's suffix number of the order#
	DistributionNumber *string `json:"distributionNumber,omitempty"`
	// Ingram Micro's last number of the order#
	ShipmentNumber *string `json:"shipmentNumber,omitempty"`
	// Ingram ID(not required)
	OperatorID *string `json:"operatorID,omitempty"`
}

type _OrderDeleteRequestServicerequestOrderDeleteRequestDetails OrderDeleteRequestServicerequestOrderDeleteRequestDetails

// NewOrderDeleteRequestServicerequestOrderDeleteRequestDetails instantiates a new OrderDeleteRequestServicerequestOrderDeleteRequestDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDeleteRequestServicerequestOrderDeleteRequestDetails(entryDate string, orderBranch string, orderNumber string) *OrderDeleteRequestServicerequestOrderDeleteRequestDetails {
	this := OrderDeleteRequestServicerequestOrderDeleteRequestDetails{}
	this.EntryDate = entryDate
	this.OrderBranch = orderBranch
	this.OrderNumber = orderNumber
	return &this
}

// NewOrderDeleteRequestServicerequestOrderDeleteRequestDetailsWithDefaults instantiates a new OrderDeleteRequestServicerequestOrderDeleteRequestDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDeleteRequestServicerequestOrderDeleteRequestDetailsWithDefaults() *OrderDeleteRequestServicerequestOrderDeleteRequestDetails {
	this := OrderDeleteRequestServicerequestOrderDeleteRequestDetails{}
	return &this
}

// GetEntryDate returns the EntryDate field value
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetEntryDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryDate
}

// GetEntryDateOk returns a tuple with the EntryDate field value
// and a boolean to check if the value has been set.
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetEntryDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryDate, true
}

// SetEntryDate sets field value
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) SetEntryDate(v string) {
	o.EntryDate = v
}

// GetOrderBranch returns the OrderBranch field value
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetOrderBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderBranch
}

// GetOrderBranchOk returns a tuple with the OrderBranch field value
// and a boolean to check if the value has been set.
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetOrderBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderBranch, true
}

// SetOrderBranch sets field value
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) SetOrderBranch(v string) {
	o.OrderBranch = v
}

// GetOrderNumber returns the OrderNumber field value
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetOrderNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderNumber
}

// GetOrderNumberOk returns a tuple with the OrderNumber field value
// and a boolean to check if the value has been set.
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetOrderNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderNumber, true
}

// SetOrderNumber sets field value
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) SetOrderNumber(v string) {
	o.OrderNumber = v
}

// GetRejectionCode returns the RejectionCode field value if set, zero value otherwise.
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetRejectionCode() string {
	if o == nil || IsNil(o.RejectionCode) {
		var ret string
		return ret
	}
	return *o.RejectionCode
}

// GetRejectionCodeOk returns a tuple with the RejectionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetRejectionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RejectionCode) {
		return nil, false
	}
	return o.RejectionCode, true
}

// HasRejectionCode returns a boolean if a field has been set.
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) HasRejectionCode() bool {
	if o != nil && !IsNil(o.RejectionCode) {
		return true
	}

	return false
}

// SetRejectionCode gets a reference to the given string and assigns it to the RejectionCode field.
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) SetRejectionCode(v string) {
	o.RejectionCode = &v
}

// GetDistributionNumber returns the DistributionNumber field value if set, zero value otherwise.
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetDistributionNumber() string {
	if o == nil || IsNil(o.DistributionNumber) {
		var ret string
		return ret
	}
	return *o.DistributionNumber
}

// GetDistributionNumberOk returns a tuple with the DistributionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetDistributionNumberOk() (*string, bool) {
	if o == nil || IsNil(o.DistributionNumber) {
		return nil, false
	}
	return o.DistributionNumber, true
}

// HasDistributionNumber returns a boolean if a field has been set.
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) HasDistributionNumber() bool {
	if o != nil && !IsNil(o.DistributionNumber) {
		return true
	}

	return false
}

// SetDistributionNumber gets a reference to the given string and assigns it to the DistributionNumber field.
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) SetDistributionNumber(v string) {
	o.DistributionNumber = &v
}

// GetShipmentNumber returns the ShipmentNumber field value if set, zero value otherwise.
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetShipmentNumber() string {
	if o == nil || IsNil(o.ShipmentNumber) {
		var ret string
		return ret
	}
	return *o.ShipmentNumber
}

// GetShipmentNumberOk returns a tuple with the ShipmentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetShipmentNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ShipmentNumber) {
		return nil, false
	}
	return o.ShipmentNumber, true
}

// HasShipmentNumber returns a boolean if a field has been set.
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) HasShipmentNumber() bool {
	if o != nil && !IsNil(o.ShipmentNumber) {
		return true
	}

	return false
}

// SetShipmentNumber gets a reference to the given string and assigns it to the ShipmentNumber field.
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) SetShipmentNumber(v string) {
	o.ShipmentNumber = &v
}

// GetOperatorID returns the OperatorID field value if set, zero value otherwise.
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetOperatorID() string {
	if o == nil || IsNil(o.OperatorID) {
		var ret string
		return ret
	}
	return *o.OperatorID
}

// GetOperatorIDOk returns a tuple with the OperatorID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetOperatorIDOk() (*string, bool) {
	if o == nil || IsNil(o.OperatorID) {
		return nil, false
	}
	return o.OperatorID, true
}

// HasOperatorID returns a boolean if a field has been set.
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) HasOperatorID() bool {
	if o != nil && !IsNil(o.OperatorID) {
		return true
	}

	return false
}

// SetOperatorID gets a reference to the given string and assigns it to the OperatorID field.
func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) SetOperatorID(v string) {
	o.OperatorID = &v
}

func (o OrderDeleteRequestServicerequestOrderDeleteRequestDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDeleteRequestServicerequestOrderDeleteRequestDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entryDate"] = o.EntryDate
	toSerialize["orderBranch"] = o.OrderBranch
	toSerialize["orderNumber"] = o.OrderNumber
	if !IsNil(o.RejectionCode) {
		toSerialize["rejectionCode"] = o.RejectionCode
	}
	if !IsNil(o.DistributionNumber) {
		toSerialize["distributionNumber"] = o.DistributionNumber
	}
	if !IsNil(o.ShipmentNumber) {
		toSerialize["shipmentNumber"] = o.ShipmentNumber
	}
	if !IsNil(o.OperatorID) {
		toSerialize["operatorID"] = o.OperatorID
	}
	return toSerialize, nil
}

func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entryDate",
		"orderBranch",
		"orderNumber",
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

	varOrderDeleteRequestServicerequestOrderDeleteRequestDetails := _OrderDeleteRequestServicerequestOrderDeleteRequestDetails{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderDeleteRequestServicerequestOrderDeleteRequestDetails)

	if err != nil {
		return err
	}

	*o = OrderDeleteRequestServicerequestOrderDeleteRequestDetails(varOrderDeleteRequestServicerequestOrderDeleteRequestDetails)

	return err
}

type NullableOrderDeleteRequestServicerequestOrderDeleteRequestDetails struct {
	value *OrderDeleteRequestServicerequestOrderDeleteRequestDetails
	isSet bool
}

func (v NullableOrderDeleteRequestServicerequestOrderDeleteRequestDetails) Get() *OrderDeleteRequestServicerequestOrderDeleteRequestDetails {
	return v.value
}

func (v *NullableOrderDeleteRequestServicerequestOrderDeleteRequestDetails) Set(val *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDeleteRequestServicerequestOrderDeleteRequestDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDeleteRequestServicerequestOrderDeleteRequestDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDeleteRequestServicerequestOrderDeleteRequestDetails(val *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) *NullableOrderDeleteRequestServicerequestOrderDeleteRequestDetails {
	return &NullableOrderDeleteRequestServicerequestOrderDeleteRequestDetails{value: val, isSet: true}
}

func (v NullableOrderDeleteRequestServicerequestOrderDeleteRequestDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDeleteRequestServicerequestOrderDeleteRequestDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


