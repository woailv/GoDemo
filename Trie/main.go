package main

import (
	"GoDemo/Echo"
	"fmt"
	"strings"
)

func main() {
	n := node{}
	path := "a/b/:p"
	n.insert(path, parsePattern(path), 0)
	path1 := "a/c/:p"
	n.insert(path, parsePattern(path1), 0)
	Echo.JsonPretty(n)
}

type node struct {
	Pattern  string
	Part     string
	Children []*node
	IsWild   bool
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
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
