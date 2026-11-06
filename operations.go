// Code generated from the CaptureKit OpenAPI document. DO NOT EDIT.

package capturekit

import "context"

// Analyze — Analyze (AI). AI Content API analyzes any webpage and returns AI insights summary, value proposition, features, categories, use cases, and optional contact signals. Credit cost: 1 credit per call.
//
// GET / POST /v1/analyze
func (c *Client) Analyze(ctx context.Context, params *AnalyzeParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["analyze"], params.params(), options)
}

// Capture — Capture (Screenshot). Browser-like Screenshot API that captures webpages as PNG/JPEG/WebP or PDF for previews, QA, reporting, monitoring, and archiving no headless setup. Credit cost: 1 credit per call.
//
// GET / POST /v1/capture
func (c *Client) Capture(ctx context.Context, params *CaptureParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["capture"], params.params(), options)
}

// Content — Content. CaptureKit Content API extracts structured webpage data, metadata, links, HTML/Markdown: in one request for SEO, monitoring, indexing and automation. Credit cost: 1 credit per call.
//
// GET / POST /v1/content
func (c *Client) Content(ctx context.Context, params *ContentParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["content"], params.params(), options)
}

// Jobs — Jobs. Poll the status, progress (0–100) and result of a job started with is_async=true. Polling is free. Poll until status is terminal: CompletedJob, Error, NotFound or BadRequest.
//
// GET /v1/jobs/{id}
func (c *Client) Jobs(ctx context.Context, params *JobsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["jobs"], params.params(), options)
}

// Stats — Stats. Infrastructure health for an engine or endpoint over the last hour. Not scoped to your account or tenant. Enabled per account: HTTP 403 means contact support to activate it.
//
// GET /v1/stats
func (c *Client) Stats(ctx context.Context, params *StatsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["stats"], params.params(), options)
}

// Usage — Usage. Retrieve real-time usage statistics and credit balance for your CaptureKit account.
//
// GET /v1/usage
func (c *Client) Usage(ctx context.Context, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["usage"], nil, options)
}

// WebsiteCrawler — Website Crawler. Website Crawler API fetches raw HTML from any webpage in request mode (no rendering) for fast scraping, audits, monitoring, and link/metadata extraction. Credit cost: 1 credit per call.
//
// GET / POST /v1/website/crawler
func (c *Client) WebsiteCrawler(ctx context.Context, params *WebsiteCrawlerParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["website_crawler"], params.params(), options)
}

// endpoints lists every documented operation by its snake_case name.
var endpoints = map[string]endpoint{
	"analyze":         {path: "/v1/analyze", methods: []string{"GET", "POST"}, json: true, required: []string{"url"}},
	"capture":         {path: "/v1/capture", methods: []string{"GET", "POST"}, json: false, required: []string{"url"}},
	"content":         {path: "/v1/content", methods: []string{"GET", "POST"}, json: true, required: []string{"url"}},
	"jobs":            {path: "/v1/jobs/{id}", methods: []string{"GET"}, json: true, required: []string{"id"}},
	"stats":           {path: "/v1/stats", methods: []string{"GET"}, json: true, required: []string{}},
	"usage":           {path: "/v1/usage", methods: []string{"GET"}, json: true, required: []string{}},
	"website_crawler": {path: "/v1/website/crawler", methods: []string{"GET", "POST"}, json: false, required: []string{"query"}},
}
