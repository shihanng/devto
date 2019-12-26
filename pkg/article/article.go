package article

import (
	"io/ioutil"

	"github.com/cockroachdb/errors"
	"github.com/shihanng/devto/pkg/frontmatter"
)

func Read(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.Wrap(err, "article: read file")
	}

	fm, md, err := frontmatter.Split(content)
	if err != nil {
		return "", err
	}

	_ = fm

	return string(append(fm, md...)), nil
}
