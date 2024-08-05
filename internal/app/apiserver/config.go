package apiserver

//Config ...
type Config struct {
	BindAddr    string `json:"bind_addr" toml:"bind_addr"`
	LogLevel    string `json:"log_level" toml:"log_level"`
	DatabaseURL string `json:"database_url" toml:"database_url"`
	SessionKey  string `json:"session_key" toml:"session_key"`
}

//NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: "user_name:password@tcp(localhost:3306)/db_name",
		LogLevel: "debug",
	}
}
