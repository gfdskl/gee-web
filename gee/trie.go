package gee

import "fmt"

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

func newNode() *node {
	return &node{}
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.isWild || child.part == part {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	children := make([]*node, 0)
	for _, child := range n.children {
		if child.isWild || child.part == part {
			children = append(children, child)
		}
	}
	return children
}

func (n *node) insert(pattern string, parts []string, height int) {
	if height == len(parts) {
		if n.pattern != "" {
			panic(fmt.Sprintf("pattern exists: %s", n.pattern))
		}
		n.pattern = pattern
		return
	}
	child := n.matchChild(parts[height])
	if child == nil {
		child = &node{
			part:   parts[height],
			isWild: parts[height][0] == ':' || parts[height][0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if height == len(parts) || len(n.part) != 0 && n.part[0] == '*' {
		if n.pattern != "" {
			return n
		}
		return nil
	}
	children := n.matchChildren(parts[height])
	for _, child := range children {
		res := child.search(parts, height+1)
		if res != nil {
			return res
		}
	}
	return nil
}

func (n *node) tranverse() {
	if n == nil {
		return
	}
	fmt.Println(n.pattern, n.part, n.isWild)
	for _, child := range n.children {
		child.tranverse()
	}
}
