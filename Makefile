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

tsl_parser_src := $(wildcard ./cmd/tsl_parser/*.go)
tsl_sqlite_src := $(wildcard ./cmd/tsl_sqlite/*.go)
tsl_gorm_src := $(wildcard ./cmd/tsl_gorm/*.go)
tsl_mongo_src := $(wildcard ./cmd/tsl_mongo/*.go)
tsl_graphql_src := $(wildcard ./cmd/tsl_graphql/*.go)
tsl_mem_src := $(wildcard ./cmd/tsl_mem/*.go)

all: fmt tsl_parser tsl_sqlite tsl_gorm tsl_mongo tsl_graphql tsl_mem

tsl_parser: $(tsl_parser_src)
	go build ./cmd/tsl_parser

tsl_sqlite: $(tsl_sqlite_src)
	go build ./cmd/tsl_sqlite

tsl_gorm: $(tsl_gorm_src)
	go build ./cmd/tsl_gorm

tsl_mongo: $(tsl_mongo_src)
	go build ./cmd/tsl_mongo

tsl_graphql: $(tsl_graphql_src)
	go build ./cmd/tsl_graphql

tsl_mem: $(tsl_mem_src)
	go build ./cmd/tsl_mem

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
	rm tsl_mem

.PHONY: test
test:
	go test ./cmd/tsl_parser
	go test ./cmd/tsl_sqlite
	go test ./cmd/tsl_gorm
	go test ./cmd/tsl_mongo
	go test ./cmd/tsl_graphql
	go test ./cmd/tsl_mem
	go test ./pkg/tsl
	go test ./pkg/walkers/sql
	go test ./pkg/walkers/mongo
	go test ./pkg/walkers/graphviz

.PHONY: generate
generate:
	antlr4 -Dlanguage=Go -o pkg/parser TSL.g4
