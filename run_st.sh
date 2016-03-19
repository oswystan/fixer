
#!/bin/bash

################################
# some basic functions for log
################################
log_start()
{
    strNow=`date +'%Y-%m-%d %H:%M:%S'`
    echo "[${strNow}]##########################################################"
    echo "[${strNow}] start program  : $0"
    echo "[${strNow}]##########################################################"
    echo ""
}

logi()
{
    strNow=`date +'%Y-%m-%d %H:%M:%S'`
    echo "[${strNow}] INFO:$*"
}

logw()
{
    strNow=`date +'%Y-%m-%d %H:%M:%S'`
    echo "[${strNow}] WARN:$*"
}

loge()
{
    strNow=`date +'%Y-%m-%d %H:%M:%S'`
    echo "[${strNow}]ERROR:$*"
}

log_end()
{
    strNow=`date +'%Y-%m-%d %H:%M:%S'`
    echo ""
    echo "[${strNow}]##########################################################"
    echo "[${strNow}] finished $0"
    echo "[${strNow}]##########################################################"
}
safe_exec()
{
    if [ $# -eq 0 ]; then
        exit 1
    fi

    $*
    if [ $? -ne 0 ]; then
        loge "fail to do [$*]"
        exit 1
    fi
}

do_bug()
{
    safe_exec curl -XGET "http://localhost:8000/filter/buglist.html?team_id=1&handler=1&created_by=2&status=1&priority=1&date_from=2012-01-01&date_to=2017-01-01&offset=0&count=10" -w %{http_code}
    echo ""
}

do_team()
{
    safe_exec curl -XGET "http://localhost:8000/filter/team.html?id=1" -w "%{http_code}\n"
    safe_exec curl -XGET "http://localhost:8000/filter/team.html?name=john-frog" -w "%{http_code}\n"
    safe_exec curl -XGET "http://localhost:8000/filter/teamlist.html?creatorid=1" -w "%{http_code}\n"
    safe_exec curl -XGET "http://localhost:8000/filter/teamlist.html?creator=john" -w "%{http_code}\n"
    safe_exec curl -XGET "http://localhost:8000/filter/teamlist.html?memberid=1"  -w "%{http_code}\n"
    safe_exec curl -XGET "http://localhost:8000/filter/teamlist.html?member=john" -w "%{http_code}\n"
}
do_member()
{
    safe_exec curl -XGET "http://localhost:8000/filter/memberlist.html?team_id=1" -w %{http_code}
    echo ""
}
do_user()
{
    safe_exec curl -XGET "http://localhost:8000/filter/user-detail.html?id=1" -w "%{http_code}\n"
    safe_exec curl -XGET "http://localhost:8000/filter/user-detail.html?nicky=sherlock" -w "%{http_code}\n"
}

do_api()
{
    safe_exec curl -XGET    "http://localhost:8000/users" -w "%{http_code}\n"
    safe_exec curl -XPOST   "http://localhost:8000/users" -w "%{http_code}\n"
    safe_exec curl -XDELETE "http://localhost:8000/users" -w "%{http_code}\n"
    safe_exec curl -XGET    "http://localhost:8000/users/1" -w "%{http_code}\n"
    safe_exec curl -XPUT    "http://localhost:8000/users/1" -w "%{http_code}\n"
    safe_exec curl -XDELETE "http://localhost:8000/users/1" -w "%{http_code}\n"
    safe_exec curl -XGET    "http://localhost:8000/users/1/teams/joined"  -w "%{http_code}\n"
    safe_exec curl -XGET    "http://localhost:8000/users/1/teams/created" -w "%{http_code}\n"
}

do_users()
{
    logi "test users releated features..."
    safe_exec curl -XGET    "http://localhost:8000/users/1/teams/joined"    -w "%{http_code}\n"
    safe_exec curl -XGET    "http://localhost:8000/users/1/teams/created"   -w "%{http_code}\n"
    safe_exec curl -XGET    "http://localhost:8000/users?q=j"               -w "%{http_code}\n"
    safe_exec curl -XGET    "http://localhost:8000/users/2"                 -w "%{http_code}\n"
    safe_exec curl -XDELETE "http://localhost:8000/users/1"                 -w "%{http_code}\n"
    safe_exec curl -XDELETE "http://localhost:8000/users"                   -w "%{http_code}\n"
    safe_exec curl -XGET    "http://localhost:8000/users"                   -w "%{http_code}\n"

    data_post='{
        "nicky":"mark", 
        "pwd":"123456a", 
        "portrait":"static/images/1.jpg", 
        "email":"mark@gmail.com"
    }'
    data_put='{
        "nicky":"mark", 
        "pwd":"123456a", 
        "portrait":"static/images/1.jpg", 
        "email":"marks@gmail.com"
    }'
    curl -XPOST -d "$data_post" "http://localhost:8000/users"               -w "%{http_code}\n"
    safe_exec curl -XGET    "http://localhost:8000/users"                   -w "%{http_code}\n"
    curl -XPUT -d "$data_put"   "http://localhost:8000/users/3"             -w "%{http_code}\n"
    logi "done."
}

