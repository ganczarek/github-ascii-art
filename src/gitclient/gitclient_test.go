package gitclient_test

import (
	"testing"
	. "."
	"github.com/google/uuid"
	"fmt"
	"os"
	"github.com/stretchr/testify/assert"
	"time"
	"github.com/libgit2/git2go"
)

func Test_InitRepo_ShouldReturnErrorIfCannotCreateRepo(t *testing.T) {
	repoName := "IncorrectTestRepo"
	// should not e able to touch root home dir
	config := setupConfig(t)
	defer cleanupConfig()
	incorrectPath := fmt.Sprintf("/root/go_gitclient_test/%s", uuid.New().String())
	gitClient, err := GitClient(repoName, incorrectPath, config)
	if err != nil {
		t.Fatal("Failed to create git client!")
	}

	gitClient.InitRepo()

	_, err = os.Stat(fmt.Sprintf("%s/%s/.git", incorrectPath, repoName))
	assert.NotNil(t, err)
	assert.True(t, os.IsNotExist(err))
}

func Test_InitRepo_ShouldCreateANewRepo(t *testing.T) {
	repoName := "TestRepo"
	config := setupConfig(t)
	defer cleanupConfig()
	repoPath := fmt.Sprintf("%s/go_gitclient_test/%s", os.TempDir(), uuid.New().String())
	gitClient, err := GitClient(repoName, repoPath, config)
	if err != nil {
		t.Fatal("Failed to create git client!")
	}

	gitClient.InitRepo()

	_, err = os.Stat(fmt.Sprintf("%s/%s/.git", repoPath, repoName))
	assert.Nil(t, err)
}

func Test_Commit_ShouldNotReturnErrorWhenCreateCommitInThePast(t *testing.T) {
	repoName := "TestRepo"
	config := setupConfig(t)
	defer cleanupConfig()
	repoPath := fmt.Sprintf("%s/go_gitclient_test/%s", os.TempDir(), uuid.New().String())
	gitClient, err := GitClient(repoName, repoPath, config)
	if err != nil {
		t.Fatal("Failed to create git client!")
	}
	date := time.Date(2016, time.June, 6, 12, 0, 0, 0, time.UTC)

	gitClient.InitRepo()
	err = gitClient.CreateCommitAtDate(date, "Test commit message")
	assert.Nil(t, err)
}

var tempConfig = os.TempDir() + "/temp.gitconfig"

func setupConfig(t *testing.T) *git.Config {
	c, err := git.OpenOndisk(nil, tempConfig)
	if err != nil {
		t.Fatalf("Failed to open %s", tempConfig)
	}

	err = c.SetString("user.name", "Test User Name")
	if err != nil {
		t.Fatal("Failed to set user.name value in test config")
	}
	err = c.SetString("user.email", "user@test.com")
	if err != nil {
		t.Fatal("Failed to set user.email value in test config")
	}
	return c
}

func cleanupConfig() {
	os.Remove(tempConfig)
}
