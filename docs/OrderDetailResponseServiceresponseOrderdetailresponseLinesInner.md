# OrderDetailResponseServiceresponseOrderdetailresponseLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Linenumber** | Pointer to **string** | Impulse line number | [optional] 
**Globallinenumber** | Pointer to **string** | Line of the Globel Sku / Customer Line Number | [optional] 
**Ordersuffix** | Pointer to **string** | Order Suffix | [optional] 
**Erpordernumber** | Pointer to **string** | Sales order number | [optional] 
**Linestatus** | Pointer to **string** | Status of the line | [optional] 
**Partnumber** | Pointer to **string** | Ingram part number | [optional] 
**Manufacturerpartnumber** | Pointer to **string** | manufacture number of the product | [optional] 
**Vendorname** | Pointer to **string** | name of the vendor | [optional] 
**Vendorcode** | Pointer to **string** | Ingram Micro assigned code for the vendor | [optional] 
**Partdescription1** | Pointer to **string** |  | [optional] 
**Partdescription2** | Pointer to **string** |  | [optional] 
**Unitweight** | Pointer to **string** | weight of the product unit | [optional] 
**Unitprice** | Pointer to **float32** | Customer price of the unit | [optional] 
**Extendedprice** | Pointer to **float32** | extended price of the order | [optional] 
**Taxamount** | Pointer to **float32** | tax amount for the order | [optional] 
**Requestedquantity** | Pointer to **string** | no. of units requested | [optional] 
**Confirmedquantity** | Pointer to **string** | no. of units confirmed available | [optional] 
**Backorderquantity** | Pointer to **string** | quantity of back order | [optional] 
**Serialnumberdetails** | Pointer to [**[]OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerSerialnumberdetailsInner**](OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerSerialnumberdetailsInner.md) |  | [optional] 
**Trackingnumber** | Pointer to **[]string** |  | [optional] 
**Shipmentdetails** | Pointer to [**[]OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner**](OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner.md) |  | [optional] 
**Productextendedspecs** | Pointer to [**[]InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner**](InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner.md) |  | [optional] 
**Backorderetadate** | Pointer to **string** | estimated date of back order | [optional] 

## Methods

### NewOrderDetailResponseServiceresponseOrderdetailresponseLinesInner

`func NewOrderDetailResponseServiceresponseOrderdetailresponseLinesInner() *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner`

NewOrderDetailResponseServiceresponseOrderdetailresponseLinesInner instantiates a new OrderDetailResponseServiceresponseOrderdetailresponseLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDetailResponseServiceresponseOrderdetailresponseLinesInnerWithDefaults

`func NewOrderDetailResponseServiceresponseOrderdetailresponseLinesInnerWithDefaults() *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner`

NewOrderDetailResponseServiceresponseOrderdetailresponseLinesInnerWithDefaults instantiates a new OrderDetailResponseServiceresponseOrderdetailresponseLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinenumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetLinenumber() string`

GetLinenumber returns the Linenumber field if non-nil, zero value otherwise.

### GetLinenumberOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetLinenumberOk() (*string, bool)`

GetLinenumberOk returns a tuple with the Linenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinenumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetLinenumber(v string)`

SetLinenumber sets Linenumber field to given value.

### HasLinenumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasLinenumber() bool`

HasLinenumber returns a boolean if a field has been set.

### GetGloballinenumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetGloballinenumber() string`

GetGloballinenumber returns the Globallinenumber field if non-nil, zero value otherwise.

### GetGloballinenumberOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetGloballinenumberOk() (*string, bool)`

GetGloballinenumberOk returns a tuple with the Globallinenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGloballinenumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetGloballinenumber(v string)`

SetGloballinenumber sets Globallinenumber field to given value.

### HasGloballinenumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasGloballinenumber() bool`

HasGloballinenumber returns a boolean if a field has been set.

### GetOrdersuffix

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetOrdersuffix() string`

GetOrdersuffix returns the Ordersuffix field if non-nil, zero value otherwise.

### GetOrdersuffixOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetOrdersuffixOk() (*string, bool)`

GetOrdersuffixOk returns a tuple with the Ordersuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersuffix

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetOrdersuffix(v string)`

SetOrdersuffix sets Ordersuffix field to given value.

### HasOrdersuffix

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasOrdersuffix() bool`

HasOrdersuffix returns a boolean if a field has been set.

### GetErpordernumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetErpordernumber() string`

GetErpordernumber returns the Erpordernumber field if non-nil, zero value otherwise.

### GetErpordernumberOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetErpordernumberOk() (*string, bool)`

GetErpordernumberOk returns a tuple with the Erpordernumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErpordernumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetErpordernumber(v string)`

SetErpordernumber sets Erpordernumber field to given value.

### HasErpordernumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasErpordernumber() bool`

HasErpordernumber returns a boolean if a field has been set.

