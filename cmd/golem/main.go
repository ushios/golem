package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
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

	if *user == "" || *repo == "" || *branch == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

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

			body, err := client.GetArtifact(a)
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

			if err := createDirectory(p); err != nil {
				log.Fatalf("faild to create directory: %s", err)
			}

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

func createDirectory(path string) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0750); err != nil {
			return errors.Wrap(err, "faild to create directory")
		}
	}
	return nil
}
