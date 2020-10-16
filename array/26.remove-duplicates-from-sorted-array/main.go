package _6_remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1
}

// 解法二
func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := len(nums)
	lastNum := nums[length-1]
	i := 0
	for i = 0; i < length-1; i++ {
		if nums[i] == lastNum {
			break
		}
		if nums[i+1] == nums[i] {
			removeElement1(nums, i+1, nums[i])
			// fmt.Printf("此时 num = %v length = %v\n", nums, length)
		}
	}
	return i + 1
}

func removeElement1(nums []int, start, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := start
	for i := start; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
				j++
			} else {
				j++
			}
		}
	}
	return j
}

//双指针法，i为慢指针，j为快指针
//当nums[i] != nums[j] 时 i向前进一位，并且nums[i]等于 nums[j]
func removeDuplicates3(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
