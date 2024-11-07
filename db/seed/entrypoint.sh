#!/bin/sh
set -e

testfixtures -d mysql -c "${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?multiStatements=true" -D ./seed --dangerous-no-test-database-check
