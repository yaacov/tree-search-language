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

echo -e "\n${GREEN}All tests completed!${NC}"
