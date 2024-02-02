package main

import (
	"fmt"
	"strings"
)

type TrieTree struct {
	Children map[string]*TrieTree
	IsEnd    bool
}

func (t *TrieTree) Insert(route string) {

	node := t
	paths := strings.Split(route, "/")

	// /v1/user/register/callback
	for _, path := range paths {
		if path == "" {
			continue
		}
		if _, ok := node.Children[path]; !ok {
			node.Children[path] = &TrieTree{Children: make(map[string]*TrieTree)}
		}
		node = node.Children[path]
	}
	node.IsEnd = true
}

func (t *TrieTree) StartWith(route string) bool {
	node := t
	paths := strings.Split(route, "/")
	for _, path := range paths {
		if path == "" {
			continue
		}
		if v, ok := node.Children[path]; ok {
			node = v
		} else {
			return false
		}
	}
	return true
}

func (t *TrieTree) Search(route string) bool {
	node := t
	paths := strings.Split(route, "/")
	for _, path := range paths {
		if path == "" {
			continue
		}
		if v, ok := node.Children[path]; ok {
			node = v
		} else {
			return false
		}
	}
	return node.IsEnd
}

func (t *TrieTree) Print(root string) {
	for path, child := range t.Children {
		fullPath := root + "/" + path
		fmt.Println(fullPath)
		child.Print(fullPath)
	}
}
func main() {
	t := &TrieTree{
		Children: make(map[string]*TrieTree),
	}
	t.Insert("/v1/user/login/callback")
	//t.Insert("/v1/user/login/event")
	//t.Insert("/v1/user/login/event")
	t.Insert("/v2/fuck/shift/ok")
	//
	//t.Print("")
	//fmt.Println(t.Search("/v2/fuck/shift/ok"))
	fmt.Println(t.StartWith("/v2/fuck/shift"))

}
