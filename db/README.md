# Database Migrations

## Setup

Run `docker-compose up` (or `./db/docker-compose -f docker-compose.yml up`) from `db` folder and it will create the `localrivet` database and run the migration scripts necessary for you to run the full test suite.

Wait for it to spin up postgres and run migrations, then open browser to http://localhost:5431/browser/ , enter user/pass as in `docker-compose.yaml` which are `postgres@foo.com`/`postgres`.

Alternatively, run the `goose` migration tool manually on the cli (described below) to bring it up to date against an already running database.

If the pgadmin webpage (http://localhost:5431/browser/) is already open it may need reload if there is <em>CSRF session token is missing</em> error.

If migration scripts are added, you'll need to either rebuild it like so `docker-compose build && docker-compose run` or `docker-compose up`.

In Kubernetes, we already create test database with `init containers`.

## Migrations

Set the env var `GOOSE_DBSTRING`
```sh
postgres://postgres:postgres@localhost:5432/localrivet?connect_timeout=180&sslmode=disable
```

_NOTE_: You only need to do this if you are just running `postgres` locally.  If you run it in `docker-compose` the migrations are run for you automatically.

To run the migrations, I'm using the [goose](https://github.com/pressly/goose) tool.  You could obviously just run the migration up/down scripts in order on the database, but the tool makes it easier.

First [install the CLI tool](https://github.com/pressly/goose#install)

To create a migration, run as:

```sh
cd path/to/db/migrations

goose create AddSomeField sql
```

*NOTE*: The `sql` arg is very important - we'll be using sql-based migration files, and not golang files, needing executed with a built `goose` binary.

To apply migrations, run as:

```sh
goose postgres "postgres://postgres@localhost:5432/localrivet?user=postgres&password=postgres&connect_timeout=180&sslmode=disable" up
```

To back out, run as (**WARNING**: do not do on production, obviously):

```sh
goose postgres "postgres://postgres@localhost:5432/localrivet?user=postgres&password=postgres&connect_timeout=180&sslmode=disable" down

GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://postgres@localhost:5428/localrivet?user=postgres&password=postgres&connect_timeout=180&sslmode=disable" goose reset
```

You can give the command specific up/down commands so it can run against a particular script and not all of them; consult the docs. For additional information, run

```sh
    goose --help
```

*PROTIP*: Each migration needs to have SQL in the up and the down, else the `goose` tool throws an error. For instance, if you have a simple migration whose only job is to alter a field to make it have a default, you don't really need to _undo_ it, but something needs put in the `down` area - so `SELECT false;` works just fine for that.

### Check Migrations Executed vs Expected

While we can use `goose status` cmd to see the status of migrations on a db, that doesn't easily tell us for sure if any migrations have been missed. Here's a relatively easy way to do this:

1. print out all migration ids (timestamps) via this script (from within the `migrations` dir)
    `ls | grep sed 's/_.*//' | awk '{print ",("$1")"}'`
This will print out a list like:
```
,(123123123)
,(123123124)
...
```
2. Copy those values, less the very first leading comma, into this query at the marker:
```
WITH needed_migration_ids(id) AS (
	values
	-- IDS INSERT HERE
), executed_migration_ids AS (
	SELECT version_id FROM goose_db_version WHERE version_id > 0
)
SELECT
	id
FROM
	needed_migration_ids
WHERE id NOT IN (SELECT id FROM executed_migration_ids);
```
Executing this query against any db should return zero results if all good, but any that are returned are migrations that need ran, that were missed in that environment.

## Dockerfile

The current Makefile supports the following commands:

`make docker` -- will create the image locally.

`make push` -- will build the image and push to ECR.

The image is based on `busybox` so you have other Linux CLI tools if needed, but the `goose` binary is loaded there in `/bin/goose` as well.  Also at root dir is a copy of this `./migrations/` folder.  You can deploy as a k8s `Job` to run goose with whatever params are needed in the cluster.

## Checklist

When adding a migration, it is imperitive that it works - both in the `up` and in the `down`. Before merging a PR with a new migration, make sure to have tested it out wholly. A good way to test it out wholly is to use a test db and do a total `reset` and `up` - **twice**. Twice so that we make sure everything needing dropped was indeed dropped and no problems arise on the second `up`. Here's an example:

*Reset*  
`$ goose postgres "postgres://postgres@localhost:5432/localrivet?connect_timeout=180&sslmode=disable" reset`  
*Up*  
`$ goose postgres "postgres://postgres@localhost:5432/localrivet?connect_timeout=180&sslmode=disable" up`  
*Reset*  
`$ goose postgres "postgres://postgres@localhost:5432/localrivet?connect_timeout=180&sslmode=disable" reset`  
*Up*  
`$ goose postgres "postgres://postgres@localhost:5432/localrivet?connect_timeout=180&sslmode=disable" up`  
