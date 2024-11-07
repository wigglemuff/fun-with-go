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

// Maximum Depth of Binary Tree
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return max(leftDepth, rightDepth) + 1
}

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

// Main function to test each problem
func main() {
	// Build a sample tree
	root := buildSampleTree()

	// Test Maximum Depth
	fmt.Println("Maximum Depth of Binary Tree:", maxDepth(root))

	// Test Symmetric Tree
	symmetricRoot := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
		},
	}
	fmt.Println("Symmetric Tree:", isSymmetric(symmetricRoot))

	// Test Invert Binary Tree
	fmt.Println("Inverted Binary Tree (Preorder Traversal):")
	invertTree(root)
	printPreorder(root) // Helper function to print the tree in preorder
	fmt.Println()

	// Test Path Sum
	fmt.Println("Has Path Sum (target 7):", hasPathSum(root, 7))

	// Test Binary Tree Paths
	fmt.Println("Binary Tree Paths:", binaryTreePaths(root))
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
