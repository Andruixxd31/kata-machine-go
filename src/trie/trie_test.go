package trie_test

import (
	"reflect"
	"testing"

	"github.com/andruixxd31/kata-machine-go/src/trie"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name  string
		given string
		trie  *trie.Trie
		want  *trie.Trie
	}{
		{
			name:  "Add word to empty trie",
			given: "cat",
			trie:  &trie.Trie{Root: node(0, false)}, // empty root
			want: &trie.Trie{
				Root: node(0, false,
					node('c', false,
						node('a', false,
							node('t', true),
						),
					),
				),
			},
		},
		{
			name:  "Add similar word when prefix already exists",
			given: "car",
			trie: &trie.Trie{
				Root: node(0, false,
					node('c', false,
						node('a', false,
							node('t', true),
						),
					),
				),
			},
			want: &trie.Trie{
				Root: node(0, false,
					node('c', false,
						node('a', false,
							node('t', true),
							node('r', true),
						),
					),
				),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.trie.Add(tt.given) // assumes you implemented Add
			if !reflect.DeepEqual(tt.trie, tt.want) {
				t.Errorf("got %+v, want %+v", tt.trie, tt.want)
			}
		})
	}
}

func TestSearch(t *testing.T) {

	tests := []struct {
		name  string
		given string
		trie  *trie.Trie
		want  bool
	}{
		{
			name:  "Word in trie",
			given: "cat",
			trie: &trie.Trie{
				Root: node(0, false,
					node('c', false,
						node('a', false,
							node('t', true),
							node('r', true),
						),
					),
				),
			},
			want: true,
		},
		{
			name:  "Word not in trie",
			given: "carbon",
			trie: &trie.Trie{
				Root: node(0, false,
					node('c', false,
						node('a', false,
							node('t', true),
							node('r', true),
						),
					),
				),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.trie.Search(tt.given)
			if got != tt.want {
				t.Errorf("got %+v, want %+v", tt.trie, tt.want)
			}
		})
	}
}

// helper to create a node
func node(ch rune, isWord bool, children ...*trie.Node) *trie.Node {
	arr := [26]*trie.Node{}
	for _, c := range children {
		arr[c.Character-'a'] = c
	}
	return &trie.Node{
		Character: ch,
		IsWord:    isWord,
		Children:  arr,
	}
}
