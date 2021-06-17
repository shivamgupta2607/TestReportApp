package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"os/exec"
	"strings"
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
	testCaseRepoRow := testCaseRows[1]
	testCaseRepoUrl := testCaseRepoRow[0]
	split := strings.Split(testCaseRepoUrl, "/")
	testCaseRepo := split[len(split)-1]
	checkoutTestCaseRepo(testCaseRepo, testCaseRepoUrl)
	for i, row := range userRepoRows {
		if i == 0 {
			continue
		}
		processRow(row, testCaseRepoRow)
	}
}

func processRow(userRepoRow []string, testCaseRepoRow []string) {
	userRepoUrl := userRepoRow[0]
	split := strings.Split(userRepoUrl, "/")
	userRepo := split[len(split)-1]
	name := split[len(split)-2]
	checkoutUserRepo(name, userRepo, userRepoUrl)
}

func checkoutTestCaseRepo(testRepo string, testRepoUrl string) {
	log.Printf("Going to checkout test caserepo for testRepo : {%s}, testRepoUrl {%s}", testRepo, testRepoUrl)
	testRepoCheckoutOutput, err := exec.Command("/bin/sh", "resources/checkout_testcase_repo.sh", testRepo, testRepoUrl).Output()
	if err != nil {
		log.Println(fmt.Sprintf("error while checking out test case repo."), err)
		return
	}
	log.Println(string(testRepoCheckoutOutput))
}

func checkoutUserRepo(name string, userRepo string, userRepoUrl string) {
	log.Printf("Going to checkout user repo for username : {%s}, userRepo : {%s}, userRepoUrl : {%s}", name, userRepo, userRepoUrl)
	userRepoCheckoutOutput, err := exec.Command("/bin/sh", "resources/checkout_user_repos.sh", name, userRepo, userRepoUrl).Output()
	if err != nil {
		log.Println(fmt.Sprintf("error while checking out user repo {%s}", name), err)
		return
	}
	log.Println(string(userRepoCheckoutOutput))
}
