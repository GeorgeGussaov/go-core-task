package main

func GetCommonValues(slice1, slice2 []int) (bool, []int) {
	firstItems := make(map[int]bool)
	for _, item := range slice1 {
		firstItems[item] = true
	}

	result := make([]int, 0)
	for _, item := range slice2 {
		if firstItems[item] {
			result = append(result, item)
		}
	}

	if len(result) != 0 {
		return true, result
	}
	return false, nil
}
