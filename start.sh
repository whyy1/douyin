#!/bin/sh
set -e

# will run some initial script: eg: db migration
echo "start the app"
exec "$@"