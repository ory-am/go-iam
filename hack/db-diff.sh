#!/bin/bash

set -o nounset
set -o errexit
set -o pipefail

# Usage
# ./hack/db-diff.sh sqlite master 649f56cc
if [ "$#" -ne 3 ]; then
	echo "Usage: $0 <sqlite|postgres|cockroach|mysql> <commit-ish> <commit-ish>"
	exit 1
fi

# Exports:
# - DB_DIALECT
# - DB_USER
# - DB_PASSWORD
# - DB_HOST
# - DB_PORT
# - DB_NAME
function hydra::util::parse-connection-url {
	local -r url=$1
	if [[ "${url}" =~ ^(.*)://([^:]*):?(.*)@\(?(.*):([0-9]*)\)?/([^?]*) ]]; then
		export DB_DIALECT="${BASH_REMATCH[1]}"
		export DB_USER="${BASH_REMATCH[2]}"
		export DB_PASSWORD="${BASH_REMATCH[3]}"
		export DB_HOST="${BASH_REMATCH[4]}"
		export DB_PORT="${BASH_REMATCH[5]}"
		export DB_DB="${BASH_REMATCH[6]}"
	else
		echo "Failed to parse URL"
		exit 1
	fi
}

function hydra::util::ensure-sqlite {
	if ! sqlite3 --version > /dev/null 2>&1; then
		echo 'Error: sqlite3 is not installed;' >&2
		exit 1
	fi
}

function hydra::util::ensure-pg_dump {
	if ! pg_dump --version > /dev/null 2>&1; then
		echo 'Error: pg_dump is not installed;' >&2
		exit 1
	fi
}

function hydra::util::ensure-mysqldump {
	if ! mysqldump --version > /dev/null 2>&1; then
		echo 'Error: mysqldump is not installed;' >&2
		exit 1
	fi
}

function dump_pg {
	hydra::util::ensure-pg_dump

	make test-resetdb >/dev/null 2>&1
	sleep 2
	yes | go run . migrate sql "$TEST_DATABASE_POSTGRESQL" > /dev/null || true
	pg_dump -s "$TEST_DATABASE_POSTGRESQL" | sed '/^--/d'
}

function dump_cockroach {
	make test-resetdb >/dev/null 2>&1
	sleep 4
	yes | go run . migrate sql "$TEST_DATABASE_COCKROACHDB" > /dev/null || true
	hydra::util::parse-connection-url "${TEST_DATABASE_COCKROACHDB}"
	docker run --rm --net=host -it cockroachdb/cockroach:v20.2.6 dump --dump-all --dump-mode=schema --insecure --user="${DB_USER}" --host="${DB_HOST}" --port="${DB_PORT}"
}

function dump_sqlite {
	hydra::util::ensure-sqlite

	rm "$SQLITE_PATH" > /dev/null 2>&1 || true
	yes | go run -tags sqlite . migrate sql "sqlite://$SQLITE_PATH?_fk=true" > /dev/null 2>&1 || true
	echo '.dump' | sqlite3 "$SQLITE_PATH"
}

function dump_mysql {
	hydra::util::ensure-mysqldump
	make test-resetdb >/dev/null 2>&1
	sleep 10
	yes | go run . migrate sql "$TEST_DATABASE_MYSQL" > /dev/null || true
	hydra::util::parse-connection-url "${TEST_DATABASE_MYSQL}"
	mysqldump --user="$DB_USER" --password="$DB_PASSWORD" --host="$DB_HOST" --port="$DB_PORT" "$DB_DB" --no-data
}

if ! git diff-index --quiet HEAD --; then
	echo 'Error: working tree is dirty;' >&2
	exit 1
fi

case $1 in
	postgres)
		DUMP_CMD=dump_pg
		;;
	cockroach)
		DUMP_CMD=dump_cockroach
		;;
	sqlite)
		DUMP_CMD=dump_sqlite
		;;
	mysql)
		DUMP_CMD=dump_mysql
		;;
	*)
		echo 'Error: unknown database type' >&2
		exit 1
		;;
esac

DIALECT=$1
COMMIT_FROM=$2
COMMIT_TO=$3
T_FROM="./output/sql/$COMMIT_FROM.$DIALECT.dump.sql"
T_TO="./output/sql/$COMMIT_TO.$DIALECT.dump.sql"

mkdir -p ./output/sql/

set -x
git checkout $COMMIT_FROM >/dev/null 2>&1
# set +x
$DUMP_CMD > $T_FROM

git checkout $COMMIT_TO >/dev/null 2>&1
$DUMP_CMD > $T_TO

git checkout @{-2} >/dev/null 2>&1
set +x

echo '---'
echo 'Use the following command to print the diff:'
echo 'git diff --no-index '"$T_FROM"' '"$T_TO"
