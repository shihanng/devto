package article

import (
	"bytes"

	"github.com/cockroachdb/errors"
	"github.com/shihanng/md"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

func SetImageLinks(filename string, images map[string]string) (string, error) {
	parsed, n, err := read(filename)
	if err != nil {
		return "", err
	}

	r := renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(&md.Renderer{}, 100)))

	if err := ast.Walk(n, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering && node.Kind() == ast.KindImage {
			n := node.(*ast.Image)
			if replace, ok := images[string(n.Destination)]; ok {
				n.Destination = []byte(replace)
			}
		}
		return ast.WalkContinue, nil
	}); err != nil {
		return "", errors.Wrap(err, "article: walk ast for Read")
	}

	var buf bytes.Buffer

	if err := r.Render(&buf, parsed.markdownSource, n); err != nil {
		return "", errors.Wrap(err, "article: render markdown")
	}

	parsed.markdownSource = buf.Bytes()

	return parsed.Content()
}

func GetImageLinks(filename string) (map[string]string, error) {
	_, n, err := read(filename)
	if err != nil {
		return nil, err
	}

	images := make(map[string]string)

	if err := ast.Walk(n, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering && node.Kind() == ast.KindImage {
			n := node.(*ast.Image)
			images[string(n.Destination)] = ""
		}
		return ast.WalkContinue, nil
	}); err != nil {
		return nil, errors.Wrap(err, "article: walk ast for GetImageLinks")
	}

	return images, nil
}

func read(filename string) (*Parsed, ast.Node, error) {
	parsed, err := Parse(filename)
	if err != nil {
		return nil, nil, err
	}

	p := parser.NewParser(
		parser.WithBlockParsers(
			[]util.PrioritizedValue{
				util.Prioritized(md.NewRawParagraphParser(), 100),
			}...),
		parser.WithInlineParsers(
			[]util.PrioritizedValue{
				util.Prioritized(parser.NewLinkParser(), 100),
			}...),
	)

	reader := text.NewReader(parsed.markdownSource)

	return parsed, p.Parse(reader), nil
}
