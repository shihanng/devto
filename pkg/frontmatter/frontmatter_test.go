package frontmatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	type args struct {
		content []byte
	}

	tests := []struct {
		name      string
		args      args
		want      []byte
		want1     []byte
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "no triple-dashed lines",
			args: args{
				content: []byte(`test`),
			},
			want:      nil,
			want1:     []byte(`test`),
			assertion: assert.NoError,
		},
		{
			name: "simple triple-dashed lines",
			args: args{
				content: []byte(`---
frontmatter
---
markdown`),
			},
			want: []byte(`---
frontmatter
---
`),
			want1:     []byte(`markdown`),
			assertion: assert.NoError,
		},
		{
			name: "no second triple-dashed",
			args: args{
				content: []byte(`---
frontmatter
markdown`),
			},
			want: []byte(`---
frontmatter
markdown`),
			want1:     nil,
			assertion: assert.NoError,
		},
		{
			name: "begin with 3 chars",
			args: args{
				content: []byte(`+++
frontmatter
---
markdown`),
			},
			want: nil,
			want1: []byte(`+++
frontmatter
---
markdown`),
			assertion: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Split(tt.args.content)

			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestGetFrontMatter(t *testing.T) {
	type args struct {
		data []byte
	}

	tests := []struct {
		name      string
		args      args
		want      *FrontMatter
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "empty",
			args: args{
				data: nil,
			},
			want:      &FrontMatter{},
			assertion: assert.NoError,
		},
		{
			name: "two triple-dashed lines",
			args: args{
				data: []byte(`---
---`),
			},
			want:      &FrontMatter{},
			assertion: assert.NoError,
		},
		{
			name: "simple",
			args: args{
				data: []byte(`---
title: test title
published: true
tags: go, tutorial
---`),
			},
			want: &FrontMatter{
				Title:     "test title",
				Published: true,
				Tags:      "go, tutorial",
			},
			assertion: assert.NoError,
		},
		{
			name: "ill-formatted",
			args: args{
				data: []byte(`- title: test title`),
			},
			want:      nil,
			assertion: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFrontMatter(tt.args.data)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
