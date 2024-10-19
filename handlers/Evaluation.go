package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"y/models"
)

// ConvertJSONToAST converts JSON data to an Abstract Syntax Tree (AST).
func ConvertJSONToAST(jsonData string) (*models.Node, error) {
	var combined models.Node
	err := json.Unmarshal([]byte(jsonData), &combined)
	if err != nil {
		return nil, err
	}
	return &combined, nil
}

// evaluateAST evaluates the combined AST against the provided data.
func evaluateAST(node *models.Node, data map[string]interface{}) (bool, error) {
	if node == nil {
		return false, nil
	}

	leftResult, err := evaluateAST(node.Left, data)
	if err != nil {
		return false, err
	}
	rightResult, err := evaluateAST(node.Right, data)
	if err != nil {
		return false, err
	}

	switch node.Type {
	case models.Operator:
		return evaluateOperator(node.Value, leftResult, rightResult)
	case models.Operand:
		return evaluateCondition(node.Value, data)
	}

	return false, nil
}

// evaluateOperator evaluates the result of logical operators.
func evaluateOperator(operator string, leftResult, rightResult bool) (bool, error) {
	switch operator {
	case "AND":
		return leftResult && rightResult, nil
	case "OR":
		return leftResult || rightResult, nil
	}
	return false, nil // Invalid operator
}

// evaluateCondition evaluates a condition string against the provided data.
func evaluateCondition(condition string, data map[string]interface{}) (bool, error) {
	parts := strings.Fields(condition)
	if len(parts) < 3 {
		return false, nil // Invalid condition format
	}

	field, operator, value := parts[0], parts[1], parts[2]

	dataValue, exists := data[field]
	if !exists {
		return false, nil // Field not found in data
	}

	return compareValues(dataValue, value, operator)
}

// compareValues compares the data value against the provided value based on the operator.
func compareValues(dataValue interface{}, value string, operator string) (bool, error) {
	switch v := dataValue.(type) {
	case float64:
		parsedValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return false, err
		}
		return compareFloats(v, parsedValue, operator), nil
	case string:
		return compareStrings(v, value, operator), nil
	default:
		return false, nil // Unsupported data type
	}
}

// compareFloats compares two float64 values based on the operator.
func compareFloats(a, b float64, operator string) bool {
	switch operator {
	case ">":
		return a > b
	case ">=":
		return a >= b
	case "<":
		return a < b
	case "<=":
		return a <= b
	case "=":
		return a == b
	case "!=":
		return a != b
	default:
		return false // Invalid operator
	}
}

// compareStrings compares two strings based on the operator.
func compareStrings(a, b string, operator string) bool {
	switch operator {
	case "=":
		return a == b
	case "!=":
		return a != b
	default:
		return false // Invalid operator
	}
}

// setCORSHeaders sets the appropriate CORS headers.
func setCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
