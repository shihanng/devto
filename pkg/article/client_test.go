package article

import (
	"bytes"
	"testing"

	"github.com/antihax/optional"
	"github.com/golang/mock/gomock"
	mock_article "github.com/shihanng/devto/pkg/article/mock"
	"github.com/shihanng/devto/pkg/devto"
	"github.com/stretchr/testify/assert"
)

const (
	apiKey             = "abc1234"
	articleID    int32 = 123
	emptyArticle       = "---\n---\n"
)

func TestSubmitArticle_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAPIClient := mock_article.NewMockapiClient(ctrl)
	mockConfig := mock_article.NewMockconfiger(ctrl)

	c, err := NewClient(apiKey, SetConfig(mockConfig))
	assert.NoError(t, err)

	c.api = mockAPIClient

	{
		mockConfig.EXPECT().ImageLinks().Return(nil)
		mockConfig.EXPECT().ArticleID().Return(int32(0))
		mockAPIClient.EXPECT().CreateArticle(c.contextWithAPIKey(), &devto.ArticlesApiCreateArticleOpts{
			ArticleCreate: optional.NewInterface(devto.ArticleCreate{
				Article: devto.ArticleCreateArticle{
					BodyMarkdown: emptyArticle,
				},
			},
			),
		}).Return(devto.ArticleShow{Id: articleID}, nil, nil)
		mockConfig.EXPECT().SetArticleID(articleID)
		mockConfig.EXPECT().Save().Return(nil)

		assert.NoError(t, c.SubmitArticle("./testdata/empty.md"))
	}
}

func TestSubmitArticle_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAPIClient := mock_article.NewMockapiClient(ctrl)
	mockConfig := mock_article.NewMockconfiger(ctrl)

	c, err := NewClient(apiKey, SetConfig(mockConfig))
	assert.NoError(t, err)

	c.api = mockAPIClient

	mockConfig.EXPECT().ImageLinks().Return(nil)
	mockConfig.EXPECT().ArticleID().Return(articleID)
	mockConfig.EXPECT().ArticleID().Return(articleID)
	mockAPIClient.EXPECT().UpdateArticle(c.contextWithAPIKey(), articleID, &devto.ArticlesApiUpdateArticleOpts{
		ArticleUpdate: optional.NewInterface(devto.ArticleUpdate{
			Article: devto.ArticleUpdateArticle{
				BodyMarkdown: emptyArticle,
			},
		},
		),
	}).Return(devto.ArticleShow{Id: articleID}, nil, nil)

	assert.NoError(t, c.SubmitArticle("./testdata/empty.md"))
}

func TestListArticle(t *testing.T) {
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

func TestGenerateImageLinks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfig := mock_article.NewMockconfiger(ctrl)

	c, err := NewClient(apiKey, SetConfig(mockConfig))
	assert.NoError(t, err)

	mockConfig.EXPECT().ImageLinks().Return(map[string]string{
		"./image.png": "image-1.png",
	})
	mockConfig.EXPECT().SetImageLinks(map[string]string{
		"./image.png":   "image-1.png",
		"./image-2.png": "",
	})
	mockConfig.EXPECT().Save().Return(nil)

	assert.NoError(t, c.GenerateImageLinks("./testdata/testdata.md"))
}
