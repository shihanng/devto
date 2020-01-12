package article

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/antihax/optional"
	"github.com/golang/mock/gomock"
	mock_article "github.com/shihanng/devto/pkg/article/mock"
	"github.com/shihanng/devto/pkg/devto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubmitArticle(t *testing.T) {
	apiKey := "abc1234"

	dir, err := ioutil.TempDir("", "devto_test")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	filename := filepath.Join(dir, "test.md")
	require.NoError(t, ioutil.WriteFile(filename, []byte("---\n---\ntest"), 0644))

	c, err := NewClient(apiKey, filename)
	assert.NoError(t, err)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAPIClient := mock_article.NewMockapiClient(ctrl)
	mockAPIClient.EXPECT().CreateArticle(c.contextWithAPIKey(), &devto.ArticlesApiCreateArticleOpts{
		ArticleCreate: optional.NewInterface(devto.ArticleCreate{
			Article: devto.ArticleCreateArticle{
				BodyMarkdown: "---\n---\ntest",
			},
		},
		),
	}).Return(devto.ArticleShow{Id: 1}, nil, nil)
	mockAPIClient.EXPECT().UpdateArticle(c.contextWithAPIKey(), int32(1), &devto.ArticlesApiUpdateArticleOpts{
		ArticleUpdate: optional.NewInterface(devto.ArticleUpdate{
			Article: devto.ArticleUpdateArticle{
				BodyMarkdown: "---\n---\ntest",
			},
		},
		),
	}).Return(devto.ArticleShow{Id: 1}, nil, nil)
	c.api = mockAPIClient

	assert.NoError(t, c.SubmitArticle())
	assert.NoError(t, c.SubmitArticle())
}
