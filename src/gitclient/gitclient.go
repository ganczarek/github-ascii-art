package gitclient

import (
	"os"
	"github.com/libgit2/git2go"
	"fmt"
	"time"
	"io/ioutil"
)

type gitClient struct {
	RepoPath  string
	UserName  string
	UserEmail string
}

func GitClient(repoPath string, config *git.Config) (*gitClient, error) {
	authorEmail, err := config.LookupString("user.email")
	if err != nil {
		return nil, err
	}
	authorName, err := config.LookupString("user.name")
	if err != nil {
		return nil, err
	}
	return &gitClient{repoPath, authorName, authorEmail}, nil
}

func (gc gitClient) RepoConfigPath() string {
	return fmt.Sprintf("%s/.git", gc.RepoPath)
}

func (gc gitClient) InitRepo() error {
	err := os.MkdirAll(gc.RepoConfigPath(), os.ModePerm)
	if err != nil {
		return err
	}
	_, err = git.InitRepository(gc.RepoConfigPath(), false)
	return err
}

func (gc gitClient) CreateCommitAtDate(commitDate time.Time, commitMessage string) error {
	tmpfile := "README"
	err := ioutil.WriteFile(gc.RepoPath + "/" + tmpfile, []byte("foo\n"), 0644)
	repo, err := git.OpenRepository(gc.RepoConfigPath())
	if err != nil {
		return err
	}
	sig := &git.Signature{Name: gc.UserName, Email: gc.UserEmail, When: commitDate}
	idx, err := repo.Index()
	err = idx.AddByPath("README")
	if err != nil {
		return err
	}

	err = idx.Write()
	if err != nil {
		return err
	}

	treeId, err := idx.WriteTree()
	if err != nil {
		return err
	}

	tree, err := repo.LookupTree(treeId)
	if err != nil {
		return err
	}
	_, err = repo.CreateCommit("HEAD", sig, sig, commitMessage, tree)
	if err != nil {
		return err
	}
	return nil
}