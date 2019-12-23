package article

import (
	"io/ioutil"

	"github.com/cockroachdb/errors"
)

func Read(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.Wrap(err, "article: read file")
	}

	return string(content), nil
}
