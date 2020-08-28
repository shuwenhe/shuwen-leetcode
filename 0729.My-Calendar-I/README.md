# [729. My Calendar I](https://leetcode.com/problems/my-calendar-i/)


## 题目

Implement a `MyCalendar` class to store your events. A new event can be added if adding the event will not cause a double booking.

Your class will have the method, `book(int start, int end)`. Formally, this represents a booking on the half open interval `[start, end)`, the range of real numbers `x` such that `start <= x < end`.

A double booking happens when two events have some non-empty intersection (ie., there is some time that is common to both events.)

For each call to the method `MyCalendar.book`, return `true` if the event can be added to the calendar successfully without causing a double booking. Otherwise, return `false` and do not add the event to the calendar.

Your class will be called like this:

`MyCalendar cal = new MyCalendar();`

`MyCalendar.book(start, end)`

**Example 1:**

    MyCalendar();
    MyCalendar.book(10, 20); // returns true
    MyCalendar.book(15, 25); // returns false
    MyCalendar.book(20, 30); // returns true
    Explanation: 
    The first event can be booked.  The second can't because time 15 is already booked by another event.
    The third event can be booked, as the first event takes every time less than 20, but not including 20.

**Note:**

- The number of calls to `MyCalendar.book` per test case will be at most `1000`.
- In calls to `MyCalendar.book(start, end)`, `start` and `end` are integers in the range `[0, 10^9]`.



## 题目大意

实现一个 MyCalendar 类来存放你的日程安排。如果要添加的时间内没有其他安排，则可以存储这个新的日程安排。

MyCalendar 有一个 book(int start, int end) 方法。它意味着在 start 到 end 时间内增加一个日程安排，注意，这里的时间是半开区间，即 [start, end), 实数 x 的范围为，  start <= x < end。

当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生重复预订。

每次调用 MyCalendar.book 方法时，如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true。否则，返回 false 并且不要将该日程安排添加到日历中。

请按照以下步骤调用 MyCalendar 类: MyCalendar cal = new MyCalendar(); MyCalendar.book(start, end)

说明:

- 每个测试用例，调用 MyCalendar.book 函数最多不超过 100次。
- 调用函数 MyCalendar.book(start, end) 时， start 和 end 的取值范围为 [0, 10^9]。


## 解题思路


- 要求实现一个日程安排的功能，如果有日程安排冲突了，就返回 false，如果不冲突则返回 ture
- 这一题有多种解法，第一种解法可以用类似第 34 题的解法。先排序每个区间，然后再这个集合中用二分搜索找到最后一个区间的左值比当前要比较的区间左值小的，如果找到，再判断能否插入进去(判断右区间是否比下一个区间的左区间小)，此方法时间复杂度 O(n log n)
- 第二种解法是用生成一个 BST 树。在插入树中先排除不能插入的情况，例如区间有重合。然后以区间左值为依据，递归插入，每次插入依次会继续判断区间是否重合。直到不能插入，则返回 fasle。整个查找的时间复杂度是 O(log n)。