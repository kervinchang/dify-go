package dify

const (
	BlockingMode  ResponseMode = "blocking"  // Blocking response.
	StreamingMode ResponseMode = "streaming" // Streaming response.
)

// ResponseMode - Response mode, `streaming` or `blocking`.
type ResponseMode string

// ChatCompletionResponse - Response body from the CreateChatMessage or CreateCompletionMessage endpoint in blocking mode.
type ChatCompletionResponse struct {
	ID             string   `json:"id,omitempty"`              // Agent thought ID. Each Agent iteration will have a unique id.
	Event          string   `json:"event,omitempty"`           // SSE event name.
	MessageID      string   `json:"message_id,omitempty"`      // Message unique ID.
	ConversationID string   `json:"conversation_id,omitempty"` // Session ID.
	Mode           string   `json:"mode,omitempty"`            // App mode, `chat` or `completion`.
	Answer         string   `json:"answer"`                    // Full response content.
	Metadata       Metadata `json:"metadata,omitempty"`        // Message metadata.
	CreatedAt      int      `json:"created_at"`                // Message creation timestamp, such as: 1705395332.
}

// ChunkChatCompletionResponse - Response body from the CreateChatMessageStream or CreateCompletionMessageStream endpoint in streaming mode.
type ChunkChatCompletionResponse struct {
	ID             string   `json:"id,omitempty"`              // Agent thought ID. Each Agent iteration will have a unique id.
	Event          string   `json:"event,omitempty"`           // SSE event name.
	TaskID         string   `json:"task_id,omitempty"`         // Task ID, used for request tracking and the stop response interface below.
	MessageID      string   `json:"message_id,omitempty"`      // Message unique ID.
	ConversationID string   `json:"conversation_id,omitempty"` // Session ID.
	Answer         string   `json:"answer,omitempty"`          // Block response content.
	Position       int      `json:"position,omitempty"`        // The position of agent_thought in the message, such as the first iteration position is 1.
	Thought        string   `json:"thought,omitempty"`         // The agent's thoughts.
	Observation    string   `json:"observation,omitempty"`     // The result returned by the tool call.
	Tool           string   `json:"tool,omitempty"`            // List of tools to use, separate multiple tools with.
	ToolInput      string   `json:"tool_input,omitempty"`      // Tool input, a string in JSON format (object).
	MessageFiles   []string `json:"message_files,omitempty"`   // Current agent_thought associated file ID.
	Type           string   `json:"type,omitempty"`            // File type, currently only image.
	BelongsTo      string   `json:"belongs_to,omitempty"`      // File owner, user or assistant. This interface returns only assistant.
	URL            string   `json:"url,omitempty"`             // File access address.
	Data           struct {
		ID                string         `json:"id,omitempty"`                  // Workflow execution ID.
		NodeID            string         `json:"node_id,omitempty"`             // Workflow node ID.
		NodeType          string         `json:"node_type,omitempty"`           // Workflow node type.
		Title             string         `json:"title,omitempty"`               // Workflow node title.
		Index             int            `json:"index,omitempty"`               // Workflow node index.
		PredecessorNodeID string         `json:"predecessor_node_id,omitempty"` // Workflow node predecessor node ID.
		Inputs            map[string]any `json:"inputs,omitempty"`              // Input content.
		Outputs           map[string]any `json:"outputs,omitempty"`             // Output content.
		Status            string         `json:"status,omitempty"`              // Execution status.
		ElapsedTime       float64        `json:"elapsed_time,omitempty"`        // Execution time.
		CreatedAt         int            `json:"created_at,omitempty"`          // Start time.
		FinishedAt        int            `json:"finished_at,omitempty"`         // End time.
	} `json:"data,omitempty"` // Details.
	Metadata  Metadata `json:"metadata,omitempty"`   // Message metadata.
	Audio     string   `json:"audio,omitempty"`      // The audio block after speech synthesis is encoded with Base64 text content.
	Status    int      `json:"status,omitempty"`     // HTTP status code.
	Code      string   `json:"code,omitempty"`       // Error code.
	Message   string   `json:"message,omitempty"`    // Error message.
	CreatedAt int      `json:"created_at,omitempty"` // Message creation timestamp, such as: 1705395332.
}

