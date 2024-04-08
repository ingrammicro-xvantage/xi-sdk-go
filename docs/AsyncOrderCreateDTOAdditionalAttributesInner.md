# AsyncOrderCreateDTOAdditionalAttributesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeName** | Pointer to **string** | The attribute name. allowDuplicateCustomerOrderNumber: Allow orders with duplicate customer PO numbers. Enables resellers to have the same PO number for multiple orders. enableCommentsAsLines:  It will enable comments as lines. | [optional] 
**AttributeValue** | Pointer to **string** | The attribute field data. | [optional] 

## Methods

### NewAsyncOrderCreateDTOAdditionalAttributesInner

`func NewAsyncOrderCreateDTOAdditionalAttributesInner() *AsyncOrderCreateDTOAdditionalAttributesInner`

NewAsyncOrderCreateDTOAdditionalAttributesInner instantiates a new AsyncOrderCreateDTOAdditionalAttributesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncOrderCreateDTOAdditionalAttributesInnerWithDefaults

`func NewAsyncOrderCreateDTOAdditionalAttributesInnerWithDefaults() *AsyncOrderCreateDTOAdditionalAttributesInner`

NewAsyncOrderCreateDTOAdditionalAttributesInnerWithDefaults instantiates a new AsyncOrderCreateDTOAdditionalAttributesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeName

`func (o *AsyncOrderCreateDTOAdditionalAttributesInner) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *AsyncOrderCreateDTOAdditionalAttributesInner) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *AsyncOrderCreateDTOAdditionalAttributesInner) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.

### HasAttributeName

`func (o *AsyncOrderCreateDTOAdditionalAttributesInner) HasAttributeName() bool`

HasAttributeName returns a boolean if a field has been set.

### GetAttributeValue

`func (o *AsyncOrderCreateDTOAdditionalAttributesInner) GetAttributeValue() string`

GetAttributeValue returns the AttributeValue field if non-nil, zero value otherwise.

### GetAttributeValueOk

`func (o *AsyncOrderCreateDTOAdditionalAttributesInner) GetAttributeValueOk() (*string, bool)`

GetAttributeValueOk returns a tuple with the AttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValue

`func (o *AsyncOrderCreateDTOAdditionalAttributesInner) SetAttributeValue(v string)`

SetAttributeValue sets AttributeValue field to given value.

### HasAttributeValue

`func (o *AsyncOrderCreateDTOAdditionalAttributesInner) HasAttributeValue() bool`

HasAttributeValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


