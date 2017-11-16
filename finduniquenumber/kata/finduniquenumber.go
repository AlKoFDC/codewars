package kata

func FindUniq(arr []float32) float32 {
	const (
		firstValue = iota
		secondValue
		thirdValue
	)

	if len(arr) < 3 {
		panic("Instructions: It’s guaranteed that array contains more than 3 numbers.")
	}

	// Cover trivial case of 3 values where 1st or 2nd are different.
	testValue := arr[firstValue]
	if result, ok := ifFirstAndSecondDifferent(testValue, arr[secondValue], arr[thirdValue]); ok {
		return result
	}
	length := len(arr)
	for idx := thirdValue; length > idx; idx++ {
		if value := arr[idx]; testValue != value {
			return value
		}
	}
	panic("no different value found")
}

func ifFirstAndSecondDifferent(first, second, third float32) (float32, bool) {
	if first != second {
		if first == third {
			return second, true
		}
		return first, true
	}
	if third != first {
		return third, true
	}
	return first, false
}
