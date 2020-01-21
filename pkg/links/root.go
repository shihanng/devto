package links

import (
	"path/filepath"
	"strings"

	"github.com/cockroachdb/errors"
	"gopkg.in/src-d/go-git.v4"
)

// Root is like git ls-files --full-name <filepath>
func Root(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", errors.Wrap(err, "links: get absolute path")
	}

	repo, err := git.PlainOpenWithOptions(filepath.Dir(absPath), &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		return "", errors.Wrap(err, "links: open git")
	}

	wt, err := repo.Worktree()

	var repoRootWorkDir string

	if err == nil {
		workDir := filepath.Dir(absPath)
		repoRootWorkDir = strings.TrimPrefix(workDir, wt.Filesystem.Root())
	}

	return repoRootWorkDir, errors.Wrap(err, "links: get git worktree")
}
