package parser

import (
	"fmt"
	"sync"
	"testing"
)

func TestParseConcurrent(t *testing.T) {
	expressions := []string{
		"name = 'alice'",
		"age > 25 and city = 'rome'",
		"status in ['active', 'pending']",
		"price between 10 and 100",
		"not (deleted = true)",
		"title like '%book%'",
		"created_at > 2023-01-01",
		"(a = 1 or b = 2) and c = 3",
	}

	var wg sync.WaitGroup
	errors := make(chan error, len(expressions)*10)

	for i := 0; i < 10; i++ {
		for _, expr := range expressions {
			wg.Add(1)
			go func(input string) {
				defer wg.Done()
				node, err := Parse(input)
				if err != nil {
					errors <- err
					return
				}
				if node == nil {
					errors <- fmt.Errorf("got nil node for input: %s", input)
				}
			}(expr)
		}
	}

	wg.Wait()
	close(errors)

	for err := range errors {
		t.Errorf("concurrent parse failed: %v", err)
	}
}
