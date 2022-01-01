#!/bin/sh
# wait-for-postgres.sh

set -e
set -x

cmd="$@"

until psql -h "$PGHOST" -U "$PGUSER" -d postgres -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd