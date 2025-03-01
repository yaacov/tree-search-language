#!/bin/bash

set -e
SCRIPT_DIR=$(dirname "$(readlink -f "$0")")
ROOT_DIR=$(dirname "$SCRIPT_DIR")
TSL_BIN="${ROOT_DIR}/tsl"
TEST_DATA_DIR="${SCRIPT_DIR}/testdata"
TEMP_DIR=$(mktemp -d)
FAILED=0
PASSED=0
TOTAL=0

# Create test data directory if it doesn't exist
mkdir -p "${TEST_DATA_DIR}"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

cleanup() {
  rm -rf "${TEMP_DIR}"
}

trap cleanup EXIT

# Check if tsl binary exists
if [ ! -f "${TSL_BIN}" ]; then
  echo -e "${RED}Error: TSL binary not found at ${TSL_BIN}${NC}"
  echo "Please build the TSL utility before running this test."
  exit 1
fi

# Function to run a test
run_test() {
  local test_name="$1"
  local test_cmd="$2"
  local expected_file="${TEST_DATA_DIR}/${test_name}.expected"
  local output_file="${TEMP_DIR}/${test_name}.output"
  local first_run=0
  
  TOTAL=$((TOTAL + 1))
  
  echo -e "${YELLOW}Running test: ${test_name}${NC}"
  
  # Run the command and capture output
  if ! eval "${test_cmd}" > "${output_file}" 2>&1; then
    echo -e "${RED}FAILED: Command execution failed${NC}"
    echo "Command: ${test_cmd}"
    cat "${output_file}"
    FAILED=$((FAILED + 1))
    return
  fi
  
  # If expected output doesn't exist, create it
  if [ ! -f "${expected_file}" ]; then
    echo -e "${YELLOW}Creating expected output file for the first time${NC}"
    cp "${output_file}" "${expected_file}"
    first_run=1
  fi
  
  # Compare with expected output
  if [ "${first_run}" -eq 0 ]; then
    if diff -q "${expected_file}" "${output_file}" > /dev/null; then
      echo -e "${GREEN}PASSED: Output matches expected result${NC}"
      PASSED=$((PASSED + 1))
    else
      echo -e "${RED}FAILED: Output doesn't match expected result${NC}"
      echo "Differences:"
      diff "${expected_file}" "${output_file}"
      FAILED=$((FAILED + 1))
    fi
  else
    echo -e "${GREEN}PASSED: First run, expected output created${NC}"
    PASSED=$((PASSED + 1))
  fi
}

# Define the tests
echo "Starting TSL stability tests..."

# Parsing tests with different query formats
run_test "01_parse_simple" "${TSL_BIN} 'name = \"alice\"'"
run_test "02_parse_complex" "${TSL_BIN} 'name = \"alice\" and (age > 30 or country in [\"US\", \"UK\"])'"

# Different operators tests
run_test "03_op_equality" "${TSL_BIN} 'name = \"bob\"'"
run_test "04_op_inequality" "${TSL_BIN} 'name != \"alice\"'"
run_test "05_op_greater_than" "${TSL_BIN} 'age > 25'"
run_test "06_op_less_equal" "${TSL_BIN} 'height <= 180'"

run_test "07_op_in" "${TSL_BIN} 'status in [\"active\", \"pending\"]'"
run_test "08_op_not_in" "${TSL_BIN} 'country not in [\"US\", \"UK\"]'"

run_test "09_op_like" "${TSL_BIN} 'name like \"%smith%\"'"
run_test "10_op_not_like" "${TSL_BIN} 'email not like \"%.gov\"'"
run_test "11_op_between" "${TSL_BIN} 'score between 50 and 100'"

# Different data types tests
run_test "12_type_number" "${TSL_BIN} 'count = 42'"
run_test "13_type_boolean" "${TSL_BIN} 'is_active = true'"
run_test "14_type_date" "${TSL_BIN} 'created_at > \"2023-01-01\"'"
run_test "15_type_date_range" "${TSL_BIN} 'updated_at between \"2023-01-01\" and \"2023-12-31\"'"

