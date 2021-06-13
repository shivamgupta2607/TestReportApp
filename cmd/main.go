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
	sh.ForEachRow(processRow)
}

func processRow(r *xlsx.Row) error {
	name, err := r.GetCell(0).FormattedValue()
	if err != nil {
		log.Println("error while getting name", err)
		return err
	}
	userRepo, err := r.GetCell(1).FormattedValue()
	if err != nil {
		log.Println("error while getting userRepo", err)
		return err
	}
	userRepoUrl, err := r.GetCell(2).FormattedValue()
	if err != nil {
		log.Println("error while getting userRepoUrl", err)
		return err
	}
	testCaseRepo, err := r.GetCell(3).FormattedValue()
	if err != nil {
		log.Println("error while getting testCaseRepo", err)
		return err
	}
	testCaseRepoUrl, err := r.GetCell(4).FormattedValue()
	if err != nil {
		log.Println("error while getting testCaseRepoUrl", err)
		return err
	}
	runTestCase(name, userRepo, userRepoUrl, testCaseRepo, testCaseRepoUrl)
	fmt.Println(r)
	return nil
}

func runTestCase(name string, userRepo string, userRepoUrl string, testRepo string, testRepoUrl string) {
	log.Printf("Going to execute test cases for username : {%s}, userRepo : {%s}, userRepoUrl : {%s}, testRepo : {%s}, testRepoUrl {%s}", name, userRepo, userRepoUrl, testRepo, testRepoUrl)
	_, err := exec.Command("/bin/sh", "resources/test_report.sh", name, userRepo, userRepoUrl, testRepo, testRepoUrl).Output()
	if err != nil {
		log.Println(fmt.Sprintf("error while executing command for user {%s}", name), err)
		return
	}
	log.Printf("Ran test cases successfully for username {%s}", name)
	//enable this comment to print output of shell script
	//log.Println(string(out))
}
