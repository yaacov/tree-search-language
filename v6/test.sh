#!/bin/bash

# Ensure the script exits on any error
set -e

# Build the TSL tool if not already built
make

# Text colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Helper function to run a test
run_test() {
    local description="$1"
    local input="$2"
    
    echo -e "\n${BLUE}Testing: ${description}${NC}"
    echo "Input: $input"
    echo -e "${GREEN}Output:${NC}"
    ./tsl "$input"
    echo "----------------------------------------"
}

# Basic comparisons
run_test "Equal comparison" "name = 'John'"
run_test "Not equal comparison" "age != 25"
run_test "Greater than" "price > 100"
run_test "Less than or equal" "quantity <= 50"
run_test "Regular expression equal" "name ~= 'John.*'"
run_test "Regular expression not equal" "name ~! 'Smith.*'"

# Arithmetic operations
run_test "Addition" "total = price + tax"
run_test "Subtraction" "balance = credit - debit"
run_test "Multiplication" "total_price = quantity * price"
run_test "Division" "price_per_unit = total / count"
run_test "Modulo" "remainder = items % 3"
run_test "Complex arithmetic" "(price * quantity) + (tax * 0.1) > 1000"

# Unary operators
run_test "Unary minus" "temperature = -5"
run_test "Unary minus with expression" "balance = -(debt + interest)"
run_test "Unary NOT" "not is_deleted"
run_test "Unary NOT with comparison" "not (price > 100)"
run_test "Unary NOT with complex condition" "not (status = 'inactive' or price < 0)"

# Logical operators
run_test "AND operation" "age > 20 and age < 30"
run_test "OR operation" "status = 'active' or status = 'pending'"
run_test "NOT operation" "not (price < 10)"
run_test "Complex logical expression" "(age > 20 and status = 'active') or (priority = 'high')"

# Arrays and IN operator
run_test "IN operator with parentheses" "status in ('active', 'pending', 'completed')"
run_test "IN operator with brackets" "color in ['red', 'blue', 'green']"
run_test "NOT IN operator" "category not in ['misc', 'unknown']"

# BETWEEN operator
run_test "BETWEEN operator" "age between 18 and 65"
run_test "NOT BETWEEN operator" "price not between 100 and 200"

# NULL checks
run_test "IS NULL check" "comment is null"
run_test "IS NOT NULL check" "email is not null"

# String patterns
run_test "LIKE operator" "name like 'J%'"
run_test "ILIKE operator" "email ilike '%@gmail.com'"
run_test "NOT LIKE operator" "description not like '%test%'"

# Boolean values
run_test "Boolean true" "is_active = true"
run_test "Boolean false" "is_deleted = false"

# Date and time
run_test "Date comparison" "created_at = '2023-01-01'"
run_test "RFC3339 timestamp" "timestamp = '2023-01-01T15:30:00Z'"
run_test "Date range" "created_at between '2023-01-01' and '2023-12-31'"

# Numbers with size suffixes
run_test "Binary size (Ki)" "size > 50Ki"
run_test "Decimal size (M)" "bandwidth < 100M"
run_test "Binary size (Gi)" "storage = 1Gi"

# Complex combinations
run_test "Complex query 1" "status in ['active', 'pending'] and (price between 10 and 100) and not is_deleted"
run_test "Complex query 2" "(name like 'Tech%' or category = 'electronics') and price > 1000 and stock > 0"
run_test "Complex query 3" "created_at > '2023-01-01' and (status = 'active' or priority in ['high', 'critical']) and size < 5Gi"

# Array notation in identifiers
run_test "Array index in identifier" "pods[0].status = 'running'"
run_test "Multiple array indices" "containers[2].ports[80].protocol = 'TCP'"
run_test "Wildcard array notation" "nodes[*].status = 'ready'"
run_test "Mixed array and dot notation" "cluster.nodes[3].pods[*].status = 'running'"
run_test "Complex array path" "deployments[0].spec.containers[*].image ~= 'nginx.*'"
run_test "Nested arrays with conditions" "pods[0].containers[*].resources.limits.memory < 1Gi"
run_test "Multiple wildcards" "nodes[*].pods[*].status != 'failed'"

# New tests for ARRAY_SUFFIX
run_test "Array suffix with index" "services[0].status = 'active'"
run_test "Array suffix with wildcard" "nodes[*].status = 'ready'"
run_test "Array suffix with identifier" "pods[my-pod].status = 'running'"
run_test "Array suffix with complex identifier" "deployments[my-deployment.spec.containers/nginx].image ~= 'nginx.*'"
run_test "Array suffix with dot" "services[my.service].status = 'active'"
run_test "Array suffix with slash" "nodes[my/node].status = 'ready'"
run_test "Array suffix with hyphen" "pods[my-pod] = 'running'"

# Identifiers ending with ()
run_test "Function call identifier" "func() = true"
run_test "Nested function call identifier" "obj.func() = 'value'"
run_test "Function call with array index" "arr[0].func() = 123"
run_test "Function call with complex path" "obj.arr[0].func() != 'value'"

echo -e "\n${GREEN}All tests completed!${NC}"
