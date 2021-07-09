package lc_merge

func merge(intervals [][]int) [][]int {
	sort(intervals)
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i+1][0] > intervals[i][1] {
			continue
		}
		if intervals[i+1][1] >= intervals[i][1] {
			intervals[i+1][0] = intervals[i][0]
		} else if intervals[i+1][1] < intervals[i][1] {
			intervals[i+1][0] = intervals[i][0]
			intervals[i+1][1] = intervals[i][1]
		}

		intervals = append(intervals[:i], intervals[i+1:]...)
		i--
	}
	return intervals
}

func sort(nums [][]int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i][0] > nums[j][0] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func qsort(from, to int, res [][]int) {
	begin := from
	end := to
	index := from
	if from >= to {
		return
	}

	for begin < end {
		for begin < end {
			if res[index][0] <= res[end][0] {
				end--
				continue
			}
			res[index], res[end] = res[end], res[index]
			index = end
			break
		}

		for begin < end {
			if res[index][0] >= res[begin][0] {
				begin++
				continue
			}
			res[index], res[begin] = res[begin], res[index]
			index = begin
			break
		}
	}

	qsort(from, index-1, res)
	qsort(index+1, to, res)

}
