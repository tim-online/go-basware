package basware

// Invoice is a business document which can contain attachments.
type InvoicesGetResponse struct {
	Data     Invoice   `json:"data"`
	FileRefs []FileRef `json:"fileRefs,omitempty"`
	Links    []Link    `json:"links,omitempty"`
	Version  string    `json:"version"`
}
