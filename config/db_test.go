package config_test

import (
	"programbuku-v3/config"
	"testing"
)

func TestConnection(t *testing.T) {
	testing.Init()
	config.OpenDB()
}
