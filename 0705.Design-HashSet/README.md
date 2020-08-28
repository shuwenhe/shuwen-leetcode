# [705. Design HashSet](https://leetcode.com/problems/design-hashset/)


## 题目

Design a HashSet without using any built-in hash table libraries.

To be specific, your design should include these functions:

- `add(value)`: Insert a value into the HashSet.
- `contains(value)` : Return whether the value exists in the HashSet or not.
- `remove(value)`: Remove a value in the HashSet. If the value does not exist in the HashSet, do nothing.

**Example:**

    MyHashSet hashSet = new MyHashSet();
    hashSet.add(1);         
    hashSet.add(2);         
    hashSet.contains(1);    // returns true
    hashSet.contains(3);    // returns false (not found)
    hashSet.add(2);          
    hashSet.contains(2);    // returns true
    hashSet.remove(2);          
    hashSet.contains(2);    // returns false (already removed)

**Note:**

- All values will be in the range of `[0, 1000000]`.
- The number of operations will be in the range of `[1, 10000]`.
- Please do not use the built-in HashSet library.


## 题目大意

不使用任何内建的哈希表库设计一个哈希集合具体地说，你的设计应该包含以下的功能：

- add(value)：向哈希集合中插入一个值。
- contains(value) ：返回哈希集合中是否存在这个值。
- remove(value)：将给定值从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。


注意：

- 所有的值都在 [1, 1000000] 的范围内。
- 操作的总数目在 [1, 10000] 范围内。
- 不要使用内建的哈希集合库。



## 解题思路


- 简单题，设计一个 hashset 的数据结构，要求有 `add(value)`，`contains(value)`，`remove(value)`，这 3 个方法。
