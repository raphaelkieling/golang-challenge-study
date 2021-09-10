#!/bin/sh

set -e
shift
  
until PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -d "$DB_NAME" -U "postgres" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done
  
>&2 echo "Postgres is up - executing command"
exec "$@"