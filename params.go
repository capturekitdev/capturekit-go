// Code generated from the CaptureKit OpenAPI document. DO NOT EDIT.

package capturekit

// AnalyzeParams holds the parameters of Analyze (GET / POST /v1/analyze).
// Required fields are plain values; optional fields are pointers (use capturekit.String, capturekit.Float, capturekit.Bool) or nil slices and maps.
type AnalyzeParams struct {
	// The URL of the webpage to capture.
	URL string `json:"url" api:"url,required"`
	// Pass custom response fields to be included in the AI response.
	CustomResponse *string `json:"custom_response" api:"custom_response"`
	// Add your own instructions to guide the AI’s analysis.
	CustomPrompt *string `json:"custom_prompt" api:"custom_prompt"`
	// Specify a proxy server to route your request through.
	Proxy *string `json:"proxy" api:"proxy"`
	// Cache the response. Default: false.
	Cache *bool `json:"cache" api:"cache"`
	// Cache the response for a custom TTL (in seconds). Default: 2592000.
	CacheTTL *float64 `json:"cache_ttl" api:"cache_ttl"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AnalyzeParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 9)
	if p.URL != "" {
		values = append(values, param{"url", p.URL})
	}
	if p.CustomResponse != nil {
		values = append(values, param{"custom_response", *p.CustomResponse})
	}
	if p.CustomPrompt != nil {
		values = append(values, param{"custom_prompt", *p.CustomPrompt})
	}
	if p.Proxy != nil {
		values = append(values, param{"proxy", *p.Proxy})
	}
	if p.Cache != nil {
		values = append(values, param{"cache", *p.Cache})
	}
	if p.CacheTTL != nil {
		values = append(values, param{"cache_ttl", *p.CacheTTL})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CaptureParams holds the parameters of Capture (GET / POST /v1/capture).
// Required fields are plain values; optional fields are pointers (use capturekit.String, capturekit.Float, capturekit.Bool) or nil slices and maps.
type CaptureParams struct {
	// The URL of the webpage to capture. Examples:  https://www.amazon.com/s?k=pokemon
	URL string `json:"url" api:"url,required"`
	// The output format of the screenshot. Allowed values: jpeg, jpg, pdf, png, webp. Default: "png".
	Format *string `json:"format" api:"format"`
	// The device type to emulate for the screenshot. Allowed values: galaxy_fold4, galaxy_s23, galaxy_s23_ultra, galaxy_s8, huawei_mate_50_pro, huawei_p60_pro, ipad, iphone_12, iphone_13_mini, iphone_13_pro_max, iphone_14_pro, iphone_14_pro_max, iphone_x, pixel_5, pixel_6a, pixel_7_pro, redmi_note_11, re…
	Device *string `json:"device" api:"device"`
	// Capture the entire page instead of just the visible viewport. Default: false.
	FullPage *bool `json:"full_page" api:"full_page"`
	// Scroll the page to fully load lazy-loaded elements before capturing a full-page screenshot. Default: false.
	FullPageScroll *bool `json:"full_page_scroll" api:"full_page_scroll"`
	// Time in milliseconds for scrolling before capturing the full page. Default: 400.
	FullPageScrollDuration *float64 `json:"full_page_scroll_duration" api:"full_page_scroll_duration"`
	// The width of the browser viewport in pixels. Default: 1280.
	ViewportWidth *float64 `json:"viewport_width" api:"viewport_width"`
	// The height of the browser viewport in pixels. Default: 1024.
	ViewportHeight *float64 `json:"viewport_height" api:"viewport_height"`
	// The scale factor for high-resolution screenshots. Default: 1.
	ScaleFactor *float64 `json:"scale_factor" api:"scale_factor"`
	// Generate images with custom quality settings. Default: "80".
	ImageQuality *string `json:"image_quality" api:"image_quality"`
	// Capture a specific element on the page instead of the full viewport.
	Selector *string `json:"selector" api:"selector"`
	// Cache the response. Default: false.
	Cache *bool `json:"cache" api:"cache"`
	// Cache the response for a custom TTL (in seconds). Default: 2592000.
	CacheTTL *float64 `json:"cache_ttl" api:"cache_ttl"`
	// A comma-separated list of elements to hide before capturing (e.g., ads, popups).
	RemoveSelectors *string `json:"remove_selectors" api:"remove_selectors"`
	// Automatically remove cookie banners before capturing. Default: false.
	RemoveCookieBanners *bool `json:"remove_cookie_banners" api:"remove_cookie_banners"`
	// Automatically remove ads before capturing. Default: false.
	RemoveAds *bool `json:"remove_ads" api:"remove_ads"`
	// Comma-separated list of resource types to block (e.g., “image,stylesheet,font”). Allowed values: document, eventsource, fetch, font, image, manifest, media, other, script, stylesheet, texttrack, websocket, xhr.
	BlockResources *string `json:"block_resources" api:"block_resources"`
	// Comma-separated list of URL patterns to block (e.g., “analytics,tracking,advertisement”).
	BlockUrls *string `json:"block_urls" api:"block_urls"`
	// Delay in seconds before capturing the screenshot. Default: 0.
	Delay *float64 `json:"delay" api:"delay"`
	// Define when to capture. Allowed values: domcontentloaded, load, networkidle0, networkidle2.
	WaitUntil []string `json:"wait_until" api:"wait_until"`
	// Wait for a specific element to appear before taking the screenshot.
	WaitForSelector *string `json:"wait_for_selector" api:"wait_for_selector"`
	// Specify a proxy server to route your request through.
	Proxy *string `json:"proxy" api:"proxy"`
	// AWS S3 secret access key.
	S3SecretKey *string `json:"s3_secret_key" api:"s3_secret_key"`
	// AWS S3 access key ID.
	S3AccessKeyID *string `json:"s3_access_key_id" api:"s3_access_key_id"`
	// The S3 bucket name where the screenshot will be uploaded.
	S3Bucket *string `json:"s3_bucket" api:"s3_bucket"`
	// The filename for the S3 object (auto-generated if not provided).
	S3ObjectKey *string `json:"s3_object_key" api:"s3_object_key"`
	// AWS S3 region for storage.
	S3Region *string `json:"s3_region" api:"s3_region"`
	// Leave this field empty for AWS S3. Only specify a value if required. S3-compatible storage services are supported, for example, use https://<accountId>.r2.cloudflarestorage.com for Cloudflare R2.
	StorageEndpoint *string `json:"storage_endpoint" api:"storage_endpoint"`
	// Use S3 storage for screenshots. The S3 URL will be returned in the response. Default: false.
	S3URL *bool `json:"s3_url" api:"s3_url"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CaptureParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 32)
	if p.URL != "" {
		values = append(values, param{"url", p.URL})
	}
	if p.Format != nil {
		values = append(values, param{"format", *p.Format})
	}
	if p.Device != nil {
		values = append(values, param{"device", *p.Device})
	}
	if p.FullPage != nil {
		values = append(values, param{"full_page", *p.FullPage})
	}
	if p.FullPageScroll != nil {
		values = append(values, param{"full_page_scroll", *p.FullPageScroll})
	}
	if p.FullPageScrollDuration != nil {
		values = append(values, param{"full_page_scroll_duration", *p.FullPageScrollDuration})
	}
	if p.ViewportWidth != nil {
		values = append(values, param{"viewport_width", *p.ViewportWidth})
	}
	if p.ViewportHeight != nil {
		values = append(values, param{"viewport_height", *p.ViewportHeight})
	}
	if p.ScaleFactor != nil {
		values = append(values, param{"scale_factor", *p.ScaleFactor})
	}
	if p.ImageQuality != nil {
		values = append(values, param{"image_quality", *p.ImageQuality})
	}
	if p.Selector != nil {
		values = append(values, param{"selector", *p.Selector})
	}
	if p.Cache != nil {
		values = append(values, param{"cache", *p.Cache})
	}
	if p.CacheTTL != nil {
		values = append(values, param{"cache_ttl", *p.CacheTTL})
	}
	if p.RemoveSelectors != nil {
		values = append(values, param{"remove_selectors", *p.RemoveSelectors})
	}
	if p.RemoveCookieBanners != nil {
		values = append(values, param{"remove_cookie_banners", *p.RemoveCookieBanners})
	}
	if p.RemoveAds != nil {
		values = append(values, param{"remove_ads", *p.RemoveAds})
	}
	if p.BlockResources != nil {
		values = append(values, param{"block_resources", *p.BlockResources})
	}
	if p.BlockUrls != nil {
		values = append(values, param{"block_urls", *p.BlockUrls})
	}
	if p.Delay != nil {
		values = append(values, param{"delay", *p.Delay})
	}
	if p.WaitUntil != nil {
		values = append(values, param{"wait_until", p.WaitUntil})
	}
	if p.WaitForSelector != nil {
		values = append(values, param{"wait_for_selector", *p.WaitForSelector})
	}
	if p.Proxy != nil {
		values = append(values, param{"proxy", *p.Proxy})
	}
	if p.S3SecretKey != nil {
		values = append(values, param{"s3_secret_key", *p.S3SecretKey})
	}
	if p.S3AccessKeyID != nil {
		values = append(values, param{"s3_access_key_id", *p.S3AccessKeyID})
	}
	if p.S3Bucket != nil {
		values = append(values, param{"s3_bucket", *p.S3Bucket})
	}
	if p.S3ObjectKey != nil {
		values = append(values, param{"s3_object_key", *p.S3ObjectKey})
	}
	if p.S3Region != nil {
		values = append(values, param{"s3_region", *p.S3Region})
	}
	if p.StorageEndpoint != nil {
		values = append(values, param{"storage_endpoint", *p.StorageEndpoint})
	}
	if p.S3URL != nil {
		values = append(values, param{"s3_url", *p.S3URL})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ContentParams holds the parameters of Content (GET / POST /v1/content).
// Required fields are plain values; optional fields are pointers (use capturekit.String, capturekit.Float, capturekit.Bool) or nil slices and maps.
type ContentParams struct {
	// The URL of the webpage to capture.
	URL string `json:"url" api:"url,required"`
	// Include the raw HTML of the webpage in the response. Default: false.
	IncludeHTML *bool `json:"include_html" api:"include_html"`
	// Include the Markdown of the webpage in the response. Default: false.
	IncludeMarkdown *bool `json:"include_markdown" api:"include_markdown"`
	// Include the sitemap of the website in the response. Default: false.
	IncludeSitemap *bool `json:"include_sitemap" api:"include_sitemap"`
	// Use Defuddle to clean and extract the main content from web pages. This popular library removes unnecessary elements like comments, sidebars, headers, footers, and other non-essential elements, leaving only the primary content. When enabled, the HTML response will be processed through Defuddle befo…
	UseDefuddle *bool `json:"use_defuddle" api:"use_defuddle"`
	// A comma-separated list of elements to hide before capturing (e.g., ads, popups).
	RemoveSelectors *string `json:"remove_selectors" api:"remove_selectors"`
	// Automatically remove cookie banners before capturing. Default: false.
	RemoveCookieBanners *bool `json:"remove_cookie_banners" api:"remove_cookie_banners"`
	// Automatically remove ads before capturing. Default: false.
	RemoveAds *bool `json:"remove_ads" api:"remove_ads"`
	// Comma-separated list of resource types to block (e.g., “image,stylesheet,font”). Allowed values: document, eventsource, fetch, font, image, manifest, media, other, script, stylesheet, texttrack, websocket, xhr.
	BlockResources *string `json:"block_resources" api:"block_resources"`
	// Comma-separated list of URL patterns to block (e.g., “analytics,tracking,advertisement”).
	BlockUrls *string `json:"block_urls" api:"block_urls"`
	// Delay in seconds before capturing the screenshot. Default: 0.
	Delay *float64 `json:"delay" api:"delay"`
	// Define when to capture. Allowed values: domcontentloaded, load, networkidle0, networkidle2.
	WaitUntil []string `json:"wait_until" api:"wait_until"`
	// Wait for a specific element to appear before taking the screenshot.
	WaitForSelector *string `json:"wait_for_selector" api:"wait_for_selector"`
	// Specify a proxy server to route your request through.
	Proxy *string `json:"proxy" api:"proxy"`
	// Cache the response. Default: false.
	Cache *bool `json:"cache" api:"cache"`
	// Cache the response for a custom TTL (in seconds). Default: 2592000.
	CacheTTL *float64 `json:"cache_ttl" api:"cache_ttl"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ContentParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 19)
	if p.URL != "" {
		values = append(values, param{"url", p.URL})
	}
	if p.IncludeHTML != nil {
		values = append(values, param{"include_html", *p.IncludeHTML})
	}
	if p.IncludeMarkdown != nil {
		values = append(values, param{"include_markdown", *p.IncludeMarkdown})
	}
	if p.IncludeSitemap != nil {
		values = append(values, param{"include_sitemap", *p.IncludeSitemap})
	}
	if p.UseDefuddle != nil {
		values = append(values, param{"use_defuddle", *p.UseDefuddle})
	}
	if p.RemoveSelectors != nil {
		values = append(values, param{"remove_selectors", *p.RemoveSelectors})
	}
	if p.RemoveCookieBanners != nil {
		values = append(values, param{"remove_cookie_banners", *p.RemoveCookieBanners})
	}
	if p.RemoveAds != nil {
		values = append(values, param{"remove_ads", *p.RemoveAds})
	}
	if p.BlockResources != nil {
		values = append(values, param{"block_resources", *p.BlockResources})
	}
	if p.BlockUrls != nil {
		values = append(values, param{"block_urls", *p.BlockUrls})
	}
	if p.Delay != nil {
		values = append(values, param{"delay", *p.Delay})
	}
	if p.WaitUntil != nil {
		values = append(values, param{"wait_until", p.WaitUntil})
	}
	if p.WaitForSelector != nil {
		values = append(values, param{"wait_for_selector", *p.WaitForSelector})
	}
	if p.Proxy != nil {
		values = append(values, param{"proxy", *p.Proxy})
	}
	if p.Cache != nil {
		values = append(values, param{"cache", *p.Cache})
	}
	if p.CacheTTL != nil {
		values = append(values, param{"cache_ttl", *p.CacheTTL})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// JobsParams holds the parameters of Jobs (GET /v1/jobs/{id}).
// Required fields are plain values; optional fields are pointers (use capturekit.String, capturekit.Float, capturekit.Bool) or nil slices and maps.
type JobsParams struct {
	// The job_id returned by an endpoint called with is_async=true.
	ID string `json:"id" api:"id,required"`
}

func (p *JobsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 1)
	if p.ID != "" {
		values = append(values, param{"id", p.ID})
	}
	return values
}

// StatsParams holds the parameters of Stats (GET /v1/stats).
// Required fields are plain values; optional fields are pointers (use capturekit.String, capturekit.Float, capturekit.Bool) or nil slices and maps.
type StatsParams struct {
	// API family to check, for example capture. Provide at least engine or path.
	Engine *string `json:"engine" api:"engine"`
	// Specific endpoint to check, for example /v1/capture. If you also pass engine, use an endpoint from that API.
	Path *string `json:"path" api:"path"`
}

func (p *StatsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 2)
	if p.Engine != nil {
		values = append(values, param{"engine", *p.Engine})
	}
	if p.Path != nil {
		values = append(values, param{"path", *p.Path})
	}
	return values
}

// WebsiteCrawlerParams holds the parameters of WebsiteCrawler (GET / POST /v1/website/crawler).
// Required fields are plain values; optional fields are pointers (use capturekit.String, capturekit.Float, capturekit.Bool) or nil slices and maps.
type WebsiteCrawlerParams struct {
	// A website URL with either the http or https protocol.
	Query string `json:"query" api:"query,required"`
	// If set to false, do not follow redirects. Default: true.
	AllowRedirects *bool `json:"allow_redirects" api:"allow_redirects"`
	// If set to true, the response will be the HTML source of the page. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WebsiteCrawlerParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.AllowRedirects != nil {
		values = append(values, param{"allow_redirects", *p.AllowRedirects})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}
