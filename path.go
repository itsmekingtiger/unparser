package unparser

import (
	"fmt"
	"os"
	"strings"
)

type PathRouter struct {
	root *node
}

type node struct {
	path     string
	children []*node
	handler  func()
	wildcard *node
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
	// fmt.Println(rawPath)
	nodestrings := strings.Split(rawPath, "/")
	if rawPath == "/" {
		nodestrings[0] = "/"
	}

	P.root.add(nodestrings[1:])

}

type MatchResult struct {
	IsMatch  bool
	Wildcard map[string]string
}

func (P PathRouter) Match(rawPath string) *MatchResult {

	mRes := &MatchResult{}
	mRes.IsMatch = false
	mRes.Wildcard = make(map[string]string)
	nodestrings := strings.Split(rawPath, "/")
	if rawPath == "/" {
		nodestrings[0] = "/"
	}

	P.root.match(nodestrings[1:], mRes)
	return mRes
}

func (P PathRouter) Print() {
	P.root.print(0, false)
}

func (N *node) parse(rawPath string) {

}

func (N *node) match(rawPath []string, mRes *MatchResult) {
	// fmt.Println(rawPath[0])
	n := N.findChild(rawPath[0])
	if n == nil {
		mRes.IsMatch = false
		if N.wildcard != nil {
			mRes.Wildcard[N.wildcard.path] = rawPath[0]
			mRes.IsMatch = true
			n = N.wildcard
		} else {
			return
		}
	}

	if len(rawPath[1:]) == 0 {
		mRes.IsMatch = true
		return
	}
	n.match(rawPath[1:], mRes)
}

func (N *node) append(rawPath string) *node {
	node := newNode(rawPath)
	N.children = append(N.children, node)
	return node
}

func (N *node) setWildcard(rawPath string) *node {
	if N.wildcard != nil {
		if N.wildcard.path != rawPath {
			fmt.Println("ONLY ONE WILDCARD NAME CAN BE USED IN EACH PATH POSITION")
			os.Exit(1)
		}
		return N.wildcard
	}
	node := newNode(rawPath)
	N.wildcard = node
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

// func (N *node) findWildCard(rawPath string) *node {
// }

func (N *node) findChild(rawPath string) *node {
	for _, c := range N.children {
		if c.path == rawPath {
			return c
		}
	}
	return nil
}

func (N *node) findOrCreateChild(rawPath string) *node {
	if rawPath[0] == '$' {
		// fmt.Printf("딸라는: %s\n", rawPath)
		return N.setWildcard(rawPath[1:])
	}

	n := N.findChild(rawPath)
	if n != nil {
		return n
	}
	return N.append(rawPath)
}

func (N node) print(depth int, wildcard bool) {
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}
	if wildcard {
		fmt.Printf("%s- ${%s}\n", indent, N.path)
	} else {
		fmt.Printf("%s- %s\n", indent, N.path)
	}

	for _, child := range N.children {
		child.print(depth+1, false)
	}
	if N.wildcard != nil {
		N.wildcard.print(depth+1, true)
	}
}
