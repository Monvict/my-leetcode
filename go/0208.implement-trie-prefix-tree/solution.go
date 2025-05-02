// Created by lotus at 2025/04/26 16:23
// leetgo: 1.4.13
// https://leetcode.com/problems/implement-trie-prefix-tree/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type Trie struct {
	root *TNode
}

type TNode struct {
	children [26]*TNode
	isEnd    bool
}

func Constructor() Trie {

	return Trie{root: &TNode{}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TNode{}
		}
		node = node.children[idx]
	}

	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root

	for _, char := range word {
		idx := char - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}

	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root

	for _, ch := range prefix {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return false
		}

		node = node.children[idx]
	}

	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	ops := Deserialize[[]string](ReadLine(stdin))
	params := MustSplitArray(ReadLine(stdin))
	output := make([]string, 0, len(ops))
	output = append(output, "null")

	obj := Constructor()

	for i := 1; i < len(ops); i++ {
		switch ops[i] {
		case "insert":
			methodParams := MustSplitArray(params[i])
			word := Deserialize[string](methodParams[0])
			obj.Insert(word)
			output = append(output, "null")
		case "search":
			methodParams := MustSplitArray(params[i])
			word := Deserialize[string](methodParams[0])
			ans := Serialize(obj.Search(word))
			output = append(output, ans)
		case "startsWith":
			methodParams := MustSplitArray(params[i])
			prefix := Deserialize[string](methodParams[0])
			ans := Serialize(obj.StartsWith(prefix))
			output = append(output, ans)
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}
