package golem

import (
	"fmt"

	circleci "github.com/jszwedko/go-circleci"
	"github.com/pkg/errors"
)

// goCircleCIClient using jszwedko/go-circleci
type goCircleCIClient struct {
	client *circleci.Client
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
