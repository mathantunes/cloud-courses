package leet

// Binary Tree Preorder Traversal
// https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/928/
// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Root -> Left -> Right
func preorderTraversal(root *TreeNode) []int {
	var out []int
	preOrderAppend(root, &out)
	return out
}

func preOrderAppend(node *TreeNode, arr *[]int) {
	if node != nil {
		*arr = append(*arr, node.Val)
		preOrderAppend(node.Left, arr)
		preOrderAppend(node.Right, arr)
	}
}

// Left -> Root -> Right
func inorderTraversal(root *TreeNode) []int {
	var out []int
	inOrderAppend(root, &out)
	return out
}

func inOrderAppend(node *TreeNode, arr *[]int) {
	if node != nil {
		inOrderAppend(node.Left, arr)
		*arr = append(*arr, node.Val)
		inOrderAppend(node.Right, arr)
	}
}

func postorderTraversal(root *TreeNode) []int {
	var out []int
	postOrderAppend(root, &out)
	return out
}

// Left -> Right -> Root
func postOrderAppend(node *TreeNode, arr *[]int) {
	if node != nil {
		postOrderAppend(node.Left, arr)
		postOrderAppend(node.Right, arr)
		*arr = append(*arr, node.Val)
	}
}

// Binary Tree Level Order Traversal
// https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/931/
func levelOrder(root *TreeNode) [][]int {
	var out [][]int
	appendPerLevel(root, &out, 0)
	return out
}

func appendPerLevel(root *TreeNode, out *[][]int, level int) {
	if root != nil {
		if len(*out) > level {
			(*out)[level] = append((*out)[level], root.Val)
		} else {
			*out = append(*out, []int{root.Val})
		}
		appendPerLevel(root.Left, out, level+1)
		appendPerLevel(root.Right, out, level+1)
	}
}

// Maximum Depth of Binary Tree
// https://leetcode.com/explore/learn/card/data-structure-tree/17/solve-problems-recursively/535/
func maxDepth(root *TreeNode) int {
	return findDepth(root, 0)
}

func findDepth(root *TreeNode, currentDepth int) int {
	if root != nil {
		return max(findDepth(root.Left, currentDepth+1), findDepth(root.Right, currentDepth+1))
	}
	return currentDepth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Symmetric Tree
// https://leetcode.com/explore/learn/card/data-structure-tree/17/solve-problems-recursively/536/
func isSymmetric(root *TreeNode) bool {
	if root != nil {
		return isSymmetricRec(root.Left, root.Right)
	}
	return true
}

func isSymmetricRec(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == nil && right == nil
	}
	if left.Val == right.Val {
		return isSymmetricRec(left.Left, right.Right) && isSymmetricRec(left.Right, right.Left)
	}
	return false
}

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// Populating Next Right Pointers in Each Node
// https://leetcode.com/explore/learn/card/data-structure-tree/133/conclusion/994/
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	root.Next = nil
	connectChildren(root.Left, root.Right)
	return root
}

func connectChildren(left, right *Node) {
	if left == nil && right == nil {
		return
	}
	left.Next = right
	connectChildren(left.Left, left.Right)
	connectChildren(left.Right, right.Left)
	connectChildren(right.Left, right.Right)
}

// Sort an Array
// Merge sort
// https://leetcode.com/explore/learn/card/recursion-ii/470/divide-and-conquer/2944/
func sortArray(nums []int) []int {
	numsLen := len(nums)
	if numsLen <= 1 {
		return nums
	}
	pivot := numsLen / 2
	return merge(sortArray(nums[:pivot]), sortArray(nums[pivot:]))
}

func merge(left, right []int) []int {
	leftLen := len(left)
	rightLen := len(right)
	out := make([]int, leftLen+rightLen)
	var leftPtr, rightPtr, mergedPtr = 0, 0, 0

	for leftPtr < leftLen && rightPtr < rightLen {
		if left[leftPtr] > right[rightPtr] {
			out[mergedPtr] = right[rightPtr]
			rightPtr++
		} else {
			out[mergedPtr] = left[leftPtr]
			leftPtr++
		}
		mergedPtr++
	}

	for leftPtr < leftLen {
		out[mergedPtr] = left[leftPtr]
		mergedPtr++
		leftPtr++
	}
	for rightPtr < rightLen {
		out[mergedPtr] = right[rightPtr]
		mergedPtr++
		rightPtr++
	}
	return out
}

// Validate Binary Search Tree
// https://leetcode.com/explore/learn/card/recursion-ii/470/divide-and-conquer/2874/
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isValid(root, nil, nil)
}

func isValid(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return isValid(root.Left, min, root) && isValid(root.Right, root, max)
}

// Search a 2D Matrix II
// https://leetcode.com/explore/learn/card/recursion-ii/470/divide-and-conquer/2872/
func searchMatrix(matrix [][]int, target int) bool {
	var m, n = len(matrix), len(matrix[0])
	var r, c = 0, n - 1
	for r < m && c >= 0 {
		if target > matrix[r][c] {
			r += 1
		} else if target < matrix[r][c] {
			c -= 1
		} else {
			return true
		}
	}
	return false
}
