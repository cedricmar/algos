package ds

type Trie struct {
	root *node
}

type node struct {
	isWord bool
	next   map[string]*node
}

func NewTrie() Trie {
	return Trie{
		root: &node{
			next: map[string]*node{},
		},
	}
}

func (t *Trie) Insert(word string) {
	curr := t.root

	for i := 0; i < len(word); i++ {
		char := string(word[i])

		if _, ok := curr.next[char]; !ok {
			curr.next[char] = &node{next: map[string]*node{}}
		}
		curr = curr.next[char]
	}

	curr.isWord = true
}

func (t *Trie) Search(word string) bool {
	n := t.findLastNode(word)
	if n == nil {
		return false
	}
	return n.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	if n := t.findLastNode(prefix); n == nil {
		return false
	}
	return true
}

func (t *Trie) findLastNode(s string) *node {
	curr := t.root

	for i := 0; i < len(s); i++ {
		char := string(s[i])

		if _, ok := curr.next[char]; !ok {
			return nil
		}

		curr = curr.next[char]
	}

	return curr
}
