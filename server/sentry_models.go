package main

// SentryWebhook contains information about the resource and will differ in content depending on the type of webhook
type SentryWebhook struct {
	Project string      `json:"project"`
	URL     string      `json:"url"`
	Event   SentryEvent `json:"event"`
}

// SentryEvent is generally an error, but it's possible a non-error related data as well
type SentryEvent struct {
	Title    string `json:"title"`
	Location string `json:"location"`
}
