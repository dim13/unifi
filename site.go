// Copyright (c) 2014 Dimitri Sokolyuk. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

//type StaMap map[string]Sta

// Station data
type Site struct {
	//u                *Unifi
	Id           string `json:"_id"` // For internal use
	AttrHiddenId string `json:"attr_hidden_id"`
	AttrNoDelete bool   `json:"attr_no_delete"`
	Desc         string // The name of the site!
	Name         string // The site-id!
	Role         string
}
