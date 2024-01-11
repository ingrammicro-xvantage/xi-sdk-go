# ReturnsDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeOfDetails** | Pointer to **string** | The type of the details. Return or Claim. | [optional] 
**RmaClaimId** | Pointer to **string** | The rmaClaimId claim id. | [optional] 
**CaseRequestNumber** | Pointer to **string** | A unique return request number. | [optional] 
**CreatedOn** | Pointer to **string** | The date on which the return request was created. | [optional] 
**ReturnReason** | Pointer to **string** | The reason for the return. | [optional] 
**ReferenceNumber** | Pointer to **string** | The reference number for the return. | [optional] 
**Status** | Pointer to **string** | The status of the request. | [optional] 
**ReturnWarehouseAddress** | Pointer to **string** | The address of the return warehouse. | [optional] 
**Products** | Pointer to [**[]ReturnsDetailsResponseProductsInner**](ReturnsDetailsResponseProductsInner.md) |  | [optional] 
**SubTotal** | Pointer to **float32** | Sub total amount of the return request. | [optional] 
**Tax** | Pointer to **float32** | The tax amount of the return request. | [optional] 
**AdditionalFees** | Pointer to **float32** | The additional fees for the return request. | [optional] 
**EstimatedTotal** | Pointer to **float32** | The total estimated amount for the return request. | [optional] 

## Methods

### NewReturnsDetailsResponse

`func NewReturnsDetailsResponse() *ReturnsDetailsResponse`

NewReturnsDetailsResponse instantiates a new ReturnsDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnsDetailsResponseWithDefaults

`func NewReturnsDetailsResponseWithDefaults() *ReturnsDetailsResponse`

NewReturnsDetailsResponseWithDefaults instantiates a new ReturnsDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeOfDetails

`func (o *ReturnsDetailsResponse) GetTypeOfDetails() string`

GetTypeOfDetails returns the TypeOfDetails field if non-nil, zero value otherwise.

### GetTypeOfDetailsOk

`func (o *ReturnsDetailsResponse) GetTypeOfDetailsOk() (*string, bool)`

GetTypeOfDetailsOk returns a tuple with the TypeOfDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfDetails

`func (o *ReturnsDetailsResponse) SetTypeOfDetails(v string)`

SetTypeOfDetails sets TypeOfDetails field to given value.

### HasTypeOfDetails

`func (o *ReturnsDetailsResponse) HasTypeOfDetails() bool`

HasTypeOfDetails returns a boolean if a field has been set.

### GetRmaClaimId

`func (o *ReturnsDetailsResponse) GetRmaClaimId() string`

GetRmaClaimId returns the RmaClaimId field if non-nil, zero value otherwise.

### GetRmaClaimIdOk

`func (o *ReturnsDetailsResponse) GetRmaClaimIdOk() (*string, bool)`

GetRmaClaimIdOk returns a tuple with the RmaClaimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmaClaimId

`func (o *ReturnsDetailsResponse) SetRmaClaimId(v string)`

SetRmaClaimId sets RmaClaimId field to given value.

### HasRmaClaimId

`func (o *ReturnsDetailsResponse) HasRmaClaimId() bool`

HasRmaClaimId returns a boolean if a field has been set.

### GetCaseRequestNumber

`func (o *ReturnsDetailsResponse) GetCaseRequestNumber() string`

GetCaseRequestNumber returns the CaseRequestNumber field if non-nil, zero value otherwise.

### GetCaseRequestNumberOk

`func (o *ReturnsDetailsResponse) GetCaseRequestNumberOk() (*string, bool)`

GetCaseRequestNumberOk returns a tuple with the CaseRequestNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseRequestNumber

`func (o *ReturnsDetailsResponse) SetCaseRequestNumber(v string)`

SetCaseRequestNumber sets CaseRequestNumber field to given value.

### HasCaseRequestNumber

`func (o *ReturnsDetailsResponse) HasCaseRequestNumber() bool`

HasCaseRequestNumber returns a boolean if a field has been set.

### GetCreatedOn

