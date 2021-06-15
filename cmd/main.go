package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"os/exec"
	"strconv"
)

func main() {
	fmt.Println("Welcome to Test Report app")
	filename := "resources/SampleTestFile.xlsx"
	UserRepoSheetName := "Sheet1"
	TestRepoSheetName := "Sheet2"
	//filename := "/Users/shivam.gupta/projects/go/personal/TestReportApp/resources/SampleTestFile.xlsx"
	f, err := excelize.OpenFile(filename)
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
		outCell := "D" + strconv.Itoa(i+1)
		errorCell := "E" + strconv.Itoa(i+1)
		output, err := processRow(row, tetCaseRepoRow)
		if err != nil {
			f.SetCellValue(UserRepoSheetName, errorCell, err.Error())
		}
		f.SetCellValue(UserRepoSheetName, outCell, output)
	}
	if err := f.SaveAs(filename); err != nil {
		println(err.Error())
	}
}

func processRow(userRepoRow []string, testCaseRepoRow []string) (string, error) {
	name := userRepoRow[0]
	userRepo := userRepoRow[1]
	userRepoUrl := userRepoRow[2]
	testCaseRepo := testCaseRepoRow[0]
	testCaseRepoUrl := testCaseRepoRow[1]

	testCaseOutput, err := runTestCase(name, userRepo, userRepoUrl, testCaseRepo, testCaseRepoUrl)
	if err != nil {
		log.Println(fmt.Sprintf("error while executing testcase for user {%s}", name), err)
		return "", err
	}
	log.Println(testCaseOutput)
	return testCaseOutput, nil
}

func runTestCase(name string, userRepo string, userRepoUrl string, testRepo string, testRepoUrl string) (string, error) {
	log.Printf("Going to execute test cases for username : {%s}, userRepo : {%s}, userRepoUrl : {%s}, testRepo : {%s}, testRepoUrl {%s}", name, userRepo, userRepoUrl, testRepo, testRepoUrl)
	out, err := exec.Command("/bin/sh", "resources/test_report.sh", name, userRepo, userRepoUrl, testRepo, testRepoUrl).Output()
	if err != nil {
		log.Println(fmt.Sprintf("error while executing command for user {%s}", name), err)
		return "", err
	}
	log.Printf("Ran test cases successfully for username {%s}", name)
	//enable this comment to print output of shell script
	return string(out), nil
}
