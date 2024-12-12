package cfg

import "github.com/spf13/pflag"

// LoadOptions defines the options used to load the configuration.
// - Filename: The path to the YAML configuration file.
// - Flags: Command-line flags provided by the application.
// - EnvPrefix: The prefix for environment variables to load into the configuration.
type LoadOptions struct {
	Filename      string
	Flags         *pflag.FlagSet
	EnvPrefix     string
	DefaultValues *map[string]interface{}
}
