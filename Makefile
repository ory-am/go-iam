SHELL=/bin/bash -o pipefail

.PHONY: tools
tools:
		npm i
		go install github.com/ory/go-acc github.com/ory/x/tools/listx github.com/go-swagger/go-swagger/cmd/swagger github.com/go-bindata/go-bindata/go-bindata github.com/sqs/goreturns github.com/ory/sdk/swagutil

# Runs full test suite including tests where databases are enabled
.PHONY: test-legacy-migrations
test-legacy-migrations:
		make test-resetdb
		make sqlbin
		source scripts/test-env.sh && go test -tags legacy_migration_test -failfast -timeout=20m ./internal/fizzmigrate
		docker rm -f hydra_test_database_mysql
		docker rm -f hydra_test_database_postgres
		docker rm -f hydra_test_database_cockroach

# Runs full test suite including tests where databases are enabled
.PHONY: test
test:
		make test-resetdb
		source scripts/test-env.sh && $$(go env GOPATH)/bin/go-acc ./... -- -failfast -timeout=20m
		docker rm -f hydra_test_database_mysql
		docker rm -f hydra_test_database_postgres
		docker rm -f hydra_test_database_cockroach

# Resets the test databases
.PHONY: test-resetdb
test-resetdb:
		docker kill hydra_test_database_mysql || true
		docker kill hydra_test_database_postgres || true
		docker kill hydra_test_database_cockroach || true
		docker rm -f hydra_test_database_mysql || true
		docker rm -f hydra_test_database_postgres || true
		docker rm -f hydra_test_database_cockroach || true
		docker run --rm --name hydra_test_database_mysql -p 3444:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql:5.7
		docker run --rm --name hydra_test_database_postgres -p 3445:5432 -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=postgres -d postgres:9.6
		docker run --rm --name hydra_test_database_cockroach -p 3446:26257 -d cockroachdb/cockroach:v2.1.6 start --insecure

# Runs tests in short mode, without database adapters
.PHONY: docker
docker:
		packr2
		CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64 go build
		packr2 clean
		docker build -t oryd/hydra:latest .
		docker build -f Dockerfile-alpine -t oryd/hydra:latest-alpine .
		rm hydra

.PHONY: e2e
e2e:
		make test-resetdb
		source ./scripts/test-env.sh
		./test/e2e/circle-ci.bash memory
		./test/e2e/circle-ci.bash memory-jwt
		./test/e2e/circle-ci.bash postgres
		./test/e2e/circle-ci.bash postgres-jwt
		./test/e2e/circle-ci.bash mysql
		./test/e2e/circle-ci.bash mysql-jwt
		./test/e2e/circle-ci.bash cockroach
		./test/e2e/circle-ci.bash cockroach-jwt
		./test/e2e/circle-ci.bash plugin
		./test/e2e/circle-ci.bash plugin-jwt

# Runs tests in short mode, without database adapters
.PHONY: quicktest
quicktest:
		go test -failfast -short ./...

# Formats the code
.PHONY: format
format:
		$$(go env GOPATH)/bin/goreturns -w -local github.com/ory $$($$(go env GOPATH)/bin/listx .)
		npm run format

# Generates mocks
.PHONY: mocks
mocks:
		mockgen -package oauth2_test -destination oauth2/oauth2_provider_mock_test.go github.com/ory/fosite OAuth2Provider

# Adds sql files to the binary using go-bindata
.PHONY: sqlbin
sqlbin:
		cd internal/fizzmigrate/client; $$(go env GOPATH)/bin/go-bindata -o sql_migration_files.go -pkg client ./migrations/sql/...
		cd internal/fizzmigrate/consent; $$(go env GOPATH)/bin/go-bindata -o sql_migration_files.go -pkg consent ./migrations/sql/...
		cd internal/fizzmigrate/jwk; $$(go env GOPATH)/bin/go-bindata -o sql_migration_files.go -pkg jwk ./migrations/sql/...
		cd internal/fizzmigrate/oauth2; $$(go env GOPATH)/bin/go-bindata -o sql_migration_files.go -pkg oauth2 ./migrations/sql/...

# Runs all code generators
.PHONY: gen
gen: mocks sqlbin sdk

# Generates the SDKs
.PHONY: sdk
sdk:
		$$(go env GOPATH)/bin/swagger generate spec -m -o ./docs/api.swagger.json -x internal/httpclient,gopkg.in/square/go-jose.v2
		$$(go env GOPATH)/bin/swagutil sanitize ./docs/api.swagger.json
		$$(go env GOPATH)/bin/swagger flatten --with-flatten=remove-unused -o ./docs/api.swagger.json ./docs/api.swagger.json
		$$(go env GOPATH)/bin/swagger validate ./docs/api.swagger.json
		rm -rf internal/httpclient
		mkdir -p internal/httpclient
		$$(go env GOPATH)/bin/swagger generate client -f ./docs/api.swagger.json -t internal/httpclient -A Ory_Hydra
		make format


.PHONY: install-stable
install-stable:
		HYDRA_LATEST=$$(git describe --abbrev=0 --tags)
		git checkout $$HYDRA_LATEST
		packr2
		GO111MODULE=on go install \
				-ldflags "-X github.com/ory/hydra/cmd.Version=$$HYDRA_LATEST -X github.com/ory/hydra/cmd.Date=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` -X github.com/ory/hydra/cmd.Commit=`git rev-parse HEAD`" \
				.
		packr2 clean
		git checkout master

.PHONY: install
install:
		packr2
		GO111MODULE=on go install .
		packr2 clean

.PHONY: init
init:
		GO111MODULE=on go get .
		GO111MODULE=on go install github.com/ory/go-acc
