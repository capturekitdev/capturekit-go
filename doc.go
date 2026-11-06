// Package capturekit is the official Go SDK for the CaptureKit API (https://www.capturekit.dev).
//
// Capture websites as PNG, JPEG, WebP or PDF, extract their content, or analyse them with AI.
//
// Create one client and reuse it:
//
//	client, err := capturekit.NewClient("", nil) // reads CAPTUREKIT_API_KEY
//	response, err := client.Content(ctx, ...)
//
// Every documented endpoint has a typed method; Client.Request reaches any other path.
package capturekit
