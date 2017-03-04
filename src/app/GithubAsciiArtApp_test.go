package main_test

import (
	. "."
	"../reader"
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

func Test_ShouldCalculateTimeForFirstMondayOfAYear(t *testing.T) {
	commitData := reader.CommitData{0, 0, 1}
	expectedTime := time.Date(2015, 0, 0, 12, 0, 0, 0, time.UTC)

	result := CalculateCommitTime(commitData, 2015, 0)

	assert.Equal(t, expectedTime, result)
}

func Test_ShouldCalculateTimeWithOffset(t *testing.T) {
	commitData := reader.CommitData{0, 0, 1}
	offset := 5
	expectedTime := time.Date(2014, 0, 7 * offset, 12, 0, 0, 0, time.UTC)

	result := CalculateCommitTime(commitData, 2014, offset)

	assert.Equal(t, expectedTime, result)
}