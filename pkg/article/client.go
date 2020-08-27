//go:generate mockgen -source=client.go -destination=mock/mock_client.go
package article

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/antihax/optional"
	"github.com/cockroachdb/errors"
	"github.com/shihanng/devto/pkg/devto"
)

type apiClient interface {
	CreateArticle(context.Context, *devto.ArticlesApiCreateArticleOpts) (devto.ArticleShow, *http.Response, error)
	UpdateArticle(context.Context, int32, *devto.ArticlesApiUpdateArticleOpts) (devto.ArticleShow, *http.Response, error)
	GetUserAllArticles(context.Context, *devto.ArticlesApiGetUserAllArticlesOpts) ([]devto.ArticleMe, *http.Response, error)
}

type configer interface {
	Save() error
	ImageLinks() map[string]string
	SetImageLinks(map[string]string)
	ArticleID() int32
	SetArticleID(int32)
}

type Client struct {
	api    apiClient
	apiKey string
	config configer
}

func NewClient(apiKey string, opts ...Option) (*Client, error) {
	c := Client{
		api:    devto.NewAPIClient(devto.NewConfiguration()).ArticlesApi,
		apiKey: apiKey,
	}

	for _, opt := range opts {
		opt(&c)
	}

	return &c, nil
}

type Option func(*Client)

func SetConfig(cfg configer) Option {
	return func(c *Client) {
		c.config = cfg
	}
}

func (c *Client) SubmitArticle(filename string) error {
	body, err := SetImageLinks(filename, c.config.ImageLinks(), CoverImageUntouch)
	if err != nil {
		return err
	}

	switch c.config.ArticleID() {
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

		c.config.SetArticleID(submitted.Id)

		return c.config.Save()
	default:
		articleID := c.config.ArticleID()

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

func (c *Client) ListArticle(w io.Writer) error {
	articles, _, err := c.api.GetUserAllArticles(c.contextWithAPIKey(), nil)
	if err != nil {
		return errors.Wrap(err, "article: list articles in dev.to")
	}

	for _, a := range articles {
		fmt.Fprintf(w, "[%d] %s\n", a.Id, a.Title)
	}

	return nil
}

func (c *Client) GenerateImageLinks(filename string) error {
	links, err := GetImageLinks(filename)
	if err != nil {
		return err
	}

	links = mergeLinks(c.config.ImageLinks(), links)
	c.config.SetImageLinks(links)

	return c.config.Save()
}

func (c *Client) contextWithAPIKey() context.Context {
	return context.WithValue(context.Background(), devto.ContextAPIKey, devto.APIKey{
		Key: c.apiKey,
	})
}

func mergeLinks(old, latest map[string]string) map[string]string {
	for k := range latest {
		if v, ok := old[k]; ok {
			latest[k] = v
		}
	}

	return latest
}
