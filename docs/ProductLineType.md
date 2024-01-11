# ProductLineType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Linenumber** | Pointer to **string** |  | [optional] 
**Linetype** | Pointer to **string** |  | [optional] 
**Partnumber** | Pointer to **string** |  | [optional] 
**Vendorpartnumber** | Pointer to **string** |  | [optional] 
**Partdescription** | Pointer to **string** |  | [optional] 
**Shipfrombranch** | Pointer to **string** |  | [optional] 
**Shippedquantity** | Pointer to **string** |  | [optional] 
**Orderedquantity** | Pointer to **string** |  | [optional] 
**Marginpercent** | Pointer to **string** |  | [optional] 
**Backorderquantity** | Pointer to **string** |  | [optional] 
**Backorderetadate** | Pointer to **string** |  | [optional] 
**Extendedprice** | Pointer to **string** |  | [optional] 
**Specialbidnumber** | Pointer to **string** |  | [optional] 
**Ordersuffix** | Pointer to **string** |  | [optional] 
**Isacopapplied** | Pointer to **string** |  | [optional] 
**Unitprice** | Pointer to **string** |  | [optional] 
**Unitofmeasure** | Pointer to **string** |  | [optional] 
**Serialnumberdetails** | Pointer to [**[]ProductLineTypeSerialnumberdetailsInner**](ProductLineTypeSerialnumberdetailsInner.md) |  | [optional] 
**Trackingnumberdetails** | Pointer to [**[]ProductLineTypeTrackingnumberdetailsInner**](ProductLineTypeTrackingnumberdetailsInner.md) |  | [optional] 
**Productextendedspecs** | Pointer to [**[]InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner**](InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner.md) |  | [optional] 

## Methods

### NewProductLineType

`func NewProductLineType() *ProductLineType`

NewProductLineType instantiates a new ProductLineType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductLineTypeWithDefaults

`func NewProductLineTypeWithDefaults() *ProductLineType`

NewProductLineTypeWithDefaults instantiates a new ProductLineType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinenumber

`func (o *ProductLineType) GetLinenumber() string`

GetLinenumber returns the Linenumber field if non-nil, zero value otherwise.

### GetLinenumberOk

`func (o *ProductLineType) GetLinenumberOk() (*string, bool)`

GetLinenumberOk returns a tuple with the Linenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinenumber

`func (o *ProductLineType) SetLinenumber(v string)`

SetLinenumber sets Linenumber field to given value.

### HasLinenumber

`func (o *ProductLineType) HasLinenumber() bool`

HasLinenumber returns a boolean if a field has been set.

### GetLinetype

`func (o *ProductLineType) GetLinetype() string`

GetLinetype returns the Linetype field if non-nil, zero value otherwise.

### GetLinetypeOk

`func (o *ProductLineType) GetLinetypeOk() (*string, bool)`

GetLinetypeOk returns a tuple with the Linetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinetype

`func (o *ProductLineType) SetLinetype(v string)`

SetLinetype sets Linetype field to given value.

### HasLinetype

`func (o *ProductLineType) HasLinetype() bool`

HasLinetype returns a boolean if a field has been set.

### GetPartnumber

`func (o *ProductLineType) GetPartnumber() string`

GetPartnumber returns the Partnumber field if non-nil, zero value otherwise.

### GetPartnumberOk

`func (o *ProductLineType) GetPartnumberOk() (*string, bool)`

GetPartnumberOk returns a tuple with the Partnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnumber

`func (o *ProductLineType) SetPartnumber(v string)`

SetPartnumber sets Partnumber field to given value.

### HasPartnumber

`func (o *ProductLineType) HasPartnumber() bool`

HasPartnumber returns a boolean if a field has been set.

### GetVendorpartnumber

`func (o *ProductLineType) GetVendorpartnumber() string`

GetVendorpartnumber returns the Vendorpartnumber field if non-nil, zero value otherwise.

### GetVendorpartnumberOk

`func (o *ProductLineType) GetVendorpartnumberOk() (*string, bool)`

GetVendorpartnumberOk returns a tuple with the Vendorpartnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorpartnumber

`func (o *ProductLineType) SetVendorpartnumber(v string)`

SetVendorpartnumber sets Vendorpartnumber field to given value.

### HasVendorpartnumber

`func (o *ProductLineType) HasVendorpartnumber() bool`

HasVendorpartnumber returns a boolean if a field has been set.

### GetPartdescription

`func (o *ProductLineType) GetPartdescription() string`

GetPartdescription returns the Partdescription field if non-nil, zero value otherwise.

### GetPartdescriptionOk

`func (o *ProductLineType) GetPartdescriptionOk() (*string, bool)`

GetPartdescriptionOk returns a tuple with the Partdescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartdescription

`func (o *ProductLineType) SetPartdescription(v string)`

SetPartdescription sets Partdescription field to given value.

### HasPartdescription

`func (o *ProductLineType) HasPartdescription() bool`

HasPartdescription returns a boolean if a field has been set.

### GetShipfrombranch

`func (o *ProductLineType) GetShipfrombranch() string`

