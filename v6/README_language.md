# Tree Search Language (TSL) Reference

A quick overview of TSL syntax, operators, literals and identifiers.

## 1. Basic Structure

- Expressions combine identifiers, literals, operators and parentheses:
  ```
  (field1 = 'value' OR field2 > 10) AND NOT field3 IN [1,2,3]
  ```

## 2. Identifiers

- Start with a letter or underscore, may include letters, digits, `_`, `.`, `/`, `-`
- Array suffixes: `[index]`, `[*]`, `[name]`
- Examples:
  ```
  name
  user.age
  pods[0].status
  services[my.service].ip
  ```

## 3. Literals

- String: `'text'`, `"text"`, `` `text` ``
- Numeric: integer, decimal, scientific, with optional SI suffix (`Ki`, `M`, etc.)
- Date/Time: `YYYY-MM-DD` or RFC3339 `YYYY-MM-DDThh:mm:ssZ`
- Boolean: `true`, `false`
- Null: `null`
- Arrays: `[expr, expr, ...]`

## 4. Operators

1. Logical
   - `AND`, `OR`, `NOT`
2. Comparison
   - `=`, `!=`, `<`, `<=`, `>`, `>=`
3. Pattern
   - `LIKE`, `ILIKE` (case‑insensitive), `~=` (regex match), `~!` (regex not match)
4. Membership
   - `IN`, `NOT IN`, `BETWEEN … AND …`
5. Arithmetic
   - `+`, `-`, `*`, `/`, `%`
6. Array functions
   - `LEN x`, `ANY x`, `ALL x`, `SUM x`

## 5. Precedence (high→low)

1. Unary: `NOT`, `LEN`, `ANY`, `ALL`, `SUM`, unary `-`
2. `*`, `/`, `%`
3. `+`, `-`
4. `IN`, `BETWEEN`, `LIKE`, `ILIKE`, `IS`, etc.
5. Comparisons: `=`, `!=`, `<`, `<=`, `>`, `>=`, `~=`, `~!`
6. `AND`
7. `OR`

## 6. Examples

```sql
# combine filters
(name LIKE '%joe%' OR city = 'milan') AND age BETWEEN 20 AND 30

# array operations
tags IN ['a','b','c']
SUM scores > 100
ANY (values > 5)

# date comparison
created_at >= '2021-01-01T00:00:00Z'
```
