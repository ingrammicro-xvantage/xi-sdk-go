# OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **float32** | quantity shipped | [optional] 
**Shipmentdate** | Pointer to **string** | date of shipment | [optional] 
**Shipfromwarehouseid** | Pointer to **string** | Warehouse product was shipped from | [optional] 
**Warehousename** | Pointer to **string** | name of the warehouse | [optional] 
**Invoicenumber** | Pointer to **string** | Invoice Number | [optional] 
**Invoicedate** | Pointer to **string** | date on the invoice generated | [optional] 
**Status** | Pointer to **string** | code for current Status of the order | [optional] 
**Statusdescription** | Pointer to **string** | Description of status | [optional] 
**Shippeddate** | Pointer to **string** | date of shipment | [optional] 
**Holdreasoncodedescription** | Pointer to **string** | Description of the code if the order is on hold | [optional] 
**Ponumber** | Pointer to **string** | Ingram PO Number to vendors for direct ship orders | [optional] 
**Carriertype** | Pointer to **string** | Helps to determine shipment type. for e.g. LTL is used for heavy shipment. SML is used for light shipment | [optional] 
**Carriercode** | Pointer to **string** |  | [optional] 
**Carriername** | Pointer to **string** | Name of the carrier. If carriername is LTL then the tracking info is in the \&quot;pronumber\&quot; data field | [optional] 
**Pronumber** | Pointer to **string** |  | [optional] 
**Packagedetails** | Pointer to [**OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInnerPackagedetails**](OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInnerPackagedetails.md) |  | [optional] 

## Methods

### NewOrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner

`func NewOrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner() *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner`

NewOrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner instantiates a new OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInnerWithDefaults

`func NewOrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInnerWithDefaults() *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner`

NewOrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInnerWithDefaults instantiates a new OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetShipmentdate

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetShipmentdate() string`

GetShipmentdate returns the Shipmentdate field if non-nil, zero value otherwise.

### GetShipmentdateOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetShipmentdateOk() (*string, bool)`

GetShipmentdateOk returns a tuple with the Shipmentdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentdate

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) SetShipmentdate(v string)`

SetShipmentdate sets Shipmentdate field to given value.

### HasShipmentdate

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) HasShipmentdate() bool`

HasShipmentdate returns a boolean if a field has been set.

### GetShipfromwarehouseid

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetShipfromwarehouseid() string`

GetShipfromwarehouseid returns the Shipfromwarehouseid field if non-nil, zero value otherwise.

### GetShipfromwarehouseidOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetShipfromwarehouseidOk() (*string, bool)`

GetShipfromwarehouseidOk returns a tuple with the Shipfromwarehouseid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipfromwarehouseid

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) SetShipfromwarehouseid(v string)`

SetShipfromwarehouseid sets Shipfromwarehouseid field to given value.

### HasShipfromwarehouseid

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) HasShipfromwarehouseid() bool`

HasShipfromwarehouseid returns a boolean if a field has been set.

### GetWarehousename

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetWarehousename() string`

GetWarehousename returns the Warehousename field if non-nil, zero value otherwise.

### GetWarehousenameOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetWarehousenameOk() (*string, bool)`

GetWarehousenameOk returns a tuple with the Warehousename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehousename

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) SetWarehousename(v string)`

SetWarehousename sets Warehousename field to given value.

### HasWarehousename

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) HasWarehousename() bool`

HasWarehousename returns a boolean if a field has been set.

### GetInvoicenumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetInvoicenumber() string`

GetInvoicenumber returns the Invoicenumber field if non-nil, zero value otherwise.

### GetInvoicenumberOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetInvoicenumberOk() (*string, bool)`

GetInvoicenumberOk returns a tuple with the Invoicenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicenumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) SetInvoicenumber(v string)`

SetInvoicenumber sets Invoicenumber field to given value.

### HasInvoicenumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) HasInvoicenumber() bool`

HasInvoicenumber returns a boolean if a field has been set.

### GetInvoicedate

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetInvoicedate() string`

GetInvoicedate returns the Invoicedate field if non-nil, zero value otherwise.

### GetInvoicedateOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetInvoicedateOk() (*string, bool)`

GetInvoicedateOk returns a tuple with the Invoicedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicedate

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) SetInvoicedate(v string)`

SetInvoicedate sets Invoicedate field to given value.

### HasInvoicedate

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) HasInvoicedate() bool`

HasInvoicedate returns a boolean if a field has been set.