### GetLinestatus

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetLinestatus() string`

GetLinestatus returns the Linestatus field if non-nil, zero value otherwise.

### GetLinestatusOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetLinestatusOk() (*string, bool)`

GetLinestatusOk returns a tuple with the Linestatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinestatus

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetLinestatus(v string)`

SetLinestatus sets Linestatus field to given value.

### HasLinestatus

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasLinestatus() bool`

HasLinestatus returns a boolean if a field has been set.

### GetPartnumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetPartnumber() string`

GetPartnumber returns the Partnumber field if non-nil, zero value otherwise.

### GetPartnumberOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetPartnumberOk() (*string, bool)`

GetPartnumberOk returns a tuple with the Partnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetPartnumber(v string)`

SetPartnumber sets Partnumber field to given value.

### HasPartnumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasPartnumber() bool`

HasPartnumber returns a boolean if a field has been set.

### GetManufacturerpartnumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetManufacturerpartnumber() string`

GetManufacturerpartnumber returns the Manufacturerpartnumber field if non-nil, zero value otherwise.

### GetManufacturerpartnumberOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetManufacturerpartnumberOk() (*string, bool)`

GetManufacturerpartnumberOk returns a tuple with the Manufacturerpartnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerpartnumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetManufacturerpartnumber(v string)`

SetManufacturerpartnumber sets Manufacturerpartnumber field to given value.

### HasManufacturerpartnumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasManufacturerpartnumber() bool`

HasManufacturerpartnumber returns a boolean if a field has been set.

### GetVendorname

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetVendorname() string`

GetVendorname returns the Vendorname field if non-nil, zero value otherwise.

### GetVendornameOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetVendornameOk() (*string, bool)`

GetVendornameOk returns a tuple with the Vendorname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorname

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetVendorname(v string)`

SetVendorname sets Vendorname field to given value.

### HasVendorname

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasVendorname() bool`

HasVendorname returns a boolean if a field has been set.

### GetVendorcode

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetVendorcode() string`

GetVendorcode returns the Vendorcode field if non-nil, zero value otherwise.

### GetVendorcodeOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetVendorcodeOk() (*string, bool)`

GetVendorcodeOk returns a tuple with the Vendorcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorcode

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetVendorcode(v string)`

SetVendorcode sets Vendorcode field to given value.

### HasVendorcode

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasVendorcode() bool`

HasVendorcode returns a boolean if a field has been set.

### GetPartdescription1

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetPartdescription1() string`

GetPartdescription1 returns the Partdescription1 field if non-nil, zero value otherwise.

### GetPartdescription1Ok

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetPartdescription1Ok() (*string, bool)`

GetPartdescription1Ok returns a tuple with the Partdescription1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartdescription1

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetPartdescription1(v string)`

SetPartdescription1 sets Partdescription1 field to given value.

### HasPartdescription1

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasPartdescription1() bool`

HasPartdescription1 returns a boolean if a field has been set.

### GetPartdescription2

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetPartdescription2() string`

GetPartdescription2 returns the Partdescription2 field if non-nil, zero value otherwise.

### GetPartdescription2Ok

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetPartdescription2Ok() (*string, bool)`

GetPartdescription2Ok returns a tuple with the Partdescription2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartdescription2

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetPartdescription2(v string)`

SetPartdescription2 sets Partdescription2 field to given value.

### HasPartdescription2

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasPartdescription2() bool`

HasPartdescription2 returns a boolean if a field has been set.

### GetUnitweight

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetUnitweight() string`

GetUnitweight returns the Unitweight field if non-nil, zero value otherwise.

### GetUnitweightOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetUnitweightOk() (*string, bool)`

GetUnitweightOk returns a tuple with the Unitweight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitweight

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetUnitweight(v string)`

SetUnitweight sets Unitweight field to given value.

### HasUnitweight

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasUnitweight() bool`

HasUnitweight returns a boolean if a field has been set.

### GetUnitprice

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetUnitprice() float32`

GetUnitprice returns the Unitprice field if non-nil, zero value otherwise.

### GetUnitpriceOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetUnitpriceOk() (*float32, bool)`

GetUnitpriceOk returns a tuple with the Unitprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitprice

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetUnitprice(v float32)`

SetUnitprice sets Unitprice field to given value.

### HasUnitprice

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasUnitprice() bool`

HasUnitprice returns a boolean if a field has been set.

### GetExtendedprice

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetExtendedprice() float32`

GetExtendedprice returns the Extendedprice field if non-nil, zero value otherwise.

### GetExtendedpriceOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetExtendedpriceOk() (*float32, bool)`

GetExtendedpriceOk returns a tuple with the Extendedprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedprice

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetExtendedprice(v float32)`

SetExtendedprice sets Extendedprice field to given value.

### HasExtendedprice

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasExtendedprice() bool`

HasExtendedprice returns a boolean if a field has been set.

