package basware

// Accounting related content.
type Accounting struct {
	// Virtual bar code can be added to the business document that should be
	// printed.
	VirtualBankBarcode VirtualBankBarcode `json:"virtualBankBarcode,omitempty"`
}

// Party that is the accountable buyer of the goods/services in the referred
// business document.
type AccountingCustomerParty struct {
	// An array holding the external system identifiers of the party. Used for
	// defining customer, supplier and delivery party data.
	Endpoint Endpoint `json:"endpoint,omitempty"`

	// An array holding the external system identifiers of the party. Used for
	// defining customer, supplier and delivery party data.
	PartyIdentification []PartyIdentificationItem `json:"partyIdentification,omitempty"`

	// A name of the party. Used for defining supplier, customer and delivery
	// party names.
	PartyName string `json:"partyName"`

	// An object containing address information. Used for defining supplier
	// party, customer party and delivery party address data.
	PostalAddress PostalAddress `json:"postalAddress,omitempty"`

	// An object containing information about contacts. Used for defining the
	// company contact data
	Contact Contact `json:"contact,omitempty"`

	// Information about taxes. Notice that only one tax scheme is used,
	// although there could be multiple.
	PartyTaxScheme PartyTaxScheme `json:"partyTaxScheme,omitempty"`
}

// Information about taxes. Notice that only one tax scheme is used, although
// there could be multiple.
type PartyTaxScheme struct {
	// Information about the company taxes.
	Company PartyTaxSchemeCompany `json:"company,omitempty"`
}

// Information about the company taxes.
type PartyTaxSchemeCompany struct {
	// A tax identifier for a company. The identifier assigned for tax purposes
	// to a party by the taxation authority.
	ID string `json:"id,omitempty"`

	// External global identifier of the endpoint identifier element. Valid
	// values: Country specific agency schema, example DK:CVR for Denmark.
	SchemeID string `json:"schemeId,omitempty"`
}

// Party that is the accountable supplier of the goods/services in the referred
// business document.
type AccountingSupplierParty struct {
	// An array holding the external system identifiers of the party. Used for
	// defining customer, supplier and delivery party data.
	Endpoint Endpoint `json:"endpoint,omitempty"`

	// An array holding the external system identifiers of the party. Used for
	// defining customer, supplier and delivery party data.
	PartyIdentification []PartyIdentificationItem `json:"partyIdentification,omitempty"`

	// A name of the party. Used for defining supplier, customer and delivery
	// party names.
	PartyName string `json:"partyName"`

	// An object containing address information. Used for defining supplier
	// party, customer party and delivery party address data.
	PostalAddress PostalAddress `json:"postalAddress,omitempty"`

	// Information about taxes. Notice that only one tax scheme is used,
	// although there could be multiple.
	PartyTaxScheme PartyTaxScheme `json:"partyTaxScheme,omitempty"`

	// An object containing information about contacts. Used for defining the
	// company contact data
	Contact Contact `json:"contact,omitempty"`
}

// An object containing information about contacts. Used for defining the
// company contact data
type Contact struct {
	// A contact name of the party.
	Name string `json:"name,omitempty"`

	// A telephone number of the contact of the party.
	Telephone string `json:"telephone,omitempty"`

	// A fax number of the contact of the party.
	Telefax string `json:"telefax,omitempty"`

	// An email of the contact of the party.
	ElectronicMail string `json:"electronicMail,omitempty"`
}

// An array holding the external system identifiers of the party. Used for
// defining customer, supplier and delivery party data.
type Endpoint struct {
	ID       string `json:"id"`
	SchemeID string `json:"schemeId,omitempty"`
}

// An object holding a party identification.
type PartyIdentificationItem struct {
	// An object holding the external system identifier of the party.
	ID string `json:"id"`

	// External global identifier of the id identifier element.
	SchemeID string `json:"schemeId,omitempty"`
}

