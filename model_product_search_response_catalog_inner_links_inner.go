/*
Reseller API Documentation

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi-sdk-resellers-go

import (
	"encoding/json"
)

// checks if the ProductSearchResponseCatalogInnerLinksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductSearchResponseCatalogInnerLinksInner{}

// ProductSearchResponseCatalogInnerLinksInner HATEOAS links for the price and availability of the sku.
type ProductSearchResponseCatalogInnerLinksInner struct {
	// Provides the details of the product.
	Topic *string `json:"topic,omitempty"`
	// The URL endpoint for accessing the relevant data..
	Href *string `json:"href,omitempty"`
	// The type of call that can be made to the href link(GET)
	Type *string `json:"type,omitempty"`
}

// NewProductSearchResponseCatalogInnerLinksInner instantiates a new ProductSearchResponseCatalogInnerLinksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductSearchResponseCatalogInnerLinksInner() *ProductSearchResponseCatalogInnerLinksInner {
	this := ProductSearchResponseCatalogInnerLinksInner{}
	return &this
}

// NewProductSearchResponseCatalogInnerLinksInnerWithDefaults instantiates a new ProductSearchResponseCatalogInnerLinksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductSearchResponseCatalogInnerLinksInnerWithDefaults() *ProductSearchResponseCatalogInnerLinksInner {
	this := ProductSearchResponseCatalogInnerLinksInner{}
	return &this
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *ProductSearchResponseCatalogInnerLinksInner) GetTopic() string {
	if o == nil || IsNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductSearchResponseCatalogInnerLinksInner) GetTopicOk() (*string, bool) {
	if o == nil || IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *ProductSearchResponseCatalogInnerLinksInner) HasTopic() bool {
	if o != nil && !IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *ProductSearchResponseCatalogInnerLinksInner) SetTopic(v string) {
	o.Topic = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ProductSearchResponseCatalogInnerLinksInner) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductSearchResponseCatalogInnerLinksInner) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ProductSearchResponseCatalogInnerLinksInner) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ProductSearchResponseCatalogInnerLinksInner) SetHref(v string) {
	o.Href = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProductSearchResponseCatalogInnerLinksInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductSearchResponseCatalogInnerLinksInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProductSearchResponseCatalogInnerLinksInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProductSearchResponseCatalogInnerLinksInner) SetType(v string) {
	o.Type = &v
}

func (o ProductSearchResponseCatalogInnerLinksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductSearchResponseCatalogInnerLinksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Topic) {
		toSerialize["topic"] = o.Topic
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableProductSearchResponseCatalogInnerLinksInner struct {
	value *ProductSearchResponseCatalogInnerLinksInner
	isSet bool
}

func (v NullableProductSearchResponseCatalogInnerLinksInner) Get() *ProductSearchResponseCatalogInnerLinksInner {
	return v.value
}

func (v *NullableProductSearchResponseCatalogInnerLinksInner) Set(val *ProductSearchResponseCatalogInnerLinksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProductSearchResponseCatalogInnerLinksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProductSearchResponseCatalogInnerLinksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductSearchResponseCatalogInnerLinksInner(val *ProductSearchResponseCatalogInnerLinksInner) *NullableProductSearchResponseCatalogInnerLinksInner {
	return &NullableProductSearchResponseCatalogInnerLinksInner{value: val, isSet: true}
}

func (v NullableProductSearchResponseCatalogInnerLinksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductSearchResponseCatalogInnerLinksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


