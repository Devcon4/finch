# Postgres-Scripts

The postgres docker image will run any scripts in this folder when it is created. The `create-multiple-postgresql-databases` is a helper script that lets us declare multiple databases to be created on startup (you normally can only create one).

The [official recommendation](https://hub.docker.com/_/postgres/) for creating multiple databases is as follows:

> If you would like to do additional initialization in an image derived from this one, add one or more *.sql, *.sql.gz, or *.sh scripts under /docker-entrypoint-initdb.d (creating the directory if necessary). After the entrypoint calls initdb to create the default postgres user and database, it will run any *.sql files and source any *.sh scripts found in that directory to do further initialization before starting the service.

## scripts

- `create-multiple-postgresql-databases.sh`: source https://github.com/mrts/docker-postgresql-multiple-databases
- `wait-for-database.sh`: docker-compose depend-on only waits until the service is started not until it is healthy. This is a helper script to wait until a database is available before starting.