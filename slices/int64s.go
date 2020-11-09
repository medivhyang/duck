package slices

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

type Int64s []int64

func (source Int64s) Unique() []int64 {
	return UniqueInt64s(source)
}

func (source Int64s) Remove(target []int64) []int64 {
	return RemoveInt64s(source, target)
}
