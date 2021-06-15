package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Welcome to Test Report app")
	filepath := "resources/SampleTestFile.xlsx"
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
	testCaseRepo := testCaseRepoRow[0]
	testCaseRepoUrl := testCaseRepoRow[1]
	checkoutTestCaseRepo(testCaseRepo, testCaseRepoUrl)
	for i, row := range userRepoRows {
		if i == 0 {
			continue
		}
		processRow(row, testCaseRepoRow)
	}
}

func processRow(userRepoRow []string, testCaseRepoRow []string) {
	name := userRepoRow[0]
	userRepo := userRepoRow[1]
	userRepoUrl := userRepoRow[2]
	checkoutUserRepo(name, userRepo, userRepoUrl)
}

func checkoutTestCaseRepo(testRepo string, testRepoUrl string) {
	log.Printf("Going to checkout test caserepo for testRepo : {%s}, testRepoUrl {%s}", testRepo, testRepoUrl)
	testRepoCheckoutOutput, err := exec.Command("/bin/sh", "resources/checkout_testcase_repo.sh", testRepo, testRepoUrl).Output()
	if err != nil {
		log.Println(fmt.Sprintf("error while checking out test case repo."), err)
		return
	}
	log.Printf("TestCase Code checkout successfully.")
	log.Println(testRepoCheckoutOutput)
}

func checkoutUserRepo(name string, userRepo string, userRepoUrl string) {
	log.Printf("Going to checkout user repo for username : {%s}, userRepo : {%s}, userRepoUrl : {%s}", name, userRepo, userRepoUrl)
	userRepoCheckoutOutput, err := exec.Command("/bin/sh", "resources/checkout_user_repos.sh", name, userRepo, userRepoUrl).Output()
	if err != nil {
		log.Println(fmt.Sprintf("error while checking out user repo {%s}", name), err)
		return
	}
	log.Printf("Code checkout successfully for username {%s}", name)
	log.Println(userRepoCheckoutOutput)
}
