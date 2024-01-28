package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	for i := m + n - 1; i >= 0; i-- {
		if p2 < 0 {
			break
		}

		if p1 < 0 {
			nums1[i] = nums2[p2]
			p2--
			continue
		}
		l := nums1[p1]
		r := nums2[p2]
		if r >= l {
			nums1[i] = r
			p2--
		} else {
			nums1[i] = l
			p1--
		}
	}
}
