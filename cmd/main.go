package main

import (
	"fmt"
	xlsx "github.com/tealeg/xlsx/v3"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Welcome to Test Report app")
	filename := "resources/SampleTestFile.xlsx"
	//filename := "/Users/shivam.gupta/projects/go/personal/TestReportApp/resources/SampleTestFile.xlsx"
	processExcelFile(filename)
	fmt.Println("File has been processed successfully")
}

func processExcelFile(filename string) {
	wb, err := xlsx.OpenFile(filename)
	if err != nil {
		panic(err)
	}
	sh := wb.Sheets[0]
	fmt.Println("No of records to process is", sh.MaxRow)
	sheet2 := wb.Sheets[1]
	testRow, err := sheet2.Row(1)
	if err != nil {
		log.Println("Error while getting testRow from test repo sheet ")
		return
	}
	testCaseRepo, err := testRow.GetCell(0).FormattedValue()
	if err != nil {
		log.Println("Error while getting test case repo name from test repo sheet ")
		return
	}

	testCaseRepoUrl, err := testRow.GetCell(1).FormattedValue()
	if err != nil {
		log.Println("Error while getting test case repo url row from test repo sheet ")
		return
	}
	//ignoring header row
	for i := 1; i < sh.MaxRow; i++ {
		userRepoRow, err := sh.Row(i)
		if err != nil {
			log.Println("Error while getting user repo row from test repo sheet ")
			continue
		}
		name, err := userRepoRow.GetCell(0).FormattedValue()
		if err != nil {
			log.Println("error while getting name", err)
			continue
		}
		userRepo, err := userRepoRow.GetCell(1).FormattedValue()
		if err != nil {
			log.Println("error while getting userRepo", err)
			continue
		}
		userRepoUrl, err := userRepoRow.GetCell(2).FormattedValue()
		if err != nil {
			log.Println("error while getting userRepoUrl", err)
			continue
		}
		testCaseOutput, err := runTestCase(name, userRepo, userRepoUrl, testCaseRepo, testCaseRepoUrl)
		if err != nil {
			log.Println(fmt.Sprintf("error while executing testcase for user {%s}", name), err)
			continue
		}
		log.Println(testCaseOutput)
	}
}

func runTestCase(name string, userRepo string, userRepoUrl string, testRepo string, testRepoUrl string) (string, error){
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
