package golem

import (
	"errors"
	"strings"
)

var (
	ErrVersionControlSystemNotFound = errors.New("version control system not found")
)

const (
	Github    VersionControlSystem = "github"
	BitBucket VersionControlSystem = "bitbucket"
)

type VersionControlSystem string

func NewVersionControlSystemFromString(s string) (VersionControlSystem, error) {
	s = strings.ToLower(s)
	switch s {
	case "bitbucket", "bb":
		return BitBucket, nil
	case "github", "gh":
		return Github, nil
	}

	return Github, ErrVersionControlSystemNotFound
}

func (v VersionControlSystem) String() string {
	return string(v)
}
