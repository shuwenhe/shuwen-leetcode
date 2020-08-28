# [981. Time Based Key-Value Store](https://leetcode.com/problems/time-based-key-value-store/)


## 题目

Create a timebased key-value store class `TimeMap`, that supports two operations.

1. `set(string key, string value, int timestamp)`

- Stores the `key` and `value`, along with the given `timestamp`.

2. `get(string key, int timestamp)`

- Returns a value such that `set(key, value, timestamp_prev)` was called previously, with `timestamp_prev <= timestamp`.
- If there are multiple such values, it returns the one with the largest `timestamp_prev`.
- If there are no values, it returns the empty string (`""`).

**Example 1:**

    Input: inputs = ["TimeMap","set","get","get","set","get","get"], inputs = [[],["foo","bar",1],["foo",1],["foo",3],["foo","bar2",4],["foo",4],["foo",5]]
    Output: [null,null,"bar","bar",null,"bar2","bar2"]
    Explanation:   
    TimeMap kv;   
    kv.set("foo", "bar", 1); // store the key "foo" and value "bar" along with timestamp = 1   
    kv.get("foo", 1);  // output "bar"   
    kv.get("foo", 3); // output "bar" since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 ie "bar"   
    kv.set("foo", "bar2", 4);   
    kv.get("foo", 4); // output "bar2"   
    kv.get("foo", 5); //output "bar2"

**Example 2:**

    Input: inputs = ["TimeMap","set","set","get","get","get","get","get"], inputs = [[],["love","high",10],["love","low",20],["love",5],["love",10],["love",15],["love",20],["love",25]]
    Output: [null,null,null,"","high","high","low","low"]

**Note:**

1. All key/value strings are lowercase.
2. All key/value strings have length in the range `[1, 100]`
3. The `timestamps` for all `TimeMap.set` operations are strictly increasing.
4. `1 <= timestamp <= 10^7`
5. `TimeMap.set` and `TimeMap.get` functions will be called a total of `120000` times (combined) per test case.

## 题目大意

创建一个基于时间的键值存储类 TimeMap，它支持下面两个操作：

1. set(string key, string value, int timestamp)

- 存储键 key、值 value，以及给定的时间戳 timestamp。

2. get(string key, int timestamp)

- 返回先前调用 set(key, value, timestamp_prev) 所存储的值，其中 timestamp_prev <= timestamp。
- 如果有多个这样的值，则返回对应最大的  timestamp_prev 的那个值。
- 如果没有值，则返回空字符串（""）。

提示：

1. 所有的键/值字符串都是小写的。
2. 所有的键/值字符串长度都在 [1, 100] 范围内。
3. 所有 TimeMap.set 操作中的时间戳 timestamps 都是严格递增的。
4. 1 <= timestamp <= 10^7
5. TimeMap.set 和 TimeMap.get 函数在每个测试用例中将（组合）调用总计 120000 次。


## 解题思路

- 要求设计一个基于时间戳的 `kv` 存储。`set()` 操作里面会会包含一个时间戳。get() 操作的时候查找时间戳小于等于 `timestamp` 的 `key` 对应的 `value`，如果有多个解，输出满足条件的最大时间戳对应的 `value` 值。
- 这一题可以用二分搜索来解答，用 `map` 存储 `kv` 数据，`key` 对应的就是 `key`，`value` 对应一个结构体，里面包含 `value` 和 `timestamp`。执行 `get()` 操作的时候，先取出 `key` 对应的结构体数组，然后在这个数组里面根据 `timestamp` 进行二分搜索。由于题意是要找小于等于 `timestamp` 的最大 `timestamp` ，这会有很多满足条件的解，变换一下，先找 `> timestamp` 的最小解，然后下标减一即是满足题意的最大解。
- 另外题目中提到“`TimeMap.set` 操作中的 `timestamp` 是严格递增的”。所以在 `map` 中存储 `value` 结构体的时候，不需要排序了，天然有序。
