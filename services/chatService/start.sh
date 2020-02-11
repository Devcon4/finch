#!/bin/sh

./wait-for-database.sh $DB_DBNAME

exec /go/bin/server
