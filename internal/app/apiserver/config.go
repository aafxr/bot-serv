package apiserver

//Config ...
type Config struct {
	BindAddr    string
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:database_url"`
	SessionKey  string `toml:"session_key"`
}

//NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: "user_name:password@tcp(localhost:3306)/db_name",
		LogLevel: "debug",
	}
}
