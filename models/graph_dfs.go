package models

// Node Represents metro station in graph
type Node struct {
	Id       int
	LineID   int
	Title    string
	IsClosed bool
	Next     []*Node

	// internal fields
	visited bool

	// black  field is used for calculating all the possible routes
	// it is true if all nodes in n.Next are visited
	black bool
}

type Path struct {
	Root     []*Node
	switches int
	length   int
}

func NewNode() *Node {
	return &Node{
		Next: make([]*Node, 0),
	}
}

func (n *Node) TraverseDFS(target int) []*Path {
	n.visited = true

	if n.Id == target {
		return []*Path{{
			Root: []*Node{n},
		}}
	}

	var outPath []*Path

	var visited []*Node
	var unvisited []*Node

	for _, node := range n.Next {
		if node.visited {
			visited = append(visited, node)
		} else {
			unvisited = append(unvisited, node)
		}
	}

	for _, node := range unvisited {
		if node.IsClosed || node.black {
			continue
		}

		path := node.TraverseDFS(target)
		for _, p := range path {
			if len(p.Root) > 0 && p.Root[len(p.Root)-1].LineID != n.LineID {
				p.switches++
			}
			p.length++
			p.Root = append(p.Root, n)
		}
		outPath = append(outPath, path...)
	}

	n.black = true

	for _, node := range visited {
		if node.IsClosed || node.black {
			continue
		}

		path := node.TraverseDFS(target)
		for _, p := range path {
			if len(p.Root) > 0 && p.Root[len(p.Root)-1].LineID != n.LineID {
				p.switches++
			}
			p.length++
			p.Root = append(p.Root, n)
		}
		outPath = append(outPath, path...)
	}

	return outPath
}
