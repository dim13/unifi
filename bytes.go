// Copyright (c) 2014 Dimitri Sokolyuk. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

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
	return fmt.Sprintf("%6.1f %sB", b, string(dim[i]))
}
