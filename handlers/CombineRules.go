package handlers

import (
	"fmt"
	"strings"
	"y/models"
)

// CombineRules takes a list of rule strings and combines them into a single rule string
func CombineRules(ruleStrings []string) (*models.Node, error) {
	if len(ruleStrings) == 0 {
		return nil, nil // No rules to combine
	}

	// Join all rule strings with " AND "
	combinedRule := strings.Join(ruleStrings, " AND ")

	// Create the AST from the combined rule string
	ast, err := CreateRule(combinedRule)
	if err != nil {
		fmt.Printf("Error creating AST for combined rule: %s, Error: %s\n", combinedRule, err)
		return nil, err // Return the error if CreateRule fails
	}

	return ast, nil
}
