package trietree

import (
	"reflect"
	"sort"
	"strconv"
	"testing"

	"github.com/motomux/pretty"
)

func TestNew(t *testing.T) {
	tests := []struct {
		in  []string
		out *TrieTree
	}{
		{
			in: []string{},
			out: &TrieTree{
				root: &Node{children: map[byte]*Node{}},
			},
		},
		{
			in: []string{"a"},
			out: &TrieTree{
				root: &Node{children: map[byte]*Node{
					'a': &Node{isEnd: true, children: map[byte]*Node{}},
				}},
			},
		},
		{
			in: []string{"a", "ab"},
			out: &TrieTree{
				root: &Node{children: map[byte]*Node{
					'a': &Node{isEnd: true, children: map[byte]*Node{
						'b': &Node{isEnd: true, children: map[byte]*Node{}},
					}},
				}},
			},
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			out := New(test.in)
			if !reflect.DeepEqual(out, test.out) {
				t.Errorf("diff=%v", pretty.Diff(out, test.out))
			}
		})
	}
}

func TestComplete(t *testing.T) {
	tests := []struct {
		tree *TrieTree
		in   string
		out  []string
	}{
		{
			tree: New([]string{
				"bad", "badger", "badminton", "badegg", "battleship", "bed", "beggar", "cat", "caterpillar", "catering",
			}),
			in:  "ba",
			out: []string{"bad", "badger", "badminton", "badegg", "battleship"},
		},
		{
			tree: New([]string{
				"bad", "badger", "badminton", "badegg", "battleship", "bed", "beggar", "cat", "caterpillar", "catering",
			}),
			in:  "bad",
			out: []string{"bad", "badger", "badminton", "badegg"},
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			out := test.tree.Complete(test.in)

			// Order can be changed since map is unsorted
			sort.Strings(out)
			sort.Strings(test.out)
			if !reflect.DeepEqual(out, test.out) {
				t.Errorf("actual=%v expected=%v", out, test.out)
			}
		})
	}
}
