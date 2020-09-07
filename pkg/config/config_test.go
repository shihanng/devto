package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	type args struct {
		filename string
	}

	tests := []struct {
		name           string
		args           args
		valueAssertion assert.ValueAssertionFunc
		errAssertion   assert.ErrorAssertionFunc
	}{
		{
			name: "missing config file",
			args: args{
				filename: "./testdata/unknown.yml",
			},
			valueAssertion: assert.NotNil,
			errAssertion:   assert.NoError,
		},
		{
			name: "empty config file",
			args: args{
				filename: "./testdata/empty.yml",
			},
			valueAssertion: assert.NotNil,
			errAssertion:   assert.NoError,
		},
		{
			name: "bad config file",
			args: args{
				filename: "./testdata/bad.yml",
			},
			valueAssertion: assert.Nil,
			errAssertion:   assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.filename)
			tt.valueAssertion(t, got)
			tt.errAssertion(t, err)
		})
	}
}

func TestConfig_Save(t *testing.T) {
	type fields struct {
		filename string
	}

	tests := []struct {
		name      string
		fields    fields
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "no config file",
			fields: fields{
				filename: "",
			},
			assertion: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := New(tt.fields.filename)
			require.NoError(t, err)
			tt.assertion(t, c.Save())
		})
	}
}

func TestGetters(t *testing.T) {
	c, err := New("./testdata/devto.yml")
	require.NoError(t, err)

	{
		var expected int32 = 1985
		assert.Equal(t, expected, c.ArticleID())
	}

	{
		expected := map[string]string{
			"key-1": "value-a",
			"key-2": "value-b",
		}
		assert.Equal(t, expected, c.ImageLinks())
	}
}

func TestSetters(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "temp.*.yml")
	require.NoError(t, err)

	defer os.Remove(tmpfile.Name()) // clean up

	c, err := New(tmpfile.Name())
	require.NoError(t, err)

	c.SetArticleID(1985)
	c.SetImageLinks(map[string]string{
		"key-1": "value-a",
		"key-2": "value-b",
	})

	require.NoError(t, c.Save())

	expected := []byte(`article_id: 1985
images:
  key-1: value-a
  key-2: value-b
`)

	actual, err := ioutil.ReadFile(tmpfile.Name())
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}
