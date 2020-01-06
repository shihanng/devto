package article

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/cockroachdb/errors"
	"github.com/shihanng/devto/pkg/frontmatter"
	"github.com/shihanng/md"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

func Read(filename string, images map[string]string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.Wrap(err, "article: read file")
	}

	fmBytes, mdBytes, err := frontmatter.Split(content)
	if err != nil {
		return "", err
	}

	mdBytes = Render(mdBytes, images)

	return string(append(fmBytes, mdBytes...)), nil
}

func Render(body []byte, images map[string]string) []byte {
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

	ast.Walk(n, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering && node.Kind() == ast.KindImage {
			n := node.(*ast.Image)
			if replace, ok := images[string(n.Destination)]; ok {
				n.Destination = []byte(replace)
			}
		}
		return ast.WalkContinue, nil
	})

	err := r.Render(&buf, body, n)
	if err != nil {
		fmt.Println(err)
	}
	return buf.Bytes()
}
