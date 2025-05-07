package http

type ServerOptions struct {
	URL           string
	Host          string
	Port          int
	AccessControl AccessControlOptions
	RateLimiter   RateLimiterOptions
	Storage       StorageOptions
}

type AccessControlOptions struct {
	EnableResourceProviderAllowlist bool
	ValidationTokenSecret           string
	ValidationTokenExpiration       int
	ValidationTokenKid              string
	AnuraAddresses                  []string
	OfferTimestampDiffSeconds       int
	EnableVersionCheck              bool
	MinimumVersion                  string
	MaximumJobOfferCapacity         int
}

type ValidationToken struct {
	JWT string
}

type RateLimiterOptions struct {
	RequestLimit int
	WindowLength int
}

type StorageOptions struct {
	MaximumFileInputsMemoryMB int
	MaximumFileInputsSizeMB   int
}

type ClientOptions struct {
	URL           string
	PrivateKey    string
	PublicAddress string
	Type          string
}
