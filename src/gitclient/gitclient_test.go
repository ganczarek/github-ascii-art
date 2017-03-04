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

const TEST_USER_NAME = "Test User Name"
const TEST_USER_EMAIL = "user@test.com"

var TEST_TEMP_DIR = os.TempDir() + "/go_gitclient_test"
var TEST_TEMP_CONFIG = TEST_TEMP_DIR + "/temp_test.gitconfig"

func Test_ClientConstruction_ShouldFail_WhenNoConfigFile(t *testing.T) {
	repoPath := fmt.Sprintf("%s/go_gitclient_test/%s", os.TempDir(), uuid.New().String())
	config, err := git.OpenOndisk(nil, "non-esiting-git-config-file")

	_, err = New(repoPath, config)

	assert.NotNil(t, err)
}

func Test_ClientConstruction_WhenConfigFileExists(t *testing.T) {
	repoPath := fmt.Sprintf("%s/go_gitclient_test/%s", os.TempDir(), uuid.New().String())
	config := setupTestConfig(t)
	defer cleanupTestConfig()

	gitClient, err := New(repoPath, config)

	assert.Nil(t, err)
	assert.Equal(t, gitClient.UserName, TEST_USER_NAME)
	assert.Equal(t, gitClient.UserEmail, TEST_USER_EMAIL)
}

func Test_ClientConstruction_ShouldReturnErrorIfCannotCreateRepo(t *testing.T) {
	config := setupTestConfig(t)
	defer cleanupTestConfig()
	// should not e able to touch root home dir
	incorrectPath := fmt.Sprintf("/root/go_gitclient_test/%s", uuid.New().String())

	_, err := New(incorrectPath, config)

	assert.NotNil(t, err)
	_, err = os.Stat(fmt.Sprintf("%s/.git", incorrectPath))
	assert.NotNil(t, err)
	assert.True(t, os.IsNotExist(err))
}

func Test_InitRepo_ShouldCreateANewRepo(t *testing.T) {
	gitClient := setupTestRepoAndClient(t)
	defer cleanupTestRepo()

	_, err := os.Stat(fmt.Sprintf("%s/.git", gitClient.RepoPath))
	assert.Nil(t, err)
}

func Test_Commit_ShouldNotReturnErrorWhenCreateCommitInThePast(t *testing.T) {
	gitClient := setupTestRepoAndClient(t)
	defer cleanupTestRepo()
	date := time.Date(2016, time.June, 6, 12, 0, 0, 0, time.UTC)

	// repository should be empty before commit
	isRepoEmpty, err := gitClient.Repo.IsEmpty()
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, isRepoEmpty)

	err = gitClient.CreateCommitAtDate(date, "Test commit message")

	assert.Nil(t, err)
	isRepoEmpty, err = gitClient.Repo.IsEmpty()
	if err != nil {
		t.Fatal(err)
	}
	assert.False(t, isRepoEmpty)
}

func setupTestConfig(t *testing.T) *git.Config {
	c, err := git.OpenOndisk(nil, TEST_TEMP_CONFIG)
	if err != nil {
		t.Fatalf("Failed to open %s", TEST_TEMP_CONFIG)
	}
	err = c.SetString("user.name", TEST_USER_NAME)
	if err != nil {
		t.Fatalf("Failed to set user.name value in %d", TEST_TEMP_CONFIG)
	}
	err = c.SetString("user.email", TEST_USER_EMAIL)
	if err != nil {
		t.Fatalf("Failed to set user.email value in %s", TEST_TEMP_CONFIG)
	}
	return c
}

func cleanupTestConfig() {
	os.Remove(TEST_TEMP_CONFIG)
}

func setupTestRepoAndClient(t *testing.T) *GitClient {
	config := setupTestConfig(t)
	repoPath := fmt.Sprintf("%s/%s/%s", TEST_TEMP_DIR, uuid.New().String(), "TestRepo")
	gitClient, err := New(repoPath, config)
	if err != nil {
		t.Fatalf("Failed to create git client!\n%s", err)
	}
	return gitClient
}

func cleanupTestRepo() {
	cleanupTestConfig()
	os.Remove(TEST_TEMP_DIR)
}