GetShipfrombranch returns the Shipfrombranch field if non-nil, zero value otherwise.

### GetShipfrombranchOk

`func (o *ProductLineType) GetShipfrombranchOk() (*string, bool)`

GetShipfrombranchOk returns a tuple with the Shipfrombranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipfrombranch

`func (o *ProductLineType) SetShipfrombranch(v string)`

SetShipfrombranch sets Shipfrombranch field to given value.

### HasShipfrombranch

`func (o *ProductLineType) HasShipfrombranch() bool`

HasShipfrombranch returns a boolean if a field has been set.

### GetShippedquantity

`func (o *ProductLineType) GetShippedquantity() string`

GetShippedquantity returns the Shippedquantity field if non-nil, zero value otherwise.

### GetShippedquantityOk

`func (o *ProductLineType) GetShippedquantityOk() (*string, bool)`

GetShippedquantityOk returns a tuple with the Shippedquantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedquantity

`func (o *ProductLineType) SetShippedquantity(v string)`

SetShippedquantity sets Shippedquantity field to given value.

### HasShippedquantity

`func (o *ProductLineType) HasShippedquantity() bool`

HasShippedquantity returns a boolean if a field has been set.

### GetOrderedquantity

`func (o *ProductLineType) GetOrderedquantity() string`

GetOrderedquantity returns the Orderedquantity field if non-nil, zero value otherwise.

### GetOrderedquantityOk

`func (o *ProductLineType) GetOrderedquantityOk() (*string, bool)`

GetOrderedquantityOk returns a tuple with the Orderedquantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderedquantity

`func (o *ProductLineType) SetOrderedquantity(v string)`

SetOrderedquantity sets Orderedquantity field to given value.

### HasOrderedquantity

`func (o *ProductLineType) HasOrderedquantity() bool`

HasOrderedquantity returns a boolean if a field has been set.

### GetMarginpercent

`func (o *ProductLineType) GetMarginpercent() string`

GetMarginpercent returns the Marginpercent field if non-nil, zero value otherwise.

### GetMarginpercentOk

`func (o *ProductLineType) GetMarginpercentOk() (*string, bool)`

GetMarginpercentOk returns a tuple with the Marginpercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginpercent

`func (o *ProductLineType) SetMarginpercent(v string)`

SetMarginpercent sets Marginpercent field to given value.

### HasMarginpercent

`func (o *ProductLineType) HasMarginpercent() bool`

HasMarginpercent returns a boolean if a field has been set.

### GetBackorderquantity

`func (o *ProductLineType) GetBackorderquantity() string`

GetBackorderquantity returns the Backorderquantity field if non-nil, zero value otherwise.

### GetBackorderquantityOk

`func (o *ProductLineType) GetBackorderquantityOk() (*string, bool)`

GetBackorderquantityOk returns a tuple with the Backorderquantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackorderquantity

`func (o *ProductLineType) SetBackorderquantity(v string)`

SetBackorderquantity sets Backorderquantity field to given value.

### HasBackorderquantity

`func (o *ProductLineType) HasBackorderquantity() bool`

HasBackorderquantity returns a boolean if a field has been set.

### GetBackorderetadate

`func (o *ProductLineType) GetBackorderetadate() string`

GetBackorderetadate returns the Backorderetadate field if non-nil, zero value otherwise.

### GetBackorderetadateOk

`func (o *ProductLineType) GetBackorderetadateOk() (*string, bool)`

GetBackorderetadateOk returns a tuple with the Backorderetadate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackorderetadate

`func (o *ProductLineType) SetBackorderetadate(v string)`

SetBackorderetadate sets Backorderetadate field to given value.

### HasBackorderetadate

`func (o *ProductLineType) HasBackorderetadate() bool`

HasBackorderetadate returns a boolean if a field has been set.

### GetExtendedprice

`func (o *ProductLineType) GetExtendedprice() string`

GetExtendedprice returns the Extendedprice field if non-nil, zero value otherwise.

### GetExtendedpriceOk

`func (o *ProductLineType) GetExtendedpriceOk() (*string, bool)`

GetExtendedpriceOk returns a tuple with the Extendedprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedprice

`func (o *ProductLineType) SetExtendedprice(v string)`

SetExtendedprice sets Extendedprice field to given value.

### HasExtendedprice

`func (o *ProductLineType) HasExtendedprice() bool`

HasExtendedprice returns a boolean if a field has been set.

### GetSpecialbidnumber

`func (o *ProductLineType) GetSpecialbidnumber() string`

GetSpecialbidnumber returns the Specialbidnumber field if non-nil, zero value otherwise.

### GetSpecialbidnumberOk

`func (o *ProductLineType) GetSpecialbidnumberOk() (*string, bool)`

GetSpecialbidnumberOk returns a tuple with the Specialbidnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialbidnumber

`func (o *ProductLineType) SetSpecialbidnumber(v string)`

SetSpecialbidnumber sets Specialbidnumber field to given value.

### HasSpecialbidnumber

`func (o *ProductLineType) HasSpecialbidnumber() bool`

HasSpecialbidnumber returns a boolean if a field has been set.

