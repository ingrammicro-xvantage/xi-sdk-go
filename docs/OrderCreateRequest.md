# OrderCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerOrderNumber** | **string** | The reseller&#39;s unique PO/Order number. | 
**EndCustomerOrderNumber** | Pointer to **string** | The end user/customer&#39;s Purchase Order number. | [optional] 
**BillToAddressId** | Pointer to **string** | Suffix used to identify billing address. Created during onboarding. Resellers are provided with one or more address IDs depending on how many bill to addresses they need for various flooring companies they are using for credit. | [optional] 
**SpecialBidNumber** | Pointer to **string** | The bid number provided to the reseller by the vendor for special pricing and discounts. Line-level bid numbers take precedence over header-level bid numbers. | [optional] 
**Notes** | Pointer to **string** | Order level notes. | [optional] 
**AcceptBackOrder** | Pointer to **bool** | ENUM [\&quot;true\&quot;,\&quot;false\&quot;] - accept order if this item is backordered. This field along with shipComplete field decides the value of backorderflag. The value of this field is ignored when shipComplete field is present. | [optional] 
**ResellerInfo** | Pointer to [**OrderCreateRequestResellerInfo**](OrderCreateRequestResellerInfo.md) |  | [optional] 
**Vmf** | Pointer to [**OrderCreateRequestVmf**](OrderCreateRequestVmf.md) |  | [optional] 
**ShipToInfo** | Pointer to [**OrderCreateRequestShipToInfo**](OrderCreateRequestShipToInfo.md) |  | [optional] 
**EndUserInfo** | Pointer to [**OrderCreateRequestEndUserInfo**](OrderCreateRequestEndUserInfo.md) |  | [optional] 
**Lines** | Pointer to [**[]OrderCreateRequestLinesInner**](OrderCreateRequestLinesInner.md) | The line-level details of the order. | [optional] 
**ShipmentDetails** | Pointer to [**OrderCreateRequestShipmentDetails**](OrderCreateRequestShipmentDetails.md) |  | [optional] 
**AdditionalAttributes** | Pointer to [**[]OrderCreateRequestAdditionalAttributesInner**](OrderCreateRequestAdditionalAttributesInner.md) | Shipment-level additional attributes. | [optional] 

## Methods

### NewOrderCreateRequest

`func NewOrderCreateRequest(customerOrderNumber string, ) *OrderCreateRequest`

NewOrderCreateRequest instantiates a new OrderCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateRequestWithDefaults

`func NewOrderCreateRequestWithDefaults() *OrderCreateRequest`

NewOrderCreateRequestWithDefaults instantiates a new OrderCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerOrderNumber

`func (o *OrderCreateRequest) GetCustomerOrderNumber() string`

GetCustomerOrderNumber returns the CustomerOrderNumber field if non-nil, zero value otherwise.

### GetCustomerOrderNumberOk

`func (o *OrderCreateRequest) GetCustomerOrderNumberOk() (*string, bool)`

GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOrderNumber

`func (o *OrderCreateRequest) SetCustomerOrderNumber(v string)`

SetCustomerOrderNumber sets CustomerOrderNumber field to given value.


### GetEndCustomerOrderNumber

`func (o *OrderCreateRequest) GetEndCustomerOrderNumber() string`

GetEndCustomerOrderNumber returns the EndCustomerOrderNumber field if non-nil, zero value otherwise.

### GetEndCustomerOrderNumberOk

`func (o *OrderCreateRequest) GetEndCustomerOrderNumberOk() (*string, bool)`

GetEndCustomerOrderNumberOk returns a tuple with the EndCustomerOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCustomerOrderNumber

`func (o *OrderCreateRequest) SetEndCustomerOrderNumber(v string)`

SetEndCustomerOrderNumber sets EndCustomerOrderNumber field to given value.

### HasEndCustomerOrderNumber

`func (o *OrderCreateRequest) HasEndCustomerOrderNumber() bool`

HasEndCustomerOrderNumber returns a boolean if a field has been set.

### GetBillToAddressId

`func (o *OrderCreateRequest) GetBillToAddressId() string`

GetBillToAddressId returns the BillToAddressId field if non-nil, zero value otherwise.

### GetBillToAddressIdOk

`func (o *OrderCreateRequest) GetBillToAddressIdOk() (*string, bool)`

GetBillToAddressIdOk returns a tuple with the BillToAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToAddressId

`func (o *OrderCreateRequest) SetBillToAddressId(v string)`

SetBillToAddressId sets BillToAddressId field to given value.

### HasBillToAddressId

`func (o *OrderCreateRequest) HasBillToAddressId() bool`

HasBillToAddressId returns a boolean if a field has been set.

### GetSpecialBidNumber

`func (o *OrderCreateRequest) GetSpecialBidNumber() string`

GetSpecialBidNumber returns the SpecialBidNumber field if non-nil, zero value otherwise.

### GetSpecialBidNumberOk

`func (o *OrderCreateRequest) GetSpecialBidNumberOk() (*string, bool)`

GetSpecialBidNumberOk returns a tuple with the SpecialBidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialBidNumber

`func (o *OrderCreateRequest) SetSpecialBidNumber(v string)`

SetSpecialBidNumber sets SpecialBidNumber field to given value.

### HasSpecialBidNumber

`func (o *OrderCreateRequest) HasSpecialBidNumber() bool`

