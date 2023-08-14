package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieSearch(t *testing.T) {
	tr := NewTrie()

	tr.Insert("hello")
	tr.Insert("hallo")

	found := tr.Search("hello")
	assert.True(t, found)

	found = tr.Search("hallo")
	assert.True(t, found)

	found = tr.Search("hal")
	assert.False(t, found)

	found = tr.Search("helio")
	assert.False(t, found)

	found = tr.Search("other")
	assert.False(t, found)
}

func TestTrieStartsWith(t *testing.T) {
	tr := NewTrie()

	tr.Insert("hello")
	tr.Insert("hallo")

	found := tr.StartsWith("hello")
	assert.True(t, found)

	found = tr.StartsWith("hal")
	assert.True(t, found)

	found = tr.StartsWith("hyl")
	assert.False(t, found)
}

func TestTrieInsert(t *testing.T) {
	tr := NewTrie()

	tr.Insert("hello")
	tr.Insert("hallo")

	_, hasH := tr.root.next["h"]
	assert.Len(t, tr.root.next, 1)
	assert.True(t, hasH)

	_, hasA := tr.root.next["h"].next["a"]
	_, hasE := tr.root.next["h"].next["e"]
	assert.Len(t, tr.root.next["h"].next, 2)
	assert.True(t, hasA)
	assert.True(t, hasE)

	// @TODO - etc...
}

func TestTrieWordEnd(t *testing.T) {
	tr := NewTrie()

	tr.Insert("dog")
	tr.Insert("dogs")

	d := tr.root.next["d"]
	assert.False(t, d.isWord)

	do := d.next["o"]
	assert.False(t, do.isWord)

	dog := do.next["g"]
	assert.True(t, dog.isWord)

	dogs := dog.next["s"]
	assert.True(t, dogs.isWord)
}
