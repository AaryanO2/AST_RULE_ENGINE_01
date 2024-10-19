package handlers

import (
	"fmt"
	"regexp"
	"strings"
	"y/models"
)

// parseRuleString splits the rule string into tokens using regular expressions.
func parseRuleString(ruleString string) []string {
	re := regexp.MustCompile(`[\w']+|[()=><!]+`)
	tokens := re.FindAllString(ruleString, -1)
	return tokens
}

// debugging function to print the AST
func printAST(node *models.Node, level int) {
	if node == nil {
		return
	}

	// Print the current node with indentation for levels
	fmt.Printf("%s%s\n", strings.Repeat("  ", level), node.Value)

	// Recursively print the left and right children, if they exist
	printAST(node.Left, level+1)
	printAST(node.Right, level+1)
}

// CreateRule converts a rule string into an AST Node.
func CreateRule(ruleString string) (*models.Node, error) {
	tokens := parseRuleString(ruleString)
	ast, err := convertTokensToAST(tokens, 0)
	if err != nil {
		return nil, err
	}
	printAST(ast, 0) // Debugging: Print the generated AST
	return ast, nil
}

func convertTokensToAST(tokens []string, index int) (*models.Node, error) {
	var currentNode *models.Node
	nodeStack := []*models.Node{} // Stack to manage nodes in nested expressions
	parenthesesCount := 0

	for index < len(tokens) {
		token := tokens[index]

		switch token {
		case "(":
			parenthesesCount++
			// Push the current node to the stack and start a new subtree
			if currentNode != nil {
				nodeStack = append(nodeStack, currentNode)
			}
			currentNode = nil // Start new subtree
		case ")":
			parenthesesCount--
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
			} else {
				return nil, fmt.Errorf("unmatched closing parenthesis at token %s", token)
			}
		case "AND", "OR":
			// Ensure there's a current node before adding an operator
			if currentNode == nil {
				return nil, fmt.Errorf("operator %s cannot be placed without a preceding operand", token)
			}

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
			// Check for valid comparisons including new operators
			if index+2 < len(tokens) && (tokens[index+1] == ">" || tokens[index+1] == "<" || tokens[index+1] == "=" || tokens[index+1] == "!=" || tokens[index+1] == ">=" || tokens[index+1] == "<=" || tokens[index+1] == "==") {
				// Ensure valid comparison
				condition += " " + tokens[index+1] + " " + tokens[index+2]
				index += 2
			} else if index+1 < len(tokens) && (tokens[index+1] == ">" || tokens[index+1] == "<" || tokens[index+1] == "=" || tokens[index+1] == "!=" || tokens[index+1] == ">=" || tokens[index+1] == "<=" || tokens[index+1] == "==") {
				return nil, fmt.Errorf("missing operand after comparison operator %s", tokens[index+1])
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
					// Attach to the left if already has a right child
					currentNode.Left = operandNode
				}
			}
		}
		index++
	}

	if parenthesesCount != 0 {
		return nil, fmt.Errorf("mismatched parentheses: %d unclosed parentheses", parenthesesCount)
	}

	// Return the fully constructed tree (currentNode should be the root)
	return currentNode, nil
}
