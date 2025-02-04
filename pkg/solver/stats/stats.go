package stats

type StatsOptions struct {
	Enabled bool
	URL     string
}

type Stats interface{}

func NewStats(options StatsOptions) (Stats, error) {
	if !options.Enabled {
		return &NoopStats{}, nil
	}
	return &HTTPStats{URL: options.URL}, nil
}

// Stats API implementation

type HTTPStats struct {
	URL string
}

// Noop implementation

type NoopStats struct{}
