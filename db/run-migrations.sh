#!/bin/sh
# run-migrations.sh

set -e
set -x

while [ true ]; do
    echo "waiting for host to accept connections...."
    pg_isready --host=$PGHOST --port=5432 --timeout=5 --username=$PGUSER && break
    sleep 1
done

echo "postgres host is ready to accept connections...."

# check of PGDATABASE already exists and generate it if not...
if [ "$( psql -tA -h $PGHOST -p 5432 -U $PGUSER -w -c "SELECT 1 FROM pg_database WHERE datname='${PGDATABASE}'" )" = '1' ]; then
    echo "${PGDATABASE} database already exists and is ready"
else
    echo "creating $PGDATABASE for you..." # must connect to postgres without db here b/c $PGDATABASE does not exist
    psql -a -h $PGHOST -p 5432 -U $PGUSER -d postgres -w <<-EOSQL
    \set ON_ERROR_STOP on
    CREATE DATABASE "${PGDATABASE}" OWNER postgres;
    ALTER USER ${PGUSER} WITH superuser;
    GRANT ALL PRIVILEGES ON DATABASE "${PGDATABASE}" TO "$PGUSER";
    \set ON_ERROR_STOP off
EOSQL
fi

echo "postgres is up and ${PGDATABASE} database is ready"

# now run migration scripts....
echo "*** executing migrations"
/bin/goose -v -dir /migrations postgres "user=$PGUSER password=$PGPASSWORD dbname=$PGDATABASE sslmode=$PGSSLMODE host=$PGHOST" up
