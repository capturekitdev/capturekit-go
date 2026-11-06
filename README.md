# CaptureKit Go SDK

The official Go SDK for the [CaptureKit](https://www.capturekit.dev) screenshot API.

Capture websites as PNG, JPEG, WebP or PDF, extract their content, or analyse them with AI.

The SDK exposes **7 typed methods**, one per documented endpoint, plus a generic `Request()`. It has no dependencies outside the standard library.

## Requirements

- Go 1.22 or newer

## Installation

```sh
go get github.com/capturekitdev/capturekit-go
```

## Get your API key

[Sign up for CaptureKit](https://app.capturekit.dev/register) or sign in, then copy your API key from the [dashboard](https://app.capturekit.dev). `NewClient("", nil)` reads the `CAPTUREKIT_API_KEY` environment variable.

## Quick start

```go
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
	data, err := response.Map()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
```

Parameter structs mirror the [API reference](https://app.capturekit.dev/openapi). Required fields are plain values; optional fields are pointers — use `capturekit.String`, `capturekit.Float`, `capturekit.Int` and `capturekit.Bool`.

Decode into your own types with `response.Decode(&value)`.

## Binary responses

Endpoints that return a file (image, PDF…) give you the bytes in `Body`; `IsJSON` tells both cases apart:

```go
response, err := client.Capture(ctx, &capturekit.CaptureParams{URL: "https://example.com", Format: capturekit.String("png")})
if err != nil {
	log.Fatal(err)
}
if !response.IsJSON {
	err = os.WriteFile("screenshot.png", response.Body, 0o644)
}
```

## Async jobs

Any call accepts `is_async: true` (useful past the 60-second server limit): the response is then a `job_id`. `WaitForJob` polls `GET /v1/jobs/{id}` until a terminal status (`CompletedJob`, `Error`, `NotFound`, `BadRequest`):

```go
submitted, err := client.Content(ctx, &capturekit.ContentParams{URL: "https://example.com", IsAsync: capturekit.Bool(true)})
if err != nil {
	log.Fatal(err)
}
var accepted struct{ JobID string `json:"job_id"` }
if err := submitted.Decode(&accepted); err != nil {
	log.Fatal(err)
}
job, err := client.WaitForJob(ctx, accepted.JobID, &capturekit.WaitForJobOptions{
	OnProgress: func(job map[string]any) { fmt.Println(job["status"], job["progress"]) },
})
```

## GET or POST

Methods use GET when the endpoint supports it (parameters in the query string); `WithMethod("POST")` sends them as a JSON body.

```go
response, err := client.Analyze(ctx, &capturekit.AnalyzeParams{URL: "test"}, capturekit.WithMethod("POST"))
```

## Any endpoint

```go
response, err := client.Request(ctx, "GET", "/v1/usage", nil)
```

## Errors

```go
response, err := client.Content(ctx, &capturekit.ContentParams{URL: "https://example.com"})
if apiError, ok := capturekit.IsAPIError(err); ok {
	log.Printf("HTTP %d: %s", apiError.StatusCode, apiError.Message)
}
```

| Error | When |
| --- | --- |
| `*ConfigurationError` | Missing API key, invalid options or unsupported method. |
| `*SerializationError` | A required parameter is missing or a value can't be sent. |
| `*TransportError` | Network failure, timeout or cancelled context (`errors.Is(err, context.DeadlineExceeded)`). |
| `*APIError` | Non-2xx status (`StatusCode`, `Message`, `Response`). |
| `*ResponseDecodingError` | A success response could not be decoded. |

## Retries

Requests are retried up to 3 times after a network error, a `408`, `500`, `502`, `503` or `504`, or a rate limit — which the API returns as `401` (or `429`) with a `Rate limit exceeded` message. Per-key quota errors (`…: quota daily`) are not retried, nor are `400`, `402`, `404` or client-side timeouts. The delay starts at 1 second and grows 1.5× per retry (capped at 30 seconds); a `Retry-After` header wins.

```go
client, err := capturekit.NewClient("", &capturekit.ClientOptions{
	MaxRetries: 5,                      // negative disables retries
	RetryDelay: 500 * time.Millisecond,
})

// or for one request
response, err := client.Content(ctx, &capturekit.ContentParams{URL: "https://example.com"}, capturekit.WithMaxRetries(0))
```

## Batch requests

`capturekit.Batch` runs a function per item with at most 10 calls in flight and returns one `BatchResult` per item, in input order. A failing item never stops the others.

```go
inputs := []string{"https://example.com", "https://www.capturekit.dev", "https://github.com"}
results := capturekit.Batch(ctx, inputs, 5, func(ctx context.Context, value string) (*capturekit.Response, error) {
	return client.Content(ctx, &capturekit.ContentParams{URL: value})
})
for _, result := range results {
	if result.Err != nil {
		log.Printf("%s: %v", inputs[result.Index], result.Err)
		continue
	}
	fmt.Println(string(result.Value.Body))
}
```

| Constant | Default | Description |
| --- | --- | --- |
| `DefaultConcurrency` | `10` | Tasks run at the same time by batch |
| `DefaultTimeout` | `120 * time.Second` | Request timeout |
| `MaxRetries` | `3` | Retry attempts after a network error, a 408/5xx or a rate limit |
| `RetryBackoffFactor` | `1.5` | Exponential backoff multiplier |
| `DefaultRetryDelay` | `time.Second` | Delay before the first retry |

## Timeouts and custom HTTP clients

```go
client, err := capturekit.NewClient("", &capturekit.ClientOptions{
	Timeout:    60 * time.Second, // default 120 s; negative disables
	HTTPClient: &http.Client{Transport: myTransport},
})
```

Every method also honours the deadline and cancellation of the `context.Context` you pass.

## Endpoints

7 methods, generated from the [OpenAPI document](https://app.capturekit.dev/openapi).

| Method | HTTP | Description |
| --- | --- | --- |
| `Analyze` | GET / POST `/v1/analyze` | AI Content API analyzes any webpage and returns AI insights summary, value proposition, features, categories, use cases, and optional conta… |
| `Capture` | GET / POST `/v1/capture` | Browser-like Screenshot API that captures webpages as PNG/JPEG/WebP or PDF for previews, QA, reporting, monitoring, and archiving no headle… |
| `Content` | GET / POST `/v1/content` | CaptureKit Content API extracts structured webpage data, metadata, links, HTML/Markdown: in one request for SEO, monitoring, indexing and a… |
| `Jobs` | GET `/v1/jobs/{id}` | Poll the status, progress (0–100) and result of a job started with is_async=true. Polling is free. Poll until status is terminal: Completed… |
| `Stats` | GET `/v1/stats` | Infrastructure health for an engine or endpoint over the last hour. Not scoped to your account or tenant. Enabled per account: HTTP 403 mea… |
| `Usage` | GET `/v1/usage` | Retrieve real-time usage statistics and credit balance for your CaptureKit account. |
| `WebsiteCrawler` | GET / POST `/v1/website/crawler` | Website Crawler API fetches raw HTML from any webpage in request mode (no rendering) for fast scraping, audits, monitoring, and link/metada… |

## License

MIT — see [LICENSE](LICENSE).
