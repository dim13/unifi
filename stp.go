// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

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
