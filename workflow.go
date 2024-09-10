package dify

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// WorkflowEndpoint - Endpoint for workflows.
const WorkflowEndpoint = "/v1/workflows"

// RunWorkflowRequest - Request body for running a workflow.
type RunWorkflowRequest struct {
	Inputs       map[string]interface{} `json:"inputs"`        // User input/question content.
	ResponseMode ResponseMode           `json:"response_mode"` // Response mode, `streaming`(recommended) or `blocking`.
	User         string                 `json:"user"`          // Identity of the end user.
	Files        []File                 `json:"files"`         // Uploaded files.
}

// RunWorkflow - Runs a workflow in blocking mode.
func (c *Client) RunWorkflow(ctx context.Context, req RunWorkflowRequest) (*CompletionResponse, error) {
	url := fmt.Sprintf("%s%s/run", c.config.BaseURL, WorkflowEndpoint)

	req.ResponseMode = BlockingMode
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	request, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	request.Header.Set("Authorization", "Bearer "+c.config.APIKey)
	request.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		buf := &bytes.Buffer{}
		_, err = buf.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %v", err)
		}
		return nil, fmt.Errorf("unexpected response status %d: %s", resp.StatusCode, buf.String())
	}

	var response CompletionResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &response, nil
}

// RunWorkflowStream - Runs a workflow in streaming mode.
func (c *Client) RunWorkflowStream(ctx context.Context, req RunWorkflowRequest) (<-chan ChunkCompletionResponse, error) {
	url := fmt.Sprintf("%s%s/run", c.config.BaseURL, WorkflowEndpoint)

	req.ResponseMode = StreamingMode
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	request, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	request.Header.Set("Authorization", "Bearer "+c.config.APIKey)
	request.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		buf := &bytes.Buffer{}
		_, err = buf.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %v", err)
		}
		return nil, fmt.Errorf("unexpected response status %d: %s", resp.StatusCode, buf.String())
	}

	stream := make(chan ChunkCompletionResponse)
	go func() {
		defer close(stream)

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Bytes()

			if bytes.HasPrefix(line, []byte("data: ")) {
				data := bytes.TrimPrefix(line, []byte("data: "))

				var chunk ChunkCompletionResponse
				if err := json.Unmarshal(data, &chunk); err != nil {
					fmt.Printf("failed to unmarshal chunk: %v\n", err)
					return
				}

				stream <- chunk
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("error reading response body: %v\n", err)
		}
	}()

	return stream, nil
}
