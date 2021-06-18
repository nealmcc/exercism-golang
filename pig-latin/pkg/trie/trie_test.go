package trie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertEvaluate(t *testing.T) {
	tt := []struct {
		name     string
		rules    []Rule
		in       string
		wantOk   bool
		wantRule Rule
	}{
		{
			name:   "an empty trie does not match a word",
			rules:  nil,
			in:     "a",
			wantOk: false,
		},
		{
			name:   "an empty trie does not match an empty string",
			rules:  nil,
			in:     "",
			wantOk: false,
		},
		{
			name:     "match on a one-letter word",
			rules:    []Rule{{"a", 0}},
			in:       "a",
			wantOk:   true,
			wantRule: Rule{"a", 0},
		},
		{
			name:     "match on a subset of a longer rule",
			rules:    []Rule{{"abc", 2}, {"ab", 1}},
			in:       "ab",
			wantOk:   true,
			wantRule: Rule{"ab", 1},
		},
		{
			name:     "match on the longer prefix",
			rules:    []Rule{{"abc", 2}, {"ab", 1}},
			in:       "abc",
			wantOk:   true,
			wantRule: Rule{"abc", 2},
		},
		{
			name:     "our word is longer than the longest prefix",
			rules:    []Rule{{"abc", 2}, {"ab", 1}},
			in:       "abcd",
			wantOk:   true,
			wantRule: Rule{"abc", 2},
		},
		{
			name:     "our word has a match shorter than an almost-match",
			rules:    []Rule{{"abc", 2}, {"ab", 1}, {"abdq", 3}},
			in:       "abder",
			wantOk:   true,
			wantRule: Rule{"ab", 1},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			trie := New()

			for _, rule := range tc.rules {
				trie.Insert(rule)
			}

			got, ok := trie.Evaluate(tc.in)

			r.Equal(tc.wantOk, ok)
			if ok {
				r.Equal(tc.wantRule, got)
			}
		})
	}
}
