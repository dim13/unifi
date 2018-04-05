// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

type PortprofileMap map[string]PortProfile

type PortProfile struct {
	ID                   string   `json:"_id"`
	AttrHidden           bool     `json:"attr_hidden,omitempty"`
	AttrHiddenID         string   `json:"attr_hidden_id,omitempty"`
	AttrNoDelete         bool     `json:"attr_no_delete,omitempty"`
	AttrNoEdit           bool     `json:"attr_no_edit,omitempty"`
	Forward              string   `json:"forward"`
	Name                 string   `json:"name"`
	SiteID               string   `json:"site_id"`
	NativeNetworkconfID  string   `json:"native_networkconf_id,omitempty"`
	Autoneg              bool     `json:"autoneg,omitempty"`
	Dot1XCtrl            string   `json:"dot1x_ctrl,omitempty"`
	LldpmedEnabled       bool     `json:"lldpmed_enabled,omitempty"`
	OpMode               string   `json:"op_mode,omitempty"`
	TaggedNetworkconfIds []string `json:"tagged_networkconf_ids,omitempty"`
}