type AdditionalDocumentReference struct {
	// An identifier for the referenced document (i.e. bumid).
	ID string `json:"id"`

	// External system specific identifier of the invoicing system identifier
	// element. If the source business document has any matching element, it
	// should be used.
	SchemeID string `json:"schemeId,omitempty"`

	// Date when the referenced document was issued. Valid values must be in
	// format: CCYY-MM-DD. If the time zone is known, it must be represented
	// with +hh:mm or -hh:mm or Z (which means UTC). If time zone is not known,
	// it must be left empty.
	IssueDate string `json:"issueDate,omitempty"`

	// The type of document being referenced, expressed as a code, for example
	// to reference to an Invoice document, code is 380.
	TypeCode string `json:"typeCode,omitempty"`
}

type BillingReference struct {
	// External system identifier of the billing entity referenced by the
	// Business Document.
	ID string `json:"id"`

	// External system specific identifier of the billing reference system
	// identifier element. If the source business document has any matching
	// element, it should be used.
	SchemeID string `json:"schemeId,omitempty"`
}

type BuyerReference struct {
	// The id value of the buyer reference
	ID string `json:"id"`
}

type ContractDocumentReference struct {
	// External system identifier of the contract referenced by the Business
	// Document (i.e. buyers contract number). Mandatory field if the customer
	// demands that the goods or services invoiced refer to a contract number
	// defined by the customer to which he wants to assign the Business
	// Document. Is demanded for example in service and maintenance agreements
	// for which there is generally no explicit order.
	ID string `json:"id"`

	// External system specific identifier of the contract system identifier
	// element. If the source business document has any matching element, it should
	// be used.
	SchemeID string `json:"schemeId,omitempty"`
}

// Object holding the business content of the Invoice. Content is at some level
// based on Universal Business Language (UBL) standard version 2.1. It has also
// been extended by Basware so it is not strictly UBL.
type Data struct {
	// External system identifier of the business document.
	ID string `json:"id"`

	// External system specific identifier of the system identifier element. If
	// the source business document has any matching element, it should be used.
	IDSchemeID string `json:"idSchemeId,omitempty"`

	// The date when the Invoice was issued. Valid values must be in format:
	// CCYY-MM-DD. If the time zone is known, it must be represented with +hh:mm
	// or -hh:mm or Z (which means UTC). If time zone is not known, it must be
	// left empty.
	IssueDate string `json:"issueDate"`

	// Currency presentation of the Invoice document. Valid values must be ISO
	// 4217 Alpha format.
	DocumentCurrencyCode string `json:"documentCurrencyCode,omitempty"`

	// Free-form text pertinent to this document, conveying information that is
	// not contained explicitly in other structures.
	Note string `json:"note,omitempty"`

	AllowanceCharge             AllowanceCharge             `json:"allowanceCharge,omitempty"`
	OrderReference              OrderReference              `json:"orderReference,omitempty"`
	BillingReference            BillingReference            `json:"billingReference,omitempty"`
	ContractDocumentReference   ContractDocumentReference   `json:"contractDocumentReference,omitempty"`
	AdditionalDocumentReference AdditionalDocumentReference `json:"additionalDocumentReference,omitempty"`

	// Party that is the accountable supplier of the goods/services in the
	// referred business document.
	AccountingSupplierParty AccountingSupplierParty `json:"accountingSupplierParty"`

	// Party that is the accountable buyer of the goods/services in the referred
	// business document.
	AccountingCustomerParty AccountingCustomerParty `json:"accountingCustomerParty"`

	BuyerReference BuyerReference `json:"buyerReference,omitempty"`
	Delivery       Delivery       `json:"delivery,omitempty"`
	DeliveryParty  DeliveryParty  `json:"deliveryParty,omitempty"`

	// An array holding the Invoice lines.
	InvoiceLine []InvoiceLineItem `json:"invoiceLine"`

	LegalMonetaryTotal LegalMonetaryTotal `json:"legalMonetaryTotal"`
	PaymentMeans       PaymentMeans       `json:"paymentMeans,omitempty"`
	PaymentTerms       PaymentTerms       `json:"paymentTerms,omitempty"`
	TaxTotal           TaxTotal           `json:"taxTotal,omitempty"`
}

