package main

import (
	"fmt"
	"strings"
)

// Tree 前缀索引树
type Tree struct {
	Children map[string]*Tree
	IsEnd    bool
}

func (t *Tree) Insert(value string) {
	node := t

	paths := strings.Split(value, "/")

	// user callback haha/
	for _, path := range paths {
		if path == "" {
			continue
		}

		if _, exists := node.Children[path]; !exists {
			node.Children[path] = &Tree{Children: make(map[string]*Tree)}
		}
		node = node.Children[path]
	}
	node.IsEnd = true
}

func (t *Tree) Match(value string) bool {
	node := t
	paths := strings.Split(value, "/")
	for _, path := range paths {
		if path == "" {
			continue
		}

		if child, exists := node.Children[path]; exists {
			//node.IsEnd = true
			node = child
		} else {
			return false
		}
	}

	return node.IsEnd // 检查是否为路径的结束
}

func (t *Tree) StartWith(prefix string) bool {
	node := t
	paths := strings.Split(prefix, "/")
	// user callback haha/
	for _, path := range paths {
		if path == "" {
			continue
		}

		if child, exists := node.Children[path]; exists {
			node = child
		} else {
			return false
		}
	}
	return node.IsEnd
}

// printTree 用于打印树的结构，辅助验证
func printTree(t *Tree, prefix string) {
	for key, child := range t.Children {
		fmt.Println(prefix + "/" + key)
		printTree(child, prefix+"/"+key)
	}
	//if t.IsEnd {
	//	fmt.Println(prefix + " [End]")
	//}
}
func main() {
	root := Tree{
		Children: make(map[string]*Tree),
		IsEnd:    false,
	}

	//r := strings.Split("/user/callback", "/")

	root.Insert("/user/callback/haha")
	root.Insert("/user/callback/zap/log")
	root.Insert("/user/event")
	root.Insert("/user/shift")

	matchResult := root.Match("/user/")
	startResult := root.StartWith("/user/")
	fmt.Println(matchResult, startResult)

}
