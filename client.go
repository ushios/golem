package golem

import (
	"github.com/jszwedko/go-circleci"
)

// Client interface for golem
type Client interface {
	ListArtifacts(v VersionControlSystem, acc, repo string, buildNum int) ([]string, error)
}

// NewClient interface for golem
func NewClient(token string) Client {
	circleclient := circleci.Client{
		Token: token,
	}

	return &goCircleCIClient{
		client: &circleclient,
	}
}
