package parser

import "testing"

func FuzzParse(f *testing.F) {
	f.Add("name = 'alice'")
	f.Add("age > 25 and city = 'rome'")
	f.Add("status in ['active', 'pending']")
	f.Add("price between 10 and 100")
	f.Add("not (deleted = true)")
	f.Add("title like '%book%'")
	f.Add("created_at > 2023-01-01")
	f.Add("(a = 1 or b = 2) and c = 3")
	f.Add("")
	f.Add("   ")
	f.Add("名前 = 'test'")
	f.Add("name = '🎉'")
	f.Add("[1, 2, 3]")
	f.Add("a ~= '.*'")

	f.Fuzz(func(t *testing.T, input string) {
		// Should not panic regardless of input
		_, _ = Parse(input)
	})
}
