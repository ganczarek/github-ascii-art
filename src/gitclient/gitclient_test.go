package gitclient_test

import (
	"testing"
	. "."
	"github.com/google/uuid"
	"fmt"
	"os"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldReturnErrorIfCannotCreateRepo(t *testing.T) {
	repoName := "IncorrectTestRepo"
	// should not e able to touch root home dir
	incorrectPath := fmt.Sprintf("/root/go_gitclient_test/%s", uuid.New().String())

	gitClient := GitClient(repoName, incorrectPath)
	gitClient.InitRepo()

	_, err := os.Stat(fmt.Sprintf("%s/%s/.git", incorrectPath, repoName))
	assert.NotNil(t, err)
	assert.True(t, os.IsNotExist(err))
}

func Test_ShouldCreateANewRepo(t *testing.T) {
	repoName := "TestRepo"
	repoPath := fmt.Sprintf("%s/go_gitclient_test/%s", os.TempDir(), uuid.New().String())

	gitClient := GitClient(repoName, repoPath)
	gitClient.InitRepo()

	_, err := os.Stat(fmt.Sprintf("%s/%s/.git", repoPath, repoName))
	assert.Nil(t, err)
}