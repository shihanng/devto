package article

import (
	"testing"

	"github.com/gohugoio/hugo/parser/metadecoders"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	nope := false
	actual, err := Parse("./testdata/testdata.md")
	expected := &Parsed{
		frontMatterFormat: metadecoders.YAML,
		content: []byte(`
![image](./image.png)
[Google](www.google.com)
![image](./image-2.png)
`),
		frontMatter: FrontMatter{
			Title:       "A title",
			Published:   &nope,
			Description: "A description",
			Tags:        "tag-one, tag-two",
			CoverImage:  "./cv.jpg",
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)

	expectedContent := `---
title: A title
published: false
description: A description
tags: tag-one, tag-two
cover_image: ./cv.jpg
---

![image](./image.png)
[Google](www.google.com)
![image](./image-2.png)
`
	actualContent, err := actual.Content()
	assert.NoError(t, err)
	assert.Equal(t, expectedContent, actualContent)
}

func TestParse_NotFound(t *testing.T) {
	actual, err := Parse("./testdata/unknown.md")
	assert.Error(t, err)
	assert.Nil(t, actual)
}
