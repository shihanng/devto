package article

import (
	"bytes"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/gohugoio/hugo/parser"
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

	cfm, err := pageparser.ParseFrontMatterAndContent(f)
	if err != nil {
		return nil, errors.Wrap(err, "article: parse file")
	}

	parsed := Parsed{
		content:           cfm.Content,
		frontMatterFormat: cfm.FrontMatterFormat,
	}

	if err := mapstructure.Decode(cfm.FrontMatter, &parsed.frontMatter); err != nil {
		return nil, errors.Wrap(err, "article: decode front matter")
	}

	return &parsed, nil
}

type Parsed struct {
	content           []byte
	frontMatterFormat metadecoders.Format
	frontMatter       FrontMatter
}

// Content merges the front mattter and markdown source and returns it as string.
func (p *Parsed) Content() (string, error) {
	eb := errBuffer{
		b: &bytes.Buffer{},
	}

	var buf bytes.Buffer

	if err := parser.InterfaceToFrontMatter(p.frontMatter, p.frontMatterFormat, &buf); err != nil {
		return "", errors.Wrap(eb.err, "article: marshal frontMatter to YAML")
	}

	eb.Write(buf.Bytes())
	eb.Write(p.content)

	if eb.err != nil {
		return "", errors.Wrap(eb.err, "article: output parsed content")
	}

	return eb.b.String(), nil
}

// FrontMatter as described in https://dev.to/p/editor_guide
type FrontMatter struct {
	Title        string `yaml:"title,omitempty"`
	Published    *bool  `yaml:"published,omitempty"`
	Description  string `yaml:"description,omitempty"`
	Tags         string `yaml:"tags,omitempty"`
	CanonicalURL string `yaml:"canonical_url,omitempty" mapstructure:"canonical_url"`
	CoverImage   string `yaml:"cover_image,omitempty" mapstructure:"canonical_url"`
	Series       string `yaml:"series,omitempty"`
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