### GetStatus

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusdescription

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetStatusdescription() string`

GetStatusdescription returns the Statusdescription field if non-nil, zero value otherwise.

### GetStatusdescriptionOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetStatusdescriptionOk() (*string, bool)`

GetStatusdescriptionOk returns a tuple with the Statusdescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusdescription

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) SetStatusdescription(v string)`

SetStatusdescription sets Statusdescription field to given value.

### HasStatusdescription

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) HasStatusdescription() bool`

HasStatusdescription returns a boolean if a field has been set.

### GetShippeddate

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetShippeddate() string`

GetShippeddate returns the Shippeddate field if non-nil, zero value otherwise.

### GetShippeddateOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetShippeddateOk() (*string, bool)`

GetShippeddateOk returns a tuple with the Shippeddate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippeddate

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) SetShippeddate(v string)`

SetShippeddate sets Shippeddate field to given value.

### HasShippeddate

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) HasShippeddate() bool`

HasShippeddate returns a boolean if a field has been set.

### GetHoldreasoncodedescription

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetHoldreasoncodedescription() string`

GetHoldreasoncodedescription returns the Holdreasoncodedescription field if non-nil, zero value otherwise.

### GetHoldreasoncodedescriptionOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetHoldreasoncodedescriptionOk() (*string, bool)`

GetHoldreasoncodedescriptionOk returns a tuple with the Holdreasoncodedescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldreasoncodedescription

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) SetHoldreasoncodedescription(v string)`

SetHoldreasoncodedescription sets Holdreasoncodedescription field to given value.

### HasHoldreasoncodedescription

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) HasHoldreasoncodedescription() bool`

HasHoldreasoncodedescription returns a boolean if a field has been set.

### GetPonumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetPonumber() string`

GetPonumber returns the Ponumber field if non-nil, zero value otherwise.

### GetPonumberOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetPonumberOk() (*string, bool)`

GetPonumberOk returns a tuple with the Ponumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPonumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) SetPonumber(v string)`

SetPonumber sets Ponumber field to given value.

### HasPonumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) HasPonumber() bool`

HasPonumber returns a boolean if a field has been set.

### GetCarriertype

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetCarriertype() string`

GetCarriertype returns the Carriertype field if non-nil, zero value otherwise.

### GetCarriertypeOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetCarriertypeOk() (*string, bool)`

GetCarriertypeOk returns a tuple with the Carriertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarriertype

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) SetCarriertype(v string)`

SetCarriertype sets Carriertype field to given value.

### HasCarriertype

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) HasCarriertype() bool`

HasCarriertype returns a boolean if a field has been set.

### GetCarriercode

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetCarriercode() string`

GetCarriercode returns the Carriercode field if non-nil, zero value otherwise.

### GetCarriercodeOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetCarriercodeOk() (*string, bool)`

GetCarriercodeOk returns a tuple with the Carriercode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarriercode

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) SetCarriercode(v string)`

SetCarriercode sets Carriercode field to given value.

### HasCarriercode

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) HasCarriercode() bool`

HasCarriercode returns a boolean if a field has been set.

### GetCarriername

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetCarriername() string`

GetCarriername returns the Carriername field if non-nil, zero value otherwise.

### GetCarriernameOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetCarriernameOk() (*string, bool)`

GetCarriernameOk returns a tuple with the Carriername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarriername

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) SetCarriername(v string)`

SetCarriername sets Carriername field to given value.

### HasCarriername

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) HasCarriername() bool`

HasCarriername returns a boolean if a field has been set.

### GetPronumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetPronumber() string`

GetPronumber returns the Pronumber field if non-nil, zero value otherwise.

### GetPronumberOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetPronumberOk() (*string, bool)`

GetPronumberOk returns a tuple with the Pronumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPronumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) SetPronumber(v string)`

SetPronumber sets Pronumber field to given value.

### HasPronumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) HasPronumber() bool`

HasPronumber returns a boolean if a field has been set.

### GetPackagedetails

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetPackagedetails() OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInnerPackagedetails`

GetPackagedetails returns the Packagedetails field if non-nil, zero value otherwise.

### GetPackagedetailsOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) GetPackagedetailsOk() (*OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInnerPackagedetails, bool)`

GetPackagedetailsOk returns a tuple with the Packagedetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagedetails

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) SetPackagedetails(v OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInnerPackagedetails)`

SetPackagedetails sets Packagedetails field to given value.

### HasPackagedetails

`func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) HasPackagedetails() bool`

HasPackagedetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


