package article

import (
	"bytes"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/gohugoio/hugo/parser/metadecoders"
	"github.com/gohugoio/hugo/parser/pageparser"
	"github.com/mitchellh/mapstructure"
)

// Parse parses the article and divides the content into front matter and markdown.
// Heavily inspired by:
// https://github.com/gohugoio/hugo/blob/94cfdf6befd657e46c9458b23f17d851cd2f7037/commands/convert.go#L207-L250
func Parse(filename string) (*Parsed, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "article: open file")
	}

	defer f.Close()

	result, err := pageparser.Parse(f, pageparser.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "article: parse file")
	}

	var parsed Parsed

	walkFn := func(item pageparser.Item) bool {
		if parsed.frontMatterSource != nil {
			parsed.markdownSource = result.Input()[item.Pos:]
			return false
		} else if item.IsFrontMatter() {
			parsed.frontMatterFormat = metadecoders.FormatFromFrontMatterType(item.Type)
			parsed.frontMatterSource = item.Val
		}

		return true
	}

	result.Iterator().PeekWalk(walkFn)

	metadata, err := metadecoders.Default.UnmarshalToMap(parsed.frontMatterSource, parsed.frontMatterFormat)
	if err != nil {
		return nil, errors.Wrap(err, "article: unmarshal front matter")
	}

	if err := mapstructure.Decode(metadata, &parsed.frontMatter); err != nil {
		return nil, errors.Wrap(err, "article: decode front matter")
	}

	return &parsed, nil
}

type Parsed struct {
	frontMatterFormat metadecoders.Format
	frontMatterSource []byte
	frontMatter       FrontMatter
	markdownSource    []byte
}

// Content merges the front mattter and markdown source and returns it as string.
func (p *Parsed) Content() (string, error) {
	eb := errBuffer{
		b: &bytes.Buffer{},
	}

	eb.WriteString("---\n")
	eb.Write(p.frontMatterSource)
	eb.WriteString("---\n")
	eb.Write(p.markdownSource)

	if eb.err != nil {
		return "", errors.Wrap(eb.err, "article: output parsed content")
	}

	return eb.b.String(), nil
}

// FrontMatter as described in https://dev.to/p/editor_guide
type FrontMatter struct {
	Title        string
	Published    bool
	Description  string
	Tags         string
	CanonicalURL string
	CoverImage   string
	Series       string
}

type errBuffer struct {
	b   *bytes.Buffer
	err error
}

func (eb *errBuffer) WriteString(s string) {
	if eb.err != nil {
		return
	}

	_, eb.err = eb.b.WriteString(s)
}

func (eb *errBuffer) Write(p []byte) {
	if eb.err != nil {
		return
	}

	_, eb.err = eb.b.Write(p)
}
