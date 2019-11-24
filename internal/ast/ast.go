package ast

import "fmt"

// Nodes type represents an array of Nodes.
type Nodes []Node

// Node represents a parsed lisp node.
type Node interface {
	fmt.Stringer
}

// StrNode represents a string node
type StrNode struct {
	s string
}

func (n *StrNode) String() string {
	return n.s
}
