package article

import (
	"testing"

	"github.com/gohugoio/hugo/parser/metadecoders"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	actual, err := Parse("./testdata/testdata.md")
	expected := &Parsed{
		frontMatterFormat: metadecoders.YAML,
		frontMatterSource: []byte(`title: "A title"
published: false
description: "A description"
tags: "tag-one, tag-two"
`),
		frontMatter: FrontMatter{
			Title:       "A title",
			Published:   false,
			Description: "A description",
			Tags:        "tag-one, tag-two",
		},
		markdownSource: []byte(`
![image](./image.png)
[Google](www.google.com)
![image](./image-2.png)
`),
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)

	expectedContent := `---
title: "A title"
published: false
description: "A description"
tags: "tag-one, tag-two"
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
