# MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Itemstatus** | Pointer to **string** |  | [optional] 
**Statusmessage** | Pointer to **string** |  | [optional] 
**Ingrampartnumber** | Pointer to **string** | SKU number for the product for which order needs to be created with Ingram Micro | [optional] 
**Vendorpartnumber** | Pointer to **string** | Vendor Part number for the product | [optional] 
**Globalskuid** | Pointer to **string** |  | [optional] 
**Customerprice** | Pointer to **string** | Customer specific price for the product, excluding taxes | [optional] 
**Partdescription1** | Pointer to **string** | Description on the part number that is being requested | [optional] 
**Partdescription2** | Pointer to **string** | Contuiation of description on the part number that is being requested | [optional] 
**Vendornumber** | Pointer to **string** | Internal four digit code assigned by Ingram | [optional] 
**Vendorname** | Pointer to **string** | Name of the vendor | [optional] 
**Cpucode** | Pointer to **string** | Ingram internal code for a product | [optional] 
**Class** | Pointer to **string** | Ingram Micro assigned product classification. | [optional] 
**Skustatus** | Pointer to **string** | Identifies if the SKU has been discontinued. Rules must be defined on the values to be sent out to partner. | [optional] 
**Mediacpu** | Pointer to **string** |  | [optional] 
**Categorysubcategory** | Pointer to **string** | Ingram&#39;s internal categorization of the product | [optional] 
**Retailprice** | **float32** | MSRP Price 0.00 | 
**Newmedia** | Pointer to **string** | Internal four-digit code assigned by Ingram to represent the item group | [optional] 
**Enduserrequired** | Pointer to **string** | Y - End user required N - Not required End user | [optional] 
**Backorderflag** | Pointer to **string** | Y- Allow Backorder Flag N- Not allowed | [optional] 
**Skuauthorized** | Pointer to **string** |  | [optional] 
**Extendedvendorpartnumber** | Pointer to **string** |  | [optional] 
**Warehousedetails** | Pointer to [**[]MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInnerWarehousedetailsInner**](MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInnerWarehousedetailsInner.md) |  | [optional] 

## Methods

### NewMultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner

`func NewMultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner(retailprice float32, ) *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner`

NewMultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner instantiates a new MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInnerWithDefaults

`func NewMultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInnerWithDefaults() *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner`

NewMultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInnerWithDefaults instantiates a new MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemstatus

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetItemstatus() string`

GetItemstatus returns the Itemstatus field if non-nil, zero value otherwise.

### GetItemstatusOk

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetItemstatusOk() (*string, bool)`

GetItemstatusOk returns a tuple with the Itemstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemstatus

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetItemstatus(v string)`

SetItemstatus sets Itemstatus field to given value.

### HasItemstatus

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasItemstatus() bool`

HasItemstatus returns a boolean if a field has been set.

### GetStatusmessage

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetStatusmessage() string`

GetStatusmessage returns the Statusmessage field if non-nil, zero value otherwise.

### GetStatusmessageOk

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetStatusmessageOk() (*string, bool)`

GetStatusmessageOk returns a tuple with the Statusmessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusmessage

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetStatusmessage(v string)`

SetStatusmessage sets Statusmessage field to given value.

### HasStatusmessage

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasStatusmessage() bool`

HasStatusmessage returns a boolean if a field has been set.

### GetIngrampartnumber

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetIngrampartnumber() string`

GetIngrampartnumber returns the Ingrampartnumber field if non-nil, zero value otherwise.

### GetIngrampartnumberOk

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetIngrampartnumberOk() (*string, bool)`

GetIngrampartnumberOk returns a tuple with the Ingrampartnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngrampartnumber

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetIngrampartnumber(v string)`

SetIngrampartnumber sets Ingrampartnumber field to given value.

### HasIngrampartnumber

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasIngrampartnumber() bool`

HasIngrampartnumber returns a boolean if a field has been set.

### GetVendorpartnumber

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetVendorpartnumber() string`

GetVendorpartnumber returns the Vendorpartnumber field if non-nil, zero value otherwise.

### GetVendorpartnumberOk

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetVendorpartnumberOk() (*string, bool)`

GetVendorpartnumberOk returns a tuple with the Vendorpartnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorpartnumber

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetVendorpartnumber(v string)`

SetVendorpartnumber sets Vendorpartnumber field to given value.

### HasVendorpartnumber

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasVendorpartnumber() bool`

