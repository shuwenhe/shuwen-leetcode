# [1105. Filling Bookcase Shelves](https://leetcode.com/problems/filling-bookcase-shelves/)


## 题目

We have a sequence of `books`: the `i`-th book has thickness `books[i][0]`and height `books[i][1]`.

We want to place these books **in order** onto bookcase shelves that have total width `shelf_width`.

We choose some of the books to place on this shelf (such that the sum of their thickness is `<= shelf_width`), then build another level of shelf of the bookcase so that the total height of the bookcase has increased by the maximum height of the books we just put down. We repeat this process until there are no more books to place.

Note again that at each step of the above process, the order of the books we place is the same order as the given sequence of books. For example, if we have an ordered list of 5 books, we might place the first and second book onto the first shelf, the third book on the second shelf, and the fourth and fifth book on the last shelf.

Return the minimum possible height that the total bookshelf can be after placing shelves in this manner.

**Example 1:**

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/06/28/shelves.png)

    Input: books = [[1,1],[2,3],[2,3],[1,1],[1,1],[1,1],[1,2]], shelf_width = 4
    Output: 6
    Explanation:
    The sum of the heights of the 3 shelves are 1 + 3 + 2 = 6.
    Notice that book number 2 does not have to be on the first shelf.

**Constraints:**

- `1 <= books.length <= 1000`
- `1 <= books[i][0] <= shelf_width <= 1000`
- `1 <= books[i][1] <= 1000`

## 题目大意

附近的家居城促销，你买回了一直心仪的可调节书架，打算把自己的书都整理到新的书架上。你把要摆放的书 books 都整理好，叠成一摞：从上往下，第 i 本书的厚度为 books[i][0]，高度为 books[i][1]。按顺序将这些书摆放到总宽度为 shelf\_width 的书架上。

先选几本书放在书架上（它们的厚度之和小于等于书架的宽度 shelf_width），然后再建一层书架。重复这个过程，直到把所有的书都放在书架上。

需要注意的是，在上述过程的每个步骤中，摆放书的顺序与你整理好的顺序相同。 例如，如果这里有 5 本书，那么可能的一种摆放情况是：第一和第二本书放在第一层书架上，第三本书放在第二层书架上，第四和第五本书放在最后一层书架上。每一层所摆放的书的最大高度就是这一层书架的层高，书架整体的高度为各层高之和。以这种方式布置书架，返回书架整体可能的最小高度。



## 解题思路

- 给出一个数组，数组里面每个元素代表的是一个本书的宽度和高度。要求按照书籍的顺序，把书摆到宽度为 `shelf_width` 的书架上。问最终放下所有书籍以后，书架的最小高度。
- 这一题的解题思路是动态规划。`dp[i]` 代表放置前 `i` 本书所需要的书架最小高度。初始值 dp[0] = 0，其他为最大值 1000*1000。遍历每一本书，把当前这本书作为书架最后一层的最后一本书，将这本书之前的书向后调整，看看是否可以减少之前的书架高度。状态转移方程为 `dp[i] = min(dp[i] , dp[j - 1] + h)`，其中 `j` 表示最后一层所能容下书籍的索引，`h` 表示最后一层最大高度。`j` 调整完一遍以后就能找出书架最小高度值了。时间复杂度 O(n^2)。