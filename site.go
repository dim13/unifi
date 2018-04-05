// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

// Station data
type Site struct {
	//u                *Unifi
	ID           string `json:"_id"` // For internal use
	AttrHiddenID string `json:"attr_hidden_id"`
	AttrNoDelete bool   `json:"attr_no_delete"`
	Desc         string // The name of the site!
	Name         string // The site-id!
	Role         string
}
