package main

import (
	"flag"
	"log"

	"github.com/ushios/golem"
)

var (
	token  = flag.String("token", "", "circle ci artifact build token")
	user   = flag.String("user", "", "github or bitbucket user name")
	repo   = flag.String("repository", "", "respository name")
	branch = flag.String("branch", "master", "branch name")
)

func main() {
	flag.Parse()
	client := golem.NewClient(*token)

	n, err := client.GetLatestSucceededBuildNum(golem.BitBucket, *user, *repo, *branch)
	if err != nil {
		log.Fatalf("faild to get latest build num: %s", err)
	}

	paths, err := client.ListArtifacts(golem.BitBucket, *user, *repo, n)
	if err != nil {
		log.Fatalf("faild to get artifact list: %s", err)
	}

	log.Println(paths)
}
