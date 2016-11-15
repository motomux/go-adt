package linkedlist

import (
	"reflect"
	"testing"
)

func TestNewNode(t *testing.T) {
	type (
		in struct {
			value int
			next  *Node
		}
	)
	tests := map[string]struct {
		in
		out *Node
	}{
		"case1": {
			in{1, &Node{2, nil}},
			&Node{1, &Node{2, nil}},
		},
		"case2": {
			in{1, nil},
			&Node{1, nil},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			in := test.in

			out := NewNode(in.value, in.next)

			if reflect.DeepEqual(out, test.out) == false {
				t.Errorf("actual: %+v, expected: %+v\n", out, test.out)
			}
		})
	}
}

func TestNodeHasNext(t *testing.T) {
	tests := map[string]struct {
		node *Node
		out  bool
	}{
		"case1": {
			NewNode(1, NewNode(2, nil)),
			true,
		},
		"case2": {
			NewNode(1, nil),
			false,
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			node := test.node
			if out := node.HasNext(); out != test.out {
				t.Errorf("actual: %v, expected: %v\n", out, test.out)
			}
		})
	}
}
