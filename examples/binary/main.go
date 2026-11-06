package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/capturekitdev/capturekit-go"
)

func main() {
	client, err := capturekit.NewClient("", nil) // reads CAPTUREKIT_API_KEY
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	response, err := client.Capture(ctx, &capturekit.CaptureParams{URL: "https://example.com", Format: capturekit.String("png")})
	if err != nil {
		log.Fatal(err)
	}
	if response.IsJSON {
		fmt.Println(string(response.Body))
		return
	}
	if err := os.WriteFile("screenshot.png", response.Body, 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Saved screenshot.png (%s, %d bytes)\n", response.ContentType, len(response.Body))
}
