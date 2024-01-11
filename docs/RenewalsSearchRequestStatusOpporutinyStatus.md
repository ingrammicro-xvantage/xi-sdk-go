# RenewalsSearchRequestStatusOpporutinyStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | The value of opportunity status, it can be either Open or Closed. | [optional] 
**SubStatus** | Pointer to **string** | The sub-status of Opportunity status. Possible sub-status values for Open opportunity status are Ready to order, Quote pending, Notification sent, Future, and Quote requested. Possible sub-status values for Closed opportunity status are All, Ordered, Quote closed-contract sales desk, Expired, Transition to new/upsell, Lost to competitior, Consolidated, Transitioned to cloud, Renewal went direct, EOL, End user out of business, Void, Other, and Cancelled. | [optional] 

## Methods

### NewRenewalsSearchRequestStatusOpporutinyStatus

`func NewRenewalsSearchRequestStatusOpporutinyStatus() *RenewalsSearchRequestStatusOpporutinyStatus`

NewRenewalsSearchRequestStatusOpporutinyStatus instantiates a new RenewalsSearchRequestStatusOpporutinyStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenewalsSearchRequestStatusOpporutinyStatusWithDefaults

`func NewRenewalsSearchRequestStatusOpporutinyStatusWithDefaults() *RenewalsSearchRequestStatusOpporutinyStatus`

NewRenewalsSearchRequestStatusOpporutinyStatusWithDefaults instantiates a new RenewalsSearchRequestStatusOpporutinyStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *RenewalsSearchRequestStatusOpporutinyStatus) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RenewalsSearchRequestStatusOpporutinyStatus) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RenewalsSearchRequestStatusOpporutinyStatus) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RenewalsSearchRequestStatusOpporutinyStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSubStatus

`func (o *RenewalsSearchRequestStatusOpporutinyStatus) GetSubStatus() string`

GetSubStatus returns the SubStatus field if non-nil, zero value otherwise.

### GetSubStatusOk

`func (o *RenewalsSearchRequestStatusOpporutinyStatus) GetSubStatusOk() (*string, bool)`

GetSubStatusOk returns a tuple with the SubStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStatus

`func (o *RenewalsSearchRequestStatusOpporutinyStatus) SetSubStatus(v string)`

SetSubStatus sets SubStatus field to given value.

### HasSubStatus

`func (o *RenewalsSearchRequestStatusOpporutinyStatus) HasSubStatus() bool`

HasSubStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


