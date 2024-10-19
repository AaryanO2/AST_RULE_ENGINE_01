package handlers

import (
	"y/models"
)

// CombineRules takes a list of rule strings and combines them into a single AST
func CombineRules(ruleStrings []string) *models.Node {
	var asts []*models.Node

	// Convert each rule string into an AST
	for _, ruleString := range ruleStrings {
		ast := CreateRule(ruleString)
		asts = append(asts, ast)
	}

	// Combine the ASTs into a single AST using "AND" as the root operator
	if len(asts) == 0 {
		return nil
	}

	combinedAST := asts[0]
	for i := 1; i < len(asts); i++ {
		combinedAST = &models.Node{
			Type:  models.Operator,
			Value: "AND", // Using AND as the combining operator
			Left:  combinedAST,
			Right: asts[i],
		}
	}

	return combinedAST
}
