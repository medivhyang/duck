package slices

func ContainInts(source []int, target ...int) bool {
	for _, t := range target {
		flag := false
		for _, s := range source {
			if t == s {
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}
	return true
}

func UniqueInts(items []int) []int {
	m := make(map[int]bool, len(items))
	for _, item := range items {
		m[item] = true
	}
	var result []int
	for k := range m {
		result = append(result, k)
	}
	return result
}

func RemoveInts(source []int, target []int) []int {
	var result []int
	for _, sourceItem := range source {
		find := false
		for _, targetItem := range target {
			if sourceItem == targetItem {
				find = true
			}
		}
		if !find {
			result = append(result, sourceItem)
		}
	}
	return result
}
