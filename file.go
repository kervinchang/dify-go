package dify

// File - Uploaded File
type File struct {
	Type           string `json:"type"`                     // Supported types, `image` only.
	TransferMethod string `json:"transfer_method"`          // Delivery method, `remote_url` or `local_file`.
	Url            string `json:"url,omitempty"`            // Image URL.
	UploadFileID   string `json:"upload_file_id,omitempty"` // Upload file ID.
}
