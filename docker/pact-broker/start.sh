#!/bin/sh

./wait-for-database.sh $PACT_BROKER_DATABASE_HOST $PACT_BROKER_DATABASE_PASSWORD $PACT_BROKER_DATABASE_NAME

./entrypoint.sh

# bundle exec puma --port $PACT_BROKER_PORT
