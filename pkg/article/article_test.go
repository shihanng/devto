package article

import (
	"testing"

	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
)

func TestSetImageLinks(t *testing.T) {
	type args struct {
		filename string
		images   map[string]string
		prefix   string
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
				prefix: "www.example.com/",
			},
			want: `---
title: A title
published: false
description: A description
tags: tag-one, tag-two
cover_image: www.example.com/./cv.jpg
---
![image](./a/image.png)
[Google](www.google.com)
![image](www.example.com/./image-2.png)
`,
			assertion: assert.NoError,
		},
		{
			name: "prefix cover_image",
			args: args{
				filename: "./testdata/testdata.md",
				images: map[string]string{
					"./cv.jpg": "test/./cv.jpg",
				},
			},
			want: `---
title: A title
published: false
description: A description
tags: tag-one, tag-two
cover_image: test/./cv.jpg
---
![image](./image.png)
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
			got, err := SetImageLinks(tt.args.filename, tt.args.images, tt.args.prefix)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSetImageLinks_golden(t *testing.T) {
	content, err := SetImageLinks("./testdata/real_article.md", map[string]string{}, "example.com/")
	assert.NoError(t, err)

	g := goldie.New(t)
	g.Assert(t, "real_article", []byte(content))
}

func TestGetImageLinks(t *testing.T) {
	type args struct {
		filename string
	}

	tests := []struct {
		name           string
		args           args
		wantLinks      map[string]string
		wantCoverImage string
		assertion      assert.ErrorAssertionFunc
	}{
		{
			name: "normal",
			args: args{filename: "./testdata/testdata.md"},
			wantLinks: map[string]string{
				"./image.png":   "",
				"./image-2.png": "",
				"./cv.jpg":      "",
			},
			wantCoverImage: "./cv.jpg",
			assertion:      assert.NoError,
		},
		{
			name:      "not found",
			args:      args{filename: "./testdata/unknown.md"},
			wantLinks: nil,
			assertion: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLinks, gotCoverImage, err := GetImageLinks(tt.args.filename)
			tt.assertion(t, err)
			assert.Equal(t, tt.wantLinks, gotLinks)
			assert.Equal(t, tt.wantCoverImage, gotCoverImage)
		})
	}
}

func TestPrefixLinks(t *testing.T) {
	type args struct {
		links  map[string]string
		prefix string
		force  bool
	}

	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "normal",
			args: args{
				links: map[string]string{
					"./image/image.png":   "image.png",
					"./image/picture.jpg": "",
				},
				prefix: "https://raw.githubusercontent.com/repo/user/",
			},
			want: map[string]string{
				"./image/image.png":   "image.png",
				"./image/picture.jpg": "https://raw.githubusercontent.com/repo/user/./image/picture.jpg",
			},
		},
		{
			name: "force",
			args: args{
				links: map[string]string{
					"./image/image.png":   "image.png",
					"./image/picture.jpg": "",
				},
				prefix: "https://raw.githubusercontent.com/repo/user/",
				force:  true,
			},
			want: map[string]string{
				"./image/image.png":   "https://raw.githubusercontent.com/repo/user/./image/image.png",
				"./image/picture.jpg": "https://raw.githubusercontent.com/repo/user/./image/picture.jpg",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, PrefixLinks(tt.args.links, tt.args.prefix, tt.args.force))
		})
	}
}
