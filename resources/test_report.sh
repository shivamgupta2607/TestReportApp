user_name=$1
user_repo=$2
user_repo_url=$3
test_case_repo=$4
cd resources/reports
if [ ! -d $user_name ] ; then
    echo "code is not checked out"
else
    cd $user_name/$user_repo
    rm -rf src/test
    cp -R ../../../0deskera/$test_case_repo/src/test src/
    mvn clean test > /dev/null 2>&1;
    if [ ! -d target/surefire-reports ] ; then
        echo "code compilation failed"
    else
        fileNames=`ls target/surefire-reports/*.txt`
        for eachfile in $fileNames
        do
            sed -n '4p' < $eachfile
        done
    fi
fi