package repositories

import (
	"fmt"
	"strings"

	"github.com/chrislusf/gleam/util"
	git "gopkg.in/src-d/go-git.v4"
)

type RepositoriesGitReader struct {
	repo *git.Repository
	read bool
}

func New(r *git.Repository) *RepositoriesGitReader {
	return &RepositoriesGitReader{
		repo: r,
		read: false,
	}
}

func (r *RepositoriesGitReader) ReadHeader() (fieldNames []string, err error) {
	return nil, nil
}

/*
root
 |-- id: string (nullable = false)
 |-- urls: array (nullable = false)
 |    |-- element: string (containsNull = false)
 |-- is_fork: boolean (nullable = true)
 |-- repository_path: string (nullable = true)
*/

func (r *RepositoriesGitReader) Read() (row *util.Row, err error) {

	if r.read {
		return nil, fmt.Errorf("repository already read")
	}

	remotes, err := r.repo.Remotes()
	if err != nil {
		return nil, err
	} else {
		r.read = true
	}

	urls := remotes[0].Config().URLs
	id := strings.TrimPrefix(urls[0], "https://")

	return util.NewRow(util.Now(), id, urls), nil
}
