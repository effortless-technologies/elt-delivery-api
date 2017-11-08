// +build prod

package config

import "os"

func MakeConfig() *Config {

	c := new(Config)
	c.PostmatesAPIKey = os.Getenv("POSTMATES_PROD_KEY")

	return c
}
