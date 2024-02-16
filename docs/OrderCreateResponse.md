# OrderCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerOrderNumber** | Pointer to **string** | The reseller&#39;s unique PO/Order number. | [optional] 
**EndCustomerOrderNumber** | Pointer to **string** | The end user/customer&#39;s Purchase Order number. | [optional] 
**BillToAddressId** | Pointer to **string** | Suffix used to identify billing address. Created during onboarding. Resellers are provided with one or more address IDs depending on how many bill to addresses they need for various flooring companies they are using for credit | [optional] 
**SpecialBidNumber** | Pointer to **string** | The bid number provided to the reseller by the vendor for special pricing and discounts. Line-level bid numbers take precedence over header-level bid numbers. | [optional] 
**OrderSplit** | Pointer to **bool** | true for multiple orders | [optional] 
**ProcessedPartially** | Pointer to **bool** | true for partial order succesfully placed | [optional] 
**PurchaseOrderTotal** | Pointer to **float32** | Total of all the orders including taxes and fees. | [optional] 
**ShipToInfo** | Pointer to [**OrderCreateResponseShipToInfo**](OrderCreateResponseShipToInfo.md) |  | [optional] 
**EndUserInfo** | Pointer to [**OrderCreateResponseEndUserInfo**](OrderCreateResponseEndUserInfo.md) |  | [optional] 
**Orders** | Pointer to [**[]OrderCreateResponseOrdersInner**](OrderCreateResponseOrdersInner.md) | Order-level details. | [optional] 

## Methods

### NewOrderCreateResponse

`func NewOrderCreateResponse() *OrderCreateResponse`

NewOrderCreateResponse instantiates a new OrderCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateResponseWithDefaults

`func NewOrderCreateResponseWithDefaults() *OrderCreateResponse`

NewOrderCreateResponseWithDefaults instantiates a new OrderCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerOrderNumber

`func (o *OrderCreateResponse) GetCustomerOrderNumber() string`

GetCustomerOrderNumber returns the CustomerOrderNumber field if non-nil, zero value otherwise.

### GetCustomerOrderNumberOk

`func (o *OrderCreateResponse) GetCustomerOrderNumberOk() (*string, bool)`

GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOrderNumber

`func (o *OrderCreateResponse) SetCustomerOrderNumber(v string)`

SetCustomerOrderNumber sets CustomerOrderNumber field to given value.

### HasCustomerOrderNumber

`func (o *OrderCreateResponse) HasCustomerOrderNumber() bool`

HasCustomerOrderNumber returns a boolean if a field has been set.

### GetEndCustomerOrderNumber

`func (o *OrderCreateResponse) GetEndCustomerOrderNumber() string`

GetEndCustomerOrderNumber returns the EndCustomerOrderNumber field if non-nil, zero value otherwise.

### GetEndCustomerOrderNumberOk

`func (o *OrderCreateResponse) GetEndCustomerOrderNumberOk() (*string, bool)`

GetEndCustomerOrderNumberOk returns a tuple with the EndCustomerOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCustomerOrderNumber

`func (o *OrderCreateResponse) SetEndCustomerOrderNumber(v string)`

SetEndCustomerOrderNumber sets EndCustomerOrderNumber field to given value.

### HasEndCustomerOrderNumber

`func (o *OrderCreateResponse) HasEndCustomerOrderNumber() bool`

HasEndCustomerOrderNumber returns a boolean if a field has been set.

### GetBillToAddressId

`func (o *OrderCreateResponse) GetBillToAddressId() string`

GetBillToAddressId returns the BillToAddressId field if non-nil, zero value otherwise.

### GetBillToAddressIdOk

`func (o *OrderCreateResponse) GetBillToAddressIdOk() (*string, bool)`

GetBillToAddressIdOk returns a tuple with the BillToAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToAddressId

`func (o *OrderCreateResponse) SetBillToAddressId(v string)`

SetBillToAddressId sets BillToAddressId field to given value.

### HasBillToAddressId

`func (o *OrderCreateResponse) HasBillToAddressId() bool`

HasBillToAddressId returns a boolean if a field has been set.

### GetSpecialBidNumber

`func (o *OrderCreateResponse) GetSpecialBidNumber() string`

GetSpecialBidNumber returns the SpecialBidNumber field if non-nil, zero value otherwise.

### GetSpecialBidNumberOk

`func (o *OrderCreateResponse) GetSpecialBidNumberOk() (*string, bool)`

