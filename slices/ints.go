package slices

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

type Ints []int

func (source Ints) Unique() []int {
	return UniqueInts(source)
}

func (source Ints) Remove(target []int) []int {
	return RemoveInts(source, target)
}
