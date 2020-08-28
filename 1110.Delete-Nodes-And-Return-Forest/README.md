# [1110. Delete Nodes And Return Forest](https://leetcode.com/problems/delete-nodes-and-return-forest/)



## 题目

Given the `root` of a binary tree, each node in the tree has a distinct value.

After deleting all nodes with a value in `to_delete`, we are left with a forest (a disjoint union of trees).

Return the roots of the trees in the remaining forest. You may return the result in any order.

**Example 1**:

![https://assets.leetcode.com/uploads/2019/07/01/screen-shot-2019-07-01-at-53836-pm.png](https://assets.leetcode.com/uploads/2019/07/01/screen-shot-2019-07-01-at-53836-pm.png)

```
Input: root = [1,2,3,4,5,6,7], to_delete = [3,5]
Output: [[1,2,null,4],[6],[7]]
```

**Constraints**:

- The number of nodes in the given tree is at most `1000`.
- Each node has a distinct value between `1` and `1000`.
- `to_delete.length <= 1000`
- `to_delete` contains distinct values between `1` and `1000`.


## 题目大意

给出二叉树的根节点 root，树上每个节点都有一个不同的值。如果节点值在 to_delete 中出现，我们就把该节点从树上删去，最后得到一个森林（一些不相交的树构成的集合）。返回森林中的每棵树。你可以按任意顺序组织答案。


提示：

- 树中的节点数最大为 1000。
- 每个节点都有一个介于 1 到 1000 之间的值，且各不相同。
- to_delete.length <= 1000
- to_delete 包含一些从 1 到 1000、各不相同的值。



## 解题思路

- 给出一棵树，再给出一个数组，要求删除数组中相同元素值的节点。输出最终删除以后的森林。
- 简单题。边遍历树，边删除数组中的元素。这里可以先把数组里面的元素放入 map 中，加速查找。遇到相同的元素就删除节点。这里需要特殊判断的是当前删除的节点是否是根节点，如果是根节点需要根据条件置空它的左节点或者右节点。

## 代码

```go
func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	if root == nil {
		return nil
	}
	res, deleteMap := []*TreeNode{}, map[int]bool{}
	for _, v := range toDelete {
		deleteMap[v] = true
	}
	dfsDelNodes(root, deleteMap, true, &res)
	return res
}

func dfsDelNodes(root *TreeNode, toDel map[int]bool, isRoot bool, res *[]*TreeNode) bool {
	if root == nil {
		return false
	}
	if isRoot && !toDel[root.Val] {
		*res = append(*res, root)
	}
	isRoot = false
	if toDel[root.Val] {
		isRoot = true
	}
	if dfsDelNodes(root.Left, toDel, isRoot, res) {
		root.Left = nil
	}
	if dfsDelNodes(root.Right, toDel, isRoot, res) {
		root.Right = nil
	}
	return isRoot
}
```