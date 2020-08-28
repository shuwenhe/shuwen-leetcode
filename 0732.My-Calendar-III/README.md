# [732. My Calendar III](https://leetcode.com/problems/my-calendar-iii/)


## 题目

Implement a `MyCalendarThree` class to store your events. A new event can **always** be added.

Your class will have one method, `book(int start, int end)`. Formally, this represents a booking on the half open interval `[start, end)`, the range of real numbers `x` such that `start <= x < end`.

A K-booking happens when **K** events have some non-empty intersection (ie., there is some time that is common to all K events.)

For each call to the method `MyCalendar.book`, return an integer `K` representing the largest integer such that there exists a `K`-booking in the calendar.

Your class will be called like this:

`MyCalendarThree cal = new MyCalendarThree();`

`MyCalendarThree.book(start, end)`

**Example 1:**

    MyCalendarThree();
    MyCalendarThree.book(10, 20); // returns 1
    MyCalendarThree.book(50, 60); // returns 1
    MyCalendarThree.book(10, 40); // returns 2
    MyCalendarThree.book(5, 15); // returns 3
    MyCalendarThree.book(5, 10); // returns 3
    MyCalendarThree.book(25, 55); // returns 3
    Explanation: 
    The first two events can be booked and are disjoint, so the maximum K-booking is a 1-booking.
    The third event [10, 40) intersects the first event, and the maximum K-booking is a 2-booking.
    The remaining events cause the maximum K-booking to be only a 3-booking.
    Note that the last event locally causes a 2-booking, but the answer is still 3 because
    eg. [10, 20), [10, 40), and [5, 15) are still triple booked.

**Note:**

- The number of calls to `MyCalendarThree.book` per test case will be at most `400`.
- In calls to `MyCalendarThree.book(start, end)`, `start` and `end` are integers in the range `[0, 10^9]`.


## 题目大意

实现一个 MyCalendar 类来存放你的日程安排，你可以一直添加新的日程安排。

MyCalendar 有一个 book(int start, int end)方法。它意味着在start到end时间内增加一个日程安排，注意，这里的时间是半开区间，即 [start, end), 实数 x 的范围为，  start <= x < end。当 K 个日程安排有一些时间上的交叉时（例如K个日程安排都在同一时间内），就会产生 K 次预订。每次调用 MyCalendar.book方法时，返回一个整数 K ，表示最大的 K 次预订。

请按照以下步骤调用MyCalendar 类: MyCalendar cal = new MyCalendar(); MyCalendar.book(start, end)

说明:

- 每个测试用例，调用 MyCalendar.book 函数最多不超过 400 次。
- 调用函数 MyCalendar.book(start, end)时， start 和 end 的取值范围为 [0, 10^9]。




## 解题思路

- 设计一个日程类，每添加一个日程，实时显示出当前排期中累计日程最多的个数，例如在一段时间内，排了 3 个日程，其他时间内都只有 0，1，2 个日程，则输出 3 。
- 拿到这个题目以后会立即想到线段树。由于题目中只有增加日程，所以这一题难度不大。这一题和第 699 题也类似，但是有区别，第 699 题中，俄罗斯方块会依次摞起来，而这一题中，俄罗斯方块也就摞起来，但是方块下面如果是空挡，方块会断掉。举个例子：依次增加区间 [10,20]，[10,40]，[5,15]，[5,10]，如果是第 699 题的规则，这 [5,10] 的这块砖块会落在 [5,15] 上，从而使得高度为 4，但是这一题是日程，日程不一样，[5,15] 这个区间内有 3 个日程，但是其他部分都没有 3 个日程，所以第三块砖块 [5,15] 中的 [5,10] 会“断裂”，掉下去，第四块砖块还是 [5,10]，落在第三块砖块断落下去的位置，它们俩落在一起的高度是 2 。
- 构造一颗线段树，这里用树来构造，如果用数组需要开辟很大的空间。当区间左右边界和查询边界完全相同的时候再累加技术，否则不加，继续划分区间。以区间的左边界作为划分区间的标准，因为区间左边界是开区间，右边是闭区间。一个区间的计数值以区间左边界的计数为准。还是上面的例子，[5,10) 计数以 5 为标准，count = 2，[10,15) 计数以 10 为标准，count = 3 。还需要再动态维护一个最大值。这个线段树的实现比较简单。
- 类似的题目有：第 715 题，第 218 题，第 699 题。第 715 题是区间更新定值(**不是增减**)，第 218 题可以用扫描线，第 732 题和第 699 题类似，也是俄罗斯方块的题目，但是第 732 题的俄罗斯方块的方块会“断裂”。
