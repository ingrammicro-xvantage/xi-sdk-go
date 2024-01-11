# OrderCreateResponseOrdersInnerAdditionalAttributesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeName** | Pointer to **string** |  allowPartialOrder: Allow orders with failed lines. (SAP) Depends on backorder settings.   dpasRating: DX rating by Department of Defense is the highest rating by the highest offices and meant to be top priority. DO any other gov offices at the federal level to priotize.   dpasProgramId: Identifies the actual agency that signed off on the DPAS priority.   storageLocation: Determine the location of the product stock in SAP for Marketplaces.  soldTo: To be used in cases when Sold-To is different than Customer ID.  Z101: Used for end customer details such as name, address, phone, etc. This information flows to SAP and is used by warehouse.  euDepId: DEP ID would be the &#39;End User DEP/ABM Organization ID&#39; up to 32 characters and is assigned by Apple.  depOrderNbr: depordernbr is &#39;End User PO to reseller&#39; Can appear in message lines or dedicated end user po#.   | [optional] 
**AttributeValue** | Pointer to **string** | Attribute value | [optional] 

## Methods

### NewOrderCreateResponseOrdersInnerAdditionalAttributesInner

`func NewOrderCreateResponseOrdersInnerAdditionalAttributesInner() *OrderCreateResponseOrdersInnerAdditionalAttributesInner`

NewOrderCreateResponseOrdersInnerAdditionalAttributesInner instantiates a new OrderCreateResponseOrdersInnerAdditionalAttributesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateResponseOrdersInnerAdditionalAttributesInnerWithDefaults

`func NewOrderCreateResponseOrdersInnerAdditionalAttributesInnerWithDefaults() *OrderCreateResponseOrdersInnerAdditionalAttributesInner`

NewOrderCreateResponseOrdersInnerAdditionalAttributesInnerWithDefaults instantiates a new OrderCreateResponseOrdersInnerAdditionalAttributesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeName

`func (o *OrderCreateResponseOrdersInnerAdditionalAttributesInner) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *OrderCreateResponseOrdersInnerAdditionalAttributesInner) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *OrderCreateResponseOrdersInnerAdditionalAttributesInner) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.

### HasAttributeName

`func (o *OrderCreateResponseOrdersInnerAdditionalAttributesInner) HasAttributeName() bool`

HasAttributeName returns a boolean if a field has been set.

### GetAttributeValue

`func (o *OrderCreateResponseOrdersInnerAdditionalAttributesInner) GetAttributeValue() string`

GetAttributeValue returns the AttributeValue field if non-nil, zero value otherwise.

### GetAttributeValueOk

`func (o *OrderCreateResponseOrdersInnerAdditionalAttributesInner) GetAttributeValueOk() (*string, bool)`

GetAttributeValueOk returns a tuple with the AttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValue

`func (o *OrderCreateResponseOrdersInnerAdditionalAttributesInner) SetAttributeValue(v string)`

SetAttributeValue sets AttributeValue field to given value.

### HasAttributeValue

`func (o *OrderCreateResponseOrdersInnerAdditionalAttributesInner) HasAttributeValue() bool`

HasAttributeValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


