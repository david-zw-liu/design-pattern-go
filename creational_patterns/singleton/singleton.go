package singleton

import "sync"

var config *Config
var once sync.Once // async lock

type Config struct {
	APIKey string
}

func GetConfig() *Config {
	once.Do(func() {
		// load config from anywhere
		config = &Config{APIKey: "Testkey"}
	})

	return config
}
