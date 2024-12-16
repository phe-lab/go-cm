package test

type Config struct {
	Env    string `koanf:"env"`
	Server struct {
		Port int    `koanf:"port"`
		Host string `koanf:"host"`
	} `koanf:"server"`
	Database struct {
		User     string `koanf:"user"`
		Password string `koanf:"password"`
	} `koanf:"database"`
}
