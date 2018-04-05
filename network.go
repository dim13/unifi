// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

type NetworkMap map[string]Network

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
