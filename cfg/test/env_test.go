package test

import (
	"os"
	"testing"

	"github.com/phe-lab/go-cm/cfg"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig_FromEnvVariables(t *testing.T) {
	// Create OS Environment Variables
	os.Setenv("APP_ENV", "EnvApp")
	os.Setenv("APP_SERVER_PORT", "9090")
	os.Setenv("APP_SERVER_HOST", "0.0.0.0")
	defer os.Unsetenv("APP_ENV")
	defer os.Unsetenv("APP_SERVER_PORT")
	defer os.Unsetenv("APP_SERVER_HOST")

	var config Config
	err := cfg.LoadConfig(&config, cfg.LoadOptions{
		EnvPrefix: "APP_",
	})
	assert.NoError(t, err)
	assert.Equal(t, "EnvApp", config.Env)
	assert.Equal(t, 9090, config.Server.Port)
	assert.Equal(t, "0.0.0.0", config.Server.Host)
}
