/*
XI Sdk Resellers

For Resellers. Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the QuoteDetailsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteDetailsResponse{}

// QuoteDetailsResponse struct for QuoteDetailsResponse
type QuoteDetailsResponse struct {
	// Quote Name given to quote by sales team or system generated.  Generally used as a reference to identify the quote.
	QuoteName *string `json:"quoteName,omitempty"`
	// Unique identifier generated by Ingram Micro's CRM specific to each quote.  When applying a filter to the quoteNumber and including a partial quote number in the filter, all quotes containing any information included in the filter can be retrieved as a subset of all available customer quotes.
	QuoteNumber *string `json:"quoteNumber,omitempty"`
	// When a quote has been revised and updated, the quote number remains the same throughout the lifecycle of the quote, however, a Revision number is updated for each revision of the quote.  The revision numbers is associated with the Unique Quote Number.
	Revision *string `json:"revision,omitempty"`
	// Date the Quote was initially Created.
	IngramQuoteDate *string `json:"ingramQuoteDate,omitempty"`
	// Date the Quote was last updated or modified.
	LastModifiedDate *string `json:"lastModifiedDate,omitempty"`
	// Quote expiration date.
	IngramQuoteExpiryDate *string `json:"ingramQuoteExpiryDate,omitempty"`
	// Three letter currency code.
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// Price discount identifyer to specify  a pricing discount that has been applied to the quote. If present - the priceDeviationStartDate and priceDeviationExpiryDate must be presented. Cisco refers to this as a Dart
	SpecialBidId *string `json:"specialBidId,omitempty"`
	// If price discount has been applied to the quote - the starting date the discount begins.
	SpecialBidEffectiveDate *string `json:"specialBidEffectiveDate,omitempty"`
	// If a price discount has been applied to the quote - The date the discount expires and will no longer be applicable.
	SpecialBidExpirationDate *string `json:"specialBidExpirationDate,omitempty"`
	// This refers to the primary status of the quote.  API responses will return
	Status *string `json:"status,omitempty"`
	// Details related to the customer's request for the quote entered by the sales representative or system generated.
	CustomerNeed *string `json:"customerNeed,omitempty"`
	// Ingram Micro proposed solution and summary of quote.
	ProposedSolution *string `json:"proposedSolution,omitempty"`
	// Introductory paragraph included in each quote.  Legally required - must be included when presenting the quote details.
	IntroPreamble *string `json:"introPreamble,omitempty"`
	// Purchase instructions.  Legally required - must be included when presenting the quote details.
	PurchaseInstructions *string `json:"purchaseInstructions,omitempty"`
	// Legal terms -  Legally required - must be included when presenting the quote details.
	LegalTerms *string `json:"legalTerms,omitempty"`
	// Lease information.
	LeaseInfo *string `json:"leaseInfo,omitempty"`
	// Leasing information
	LeasingInstructions *string `json:"leasingInstructions,omitempty"`
	ResellerInfo *QuoteDetailsResponseResellerInfo `json:"resellerInfo,omitempty"`
	EndUserInfo *QuoteDetailsResponseEndUserInfo `json:"endUserInfo,omitempty"`
	Products []QuoteDetailsResponseProductsInner `json:"products,omitempty"`
	// Total number of products included in the quote
	ProductsCount *int32 `json:"productsCount,omitempty"`
	// Total extended MSRP for all products included in the quote
	ExtendedMsrpTotal *int32 `json:"extendedMsrpTotal,omitempty"`
	// Total quantity of all items in the quote.
	QuantityTotal *int32 `json:"quantityTotal,omitempty"`
	// Total amount of quoted price for all products in the quote including both solution products and suggested products.
	ExtendedQuotePriceTotal *int32 `json:"extendedQuotePriceTotal,omitempty"`
	AdditionalAttributes []QuoteDetailsResponseAdditionalAttributesInner `json:"additionalAttributes,omitempty"`
}

// NewQuoteDetailsResponse instantiates a new QuoteDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteDetailsResponse() *QuoteDetailsResponse {
	this := QuoteDetailsResponse{}
	return &this
}

// NewQuoteDetailsResponseWithDefaults instantiates a new QuoteDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteDetailsResponseWithDefaults() *QuoteDetailsResponse {
	this := QuoteDetailsResponse{}
	return &this
}

// GetQuoteName returns the QuoteName field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetQuoteName() string {
	if o == nil || IsNil(o.QuoteName) {
		var ret string
		return ret
	}
	return *o.QuoteName
}

// GetQuoteNameOk returns a tuple with the QuoteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetQuoteNameOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteName) {
		return nil, false
	}
	return o.QuoteName, true
}

// HasQuoteName returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasQuoteName() bool {
	if o != nil && !IsNil(o.QuoteName) {
		return true
	}

	return false
}

// SetQuoteName gets a reference to the given string and assigns it to the QuoteName field.
func (o *QuoteDetailsResponse) SetQuoteName(v string) {
	o.QuoteName = &v
}

// GetQuoteNumber returns the QuoteNumber field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetQuoteNumber() string {
	if o == nil || IsNil(o.QuoteNumber) {
		var ret string
		return ret
	}
	return *o.QuoteNumber
}

// GetQuoteNumberOk returns a tuple with the QuoteNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetQuoteNumberOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteNumber) {
		return nil, false
	}
	return o.QuoteNumber, true
}

// HasQuoteNumber returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasQuoteNumber() bool {
	if o != nil && !IsNil(o.QuoteNumber) {
		return true
	}

	return false
}

// SetQuoteNumber gets a reference to the given string and assigns it to the QuoteNumber field.
func (o *QuoteDetailsResponse) SetQuoteNumber(v string) {
	o.QuoteNumber = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetRevision() string {
	if o == nil || IsNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *QuoteDetailsResponse) SetRevision(v string) {
	o.Revision = &v
}

// GetIngramQuoteDate returns the IngramQuoteDate field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetIngramQuoteDate() string {
	if o == nil || IsNil(o.IngramQuoteDate) {
		var ret string
		return ret
	}
	return *o.IngramQuoteDate
}

// GetIngramQuoteDateOk returns a tuple with the IngramQuoteDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetIngramQuoteDateOk() (*string, bool) {
	if o == nil || IsNil(o.IngramQuoteDate) {
		return nil, false
	}
	return o.IngramQuoteDate, true
}

// HasIngramQuoteDate returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasIngramQuoteDate() bool {
	if o != nil && !IsNil(o.IngramQuoteDate) {
		return true
	}

	return false
}

// SetIngramQuoteDate gets a reference to the given string and assigns it to the IngramQuoteDate field.
func (o *QuoteDetailsResponse) SetIngramQuoteDate(v string) {
	o.IngramQuoteDate = &v
}

// GetLastModifiedDate returns the LastModifiedDate field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetLastModifiedDate() string {
	if o == nil || IsNil(o.LastModifiedDate) {
		var ret string
		return ret
	}
	return *o.LastModifiedDate
}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetLastModifiedDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedDate) {
		return nil, false
	}
	return o.LastModifiedDate, true
}

// HasLastModifiedDate returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasLastModifiedDate() bool {
	if o != nil && !IsNil(o.LastModifiedDate) {
		return true
	}

	return false
}

// SetLastModifiedDate gets a reference to the given string and assigns it to the LastModifiedDate field.
func (o *QuoteDetailsResponse) SetLastModifiedDate(v string) {
	o.LastModifiedDate = &v
}

// GetIngramQuoteExpiryDate returns the IngramQuoteExpiryDate field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetIngramQuoteExpiryDate() string {
	if o == nil || IsNil(o.IngramQuoteExpiryDate) {
		var ret string
		return ret
	}
	return *o.IngramQuoteExpiryDate
}

// GetIngramQuoteExpiryDateOk returns a tuple with the IngramQuoteExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetIngramQuoteExpiryDateOk() (*string, bool) {
	if o == nil || IsNil(o.IngramQuoteExpiryDate) {
		return nil, false
	}
	return o.IngramQuoteExpiryDate, true
}

// HasIngramQuoteExpiryDate returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasIngramQuoteExpiryDate() bool {
	if o != nil && !IsNil(o.IngramQuoteExpiryDate) {
		return true
	}

	return false
}

// SetIngramQuoteExpiryDate gets a reference to the given string and assigns it to the IngramQuoteExpiryDate field.
func (o *QuoteDetailsResponse) SetIngramQuoteExpiryDate(v string) {
	o.IngramQuoteExpiryDate = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *QuoteDetailsResponse) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetSpecialBidId returns the SpecialBidId field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetSpecialBidId() string {
	if o == nil || IsNil(o.SpecialBidId) {
		var ret string
		return ret
	}
	return *o.SpecialBidId
}

// GetSpecialBidIdOk returns a tuple with the SpecialBidId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetSpecialBidIdOk() (*string, bool) {
	if o == nil || IsNil(o.SpecialBidId) {
		return nil, false
	}
	return o.SpecialBidId, true
}

// HasSpecialBidId returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasSpecialBidId() bool {
	if o != nil && !IsNil(o.SpecialBidId) {
		return true
	}

	return false
}

// SetSpecialBidId gets a reference to the given string and assigns it to the SpecialBidId field.
func (o *QuoteDetailsResponse) SetSpecialBidId(v string) {
	o.SpecialBidId = &v
}

// GetSpecialBidEffectiveDate returns the SpecialBidEffectiveDate field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetSpecialBidEffectiveDate() string {
	if o == nil || IsNil(o.SpecialBidEffectiveDate) {
		var ret string
		return ret
	}
	return *o.SpecialBidEffectiveDate
}

// GetSpecialBidEffectiveDateOk returns a tuple with the SpecialBidEffectiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetSpecialBidEffectiveDateOk() (*string, bool) {
	if o == nil || IsNil(o.SpecialBidEffectiveDate) {
		return nil, false
	}
	return o.SpecialBidEffectiveDate, true
}

// HasSpecialBidEffectiveDate returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasSpecialBidEffectiveDate() bool {
	if o != nil && !IsNil(o.SpecialBidEffectiveDate) {
		return true
	}

	return false
}

// SetSpecialBidEffectiveDate gets a reference to the given string and assigns it to the SpecialBidEffectiveDate field.
func (o *QuoteDetailsResponse) SetSpecialBidEffectiveDate(v string) {
	o.SpecialBidEffectiveDate = &v
}

// GetSpecialBidExpirationDate returns the SpecialBidExpirationDate field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetSpecialBidExpirationDate() string {
	if o == nil || IsNil(o.SpecialBidExpirationDate) {
		var ret string
		return ret
	}
	return *o.SpecialBidExpirationDate
}

// GetSpecialBidExpirationDateOk returns a tuple with the SpecialBidExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetSpecialBidExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.SpecialBidExpirationDate) {
		return nil, false
	}
	return o.SpecialBidExpirationDate, true
}

// HasSpecialBidExpirationDate returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasSpecialBidExpirationDate() bool {
	if o != nil && !IsNil(o.SpecialBidExpirationDate) {
		return true
	}

	return false
}

// SetSpecialBidExpirationDate gets a reference to the given string and assigns it to the SpecialBidExpirationDate field.
func (o *QuoteDetailsResponse) SetSpecialBidExpirationDate(v string) {
	o.SpecialBidExpirationDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QuoteDetailsResponse) SetStatus(v string) {
	o.Status = &v
}

// GetCustomerNeed returns the CustomerNeed field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetCustomerNeed() string {
	if o == nil || IsNil(o.CustomerNeed) {
		var ret string
		return ret
	}
	return *o.CustomerNeed
}

// GetCustomerNeedOk returns a tuple with the CustomerNeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetCustomerNeedOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerNeed) {
		return nil, false
	}
	return o.CustomerNeed, true
}

// HasCustomerNeed returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasCustomerNeed() bool {
	if o != nil && !IsNil(o.CustomerNeed) {
		return true
	}

	return false
}

// SetCustomerNeed gets a reference to the given string and assigns it to the CustomerNeed field.
func (o *QuoteDetailsResponse) SetCustomerNeed(v string) {
	o.CustomerNeed = &v
}

// GetProposedSolution returns the ProposedSolution field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetProposedSolution() string {
	if o == nil || IsNil(o.ProposedSolution) {
		var ret string
		return ret
	}
	return *o.ProposedSolution
}

// GetProposedSolutionOk returns a tuple with the ProposedSolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetProposedSolutionOk() (*string, bool) {
	if o == nil || IsNil(o.ProposedSolution) {
		return nil, false
	}
	return o.ProposedSolution, true
}

// HasProposedSolution returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasProposedSolution() bool {
	if o != nil && !IsNil(o.ProposedSolution) {
		return true
	}

	return false
}

// SetProposedSolution gets a reference to the given string and assigns it to the ProposedSolution field.
func (o *QuoteDetailsResponse) SetProposedSolution(v string) {
	o.ProposedSolution = &v
}

// GetIntroPreamble returns the IntroPreamble field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetIntroPreamble() string {
	if o == nil || IsNil(o.IntroPreamble) {
		var ret string
		return ret
	}
	return *o.IntroPreamble
}

// GetIntroPreambleOk returns a tuple with the IntroPreamble field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetIntroPreambleOk() (*string, bool) {
	if o == nil || IsNil(o.IntroPreamble) {
		return nil, false
	}
	return o.IntroPreamble, true
}

// HasIntroPreamble returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasIntroPreamble() bool {
	if o != nil && !IsNil(o.IntroPreamble) {
		return true
	}

	return false
}

// SetIntroPreamble gets a reference to the given string and assigns it to the IntroPreamble field.
func (o *QuoteDetailsResponse) SetIntroPreamble(v string) {
	o.IntroPreamble = &v
}

// GetPurchaseInstructions returns the PurchaseInstructions field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetPurchaseInstructions() string {
	if o == nil || IsNil(o.PurchaseInstructions) {
		var ret string
		return ret
	}
	return *o.PurchaseInstructions
}

// GetPurchaseInstructionsOk returns a tuple with the PurchaseInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetPurchaseInstructionsOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseInstructions) {
		return nil, false
	}
	return o.PurchaseInstructions, true
}

// HasPurchaseInstructions returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasPurchaseInstructions() bool {
	if o != nil && !IsNil(o.PurchaseInstructions) {
		return true
	}

	return false
}

// SetPurchaseInstructions gets a reference to the given string and assigns it to the PurchaseInstructions field.
func (o *QuoteDetailsResponse) SetPurchaseInstructions(v string) {
	o.PurchaseInstructions = &v
}

// GetLegalTerms returns the LegalTerms field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetLegalTerms() string {
	if o == nil || IsNil(o.LegalTerms) {
		var ret string
		return ret
	}
	return *o.LegalTerms
}

// GetLegalTermsOk returns a tuple with the LegalTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetLegalTermsOk() (*string, bool) {
	if o == nil || IsNil(o.LegalTerms) {
		return nil, false
	}
	return o.LegalTerms, true
}

// HasLegalTerms returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasLegalTerms() bool {
	if o != nil && !IsNil(o.LegalTerms) {
		return true
	}

	return false
}

// SetLegalTerms gets a reference to the given string and assigns it to the LegalTerms field.
func (o *QuoteDetailsResponse) SetLegalTerms(v string) {
	o.LegalTerms = &v
}

// GetLeaseInfo returns the LeaseInfo field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetLeaseInfo() string {
	if o == nil || IsNil(o.LeaseInfo) {
		var ret string
		return ret
	}
	return *o.LeaseInfo
}

// GetLeaseInfoOk returns a tuple with the LeaseInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetLeaseInfoOk() (*string, bool) {
	if o == nil || IsNil(o.LeaseInfo) {
		return nil, false
	}
	return o.LeaseInfo, true
}

// HasLeaseInfo returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasLeaseInfo() bool {
	if o != nil && !IsNil(o.LeaseInfo) {
		return true
	}

	return false
}

// SetLeaseInfo gets a reference to the given string and assigns it to the LeaseInfo field.
func (o *QuoteDetailsResponse) SetLeaseInfo(v string) {
	o.LeaseInfo = &v
}

// GetLeasingInstructions returns the LeasingInstructions field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetLeasingInstructions() string {
	if o == nil || IsNil(o.LeasingInstructions) {
		var ret string
		return ret
	}
	return *o.LeasingInstructions
}

// GetLeasingInstructionsOk returns a tuple with the LeasingInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetLeasingInstructionsOk() (*string, bool) {
	if o == nil || IsNil(o.LeasingInstructions) {
		return nil, false
	}
	return o.LeasingInstructions, true
}

// HasLeasingInstructions returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasLeasingInstructions() bool {
	if o != nil && !IsNil(o.LeasingInstructions) {
		return true
	}

	return false
}

// SetLeasingInstructions gets a reference to the given string and assigns it to the LeasingInstructions field.
func (o *QuoteDetailsResponse) SetLeasingInstructions(v string) {
	o.LeasingInstructions = &v
}

// GetResellerInfo returns the ResellerInfo field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetResellerInfo() QuoteDetailsResponseResellerInfo {
	if o == nil || IsNil(o.ResellerInfo) {
		var ret QuoteDetailsResponseResellerInfo
		return ret
	}
	return *o.ResellerInfo
}

// GetResellerInfoOk returns a tuple with the ResellerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetResellerInfoOk() (*QuoteDetailsResponseResellerInfo, bool) {
	if o == nil || IsNil(o.ResellerInfo) {
		return nil, false
	}
	return o.ResellerInfo, true
}

// HasResellerInfo returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasResellerInfo() bool {
	if o != nil && !IsNil(o.ResellerInfo) {
		return true
	}

	return false
}

// SetResellerInfo gets a reference to the given QuoteDetailsResponseResellerInfo and assigns it to the ResellerInfo field.
func (o *QuoteDetailsResponse) SetResellerInfo(v QuoteDetailsResponseResellerInfo) {
	o.ResellerInfo = &v
}

// GetEndUserInfo returns the EndUserInfo field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetEndUserInfo() QuoteDetailsResponseEndUserInfo {
	if o == nil || IsNil(o.EndUserInfo) {
		var ret QuoteDetailsResponseEndUserInfo
		return ret
	}
	return *o.EndUserInfo
}

// GetEndUserInfoOk returns a tuple with the EndUserInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetEndUserInfoOk() (*QuoteDetailsResponseEndUserInfo, bool) {
	if o == nil || IsNil(o.EndUserInfo) {
		return nil, false
	}
	return o.EndUserInfo, true
}

// HasEndUserInfo returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasEndUserInfo() bool {
	if o != nil && !IsNil(o.EndUserInfo) {
		return true
	}

	return false
}

// SetEndUserInfo gets a reference to the given QuoteDetailsResponseEndUserInfo and assigns it to the EndUserInfo field.
func (o *QuoteDetailsResponse) SetEndUserInfo(v QuoteDetailsResponseEndUserInfo) {
	o.EndUserInfo = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetProducts() []QuoteDetailsResponseProductsInner {
	if o == nil || IsNil(o.Products) {
		var ret []QuoteDetailsResponseProductsInner
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetProductsOk() ([]QuoteDetailsResponseProductsInner, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []QuoteDetailsResponseProductsInner and assigns it to the Products field.
func (o *QuoteDetailsResponse) SetProducts(v []QuoteDetailsResponseProductsInner) {
	o.Products = v
}

// GetProductsCount returns the ProductsCount field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetProductsCount() int32 {
	if o == nil || IsNil(o.ProductsCount) {
		var ret int32
		return ret
	}
	return *o.ProductsCount
}

// GetProductsCountOk returns a tuple with the ProductsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetProductsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ProductsCount) {
		return nil, false
	}
	return o.ProductsCount, true
}

// HasProductsCount returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasProductsCount() bool {
	if o != nil && !IsNil(o.ProductsCount) {
		return true
	}

	return false
}

// SetProductsCount gets a reference to the given int32 and assigns it to the ProductsCount field.
func (o *QuoteDetailsResponse) SetProductsCount(v int32) {
	o.ProductsCount = &v
}

// GetExtendedMsrpTotal returns the ExtendedMsrpTotal field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetExtendedMsrpTotal() int32 {
	if o == nil || IsNil(o.ExtendedMsrpTotal) {
		var ret int32
		return ret
	}
	return *o.ExtendedMsrpTotal
}

// GetExtendedMsrpTotalOk returns a tuple with the ExtendedMsrpTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetExtendedMsrpTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.ExtendedMsrpTotal) {
		return nil, false
	}
	return o.ExtendedMsrpTotal, true
}

// HasExtendedMsrpTotal returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasExtendedMsrpTotal() bool {
	if o != nil && !IsNil(o.ExtendedMsrpTotal) {
		return true
	}

	return false
}

// SetExtendedMsrpTotal gets a reference to the given int32 and assigns it to the ExtendedMsrpTotal field.
func (o *QuoteDetailsResponse) SetExtendedMsrpTotal(v int32) {
	o.ExtendedMsrpTotal = &v
}

// GetQuantityTotal returns the QuantityTotal field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetQuantityTotal() int32 {
	if o == nil || IsNil(o.QuantityTotal) {
		var ret int32
		return ret
	}
	return *o.QuantityTotal
}

// GetQuantityTotalOk returns a tuple with the QuantityTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetQuantityTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityTotal) {
		return nil, false
	}
	return o.QuantityTotal, true
}

// HasQuantityTotal returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasQuantityTotal() bool {
	if o != nil && !IsNil(o.QuantityTotal) {
		return true
	}

	return false
}

// SetQuantityTotal gets a reference to the given int32 and assigns it to the QuantityTotal field.
func (o *QuoteDetailsResponse) SetQuantityTotal(v int32) {
	o.QuantityTotal = &v
}

// GetExtendedQuotePriceTotal returns the ExtendedQuotePriceTotal field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetExtendedQuotePriceTotal() int32 {
	if o == nil || IsNil(o.ExtendedQuotePriceTotal) {
		var ret int32
		return ret
	}
	return *o.ExtendedQuotePriceTotal
}

// GetExtendedQuotePriceTotalOk returns a tuple with the ExtendedQuotePriceTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetExtendedQuotePriceTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.ExtendedQuotePriceTotal) {
		return nil, false
	}
	return o.ExtendedQuotePriceTotal, true
}

// HasExtendedQuotePriceTotal returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasExtendedQuotePriceTotal() bool {
	if o != nil && !IsNil(o.ExtendedQuotePriceTotal) {
		return true
	}

	return false
}

// SetExtendedQuotePriceTotal gets a reference to the given int32 and assigns it to the ExtendedQuotePriceTotal field.
func (o *QuoteDetailsResponse) SetExtendedQuotePriceTotal(v int32) {
	o.ExtendedQuotePriceTotal = &v
}

// GetAdditionalAttributes returns the AdditionalAttributes field value if set, zero value otherwise.
func (o *QuoteDetailsResponse) GetAdditionalAttributes() []QuoteDetailsResponseAdditionalAttributesInner {
	if o == nil || IsNil(o.AdditionalAttributes) {
		var ret []QuoteDetailsResponseAdditionalAttributesInner
		return ret
	}
	return o.AdditionalAttributes
}

// GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponse) GetAdditionalAttributesOk() ([]QuoteDetailsResponseAdditionalAttributesInner, bool) {
	if o == nil || IsNil(o.AdditionalAttributes) {
		return nil, false
	}
	return o.AdditionalAttributes, true
}

// HasAdditionalAttributes returns a boolean if a field has been set.
func (o *QuoteDetailsResponse) HasAdditionalAttributes() bool {
	if o != nil && !IsNil(o.AdditionalAttributes) {
		return true
	}

	return false
}

// SetAdditionalAttributes gets a reference to the given []QuoteDetailsResponseAdditionalAttributesInner and assigns it to the AdditionalAttributes field.
func (o *QuoteDetailsResponse) SetAdditionalAttributes(v []QuoteDetailsResponseAdditionalAttributesInner) {
	o.AdditionalAttributes = v
}

func (o QuoteDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QuoteName) {
		toSerialize["quoteName"] = o.QuoteName
	}
	if !IsNil(o.QuoteNumber) {
		toSerialize["quoteNumber"] = o.QuoteNumber
	}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !IsNil(o.IngramQuoteDate) {
		toSerialize["ingramQuoteDate"] = o.IngramQuoteDate
	}
	if !IsNil(o.LastModifiedDate) {
		toSerialize["lastModifiedDate"] = o.LastModifiedDate
	}
	if !IsNil(o.IngramQuoteExpiryDate) {
		toSerialize["ingramQuoteExpiryDate"] = o.IngramQuoteExpiryDate
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.SpecialBidId) {
		toSerialize["specialBidId"] = o.SpecialBidId
	}
	if !IsNil(o.SpecialBidEffectiveDate) {
		toSerialize["specialBidEffectiveDate"] = o.SpecialBidEffectiveDate
	}
	if !IsNil(o.SpecialBidExpirationDate) {
		toSerialize["specialBidExpirationDate"] = o.SpecialBidExpirationDate
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CustomerNeed) {
		toSerialize["customerNeed"] = o.CustomerNeed
	}
	if !IsNil(o.ProposedSolution) {
		toSerialize["proposedSolution"] = o.ProposedSolution
	}
	if !IsNil(o.IntroPreamble) {
		toSerialize["introPreamble"] = o.IntroPreamble
	}
	if !IsNil(o.PurchaseInstructions) {
		toSerialize["purchaseInstructions"] = o.PurchaseInstructions
	}
	if !IsNil(o.LegalTerms) {
		toSerialize["legalTerms"] = o.LegalTerms
	}
	if !IsNil(o.LeaseInfo) {
		toSerialize["leaseInfo"] = o.LeaseInfo
	}
	if !IsNil(o.LeasingInstructions) {
		toSerialize["leasingInstructions"] = o.LeasingInstructions
	}
	if !IsNil(o.ResellerInfo) {
		toSerialize["resellerInfo"] = o.ResellerInfo
	}
	if !IsNil(o.EndUserInfo) {
		toSerialize["endUserInfo"] = o.EndUserInfo
	}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	if !IsNil(o.ProductsCount) {
		toSerialize["productsCount"] = o.ProductsCount
	}
	if !IsNil(o.ExtendedMsrpTotal) {
		toSerialize["extendedMsrpTotal"] = o.ExtendedMsrpTotal
	}
	if !IsNil(o.QuantityTotal) {
		toSerialize["quantityTotal"] = o.QuantityTotal
	}
	if !IsNil(o.ExtendedQuotePriceTotal) {
		toSerialize["extendedQuotePriceTotal"] = o.ExtendedQuotePriceTotal
	}
	if !IsNil(o.AdditionalAttributes) {
		toSerialize["additionalAttributes"] = o.AdditionalAttributes
	}
	return toSerialize, nil
}

type NullableQuoteDetailsResponse struct {
	value *QuoteDetailsResponse
	isSet bool
}

func (v NullableQuoteDetailsResponse) Get() *QuoteDetailsResponse {
	return v.value
}

func (v *NullableQuoteDetailsResponse) Set(val *QuoteDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteDetailsResponse(val *QuoteDetailsResponse) *NullableQuoteDetailsResponse {
	return &NullableQuoteDetailsResponse{value: val, isSet: true}
}

func (v NullableQuoteDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