type AllowanceCharge struct {
	// Freight charge.
	Freight float64 `json:"freight,omitempty"`

	// Handling charge.
	Handling float64 `json:"handling,omitempty"`
}

type Delivery struct {
	ActualDeliveryDate string `json:"actualDeliveryDate,omitempty"`
}

// Party that is responsible for the delivery of the goods/services in the
// referred business document.
type DeliveryParty struct {
	Contact             Contact                   `json:"contact,omitempty"`
	Endpoint            Endpoint                  `json:"endpoint,omitempty"`
	PartyIdentification []PartyIdentificationItem `json:"partyIdentification,omitempty"`
	PartyName           string                    `json:"partyName"`
	PartyTaxScheme      PartyTaxScheme            `json:"partyTaxScheme,omitempty"`
	PostalAddress       PostalAddress             `json:"postalAddress,omitempty"`
}

// Description of the Business Document line item.
type DescriptionItem string

type FileRef struct {
	FileType string `json:"fileType"`
	RefID    string `json:"refId"`
}

// Object holding the financial account data
type FinancialAccountItem struct {
	// The name of financial institution.
	FinancialInstitutionName string `json:"financialInstitutionName,omitempty"`

	// Identifier of financial institution.
	FinancialInstitutionID string `json:"financialInstitutionId,omitempty"`

	// The external identifier of the financial institution id identifier
	// element.
	FinancialInstitutionIDSchemeID string `json:"financialInstitutionIdSchemeId,omitempty"`

	// The identifier of financial institution branch, for example 342-085. This
	// field is typically used by institutions in Australia and New Zealand.
	FinancialInstitutionBranchID string `json:"financialInstitutionBranchId,omitempty"`

	// The scheme identifier of financial institution branch. For example for an
	// Australian institutions, possible scheme is BSB.
	FinancialInstitutionBranchSchemeID string `json:"financialInstitutionBranchSchemeId,omitempty"`

	// Array holding ids
	Ids []ID `json:"ids,omitempty"`

	// Accounting related content.
	Accounting Accounting `json:"accounting,omitempty"`
}

// Object holding identifier data
type ID struct {
	// Identifier.
	ID string `json:"id"`

	// External identifier.
	SchemeID string `json:"schemeId,omitempty"`
}

// An object holding a Invoice line.
type InvoiceLineItem struct {
	// External system identifier for the Invoice line.
	ID string `json:"id"`

	// Internal identifier for the Invoice line.
	InternalID string `json:"internalId,omitempty"`

	Quantity Quantity `json:"quantity,omitempty"`

	// Flag indicating whether the line represents goods or services (true if
	// services, false if goods).
	ServiceIndicator bool `json:"serviceIndicator,omitempty"`

	LineExtension LineExtension `json:"lineExtension"`

	Item Item `json:"item"`

	TaxTotal []TaxTotalItem `json:"taxTotal,omitempty"`

	AllowanceCharge    InvoiceLineItemAllowanceCharge `json:"allowanceCharge,omitempty"`
	Delivery           InvoiceLineItemDelivery        `json:"delivery,omitempty"`
	OrderLineReference OrderLineReference             `json:"orderLineReference,omitempty"`
	Price              Price                          `json:"price,omitempty"`
}

type InvoiceLineItemAllowanceCharge struct {
	Amount                  float64 `json:"amount"`
	ChargeIndicator         bool    `json:"chargeIndicator"`
	MultiplierFactorNumeric float64 `json:"multiplierFactorNumeric,omitempty"`
}

type InvoiceLineItemDelivery struct {
	ActualDeliveryDate string `json:"actualDeliveryDate,omitempty"`
}

type Item struct {
	// An array holding the descriptions of the Business Document line items.
	Description []DescriptionItem `json:"description,omitempty"`

	// Name of the Business Document line item.  A short name optionally given
	// to an item, such as a name from a catalogue, as distinct from a
	// description.
	Name string `json:"name,omitempty"`

	// An object holding a identification of the Business Documents line item as
	// it is in sellers system.
	SellersItem SellersItem `json:"sellersItem,omitempty"`

	// Tax amount for the item
	TaxPercent float64 `json:"taxPercent,omitempty"`
}

