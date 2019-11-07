package config

import (
	"os"
	"strconv"
	"testing"
)

func setupTestCase(t *testing.T) func(t *testing.T) {
	os.Clearenv()
	return func(t *testing.T) {
		os.Clearenv()
	}
}
func TestGetEnvIntValueDefault(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	env := "Port"
	defaultPort := 8080
	port := getEnvIntValue(env, defaultPort)
	if port != defaultPort {
		t.Errorf("Get env var was incorrect, got: %d, want: %d.", port, defaultPort)
	}
}

func TestGetEnvIntValueNoneDefault(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	env := "Port"
	specifiedPort := 3000
	defaultPort := 8080
	os.Setenv(env, strconv.Itoa(specifiedPort))
	port := getEnvIntValue(env, defaultPort)
	if port != specifiedPort {
		t.Errorf("Get env var was incorrect, got: %d, want: %d.", port, specifiedPort)
	}
}

func TestGetEnvStrValueDefault(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	env := "RedisHost"
	defaultRedisHost := "127.0.0.1"
	redisHost := getEnvStrValue(env, defaultRedisHost)
	if redisHost != "127.0.0.1" {
		t.Errorf("Get env var was incorrect, got: %s, want: %s.", redisHost, "127.0.0.1")
	}
}

func TestGetEnvStrValueNoneDefault(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	env := "RedisHost"
	defaultRedisHost := "127.0.0.1"
	specifiedRedisHost := "172.16.1.2"
	os.Setenv(env, specifiedRedisHost)
	redisHost := getEnvStrValue(env, defaultRedisHost)
	if redisHost != specifiedRedisHost {
		t.Errorf("Get env var was incorrect, got: %s, want: %s.", redisHost, specifiedRedisHost)
	}
}

func TestLoadConfiguration(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	defaultPort := 8080
	conf := loadConfiguration()
	if conf.Port != defaultPort {
		t.Errorf("Get default settings was incorrect, got: %d, want: %d.", conf.Port, defaultPort)
	}
}
