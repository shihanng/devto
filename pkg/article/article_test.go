package article

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	images := map[string]string{
		"./image.png": "./a/image.png",
	}

	expected := `---
title: "A title"
published: false
description: "A description"
tags: "tag-one, tag-two"
---
![image](./a/image.png)
`

	actual, err := Read("./testdata/testdata.md", images)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
