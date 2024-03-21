package config

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Port        int    `envconfig:"PORT" default:"8080"`
	LogLevel    string `envconfig:"LOG_LEVEL" default:"info"`
	LogJson     bool   `envconfig:"LOG_JSON" default:"false"`
	AccessToken string `envconfig:"ACCESS_TOKEN" required:"true"`
}

var (
	cfg  config
	once sync.Once
)

func GetDefault() *config {
	once.Do(func() {
		if err := envconfig.Process("", &cfg); err != nil {
			log.Fatal(err)
		}
	})
	return &cfg
}

func (c config) String() string {
	str, err := json.Marshal(c)
	if err != nil {
		log.Fatalf("failed to parse config to string: %s\n", err.Error())
		return ""
	}
	return string(str)
}
