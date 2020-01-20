package links

import (
	"net/url"
	"path"
)

type GitHub struct {
	Repo   string
	Branch string
}

func NewGitHub(repo, branch string) *GitHub {
	return &GitHub{
		Repo:   repo,
		Branch: branch,
	}
}

const (
	githubScheme = "https"
	githubHost   = "github.com"
)

func (g *GitHub) Create(relPath string) string {
	relPath = path.Join("/", relPath)
	p := path.Join(g.Repo, "raw", g.Branch, relPath)

	url := url.URL{
		Scheme: githubScheme,
		Host:   githubHost,
		Path:   p,
	}

	return url.String()
}
