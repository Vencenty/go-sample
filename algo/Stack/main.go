package main

import (
	"fmt"
	"strings"
)

type Tree struct {
	Children map[string]*Tree
	IsEnd    bool
}

// /v1/user/callback
func (t *Tree) Insert(word string) {
	paths := strings.Split(word, "/")
	node := t
	for _, path := range paths {
		if path == "" {
			continue
		}
		if _, ok := node.Children[path]; !ok {
			node.Children[path] = &Tree{
				Children: make(map[string]*Tree),
			}
		}
		node = node.Children[path]
	}
	node.IsEnd = true
}

func (t *Tree) Search(word string) bool {
	paths := strings.Split(word, "/")
	node := t
	for _, path := range paths {
		if path == "" {
			continue
		}
		if searched, ok := node.Children[path]; ok {
			node = searched
		} else {
			node.IsEnd = false
		}
	}
	return node.IsEnd
}

func (t *Tree) StartWith(word string) bool {
	paths := strings.Split(word, "/")
	node := t
	for _, path := range paths {
		if path == "" {
			continue
		}
		if searched, ok := node.Children[path]; ok {
			node = searched
		} else {
			node.IsEnd = false
		}
	}
	return node.IsEnd
}

func (t *Tree) PrintTree() {
	for path, v := range t.Children {
		if v.IsEnd {
			fmt.Println("这是最后的路径了", path)
		}
		v.PrintTree()
	}
}

// 后入先出
func main() {
	t := &Tree{Children: make(map[string]*Tree)}
	t.Insert("/v1/auth/callback")
	t.Insert("/v1/user/register")
	t.Insert("/v1/user/get/name")
	t.Insert("/v1/user/resetPassword")
	t.Insert("/v2/user/resetPassword")
	t.Insert("/v3/user/resetPassword")

	r := t.Search("/v1/user/resetPassword")
	fmt.Println(r)
	t.PrintTree()

}
