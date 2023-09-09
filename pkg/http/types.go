package http

type ServerOptions struct {
	URL  string
	Host string
	Port int
}

type ClientOptions struct {
	URL        string
	PrivateKey string
}
