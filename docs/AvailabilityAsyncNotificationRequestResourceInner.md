# AvailabilityAsyncNotificationRequestResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | The event name sent in the event request. | [optional] 
**IngramPartNumber** | Pointer to **string** | The Unique IngramMicro part number for the product. | [optional] 
**VendorPartNumber** | Pointer to **string** | The vendors part number for the product. | [optional] 
**VendorName** | Pointer to **string** | The name of the vendor/manufacturer of the product. | [optional] 
**UpcCode** | Pointer to **string** | The UPC code for the product. Consists of 12 numeric digits that are uniquly assigned to each trade item. | [optional] 
**SkuStatus** | Pointer to **string** | Status returned saying whether sku is active. | [optional] 
**BackOrderFlag** | Pointer to **string** | Backordered Flag. | [optional] 
**TotalAvailability** | Pointer to **string** | totalAvailability. | [optional] 
**Links** | Pointer to [**[]AvailabilityAsyncNotificationRequestResourceInnerLinksInner**](AvailabilityAsyncNotificationRequestResourceInnerLinksInner.md) | Link to Order Details for the order(s). | [optional] 

## Methods

### NewAvailabilityAsyncNotificationRequestResourceInner

`func NewAvailabilityAsyncNotificationRequestResourceInner() *AvailabilityAsyncNotificationRequestResourceInner`

NewAvailabilityAsyncNotificationRequestResourceInner instantiates a new AvailabilityAsyncNotificationRequestResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityAsyncNotificationRequestResourceInnerWithDefaults

`func NewAvailabilityAsyncNotificationRequestResourceInnerWithDefaults() *AvailabilityAsyncNotificationRequestResourceInner`

NewAvailabilityAsyncNotificationRequestResourceInnerWithDefaults instantiates a new AvailabilityAsyncNotificationRequestResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *AvailabilityAsyncNotificationRequestResourceInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *AvailabilityAsyncNotificationRequestResourceInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *AvailabilityAsyncNotificationRequestResourceInner) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *AvailabilityAsyncNotificationRequestResourceInner) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *AvailabilityAsyncNotificationRequestResourceInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *AvailabilityAsyncNotificationRequestResourceInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *AvailabilityAsyncNotificationRequestResourceInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *AvailabilityAsyncNotificationRequestResourceInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *AvailabilityAsyncNotificationRequestResourceInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *AvailabilityAsyncNotificationRequestResourceInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *AvailabilityAsyncNotificationRequestResourceInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *AvailabilityAsyncNotificationRequestResourceInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetVendorName

`func (o *AvailabilityAsyncNotificationRequestResourceInner) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *AvailabilityAsyncNotificationRequestResourceInner) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *AvailabilityAsyncNotificationRequestResourceInner) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *AvailabilityAsyncNotificationRequestResourceInner) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetUpcCode

`func (o *AvailabilityAsyncNotificationRequestResourceInner) GetUpcCode() string`

GetUpcCode returns the UpcCode field if non-nil, zero value otherwise.

### GetUpcCodeOk

`func (o *AvailabilityAsyncNotificationRequestResourceInner) GetUpcCodeOk() (*string, bool)`

GetUpcCodeOk returns a tuple with the UpcCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcCode

`func (o *AvailabilityAsyncNotificationRequestResourceInner) SetUpcCode(v string)`

SetUpcCode sets UpcCode field to given value.

### HasUpcCode

`func (o *AvailabilityAsyncNotificationRequestResourceInner) HasUpcCode() bool`

HasUpcCode returns a boolean if a field has been set.

### GetSkuStatus

`func (o *AvailabilityAsyncNotificationRequestResourceInner) GetSkuStatus() string`

GetSkuStatus returns the SkuStatus field if non-nil, zero value otherwise.

### GetSkuStatusOk

`func (o *AvailabilityAsyncNotificationRequestResourceInner) GetSkuStatusOk() (*string, bool)`

GetSkuStatusOk returns a tuple with the SkuStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuStatus

`func (o *AvailabilityAsyncNotificationRequestResourceInner) SetSkuStatus(v string)`

SetSkuStatus sets SkuStatus field to given value.

### HasSkuStatus

`func (o *AvailabilityAsyncNotificationRequestResourceInner) HasSkuStatus() bool`

HasSkuStatus returns a boolean if a field has been set.

### GetBackOrderFlag

`func (o *AvailabilityAsyncNotificationRequestResourceInner) GetBackOrderFlag() string`

GetBackOrderFlag returns the BackOrderFlag field if non-nil, zero value otherwise.

### GetBackOrderFlagOk

`func (o *AvailabilityAsyncNotificationRequestResourceInner) GetBackOrderFlagOk() (*string, bool)`

GetBackOrderFlagOk returns a tuple with the BackOrderFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackOrderFlag

`func (o *AvailabilityAsyncNotificationRequestResourceInner) SetBackOrderFlag(v string)`

SetBackOrderFlag sets BackOrderFlag field to given value.

### HasBackOrderFlag

`func (o *AvailabilityAsyncNotificationRequestResourceInner) HasBackOrderFlag() bool`

HasBackOrderFlag returns a boolean if a field has been set.

### GetTotalAvailability

`func (o *AvailabilityAsyncNotificationRequestResourceInner) GetTotalAvailability() string`

GetTotalAvailability returns the TotalAvailability field if non-nil, zero value otherwise.

### GetTotalAvailabilityOk

`func (o *AvailabilityAsyncNotificationRequestResourceInner) GetTotalAvailabilityOk() (*string, bool)`

GetTotalAvailabilityOk returns a tuple with the TotalAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAvailability

`func (o *AvailabilityAsyncNotificationRequestResourceInner) SetTotalAvailability(v string)`

SetTotalAvailability sets TotalAvailability field to given value.

### HasTotalAvailability

`func (o *AvailabilityAsyncNotificationRequestResourceInner) HasTotalAvailability() bool`

HasTotalAvailability returns a boolean if a field has been set.

### GetLinks

`func (o *AvailabilityAsyncNotificationRequestResourceInner) GetLinks() []AvailabilityAsyncNotificationRequestResourceInnerLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AvailabilityAsyncNotificationRequestResourceInner) GetLinksOk() (*[]AvailabilityAsyncNotificationRequestResourceInnerLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AvailabilityAsyncNotificationRequestResourceInner) SetLinks(v []AvailabilityAsyncNotificationRequestResourceInnerLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AvailabilityAsyncNotificationRequestResourceInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


