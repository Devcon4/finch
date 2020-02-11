#!/bin/sh

echo "entrypoint!"

./healthcheck.sh

./wait-for-database.sh $PACT_BROKER_DATABASE_NAME

bundle exec puma --port $PACT_BROKER_PORT
