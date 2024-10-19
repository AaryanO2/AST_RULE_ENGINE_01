package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"y/models"
)

// EvaluateRule takes multiple rule strings, combines them, and evaluates against the provided data
func EvaluateRule(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	// Handle actual POST request
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// Parse the JSON body
	var requestBody struct {
		Rules []string               `json:"rules"`
		Data  map[string]interface{} `json:"data"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Combine the rules into a single AST using the existing CombineRules function
	combinedAST := CombineRules(requestBody.Rules)

	// Evaluate the combined rule against the provided data
	result, err := evaluateAST(combinedAST, requestBody.Data)
	if err != nil {
		http.Error(w, "Error evaluating rule", http.StatusInternalServerError)
		return
	}

	// Return the result as a JSON response
	response := map[string]bool{"result": result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// evaluateAST evaluates the combined AST against the provided data
func evaluateAST(node *models.Node, data map[string]interface{}) (bool, error) {
	if node == nil {
		return false, nil
	}

	switch node.Type {
	case models.Operator:
		leftResult, _ := evaluateAST(node.Left, data)
		rightResult, _ := evaluateAST(node.Right, data)

		switch node.Value {
		case "AND":
			return leftResult && rightResult, nil
		case "OR":
			return leftResult || rightResult, nil
		}
	case models.Operand:
		// Evaluate the condition, e.g., "age > 30"
		return evaluateCondition(node.Value, data)
	}

	return false, nil
}

// evaluateCondition evaluates a condition string against the provided data
func evaluateCondition(condition string, data map[string]interface{}) (bool, error) {
	// Split the condition string, e.g., "age > 30"
	parts := strings.Fields(condition)

	if len(parts) < 3 {
		return false, nil
	}

	field := parts[0]
	operator := parts[1]
	value := parts[2]

	dataValue, ok := data[field]
	if !ok {
		return false, nil
	}

	if(field == "department"){
		if(dataValue==value){
			return true,nil
		}else{
			return false,nil
		}
	}

	// Convert and compare data value based on the operator
	switch operator {
	case ">":
		return dataValue.(float64) > parseValue(value), nil
	case "<":
		return dataValue.(float64) < parseValue(value), nil
	case "=":
		return dataValue.(float64) == parseValue(value), nil
	case "!=":
		return dataValue.(float64) != parseValue(value), nil
	}

	return false, nil
}

// Helper function to parse a value from a string
func parseValue(value string) float64 {
	parsedValue, _ := strconv.ParseFloat(value, 64)
	return parsedValue
}