### GetTaxamount

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetTaxamount() float32`

GetTaxamount returns the Taxamount field if non-nil, zero value otherwise.

### GetTaxamountOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetTaxamountOk() (*float32, bool)`

GetTaxamountOk returns a tuple with the Taxamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxamount

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetTaxamount(v float32)`

SetTaxamount sets Taxamount field to given value.

### HasTaxamount

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasTaxamount() bool`

HasTaxamount returns a boolean if a field has been set.

### GetRequestedquantity

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetRequestedquantity() string`

GetRequestedquantity returns the Requestedquantity field if non-nil, zero value otherwise.

### GetRequestedquantityOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetRequestedquantityOk() (*string, bool)`

GetRequestedquantityOk returns a tuple with the Requestedquantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedquantity

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetRequestedquantity(v string)`

SetRequestedquantity sets Requestedquantity field to given value.

### HasRequestedquantity

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasRequestedquantity() bool`

HasRequestedquantity returns a boolean if a field has been set.

### GetConfirmedquantity

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetConfirmedquantity() string`

GetConfirmedquantity returns the Confirmedquantity field if non-nil, zero value otherwise.

### GetConfirmedquantityOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetConfirmedquantityOk() (*string, bool)`

GetConfirmedquantityOk returns a tuple with the Confirmedquantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedquantity

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetConfirmedquantity(v string)`

SetConfirmedquantity sets Confirmedquantity field to given value.

### HasConfirmedquantity

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasConfirmedquantity() bool`

HasConfirmedquantity returns a boolean if a field has been set.

### GetBackorderquantity

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetBackorderquantity() string`

GetBackorderquantity returns the Backorderquantity field if non-nil, zero value otherwise.

### GetBackorderquantityOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetBackorderquantityOk() (*string, bool)`

GetBackorderquantityOk returns a tuple with the Backorderquantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackorderquantity

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetBackorderquantity(v string)`

SetBackorderquantity sets Backorderquantity field to given value.

### HasBackorderquantity

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasBackorderquantity() bool`

HasBackorderquantity returns a boolean if a field has been set.

### GetSerialnumberdetails

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetSerialnumberdetails() []OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerSerialnumberdetailsInner`

GetSerialnumberdetails returns the Serialnumberdetails field if non-nil, zero value otherwise.

### GetSerialnumberdetailsOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetSerialnumberdetailsOk() (*[]OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerSerialnumberdetailsInner, bool)`

GetSerialnumberdetailsOk returns a tuple with the Serialnumberdetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialnumberdetails

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetSerialnumberdetails(v []OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerSerialnumberdetailsInner)`

SetSerialnumberdetails sets Serialnumberdetails field to given value.

### HasSerialnumberdetails

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasSerialnumberdetails() bool`

HasSerialnumberdetails returns a boolean if a field has been set.

### GetTrackingnumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetTrackingnumber() []string`

GetTrackingnumber returns the Trackingnumber field if non-nil, zero value otherwise.

### GetTrackingnumberOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetTrackingnumberOk() (*[]string, bool)`

GetTrackingnumberOk returns a tuple with the Trackingnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingnumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetTrackingnumber(v []string)`

SetTrackingnumber sets Trackingnumber field to given value.

### HasTrackingnumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasTrackingnumber() bool`

HasTrackingnumber returns a boolean if a field has been set.

### GetShipmentdetails

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetShipmentdetails() []OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner`

GetShipmentdetails returns the Shipmentdetails field if non-nil, zero value otherwise.

### GetShipmentdetailsOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetShipmentdetailsOk() (*[]OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner, bool)`

GetShipmentdetailsOk returns a tuple with the Shipmentdetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentdetails

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetShipmentdetails(v []OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner)`

SetShipmentdetails sets Shipmentdetails field to given value.

### HasShipmentdetails

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasShipmentdetails() bool`

HasShipmentdetails returns a boolean if a field has been set.

### GetProductextendedspecs

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetProductextendedspecs() []InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner`

GetProductextendedspecs returns the Productextendedspecs field if non-nil, zero value otherwise.

### GetProductextendedspecsOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetProductextendedspecsOk() (*[]InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner, bool)`

GetProductextendedspecsOk returns a tuple with the Productextendedspecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductextendedspecs

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetProductextendedspecs(v []InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner)`

SetProductextendedspecs sets Productextendedspecs field to given value.

### HasProductextendedspecs

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasProductextendedspecs() bool`

HasProductextendedspecs returns a boolean if a field has been set.

### GetBackorderetadate

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetBackorderetadate() string`

GetBackorderetadate returns the Backorderetadate field if non-nil, zero value otherwise.

### GetBackorderetadateOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetBackorderetadateOk() (*string, bool)`

GetBackorderetadateOk returns a tuple with the Backorderetadate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackorderetadate

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetBackorderetadate(v string)`

SetBackorderetadate sets Backorderetadate field to given value.

### HasBackorderetadate

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasBackorderetadate() bool`

HasBackorderetadate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