GetSpecialBidNumberOk returns a tuple with the SpecialBidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialBidNumber

`func (o *OrderCreateResponse) SetSpecialBidNumber(v string)`

SetSpecialBidNumber sets SpecialBidNumber field to given value.

### HasSpecialBidNumber

`func (o *OrderCreateResponse) HasSpecialBidNumber() bool`

HasSpecialBidNumber returns a boolean if a field has been set.

### GetOrderSplit

`func (o *OrderCreateResponse) GetOrderSplit() bool`

GetOrderSplit returns the OrderSplit field if non-nil, zero value otherwise.

### GetOrderSplitOk

`func (o *OrderCreateResponse) GetOrderSplitOk() (*bool, bool)`

GetOrderSplitOk returns a tuple with the OrderSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSplit

`func (o *OrderCreateResponse) SetOrderSplit(v bool)`

SetOrderSplit sets OrderSplit field to given value.

### HasOrderSplit

`func (o *OrderCreateResponse) HasOrderSplit() bool`

HasOrderSplit returns a boolean if a field has been set.

### GetProcessedPartially

`func (o *OrderCreateResponse) GetProcessedPartially() bool`

GetProcessedPartially returns the ProcessedPartially field if non-nil, zero value otherwise.

### GetProcessedPartiallyOk

`func (o *OrderCreateResponse) GetProcessedPartiallyOk() (*bool, bool)`

GetProcessedPartiallyOk returns a tuple with the ProcessedPartially field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedPartially

`func (o *OrderCreateResponse) SetProcessedPartially(v bool)`

SetProcessedPartially sets ProcessedPartially field to given value.

### HasProcessedPartially

`func (o *OrderCreateResponse) HasProcessedPartially() bool`

HasProcessedPartially returns a boolean if a field has been set.

### GetPurchaseOrderTotal

`func (o *OrderCreateResponse) GetPurchaseOrderTotal() float32`

GetPurchaseOrderTotal returns the PurchaseOrderTotal field if non-nil, zero value otherwise.

### GetPurchaseOrderTotalOk

`func (o *OrderCreateResponse) GetPurchaseOrderTotalOk() (*float32, bool)`

GetPurchaseOrderTotalOk returns a tuple with the PurchaseOrderTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderTotal

`func (o *OrderCreateResponse) SetPurchaseOrderTotal(v float32)`

SetPurchaseOrderTotal sets PurchaseOrderTotal field to given value.

### HasPurchaseOrderTotal

`func (o *OrderCreateResponse) HasPurchaseOrderTotal() bool`

HasPurchaseOrderTotal returns a boolean if a field has been set.

### GetShipToInfo

`func (o *OrderCreateResponse) GetShipToInfo() OrderCreateResponseShipToInfo`

GetShipToInfo returns the ShipToInfo field if non-nil, zero value otherwise.

### GetShipToInfoOk

`func (o *OrderCreateResponse) GetShipToInfoOk() (*OrderCreateResponseShipToInfo, bool)`

GetShipToInfoOk returns a tuple with the ShipToInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToInfo

`func (o *OrderCreateResponse) SetShipToInfo(v OrderCreateResponseShipToInfo)`

SetShipToInfo sets ShipToInfo field to given value.

### HasShipToInfo

`func (o *OrderCreateResponse) HasShipToInfo() bool`

HasShipToInfo returns a boolean if a field has been set.

### GetEndUserInfo

`func (o *OrderCreateResponse) GetEndUserInfo() OrderCreateResponseEndUserInfo`

GetEndUserInfo returns the EndUserInfo field if non-nil, zero value otherwise.

### GetEndUserInfoOk

`func (o *OrderCreateResponse) GetEndUserInfoOk() (*OrderCreateResponseEndUserInfo, bool)`

GetEndUserInfoOk returns a tuple with the EndUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserInfo

`func (o *OrderCreateResponse) SetEndUserInfo(v OrderCreateResponseEndUserInfo)`

SetEndUserInfo sets EndUserInfo field to given value.

### HasEndUserInfo

`func (o *OrderCreateResponse) HasEndUserInfo() bool`

HasEndUserInfo returns a boolean if a field has been set.

### GetOrders

`func (o *OrderCreateResponse) GetOrders() []OrderCreateResponseOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *OrderCreateResponse) GetOrdersOk() (*[]OrderCreateResponseOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *OrderCreateResponse) SetOrders(v []OrderCreateResponseOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *OrderCreateResponse) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


