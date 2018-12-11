#
# Copyright 2018 Yaacov Zamir <kobi.zamir@gmail.com>
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

all: fmt tsl_parser tsl_sqlite tsl_gorm tsl_mongo tsl_graphql

tsl_parser:
	go build ./cmd/tsl_parser

tsl_sqlite:
	go build ./cmd/tsl_sqlite

tsl_gorm:
	go build ./cmd/tsl_gorm

tsl_mongo:
	go build ./cmd/tsl_mongo

tsl_graphql:
	go build ./cmd/tsl_graphql

.PHONY: lint
lint:
	golangci-lint \
		run \
		--skip-dirs=/pkg/parser \
		--no-config \
		--issues-exit-code=1 \
		--deadline=15m \
		--disable-all \
		--enable=deadcode \
		--enable=gas \
		--enable=goconst \
		--enable=gocyclo \
		--enable=gofmt \
		--enable=golint \
		--enable=ineffassign \
		--enable=interfacer \
		--enable=lll \
		--enable=maligned \
		--enable=megacheck \
		--enable=misspell \
		--enable=structcheck \
		--enable=unconvert \
		--enable=varcheck \
		$(NULL)

.PHONY: fmt
fmt:
	gofmt -s -l -w ./pkg/ ./cmd/

.PHONY: clean
clean:
	rm -rf vendor
	rm tsl_parser
	rm tsl_gorm
	rm tsl_mongo
	rm tsl_sqlite
	rm tsl_graphql

.PHONY: test
test:
	go test ./cmd/tsl_parser
	go test ./cmd/tsl_sqlite
	go test ./cmd/tsl_gorm
	go test ./cmd/tsl_mongo
	go test ./cmd/tsl_graphql
	go test ./pkg/tsl
	go test ./pkg/walkers/sql
	go test ./pkg/walkers/mongo
	go test ./pkg/walkers/graphviz

.PHONY: generate
generate:
	antlr4 -Dlanguage=Go -o pkg/parser TSL.g4
