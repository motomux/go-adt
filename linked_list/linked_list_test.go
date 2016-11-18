package linkedlist

import (
	"reflect"
	"testing"
)

func TestNewSinglyLinkedList(t *testing.T) {
	tests := map[string]struct {
		out *SinglyLinkedList
	}{
		"case1": {
			out: &SinglyLinkedList{size: 0, head: nil},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := NewSinglyLinkedList()
			if !reflect.DeepEqual(out, test.out) {
				t.Errorf("actual: %+v, expected: %+v", out, test.out)
			}
		})
	}
}

func TestSinglyLinkedListAdd(t *testing.T) {
	tests := map[string]struct {
		pre  *SinglyLinkedList
		in   int
		post *SinglyLinkedList
	}{
		"case1": {
			pre:  &SinglyLinkedList{size: 0, head: nil},
			in:   1,
			post: &SinglyLinkedList{size: 1, head: &Node{Value: 1, Next: nil}},
		},
		"case2": {
			pre:  &SinglyLinkedList{size: 1, head: &Node{Value: 1, Next: nil}},
			in:   2,
			post: &SinglyLinkedList{size: 2, head: &Node{Value: 2, Next: &Node{Value: 1, Next: nil}}},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			pre, in, post := test.pre, test.in, test.post
			pre.Add(in)
			if !reflect.DeepEqual(pre, post) {
				t.Errorf("actual: %+v, expected: %+v", pre, post)
			}
		})
	}
}

func TestSinglyLinkedListAddLast(t *testing.T) {
	tests := map[string]struct {
		pre  *SinglyLinkedList
		in   int
		post *SinglyLinkedList
	}{
		"case1": {
			pre:  &SinglyLinkedList{size: 0, head: nil},
			in:   1,
			post: &SinglyLinkedList{size: 1, head: &Node{Value: 1, Next: nil}},
		},
		"case2": {
			pre:  &SinglyLinkedList{size: 1, head: &Node{Value: 1, Next: nil}},
			in:   2,
			post: &SinglyLinkedList{size: 2, head: &Node{Value: 1, Next: &Node{Value: 2, Next: nil}}},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			pre, in, post := test.pre, test.in, test.post
			pre.AddLast(in)
			if !reflect.DeepEqual(pre, post) {
				t.Errorf("actual: %+v, expected: %+v, diff: %v", pre, post)
			}
		})
	}
}