### GetOrdersuffix

`func (o *ProductLineType) GetOrdersuffix() string`

GetOrdersuffix returns the Ordersuffix field if non-nil, zero value otherwise.

### GetOrdersuffixOk

`func (o *ProductLineType) GetOrdersuffixOk() (*string, bool)`

GetOrdersuffixOk returns a tuple with the Ordersuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersuffix

`func (o *ProductLineType) SetOrdersuffix(v string)`

SetOrdersuffix sets Ordersuffix field to given value.

### HasOrdersuffix

`func (o *ProductLineType) HasOrdersuffix() bool`

HasOrdersuffix returns a boolean if a field has been set.

### GetIsacopapplied

`func (o *ProductLineType) GetIsacopapplied() string`

GetIsacopapplied returns the Isacopapplied field if non-nil, zero value otherwise.

### GetIsacopappliedOk

`func (o *ProductLineType) GetIsacopappliedOk() (*string, bool)`

GetIsacopappliedOk returns a tuple with the Isacopapplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsacopapplied

`func (o *ProductLineType) SetIsacopapplied(v string)`

SetIsacopapplied sets Isacopapplied field to given value.

### HasIsacopapplied

`func (o *ProductLineType) HasIsacopapplied() bool`

HasIsacopapplied returns a boolean if a field has been set.

### GetUnitprice

`func (o *ProductLineType) GetUnitprice() string`

GetUnitprice returns the Unitprice field if non-nil, zero value otherwise.

### GetUnitpriceOk

`func (o *ProductLineType) GetUnitpriceOk() (*string, bool)`

GetUnitpriceOk returns a tuple with the Unitprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitprice

`func (o *ProductLineType) SetUnitprice(v string)`

SetUnitprice sets Unitprice field to given value.

### HasUnitprice

`func (o *ProductLineType) HasUnitprice() bool`

HasUnitprice returns a boolean if a field has been set.

### GetUnitofmeasure

`func (o *ProductLineType) GetUnitofmeasure() string`

GetUnitofmeasure returns the Unitofmeasure field if non-nil, zero value otherwise.

### GetUnitofmeasureOk

`func (o *ProductLineType) GetUnitofmeasureOk() (*string, bool)`

GetUnitofmeasureOk returns a tuple with the Unitofmeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitofmeasure

`func (o *ProductLineType) SetUnitofmeasure(v string)`

SetUnitofmeasure sets Unitofmeasure field to given value.

### HasUnitofmeasure

`func (o *ProductLineType) HasUnitofmeasure() bool`

HasUnitofmeasure returns a boolean if a field has been set.

### GetSerialnumberdetails

`func (o *ProductLineType) GetSerialnumberdetails() []ProductLineTypeSerialnumberdetailsInner`

GetSerialnumberdetails returns the Serialnumberdetails field if non-nil, zero value otherwise.

### GetSerialnumberdetailsOk

`func (o *ProductLineType) GetSerialnumberdetailsOk() (*[]ProductLineTypeSerialnumberdetailsInner, bool)`

GetSerialnumberdetailsOk returns a tuple with the Serialnumberdetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialnumberdetails

`func (o *ProductLineType) SetSerialnumberdetails(v []ProductLineTypeSerialnumberdetailsInner)`

SetSerialnumberdetails sets Serialnumberdetails field to given value.

### HasSerialnumberdetails

`func (o *ProductLineType) HasSerialnumberdetails() bool`

HasSerialnumberdetails returns a boolean if a field has been set.

### GetTrackingnumberdetails

`func (o *ProductLineType) GetTrackingnumberdetails() []ProductLineTypeTrackingnumberdetailsInner`

GetTrackingnumberdetails returns the Trackingnumberdetails field if non-nil, zero value otherwise.

### GetTrackingnumberdetailsOk

`func (o *ProductLineType) GetTrackingnumberdetailsOk() (*[]ProductLineTypeTrackingnumberdetailsInner, bool)`

GetTrackingnumberdetailsOk returns a tuple with the Trackingnumberdetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingnumberdetails

`func (o *ProductLineType) SetTrackingnumberdetails(v []ProductLineTypeTrackingnumberdetailsInner)`

SetTrackingnumberdetails sets Trackingnumberdetails field to given value.

### HasTrackingnumberdetails

`func (o *ProductLineType) HasTrackingnumberdetails() bool`

HasTrackingnumberdetails returns a boolean if a field has been set.

### GetProductextendedspecs

`func (o *ProductLineType) GetProductextendedspecs() []InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner`

GetProductextendedspecs returns the Productextendedspecs field if non-nil, zero value otherwise.

### GetProductextendedspecsOk

`func (o *ProductLineType) GetProductextendedspecsOk() (*[]InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner, bool)`

GetProductextendedspecsOk returns a tuple with the Productextendedspecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductextendedspecs

`func (o *ProductLineType) SetProductextendedspecs(v []InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner)`

SetProductextendedspecs sets Productextendedspecs field to given value.

### HasProductextendedspecs

`func (o *ProductLineType) HasProductextendedspecs() bool`

HasProductextendedspecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


