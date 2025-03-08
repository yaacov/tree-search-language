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

# Basic Operators (01-10)
run_test "01_op_equality" "${TSL_BIN} 'name = \"bob\"'"
run_test "02_op_inequality" "${TSL_BIN} 'name != \"alice\"'"
run_test "03_op_greater_than" "${TSL_BIN} 'age > 25'"
run_test "04_op_less_than" "${TSL_BIN} 'age < 25'"
run_test "05_op_greater_equal" "${TSL_BIN} 'age >= 25'"
run_test "06_op_less_equal" "${TSL_BIN} 'age <= 25'"
run_test "07_op_like" "${TSL_BIN} 'name like \"%smith%\"'"
run_test "08_op_not_like" "${TSL_BIN} 'email not like \"%.gov\"'"
run_test "09_op_regex_match" "${TSL_BIN} 'name ~= \"^[A-Z].*\"'"
run_test "10_op_regex_not_match" "${TSL_BIN} 'name ~! \"^[0-9].*\"'"

# Array Operations (11-15)
run_test "11_op_in_array" "${TSL_BIN} 'status in [\"active\", \"pending\"]'"
run_test "12_op_not_in_array" "${TSL_BIN} 'country not in [\"US\", \"UK\"]'"
run_test "13_array_single" "${TSL_BIN} 'category in [\"urgent\"]'"
run_test "14_array_numbers" "${TSL_BIN} 'priority in [1, 2, 3]'"
run_test "15_array_empty" "${TSL_BIN} 'status in []'"
run_test "16_array_mixed" "${TSL_BIN} 'status in [\"normal\", \"error\\nmessage\", \"\\\"quoted\\\"\"]'"

# Data Types (16-25)
run_test "16_type_string" "${TSL_BIN} 'name = \"alice\"'"
run_test "17_type_integer" "${TSL_BIN} 'count = 42'"
run_test "18_type_float" "${TSL_BIN} 'price = 99.99'"
run_test "19_type_scientific" "${TSL_BIN} 'value = 1.23e-4'"
run_test "20_type_boolean_true" "${TSL_BIN} 'is_active = true'"
run_test "21_type_boolean_false" "${TSL_BIN} 'is_deleted = false'"
run_test "22_type_null" "${TSL_BIN} 'description is null'"
run_test "23_type_date" "${TSL_BIN} 'created_at = \"2023-12-31\"'"
run_test "24_type_rfc3339" "${TSL_BIN} 'timestamp = \"2023-12-31T23:59:59Z\"'"
run_test "25_type_rfc3339_offset" "${TSL_BIN} 'updated_at = \"2023-12-31T23:59:59+02:00\"'"

# Date Operations (26-30)
run_test "26_date_greater" "${TSL_BIN} 'date > \"2023-01-01\"'"
run_test "27_date_between" "${TSL_BIN} 'date between \"2023-01-01\" and \"2023-12-31\"'"
run_test "28_rfc3339_between" "${TSL_BIN} 'timestamp between \"2023-01-01T00:00:00Z\" and \"2023-12-31T23:59:59Z\"'"
run_test "29_date_mixed_formats" "${TSL_BIN} 'created_at > \"2023-01-01\" and updated_at < \"2023-12-31T23:59:59Z\"'"
run_test "30_date_timezone" "${TSL_BIN} 'timestamp = \"2023-12-31T23:59:59+05:30\"'"

# Logical Operations (31-35)
run_test "31_logic_and" "${TSL_BIN} 'status = \"active\" and age > 18'"
run_test "32_logic_or" "${TSL_BIN} 'type = \"admin\" or type = \"superuser\"'"
run_test "33_logic_not" "${TSL_BIN} 'not (status = \"deleted\")'"
run_test "34_logic_complex" "${TSL_BIN} '(a = 1 or b = 2) and (c = 3 or d = 4)'"
run_test "35_logic_nested_not" "${TSL_BIN} 'not (not (status = \"active\"))'"

# Size Operations (36-45)
run_test "36_size_kilobytes" "${TSL_BIN} 'size = 5k'"
run_test "37_size_megabytes" "${TSL_BIN} 'memory = 2M'"
run_test "38_size_gigabytes" "${TSL_BIN} 'storage = 1G'"
run_test "39_size_terabytes" "${TSL_BIN} 'capacity = 1T'"
run_test "40_size_binary" "${TSL_BIN} 'memory = 2Mi'"
run_test "41_size_kibi" "${TSL_BIN} 'size = 1Ki'"
run_test "42_size_mebi" "${TSL_BIN} 'size = 1Mi'"
run_test "43_size_gibi" "${TSL_BIN} 'size = 1Gi'"
run_test "44_size_decimal" "${TSL_BIN} 'size = 1.5Gi'"
run_test "45_size_comparison" "${TSL_BIN} 'size > 1.5Gi and size < 4Gi'"

# Operator Precedence (46-50)
run_test "46_precedence_and_or" "${TSL_BIN} 'a = 1 and b = 2 or c = 3'"
run_test "47_precedence_not_and" "${TSL_BIN} 'not a = 1 and b = 2'"
run_test "48_precedence_comparison_and" "${TSL_BIN} 'a > 1 and b < 2 or c >= 3'"
run_test "49_precedence_between_and" "${TSL_BIN} 'x between 1 and 10 and y = 5'"
run_test "50_precedence_complex" "${TSL_BIN} 'not a in [1,2,3] and b like \"test%\" or c between 5 and 10'"

# String Escaping (51-57)
run_test "51_escape_quotes" "${TSL_BIN} 'message = \"hello \\\"world\\\"\"'"
run_test "52_escape_backslash" "${TSL_BIN} 'path = \"C:\\\\Program Files\\\\App\"'"
run_test "53_escape_newline" "${TSL_BIN} 'text = \"line1\\nline2\"'"
run_test "54_escape_tab" "${TSL_BIN} 'format = \"column1\\tcolumn2\"'"
run_test "55_escape_mixed" "${TSL_BIN} 'content = \"Hello\\n\\\"User\\\"\\tWelcome\"'"
run_test "56_escape_in_regex" "${TSL_BIN} 'text ~= \"\\\\b\\\\w+\\\\b\"'"
run_test "57_escape_complex" "${TSL_BIN} 'description = \"Test\\\\Case\\n\\\"Special\\\"\\tChars\"'"

# Complex Expressions (58-62)
run_test "58_complex_nested_logic" "${TSL_BIN} 'not (a > 10 and (b < 20 or c = 30)) or (d in [1,2,3] and not e like \"test%\")'"
run_test "59_complex_mixed_types" "${TSL_BIN} '(size > 1Gi and name ~= \"^srv\") or (count between 1 and 10 and not status in [\"error\", \"warn\"])'"
run_test "60_complex_date_size" "${TSL_BIN} '(created_at > \"2023-01-01\" and size < 5Gi) or (updated_at < \"2023-12-31\" and memory >= 2Mi)'"
run_test "61_complex_triple_or" "${TSL_BIN} '(a = 1 and b = 2) or (c = 3 and d = 4) or (e = 5 and f = 6)'"
run_test "62_complex_mixed_arrays" "${TSL_BIN} 'tags in [\"critical\", \"high\"] and (size > 1Gi or count > 100) and not status in [\"deleted\", \"archived\"]'"

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
