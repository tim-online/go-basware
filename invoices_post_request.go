package basware

// Invoice is a business document which can contain attachments.
type InvoicesPostRequestBody struct {
	// Token generated by client (uuid). Used to verify that specific Invoice is
	// only sent and processed once, if response time-outs, retry should be
	// executed with the same clientToken.
	ClientToken string `json:"clientToken"`

	// Object holding the business content of the Invoice. Content is at some
	// level based on Universal Business Language (UBL) standard version 2.1. It
	// has also been extended by Basware so it is not strictly UBL.
	Data Data `json:"data"`

	// The way document to be routed, printing-always goes for printing as
	// sender specific processing, only-eInvoicing goes for normal processing as
	// receiver specific processing, printing-allowed goes first for
	// only-eInvoicing if fails then for printing-always, empty value goes
	// by-default for only-eInvoicing case
	DeliveryChannelPreference string `json:"deliveryChannelPreference,omitempty"`

	// Invoice file/attachment reference identifiers.
	FileRefs []FileRef `json:"fileRefs,omitempty"`

	// Identifier for the intermediate service provider.
	ServiceProviderID string `json:"serviceProviderId,omitempty"`
}