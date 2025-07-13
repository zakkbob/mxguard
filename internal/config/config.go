package config

type PostgresConfig struct {
	Url      string `yaml:"url"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SSLmode  string `yaml:"sslmode"`
}

type Config struct {
	Verbose  bool `yaml:"verbose"`
	Debug    bool `yaml:"debug"`
	Postgres PostgresConfig
}