do_teams()
{
    logi "test teams releated features..."
    safe_exec curl -XGET "http://localhost:8000/teams/1"                    -w "%{http_code}\n"
    safe_exec curl -XGET "http://localhost:8000/teams/1/users"              -w "%{http_code}\n"
    safe_exec curl -XGET "http://localhost:8000/teams?q=j&offset=1&limit=2" -w "%{http_code}\n"

    data_post='{
        "name":"sherlock-fox", 
        "leader_id":2, 
        "goal":"make team like a fox", 
        "logo":"static/images/1.jpg"
    }'
    data_put='{
        "name":"sherlock-jaguar", 
        "goal":"make team like a jar", 
        "logo":"static/images/1.jpg"
    }'
    curl -XPOST -d "$data_post" "http://localhost:8000/teams"               -w "%{http_code}\n"
    curl -XPUT -d "$data_put" "http://localhost:8000/teams/6"               -w "%{http_code}\n"

    curl -XPUT "http://localhost:8000/teams/6/users/1"              -w "%{http_code}\n"
    safe_exec curl -XDELETE "http://localhost:8000/teams/5/users"   -w "%{http_code}\n"
    safe_exec curl -XDELETE "http://localhost:8000/teams/6"         -w "%{http_code}\n"
    safe_exec curl -XGET    "http://localhost:8000/teams"           -w "%{http_code}\n"
    safe_exec curl -XDELETE "http://localhost:8000/teams"           -w "%{http_code}\n"
}

do_bugs()
{
    url="http://localhost:8000/teams/1/bugs?offset=0&limit=9&handler=1&date_from=20150101000001&date_to=20170101000001" 
    safe_exec curl -XGET $url -w "%{http_code}\n"
    safe_exec curl -XGET "http://localhost:8000/teams/1/bugs/1" -w "%{http_code}\n"
    safe_exec curl -XDELETE "http://localhost:8000/teams/1/bugs/1" -w "%{http_code}\n"

    data_post='{
        "created_by":1, 
        "current_handler":1, 
        "priority":2, 
        "status":1, 
        "title":"system crashed", 
        "attachments":"", 
        "detail":"system core dump"
    }'
    data_put='{
        "created_by":1, 
        "current_handler":1, 
        "priority":1, 
        "status":1, 
        "title":"system crashed", 
        "attachments":"", 
        "detail":"system core dump"
    }'
    curl -XPOST -d "$data_post" "http://localhost:8000/teams/1/bugs"    -w "%{http_code}\n"
    curl -XPUT  -d "$data_put"  "http://localhost:8000/teams/1/bugs/2"  -w "%{http_code}\n"
}

do_buglogs()
{
    safe_exec curl -XGET "http://localhost:8000/teams/1/bugs/1/logs?offset=0&limit=9" -w "%{http_code}\n"
    data_post='{
        "bug_id":1, 
        "who":1, 
        "action_type":1, 
        "action":"raised the bug"
    }'
    curl -XPOST -d "${data_post}" "http://localhost:8000/teams/1/bugs/1/logs" -w "%{http_code}\n"
    safe_exec curl -XDELETE "http://localhost:8000/teams/1/bugs/1/logs" -w "%{http_code}\n"
}

do_work()
{
    log_start

    #do_bug
    #do_team
    #do_member
    #do_user
    #do_api
    #do_users
    #do_teams
    do_bugs
    #do_buglogs

    log_end
}

################################
## main
################################

if [ $# -eq 0 ]; then
    do_work
else
    $*
fi
