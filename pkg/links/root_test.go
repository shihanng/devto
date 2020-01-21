package links

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRoot(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "example")
	require.NoError(t, err)

	defer os.Remove(tmpfile.Name())

	type args struct {
		path string
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
				path: "./root.go",
			},
			want:      "/pkg/links",
			assertion: assert.NoError,
		},
		{
			name: "not git",
			args: args{
				path: tmpfile.Name(),
			},
			want:      "",
			assertion: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Root(tt.args.path)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
