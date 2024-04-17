package config

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	LogLevel    string        `envconfig:"LOG_LEVEL" default:"info"`
	LogJson     bool          `envconfig:"LOG_JSON" default:"false"`
	AccessToken string        `envconfig:"ACCESS_TOKEN" required:"true"`
	ApiTimeout  time.Duration `envconfig:"API_TIMEOUT" default:"10s"`
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