HasSpecialBidNumber returns a boolean if a field has been set.

### GetNotes

`func (o *OrderCreateRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OrderCreateRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OrderCreateRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OrderCreateRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAcceptBackOrder

`func (o *OrderCreateRequest) GetAcceptBackOrder() bool`

GetAcceptBackOrder returns the AcceptBackOrder field if non-nil, zero value otherwise.

### GetAcceptBackOrderOk

`func (o *OrderCreateRequest) GetAcceptBackOrderOk() (*bool, bool)`

GetAcceptBackOrderOk returns a tuple with the AcceptBackOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptBackOrder

`func (o *OrderCreateRequest) SetAcceptBackOrder(v bool)`

SetAcceptBackOrder sets AcceptBackOrder field to given value.

### HasAcceptBackOrder

`func (o *OrderCreateRequest) HasAcceptBackOrder() bool`

HasAcceptBackOrder returns a boolean if a field has been set.

### GetResellerInfo

`func (o *OrderCreateRequest) GetResellerInfo() OrderCreateRequestResellerInfo`

GetResellerInfo returns the ResellerInfo field if non-nil, zero value otherwise.

### GetResellerInfoOk

`func (o *OrderCreateRequest) GetResellerInfoOk() (*OrderCreateRequestResellerInfo, bool)`

GetResellerInfoOk returns a tuple with the ResellerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerInfo

`func (o *OrderCreateRequest) SetResellerInfo(v OrderCreateRequestResellerInfo)`

SetResellerInfo sets ResellerInfo field to given value.

### HasResellerInfo

`func (o *OrderCreateRequest) HasResellerInfo() bool`

HasResellerInfo returns a boolean if a field has been set.

### GetVmf

`func (o *OrderCreateRequest) GetVmf() OrderCreateRequestVmf`

GetVmf returns the Vmf field if non-nil, zero value otherwise.

### GetVmfOk

`func (o *OrderCreateRequest) GetVmfOk() (*OrderCreateRequestVmf, bool)`

GetVmfOk returns a tuple with the Vmf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmf

`func (o *OrderCreateRequest) SetVmf(v OrderCreateRequestVmf)`

SetVmf sets Vmf field to given value.

### HasVmf

`func (o *OrderCreateRequest) HasVmf() bool`

HasVmf returns a boolean if a field has been set.

### GetShipToInfo

`func (o *OrderCreateRequest) GetShipToInfo() OrderCreateRequestShipToInfo`

GetShipToInfo returns the ShipToInfo field if non-nil, zero value otherwise.

### GetShipToInfoOk

`func (o *OrderCreateRequest) GetShipToInfoOk() (*OrderCreateRequestShipToInfo, bool)`

GetShipToInfoOk returns a tuple with the ShipToInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToInfo

`func (o *OrderCreateRequest) SetShipToInfo(v OrderCreateRequestShipToInfo)`

SetShipToInfo sets ShipToInfo field to given value.

### HasShipToInfo

`func (o *OrderCreateRequest) HasShipToInfo() bool`

HasShipToInfo returns a boolean if a field has been set.

### GetEndUserInfo

`func (o *OrderCreateRequest) GetEndUserInfo() OrderCreateRequestEndUserInfo`

GetEndUserInfo returns the EndUserInfo field if non-nil, zero value otherwise.

### GetEndUserInfoOk

`func (o *OrderCreateRequest) GetEndUserInfoOk() (*OrderCreateRequestEndUserInfo, bool)`

GetEndUserInfoOk returns a tuple with the EndUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserInfo

`func (o *OrderCreateRequest) SetEndUserInfo(v OrderCreateRequestEndUserInfo)`

SetEndUserInfo sets EndUserInfo field to given value.

### HasEndUserInfo

`func (o *OrderCreateRequest) HasEndUserInfo() bool`

HasEndUserInfo returns a boolean if a field has been set.

### GetLines

`func (o *OrderCreateRequest) GetLines() []OrderCreateRequestLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *OrderCreateRequest) GetLinesOk() (*[]OrderCreateRequestLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *OrderCreateRequest) SetLines(v []OrderCreateRequestLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *OrderCreateRequest) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetShipmentDetails

`func (o *OrderCreateRequest) GetShipmentDetails() OrderCreateRequestShipmentDetails`

GetShipmentDetails returns the ShipmentDetails field if non-nil, zero value otherwise.

### GetShipmentDetailsOk

`func (o *OrderCreateRequest) GetShipmentDetailsOk() (*OrderCreateRequestShipmentDetails, bool)`

GetShipmentDetailsOk returns a tuple with the ShipmentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentDetails

`func (o *OrderCreateRequest) SetShipmentDetails(v OrderCreateRequestShipmentDetails)`

SetShipmentDetails sets ShipmentDetails field to given value.

### HasShipmentDetails

`func (o *OrderCreateRequest) HasShipmentDetails() bool`

HasShipmentDetails returns a boolean if a field has been set.

### GetAdditionalAttributes

`func (o *OrderCreateRequest) GetAdditionalAttributes() []OrderCreateRequestAdditionalAttributesInner`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *OrderCreateRequest) GetAdditionalAttributesOk() (*[]OrderCreateRequestAdditionalAttributesInner, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *OrderCreateRequest) SetAdditionalAttributes(v []OrderCreateRequestAdditionalAttributesInner)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *OrderCreateRequest) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


