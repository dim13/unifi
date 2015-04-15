package unifi

import "fmt"

type Bytes float64

func (b Bytes) String() string {
	var i int
	dim := ` kMGTPEZY`
	for b > 1024 {
		b /= 1024
		i++
	}
	return fmt.Sprintf("%5.1f %sB", b, string(dim[i]))
}
