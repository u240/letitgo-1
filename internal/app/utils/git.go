package utils

import (
	"context"
	"errors"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	giturls "github.com/whilp/git-urls"
	"golang.org/x/oauth2"
)

func Run(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	bts, err := cmd.CombinedOutput()
	logrus.WithField("output", string(bts)).Trace("git result")
	if err != nil {
		return "", errors.New(string(bts))
	}

	output := strings.Replace(strings.Split(string(bts), "\n")[0], "'", "", -1)
	if err != nil {
		err = errors.New(strings.TrimSuffix(err.Error(), "\n"))
	}
	return output, err
}

func PublishFile(repoURL, path, content, message string) error {
	url, err := giturls.Parse(repoURL)
	if err != nil {
		return err
	}

	switch url.Hostname() {
	case "github.com":
		return publishFileGithub(url, path, content, message)
	}

	return errors.New("Unsupported platform")
}

func publishFileGithub(url *url.URL, path, content, message string) error {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return errors.New("Make sure to set GITHUB_TOKEN")
	}

	owner := filepath.Dir(url.Path)
	basename := filepath.Base(url.Path)
	repo := strings.TrimSuffix(basename, filepath.Ext(basename))

	ctx := context.Background()
	tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))

	client := github.NewClient(tc)

	var sha string
	fileContent, _, _, err := client.Repositories.GetContents(
		ctx,
		owner,
		repo,
		path,
		nil,
	)
	if err == nil {
		sha = fileContent.GetSHA()
	}

	_, _, err = client.Repositories.CreateFile(
		ctx,
		owner,
		repo,
		path,
		&github.RepositoryContentFileOptions{
			SHA:     &sha,
			Message: &message,
			Content: []byte(content),
		},
	)

	return err
}