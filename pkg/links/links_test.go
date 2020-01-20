package links

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitHub_Create(t *testing.T) {
	type fields struct {
		Repo   string
		Branch string
	}

	type args struct {
		path string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "normal",
			fields: fields{
				Repo:   "shihanng/devto",
				Branch: "test",
			},
			args: args{
				path: "test.png",
			},
			want: "https://github.com/shihanng/devto/raw/test/test.png",
		},
		{
			name: "ignore .. at beginning",
			fields: fields{
				Repo:   "shihanng/devto",
				Branch: "test",
			},
			args: args{
				path: "../test.png",
			},
			want: "https://github.com/shihanng/devto/raw/test/test.png",
		},
		{
			name: "can have .. inside",
			fields: fields{
				Repo:   "shihanng/devto",
				Branch: "test",
			},
			args: args{
				path: "some/tests/../test.png",
			},
			want: "https://github.com/shihanng/devto/raw/test/some/test.png",
		},
		{
			name: "can have .. inside but not exceed root",
			fields: fields{
				Repo:   "shihanng/devto",
				Branch: "test",
			},
			args: args{
				path: "tests/../../test.png",
			},
			want: "https://github.com/shihanng/devto/raw/test/test.png",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGitHub(tt.fields.Repo, tt.fields.Branch)
			assert.Equal(t, tt.want, g.Create(tt.args.path))
		})
	}
}
