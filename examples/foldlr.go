package examples

// foldr applies a function from right to left over a slice
func foldr[T, U any](f func(T, U) U, acc U, xs []T) U {
	for i := len(xs) - 1; i >= 0; i-- {
		acc = f(xs[i], acc)
	}
	return acc
}

// foldl applies a function from left to right over a slice
func foldl[T, U any](f func(U, T) U, acc U, xs []T) U {
	for _, x := range xs {
		acc = f(acc, x)
	}
	return acc
}
