package golem

import "testing"

func TestNewVersionControlSystemFromString(t *testing.T) {
	table := []struct {
		s      string
		err    error
		result VersionControlSystem
	}{
		{
			s:      "gh",
			err:    nil,
			result: Github,
		},
		{
			s:      "bb",
			err:    nil,
			result: BitBucket,
		},
		{
			s:      "bitbucket",
			err:    nil,
			result: BitBucket,
		},
		{
			s:      "github",
			err:    nil,
			result: Github,
		},
		{
			s:      "GitHub",
			err:    nil,
			result: Github,
		},
	}

	for _, row := range table {
		res, err := NewVersionControlSystemFromString(row.s)
		if err != row.err {
			t.Fatalf("err exptected (%s) but (%s)", row.err, err)
		}

		if row.result != res {
			t.Errorf("result expected (%s) but (%s)", row.result, res)
		}
	}

}
