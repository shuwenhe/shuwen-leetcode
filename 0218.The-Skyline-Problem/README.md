# [218. The Skyline Problem](https://leetcode.com/problems/the-skyline-problem/)

## 题目

A city's skyline is the outer contour of the silhouette formed by all the buildings in that city when viewed from a distance. Now suppose you are **given the locations and height of all the buildings** as shown on a cityscape photo (Figure A), write a program to **output the skyline** formed by these buildings collectively (Figure B).

![](https://img.halfrost.com/Leetcode/leetcode_218_0.png)

![](https://img.halfrost.com/Leetcode/leetcode_218_1.png)

The geometric information of each building is represented by a triplet of integers `[Li, Ri, Hi]`, where `Li` and `Ri` are the x coordinates of the left and right edge of the ith building, respectively, and `Hi` is its height. It is guaranteed that `0 ≤ Li, Ri ≤ INT_MAX`, `0 < Hi ≤ INT_MAX`, and `Ri - Li > 0`. You may assume all buildings are perfect rectangles grounded on an absolutely flat surface at height 0.

For instance, the dimensions of all buildings in Figure A are recorded as: `[ [2 9 10], [3 7 15], [5 12 12], [15 20 10], [19 24 8] ]` .

The output is a list of "**key points**" (red dots in Figure B) in the format of `[ [x1,y1], [x2, y2], [x3, y3], ... ]` that uniquely defines a skyline. **A key point is the left endpoint of a horizontal line segment**. Note that the last key point, where the rightmost building ends, is merely used to mark the termination of the skyline, and always has zero height. Also, the ground in between any two adjacent buildings should be considered part of the skyline contour.

For instance, the skyline in Figure B should be represented as:`[ [2 10], [3 15], [7 12], [12 0], [15 10], [20 8], [24, 0] ]`.

**Notes:**

- The number of buildings in any input list is guaranteed to be in the range `[0, 10000]`.
- The input list is already sorted in ascending order by the left x position `Li`.
- The output list must be sorted by the x position.
- There must be no consecutive horizontal lines of equal height in the output skyline. For instance, `[...[2 3], [4 5], [7 5], [11 5], [12 7]...]` is not acceptable; the three lines of height 5 should be merged into one in the final output as such: `[...[2 3], [4 5], [12 7], ...]`


## 题目大意

城市的天际线是从远处观看该城市中所有建筑物形成的轮廓的外部轮廓。现在，假设您获得了城市风光照片（图A）上显示的所有建筑物的位置和高度，请编写一个程序以输出由这些建筑物形成的天际线（图B）。

每个建筑物的几何信息用三元组 [Li，Ri，Hi] 表示，其中 Li 和 Ri 分别是第 i 座建筑物左右边缘的 x 坐标，Hi 是其高度。可以保证 0 ≤ Li, Ri ≤ INT\_MAX, 0 < Hi ≤ INT\_MAX 和 Ri - Li > 0。您可以假设所有建筑物都是在绝对平坦且高度为 0 的表面上的完美矩形。

例如，图 A 中所有建筑物的尺寸记录为：[ [2 9 10], [3 7 15], [5 12 12], [15 20 10], [19 24 8] ] 。

输出是以 [ [x1,y1], [x2, y2], [x3, y3], ... ] 格式的“关键点”（图 B 中的红点）的列表，它们唯一地定义了天际线。关键点是水平线段的左端点。请注意，最右侧建筑物的最后一个关键点仅用于标记天际线的终点，并始终为零高度。此外，任何两个相邻建筑物之间的地面都应被视为天际线轮廓的一部分。

例如，图 B 中的天际线应该表示为：[ [2 10], [3 15], [7 12], [12 0], [15 10], [20 8], [24, 0] ]。

说明:

- 任何输入列表中的建筑物数量保证在 [0, 10000] 范围内。
- 输入列表已经按左 x 坐标 Li 进行升序排列。
- 输出列表必须按 x 位排序。
- 输出天际线中不得有连续的相同高度的水平线。例如 [...[2 3], [4 5], [7 5], [11 5], [12 7]...] 是不正确的答案；三条高度为 5 的线应该在最终输出中合并为一个：[...[2 3], [4 5], [12 7], ...]


## 解题思路


- 给出一个二维数组，每个子数组里面代表一个高楼的信息，一个高楼的信息包含 3 个信息，高楼起始坐标，高楼终止坐标，高楼高度。要求找到这些高楼的边际点，并输出这些边际点的高度信息。
- 这一题可以用线段树来解。用线段树来解答，可以不用关心“楼挡住楼”的情况。由于楼的坐标是离散的，所以先把楼在 X 轴上两个坐标离散化。同第 699 题一样，楼的宽度是一个区间，但是离散的过程中，楼的宽度右边界需要减一，不然查询一个区间会包含两个点，导致错误的结果，例如，第一个楼是 [1,3)，楼高 10，第二个楼是 [3,6)，楼高 20 。第一个楼如果算上右边界 3，查询 [1,3] 的结果是 20，因为 [3,3] 这个点会查询到第二个楼上面去。所以每个楼的右边界应该减一。但是每个楼的右边界也要加入到 query 中，因为最终 query 的结果需要包含这些边界。将离散的数据排序以后，按照楼的信息，每个区间依次 update。最后统计的时候依次统计每个区间，如果当前区间的高度和前一个区间的高度一样，就算是等高的楼。当高度与前一个高度不相同的时候就算是天际线的边缘，就要添加到最后输出数组中。
- 类似的线段树的题目有：第 715 题，第 732 题，第 699 题。第 715 题是区间更新定值(**不是增减**)，第 218 题可以用扫描线，第 732 题和第 699 题类似，也是俄罗斯方块的题目，但是第 732 题的俄罗斯方块的方块会“断裂”。
- 这一题用线段树做时间复杂度有点高，可以用扫描线解题。扫描线的思路很简单，用一根根垂直于 X 轴的竖线，从最左边依次扫到最右边，扫描每一条大楼的边界，当进入大楼的左边界的时候，如果没有比这个左边界最高点更高的点，就记录下这个最高点 keyPoint，状态是进入状态。如果扫到一个大楼的左边界，有比它更高的高度，就不记录，因为它不是天际线，它被楼挡楼，挡在其他楼后面了。当扫到一个大楼的右边界的时候，如果是最高点，那么记录下它的状态是离开状态，此时还需要记录下第二高的点。在扫描线扫描的过程中，动态的维护大楼的高度，同时维护最高的高度和第二高的高度。其实只需要维护最高的高度这一个高度，因为当离开状态到来的时候，移除掉当前最高的，剩下的高度里面最高的就是第二高的高度。描述的伪代码如下：

        // 扫描线伪代码
        events = {{x: L , height: H , type: entering},
        		  {x: R , height: H , type: leaving}}
        event.SortByX()
        ds = new DS()
        
        for e in events:
        	if entering(e):
        		if e.height > ds.max(): ans += [e.height]
        		ds.add(e.height)
        	if leaving(e):
        		ds.remove(e.height)
        		if e.height > ds.max(): ans += [ds.max()]

- 动态插入，查找最大值可以选用的数据结构有，最大堆和二叉搜索树。最大堆找最大值 O(1)，插入 O(log n)，但是 remove_by_key 需要 O(n) 的时间复杂度，并且需要自己实现。二叉搜索树，查找 max，添加和删除元素都是 O(log n) 的时间复杂度。
- 排序的时候也需要注意几个问题：如果大楼的边界相等，并且是进入状态，那么再按照高度从大到小进行排序；如果大楼的边界相等，并且是离开状态，那么高度按照从小到大进行排序。
