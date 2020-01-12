//go:generate mockgen -source=client.go -destination=mock/mock_client.go
package article

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	"github.com/antihax/optional"
	"github.com/cockroachdb/errors"
	"github.com/shihanng/devto/pkg/devto"
	"github.com/spf13/viper"
)

type apiClient interface {
	CreateArticle(context.Context, *devto.ArticlesApiCreateArticleOpts) (devto.ArticleShow, *http.Response, error)
	UpdateArticle(context.Context, int32, *devto.ArticlesApiUpdateArticleOpts) (devto.ArticleShow, *http.Response, error)
}

type Client struct {
	api      apiClient
	apiKey   string
	viper    *viper.Viper
	filename string
}

func NewClient(apiKey, filename string) (*Client, error) {
	c := Client{
		api:      devto.NewAPIClient(devto.NewConfiguration()).ArticlesApi,
		apiKey:   apiKey,
		viper:    viper.New(),
		filename: filename,
	}

	c.viper = viper.New()
	c.viper.SetConfigFile(configFrom(filename))

	if err := c.viper.ReadInConfig(); err != nil {
		if !os.IsNotExist(err) {
			return nil, errors.Wrap(err, "article: read config")
		}
	}

	return &c, nil
}

func (c *Client) SubmitArticle() error {
	body, err := Read(c.filename, c.configImageLinks())
	if err != nil {
		return err
	}

	switch c.configArticleID() {
	case 0:
		article := &devto.ArticlesApiCreateArticleOpts{
			ArticleCreate: optional.NewInterface(devto.ArticleCreate{
				Article: devto.ArticleCreateArticle{
					BodyMarkdown: body,
				},
			},
			),
		}

		submitted, _, err := c.api.CreateArticle(c.contextWithAPIKey(), article)
		if err != nil {
			return errors.Wrap(err, "article: create article in dev.to")
		}

		c.setConfigArticleID(submitted.Id)
		return c.updateConfig()
	default:
		articleID := c.configArticleID()

		article := &devto.ArticlesApiUpdateArticleOpts{
			ArticleUpdate: optional.NewInterface(devto.ArticleUpdate{
				Article: devto.ArticleUpdateArticle{
					BodyMarkdown: body,
				},
			},
			),
		}

		_, _, err := c.api.UpdateArticle(c.contextWithAPIKey(), articleID, article)
		return errors.Wrapf(err, "article: update article %d in dev.to", articleID)
	}
}

func (c *Client) contextWithAPIKey() context.Context {
	return context.WithValue(context.Background(), devto.ContextAPIKey, devto.APIKey{
		Key: c.apiKey,
	})
}

func (c *Client) configImageLinks() map[string]string {
	return c.viper.GetStringMapString("images")
}

func (c *Client) configArticleID() int32 {
	return c.viper.GetInt32("article_id")
}

func (c *Client) setConfigArticleID(id int32) {
	c.viper.Set("article_id", id)
}

func (c *Client) updateConfig() error {
	return errors.Wrap(c.viper.WriteConfigAs(configFrom(c.filename)), "article: update config")
}

func configFrom(filename string) string {
	return filepath.Join(filepath.Dir(filename), "devto.yml")
}