HasVendorpartnumber returns a boolean if a field has been set.

### GetGlobalskuid

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetGlobalskuid() string`

GetGlobalskuid returns the Globalskuid field if non-nil, zero value otherwise.

### GetGlobalskuidOk

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetGlobalskuidOk() (*string, bool)`

GetGlobalskuidOk returns a tuple with the Globalskuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalskuid

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetGlobalskuid(v string)`

SetGlobalskuid sets Globalskuid field to given value.

### HasGlobalskuid

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasGlobalskuid() bool`

HasGlobalskuid returns a boolean if a field has been set.

### GetCustomerprice

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetCustomerprice() string`

GetCustomerprice returns the Customerprice field if non-nil, zero value otherwise.

### GetCustomerpriceOk

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetCustomerpriceOk() (*string, bool)`

GetCustomerpriceOk returns a tuple with the Customerprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerprice

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetCustomerprice(v string)`

SetCustomerprice sets Customerprice field to given value.

### HasCustomerprice

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasCustomerprice() bool`

HasCustomerprice returns a boolean if a field has been set.

### GetPartdescription1

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetPartdescription1() string`

GetPartdescription1 returns the Partdescription1 field if non-nil, zero value otherwise.

### GetPartdescription1Ok

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetPartdescription1Ok() (*string, bool)`

GetPartdescription1Ok returns a tuple with the Partdescription1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartdescription1

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetPartdescription1(v string)`

SetPartdescription1 sets Partdescription1 field to given value.

### HasPartdescription1

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasPartdescription1() bool`

HasPartdescription1 returns a boolean if a field has been set.

### GetPartdescription2

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetPartdescription2() string`

GetPartdescription2 returns the Partdescription2 field if non-nil, zero value otherwise.

### GetPartdescription2Ok

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetPartdescription2Ok() (*string, bool)`

GetPartdescription2Ok returns a tuple with the Partdescription2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartdescription2

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetPartdescription2(v string)`

SetPartdescription2 sets Partdescription2 field to given value.

### HasPartdescription2

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasPartdescription2() bool`

HasPartdescription2 returns a boolean if a field has been set.

### GetVendornumber

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetVendornumber() string`

GetVendornumber returns the Vendornumber field if non-nil, zero value otherwise.

### GetVendornumberOk

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetVendornumberOk() (*string, bool)`

GetVendornumberOk returns a tuple with the Vendornumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendornumber

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetVendornumber(v string)`

SetVendornumber sets Vendornumber field to given value.

### HasVendornumber

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasVendornumber() bool`

HasVendornumber returns a boolean if a field has been set.

### GetVendorname

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetVendorname() string`

GetVendorname returns the Vendorname field if non-nil, zero value otherwise.

### GetVendornameOk

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetVendornameOk() (*string, bool)`

GetVendornameOk returns a tuple with the Vendorname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorname

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetVendorname(v string)`

SetVendorname sets Vendorname field to given value.

### HasVendorname

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasVendorname() bool`

HasVendorname returns a boolean if a field has been set.

### GetCpucode

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetCpucode() string`

GetCpucode returns the Cpucode field if non-nil, zero value otherwise.

### GetCpucodeOk

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetCpucodeOk() (*string, bool)`

GetCpucodeOk returns a tuple with the Cpucode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpucode

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetCpucode(v string)`

SetCpucode sets Cpucode field to given value.

### HasCpucode

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasCpucode() bool`

HasCpucode returns a boolean if a field has been set.

### GetClass

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetSkustatus

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetSkustatus() string`

GetSkustatus returns the Skustatus field if non-nil, zero value otherwise.

### GetSkustatusOk

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetSkustatusOk() (*string, bool)`

GetSkustatusOk returns a tuple with the Skustatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkustatus

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetSkustatus(v string)`

SetSkustatus sets Skustatus field to given value.

### HasSkustatus

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasSkustatus() bool`

HasSkustatus returns a boolean if a field has been set.

### GetMediacpu

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetMediacpu() string`

GetMediacpu returns the Mediacpu field if non-nil, zero value otherwise.

### GetMediacpuOk

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetMediacpuOk() (*string, bool)`

GetMediacpuOk returns a tuple with the Mediacpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediacpu

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetMediacpu(v string)`

SetMediacpu sets Mediacpu field to given value.

### HasMediacpu

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasMediacpu() bool`

HasMediacpu returns a boolean if a field has been set.

