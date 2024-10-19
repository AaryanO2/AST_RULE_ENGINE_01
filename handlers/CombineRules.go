package handlers

import (
	"y/models"
)

// CombineRules takes a list of rule strings and combines them into a balanced AST
func CombineRules(ruleStrings []string) *models.Node {
	if len(ruleStrings) == 0 {
		return nil
	}

	if len(ruleStrings) == 1 {
		// Only one rule, convert it into an AST and return
		res,_ := CreateRule(ruleStrings[0])
		return res // Assuming CreateRule generates the AST for a single rule
	}

	// Divide the rule strings into two halves
	mid := len(ruleStrings) / 2
	leftAST := CombineRules(ruleStrings[:mid])
	rightAST := CombineRules(ruleStrings[mid:])

	// Combine the two halves with an "AND" operator
	combinedAST := &models.Node{
		Type:  models.Operator,
		Value: "AND",
		Left:  leftAST,
		Right: rightAST,
	}

	// Print the resulting AST for debugging purposes
	printAST(combinedAST, 0)

	return combinedAST
}
