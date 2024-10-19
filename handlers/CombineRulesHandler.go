package handlers

import (
	"encoding/json"
	"net/http"
	"y/models"
)

// CombineRulesRequest represents the structure of the incoming request for combining rules
type CombineRulesRequest struct {
	RuleStrings []string `json:"rule_strings"`
}

// CombineRulesResponse represents the structure of the response containing the combined AST
type CombineRulesResponse struct {
	CombinedAST *models.Node `json:"combined_ast"`
}

// CombineRulesHandler handles the POST request to combine multiple rules into a single AST
func CombineRulesHandler(w http.ResponseWriter, r *http.Request) {
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
	var req CombineRulesRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Combine the rules into a single AST
	combinedAST := CombineRules(req.RuleStrings)

	// Return the combined AST as a response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CombineRulesResponse{CombinedAST: combinedAST})
}
