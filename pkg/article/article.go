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

func Read(filename string, images map[string]string) (string, error) {
	parsed, err := Parse(filename)
	if err != nil {
		return "", err
	}

	parsed.markdownSource, err = Render(parsed.markdownSource, images)
	if err != nil {
		return "", err
	}

	return parsed.Content()
}

func Render(body []byte, images map[string]string) ([]byte, error) {
	var buf bytes.Buffer

	reader := text.NewReader(body)

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

	n := p.Parse(reader)
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
		return nil, errors.Wrap(err, "article: walk ast for Render")
	}

	if err := r.Render(&buf, body, n); err != nil {
		return nil, errors.Wrap(err, "article: render markdown")
	}

	return buf.Bytes(), nil
}

func GetImageLinks(filename string) (map[string]string, error) {
	parsed, err := Parse(filename)
	if err != nil {
		return nil, err
	}

	reader := text.NewReader(parsed.markdownSource)

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

	n := p.Parse(reader)

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
