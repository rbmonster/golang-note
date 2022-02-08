package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var inorderMap = make(map[int]int)

func buildTree(preorder []int, inorder []int) *TreeNode {
	for i, val := range inorder {
		inorderMap[val] = i
	}
	len := len(preorder)
	return buildHelper(preorder, 0, len-1, 0, len-1)
}

func buildHelper(preorder []int, pl int, pr int, il int, ir int) *TreeNode {
	if pl > pr {
		return nil
	}
	root := TreeNode{Val: preorder[pl]}
	index := inorderMap[preorder[pl]]
	len := index - il
	root.Left = buildHelper(preorder, pl+1, pl+len, il, index-1)
	root.Right = buildHelper(preorder, pl+len+1, pr, index+1, ir)
	return &root
}
