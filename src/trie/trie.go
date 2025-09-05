package trie

import (
	"fmt"
	"unicode"
)

type Trie struct {
	Root *Node
}

type Node struct {
	Character rune
	IsWord    bool
	Children  [26]*Node
}

func (trie *Trie) Add(word string) {
	cur := trie.Root
	var idx int
	for _, char := range word {
		idx = getCharIndex(char)

		if cur.Children[idx] != nil {
			cur = cur.Children[idx]
			continue
		}

		cur.Children[idx] = &Node{
			Character: char,
			IsWord:    false,
			Children:  [26]*Node{},
		}

		cur = cur.Children[idx]
		fmt.Println(string(char))
	}

	fmt.Println(cur)
	cur.IsWord = true
	fmt.Println(cur)
}

func (trie *Trie) Search(word string) bool {

	cur := trie.Root
	var idx int
	for _, char := range word {
		idx = getCharIndex(char)

		if cur.Children[idx] == nil {
			return false
		}

		cur = cur.Children[idx]
	}

	return cur.IsWord
}

func getCharIndex(character rune) (idx int) {
	idx = int(unicode.ToLower(character) - 'a')
	return
}

func createNode(character rune, isWord bool, children ...*Node) *Node {
	return &Node{}

}
