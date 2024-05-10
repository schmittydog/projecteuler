package projecteuler

import (
	"fmt"
	"log"
	"math"
	"sort"
)

var N = 10000000 
var sqrt int
var sieve []int16

func Pow(x, y int) int {
	f := math.Pow(float64(x), float64(y))
	return int(f)
}

func PowMod(x, y, mod int) int {
	n := 1
	ret := 1
	for n <= y {
		if n&y != 0 {
			ret *= x
			ret %= mod
		}
		x *= x
		x %= mod
		n *= 2
	}
	return ret
}

func Factors(n int) []int {
	if n > N {
		log.Printf("Increasing sieve to %d", n*2)
		InitSieve(n*2)
	}
	if n == 0 {
		panic("Can't factor 0 you idiot")
	}
	ret := []int{}

	// Since we have an odd sieve, gotta check for 2's
	for n%2 == 0 {
		ret = append(ret, 2)
		n /= 2
	}

	for n > 1 {
		idx := n / 2
		if sieve[idx] == 0 {
			ret = append(ret, n)
			break
		}
		ret = append(ret, int(sieve[idx]))
		n /= int(sieve[idx])
	}

	sort.Ints(ret)
	return ret
}

func FactorTups(n int) [][]int {
	facs := Factors(n)
	arr := [][]int{}
	cur := facs[0]
	count := 1
	for _, f := range facs[1:] {
		if f != cur {
			arr = append(arr, []int{cur, count})
			cur, count = f, 1
		} else {
			count++
		}
	}
	arr = append(arr, []int{cur, count})
	return arr
}

func Divisors(n int) []int {
	facs := FactorTups(n)
	return divisorHelper(facs, 1)
}

func divisorHelper(facTups [][]int, n int) []int {
	if len(facTups) == 0 {
		return []int{n}
	}
	ret := []int{}
	for i := 0; i <= facTups[0][1]; i++ {
		ret = append(ret, divisorHelper(facTups[1:], n*Pow(facTups[0][0], i))...)
	}
	sort.Ints(ret)
	return ret
}

func Totient(n int) int {
	for _, f := range FactorTups(n) {
		n *= f[0]-1
		n /= f[0]
	}
	return n
}

func InitSieve(n int) {
	N = n / 2 // Doing an odd number sieve, so make 1/2 size
	sqrt = int(math.Sqrt(float64(N)))
	sieve = make([]int16, N)
	for i := 1; i <= sqrt; i++ {
		n := 2*i + 1
		if sieve[i] == 0 {
			for k := i + n; k < N; k += n {
				sieve[k] = int16(n)
			}
		}
	}
}

func init() {
	InitSieve(N)
}