// CompletionResponse - Response body from the RunWorkflow endpoint in blocking mode.
type CompletionResponse struct {
	WorkflowRunID string `json:"workflow_run_id"` // Workflow execution ID.
	TaskID        string `json:"task_id"`         // Task ID, used for request tracking and the stop response interface below.
	Data          struct {
		ID          string                 `json:"id"`           // Workflow execution ID.
		WorkflowID  string                 `json:"workflow_id"`  // Associated Workflow ID.
		Status      string                 `json:"status"`       // Execution status , running/succeeded/failed/stopped.
		Outputs     map[string]interface{} `json:"outputs"`      // Optional Output content.
		Error       string                 `json:"error"`        // Optional The reason for the error.
		ElapsedTime float64                `json:"elapsed_time"` // Optional time consumed (s).
		TotalTokens int                    `json:"total_tokens"` // Optional Total tokens used.
		TotalSteps  int                    `json:"total_steps"`  // Total number of steps (redundant), default 0.
		CreatedAt   int                    `json:"created_at"`   // Start time.
		FinishedAt  int                    `json:"finished_at"`  // End time.
	} `json:"data"` // Details.
}

// ChunkCompletionResponse - Response body from the RunWorkflowStream endpoint in streaming mode.
type ChunkCompletionResponse struct {
	TaskID        string `json:"task_id"`                   // Task ID, used for request tracking and the stop response interface below.
	WorkflowRunID string `json:"workflow_run_id,omitempty"` // Workflow execution ID.
	MessageID     string `json:"message_id,omitempty"`      // Message unique ID.
	Event         string `json:"event,omitempty"`           // Fixed to workflow_started.
	Audio         string `json:"audio,omitempty"`           // The audio block after speech synthesis is encoded with Base64 text content.
	Data          struct {
		ID                string                   `json:"id"`                            // Workflow execution ID or Node execution ID.
		WorkflowID        string                   `json:"workflow_id,omitempty"`         // Associated Workflow ID.
		SequenceNumber    int                      `json:"sequence_number,omitempty"`     // Self-incrementing sequence number, self-incrementing within the App, starting from 1.
		NodeID            string                   `json:"node_id,omitempty"`             // Node ID.
		NodeType          string                   `json:"node_type,omitempty"`           // Node type, such as: chat, completion, tool, etc.
		Title             string                   `json:"title,omitempty"`               // Node name.
		Index             int                      `json:"index,omitempty"`               // Execution sequence number, used to display the Tracing Node sequence.
		PredecessorNodeID string                   `json:"predecessor_node_id,omitempty"` // Prefix node ID, used to display the execution path on the canvas.
		Inputs            []map[string]interface{} `json:"inputs,omitempty"`              // All the previous node variables used in the node.
		ProcessData       map[string]interface{}   `json:"process_data,omitempty"`        // Optional Node process data.
		Outputs           map[string]interface{}   `json:"outputs,omitempty"`             // Optional Output content.
		Status            string                   `json:"status,omitempty"`              // Execution status running/succeeded/failed/stopped.
		Error             string                   `json:"error,omitempty"`               // Optional The reason for the error.
		ElapsedTime       float64                  `json:"elapsed_time,omitempty"`        // Optional time consumed (s).
		ExecutionMetadata map[string]interface{}   `json:"execution_metadata,omitempty"`  // Metadata.
		TotalTokens       int                      `json:"total_tokens,omitempty"`        // Optional Total tokens used.
		TotalPrice        float64                  `json:"total_price,omitempty"`         // Optional Total cost.
		Currency          string                   `json:"currency,omitempty"`            // Currency, such as USD/RMB.
		TotalSteps        int                      `json:"total_steps,omitempty"`         // Total number of steps (redundant), default 0.
		CreatedAt         int                      `json:"created_at"`                    // Start time.
		FinishedAt        int                      `json:"finished_at,omitempty"`         // End time.
	} `json:"data,omitempty"` // Details.
	CreatedAt int `json:"created_at,omitempty"` // Creation timestamp, such as: 1705395332.
}
