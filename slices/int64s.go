package slices

func ContainInt64s(source []int64, target ...int64) bool {
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

func UniqueInt64s(items []int64) []int64 {
	m := make(map[int64]bool, len(items))
	for _, item := range items {
		m[item] = true
	}
	var result []int64
	for k := range m {
		result = append(result, k)
	}
	return result
}

func RemoveInt64s(source []int64, target []int64) []int64 {
	var result []int64
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
