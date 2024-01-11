# PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Itemstatus** | Pointer to **string** | SUCCESS or FAILED | [optional] 
**Statusmessage** | Pointer to **string** | Description of itemstatus | [optional] 
**Ingrampartnumber** | Pointer to **string** | Ingram Micro part number | [optional] 
**Vendorpartnumber** | Pointer to **string** | Manufacturer/Vendor part number | [optional] 
**Globalskuid** | Pointer to **string** |  | [optional] 
**Customerprice** | Pointer to **float32** | Customer specific price for the product, excluding taxes | [optional] 
**Partdescription1** | Pointer to **string** | Product description part 1 | [optional] 
**Partdescription2** | Pointer to **string** | Product description part 2 | [optional] 
**Vendornumber** | Pointer to **string** |  | [optional] 
**Vendorname** | Pointer to **string** | Name of the vendor | [optional] 
**Cpucode** | Pointer to **string** |  | [optional] 
**Class** | Pointer to **string** | Ingram Micro assigned product classification -  A-Stocked product in all IM warehouses, B-Limited stock in IM warehouses, C-Stocked in fewer wareshouses, D-Ingram discontinued, E-Planned to be phased out as per the vendor, F-Carried for specific customer as per the contract, N-New SKU, O-Discontinued to be liquidated, S-Order for specialized demand, V-Discontinued by vendor, X-Direct Ship products from vendor | [optional] 
**Skustatus** | Pointer to **string** | Identifies if the SKU has been discontinued. | [optional] 
**Mediacpu** | Pointer to **string** |  | [optional] 
**Categorysubcategory** | Pointer to **string** |  | [optional] 
**Retailprice** | Pointer to **float32** |  | [optional] 
**Newmedia** | Pointer to **string** |  | [optional] 
**Enduserrequired** | Pointer to **string** | Y - End user required N - Not required End user | [optional] 
**Backorderflag** | Pointer to **string** | Y- Allow Backorder Flag N- Not allowed | [optional] 
**Skuauthorized** | Pointer to **string** |  | [optional] 
**Extendedvendorpartnumber** | Pointer to **string** |  | [optional] 
**Warehousedetails** | Pointer to [**[]WarehouseListType**](WarehouseListType.md) |  | [optional] 

## Methods

### NewPriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner

`func NewPriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner() *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner`

NewPriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner instantiates a new PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInnerWithDefaults

`func NewPriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInnerWithDefaults() *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner`

NewPriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInnerWithDefaults instantiates a new PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemstatus

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetItemstatus() string`

GetItemstatus returns the Itemstatus field if non-nil, zero value otherwise.

### GetItemstatusOk

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetItemstatusOk() (*string, bool)`

GetItemstatusOk returns a tuple with the Itemstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemstatus

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetItemstatus(v string)`

SetItemstatus sets Itemstatus field to given value.

### HasItemstatus

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasItemstatus() bool`

HasItemstatus returns a boolean if a field has been set.

### GetStatusmessage

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetStatusmessage() string`

GetStatusmessage returns the Statusmessage field if non-nil, zero value otherwise.

### GetStatusmessageOk

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetStatusmessageOk() (*string, bool)`

GetStatusmessageOk returns a tuple with the Statusmessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusmessage

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetStatusmessage(v string)`

SetStatusmessage sets Statusmessage field to given value.

### HasStatusmessage

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasStatusmessage() bool`

HasStatusmessage returns a boolean if a field has been set.

### GetIngrampartnumber

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetIngrampartnumber() string`

GetIngrampartnumber returns the Ingrampartnumber field if non-nil, zero value otherwise.

### GetIngrampartnumberOk

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetIngrampartnumberOk() (*string, bool)`

GetIngrampartnumberOk returns a tuple with the Ingrampartnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngrampartnumber

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetIngrampartnumber(v string)`

SetIngrampartnumber sets Ingrampartnumber field to given value.

### HasIngrampartnumber

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasIngrampartnumber() bool`

HasIngrampartnumber returns a boolean if a field has been set.

### GetVendorpartnumber

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetVendorpartnumber() string`

GetVendorpartnumber returns the Vendorpartnumber field if non-nil, zero value otherwise.

### GetVendorpartnumberOk

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetVendorpartnumberOk() (*string, bool)`

GetVendorpartnumberOk returns a tuple with the Vendorpartnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorpartnumber

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetVendorpartnumber(v string)`

SetVendorpartnumber sets Vendorpartnumber field to given value.

### HasVendorpartnumber

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasVendorpartnumber() bool`