type LegalMonetaryTotal struct {
	// An object holding total amount of line extensions.
	LineExtensionAmount Amount `json:"lineExtensionAmount,omitempty"`

	// An object holding total payable amount of line extensions.
	PayableAmount Amount `json:"payableAmount"`
}

type LineExtension struct {
	// The currency of the amount.
	CurrencyID string `json:"currencyId"`

	// The total amount for the line item, including allowance charges but net
	// of taxes.
	Amount float64 `json:"amount"`
}

// An object holding total amount of line extensions.
type LineExtensionAmount struct {
	Amount     float64 `json:"amount"`
	CurrencyID string  `json:"currencyId"`
}

// Links related to the business document
type Link struct {
	Href   string `json:"href,omitempty"`
	Method string `json:"method,omitempty"`
	Rel    string `json:"rel,omitempty"`
}

type OrderLineReference struct {
	LineID         string `json:"lineId"`
	OrderReference string `json:"orderReference,omitempty"`
}

type OrderReference struct {
	// Order number reference on the business document. Identifies the referenced
	// order assigned by the buyer.
	ID string `json:"id"`

	// External system specific identifier of the order system identifier element. If
	// the source business document has any matching element, it should be used.
	SchemeID string `json:"schemeId,omitempty"`

	// Customer Reference Identifier (CRI) when using a purchasing card.
	CustomerReference string `json:"customerReference,omitempty"`

	// Sales order identifier.
	SalesOrderID string `json:"salesOrderId,omitempty"`
}

// An object holding total payable amount of line extensions.
type Amount struct {
	// Total amount of line extensions.
	Amount float64 `json:"amount"`

	// A code that identifies the currency of the total line extension amount.
	// Valid values: ISO 4217 code represented as string.
	CurrencyID string `json:"currencyId"`
}

// An identifier for a payment made using this means of payment.
type PaymentIdentifier struct {
	ID       string `json:"id"`
	SchemeID string `json:"schemeId,omitempty"`
}

// An object holding the available payment means.
type PaymentMeans struct {
	// A code that identifies how the payment can be done. Valid values: UN/ECE
	// 4461 code represented as string.
	PaymentMeansCode string `json:"paymentMeansCode"`

	// Date when the business document is due for the payment means. Valid
	// values must be in format: CCYY-MM-DD. If the time zone is known, it must
	// be represented with +hh:mm or -hh:mm or Z (which means UTC). If time zone
	// is not known, it must be left empty.
	PaymentDueDate string `json:"paymentDueDate,omitempty"`

	// An identifier for a payment made using this means of payment.
	PaymentIdentifier PaymentIdentifier `json:"paymentIdentifier,omitempty"`

	// Array holding the financial account data
	FinancialAccount []FinancialAccountItem `json:"financialAccount,omitempty"`
}

type PaymentTerms struct {
	Note                    string           `json:"note,omitempty"`
	PenaltySurchargePercent float64          `json:"penaltySurchargePercent,omitempty"`
	SettlementPeriod        SettlementPeriod `json:"settlementPeriod,omitempty"`
}

type Price struct {
	Amount     float64 `json:"amount"`
	CurrencyID string  `json:"currencyId"`
}

type Quantity struct {
	// The quantity of the target Business Document line items.
	Amount float64 `json:"amount,omitempty"`

	// The available quantity of the target Business Document line item which
	// has not been invoiced.
	AmountUninvoiced float64 `json:"amountUninvoiced,omitempty"`

	// The unit code of the quantity of the target Business Document line item.
	// Valid values: UN/ECE CEFACT Trade Facilitation Recommendation No.20
	// common code value represented as string.
	UnitCode string `json:"unitCode,omitempty"`
}

