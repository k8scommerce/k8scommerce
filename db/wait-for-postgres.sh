#!/bin/sh
# wait-for-postgres.sh
# modification of https://gist.github.com/nicerobot/1136dcfba6ce3da67ce3ded5101a4078

set -e
set -x

cmd="$@"

until psql -h "$PGHOST" -U "$PGUSER" -d postgres -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd