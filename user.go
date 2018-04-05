// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

type User struct {
	FirstSeen int `json:"first_seen"`
	Hostname  string
	IsGuest   bool `json:"is_guest"`
	LastSeen  int  `json:"last_seen"`
	Mac       string
	Oui       string
}
