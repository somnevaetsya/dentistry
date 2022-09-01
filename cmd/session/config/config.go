package config

import (
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"os"
)

const (
	redisPass = "REDIS_PASS"
)

type Config struct {
	SessionContainer string `json:"session_container"`
	RedisContainer   string `json:"redis_container"`
	RedisPass        string
	CookieExpiration int64 `json:"cookie_expiration"`
}

func ParseConfig(path string) (Config, error) {
	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	var cfg Config
	err = jsoniter.Unmarshal(jsonFile, &cfg)
	if err != nil {
		return Config{}, err
	}
	cfg.RedisPass = os.Getenv(redisPass)
	return cfg, nil
}
