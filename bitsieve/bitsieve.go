// bitsieve creates a primesieve with bits for lower memory usage
package bitsieve

type BitSieve struct {
	Array []uint64
	Exp   []uint64
}

// NewBitSieve returns a bitsieve for primes up to n
func NewBitSieve(n int) *BitSieve {
	bs := new(BitSieve)
	bs.Exp = make([]uint64, 64)
	bs.Exp[0] = 1
	for i := 1; i < 64; i++ {
		bs.Exp[i] = 2 * bs.Exp[i-1]
	}
	bs.Array = make([]uint64, n/64+1)
	bs.Init()
	return bs
}

func (bs *BitSieve) SetBit(n int) {
	bs.Array[n/64] |= bs.Exp[n%64]
}

func (bs *BitSieve) IsIn(n int) bool {
	if n < 2 {
		return false
	}
	return bs.Array[n/64]&bs.Exp[n%64] == 0
}

// Init initializes the bitsieve for usage
func (bs *BitSieve) Init() {
	l := len(bs.Array) * 64
	for i := 2; i < l; i++ {
		if !bs.IsIn(i) {
			continue
		}
		for j := i; i*j < l && j > 0; j++ {
			bs.SetBit(i * j)
		}
	}
}

// GetArray returns an integer slice of primes
func (bs *BitSieve) GetArray(n int) []int {
	arr := []int{}
	for x := 2; x <= n; x++ {
		if bs.IsIn(x) {
			arr = append(arr, x)
		}
	}
	return arr
}
