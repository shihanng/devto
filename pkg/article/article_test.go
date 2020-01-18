package article

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetImageLinks(t *testing.T) {
	type args struct {
		filename string
		images   map[string]string
	}

	tests := []struct {
		name      string
		args      args
		want      string
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "normal",
			args: args{
				filename: "./testdata/testdata.md",
				images: map[string]string{
					"./image.png":   "./a/image.png",
					"./image-2.png": "",
				},
			},
			want: `---
title: "A title"
published: false
description: "A description"
tags: "tag-one, tag-two"
---
![image](./a/image.png)
[Google](www.google.com)
![image](./image-2.png)
`,
			assertion: assert.NoError,
		},
		{
			name: "not found",
			args: args{
				filename: "./testdata/unknown.md",
				images:   map[string]string{"./image.png": "./a/image.png"},
			},
			want:      "",
			assertion: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SetImageLinks(tt.args.filename, tt.args.images)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetImageLinks(t *testing.T) {
	type args struct {
		filename string
	}

	tests := []struct {
		name      string
		args      args
		want      map[string]string
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "normal",
			args: args{filename: "./testdata/testdata.md"},
			want: map[string]string{
				"./image.png":   "",
				"./image-2.png": "",
			},
			assertion: assert.NoError,
		},
		{
			name:      "not found",
			args:      args{filename: "./testdata/unknown.md"},
			want:      nil,
			assertion: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetImageLinks(tt.args.filename)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
