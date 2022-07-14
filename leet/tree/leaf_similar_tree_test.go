package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("leaf_similar_tree", func() {
	It("ShouldBeSimilar", func() {
		tree1 := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		}
		tree2 := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
		}
		similarFlag := LeafSimilar(tree1, tree2)
		Ω(similarFlag).Should(BeTrue())
	})
})
