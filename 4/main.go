package main

func FirstSliceUniqElements(slice1, slice2 []string) []string { //норм название не придумал :)
	secondItems := make(map[string]bool)
	for _, item := range slice2 {
		secondItems[item] = true
	}

	result := make([]string, 0)
	for _, item := range slice1 {
		if !secondItems[item] {
			result = append(result, item)
		}
	}

	return result
}
