package main

import (
	"../reader"
	"../gitclient"
	"flag"
	"github.com/libgit2/git2go"
	"os/user"
	"strings"
	"time"
)

func main() {
	modelFile, outputRepoPath, gitConfigPath, year, weekOffset := setupArgs()
	commitDataChan, err := reader.ReadCommitDataFromFileToChannel(modelFile)
	checkError(err)
	repoClient, err := gitclient.New(outputRepoPath, readGitConfig(gitConfigPath))
	checkError(err)

	for commitData := range commitDataChan {
		commitTimes := CalculateCommitTimes(commitData, year, weekOffset)
		commitDataAtTimes(repoClient, commitTimes...)
	}

}

func setupArgs() (string, string, string, int, int) {

	inputModel := flag.String("input-model", "./models/mario_head.txt", "File with commit model")
	outputRepoPath := flag.String("output-repo", "./output_repo", "Output repo path")
	gitConfigPath := flag.String("git-config-path", "~/.gitconfig", "Git config file")
	year := flag.Int("year", 2015, "Year of commit messages")
	weekOffset := flag.Int("week-offset", 0, "Offest of the image from the begginning of the year")

	flag.Parse()

	return *inputModel, *outputRepoPath, *gitConfigPath, *year, *weekOffset
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readGitConfig(configPath string) *git.Config {
	usr, err := user.Current()
	checkError(err)
	config, err := git.OpenOndisk(nil, strings.Replace(configPath, "~", usr.HomeDir, 1))
	checkError(err)
	return config
}

func CalculateCommitTimes(commitData reader.CommitData, year int, weekOffset int) []time.Time {
	timeArray := make([]time.Time, commitData.NumberOfCommits)
	for i := 0; i < len(timeArray); i++ {
		timeArray[i] = CalculateCommitTime(commitData, year, weekOffset).Add(time.Duration(i) * time.Minute)
	}
	return timeArray
}

func CalculateCommitTime(commitData reader.CommitData, year int, weekOffset int) time.Time {
	firstSundayOfYearAtNoon := FirstSunday(year).Add(12 * time.Hour)
	shiftedBeginningOfYear := firstSundayOfYearAtNoon.AddDate(0, 0, weekOffset * 7)
	return shiftedBeginningOfYear.AddDate(0, 0, commitData.DayOfWeek + (7 * commitData.WeekOfYear))
}
func FirstSunday(year int) time.Time {
	beginningOfYear := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	daysToLastSunday := int(time.Sunday) - int(beginningOfYear.Weekday())
	return beginningOfYear.AddDate(0, 0, daysToLastSunday)
}

func commitDataAtTimes(gc *gitclient.GitClient, commitTimes ...time.Time) {
	for i := range commitTimes {
		gc.CreateCommitAtDate(commitTimes[i], "Auto-generated commit")
	}
}