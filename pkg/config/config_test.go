package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	if c, err := LoadDefaultConfig(); err != nil {
		t.Errorf("Failed to load default config: %v", err)
	} else {
		t.Logf("Loaded default config: %#v", c)
	}
}