// An object holding a identification of the Business Documents line item as it
// is in sellers system.
type SellersItem struct {
	ID       string `json:"id"`
	SchemeID string `json:"schemeId,omitempty"`
}

// An object holding the settlement period dates.
type SettlementPeriod struct {
	// Date when the payment terms starts. Valid values must be in format:
	// CCYY-MM-DD. If the time zone is known, it must be represented with +hh:mm
	// or -hh:mm or Z (which means UTC). If time zone is not known, it must be
	// left empty.
	StartDate string `json:"startDate,omitempty"`

	// Date when the payment terms ends. Valid values must be in format:
	// CCYY-MM-DD. If the time zone is known, it must be represented with +hh:mm
	// or -hh:mm or Z (which means UTC). If time zone is not known, it must be
	// left empty.
	EndDate string `json:"endDate,omitempty"`
}

type TaxTotal struct {
	// A code that identifies the currency of the total payable amount. Valid
	// values: ISO 4217 code represented as string.
	CurrencyID string `json:"currencyId"`

	// Total amount of the taxes. The total tax amount for particular tax scheme
	// e.g. VAT; the sum of each of the tax subtotals for each tax category
	// within the tax scheme.
	Amount float64 `json:"amount"`

	TaxSubTotal []TaxSubTotalItem `json:"taxSubTotal,omitempty"`
}

// An object holding the information about tax.
type TaxTotalItem struct {
	// Total amount of the taxes. The total tax amount for particular tax scheme
	// e.g. VAT; the sum of each of the tax subtotals for each tax category
	// within the tax scheme.
	Amount float64 `json:"amount"`

	// A code that identifies the currency of the total payable amount. Valid
	// values: ISO 4217 code represented as string.
	CurrencyID string `json:"currencyId"`

	// An object holding transaction tax.
	TransactionCurrencyTax TransactionCurrencyTax `json:"transactionCurrencyTax,omitempty"`

	TaxSubTotal []TaxSubTotalItem `json:"taxSubTotal,omitempty"`
}

// An object holding the information about tax.
type TaxSubTotalItem struct {
	// A code that identifies the currency of the tax subtotal. Valid values:
	// ISO 4217 code represented as string.
	CurrencyID string `json:"currencyId"`

	// Total amount of the taxes.
	Amount float64 `json:"amount"`

	// The tax rate for the category, expressed as a percentage.
	Percent float64 `json:"percent,omitempty"`

	// Basis of the taxes. The net amount to which the tax percent (rate) is
	// applied to calculate the tax amount.
	TaxableAmount float64 `json:"taxableAmount,omitempty"`
}

// An object holding transaction tax.
type TransactionCurrencyTax struct {
	Amount float64 `json:"amount"`
}

// Virtual bar code can be added to the business document that should be printed.
type VirtualBankBarcode struct {
	// Identifier of the virtual bar code.
	VirtualBankBarCode string `json:"id,omitempty"`

	// Scheme identifier of the virtual bank bar code, typically country code
	// according to ISO3166-1 alpha-2. Possible values: FI
	SchemeIDForVirtualBankBarCode string `json:"schemeId,omitempty"`
}

// An object containing address information. Used for defining supplier party,
// customer party and delivery party address data.
type PostalAddress struct {
	// The name of the city, town or village in the postal address of the party.
	CityName string `json:"cityName"`

	// The postal code of the area in the postal address of the party. The
	// identifier for an addressable group of properties according to the
	// relevant national postal service, such as a ZIP code or Post Code.
	PostalZOne string `json:"postalZone"`

	// The address line of the postal address of the party.
	AddressLine string `json:"addressLine"`

	// The second address line of the postal address of the party.
	AddressLine2 string `json:"addressLine2"`

	// Neighbourhood or district within town or city. Required in UK if a
	// similar road name exists within a post town area.
	Locality string `json:"locality"`

	// The sub-entity of the area in the postal address.
	CountrySubentity string `json:"countrySubentity"`

	// The country of the postal address of party. Valid values: ISO3166-1
	// alpha-2 values can be used.
	CountryID string `json:"countryId"`
}
