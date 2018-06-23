package golem

import (
	"fmt"

	circleci "github.com/jszwedko/go-circleci"
	"github.com/pkg/errors"
)

var (
	ErrBuildNotFound = errors.New("build not found")
)

// goCircleCIClient using jszwedko/go-circleci
type goCircleCIClient struct {
	client *circleci.Client
}

// GetLatestSucceededBuildNum .
func (c *goCircleCIClient) GetLatestSucceededBuildNum(v VersionControlSystem, acc, repo, branch string) (int, error) {
	acc = fmt.Sprintf("%s/%s", v, acc)
	list, err := c.client.ListRecentBuildsForProject(acc, repo, branch, "success", 1, 0)
	if err != nil {
		return 0, errors.Wrap(err, "circleci.ListRecentBuildsForProject got error")
	}

	if len(list) < 1 {
		return 0, ErrBuildNotFound
	}

	return list[0].BuildNum, nil
}

// ListArtifacts path
func (c *goCircleCIClient) ListArtifacts(v VersionControlSystem, acc, repo string, buildNum int) ([]string, error) {
	acc = fmt.Sprintf("%s/%s", v, acc)

	list, err := c.client.ListBuildArtifacts(acc, repo, buildNum)
	if err != nil {
		return nil, errors.Wrap(err, "circleci.ListBuildArtifacts got error")
	}

	res := make([]string, len(list))
	for i, a := range list {
		res[i] = a.Path
	}
	return res, nil
}
