package gitclient

import (
	"os"
	"github.com/libgit2/git2go"
	"fmt"
)

type gitClient struct {
	repoName       string
	parentRepoPath string
}

func GitClient(repoName string, repoPath string) gitClient {
	return gitClient{repoName, repoPath}
}

func (gc gitClient) RepoConfigPath() string {
	return fmt.Sprintf("%s/%s/.git", gc.parentRepoPath, gc.repoName)
}

func (gc gitClient) InitRepo() error {
	err := os.MkdirAll(gc.RepoConfigPath(), os.ModePerm)
	if err != nil {
		return err
	}
	git.InitRepository(gc.RepoConfigPath(), true)
	return nil
}
