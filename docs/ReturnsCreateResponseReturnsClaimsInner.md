# ReturnsCreateResponseReturnsClaimsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RmaClaimId** | Pointer to **string** | The rmaClaimId claim id. | [optional] 
**CaseRequestNumber** | Pointer to **string** | A unique return request number. | [optional] 
**ReferenceNumber** | Pointer to **string** | The reference number for the return. | [optional] 
**CreatedOn** | Pointer to **string** | The date on which the return request was created.  | [optional] 
**Type** | Pointer to **string** | Type of request. | [optional] 
**ReturnReason** | Pointer to **string** | The reason for the return. | [optional] 
**IngramPartNumber** | Pointer to **string** | Unique line number from Ingram. | [optional] 
**VendorPartNumber** | Pointer to **string** | Vendor Part Number. | [optional] 
**Quantity** | Pointer to **int32** | Return quantity of the product. | [optional] 
**Notes** | Pointer to **string** | Return notes. | [optional] 
**EstimatedTotalValue** | Pointer to **float32** | The estimated total value of the return. | [optional] 
**Credit** | Pointer to **float32** | The amount of credit. | [optional] 
**Status** | Pointer to **string** | The status of the request. | [optional] 
**Links** | Pointer to [**[]ReturnsSearchResponseReturnsClaimsInnerLinksInner**](ReturnsSearchResponseReturnsClaimsInnerLinksInner.md) |  | [optional] 

## Methods

### NewReturnsCreateResponseReturnsClaimsInner

`func NewReturnsCreateResponseReturnsClaimsInner() *ReturnsCreateResponseReturnsClaimsInner`

NewReturnsCreateResponseReturnsClaimsInner instantiates a new ReturnsCreateResponseReturnsClaimsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnsCreateResponseReturnsClaimsInnerWithDefaults

`func NewReturnsCreateResponseReturnsClaimsInnerWithDefaults() *ReturnsCreateResponseReturnsClaimsInner`

NewReturnsCreateResponseReturnsClaimsInnerWithDefaults instantiates a new ReturnsCreateResponseReturnsClaimsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRmaClaimId

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetRmaClaimId() string`

GetRmaClaimId returns the RmaClaimId field if non-nil, zero value otherwise.

### GetRmaClaimIdOk

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetRmaClaimIdOk() (*string, bool)`

GetRmaClaimIdOk returns a tuple with the RmaClaimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmaClaimId

`func (o *ReturnsCreateResponseReturnsClaimsInner) SetRmaClaimId(v string)`

SetRmaClaimId sets RmaClaimId field to given value.

### HasRmaClaimId

`func (o *ReturnsCreateResponseReturnsClaimsInner) HasRmaClaimId() bool`

HasRmaClaimId returns a boolean if a field has been set.

### GetCaseRequestNumber

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetCaseRequestNumber() string`

GetCaseRequestNumber returns the CaseRequestNumber field if non-nil, zero value otherwise.

### GetCaseRequestNumberOk

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetCaseRequestNumberOk() (*string, bool)`

GetCaseRequestNumberOk returns a tuple with the CaseRequestNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseRequestNumber

`func (o *ReturnsCreateResponseReturnsClaimsInner) SetCaseRequestNumber(v string)`

SetCaseRequestNumber sets CaseRequestNumber field to given value.

### HasCaseRequestNumber

`func (o *ReturnsCreateResponseReturnsClaimsInner) HasCaseRequestNumber() bool`

HasCaseRequestNumber returns a boolean if a field has been set.

### GetReferenceNumber

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *ReturnsCreateResponseReturnsClaimsInner) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *ReturnsCreateResponseReturnsClaimsInner) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### GetCreatedOn

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *ReturnsCreateResponseReturnsClaimsInner) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *ReturnsCreateResponseReturnsClaimsInner) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetType

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReturnsCreateResponseReturnsClaimsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReturnsCreateResponseReturnsClaimsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReturnReason

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetReturnReason() string`

GetReturnReason returns the ReturnReason field if non-nil, zero value otherwise.

### GetReturnReasonOk

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetReturnReasonOk() (*string, bool)`

GetReturnReasonOk returns a tuple with the ReturnReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnReason

`func (o *ReturnsCreateResponseReturnsClaimsInner) SetReturnReason(v string)`

SetReturnReason sets ReturnReason field to given value.

### HasReturnReason

`func (o *ReturnsCreateResponseReturnsClaimsInner) HasReturnReason() bool`

HasReturnReason returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *ReturnsCreateResponseReturnsClaimsInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *ReturnsCreateResponseReturnsClaimsInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *ReturnsCreateResponseReturnsClaimsInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *ReturnsCreateResponseReturnsClaimsInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetQuantity

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ReturnsCreateResponseReturnsClaimsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ReturnsCreateResponseReturnsClaimsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetNotes

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ReturnsCreateResponseReturnsClaimsInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ReturnsCreateResponseReturnsClaimsInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetEstimatedTotalValue

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetEstimatedTotalValue() float32`

GetEstimatedTotalValue returns the EstimatedTotalValue field if non-nil, zero value otherwise.

### GetEstimatedTotalValueOk

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetEstimatedTotalValueOk() (*float32, bool)`

GetEstimatedTotalValueOk returns a tuple with the EstimatedTotalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTotalValue

`func (o *ReturnsCreateResponseReturnsClaimsInner) SetEstimatedTotalValue(v float32)`

SetEstimatedTotalValue sets EstimatedTotalValue field to given value.

### HasEstimatedTotalValue

`func (o *ReturnsCreateResponseReturnsClaimsInner) HasEstimatedTotalValue() bool`

HasEstimatedTotalValue returns a boolean if a field has been set.

### GetCredit

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetCredit() float32`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetCreditOk() (*float32, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *ReturnsCreateResponseReturnsClaimsInner) SetCredit(v float32)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *ReturnsCreateResponseReturnsClaimsInner) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetStatus

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReturnsCreateResponseReturnsClaimsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReturnsCreateResponseReturnsClaimsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetLinks() []ReturnsSearchResponseReturnsClaimsInnerLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ReturnsCreateResponseReturnsClaimsInner) GetLinksOk() (*[]ReturnsSearchResponseReturnsClaimsInnerLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ReturnsCreateResponseReturnsClaimsInner) SetLinks(v []ReturnsSearchResponseReturnsClaimsInnerLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ReturnsCreateResponseReturnsClaimsInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


