package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/ushios/golem"
)

var (
	token  = flag.String("token", "", "circle ci artifact build token")
	vcs    = flag.String("vcs", "github", "github or bitbucket")
	user   = flag.String("user", "", "github or bitbucket user name")
	repo   = flag.String("repository", "", "respository name")
	branch = flag.String("branch", "master", "branch name")
	prefix = flag.String("prefix", "", "download file's prefix")
	output = flag.String("output", "./", "output files destination directory")
)

func main() {
	flag.Parse()
	s, err := golem.NewVersionControlSystemFromString(*vcs)
	if err != nil {
		log.Fatalf("faild to parse vcs name: %s", err)
	}

	client := golem.NewClient(*token)

	n, err := client.GetLatestSucceededBuildNum(s, *user, *repo, *branch)
	if err != nil {
		log.Fatalf("faild to get latest build num: %s", err)
	}

	artifacts, err := client.ListArtifacts(s, *user, *repo, n)
	if err != nil {
		log.Fatalf("faild to get artifact list: %s", err)
	}

	for _, a := range artifacts {
		if strings.HasPrefix(a.Path(), *prefix) {
			log.Println(a.Path())
			body, err := client.GetArtifact(a, *output)
			if err != nil {
				log.Fatalf("faild to get artifact: %s", err)
			}
			defer body.Close()

			b, err := ioutil.ReadAll(body)
			if err != nil {
				log.Fatalf("faild to read body: %s", err)
			}

			p := path.Join(*output, a.Path())
			log.Println("save to:", p)
			file, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.Fatalf("faild to create file: %s", err)
			}

			if _, err := file.Write(b); err != nil {
				log.Fatalf("faild to write file: %s", err)
			}
		}
	}
}
