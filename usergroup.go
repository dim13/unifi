// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

type UserGroup struct {
	AttrHiddenID    string `json:"attr_hidden_id"`
	AttrNoDelete    bool   `json:"attr_no_delete"`
	DownRateEnabled bool   `json:"downrate_enabled"`
	Name            string
	QosRateMaxDown  int  `json:"qos_rate_max_down"`
	QosRateMaxUp    int  `json:"qos_rate_max_up"`
	UpRateEnabled   bool `json:"uprate_enabled"`
}
