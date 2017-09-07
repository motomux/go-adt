package trietree

// Node describes tri tree node
type Node struct {
	children map[byte]*Node
	isEnd    bool
}

// NewNode initiates Node for TrieTree
func NewNode() *Node {
	return &Node{
		children: make(map[byte]*Node),
	}
}

// TrieTree describes tri tree
type TrieTree struct {
	root *Node
}

// New initiates TrieTree
func New(words []string) *TrieTree {
	tree := &TrieTree{
		root: NewNode(),
	}

	for _, word := range words {
		tree.Add(word)
	}

	return tree
}

// Add adds word into the tri tree
func (tree *TrieTree) Add(word string) {
	node := tree.root
	for i := 0; i < len(word); i++ {
		char := word[i]

		next := node.children[char]
		if next == nil {
			next = NewNode()
			node.children[char] = next
		}

		if i == len(word)-1 {
			next.isEnd = true
		}
		node = next
	}
}

// Complete returns words which starts with the give prefix
func (tree *TrieTree) Complete(prefix string) []string {
	var words []string

	node := tree.root
	for i := 0; i < len(prefix); i++ {
		char := prefix[i]
		node = node.children[char]
		if node == nil {
			return words
		}
	}

	tree.getWords(node, prefix, &words)
	return words
}

func (tree *TrieTree) getWords(node *Node, prefix string, words *[]string) {
	if node == nil {
		return
	}
	if node.isEnd {
		*words = append(*words, prefix)
	}

	for k, cNode := range node.children {
		tree.getWords(cNode, prefix+string(k), words)
	}
}
