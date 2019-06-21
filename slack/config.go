package slack

import "github.com/nlopes/slack"

type Config struct {
	APIToken string
}

func (c *Config) Client() (interface{}, error) {
	return slack.New(c.APIToken), nil
}
