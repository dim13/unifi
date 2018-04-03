package unifi

type STP int

//go:generate stringer -type=STP

// Spanning-Tree states
const (
	StpDisabled   STP = 0
	StpBlocking   STP = 1
	StpListening  STP = 2
	StpLearning   STP = 3
	StpForwarding STP = 4
	StpUnknown    STP = 9
)
