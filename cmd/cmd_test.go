package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testContent = `---
title: HELLO!
description: "hallo"
cover_image: "./cv.jpg"
---

![lili](./image.png)
![lili](./image.png)
`

func TestGenerate(t *testing.T) {
	dir, err := ioutil.TempDir("", "devto-cmd-test-")
	require.NoError(t, err)

	defer os.RemoveAll(dir)

	tmpfn := filepath.Join(dir, "test_article.md")
	require.NoError(t, ioutil.WriteFile(tmpfn, []byte(testContent), 0666))

	os.Args = []string{"devto", "generate", tmpfn}

	cmd, sync := New()
	defer sync()

	require.NoError(t, cmd.Execute())

	actual, err := ioutil.ReadFile(filepath.Join(dir, "devto.yml"))
	require.NoError(t, err)

	expected := []byte(`cover_image: ./cv.jpg
images:
  ./image.png: ""
`)
	assert.Equal(t, expected, actual)
}

func TestGenerate_Prefix(t *testing.T) {
	dir, err := ioutil.TempDir("", "devto-cmd-test-")
	require.NoError(t, err)

	defer os.RemoveAll(dir)

	tmpfn := filepath.Join(dir, "test_article.md")
	require.NoError(t, ioutil.WriteFile(tmpfn, []byte(testContent), 0666))

	os.Args = []string{"devto", "generate", "-p", "test/", tmpfn}

	cmd, sync := New()
	defer sync()

	require.NoError(t, cmd.Execute())

	actual, err := ioutil.ReadFile(filepath.Join(dir, "devto.yml"))
	require.NoError(t, err)

	expected := []byte(`cover_image: test/./cv.jpg
images:
  ./image.png: test/./image.png
`)
	assert.Equal(t, expected, actual)
}
