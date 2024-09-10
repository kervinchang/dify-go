package dify

// RetrieverResource - Message references and attributed segments.
type RetrieverResource struct {
	Position     int     `json:"position"`      // Position of the reference in the message.
	DatasetID    string  `json:"dataset_id"`    // ID of the dataset.
	DatasetName  string  `json:"dataset_name"`  // Name of the dataset.
	DocumentID   string  `json:"document_id"`   // ID of the document.
	DocumentName string  `json:"document_name"` // Name of the document.
	SegmentID    string  `json:"segment_id"`    // ID of the segment.
	Score        float64 `json:"score"`         // Score of the segment.
	Content      string  `json:"content"`       // Content of the segment.
}
