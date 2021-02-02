package unparser

import (
	"fmt"
	"strings"
)

type PathRouter struct {
	root *node
}

type node struct {
	path     string
	children []*node
	handler  func()
}

/*
   user/1
   /user/1
   /user/1/post/1
*/
//
// /user/:id
// /attr/:attr/prop/:prop
func DefaultPathRouter() *PathRouter {
	return &PathRouter{
		root: &node{
			path:     "/",
			children: make([]*node, 0, 2),
		},
	}
}

func (P *PathRouter) Parse(rawPath string) {
	if len(rawPath) == 0 && rawPath == "/" {

	}
}

func (P *PathRouter) Add(rawPath string) {
	fmt.Println(rawPath)
	nodestrings := strings.Split(rawPath, "/")
	if rawPath == "/" {
		nodestrings[0] = "/"
	}

	P.root.add(nodestrings[1:])

}

func (P PathRouter) Match(rawPath string) {
	nodestrings := strings.Split(rawPath, "/")
	if rawPath == "/" {
		nodestrings[0] = "/"
	}

	P.root.match(nodestrings)
}

func (P PathRouter) Print() {
	P.root.print(0)
}

func (N *node) parse(rawPath string) {

}

func (N *node) match(rawPath []string) bool {
	if N.path == rawPath[0] {
		if len(rawPath) == 0 {
			return true
		}
		for _, c := range N.children {
			if c.match(rawPath) {
				return true
			}
		}
	}
	return false
}
func (N *node) append(rawPath string) *node {
	node := newNode(rawPath)
	N.children = append(N.children, node)
	return node
}

func newNode(rawPath string) *node {
	return &node{
		path:     rawPath,
		children: make([]*node, 0, 2),
	}
}

// 주어진 path 배열 재귀적으로 추가
func (N *node) add(rawPath []string) {
	if len(rawPath) == 0 {
		return
	}

	n := N.findOrCreateChild(rawPath[0])
	n.add(rawPath[1:])
}

func (N *node) hasChild(rawPath string) bool {
	for _, c := range N.children {
		if c.path == rawPath {
			return true
		}
	}
	return false
}

func (N *node) findChild(rawPath string) *node {
	for _, c := range N.children {
		if c.path == rawPath {
			return c
		}
	}
	return nil
}

func (N *node) findOrCreateChild(rawPath string) *node {
	/// 1
	n := N.findChild(rawPath)
	if n != nil {
		return n
	}
	return N.append(rawPath)

	/// 2
	// n := N.findChild(rawPath)
	// if n == nil {
	// 	n = N.append(rawPath)
	// }
	// return n
}

func (N node) print(depth int) {
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}
	fmt.Printf("%s- %s\n", indent, N.path)

	for _, child := range N.children {
		child.print(depth + 1)
	}
}
