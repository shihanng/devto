package frontmatter

import (
	"bufio"
	"bytes"
	"strings"

	"github.com/cockroachdb/errors"
	"gopkg.in/yaml.v2"
)

const dashes = `---`

// Split separates the Front Matter from the actual Markdown content.
func Split(content []byte) ([]byte, []byte, error) {
	if !bytes.HasPrefix(content, []byte(dashes)) {
		return nil, content, nil
	}

	r := bytes.NewReader(content)

	scanner := bufio.NewScanner(r)
	scanner.Split(scanLines)

	fm := &bytes.Buffer{}

	for scanner.Scan() {
		line := scanner.Text()

		if _, err := fm.WriteString(line); err != nil {
			return nil, nil, errors.Wrap(err, "frontmatter: writing front matter")
		}

		if strings.TrimSpace(line) == dashes && fm.Len() > 4 {
			break
		}
	}

	md := &bytes.Buffer{}

	for scanner.Scan() {
		if _, err := md.Write(scanner.Bytes()); err != nil {
			return nil, nil, errors.Wrap(err, "frontmatter: writing markdown")
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, errors.Wrap(err, "frontmatter: scanning content")
	}

	return fm.Bytes(), md.Bytes(), nil
}

// Modified from https://golang.org/src/bufio/scan.go?s=11802:11880#L335
// to include newline character
func scanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, data[0 : i+1], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

// FrontMatter follows the specifications described in https://dev.to/p/editor_guide
type FrontMatter struct {
	Title        string
	Published    bool
	Description  string
	Tags         string
	CanonicalURL string
	CoverImage   string
	Series       string
}

func GetFrontMatter(data []byte) (*FrontMatter, error) {
	fm := FrontMatter{}

	if err := yaml.Unmarshal(data, &fm); err != nil {
		return nil, errors.Wrap(err, "frontmatter: unmarshal front matter part")
	}

	return &fm, nil
}
