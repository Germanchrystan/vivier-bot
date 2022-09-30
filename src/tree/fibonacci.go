package tree

func GetFibValueByIndex(index int) int {
	if index == 1 {
		return 1
	}

	if index < 1 {
		return 0
	}

	return GetFibValueByIndex(index-1) + GetFibValueByIndex(index-2)
}
