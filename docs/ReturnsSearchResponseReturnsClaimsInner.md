# ReturnsSearchResponseReturnsClaimsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnClaimId** | Pointer to **string** | A unique return claim Id. | [optional] 
**CaseRequestNumber** | Pointer to **string** | A unique return request number. | [optional] 
**CreatedOn** | Pointer to **string** | The date on which the return request was created.  | [optional] 
**Type** | Pointer to **string** | Type of request. | [optional] 
**ReturnReason** | Pointer to **string** | The reason for the return. | [optional] 
**ReferenceNumber** | Pointer to **string** | The reference number for the return. | [optional] 
**EstimatedTotalValue** | Pointer to **string** | The estimated total value of the return. | [optional] 
**Credit** | Pointer to **float32** | The amount of credit. | [optional] 
**ModifiedOn** | Pointer to **string** | The date on which the return request was last updated. | [optional] 
**Status** | Pointer to **string** | The status of the request. | [optional] 
**Links** | Pointer to [**[]ReturnsSearchResponseReturnsClaimsInnerLinksInner**](ReturnsSearchResponseReturnsClaimsInnerLinksInner.md) |  | [optional] 

## Methods

### NewReturnsSearchResponseReturnsClaimsInner

`func NewReturnsSearchResponseReturnsClaimsInner() *ReturnsSearchResponseReturnsClaimsInner`

NewReturnsSearchResponseReturnsClaimsInner instantiates a new ReturnsSearchResponseReturnsClaimsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnsSearchResponseReturnsClaimsInnerWithDefaults

`func NewReturnsSearchResponseReturnsClaimsInnerWithDefaults() *ReturnsSearchResponseReturnsClaimsInner`

NewReturnsSearchResponseReturnsClaimsInnerWithDefaults instantiates a new ReturnsSearchResponseReturnsClaimsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnClaimId

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetReturnClaimId() string`

GetReturnClaimId returns the ReturnClaimId field if non-nil, zero value otherwise.

### GetReturnClaimIdOk

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetReturnClaimIdOk() (*string, bool)`

GetReturnClaimIdOk returns a tuple with the ReturnClaimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnClaimId

`func (o *ReturnsSearchResponseReturnsClaimsInner) SetReturnClaimId(v string)`

SetReturnClaimId sets ReturnClaimId field to given value.

### HasReturnClaimId

`func (o *ReturnsSearchResponseReturnsClaimsInner) HasReturnClaimId() bool`

HasReturnClaimId returns a boolean if a field has been set.

### GetCaseRequestNumber

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetCaseRequestNumber() string`

GetCaseRequestNumber returns the CaseRequestNumber field if non-nil, zero value otherwise.

### GetCaseRequestNumberOk

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetCaseRequestNumberOk() (*string, bool)`

GetCaseRequestNumberOk returns a tuple with the CaseRequestNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseRequestNumber

`func (o *ReturnsSearchResponseReturnsClaimsInner) SetCaseRequestNumber(v string)`

SetCaseRequestNumber sets CaseRequestNumber field to given value.

### HasCaseRequestNumber

`func (o *ReturnsSearchResponseReturnsClaimsInner) HasCaseRequestNumber() bool`

HasCaseRequestNumber returns a boolean if a field has been set.

### GetCreatedOn

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *ReturnsSearchResponseReturnsClaimsInner) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *ReturnsSearchResponseReturnsClaimsInner) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetType

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReturnsSearchResponseReturnsClaimsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReturnsSearchResponseReturnsClaimsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReturnReason

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetReturnReason() string`

GetReturnReason returns the ReturnReason field if non-nil, zero value otherwise.

### GetReturnReasonOk

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetReturnReasonOk() (*string, bool)`

GetReturnReasonOk returns a tuple with the ReturnReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnReason

`func (o *ReturnsSearchResponseReturnsClaimsInner) SetReturnReason(v string)`

SetReturnReason sets ReturnReason field to given value.

### HasReturnReason

`func (o *ReturnsSearchResponseReturnsClaimsInner) HasReturnReason() bool`

HasReturnReason returns a boolean if a field has been set.

### GetReferenceNumber

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *ReturnsSearchResponseReturnsClaimsInner) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *ReturnsSearchResponseReturnsClaimsInner) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### GetEstimatedTotalValue

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetEstimatedTotalValue() string`

GetEstimatedTotalValue returns the EstimatedTotalValue field if non-nil, zero value otherwise.

### GetEstimatedTotalValueOk

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetEstimatedTotalValueOk() (*string, bool)`

GetEstimatedTotalValueOk returns a tuple with the EstimatedTotalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTotalValue

`func (o *ReturnsSearchResponseReturnsClaimsInner) SetEstimatedTotalValue(v string)`

SetEstimatedTotalValue sets EstimatedTotalValue field to given value.

### HasEstimatedTotalValue

`func (o *ReturnsSearchResponseReturnsClaimsInner) HasEstimatedTotalValue() bool`

HasEstimatedTotalValue returns a boolean if a field has been set.

### GetCredit

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetCredit() float32`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetCreditOk() (*float32, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *ReturnsSearchResponseReturnsClaimsInner) SetCredit(v float32)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *ReturnsSearchResponseReturnsClaimsInner) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetModifiedOn

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetModifiedOn() string`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetModifiedOnOk() (*string, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *ReturnsSearchResponseReturnsClaimsInner) SetModifiedOn(v string)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *ReturnsSearchResponseReturnsClaimsInner) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetStatus

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReturnsSearchResponseReturnsClaimsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReturnsSearchResponseReturnsClaimsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetLinks() []ReturnsSearchResponseReturnsClaimsInnerLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ReturnsSearchResponseReturnsClaimsInner) GetLinksOk() (*[]ReturnsSearchResponseReturnsClaimsInnerLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ReturnsSearchResponseReturnsClaimsInner) SetLinks(v []ReturnsSearchResponseReturnsClaimsInnerLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ReturnsSearchResponseReturnsClaimsInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


