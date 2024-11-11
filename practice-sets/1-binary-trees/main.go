package main

import (
	"fmt"
	"strconv"
)

// TreeNode structure for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Helper function to build a sample tree
func buildSampleTree() *TreeNode {
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
}

// Helper function to build a sample symmetric tree
func buildSampleSymmetricTree() *TreeNode {
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
}

// Helper function to build tree for pathsum2 problem
func buildTreeForPathSum2() *TreeNode {
	return &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
}

// https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
// Maximum Depth of Binary Tree
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

// https://leetcode.com/problems/symmetric-tree/description/
// Symmetric Tree
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

// https://leetcode.com/problems/invert-binary-tree/description/
// Invert Binary Tree
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// https://leetcode.com/problems/path-sum/description/
// Path Sum
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	targetSum -= root.Val
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

// https://leetcode.com/problems/binary-tree-paths/description/
// Binary Tree Paths
func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	if root == nil {
		return paths
	}
	var path string
	findPaths(root, path, &paths)
	return paths
}

func findPaths(node *TreeNode, path string, paths *[]string) {
	if node == nil {
		return
	}
	path += strconv.Itoa(node.Val)
	if node.Left == nil && node.Right == nil { // Leaf node
		*paths = append(*paths, path)
	} else {
		path += "->"
		findPaths(node.Left, path, paths)
		findPaths(node.Right, path, paths)
	}
}

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
// Lowest Common Ancestor of a Binary Search Tree
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

// https://leetcode.com/problems/diameter-of-binary-tree/description/
// Diameter of Binary Tree
func diameterOfBinaryTree(root *TreeNode) int {
	_, diameter := depthAndDiameter(root)
	return diameter
}

func depthAndDiameter(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}

	leftDepth, leftDiameter := depthAndDiameter(node.Left)
	rightDepth, rightDiameter := depthAndDiameter(node.Right)

	return 1 + max(leftDepth, rightDepth), max(leftDiameter, rightDiameter, leftDepth+rightDepth)
}

// https://leetcode.com/problems/merge-two-binary-trees/description/
// Merge Two Binary Trees
func mergeTrees(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

// https://leetcode.com/problems/minimum-depth-of-binary-tree/
// Min depth
func minDepth(root *TreeNode) int {
	return -1 * minDepthHelper(root)
}

// Min depth helper
func minDepthHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(minDepthHelper(root.Left), minDepthHelper(root.Right)) - 1
}

// Min depth alt
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return min(minDepth2(root.Left), minDepth2(root.Right)) + 1
}

// https://leetcode.com/problems/range-sum-of-bst/description/
// Range Sum of BST
func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	sum := 0
	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}
	if root.Val > L {
		sum += rangeSumBST(root.Left, L, R)
	}
	if root.Val < R {
		sum += rangeSumBST(root.Right, L, R)
	}
	return sum
}

// Helper function to print the tree in preorder for visual confirmation
func printPreorder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Val, " ")
	printPreorder(node.Left)
	printPreorder(node.Right)
}

// Level order traversal of a binary tree
func levelOrderTraversal(root *TreeNode) {
	q := []*TreeNode{root}

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]

			fmt.Print(node.Val, " ")
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
}

// https://leetcode.com/problems/path-sum-ii/description/
// Path Sum 2
func pathSum2(root *TreeNode, targetSum int, tmp []int, res *[][]int) {
	if root == nil {
		return
	}
	tmp = append(tmp, root.Val)
	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			path := make([]int, len(tmp))
			copy(path, tmp)
			*res = append(*res, path)
			return
		}
	}
	newTarget := targetSum - root.Val
	pathSum2(root.Left, newTarget, tmp, res)
	pathSum2(root.Right, newTarget, tmp, res)
}

// Main function to test each problem
func main() {
	// Build a sample tree
	root := buildSampleTree()

	// Test Maximum Depth
	fmt.Println("Maximum Depth of Binary Tree:", maxDepth(root))

	// Test Symmetric Tree
	symmetricRoot := buildSampleSymmetricTree()
	fmt.Println("Symmetric Tree:", isSymmetric(symmetricRoot))

	// Test Invert Binary Tree
	fmt.Println("Inverted Binary Tree (Preorder Traversal):")
	invertTree(root)
	printPreorder(root)
	fmt.Println()

	fmt.Println("Binary Tree (Level Order Traversal):")
	levelOrderTraversal(buildSampleTree())
	fmt.Println()

	// Test Path Sum
	fmt.Println("Has Path Sum (target 10):", hasPathSum(root, 10))

	// Test Path Sum 2
	var pathSumRes [][]int
	pathSum2(buildTreeForPathSum2(), 22, []int{}, &pathSumRes)
	fmt.Println("PathSum2 result: ", pathSumRes)

	// Test Binary Tree Paths
	fmt.Println("Binary Tree Paths:", binaryTreePaths(root))

	// Test Lowest Common Ancestor of a BST
	p, q := &TreeNode{Val: 2}, &TreeNode{Val: 4}
	fmt.Println("Lowest Common Ancestor of 2 and 4:", lowestCommonAncestor(root, p, q).Val)

	// Test Diameter of Binary Tree
	fmt.Println("Diameter of Binary Tree:", diameterOfBinaryTree(root))

	// Test Merge Two Binary Trees
	mergedTree := mergeTrees(&TreeNode{Val: 1, Left: &TreeNode{Val: 3}}, &TreeNode{Val: 2, Right: &TreeNode{Val: 4}})
	fmt.Println("Merged Tree (Preorder Traversal):")
	printPreorder(mergedTree)
	fmt.Println()

	// Test Minimum Depth of Binary Tree
	fmt.Println("Minimum Depth of Binary Tree:", minDepth(root))

	// Test Minimum Depth of Binary Tree Alt
	fmt.Println("Minimum Depth2 of Binary Tree:", minDepth2(root))

	// Test Range Sum of BST
	fmt.Println("Range Sum of BST (Range 4 to 9):", rangeSumBST(root, 4, 9))

}
