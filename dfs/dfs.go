package main

import (
	"fmt"
)

type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }


func maxPathSum(root *TreeNode) int {

    maxPath := root.Val

     max := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }
	
    var dfs func (r *TreeNode) int 
	dfs = func (r *TreeNode) int {
		if r == nil {
			return 0
        }
		
        
         leftMax := dfs(r.Left)
         rightMax := dfs(r.Right)

         leftMax = max(0, leftMax)
         rightMax = max(0, rightMax)
         fmt.Println(leftMax)
         
         pathSplit := leftMax + rightMax + r.Val
         

         if pathSplit > maxPath {
            maxPath = pathSplit
         }
        
         if leftMax > rightMax {
            return r.Val + leftMax
         } else {
            return r.Val + rightMax
         }
        
    }

    dfs(root)
    return maxPath
}


func main () {

    root2 := &TreeNode{
        Val: 1,
        Left: &TreeNode{
            Val: 10,
            Left: &TreeNode{Val: 20},
            Right: &TreeNode{Val: 30},
        },
        Right: &TreeNode{
            Val: 100,
            Left: &TreeNode{Val: 200},
            Right: &TreeNode{Val: 300},
        },
    }
    
	fmt.Println(maxPathSum(root2))
}