# Nested logical expressions
run_test "16_logic_nested_or_and" "${TSL_BIN} '(name = \"alice\" or name = \"bob\") and age > 25'"
run_test "17_logic_not_with_or" "${TSL_BIN} 'not (country = \"US\" or country = \"UK\") and status = \"active\"'"
run_test "18_logic_multiple_and" "${TSL_BIN} 'age > 20 and age < 30 and status = \"active\"'"
run_test "19_logic_multiple_or" "${TSL_BIN} 'category = \"A\" or category = \"B\" or category = \"C\"'"

# Special characters and edge cases
run_test "20_special_empty" "${TSL_BIN} 'notes = \"\"'"
run_test "21_special_null" "${TSL_BIN} 'tags is null'"
run_test "22_special_characters" "${TSL_BIN} 'filename = \"report-2023_Q1.pdf\"'"

# Complex queries
run_test "23_complex_search" "${TSL_BIN} '(name like \"%john%\" or email like \"%john%\") and age >= 18 and (status = \"active\" or status = \"pending\") and country not in [\"restricted\", \"blocked\"]'"
run_test "24_complex_date_and_categories" "${TSL_BIN} '(created_at between \"2023-01-01\" and \"2023-12-31\") and ((category = \"A\" and price > 100) or (category = \"B\" and price > 50))'"
run_test "25_complex_nested_groups" "${TSL_BIN} '((a = 1 or b = 2) and (c = 3 or d = 4)) or ((e = 5 or f = 6) and (g = 7 or h = 8))'"
run_test "26_complex_mixed_operators" "${TSL_BIN} 'name like \"%smith%\" and age > 30 and created_at > \"2023-01-01\" and (status in (\"active\", \"pending\") or is_featured = true)'"

# Additional tests for array notation in identifiers
run_test "27_array_notation_simple" "${TSL_BIN} 'pods[0].status = \"running\"'"
run_test "28_array_notation_wildcard" "${TSL_BIN} 'nodes[*].status = \"ready\"'"
run_test "29_array_notation_complex" "${TSL_BIN} 'deployments[0].spec.containers[*].image ~= \"nginx.*\"'"

# Function call identifier tests
run_test "30_function_call_simple" "${TSL_BIN} 'func() = true'"
run_test "31_function_call_nested" "${TSL_BIN} 'obj.func() = \"value\"'"
run_test "32_function_call_array" "${TSL_BIN} 'arr[0].func() = 123'"

# Arithmetic operation tests
run_test "33_arithmetic_add" "${TSL_BIN} 'total = price + tax'"
run_test "34_arithmetic_subtract" "${TSL_BIN} 'balance = credit - debit'"
run_test "35_arithmetic_multiply" "${TSL_BIN} 'total_price = quantity * price'"
run_test "36_arithmetic_divide" "${TSL_BIN} 'price_per_unit = total / count'"
run_test "37_arithmetic_modulo" "${TSL_BIN} 'remainder = items % 3'"
run_test "38_arithmetic_complex" "${TSL_BIN} '(price * quantity) + (tax * 0.1) > 1000'"

# Regex operator tests
run_test "39_regex_equal" "${TSL_BIN} 'name ~= \"John.*\"'"
run_test "40_regex_not_equal" "${TSL_BIN} 'name ~! \"Smith.*\"'"

# Unary operator tests
run_test "41_unary_minus" "${TSL_BIN} 'temperature = -5'"
run_test "42_unary_minus_expr" "${TSL_BIN} 'balance = -(debt + interest)'"

# Additional logical expression tests
run_test "43_logic_not_complex" "${TSL_BIN} 'not ((status = \"inactive\" and price < 0) or category = \"forbidden\")'"
run_test "44_logic_mixed_ops" "${TSL_BIN} '(age > 18 and verified = true) or (guardian_approved = true and age > 13)'"

# Report results
echo "======================================"
echo -e "Tests completed: ${TOTAL}"
echo -e "${GREEN}Tests passed: ${PASSED}${NC}"
echo -e "${RED}Tests failed: ${FAILED}${NC}"

if [ ${FAILED} -gt 0 ]; then
  echo -e "${RED}Some tests failed!${NC}"
  exit 1
else
  echo -e "${GREEN}All tests passed!${NC}"
  exit 0
fi
