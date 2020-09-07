package cmd

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/sebdah/goldie/v2"
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

func TestSubmit_DryRun(t *testing.T) {
	dir, err := ioutil.TempDir("", "devto-cmd-test-")
	require.NoError(t, err)

	defer os.RemoveAll(dir)

	tmpfn := filepath.Join(dir, "test_article.md")
	require.NoError(t, ioutil.WriteFile(tmpfn, []byte(testContent), 0666))

	os.Args = []string{"devto", "submit", "--dry-run", "-p", "test/", tmpfn}

	var out bytes.Buffer

	cmd, sync := New(&out)
	defer sync()

	require.NoError(t, cmd.Execute())

	data := struct {
		Dir string
	}{
		Dir: dir,
	}

	g := goldie.New(t)
	g.AssertWithTemplate(t, "submit_dryrun_out", data, out.Bytes())
}

func TestGenerate(t *testing.T) {
	dir, err := ioutil.TempDir("", "devto-cmd-test-")
	require.NoError(t, err)

	defer os.RemoveAll(dir)

	tmpfn := filepath.Join(dir, "test_article.md")
	require.NoError(t, ioutil.WriteFile(tmpfn, []byte(testContent), 0666))

	os.Args = []string{"devto", "generate", tmpfn}

	cmd, sync := New(os.Stdout)
	defer sync()

	require.NoError(t, cmd.Execute())

	actual, err := ioutil.ReadFile(filepath.Join(dir, "devto.yml"))
	require.NoError(t, err)

	expected := []byte(`images:
  ./cv.jpg: ""
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

	cmd, sync := New(os.Stdout)
	defer sync()

	require.NoError(t, cmd.Execute())

	actual, err := ioutil.ReadFile(filepath.Join(dir, "devto.yml"))
	require.NoError(t, err)

	expected := []byte(`images:
  ./cv.jpg: test/./cv.jpg
  ./image.png: test/./image.png
`)
	assert.Equal(t, expected, actual)
}
