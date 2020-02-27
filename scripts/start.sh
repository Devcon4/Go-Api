#!/bin/sh

./wait-for-database.sh  $DB_HOST+':'+$DB_PORT $DB_PASSWORD $DB_DBNAME

exec /go/bin/server
