package config

import (
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"os"
)

const (
	mailPass = "MAIL_PASS"
)

type Config struct {
	MailHost       string `json:"mail_host"`
	MailPort       int    `json:"mail_port"`
	MailUsername   string `json:"mail_username"`
	MailPassword   string
	EmailContainer string `json:"email_container"`
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
	cfg.MailPassword = os.Getenv(mailPass)
	return cfg, nil
}
