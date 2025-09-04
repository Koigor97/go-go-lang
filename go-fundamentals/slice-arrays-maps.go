package main

func sliceManipulation(theArray []int, action string) ([]int, int) {
	switch action {
	case "len":
		return nil, len(theArray)
	case "cap":
		return nil, cap(theArray)
	}
	return theArray, len(theArray)
}
