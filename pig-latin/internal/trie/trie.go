package trie

import (
	"unicode/utf8"
)

// Rule defines a rule for pig latin.
// If a word matches the given prefix, then n runes should be moved
// from the beginning of the word to the end, before appending 'ay'
// Longer prefixes will override shorter ones.
type Rule struct {
	prefix string
	n      int
}

// a Trie is a tree structure which allows us to store and evaluate rules
type Trie interface {
	// Insert is used when preparing the Trie for use.
	Insert(r Rule)
	// Evaluate uses the previously inserted prefixes to determine
	// the rule that best matches the given word.
	// If none of the rules match, then okay will be false.
	Evaluate(word string) (r Rule, okay bool)
}

// compile-time interface check
var _ Trie = &trie{}

type trie struct {
	parent   *trie
	children map[rune]*trie
	prefix   rune
	hasValue bool
	move     int
}

// New initializes an empty Trie, ready to insert Rules
func New() Trie {
	return &trie{
		children: make(map[rune]*trie),
	}
}

func (t *trie) Insert(r Rule) {
	t.insert([]rune(r.prefix), r.n)
}

func (t *trie) insert(prefix []rune, move int) {
	// base case: no more input
	if len(prefix) == 0 {
		t.move = move
		t.hasValue = true
		return
	}

	r := prefix[0]
	child, ok := t.children[r]
	if !ok {
		// add a new child node if we don't have a matching one yet
		child = &trie{
			prefix:   r,
			parent:   t,
			children: make(map[rune]*trie),
			move:     move,
		}
		t.children[r] = child
	}

	// recursive case
	child.insert(prefix[1:], move)
}

func (t *trie) Evaluate(word string) (match Rule, ok bool) {
	return t.evaluate([]rune(word))
}

func (t *trie) evaluate(prefix []rune) (match Rule, ok bool) {
	// base case - no more input
	if len(prefix) == 0 {
		return t.rule()
	}

	r := prefix[0]
	child, ok := t.children[r]
	if !ok {
		// base case - we do not have a node for the next rune
		return t.rule()
	}

	// recursive case - get the best rule from our children
	match, ok = child.evaluate(prefix[1:])
	if ok {
		return
	}

	// base case - none of our children match
	return t.rule()
}

// word returns the sequence of bytes that led to this node
func (t *trie) word() []byte {
	if t.parent == nil {
		return nil
	}
	size := utf8.RuneLen(t.prefix)
	bytes := make([]byte, size)
	utf8.EncodeRune(bytes, t.prefix)
	return append(t.parent.word(), bytes...)
}

// rule returns the rule that was inserted to this node, if there is one
func (t *trie) rule() (match Rule, ok bool) {
	if t.hasValue {
		return Rule{string(t.word()), t.move}, true
	}
	return Rule{}, false
}
