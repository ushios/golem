package golem

import (
	"net/url"

	"github.com/jszwedko/go-circleci"
)

// Client interface for golem
type Client interface {
	GetLatestSucceededBuildNum(v VersionControlSystem, acc, repo, branch string) (int, error)
	ListArtifacts(v VersionControlSystem, acc, repo string, buildNum int) ([]string, error)
}

// NewClient interface for golem
func NewClient(token string) Client {
	circleclient := circleci.Client{
		Token: token,
		BaseURL: &url.URL{
			Host:   "circleci.com",
			Scheme: "https",
			Path:   "/api/v1.1/",
		},
	}

	return &goCircleCIClient{
		client: &circleclient,
	}
}
