package reader

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"unicode"
	"sync"
)

const MAX_DAY_OF_WEEK int = 6

type CommitData struct {
	DayOfWeek       int
	WeekOfYear      int
	NumberOfCommits int
}

func ReadCommitDataFromFileToChannel(filepath string) (<-chan CommitData, <-chan bool, error) {
	commits := make(chan CommitData)
	done := make(chan bool)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}

	var wg sync.WaitGroup
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		if i > MAX_DAY_OF_WEEK {
			return nil, nil, errors.New(fmt.Sprintf("File shouldn't have more than %d lines", MAX_DAY_OF_WEEK + 1))
		}
		wg.Add(1)
		go ReadCommitDataFromString(scanner.Text(), i, commits, &wg)
	}
	go allReadNotifier(done, &wg)
	return commits, done, nil
}

func ReadCommitDataFromString(line string, line_index int, commits chan <- CommitData, wg *sync.WaitGroup) {
	defer wg.Done()
	for i, char := range line {
		if unicode.IsDigit(char) {
			commits <- CommitData{line_index, i, int(char) - '0'}
		}
	}
}

func allReadNotifier(done chan <- bool, wg *sync.WaitGroup) {
	wg.Wait()
	done <- true
}