`func (o *ReturnsDetailsResponse) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *ReturnsDetailsResponse) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *ReturnsDetailsResponse) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *ReturnsDetailsResponse) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetReturnReason

`func (o *ReturnsDetailsResponse) GetReturnReason() string`

GetReturnReason returns the ReturnReason field if non-nil, zero value otherwise.

### GetReturnReasonOk

`func (o *ReturnsDetailsResponse) GetReturnReasonOk() (*string, bool)`

GetReturnReasonOk returns a tuple with the ReturnReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnReason

`func (o *ReturnsDetailsResponse) SetReturnReason(v string)`

SetReturnReason sets ReturnReason field to given value.

### HasReturnReason

`func (o *ReturnsDetailsResponse) HasReturnReason() bool`

HasReturnReason returns a boolean if a field has been set.

### GetReferenceNumber

`func (o *ReturnsDetailsResponse) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *ReturnsDetailsResponse) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *ReturnsDetailsResponse) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *ReturnsDetailsResponse) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### GetStatus

`func (o *ReturnsDetailsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReturnsDetailsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReturnsDetailsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReturnsDetailsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReturnWarehouseAddress

`func (o *ReturnsDetailsResponse) GetReturnWarehouseAddress() string`

GetReturnWarehouseAddress returns the ReturnWarehouseAddress field if non-nil, zero value otherwise.

### GetReturnWarehouseAddressOk

`func (o *ReturnsDetailsResponse) GetReturnWarehouseAddressOk() (*string, bool)`

GetReturnWarehouseAddressOk returns a tuple with the ReturnWarehouseAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnWarehouseAddress

`func (o *ReturnsDetailsResponse) SetReturnWarehouseAddress(v string)`

SetReturnWarehouseAddress sets ReturnWarehouseAddress field to given value.

### HasReturnWarehouseAddress

`func (o *ReturnsDetailsResponse) HasReturnWarehouseAddress() bool`

HasReturnWarehouseAddress returns a boolean if a field has been set.

### GetProducts

`func (o *ReturnsDetailsResponse) GetProducts() []ReturnsDetailsResponseProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ReturnsDetailsResponse) GetProductsOk() (*[]ReturnsDetailsResponseProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ReturnsDetailsResponse) SetProducts(v []ReturnsDetailsResponseProductsInner)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *ReturnsDetailsResponse) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetSubTotal

`func (o *ReturnsDetailsResponse) GetSubTotal() float32`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *ReturnsDetailsResponse) GetSubTotalOk() (*float32, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *ReturnsDetailsResponse) SetSubTotal(v float32)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotal

`func (o *ReturnsDetailsResponse) HasSubTotal() bool`

HasSubTotal returns a boolean if a field has been set.

### GetTax

`func (o *ReturnsDetailsResponse) GetTax() float32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *ReturnsDetailsResponse) GetTaxOk() (*float32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *ReturnsDetailsResponse) SetTax(v float32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *ReturnsDetailsResponse) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetAdditionalFees

`func (o *ReturnsDetailsResponse) GetAdditionalFees() float32`

GetAdditionalFees returns the AdditionalFees field if non-nil, zero value otherwise.

### GetAdditionalFeesOk

`func (o *ReturnsDetailsResponse) GetAdditionalFeesOk() (*float32, bool)`

GetAdditionalFeesOk returns a tuple with the AdditionalFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFees

`func (o *ReturnsDetailsResponse) SetAdditionalFees(v float32)`

SetAdditionalFees sets AdditionalFees field to given value.

### HasAdditionalFees

`func (o *ReturnsDetailsResponse) HasAdditionalFees() bool`

HasAdditionalFees returns a boolean if a field has been set.

### GetEstimatedTotal

`func (o *ReturnsDetailsResponse) GetEstimatedTotal() float32`

GetEstimatedTotal returns the EstimatedTotal field if non-nil, zero value otherwise.

### GetEstimatedTotalOk

`func (o *ReturnsDetailsResponse) GetEstimatedTotalOk() (*float32, bool)`

GetEstimatedTotalOk returns a tuple with the EstimatedTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTotal

`func (o *ReturnsDetailsResponse) SetEstimatedTotal(v float32)`

SetEstimatedTotal sets EstimatedTotal field to given value.

### HasEstimatedTotal

`func (o *ReturnsDetailsResponse) HasEstimatedTotal() bool`

HasEstimatedTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


