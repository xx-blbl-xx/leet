package main

func minSum(nums1 []int, nums2 []int) int64 {
	s1 := 0
	s2 := 0
	n1 := 0
	n2 := 0

	for _, v := range nums1 {
		if v == 0 {
			n1++
			continue
		}
		s1 += v
	}
	for _, v := range nums2 {
		if v == 0 {
			n2++
			continue
		}
		s2 += v
	}

	s11 := s1 + n1
	s21 := s2 + n2

	if (n1 == 0 && s21 > s11) || (n2 == 0 && s11 > s21) {
		return -1
	}

	return int64(max(s11, s21))
}
