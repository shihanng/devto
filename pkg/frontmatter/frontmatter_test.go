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
