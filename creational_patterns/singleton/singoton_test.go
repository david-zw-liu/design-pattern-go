package singleton

import "testing"

func TestGetConfig(t *testing.T) {
	cc := make(chan *Config)

	go func(cc chan *Config) {
		cc <- GetConfig()
	}(cc)

	go func(cc chan *Config) {
		cc <- GetConfig()
	}(cc)

	config1, config2 := <-cc, <-cc

	if config1 != config2 {
		t.Errorf("Expect that two config pointer are same, but not the same")
	}

	if config1.APIKey != "Testkey" {
		t.Errorf("Expect that APIKey is \"testkey\", but got %v", config1.APIKey)
	}
}
