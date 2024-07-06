package options

type AllowlistOptions struct {
	AllowlistCmd bool
}

func NewAllowlistOptions() AllowlistOptions {
	options := AllowlistOptions{
		AllowlistCmd: false,
	}
	return options
}
func CheckAllowlistOptions(options AllowlistOptions) error {
	return nil
}
