package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// Interval define
type Interval = structures.Interval

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

func insert(intervals []Interval, newInterval Interval) []Interval {
	res := make([]Interval, 0)
	if len(intervals) == 0 {
		res = append(res, newInterval)
		return res
	}
	curIndex := 0
	for curIndex < len(intervals) && intervals[curIndex].End < newInterval.Start {
		res = append(res, intervals[curIndex])
		curIndex++
	}

	for curIndex < len(intervals) && intervals[curIndex].Start <= newInterval.End {
		newInterval = Interval{Start: min(newInterval.Start, intervals[curIndex].Start), End: max(newInterval.End, intervals[curIndex].End)}
		curIndex++
	}
	res = append(res, newInterval)

	for curIndex < len(intervals) {
		res = append(res, intervals[curIndex])
		curIndex++
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