HasVendorpartnumber returns a boolean if a field has been set.

### GetGlobalskuid

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetGlobalskuid() string`

GetGlobalskuid returns the Globalskuid field if non-nil, zero value otherwise.

### GetGlobalskuidOk

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetGlobalskuidOk() (*string, bool)`

GetGlobalskuidOk returns a tuple with the Globalskuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalskuid

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetGlobalskuid(v string)`

SetGlobalskuid sets Globalskuid field to given value.

### HasGlobalskuid

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasGlobalskuid() bool`

HasGlobalskuid returns a boolean if a field has been set.

### GetCustomerprice

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetCustomerprice() float32`

GetCustomerprice returns the Customerprice field if non-nil, zero value otherwise.

### GetCustomerpriceOk

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetCustomerpriceOk() (*float32, bool)`

GetCustomerpriceOk returns a tuple with the Customerprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerprice

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetCustomerprice(v float32)`

SetCustomerprice sets Customerprice field to given value.

### HasCustomerprice

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasCustomerprice() bool`

HasCustomerprice returns a boolean if a field has been set.

### GetPartdescription1

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetPartdescription1() string`

GetPartdescription1 returns the Partdescription1 field if non-nil, zero value otherwise.

### GetPartdescription1Ok

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetPartdescription1Ok() (*string, bool)`

GetPartdescription1Ok returns a tuple with the Partdescription1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartdescription1

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetPartdescription1(v string)`

SetPartdescription1 sets Partdescription1 field to given value.

### HasPartdescription1

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasPartdescription1() bool`

HasPartdescription1 returns a boolean if a field has been set.

### GetPartdescription2

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetPartdescription2() string`

GetPartdescription2 returns the Partdescription2 field if non-nil, zero value otherwise.

### GetPartdescription2Ok

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetPartdescription2Ok() (*string, bool)`

GetPartdescription2Ok returns a tuple with the Partdescription2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartdescription2

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetPartdescription2(v string)`

SetPartdescription2 sets Partdescription2 field to given value.

### HasPartdescription2

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasPartdescription2() bool`

HasPartdescription2 returns a boolean if a field has been set.

### GetVendornumber

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetVendornumber() string`

GetVendornumber returns the Vendornumber field if non-nil, zero value otherwise.

### GetVendornumberOk

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetVendornumberOk() (*string, bool)`

GetVendornumberOk returns a tuple with the Vendornumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendornumber

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetVendornumber(v string)`

SetVendornumber sets Vendornumber field to given value.

### HasVendornumber

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasVendornumber() bool`

HasVendornumber returns a boolean if a field has been set.

### GetVendorname

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetVendorname() string`

GetVendorname returns the Vendorname field if non-nil, zero value otherwise.

### GetVendornameOk

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetVendornameOk() (*string, bool)`

GetVendornameOk returns a tuple with the Vendorname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorname

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetVendorname(v string)`

SetVendorname sets Vendorname field to given value.

### HasVendorname

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasVendorname() bool`

HasVendorname returns a boolean if a field has been set.

### GetCpucode

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetCpucode() string`

GetCpucode returns the Cpucode field if non-nil, zero value otherwise.

### GetCpucodeOk

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetCpucodeOk() (*string, bool)`

GetCpucodeOk returns a tuple with the Cpucode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpucode

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetCpucode(v string)`

SetCpucode sets Cpucode field to given value.

### HasCpucode

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasCpucode() bool`

HasCpucode returns a boolean if a field has been set.

### GetClass

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetSkustatus

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetSkustatus() string`

GetSkustatus returns the Skustatus field if non-nil, zero value otherwise.

### GetSkustatusOk

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetSkustatusOk() (*string, bool)`

GetSkustatusOk returns a tuple with the Skustatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkustatus

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetSkustatus(v string)`

SetSkustatus sets Skustatus field to given value.

### HasSkustatus

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasSkustatus() bool`

HasSkustatus returns a boolean if a field has been set.

### GetMediacpu

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetMediacpu() string`

GetMediacpu returns the Mediacpu field if non-nil, zero value otherwise.

### GetMediacpuOk

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetMediacpuOk() (*string, bool)`

GetMediacpuOk returns a tuple with the Mediacpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediacpu

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetMediacpu(v string)`

SetMediacpu sets Mediacpu field to given value.

### HasMediacpu

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasMediacpu() bool`

HasMediacpu returns a boolean if a field has been set.

