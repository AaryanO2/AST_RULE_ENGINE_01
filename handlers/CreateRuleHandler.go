package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"y/models"
)

// RuleRequest represents the structure of the incoming request for rule creation
type RuleRequest struct {
	RuleString string `json:"rule_string"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// RuleResponse represents the structure of the response containing the AST
type RuleResponse struct {
	AST *models.Node `json:"ast"`
}

func isValidParentheses(ruleString string) bool {
	count := 0
	for _, char := range ruleString {
		if char == '(' {
			count++
		} else if char == ')' {
			count--
		}
		if count < 0 {
			return false // More closing parnatheses
		}
	}
	return count == 0
}

// CreateRuleHandler handles the POST request to create an AST from a rule string
func CreateRuleHandler(w http.ResponseWriter, r *http.Request) {
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
	var req RuleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(req)
	// Create the AST from the rule string
	if !isValidParentheses(req.RuleString) {
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Unbalanced parentheses in rule string"})
		return
	}
	ast := CreateRule(req.RuleString)

	// Return the AST as a response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RuleResponse{AST: ast})
}
