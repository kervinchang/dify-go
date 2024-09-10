package dify

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// CompletionMessageEndpoint - Endpoint for creating a completion message.
const CompletionMessageEndpoint = "/v1/completion-messages"

// CompletionMessageRequest - Request body for creating a completion message.
type CompletionMessageRequest struct {
	Inputs       map[string]interface{} `json:"inputs"`        // User input/question content.
	ResponseMode ResponseMode           `json:"response_mode"` // Response mode, `streaming`(recommended) or `blocking`.
	User         string                 `json:"user"`          // Identity of the end user.
	Files        []File                 `json:"files"`         // Uploaded files.
}

// CreateCompletionMessage - Creates a completion message in blocking mode.
func (c *Client) CreateCompletionMessage(ctx context.Context, req CompletionMessageRequest) (*ChatCompletionResponse, error) {
	url := fmt.Sprintf("%s%s", c.config.BaseURL, CompletionMessageEndpoint)

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

	var response ChatCompletionResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &response, nil
}

// CreateCompletionMessageStream - Creates a completion message in streaming mode.
func (c *Client) CreateCompletionMessageStream(ctx context.Context, req CompletionMessageRequest) (<-chan ChunkChatCompletionResponse, error) {
	url := fmt.Sprintf("%s%s", c.config.BaseURL, CompletionMessageEndpoint)

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

	if resp.StatusCode != http.StatusOK {
		buf := &bytes.Buffer{}
		_, err = buf.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %v", err)
		}
		return nil, fmt.Errorf("unexpected response status %d: %s", resp.StatusCode, buf.String())
	}

	stream := make(chan ChunkChatCompletionResponse)
	go func() {
		defer resp.Body.Close()
		defer close(stream)

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Bytes()

			if bytes.HasPrefix(line, []byte("data: ")) {
				data := bytes.TrimPrefix(line, []byte("data: "))

				var chunk ChunkChatCompletionResponse
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
