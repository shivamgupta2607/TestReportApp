package main

import (
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to Test Report app")
	filepath := "resources/file.xlsx"
	UserRepoSheetName := "Sheet1"
	TestRepoSheetName := "Sheet2"
	//filepath := "/Users/shivam.gupta/projects/go/personal/TestReportApp/resources/SampleTestFile.xlsx"
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	userRepoRows := f.GetRows(UserRepoSheetName)
	testCaseRows := f.GetRows(TestRepoSheetName)
	tetCaseRepoRow := testCaseRows[1]
	for i, row := range userRepoRows {
		if i == 0 {
			continue
		}
		ranCell := "B" + strconv.Itoa(i+1)
		failedCell := "C" + strconv.Itoa(i+1)
		testErrorCell := "D" + strconv.Itoa(i+1)
		skippedCell := "E" + strconv.Itoa(i+1)
		successPercentCell := "F" + strconv.Itoa(i+1)
		errorCell := "G" + strconv.Itoa(i+1)
		if f.GetCellValue(UserRepoSheetName, ranCell) != "" {
			log.Println(fmt.Sprintf("Not running test case for {%s}", row[0]))
		} else {
			log.Println(fmt.Sprintf("Going to run test cases for {%s}", row[0]))
			output, err := processRow(row, tetCaseRepoRow)
			if err != nil {
				f.SetCellValue(UserRepoSheetName, errorCell, err.Error())
				f.SetCellValue(UserRepoSheetName, ranCell, 0)
				f.SetCellValue(UserRepoSheetName, failedCell, 0)
				f.SetCellValue(UserRepoSheetName, testErrorCell, 0)
				f.SetCellValue(UserRepoSheetName, skippedCell, 0)
				f.SetCellValue(UserRepoSheetName, successPercentCell, "0.00%")

			} else {
				counts, err := getCounts(output)
				if err != nil {
					f.SetCellValue(UserRepoSheetName, errorCell, err.Error())
					f.SetCellValue(UserRepoSheetName, ranCell, 0)
					f.SetCellValue(UserRepoSheetName, failedCell, 0)
					f.SetCellValue(UserRepoSheetName, testErrorCell, 0)
					f.SetCellValue(UserRepoSheetName, skippedCell, 0)
					f.SetCellValue(UserRepoSheetName, successPercentCell, "0.00%")
				} else {
					percentage := getSuccessPercentage(counts)
					f.SetCellValue(UserRepoSheetName, ranCell, counts[0])
					f.SetCellValue(UserRepoSheetName, failedCell, counts[1])
					f.SetCellValue(UserRepoSheetName, testErrorCell, counts[2])
					f.SetCellValue(UserRepoSheetName, skippedCell, counts[3])
					f.SetCellValue(UserRepoSheetName, successPercentCell, percentage)
				}
			}
		}
	}
	filePathParts := strings.Split(filepath, ".")
	newFilePath := filePathParts[0] + "_" + strconv.FormatInt(time.Now().Unix(), 10) + "." + filePathParts[1]
	if err := f.SaveAs(newFilePath); err != nil {
		println(err.Error())
	}
}

func getSuccessPercentage(counts []int) string {
	var per float32
	if counts[0] == 0 {
		per = 0
	} else {
		per = float32(counts[1]+counts[2]+counts[3]) / float32(counts[0])
		per = per * 100
		per = 100 - per
	}

	return fmt.Sprintf("%.2f%%", per)
}

func processRow(userRepoRow []string, testCaseRepoRow []string) (string, error) {
	userRepoUrl := userRepoRow[0]
	split := strings.Split(userRepoUrl, "/")
	userRepo := split[len(split)-1]
	name := split[len(split)-2]
	testCaseRepoUrl := testCaseRepoRow[0]
	tsplit := strings.Split(testCaseRepoUrl, "/")
	testCaseRepo := tsplit[len(tsplit)-1]
	testCaseOutput, err := runTestCase(name, userRepo, userRepoUrl, testCaseRepo)
	if err != nil {
		log.Println(fmt.Sprintf("error while executing testcase for user {%s}", name), err)
		return "", err
	}
	log.Println(testCaseOutput)
	return testCaseOutput, nil
}

func runTestCase(name string, userRepo string, userRepoUrl string, testRepo string) (string, error) {
	log.Printf("Going to execute test cases for username : {%s}, userRepo : {%s}, userRepoUrl : {%s}, testRepo : {%s}", name, userRepo, userRepoUrl, testRepo)
	out, err := exec.Command("/bin/sh", "resources/run_test_cases_report.sh", name, userRepo, userRepoUrl, testRepo).Output()
	if err != nil {
		log.Println(fmt.Sprintf("error while executing command for user {%s}", name), err)
		return "", err
	}
	log.Printf("Ran test cases successfully for username {%s}", name)
	//enable this comment to print output of shell script
	return string(out), nil
}

func getCounts(s string) ([]int, error) {
	splits := strings.Split(s, "\n")
	var slice = make([]int, 4)
	for i := 0; i < len(splits); i++ {
		counts, err := getNumbersFromATestFile(splits[i])
		if err != nil {
			return slice, errors.New(fmt.Sprintf("Error getting counts {%s}", splits[i]))
		}
		slice[0] = slice[0] + counts[0]
		slice[1] = slice[1] + counts[1]
		slice[2] = slice[2] + counts[2]
		slice[3] = slice[3] + counts[3]
	}
	return slice, nil
}

func getNumbersFromATestFile(output string) ([]int, error) {
	var slice = make([]int, 4)
	if output == "" {
		return slice, nil
	}
	splitByCommaSlice := strings.Split(output, ", ")
	if len(splitByCommaSlice) != 5 {
		return nil, errors.New(fmt.Sprintf("Error parsing output {%s}", output))
	}
	for i := 0; i <= 3; i++ {
		count, err := splitByColon(splitByCommaSlice[i])
		if err != nil {
			return slice, errors.New(fmt.Sprintf("Error getting count {%s}", splitByCommaSlice[i]))
		}
		slice[i] = count
	}
	return slice, nil
}

func splitByColon(s string) (int, error) {
	split := strings.Split(s, ": ")
	if len(split) != 2 {
		return 0, errors.New(fmt.Sprintf("Error getting count {%s}", s))
	}
	i, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Error getting count {%s}", s))
	}
	return i, nil
}
