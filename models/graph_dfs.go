package models

// Node Represents metro station in graph
type Node struct {
	Id       int
	LineID   int
	Title    string
	IsClosed bool
	Next     []*Node

	visited bool
}

func NewNode() *Node {
	return &Node{
		Next: make([]*Node, 0),
	}
}

func (n *Node) TraverseDFS(target int) ([]*Node, bool) {
	n.visited = true

	if n.Id == target {
		return []*Node{n}, true
	}

	for _, node := range n.Next {
		if node.IsClosed || node.visited {
			continue
		}

		path, found := node.TraverseDFS(target)
		if found {
			return append(path, n), true
		}
	}
	return nil, false
}
