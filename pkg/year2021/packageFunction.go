package year2021

func first[T, U any](val T, _ U) T {
	return val
}

func reverse[T any](lst []T) []T {
	var reversed = make([]T, 0)
	for i := (len(lst) - 1); i >= 0; i-- {
		reversed = append(reversed, lst[i])
	}
	return reversed
}
