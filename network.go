// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

// A map of UniFi Networks with a given attribute as key
type NetworkMap map[string]Network

// UniFi Network object
type Network struct {
	ID                string `json:"_id"`
	AttrHiddenID      string `json:"attr_hidden_id"`
	AttrNoDelete      bool   `json:"attr_no_delete"`
	DhcpdEnabled      bool   `json:"dhcpd_enabled"`
	DhcpdStart        string `json:"dhcpd_start"`
	DhcpdStop         string `json:"dhcpd_stop"`
	Dhcpdv6Enabled    bool   `json:"dhcpdv6_enabled"`
	DomainName        string `json:"domain_name"`
	IPSubnet          string `json:"ip_subnet"`
	Ipv6InterfaceType string `json:"ipv6_interface_type"`
	Ipv6RaEnabled     bool   `json:"ipv6_ra_enabled"`
	IsNat             bool   `json:"is_nat"`
	Name              string `json:"name"`
	Networkgroup      string `json:"networkgroup"`
	Purpose           string `json:"purpose"`
	SiteID            string `json:"site_id"`
	VlanEnabled       bool   `json:"vlan_enabled"`
}

// Networks Returns a slice of known networks
func (u *Unifi) Networks(site *Site) ([]Network, error) {
	var response struct {
		Data []Network
		Meta meta
	}
	err := u.parse(site, "rest/networkconf", nil, &response)
	return response.Data, err
}

// NetworkMap Returns a map of networkconfigs of a given site with ID as key
func (u *Unifi) NetworkMap(site *Site) (NetworkMap, error) {
	networks, err := u.Networks(site)
	if err != nil {
		return nil, err
	}
	m := make(NetworkMap)
	for _, n := range networks {
		m[n.ID] = n
	}
	return m, nil
}
