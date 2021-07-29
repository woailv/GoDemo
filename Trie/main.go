package main

import (
	"GoDemo/Echo"
	"fmt"
	"strings"
)

func main() {
	n := node{}
	n.insert("/a/:p", []string{"a", ":p"}, 0)
	Echo.Json(n)
	ns := n.search([]string{"a", "b"}, 0)
	Echo.Json(ns)
}

type node struct {
	Pattern  string
	Part     string
	Children []*node
	IsWild   bool
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.Pattern, n.Part, n.IsWild)
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.Pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{Part: part, IsWild: part[0] == ':' || part[0] == '*'}
		n.Children = append(n.Children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.Part, "*") {
		if n.Pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

func (n *node) travel(list *([]*node)) {
	if n.Pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.Children {
		child.travel(list)
	}
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.Children {
		if child.Part == part || child.IsWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.Children {
		if child.Part == part || child.IsWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
