package util

import "testing"

func TestLoadConfig (t *testing.T) { 
	config, err := LoadConfig("../")
	if err != nil {
		t.Errorf("cannot load config: %v", err)
	}
	t.Logf("config: %v", config)
}