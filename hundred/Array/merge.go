package Array

import "sort"

type Interval struct {
	Start int
	End   int
}

// merge 合并区间
func merge(intervals [][]int) [][]int {
	var merged []Interval
	for _, interval := range intervals {
		merged = append(merged, Interval{
			Start: interval[0],
			End:   interval[1],
		})
	}
	sort.Slice(merged, func(i, j int) bool {
		return merged[i].Start < merged[j].Start
	})

	var result []Interval
	result = append(result, merged[0])

	for _, interval := range merged[1:] {
		last := result[len(result)-1]
		if interval.Start <= last.End {
			last.End = max(last.End, interval.End)
			result[len(result)-1] = last
		} else {
			result = append(result, interval)
		}
	}

	var finalResult [][]int
	for _, interval := range result {
		finalResult = append(finalResult, []int{
			interval.Start,
			interval.End,
		})
	}
	return finalResult
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
