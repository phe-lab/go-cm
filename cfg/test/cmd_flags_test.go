package test

import (
	"testing"

	"github.com/phe-lab/go-cm/cfg"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig_FromCommandLineFlags(t *testing.T) {
	// Setup command-line flags
	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
	flags.String("env", "DefaultApp", "application name")
	flags.Int("server.port", 8081, "server port")
	flags.String("server.host", "localhost", "server host")
	flags.Parse([]string{
		"--env=CLIApp",
		"--server.port=9091",
		"--server.host=192.168.1.100",
	})

	var config Config
	err := cfg.LoadConfig(&config, cfg.LoadOptions{
		Flags: flags,
	})
	assert.NoError(t, err)
	assert.Equal(t, "CLIApp", config.Env)
	assert.Equal(t, 9091, config.Server.Port)
	assert.Equal(t, "192.168.1.100", config.Server.Host)
}
