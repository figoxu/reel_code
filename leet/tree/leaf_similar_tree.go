package main

import "fmt"

//https://leetcode.cn/problems/leaf-similar-trees/
func main() {
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
	similarFlag := leafSimilar(tree1, tree2)
	fmt.Println(similarFlag)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	list1 := make([]int, 0)
	list2 := make([]int, 0)
	getLeaf(&list1, root1)
	getLeaf(&list2, root2)
	if len(list1) != len(list2) {
		return false
	}

	for i := 0; i < len(list1); i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

func getLeaf(list *[]int, root *TreeNode) {
	if root == nil {
		return
	}

	leafFlag := root.Left == nil && root.Right == nil
	if leafFlag {
		*list = append(*list, root.Val)
	} else {
		getLeaf(list, root.Left)
		getLeaf(list, root.Right)
	}
}
