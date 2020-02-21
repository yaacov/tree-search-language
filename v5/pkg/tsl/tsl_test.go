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

package tsl

import (
	"encoding/json"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/yaacov/tree-search-language/v5/pkg/parser"
)

func TestParser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TSL")
}

var _ = Describe("Parser", func() {
	DescribeTable("Returns the expected tree",
		func(input, expected string) {
			// Parse the input:
			tree, err := parseTSL(input)
			Expect(err).ToNot(HaveOccurred())

			// Convert the tree to JSON:
			actual, err := json.Marshal(tree)
			Expect(err).ToNot(HaveOccurred())

			// Check that the actual JSON and the expected JSON are structurally
			// identical:
			Expect(actual).To(MatchJSON(expected))
		},

		Entry(
			"Simple comparison",
			"a = 'hello'",
			`{
				"func": "$eq",
				"left": {
					"func": "$ident",
					"left": "a"
				},
				"right": {
					"func": "$string",
					"left": "hello"
				}
			}`,
		),

		Entry(
			"Order of operators",
			"a = 12.3e1 or b = 'world' and c = 'hello'",
			`{
				"func": "$or",
				"left": {
					"func": "$eq",
					"left": {
						"func": "$ident",
						"left": "a"
					},
					"right": {
						"func": "$number",
						"left": 123
					}
				},
				"right": {
					"func": "$and",
					"left": {
						"func": "$eq",
						"left": {
							"func": "$ident",
							"left": "b"
						},
						"right": {
							"func": "$string",
							"left": "world"
						}
					},
					"right": {
						"func": "$eq",
						"left": {
							"func": "$ident",
							"left": "c"
						},
						"right": {
							"func": "$string",
							"left":"hello"
						}
					}
				}
			}`,
		),

		Entry(
			"Simple use of 'like'",
			"a like 'my%'",
			`{
				"func": "$like",
				"left": {
					"func": "$ident",
					"left": "a"
				},
				"right": {
					"func": "$string",
					"left": "my%"
				}
			}`,
		),

		Entry(
			"Simple use of 'ilike'",
			"a ilike 'my%'",
			`{
				"func": "$ilike",
				"left": {
					"func": "$ident",
					"left": "a"
				},
				"right": {
					"func": "$string",
					"left": "my%"
				}
			}`,
		),

		Entry(
			"Use of 'not' with 'like'",
			"a not like 'my%'",
			`{
				"func": "$not",
				"left": {
					"func": "$like",
					"left": {
						"func": "$ident",
						"left": "a"
					},
					"right": {
						"func": "$string",
						"left": "my%"
					}
				}
			}`,
		),

		Entry(
			"Use of 'not' with 'ilike'",
			"a not ilike 'my%'",
			`{
				"func": "$not",
				"left": {
					"func": "$ilike",
					"left": {
						"func": "$ident",
						"left": "a"
					},
					"right": {
						"func": "$string",
						"left": "my%"
					}
				}
			}`,
		),

		Entry(
			"Use of . and / as valid identifier",
			"a.b.c.d.f/g.h not ilike 'my%'",
			`{
				"func": "$not",
				"left": {
					"func": "$ilike",
					"left": {
						"func": "$ident",
						"left": "a.b.c.d.f/g.h"
					},
					"right": {
						"func": "$string",
						"left": "my%"
					}
				}
			}`,
		),

		Entry(
			"Use escaping with identifier",
			"`a` > [b] and \"c\" is not null",
			`{
				"func": "$and",
				"left": {
				  "func": "$gt",
				  "left": {
					"func": "$ident",
					"left": "a"
				  },
				  "right": {
					"func": "$ident",
					"left": "b"
				  }
				},
				"right": {
				  "func": "$exists",
				  "left": {
					"func": "$ident",
					"left": "c"
				  }
				}
			  }`,
		),

		Entry(
			"Use identifiers with math operators on right hand side",
			"a > (b + 1)",
			`{
				"func": "$gt",
				"left": {
				  "func": "$ident",
				  "left": "a"
				},
				"right": {
				  "func": "$add",
				  "left": {
					"func": "$ident",
					"left": "b"
				  },
				  "right": {
					"func": "$number",
					"left": 1
				  }
				}
			  }`,
		),

		Entry(
			"Use SI units",
			"a > 1Ki and b < 2Mi and c = 3Gi or d != 4Ti",
			`{
				"func": "$or",
				"left": {
				  "func": "$and",
				  "left": {
					"func": "$and",
					"left": {
					  "func": "$gt",
					  "left": {
						"func": "$ident",
						"left": "a"
					  },
					  "right": {
						"func": "$number",
						"left": 1024
					  }
					},
					"right": {
					  "func": "$lt",
					  "left": {
						"func": "$ident",
						"left": "b"
					  },
					  "right": {
						"func": "$number",
						"left": 2097152
					  }
					}
				  },
				  "right": {
					"func": "$eq",
					"left": {
					  "func": "$ident",
					  "left": "c"
					},
					"right": {
					  "func": "$number",
					  "left": 3221225472
					}
				  }
				},
				"right": {
				  "func": "$ne",
				  "left": {
					"func": "$ident",
					"left": "d"
				  },
				  "right": {
					"func": "$number",
					"left": 4398046511104
				  }
				}
			  }`,
		),
	)
})

// parseTSL parse the TSL.
func parseTSL(input string) (n Node, err error) {
	// Setup the TSL ErrorListener.
	errorListener := NewErrorListener()

	// Setup the input.
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewTSLLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser.
	p := parser.NewTSLParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	// Finally parse the expression (by walking the tree).
	var listener Listener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	err = errorListener.Err
	if err != nil {
		return
	}

	n, err = listener.GetTree()

	return
}
