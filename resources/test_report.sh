user_name=$1
user_repo=$2
user_repo_url=$3
test_case_repo=$4
test_case_repo_url=$5
cd resources/reports
echo "creating resource dir for user :" $user_name
mkdir $user_name
cd $user_name
echo "cloning user project"
git clone $user_repo_url
echo "cloning test case project"
git clone $test_case_repo_url
echo "both projects has been cloned"
cd $user_repo
echo "going to delete test cases from userapp"
rm -rf src/test/java
echo "going to copy test resources in user repo"
cp -R ../$test_case_repo/src/test src/
echo "Copy successful"
echo "going to run test cases for app"
testcaseCommand=`mvn clean test`
echo $testcaseCommand