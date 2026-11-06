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
	response, err := client.Content(ctx, &capturekit.ContentParams{URL: "https://example.com"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(response.Body))
}
