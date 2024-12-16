package test

import (
	"os"
	"testing"

	"github.com/phe-lab/go-cm/cfg"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig_PriorityOrder(t *testing.T) {
	// Setup default values
	defaultValues := map[string]interface{}{
		"env":         "DefaultApp",
		"server.port": 3000,
		"server.host": "localhost",
	}

	// Setup YAML file
	configFileContent := `
env: "FileApp"
server:
  port: 8080
  host: "127.0.0.1"
`
	fileName := "test_priority_config.yaml"
	err := os.WriteFile(fileName, []byte(configFileContent), 0644)
	assert.NoError(t, err)
	defer os.Remove(fileName) // Xóa file sau khi test xong

	// Setup environment variables
	os.Setenv("APP_SERVER_PORT", "9090")
	os.Setenv("APP_SERVER_HOST", "0.0.0.0")
	defer os.Unsetenv("APP_SERVER_PORT")
	defer os.Unsetenv("APP_SERVER_HOST")

	// Setup command line flags
	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
	flags.String("env", "DefaultApp", "application name")
	flags.Int("server.port", 7000, "server port")
	flags.String("server.host", "192.168.1.100", "server host")
	flags.Parse([]string{
		"--env=CLIApp",
	})

	var config Config
	err = cfg.LoadConfig(&config, cfg.LoadOptions{
		Filename:      fileName,
		EnvPrefix:     "APP_",
		Flags:         flags,
		DefaultValues: &defaultValues,
	})
	assert.NoError(t, err)

	// Check priority order
	assert.Equal(t, "CLIApp", config.Env)          // Lệnh CLI có ưu tiên cao nhất
	assert.Equal(t, 9090, config.Server.Port)      // Biến môi trường có ưu tiên cao hơn file
	assert.Equal(t, "0.0.0.0", config.Server.Host) // Biến môi trường có ưu tiên cao hơn file
}
