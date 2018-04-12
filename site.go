// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

import "fmt"

// UniFi Site object
type Site struct {
	ID           string `json:"_id"` // For internal use
	AttrHiddenID string `json:"attr_hidden_id"`
	AttrNoDelete bool   `json:"attr_no_delete"`
	Desc         string // The name of the site! (Friendly name)
	Name         string // The site-id!
	Role         string
}

// Site returns a site with given name or description
func (u *Unifi) Site(desc string) (*Site, error) {

	sites, err := u.Sites()
	if err != nil {
		return nil, err
	}

	// First, search site by description (friendly name)
	for _, s := range sites {
		if s.Desc == desc {
			return &s, nil
		}
	}

	// If not found, search site by name (id used in url)
	for _, s := range sites {
		if s.Name == desc {
			return &s, nil
		}
	}

	return nil, fmt.Errorf("Site %s not found", desc)
}

// Sites returns a slice with all sites of the controller
func (u *Unifi) Sites() ([]Site, error) {
	var response struct {
		Data []Site
		Meta meta
	}
	err := u.parse(nil, "self/sites", &response)
	return response.Data, err
}

/*// SiteNameByDesc returns the name (id) of a site, searched by its description (user friendly name)
// So far only for internal use
func (u *Unifi) siteNameByDesc(desc string) (string, error) {
	sites, err := u.Sites()
	if err != nil {
		return "", err
	}

	for _, s := range sites {
		if s.Desc == desc {
			return s.Name, nil
		}
	}

	return "", errors.New("No site with desc: " + desc)
} */
