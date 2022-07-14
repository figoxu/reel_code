package tree_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"leet/tree"
)

var _ = Describe("implement-trie-prefix-tree", func() {
	It("ShouldBeMatch", func() {
		trie := tree.NewTrie()
		word := "hello"
		trie.Insert(word)
		Ω(trie.Search(word)).Should(BeTrue())
		Ω(trie.StartsWith("he")).Should(BeTrue())
		Ω(trie.StartsWith("fine")).Should(BeFalse())
	})
})
