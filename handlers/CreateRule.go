package handlers

import (
	"regexp"
	"strings"
	"y/models"
)

// parseRuleString splits the rule string into tokens using regular expressions.
func parseRuleString(ruleString string) []string {
	re := regexp.MustCompile(`[\w']+|[()=><!]`)
	tokens := re.FindAllString(ruleString, -1)
	return tokens
}

// CreateRule converts a rule string into an AST Node.
func CreateRule(ruleString string) *models.Node {
	tokens := parseRuleString(ruleString)
	ast, _ := convertTokensToAST(tokens, 0)
	return ast
}


func convertTokensToAST(tokens []string, index int) (*models.Node, int) {
	var currentNode *models.Node
	nodeStack := []*models.Node{} // Stack to manage nodes in nested expressions

	for index < len(tokens) {
		token := tokens[index]

		switch token {
		case "(":
			// Push the current node to the stack and start a new subtree
			if currentNode != nil {
				nodeStack = append(nodeStack, currentNode)
			}
			currentNode = nil // Start new subtree
		case ")":
			// Close the current subtree and attach it to the parent node from the stack
			if len(nodeStack) > 0 {
				top := nodeStack[len(nodeStack)-1] // Parent node
				nodeStack = nodeStack[:len(nodeStack)-1]

				// Attach current subtree to the parent node
				if top.Left == nil {
					top.Left = currentNode
				} else {
					top.Right = currentNode
				}
				currentNode = top // Current node is now the parent node
			}
		case "AND", "OR":
			// Set the operator, create a new node, and attach the left operand
			newNode := &models.Node{
				Type:  models.Operator,
				Value: strings.ToUpper(token),
				Left:  currentNode,
			}
			currentNode = newNode // Now expecting to attach the right operand
		default:
			// It's an operand, so attach it to the current node
			condition := token
			if index+2 < len(tokens) && (tokens[index+1] == ">" || tokens[index+1] == "<" || tokens[index+1] == "=" || tokens[index+1] == "!=") {
				condition += " " + tokens[index+1] + " " + tokens[index+2]
				index += 2
			}
			operandNode := &models.Node{
				Type:  models.Operand,
				Value: condition,
			}
			if currentNode == nil {
				currentNode = operandNode // This is the first operand in the expression
			} else {
				// Attach as the right node if currentNode is an operator
				if currentNode.Right == nil {
					currentNode.Right = operandNode
				} else {
					currentNode.Left = operandNode
				}
			}
		}
		index++
	}

	// Return the fully constructed tree (currentNode should be the root)
	return currentNode, index
}
