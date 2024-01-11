/*
Reseller API Documentation

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi-sdk-resellers-go

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the InvoiceDetailRequestServicerequestInvoicedetailrequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceDetailRequestServicerequestInvoicedetailrequest{}

// InvoiceDetailRequestServicerequestInvoicedetailrequest struct for InvoiceDetailRequestServicerequestInvoicedetailrequest
type InvoiceDetailRequestServicerequestInvoicedetailrequest struct {
	Invoicenumber string `json:"invoicenumber"`
}

type _InvoiceDetailRequestServicerequestInvoicedetailrequest InvoiceDetailRequestServicerequestInvoicedetailrequest

// NewInvoiceDetailRequestServicerequestInvoicedetailrequest instantiates a new InvoiceDetailRequestServicerequestInvoicedetailrequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceDetailRequestServicerequestInvoicedetailrequest(invoicenumber string) *InvoiceDetailRequestServicerequestInvoicedetailrequest {
	this := InvoiceDetailRequestServicerequestInvoicedetailrequest{}
	this.Invoicenumber = invoicenumber
	return &this
}

// NewInvoiceDetailRequestServicerequestInvoicedetailrequestWithDefaults instantiates a new InvoiceDetailRequestServicerequestInvoicedetailrequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceDetailRequestServicerequestInvoicedetailrequestWithDefaults() *InvoiceDetailRequestServicerequestInvoicedetailrequest {
	this := InvoiceDetailRequestServicerequestInvoicedetailrequest{}
	return &this
}

// GetInvoicenumber returns the Invoicenumber field value
func (o *InvoiceDetailRequestServicerequestInvoicedetailrequest) GetInvoicenumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Invoicenumber
}

// GetInvoicenumberOk returns a tuple with the Invoicenumber field value
// and a boolean to check if the value has been set.
func (o *InvoiceDetailRequestServicerequestInvoicedetailrequest) GetInvoicenumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Invoicenumber, true
}

// SetInvoicenumber sets field value
func (o *InvoiceDetailRequestServicerequestInvoicedetailrequest) SetInvoicenumber(v string) {
	o.Invoicenumber = v
}

func (o InvoiceDetailRequestServicerequestInvoicedetailrequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceDetailRequestServicerequestInvoicedetailrequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoicenumber"] = o.Invoicenumber
	return toSerialize, nil
}

func (o *InvoiceDetailRequestServicerequestInvoicedetailrequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invoicenumber",
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

	varInvoiceDetailRequestServicerequestInvoicedetailrequest := _InvoiceDetailRequestServicerequestInvoicedetailrequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInvoiceDetailRequestServicerequestInvoicedetailrequest)

	if err != nil {
		return err
	}

	*o = InvoiceDetailRequestServicerequestInvoicedetailrequest(varInvoiceDetailRequestServicerequestInvoicedetailrequest)

	return err
}

type NullableInvoiceDetailRequestServicerequestInvoicedetailrequest struct {
	value *InvoiceDetailRequestServicerequestInvoicedetailrequest
	isSet bool
}

func (v NullableInvoiceDetailRequestServicerequestInvoicedetailrequest) Get() *InvoiceDetailRequestServicerequestInvoicedetailrequest {
	return v.value
}

func (v *NullableInvoiceDetailRequestServicerequestInvoicedetailrequest) Set(val *InvoiceDetailRequestServicerequestInvoicedetailrequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceDetailRequestServicerequestInvoicedetailrequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceDetailRequestServicerequestInvoicedetailrequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceDetailRequestServicerequestInvoicedetailrequest(val *InvoiceDetailRequestServicerequestInvoicedetailrequest) *NullableInvoiceDetailRequestServicerequestInvoicedetailrequest {
	return &NullableInvoiceDetailRequestServicerequestInvoicedetailrequest{value: val, isSet: true}
}

func (v NullableInvoiceDetailRequestServicerequestInvoicedetailrequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceDetailRequestServicerequestInvoicedetailrequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


