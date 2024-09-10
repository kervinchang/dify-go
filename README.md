# dify-go

Go SDK for langgenius/dify.

### Quick Start
Install the Dify Go package:
```bash
go get github.com/kervinchang/dify-go
```
Initialize a Dify client:
```go
config := dify.ClientConfig{
	BaseURL: "https://your-dify-server-endpoint.com",
	APIKey: "your-api-key",
}
client, err := dify.NewClient(config)
if err != nil {
	log.Fatalf("failed to create Dify client: %v\n", err)
}
```
Send a request to the CreateChatMessage API:
```go
request := dify.ChatMessageRequest{
	Inputs:        make(map[string]interface{}),
	query:         query,
	User:          "your-user-id",
}
response, err := client.CreateChatMessage(ctx, request)
if err != nil {
	log.Fatalf("failed to create chat message in blocking mode: %v\n", err)
}
log.Printf("response: %v\n", response) // dify.ChatCompletionResponse
```
Send a request to the CreateChatMessage API for streaming responses:
```go
request := dify.ChatMessageRequest{
	Inputs:       make(map[string]interface{}),
	query:        query,
	User:         "your-user-id",
}
stream, err := client.CreateChatMessageStream(ctx, request)
if err != nil {
	log.Fatalf("failed to create chat message in streaming mode: %v\n", err)
}
for reponse := range stream {
	log.Printf("response: %v\n", response) // dify.ChunkChatCompletionResponse
}
```
Send a request to the CreateCompletionMessage API:
```go
request := dify.CompletionMessageRequest{
	Inputs:       map[string]interface{}{"query": query}, 
	User:         "your-user-id",
}
response, err := client.CreateCompletionMessage(ctx, request)
if err != nil {
	log.Fatalf("failed to create completion message in blocking mode: %v\n", err)
}
log.Printf("response: %v\n", response) // dify.ChatCompletionResponse
```
Send a request to the CreateCompletionMessage API for streaming responses:
```go
request := dify.CompletionMessageRequest{
	Inputs:       map[string]interface{}{"query": query},
	User:         "your-user-id",
}
stream, err := client.CreateCompletionMessageStream(ctx, request)
if err != nil {
	log.Fatalf("failed to create completion message in streaming mode: %v\n", err)
}
for response := range stream {
	log.Printf("response: %v\n", response) // dify.ChunkChatCompletionResponse
}
```
Send a requrest to the RunWorkflow API:
```go
request := dify.RunWorkflowRequest{
	Inputs:       map[string]interface{}{}, 
	User:         "your-user-id",
}
response, err := client.RunWorkflow(ctx, request)
if err != nil {
	log.Fatalf("failed to run workflow in blocking mode: %v\n", err)
}
log.Printf("response: %v\n", response) // dify.CompletionResponse
```
Send a request to the RunWorkflow API for streaming responses:
```go
request := dify.RunWorkflowRequest{
	Inputs:       map[string]interface{}{}, 
	User:         "your-user-id",
}
stream, err := client.RunWorkflowStream(ctx, request)
if err != nil {
	log.Fatalf("failed to run workflow in streaming mode: %v\n", err)
}
for response := range stream {
	log.Printf("response: %v\n", response) // dify.ChunkCompletionResponse
}
```