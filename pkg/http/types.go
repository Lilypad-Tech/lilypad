package http

type ServerOptions struct {
	URL           string
	Host          string
	Port          int
	AccessControl AccessControlOptions
	RateLimiter   RateLimiterOptions
}

type AccessControlOptions struct {
	ValidationTokenSecret     string
	ValidationTokenExpiration int
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
