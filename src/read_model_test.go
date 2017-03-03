package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldFailToReadFileThatDoesNotExist(t *testing.T) {
	_, err := ReadCommitDataFromFileToChannel("./non_existing_file")
	assert.NotNil(t, err, "It should fail to read a file that doesn't exist")
}

// there are only 7 days in a week
func Test_ShouldFailIfAFileHasMoreThan7Lines(t *testing.T) {
	_, err := ReadCommitDataFromFileToChannel("./test-resources/more_than_8_lines.txt")
	assert.NotNil(t, err, "It should fail to read a file that has more than 8 lines")
}

func Test_ShouldReadOnlyDigits(t *testing.T) {
	expectedData := CommitData{2, 5, 9}
	commits, err := ReadCommitDataFromFileToChannel("./test-resources/many_chars_single_digit.txt")
	assert.Nil(t, err)
	result := <-commits
	assert.Equal(t, result, expectedData)
}

