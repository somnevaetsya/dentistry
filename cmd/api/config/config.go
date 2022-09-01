package config

import (
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"os"
)

const (
	host = "PG_HOST"
	user = "PG_USER"
	pass = "PG_PASS"
	name = "PG_NAME"
)

type Config struct {
	ServerPort   string `json:"server_port"`
	PostgresPort string `json:"postgres_port"`
	Urls         struct {
		RootUrl     string `json:"root_url"`
		ProfileUrl  string `json:"profile_url"`
		LoginUrl    string `json:"login_url"`
		RegisterUrl string `json:"register_url"`
		LogoutUrl   string `json:"logout_url"`
		RefactorUrl string `json:"refactor_url"`
		UploadUrl   string `json:"upload_url"`
		VerifyUrl   string `json:"verify_url"`
	} `json:"urls"`
	Db               Postgres
	SessionContainer string `json:"session_container"`
	EmailContainer   string `json:"email_container"`
	LogFile          string `json:"log_file"`
	Expiration       int64  `json:"cookie_expiration"`
}

type Postgres struct {
	Host string
	User string
	Pass string
	Name string
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
	cfg.Db.Host = os.Getenv(host)
	cfg.Db.User = os.Getenv(user)
	cfg.Db.Pass = os.Getenv(pass)
	cfg.Db.Name = os.Getenv(name)

	return cfg, nil
}
