package main

import (
	"flag"
	"log"

	"github.com/ushios/golem"
)

var (
	token = flag.String("token", "", "circle ci artifact build token")
	user  = flag.String("user", "", "github or bitbucket user name")
	repo  = flag.String("repository", "", "respository name")
)

func main() {
	flag.Parse()
	client := golem.NewClient(*token)

	log.Println(client.ListArtifacts(golem.BitBucket, *user, *repo, 0))
}
