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

func SetImageLinks(filename string, images map[string]string, coverImage, prefix string) (string, error) {
	parsed, n, err := read(filename)
	if err != nil {
		return "", err
	}

	if coverImage != "" {
		parsed.frontMatter.CoverImage = coverImage
	} else if parsed.frontMatter.CoverImage != "" {
		parsed.frontMatter.CoverImage = prefix + parsed.frontMatter.CoverImage
	}

	r := renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(&mdRender{}, 100)))

	if err := ast.Walk(n, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering && node.Kind() == ast.KindImage {
			n := node.(*ast.Image)
			if replace, ok := images[string(n.Destination)]; ok && replace != "" {
				n.Destination = []byte(replace)
			} else {
				n.Destination = []byte(prefix + string(n.Destination))
			}
		}
		return ast.WalkContinue, nil
	}); err != nil {
		return "", errors.Wrap(err, "article: walk ast for Read")
	}

	var buf bytes.Buffer

	if err := r.Render(&buf, parsed.content, n); err != nil {
		return "", errors.Wrap(err, "article: render markdown")
	}

	parsed.content = buf.Bytes()

	return parsed.Content()
}

func CoverImageUntouch(original string) string {
	return original
}

func CoverImagePrefixed(prefix string) func(string) string {
	return func(original string) string {
		return prefix + original
	}
}

func GetImageLinks(filename string) (map[string]string, string, error) {
	parsed, n, err := read(filename)
	if err != nil {
		return nil, "", err
	}

	images := make(map[string]string)

	if parsed.frontMatter.CoverImage != "" {
		images[parsed.frontMatter.CoverImage] = ""
	}

	if err := ast.Walk(n, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering && node.Kind() == ast.KindImage {
			n := node.(*ast.Image)
			images[string(n.Destination)] = ""
		}
		return ast.WalkContinue, nil
	}); err != nil {
		return nil, "", errors.Wrap(err, "article: walk ast for GetImageLinks")
	}

	return images, parsed.frontMatter.CoverImage, nil
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

	reader := text.NewReader(parsed.content)

	return parsed, p.Parse(reader), nil
}

func PrefixLinks(links map[string]string, prefix string, force bool) map[string]string {
	results := make(map[string]string, len(links))

	for k, v := range links {
		if v != "" && !force {
			results[k] = v
			continue
		}

		results[k] = prefix + k
	}

	return results
}

type mdRender struct{}

// RegisterFuncs implements github.com/yuin/goldmark/renderer NodeRenderer.RegisterFuncs.
func (r *mdRender) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindParagraph, md.RawRenderParagraph)

	reg.Register(ast.KindImage, md.RenderImage)
	reg.Register(ast.KindLink, md.RenderLink)
	reg.Register(ast.KindText, md.RawRenderText)
}
