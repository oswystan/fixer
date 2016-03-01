
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
    safe_exec curl -XGET    "http://localhost:8000/users/1/teams/joined" -w "%{http_code}\n"
    safe_exec curl -XGET    "http://localhost:8000/users/1/teams/created" -w "%{http_code}\n"
}

do_users()
{
    #safe_exec curl -XGET    "http://localhost:8000/users/1/teams/joined"    -w "%{http_code}\n"
    #safe_exec curl -XGET    "http://localhost:8000/users/1/teams/created"   -w "%{http_code}\n"
    #safe_exec curl -XGET    "http://localhost:8000/teams/1"                 -w "%{http_code}\n"
    #safe_exec curl -XGET    "http://localhost:8000/teams/1/users"           -w "%{http_code}\n"
    safe_exec curl -XGET    "http://localhost:8000/teams?offset=1&limit=2"           -w "%{http_code}\n"
}

do_work()
{
    log_start

    #do_bug
    #do_team
    #do_member
    #do_user
    #do_api
    do_users

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
