## 知识点
### 二叉树遍历
* 前序遍历：先访问根结点，再前序遍历左子树，再前序遍历右子树
* 中序遍历：先中序遍历左子树，再访问根结点，再中序遍历右子树
* 后序遍历：先后序遍历左子树，再访问根结点，再后序遍历右子树

注意点：
    * 以根访问顺序决定是什么遍历
    * 左子树都是优先右子树
* 前序递归
```
func preorderTraversal(root *TreeNode)  {
    if root==nil{
        return
    }
    // 先访问根再访问左右
    fmt.Println(root.Val)
    preorderTraversal(root.Left)
    preorderTraversal(root.Right)
}
```
* 前序非递归
```
// V3：通过非递归遍历
func preorderTraversal(root *TreeNode) []int {
    // 非递归
    if root == nil{
        return nil
    }
    result:=make([]int,0)
    stack:=make([]*TreeNode,0)

    for root!=nil || len(stack)!=0{
        for root !=nil{
            // 前序遍历，所以先保存结果
            result=append(result,root.Val)
            stack=append(stack,root)
            root=root.Left
        }
        // pop
        node:=stack[len(stack)-1]
        stack=stack[:len(stack)-1]
        root=node.Right
    }
    return result
}
```
* 中序非递归
```
// 思路：通过stack 保存已经访问的元素，用于原路返回
func inorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    if root == nil {
        return result
    }
    stack := make([]*TreeNode, 0)
    for len(stack) > 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left // 一直向左
        }
        // 弹出
        val := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, val.Val)
        root = val.Right
    }
    return result
}
```
* 后序非递归
```
func postorderTraversal(root *TreeNode) []int {
    // 通过lastVisit标识右子节点是否已经弹出
    if root == nil {
        return nil
    }
    result := make([]int, 0)
    stack := make([]*TreeNode, 0)
    var lastVisit *TreeNode
    for root != nil || len(stack) != 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        // 这里先看看，先不弹出
        node:= stack[len(stack)-1]
        // 根节点必须在右节点弹出之后，再弹出
        if node.Right == nil || node.Right == lastVisit {
            stack = stack[:len(stack)-1] // pop
            result = append(result, node.Val)
            // 标记当前这个节点已经弹出过
            lastVisit = node
        } else {
            root = node.Right
        }
    }
    return result
}
```

注意点
* 核心就是：根节点必须在右节点弹出之后，再弹出
### 分治法应用
### BFS层次应用
### 二叉搜索树应用
## 总结
