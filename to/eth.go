package to

import "math/big"

const (
	gwei  = 1e9
	ether = 1e18
)

func ToGWei(v uint64) float64 {
	return float64(v) / gwei
}

func ToEther(v *big.Int) float64 {
	dd, _ := v.Float64()
	return dd / ether
}

func FromGWei(v float64) uint64 {
	return uint64(v * gwei)
}

func FromEther(v float64) *big.Int {
	ret := new(big.Int)
	ret.SetUint64(uint64(v * gwei))
	return ret.Mul(ret, big.NewInt(gwei))
}
