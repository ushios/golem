package golem

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	circleci "github.com/jszwedko/go-circleci"
	"github.com/pkg/errors"
)

var (
	ErrBuildNotFound = errors.New("build not found")
)

type CircleCIArtifact struct {
	path string
	u    url.URL
}

// Path to artifact
func (c *CircleCIArtifact) Path() string {
	return c.path
}

// URL of artifact
func (c *CircleCIArtifact) URL() url.URL {
	return c.u
}

// goCircleCIClient using jszwedko/go-circleci
type goCircleCIClient struct {
	client *circleci.Client
}

func (c *goCircleCIClient) Account(v VersionControlSystem, username string) string {
	return fmt.Sprintf("%s/%s", v, username)
}

// GetLatestSucceededBuildNum .
func (c *goCircleCIClient) GetLatestSucceededBuildNum(v VersionControlSystem, acc, repo, branch string) (int, error) {
	acc = c.Account(v, acc)
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
func (c *goCircleCIClient) ListArtifacts(v VersionControlSystem, acc, repo string, buildNum int) ([]Artifact, error) {
	acc = c.Account(v, acc)

	list, err := c.client.ListBuildArtifacts(acc, repo, buildNum)
	if err != nil {
		return nil, errors.Wrap(err, "circleci.ListBuildArtifacts got error")
	}

	res := make([]Artifact, len(list))
	for i, a := range list {
		u, err := url.Parse(a.URL)
		if err != nil {
			return nil, errors.Wrap(err, "url.Parse got error")
		}
		res[i] = &CircleCIArtifact{
			path: a.Path,
			u:    *u,
		}
	}

	return res, nil
}

func (c *goCircleCIClient) GetArtifact(a Artifact, output string) (io.ReadCloser, error) {
	u := a.URL()
	q := u.Query()
	q.Add("circle-token", c.client.Token)
	u.RawQuery = q.Encode()

	if _, err := os.Stat(output); os.IsNotExist(err) {
		if err := os.MkdirAll(output, 0750); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("faild to create directory %s ", output))
		}
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, errors.Wrap(err, "faild to http get file")
	}

	return resp.Body, nil
}
