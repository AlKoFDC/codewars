package kata

//     P\N 0,  1,  2,  3,  4,  5,  6,  7,  8,  9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21
//     0:  1,  2,  3,  4,  5,  6,  7,  8,  9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22
//     1:      1,  3,  6, 10, 15, 21, 28, 36, 45, 55, 66, 78, 91,105,120,136,153,171,190,210,231
//     2:          1,  4, 10, 20, 35, 56

type np struct {
	n int
	p int
}

var cache map[np]int = make(map[np]int)

func Diagonal(n, p int) (result int) {
	key := np{n: n, p: p}
	defer func() {
		cache[key] = result
	}()
	if result, ok := cache[key]; ok {
		return result
	}

	// Trivial cases
	switch {
	case p > n, n < 0, p < 0:
		return 0
	case n == p:
		return 1
	case p == 0:
		return n + 1
	}

	n64, p64 := uint64(n), uint64(p)

	// Factorial calculation is too big.
	if n64 >= 20 {
		return Diagonal(n-1, p) + Diagonal(n-1, p-1)
	}

	// Default calculation algorithm.
	nChoosePSum := uint64(0)
	for nLoop := n64; nLoop >= p64; nLoop = nLoop - 1 {
		nChoosePSum += factorial(nLoop) / (factorial(nLoop-p64) * factorial(p64))
	}
	return int(nChoosePSum)
}

func factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}
