package config

import (
	env "github.com/caarlos0/env/v6"
)

type environment struct {
	AppName    string `env:"APP_NAME,required"`
	AppVersion string `env:"APP_VERSION,required"`
	AppAuthor  string `env:"APP_AUTHOR,required"`
}

func New() (*Wrapper, error) {
	environment := environment{}
	if err := env.Parse(&environment); err != nil {
		return &Wrapper{}, err
	}

	return &Wrapper{
		App: app{
			Name:    environment.AppName,
			Version: environment.AppVersion,
			Author:  environment.AppAuthor,
		},
	}, nil
}

type Wrapper struct {
	App app
}

type app struct {
	Name    string
	Version string
	Author  string
}
