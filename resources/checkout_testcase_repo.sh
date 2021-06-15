test_case_repo=$1
test_case_repo_url=$2
cd resources/testcaserepo
if [ ! -d $test_case_repo ] ; then
    git clone $test_case_repo_url
fi