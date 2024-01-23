# RenewalsSearchResponseRenewalsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RenewalId** | Pointer to **string** | Unique renewal ID. | [optional] 
**CustomerOrderNumber** | Pointer to **string** | The reseller&#39;s order number for reference in their system. | [optional] 
**ReferenceNumber** | Pointer to **string** | Renewal reference number. It could be notification id or quote number. | [optional] 
**EndUser** | Pointer to **string** | The company name for the end user/customer. | [optional] 
**Vendor** | Pointer to **string** | The name of the vendor. | [optional] 
**ExpirationDate** | Pointer to **string** | Renewal expiration date. | [optional] 
**RenewalValue** | Pointer to **string** | The value of the renewal. | [optional] 
**Status** | Pointer to **string** | The status of the renewal. | [optional] 
**Links** | Pointer to [**[]RenewalsSearchResponseRenewalsInnerLinksInner**](RenewalsSearchResponseRenewalsInnerLinksInner.md) |  | [optional] 

## Methods

### NewRenewalsSearchResponseRenewalsInner

`func NewRenewalsSearchResponseRenewalsInner() *RenewalsSearchResponseRenewalsInner`

NewRenewalsSearchResponseRenewalsInner instantiates a new RenewalsSearchResponseRenewalsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenewalsSearchResponseRenewalsInnerWithDefaults

`func NewRenewalsSearchResponseRenewalsInnerWithDefaults() *RenewalsSearchResponseRenewalsInner`

NewRenewalsSearchResponseRenewalsInnerWithDefaults instantiates a new RenewalsSearchResponseRenewalsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRenewalId

`func (o *RenewalsSearchResponseRenewalsInner) GetRenewalId() string`

GetRenewalId returns the RenewalId field if non-nil, zero value otherwise.

### GetRenewalIdOk

`func (o *RenewalsSearchResponseRenewalsInner) GetRenewalIdOk() (*string, bool)`

GetRenewalIdOk returns a tuple with the RenewalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalId

`func (o *RenewalsSearchResponseRenewalsInner) SetRenewalId(v string)`

SetRenewalId sets RenewalId field to given value.

### HasRenewalId

`func (o *RenewalsSearchResponseRenewalsInner) HasRenewalId() bool`

HasRenewalId returns a boolean if a field has been set.

### GetCustomerOrderNumber

`func (o *RenewalsSearchResponseRenewalsInner) GetCustomerOrderNumber() string`

GetCustomerOrderNumber returns the CustomerOrderNumber field if non-nil, zero value otherwise.

### GetCustomerOrderNumberOk

`func (o *RenewalsSearchResponseRenewalsInner) GetCustomerOrderNumberOk() (*string, bool)`

GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOrderNumber

`func (o *RenewalsSearchResponseRenewalsInner) SetCustomerOrderNumber(v string)`

SetCustomerOrderNumber sets CustomerOrderNumber field to given value.

### HasCustomerOrderNumber

`func (o *RenewalsSearchResponseRenewalsInner) HasCustomerOrderNumber() bool`

HasCustomerOrderNumber returns a boolean if a field has been set.

### GetReferenceNumber

`func (o *RenewalsSearchResponseRenewalsInner) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *RenewalsSearchResponseRenewalsInner) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *RenewalsSearchResponseRenewalsInner) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *RenewalsSearchResponseRenewalsInner) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### GetEndUser

`func (o *RenewalsSearchResponseRenewalsInner) GetEndUser() string`

GetEndUser returns the EndUser field if non-nil, zero value otherwise.

### GetEndUserOk

`func (o *RenewalsSearchResponseRenewalsInner) GetEndUserOk() (*string, bool)`

GetEndUserOk returns a tuple with the EndUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUser

`func (o *RenewalsSearchResponseRenewalsInner) SetEndUser(v string)`

SetEndUser sets EndUser field to given value.

### HasEndUser

`func (o *RenewalsSearchResponseRenewalsInner) HasEndUser() bool`

HasEndUser returns a boolean if a field has been set.

### GetVendor

`func (o *RenewalsSearchResponseRenewalsInner) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *RenewalsSearchResponseRenewalsInner) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *RenewalsSearchResponseRenewalsInner) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *RenewalsSearchResponseRenewalsInner) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetExpirationDate

`func (o *RenewalsSearchResponseRenewalsInner) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *RenewalsSearchResponseRenewalsInner) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *RenewalsSearchResponseRenewalsInner) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *RenewalsSearchResponseRenewalsInner) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetRenewalValue

`func (o *RenewalsSearchResponseRenewalsInner) GetRenewalValue() string`

GetRenewalValue returns the RenewalValue field if non-nil, zero value otherwise.

### GetRenewalValueOk

`func (o *RenewalsSearchResponseRenewalsInner) GetRenewalValueOk() (*string, bool)`

GetRenewalValueOk returns a tuple with the RenewalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalValue

`func (o *RenewalsSearchResponseRenewalsInner) SetRenewalValue(v string)`

SetRenewalValue sets RenewalValue field to given value.

### HasRenewalValue

`func (o *RenewalsSearchResponseRenewalsInner) HasRenewalValue() bool`

HasRenewalValue returns a boolean if a field has been set.

### GetStatus

`func (o *RenewalsSearchResponseRenewalsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RenewalsSearchResponseRenewalsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RenewalsSearchResponseRenewalsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RenewalsSearchResponseRenewalsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *RenewalsSearchResponseRenewalsInner) GetLinks() []RenewalsSearchResponseRenewalsInnerLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RenewalsSearchResponseRenewalsInner) GetLinksOk() (*[]RenewalsSearchResponseRenewalsInnerLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RenewalsSearchResponseRenewalsInner) SetLinks(v []RenewalsSearchResponseRenewalsInnerLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RenewalsSearchResponseRenewalsInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


