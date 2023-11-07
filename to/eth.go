package to

const (
	gwei  = 1e9
	ether = 1e18
)

func ToGWei(v uint64) float64 {
	return float64(v) / gwei
}

func ToEther(v uint64) float64 {
	return float64(v) / ether
}

func FromGWei(v float64) uint64 {
	return uint64(v * gwei)
}

func FromEther(v float64) uint64 {
	return uint64(v * ether)
}
