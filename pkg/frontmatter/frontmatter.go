package frontmatter

import (
	"bufio"
	"bytes"
	"strings"

	"github.com/cockroachdb/errors"
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
