package slices

import (
	"strconv"
	"strings"
)

func UniqueStrings(items []string) []string {
	m := make(map[string]bool, len(items))
	for _, item := range items {
		m[item] = true
	}
	var result []string
	for k := range m {
		result = append(result, k)
	}
	return result
}

func TrimStrings(items []string) []string {
	var result []string
	for _, item := range items {
		item := strings.TrimSpace(item)
		if item != "" {
			result = append(result, item)
		}
	}
	return result
}

func RemoveStrings(source []string, target []string) []string {
	var result []string
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

func ToInts(source []string) ([]int, error) {
	var result []int
	for _, item := range source {
		v, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}

func ToInt64s(source []string) ([]int64, error) {
	var result []int64
	for _, item := range source {
		v, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}

func ToFloat64s(source []string) ([]float64, error) {
	var result []float64
	for _, item := range source {
		v, err := strconv.ParseFloat(item, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}

func ToBools(source []string) ([]bool, error) {
	var result []bool
	for _, item := range source {
		v, err := strconv.ParseBool(item)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}

type Strings []string

func (source Strings) Unique() []string {

	return UniqueStrings(source)
}

func (source Strings) Trim() []string {
	return TrimStrings(source)
}

func (source Strings) Remove(target []string) []string {
	return RemoveStrings(source, target)
}

func (source Strings) ToInts() ([]int, error) {
	var result []int
	for _, item := range source {
		v, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}

func (source Strings) ToInt64s() ([]int64, error) {
	var result []int64
	for _, item := range source {
		v, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}

func (source Strings) ToFloat64s() ([]float64, error) {
	var result []float64
	for _, item := range source {
		v, err := strconv.ParseFloat(item, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}

func (source Strings) ToBools() ([]bool, error) {
	var result []bool
	for _, item := range source {
		v, err := strconv.ParseBool(item)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}
