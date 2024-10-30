package http

type ServerOptions struct {
	URL         string
	Host        string
	Port        int
	RateLimiter RateLimiterOptions
}

type RateLimiterOptions struct {
	RequestLimit int
	WindowLength int
}

type ClientOptions struct {
	URL           string
	PrivateKey    string
	PublicAddress string
	Type          string
}
