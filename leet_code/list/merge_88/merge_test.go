package merge_88

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	index := m + n - 1
	for ; index >= 0; index-- {
		if j < 0 || (i >= 0 && nums1[i] >= nums2[j]) {
			nums1[index] = nums1[i]
			i--
		} else {
			nums1[index] = nums2[j]
			j--
		}
	}
}
