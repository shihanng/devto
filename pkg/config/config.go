package config

import (
	"os"

	"github.com/cockroachdb/errors"
	"github.com/spf13/viper"
)

type Config struct {
	filename string
	viper    *viper.Viper
}

func New(filename string) (*Config, error) {
	c := Config{
		filename: filename,
		viper:    viper.New(),
	}

	c.viper.SetConfigFile(filename)

	if err := c.viper.ReadInConfig(); err != nil {
		if !os.IsNotExist(err) && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return nil, errors.Wrap(err, "config: read")
		}
	}

	return &c, nil
}

func (c *Config) Save() error {
	return errors.Wrap(c.viper.WriteConfigAs(c.filename), "config: update")
}

func (c *Config) ImageLinks() map[string]string {
	return c.viper.GetStringMapString("images")
}

func (c *Config) SetImageLinks(links map[string]string) {
	c.viper.Set("images", links)
}

func (c *Config) ArticleID() int32 {
	return c.viper.GetInt32("article_id")
}

func (c *Config) SetArticleID(id int32) {
	c.viper.Set("article_id", id)
}
