package tree_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"leet/tree"
)

var _ = Describe("leaf_similar_tree", func() {
	It("ShouldBeSimilar", func() {
		tree1 := &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 5,
				Left: &tree.TreeNode{
					Val: 6,
				},
				Right: &tree.TreeNode{
					Val: 2,
					Left: &tree.TreeNode{
						Val: 7,
					},
					Right: &tree.TreeNode{
						Val: 4,
					},
				},
			},
			Right: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 9,
				},
				Right: &tree.TreeNode{
					Val: 8,
				},
			},
		}
		tree2 := &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 5,
				Left: &tree.TreeNode{
					Val: 6,
				},
				Right: &tree.TreeNode{
					Val: 7,
				},
			},
			Right: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 4,
				},
				Right: &tree.TreeNode{
					Val: 2,
					Left: &tree.TreeNode{
						Val: 9,
					},
					Right: &tree.TreeNode{
						Val: 8,
					},
				},
			},
		}
		similarFlag := tree.LeafSimilar(tree1, tree2)
		Ω(similarFlag).Should(BeTrue())
	})
})
