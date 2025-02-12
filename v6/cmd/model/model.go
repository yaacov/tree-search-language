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

// Package model for demo.
package model

// BookSpecs for demo.
type BookSpecs struct {
	Pages  uint `bson:"pages,omitempty" json:"pages,omitempty"`
	Rating uint `bson:"rating,omitempty" json:"rating,omitempty"`
}

// Book for demo.
type Book struct {
	Title  string    `bson:"title,omitempty" json:"title,omitempty"`
	Author string    `bson:"author,omitempty" json:"author,omitempty"`
	Spec   BookSpecs `bson:"spec,omitempty" json:"spec,omitempty"`
	OnLoan bool      `bson:"on-loan" json:"on-loan"`
}

// Books are the demo list of books.
var Books = []interface{}{
	Book{Title: "Book", Author: "Joe", Spec: BookSpecs{Pages: 100, Rating: 4}, OnLoan: true},
	Book{Title: "Other Book", Author: "Jane", Spec: BookSpecs{Pages: 200, Rating: 3}, OnLoan: true},
	Book{Title: "Some Book", Author: "Jane", Spec: BookSpecs{Pages: 50, Rating: 5}, OnLoan: true},
	Book{Title: "Some Other Book", Author: "Jane", Spec: BookSpecs{Pages: 50}, OnLoan: false},
	Book{Title: "Good Book", Author: "Joe", Spec: BookSpecs{Pages: 150, Rating: 4}, OnLoan: true},
	Book{Title: "Some Great Book", Author: "Jane", Spec: BookSpecs{Pages: 550, Rating: 2}, OnLoan: true},
	Book{Title: "Other Great Book", Author: "Jane", Spec: BookSpecs{Pages: 250}, OnLoan: false},
	Book{Title: "My Big Book", Author: "Joe", Spec: BookSpecs{Pages: 15, Rating: 5}, OnLoan: false},
}
