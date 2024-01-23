# OrderDeleteRequestServicerequestOrderDeleteRequestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryDate** | **string** | Date order entered | 
**OrderBranch** | **string** | Ingram Micro&#39;s first 2 numbers of the order number | 
**OrderNumber** | **string** | Ingram Micro&#39;s middle 6 numbers of the order# | 
**RejectionCode** | Pointer to **string** |  | [optional] 
**DistributionNumber** | Pointer to **string** | Ingram Micro&#39;s suffix number of the order# | [optional] 
**ShipmentNumber** | Pointer to **string** | Ingram Micro&#39;s last number of the order# | [optional] 
**OperatorID** | Pointer to **string** | Ingram ID(not required) | [optional] 

## Methods

### NewOrderDeleteRequestServicerequestOrderDeleteRequestDetails

`func NewOrderDeleteRequestServicerequestOrderDeleteRequestDetails(entryDate string, orderBranch string, orderNumber string, ) *OrderDeleteRequestServicerequestOrderDeleteRequestDetails`

NewOrderDeleteRequestServicerequestOrderDeleteRequestDetails instantiates a new OrderDeleteRequestServicerequestOrderDeleteRequestDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDeleteRequestServicerequestOrderDeleteRequestDetailsWithDefaults

`func NewOrderDeleteRequestServicerequestOrderDeleteRequestDetailsWithDefaults() *OrderDeleteRequestServicerequestOrderDeleteRequestDetails`

NewOrderDeleteRequestServicerequestOrderDeleteRequestDetailsWithDefaults instantiates a new OrderDeleteRequestServicerequestOrderDeleteRequestDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryDate

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetEntryDate() string`

GetEntryDate returns the EntryDate field if non-nil, zero value otherwise.

### GetEntryDateOk

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetEntryDateOk() (*string, bool)`

GetEntryDateOk returns a tuple with the EntryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryDate

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) SetEntryDate(v string)`

SetEntryDate sets EntryDate field to given value.


### GetOrderBranch

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetOrderBranch() string`

GetOrderBranch returns the OrderBranch field if non-nil, zero value otherwise.

### GetOrderBranchOk

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetOrderBranchOk() (*string, bool)`

GetOrderBranchOk returns a tuple with the OrderBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBranch

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) SetOrderBranch(v string)`

SetOrderBranch sets OrderBranch field to given value.


### GetOrderNumber

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.


### GetRejectionCode

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetRejectionCode() string`

GetRejectionCode returns the RejectionCode field if non-nil, zero value otherwise.

### GetRejectionCodeOk

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetRejectionCodeOk() (*string, bool)`

GetRejectionCodeOk returns a tuple with the RejectionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionCode

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) SetRejectionCode(v string)`

SetRejectionCode sets RejectionCode field to given value.

### HasRejectionCode

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) HasRejectionCode() bool`

HasRejectionCode returns a boolean if a field has been set.

### GetDistributionNumber

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetDistributionNumber() string`

GetDistributionNumber returns the DistributionNumber field if non-nil, zero value otherwise.

### GetDistributionNumberOk

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetDistributionNumberOk() (*string, bool)`

GetDistributionNumberOk returns a tuple with the DistributionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionNumber

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) SetDistributionNumber(v string)`

SetDistributionNumber sets DistributionNumber field to given value.

### HasDistributionNumber

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) HasDistributionNumber() bool`

HasDistributionNumber returns a boolean if a field has been set.

### GetShipmentNumber

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetShipmentNumber() string`

GetShipmentNumber returns the ShipmentNumber field if non-nil, zero value otherwise.

### GetShipmentNumberOk

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetShipmentNumberOk() (*string, bool)`

GetShipmentNumberOk returns a tuple with the ShipmentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentNumber

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) SetShipmentNumber(v string)`

SetShipmentNumber sets ShipmentNumber field to given value.

### HasShipmentNumber

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) HasShipmentNumber() bool`

HasShipmentNumber returns a boolean if a field has been set.

### GetOperatorID

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetOperatorID() string`

GetOperatorID returns the OperatorID field if non-nil, zero value otherwise.

### GetOperatorIDOk

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) GetOperatorIDOk() (*string, bool)`

GetOperatorIDOk returns a tuple with the OperatorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorID

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) SetOperatorID(v string)`

SetOperatorID sets OperatorID field to given value.

### HasOperatorID

`func (o *OrderDeleteRequestServicerequestOrderDeleteRequestDetails) HasOperatorID() bool`

HasOperatorID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


