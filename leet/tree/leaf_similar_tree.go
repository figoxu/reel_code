package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
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
