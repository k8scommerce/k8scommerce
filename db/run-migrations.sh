#!/bin/sh
# run-migrations.sh

set -e
set -x

while [ true ]; do
    echo "waiting for host to accept connections"
    pg_isready --host=$PG_HOST --port=5432 --timeout=5 --username=$PG_USER && break
    sleep 1
done

echo "> postgres host is ready to accept connections"

if [ "$( psql -tA -h $PG_HOST -p 5432 -U $PG_USER -w -c "SELECT 1 FROM pg_database WHERE datname='${PG_DATABASE}'" )" = '1' ]; then
    echo "> ${PG_DATABASE} database already exists and is ready"
else
    echo "> creating $PG_DATABASE for you..."
    psql -a -h $PG_HOST -p 5432 -U $PG_USER -d postgres -w <<-EOSQL
    \set ON_ERROR_STOP on
    CREATE DATABASE "${PG_DATABASE}" OWNER postgres;
    ALTER USER ${PG_USER} WITH superuser;
    GRANT ALL PRIVILEGES ON DATABASE "${PG_DATABASE}" TO "$PG_USER";
    \set ON_ERROR_STOP off
EOSQL
fi

echo "> postgres is up and ${PG_DATABASE} database is ready"

# run migration scripts
echo "> executing migrations"
/bin/goose -v -dir /migrations postgres "user=$PG_USER password=$PG_PASSWORD dbname=$PG_DATABASE sslmode=$PG_SSL_MODE host=$PG_HOST" up
