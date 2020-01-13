package article

import (
	"bytes"
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
	const (
		apiKey    string = "abc1234"
		articleID int32  = 123
	)

	dir, err := ioutil.TempDir("", "devto_test")
	require.NoError(t, err)

	defer os.RemoveAll(dir)

	filename := filepath.Join(dir, "test.md")
	require.NoError(t, ioutil.WriteFile(filename, []byte("---\n---\ntest"), 0644))

	c, err := NewClient(apiKey, SetConfig(filename))
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
	}).Return(devto.ArticleShow{Id: articleID}, nil, nil)
	mockAPIClient.EXPECT().UpdateArticle(c.contextWithAPIKey(), articleID, &devto.ArticlesApiUpdateArticleOpts{
		ArticleUpdate: optional.NewInterface(devto.ArticleUpdate{
			Article: devto.ArticleUpdateArticle{
				BodyMarkdown: "---\n---\ntest",
			},
		},
		),
	}).Return(devto.ArticleShow{Id: articleID}, nil, nil)

	c.api = mockAPIClient

	assert.NoError(t, c.SubmitArticle(filename))
	assert.NoError(t, c.SubmitArticle(filename))
}

func TestListArticle(t *testing.T) {
	const apiKey = "abc1234"

	c, err := NewClient(apiKey)
	assert.NoError(t, err)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAPIClient := mock_article.NewMockapiClient(ctrl)
	mockAPIClient.EXPECT().GetUserAllArticles(c.contextWithAPIKey(), nil).
		Return([]devto.ArticleMe{{Title: "A title", Id: 1}}, nil, nil)

	c.api = mockAPIClient

	actual := bytes.Buffer{}
	expected := "[1] A title\n"

	assert.NoError(t, c.ListArticle(&actual))
	assert.Equal(t, expected, actual.String())
}
