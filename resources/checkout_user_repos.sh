user_name=$1
user_repo=$2
user_repo_url=$3
cd resources/reports
if [ ! -d $user_name ] ; then
    mkdir $user_name
    cd $user_name
    git clone $user_repo_url
fi