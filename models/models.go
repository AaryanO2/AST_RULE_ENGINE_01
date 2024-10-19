package models

// NodeType represents the type of node: operator or operand.
type NodeType string

const (
	Operator NodeType = "operator"
	Operand  NodeType = "operand"
)

// Node represents an AST node.
type Node struct {
	Type  NodeType // "operator" or "operand"
	Left  *Node    // Left child node
	Right *Node    // Right child node
	Value string   // Value of the node, e.g., "AND", "age > 30"
}
