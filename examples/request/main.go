package main

import (
	"context"
	"fmt"
	"log"

	"github.com/capturekitdev/capturekit-go"
)

func main() {
	client, err := capturekit.NewClient("", nil) // reads CAPTUREKIT_API_KEY
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	response, err := client.Request(ctx, "GET", "/v1/usage", nil)
	if apiError, ok := capturekit.IsAPIError(err); ok {
		log.Fatalf("HTTP %d: %s", apiError.StatusCode, apiError.Message)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(response.Body))
}
