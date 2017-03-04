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

func Test_ClientConstruction_ShouldFailIfNoConfigFile(t *testing.T) {
	repoPath := fmt.Sprintf("%s/go_gitclient_test/%s", os.TempDir(), uuid.New().String())
	config, err := git.OpenOndisk(nil, "non-esiting-git-config-file")
	_, err = GitClient(repoPath, config)
	assert.NotNil(t, err)
}

func Test_ClientConstruction_WhenConfigFileExists(t *testing.T) {
	repoPath := fmt.Sprintf("%s/go_gitclient_test/%s", os.TempDir(), uuid.New().String())
	config := setupConfig(t)
	defer cleanupConfig()
	gitClient, err := GitClient(repoPath, config)
	assert.Nil(t, err)
	assert.Equal(t, gitClient.UserName, TEST_USER_NAME)
	assert.Equal(t, gitClient.UserEmail, TEST_USER_EMAIL)
}

func Test_InitRepo_ShouldReturnErrorIfCannotCreateRepo(t *testing.T) {
	repoName := "IncorrectTestRepo"
	// should not e able to touch root home dir
	config := setupConfig(t)
	defer cleanupConfig()
	incorrectPath := fmt.Sprintf("/root/go_gitclient_test/%s", uuid.New().String())
	gitClient, err := GitClient(incorrectPath, config)
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
	repoPath := fmt.Sprintf("%s/go_gitclient_test/%s/%s", os.TempDir(), uuid.New().String(), repoName)
	gitClient, err := GitClient(repoPath, config)
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
	repoPath := fmt.Sprintf("%s/go_gitclient_test/%s/%s", os.TempDir(), uuid.New().String(), repoName)
	gitClient, err := GitClient(repoPath, config)
	if err != nil {
		t.Fatal("Failed to create git client!")
	}
	date := time.Date(2016, time.June, 6, 12, 0, 0, 0, time.UTC)

	gitClient.InitRepo()
	err = gitClient.CreateCommitAtDate(date, "Test commit message")
	assert.Nil(t, err)
}

var TEMP_TEST_CONFIG = os.TempDir() + "/temp_test.gitconfig"
const TEST_USER_NAME = "Test User Name"
const TEST_USER_EMAIL = "user@test.com"

func setupConfig(t *testing.T) *git.Config {
	c, err := git.OpenOndisk(nil, TEMP_TEST_CONFIG)
	if err != nil {
		t.Fatalf("Failed to open %s", TEMP_TEST_CONFIG)
	}
	err = c.SetString("user.name", TEST_USER_NAME)
	if err != nil {
		t.Fatalf("Failed to set user.name value in %d", TEMP_TEST_CONFIG)
	}
	err = c.SetString("user.email", TEST_USER_EMAIL)
	if err != nil {
		t.Fatalf("Failed to set user.email value in %s", TEMP_TEST_CONFIG)
	}
	return c
}

func cleanupConfig() {
	os.Remove(TEMP_TEST_CONFIG)
}
