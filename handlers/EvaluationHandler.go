package handlers

import (
	"encoding/json"
	"net/http"
)

// EvaluateRuleHandler takes multiple rule strings, combines them, and evaluates against the provided data.
func EvaluateRuleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		setCORSHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	setCORSHeaders(w)

	var requestBody struct {
		Rules []string               `json:"rules"`
		Data  map[string]interface{} `json:"data"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	combinedAST, err := CombineRules(requestBody.Rules)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	result, err := evaluateAST(combinedAST, requestBody.Data)
	if err != nil {
		http.Error(w, "Error evaluating rule", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"result": result})
}
