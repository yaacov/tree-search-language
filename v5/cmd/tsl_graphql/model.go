// Copyright 2018 Yaacov Zamir <kobi.zamir@gmail.com>
// and other contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/yaacov/tree-search-language/v5/cmd/model"
)

const sqlStmt = `
create table if not exists books (
	id integer not null primary key,
	title text,
	author text,
	pages integer,
	rating integer
);
delete from books;
`

func connect(ctx context.Context, url string) (tx *sql.Tx, err error) {
	var db *sql.DB

	db, err = sql.Open("sqlite3", url)
	check(err)

	tx, err = db.BeginTx(ctx, nil)

	return
}

func prepareCollection(ctx context.Context, tx *sql.Tx) (err error) {
	// Create table.
	_, err = tx.ExecContext(ctx, sqlStmt)
	check(err)

	// Insert new books into the table.
	stmt, err := tx.PrepareContext(ctx, "insert into books(title, author, pages, rating) values(?, ?, ?, ?)")
	check(err)

	defer stmt.Close()

	for _, b := range model.Books {
		_, err = stmt.ExecContext(ctx,
			b.(model.Book).Title,
			b.(model.Book).Author,
			b.(model.Book).Spec.Pages,
			b.(model.Book).Spec.Rating)

		check(err)
	}

	return
}
