package cfg

import (
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/v2"
)

// k is the global instance of the Koanf library used for managing configuration.
var k = koanf.New(".")

// LoadConfig loads configuration from multiple sources (file, environment variables, and flags)
// and unmarshals it into the given interface.
//
// The configuration is loaded in the following order:
// 1. YAML file: Reads configuration from the file specified in LoadOptions.Filename.
// 2. Environment variables: Reads environment variables with the specified prefix (LoadOptions.EnvPrefix).
// 3. Command-line flags: Reads configuration from the provided pflag.FlagSet.
//
// The final configuration is unmarshaled into the provided object `o`.
//
// Parameters:
// - o: A pointer to the structure where the configuration will be unmarshaled.
// - opt: Options for loading configuration, such as file path, flag set, and environment variable prefix.
//
// Returns:
// An error if any of the sources fail to load or if the unmarshaling process fails.
func LoadConfig(o interface{}, opt LoadOptions) error {
	var finalErr error

	// Read from YAML file
	if err := k.Load(file.Provider(opt.Filename), yaml.Parser()); err != nil {
		finalErr = err
	}

	// Read from environment variables
	if err := k.Load(env.Provider(opt.EnvPrefix, ".", envParser(opt.EnvPrefix)), nil); err != nil {
		finalErr = err
	}

	// Read from command-line flags
	if err := k.Load(posflag.Provider(opt.Flags, ".", k), nil); err != nil {
		finalErr = err
	}

	// Unmarshal config
	if err := k.UnmarshalWithConf("", o, koanf.UnmarshalConf{FlatPaths: false}); err != nil {
		finalErr = err
	}

	return finalErr
}

// envParser returns a function to parse environment variable keys into a dot-separated format.
//
// The parser removes the provided prefix and replaces underscores with dots.
// This is used to convert environment variable names into nested configuration keys.
//
// Example:
// Prefix: "APP_"
// Input: "APP_DATABASE_HOST"
// Output: "database.host"
//
// Parameters:
// - prefix: The prefix to strip from environment variable names.
//
// Returns:
// A function that transforms an environment variable name into a configuration key.
func envParser(prefix string) func(s string) string {
	return func(s string) string {
		return strings.Replace(
			strings.ToLower(strings.TrimPrefix(s, prefix)),
			"_",
			".",
			-1,
		)
	}
}
