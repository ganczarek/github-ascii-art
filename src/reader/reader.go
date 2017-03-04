package reader

import (
	"bufio"
	"os"
	"errors"
	"fmt"
	"unicode"
)

const MAX_DAY_OF_WEEK int = 6

type CommitData struct {
	DayOfWeek       int
	WeekOfYear      int
	NumberOfCommits int
}

func ReadCommitDataFromFileToChannel(filepath string) (<-chan CommitData, error) {
	commits := make(chan CommitData)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		if i > MAX_DAY_OF_WEEK {
			return nil, errors.New(fmt.Sprintf("File shouldn't have more than %d lines", MAX_DAY_OF_WEEK + 1)	)
		}
		go ReadCommitDataFromString(scanner.Text(), i, commits)
	}
	return commits, nil
}

func ReadCommitDataFromString(line string, line_index int, commits chan <- CommitData) {
	for i, char := range line {
		if unicode.IsDigit(char) {
			commits <- CommitData{line_index, i, int(char) - '0'}
		}
	}
}