### GetCategorysubcategory

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetCategorysubcategory() string`

GetCategorysubcategory returns the Categorysubcategory field if non-nil, zero value otherwise.

### GetCategorysubcategoryOk

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetCategorysubcategoryOk() (*string, bool)`

GetCategorysubcategoryOk returns a tuple with the Categorysubcategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorysubcategory

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetCategorysubcategory(v string)`

SetCategorysubcategory sets Categorysubcategory field to given value.

### HasCategorysubcategory

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasCategorysubcategory() bool`

HasCategorysubcategory returns a boolean if a field has been set.

### GetRetailprice

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetRetailprice() float32`

GetRetailprice returns the Retailprice field if non-nil, zero value otherwise.

### GetRetailpriceOk

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetRetailpriceOk() (*float32, bool)`

GetRetailpriceOk returns a tuple with the Retailprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailprice

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetRetailprice(v float32)`

SetRetailprice sets Retailprice field to given value.


### GetNewmedia

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetNewmedia() string`

GetNewmedia returns the Newmedia field if non-nil, zero value otherwise.

### GetNewmediaOk

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetNewmediaOk() (*string, bool)`

GetNewmediaOk returns a tuple with the Newmedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewmedia

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetNewmedia(v string)`

SetNewmedia sets Newmedia field to given value.

### HasNewmedia

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasNewmedia() bool`

HasNewmedia returns a boolean if a field has been set.

### GetEnduserrequired

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetEnduserrequired() string`

GetEnduserrequired returns the Enduserrequired field if non-nil, zero value otherwise.

### GetEnduserrequiredOk

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetEnduserrequiredOk() (*string, bool)`

GetEnduserrequiredOk returns a tuple with the Enduserrequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnduserrequired

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetEnduserrequired(v string)`

SetEnduserrequired sets Enduserrequired field to given value.

### HasEnduserrequired

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasEnduserrequired() bool`

HasEnduserrequired returns a boolean if a field has been set.

### GetBackorderflag

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetBackorderflag() string`

GetBackorderflag returns the Backorderflag field if non-nil, zero value otherwise.

### GetBackorderflagOk

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetBackorderflagOk() (*string, bool)`

GetBackorderflagOk returns a tuple with the Backorderflag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackorderflag

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetBackorderflag(v string)`

SetBackorderflag sets Backorderflag field to given value.

### HasBackorderflag

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasBackorderflag() bool`

HasBackorderflag returns a boolean if a field has been set.

### GetSkuauthorized

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetSkuauthorized() string`

GetSkuauthorized returns the Skuauthorized field if non-nil, zero value otherwise.

### GetSkuauthorizedOk

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetSkuauthorizedOk() (*string, bool)`

GetSkuauthorizedOk returns a tuple with the Skuauthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuauthorized

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetSkuauthorized(v string)`

SetSkuauthorized sets Skuauthorized field to given value.

### HasSkuauthorized

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasSkuauthorized() bool`

HasSkuauthorized returns a boolean if a field has been set.

### GetExtendedvendorpartnumber

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetExtendedvendorpartnumber() string`

GetExtendedvendorpartnumber returns the Extendedvendorpartnumber field if non-nil, zero value otherwise.

### GetExtendedvendorpartnumberOk

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetExtendedvendorpartnumberOk() (*string, bool)`

GetExtendedvendorpartnumberOk returns a tuple with the Extendedvendorpartnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedvendorpartnumber

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetExtendedvendorpartnumber(v string)`

SetExtendedvendorpartnumber sets Extendedvendorpartnumber field to given value.

### HasExtendedvendorpartnumber

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasExtendedvendorpartnumber() bool`

HasExtendedvendorpartnumber returns a boolean if a field has been set.

### GetWarehousedetails

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetWarehousedetails() []MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInnerWarehousedetailsInner`

GetWarehousedetails returns the Warehousedetails field if non-nil, zero value otherwise.

### GetWarehousedetailsOk

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) GetWarehousedetailsOk() (*[]MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInnerWarehousedetailsInner, bool)`

GetWarehousedetailsOk returns a tuple with the Warehousedetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehousedetails

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) SetWarehousedetails(v []MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInnerWarehousedetailsInner)`

SetWarehousedetails sets Warehousedetails field to given value.

### HasWarehousedetails

`func (o *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponseDetailsInner) HasWarehousedetails() bool`

HasWarehousedetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