### GetCategorysubcategory

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetCategorysubcategory() string`

GetCategorysubcategory returns the Categorysubcategory field if non-nil, zero value otherwise.

### GetCategorysubcategoryOk

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetCategorysubcategoryOk() (*string, bool)`

GetCategorysubcategoryOk returns a tuple with the Categorysubcategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorysubcategory

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetCategorysubcategory(v string)`

SetCategorysubcategory sets Categorysubcategory field to given value.

### HasCategorysubcategory

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasCategorysubcategory() bool`

HasCategorysubcategory returns a boolean if a field has been set.

### GetRetailprice

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetRetailprice() float32`

GetRetailprice returns the Retailprice field if non-nil, zero value otherwise.

### GetRetailpriceOk

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetRetailpriceOk() (*float32, bool)`

GetRetailpriceOk returns a tuple with the Retailprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailprice

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetRetailprice(v float32)`

SetRetailprice sets Retailprice field to given value.

### HasRetailprice

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasRetailprice() bool`

HasRetailprice returns a boolean if a field has been set.

### GetNewmedia

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetNewmedia() string`

GetNewmedia returns the Newmedia field if non-nil, zero value otherwise.

### GetNewmediaOk

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetNewmediaOk() (*string, bool)`

GetNewmediaOk returns a tuple with the Newmedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewmedia

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetNewmedia(v string)`

SetNewmedia sets Newmedia field to given value.

### HasNewmedia

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasNewmedia() bool`

HasNewmedia returns a boolean if a field has been set.

### GetEnduserrequired

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetEnduserrequired() string`

GetEnduserrequired returns the Enduserrequired field if non-nil, zero value otherwise.

### GetEnduserrequiredOk

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetEnduserrequiredOk() (*string, bool)`

GetEnduserrequiredOk returns a tuple with the Enduserrequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnduserrequired

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetEnduserrequired(v string)`

SetEnduserrequired sets Enduserrequired field to given value.

### HasEnduserrequired

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasEnduserrequired() bool`

HasEnduserrequired returns a boolean if a field has been set.

### GetBackorderflag

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetBackorderflag() string`

GetBackorderflag returns the Backorderflag field if non-nil, zero value otherwise.

### GetBackorderflagOk

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetBackorderflagOk() (*string, bool)`

GetBackorderflagOk returns a tuple with the Backorderflag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackorderflag

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetBackorderflag(v string)`

SetBackorderflag sets Backorderflag field to given value.

### HasBackorderflag

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasBackorderflag() bool`

HasBackorderflag returns a boolean if a field has been set.

### GetSkuauthorized

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetSkuauthorized() string`

GetSkuauthorized returns the Skuauthorized field if non-nil, zero value otherwise.

### GetSkuauthorizedOk

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetSkuauthorizedOk() (*string, bool)`

GetSkuauthorizedOk returns a tuple with the Skuauthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuauthorized

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetSkuauthorized(v string)`

SetSkuauthorized sets Skuauthorized field to given value.

### HasSkuauthorized

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasSkuauthorized() bool`

HasSkuauthorized returns a boolean if a field has been set.

### GetExtendedvendorpartnumber

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetExtendedvendorpartnumber() string`

GetExtendedvendorpartnumber returns the Extendedvendorpartnumber field if non-nil, zero value otherwise.

### GetExtendedvendorpartnumberOk

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetExtendedvendorpartnumberOk() (*string, bool)`

GetExtendedvendorpartnumberOk returns a tuple with the Extendedvendorpartnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedvendorpartnumber

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetExtendedvendorpartnumber(v string)`

SetExtendedvendorpartnumber sets Extendedvendorpartnumber field to given value.

### HasExtendedvendorpartnumber

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasExtendedvendorpartnumber() bool`

HasExtendedvendorpartnumber returns a boolean if a field has been set.

### GetWarehousedetails

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetWarehousedetails() []WarehouseListType`

GetWarehousedetails returns the Warehousedetails field if non-nil, zero value otherwise.

### GetWarehousedetailsOk

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) GetWarehousedetailsOk() (*[]WarehouseListType, bool)`

GetWarehousedetailsOk returns a tuple with the Warehousedetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehousedetails

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) SetWarehousedetails(v []WarehouseListType)`

SetWarehousedetails sets Warehousedetails field to given value.

### HasWarehousedetails

`func (o *PriceAndAvailabilityResponseServiceresponsePriceandstockresponseDetailsInner) HasWarehousedetails() bool`

HasWarehousedetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


