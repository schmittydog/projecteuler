//General Utilities for Project Euler
package peuler

import (
	"io/ioutil"
	"os"
	"math/big"

	"github.com/schmittydog/projecteuler/bitsieve"
)

var (
	BS     *bitsieve.BitSieve
	Primes []int
	N      int
)

func GetFile(fname string) string {
	f, _ := os.Open(fname)
	defer f.Close()
	b, _ := ioutil.ReadAll(f)
	return string(b[:len(b)-1])
}

func PrimesTo(n int) {
	BS = bitsieve.NewBitSieve(n)
	Primes = BS.GetArray(n)
	N = n
}

func IsPrime(n int) bool {
	if n <= N && BS.IsIn(n) {
		return true
	}
	bn := big.NewInt(int64(n))
	return bn.ProbablyPrime(15)
}

func GetPrimeFactors(n int) []int {
	factors := []int{}
	for _, p := range Primes {
		if p*p > n {
			break
		}
		if n%p == 0 {
			factors = append(factors, p)
		}
		for n%p == 0 {
			n /= p
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

func GetPrimeTups(n int) [][]int {
	primes := GetPrimeFactors(n)
	ret := [][]int{}
	for _, prime := range primes {
		pcount := 0
		for n%prime == 0 {
			pcount += 1
			n /= prime
		}
		ret = append(ret, []int{prime, pcount})
	}
	return ret
}

func PowMod(a, b, p int) int {
	a %= p
	if a == 0 {
		return 0
	}
	product := 1
	for b > 0 {
		if b&1 == 1 {
			product *= a
			product %= p
			b--
		}
		a *= a
		a %= p
		b >>= 1
	}
	return product
}

func InvPowMod(a, p int) int {
	return PowMod(a, p-2, p)
}

func CombiMod(fact_n, fact_k, fact_n_k, p int) int {
	return ((fact_n * InvPowMod(fact_k, p) % p) * InvPowMod(fact_n_k, p)) % p
}
