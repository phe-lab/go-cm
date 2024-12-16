package test

import (
	"os"
	"testing"

	"github.com/phe-lab/go-cm/cfg"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig_FromFile(t *testing.T) {
	// Create YAML configuration file
	configFileContent := `
env: "TestApp"
server:
  port: 8080
  host: "127.0.0.1"
database:
  user: "dbuser"
  password: "dbpass"
`
	fileName := "test_config.yaml"
	err := os.WriteFile(fileName, []byte(configFileContent), 0644)
	assert.NoError(t, err)
	defer os.Remove(fileName) // Delete YAML file

	var config Config
	err = cfg.LoadConfig(&config, cfg.LoadOptions{
		Filename: fileName,
	})
	assert.NoError(t, err)
	assert.Equal(t, "TestApp", config.Env)
	assert.Equal(t, 8080, config.Server.Port)
	assert.Equal(t, "127.0.0.1", config.Server.Host)
	assert.Equal(t, "dbuser", config.Database.User)
	assert.Equal(t, "dbpass", config.Database.Password)
}
