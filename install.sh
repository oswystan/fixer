#!/bin/bash
log_file="log.txt"

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

log()
{
    strNow=`date +'%Y-%m-%d %H:%M:%S'`
    echo "[${strNow}] INFO: $*"
}

log_end()
{
    strNow=`date +'%Y-%m-%d %H:%M:%S'`
    echo ""
    echo "[${strNow}]##########################################################"
    echo "[${strNow}] finished $0"
    echo "[${strNow}]##########################################################"
}

do_work()
{
    log_start

    ## install the database
    psql -f ./db/pg.sql

    log_end
}

rm_db()
{
    psql -c "drop database if exists fixer;"
    psql -c "drop user if exists pgtest;"
}

################################
## main
################################
if [ $# -eq 0 ]; then
    do_work
else
    $*
fi

