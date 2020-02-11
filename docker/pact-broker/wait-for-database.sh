#!/bin/sh
# wait-for-database.sh

set -e

host="$1"
shift
pass="$2"
shift
database="$3"
shift
cmd="$@"

# psql -h "$host" -U "postgres" -c '\q'
until PGPASSWORD="$pass" psql -h "$host" -lqt | cut -d \| -f 1 | grep -qw "$database"; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec cmd