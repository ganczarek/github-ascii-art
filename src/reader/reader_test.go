package reader_test

import (
	"fmt"
	"testing"
	"time"
	. "."
	"github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"
)

const DEFAULT_TEST_TIMEOUT time.Duration = 5 * time.Second

// it's global, tests cannot run in parallel because of this
var TEST_TIMEOUT_CHAN = make(chan struct{})

func FailTestAfter(duration time.Duration) {
	time.Sleep(duration)
	TEST_TIMEOUT_CHAN <- struct{}{}
}

func Test_ShouldFailToReadFileThatDoesNotExist(t *testing.T) {
	_, _, err := ReadCommitDataFromFileToChannel("./non_existing_file")
	assert.NotNil(t, err, "It should fail to read a file that doesn't exist")
}

// there are only 7 days in a week
func Test_ShouldFailIfAFileHasMoreThan7Lines(t *testing.T) {
	_, _, err := ReadCommitDataFromFileToChannel("./testdata/more_than_8_lines.txt")
	assert.NotNil(t, err, "It should fail to read a file that has more than 8 lines")
}

func Test_ShouldReadSingleCommitDataFromAFile(t *testing.T) {
	expectedData := CommitData{2, 5, 9}
	commits, _, err := ReadCommitDataFromFileToChannel("./testdata/many_chars_single_digit.txt")
	assert.Nil(t, err)
	result := <-commits
	assert.Equal(t, result, expectedData)
}

func Test_ShouldReadAllCommitDataFromAFile(t *testing.T) {
	go FailTestAfter(DEFAULT_TEST_TIMEOUT)
	expectedData := mapset.NewSet(CommitData{0, 0, 1}, CommitData{0, 2, 2}, CommitData{0, 4, 3},
		CommitData{1, 0, 4}, CommitData{1, 2, 5}, CommitData{2, 0, 6})
	commits, _, err := ReadCommitDataFromFileToChannel("./testdata/simple_model.txt")
	assert.Nil(t, err)
	for {
		select {
		case result := <-commits:
			if expectedData.Contains(result) {
				expectedData.Remove(result)
				if expectedData.Cardinality() == 0 {
					return
				}
			} else {
				t.Error(fmt.Sprintf("Channel returned unexpected data %s\n", result))
				t.FailNow()
			}
		case <-TEST_TIMEOUT_CHAN:
			t.Error(fmt.Sprintf("Test timed out. Didn't find expected data %s\n", expectedData.String()))
			t.FailNow()
		}
	}
}

func Test_ReaderShouldNotifyThatIsDoneWriting(t *testing.T) {
	go FailTestAfter(DEFAULT_TEST_TIMEOUT)
	expectedNumerOfCommits := 6
	commits, done, err := ReadCommitDataFromFileToChannel("./testdata/simple_model.txt")
	assert.Nil(t, err)
	commitCount := 0
	for {
		select {
		case <-commits:
			commitCount++
		case <-done:
			return
		case <-TEST_TIMEOUT_CHAN:
			t.Error(fmt.Sprintf("Test timed out. Counted %d commits", commitCount))
			t.FailNow()
		}
	}
	assert.Equal(t, expectedNumerOfCommits, commitCount)
}
