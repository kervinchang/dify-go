package dify

// Metadata - Message metadata.
type Metadata struct {
	Usage              Usage               `json:"usage"`                         // Model usage information.
	RetrieverResources []RetrieverResource `json:"retriever_resources,omitempty"` // List of references and attributed segments.
}
