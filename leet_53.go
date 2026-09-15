package main

func maxSubArray2(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		res = max(res, nums[i])
	}
	return res
}

func maxSubArray(nums []int) int {
	ll := len(nums)
	if ll == 1 {
		return nums[0]
	}
	nums2 := make([]int, 0)

	f := 1
	if nums[0] < 0 {
		f = -1
	}
	m := nums[0]
	mm := m
	for i := 1; i < ll; i++ {
		if nums[i]*f >= 0 {
			mm = max(mm, m, nums[i])
			m += nums[i]
		} else {
			f *= -1
			nums2 = append(nums2, m)
			mm = max(mm, m, nums[i])
			m = nums[i]
		}
	}

	nums2 = append(nums2, m)
	mm = max(mm, m)
	ll = len(nums2)
	if ll == 1 {
		return mm
	}

	s := 0
	e := ll
	if nums2[0] < 0 {
		s = 1
	}
	if nums2[ll-1] < 0 {
		e = ll - 1
	}
	nums2 = nums2[s:e]

	m = nums2[0]
	for i := 1; i < len(nums2); i++ {
		m = nums2[i] + m
		mm = max(m, nums2[i], mm)
		if m < 0 {
			m = 0
		}
	}

	return mm
}
