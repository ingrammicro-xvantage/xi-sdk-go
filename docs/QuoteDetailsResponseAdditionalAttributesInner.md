# QuoteDetailsResponseAdditionalAttributesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeName** | Pointer to **string** | estimateId - is the identification number for an estimate provided by Cisco for a quote.  dealId - is the identification number for the specific deal pricing related to a Cisco quote  vendorName - Name of Vendor associated with the quote.  vendorMessage - Vendor Message is associated with primary vendor in the quote.  In cases where a vendor requires a message be presented in the quote, the vendor name and message will be retreived and must be included in the quote vendor message fields. | [optional] 
**AttributeValue** | Pointer to **string** | The attribute field data. | [optional] 

## Methods

### NewQuoteDetailsResponseAdditionalAttributesInner

`func NewQuoteDetailsResponseAdditionalAttributesInner() *QuoteDetailsResponseAdditionalAttributesInner`

NewQuoteDetailsResponseAdditionalAttributesInner instantiates a new QuoteDetailsResponseAdditionalAttributesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteDetailsResponseAdditionalAttributesInnerWithDefaults

`func NewQuoteDetailsResponseAdditionalAttributesInnerWithDefaults() *QuoteDetailsResponseAdditionalAttributesInner`

NewQuoteDetailsResponseAdditionalAttributesInnerWithDefaults instantiates a new QuoteDetailsResponseAdditionalAttributesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeName

`func (o *QuoteDetailsResponseAdditionalAttributesInner) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *QuoteDetailsResponseAdditionalAttributesInner) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *QuoteDetailsResponseAdditionalAttributesInner) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.

### HasAttributeName

`func (o *QuoteDetailsResponseAdditionalAttributesInner) HasAttributeName() bool`

HasAttributeName returns a boolean if a field has been set.

### GetAttributeValue

`func (o *QuoteDetailsResponseAdditionalAttributesInner) GetAttributeValue() string`

GetAttributeValue returns the AttributeValue field if non-nil, zero value otherwise.

### GetAttributeValueOk

`func (o *QuoteDetailsResponseAdditionalAttributesInner) GetAttributeValueOk() (*string, bool)`

GetAttributeValueOk returns a tuple with the AttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValue

`func (o *QuoteDetailsResponseAdditionalAttributesInner) SetAttributeValue(v string)`

SetAttributeValue sets AttributeValue field to given value.

### HasAttributeValue

`func (o *QuoteDetailsResponseAdditionalAttributesInner) HasAttributeValue() bool`

HasAttributeValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


