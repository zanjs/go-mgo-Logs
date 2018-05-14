package conf

import (
	"fmt"

	"github.com/jinzhu/configor"
)

// Config is
var Config = struct {
	APP struct {
		Name string `default:"app name"`
		Port string `default:"18000"`
	}

	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}

	Contacts []struct {
		Name  string
		Email string `required:"true"`
	}
}{}

func init() {
	configor.Load(&Config, "config.yml")
	fmt.Printf("config: %#v", Config)
}
