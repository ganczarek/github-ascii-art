package main_test

import (
	. "."
	"../reader"
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

func Test_GetSundayOfFirstWeekOfAYear_WhenItsPreviousYear(t *testing.T) {
	expectedTime := time.Date(2015, 12, 27, 0, 0, 0, 0, time.UTC)

	result := FirstSunday(2016)

	assert.Equal(t, expectedTime, result)
}

func Test_GetSundayOfFirstWeekOfAYear_WhenItsTheSameYear(t *testing.T) {
	expectedTime := time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)

	result := FirstSunday(2017)

	assert.Equal(t, expectedTime, result)
}

func Test_ShouldCalculateTimeForFirstSundayOfGithubContributionYear(t *testing.T) {
	commitData := reader.CommitData{0, 0, 1}
	expectedTime := time.Date(2015, 12, 27, 12, 0, 0, 0, time.UTC)

	result := CalculateCommitTime(commitData, 2016, 0)

	assert.Equal(t, expectedTime, result)
}

func Test_ShouldCalculateTimeForFirstSundayOfAYearWithOffset(t *testing.T) {
	commitData := reader.CommitData{0, 0, 1}
	offset := 2
	expectedTime := time.Date(2013, 12, 29, 12, 0, 0, 0, time.UTC).AddDate(0, 0, 7 * offset)

	result := CalculateCommitTime(commitData, 2014, offset)

	assert.Equal(t, expectedTime, result)
}