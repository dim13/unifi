// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

import (
	"encoding/json"
	"fmt"
)

type LAN struct {
	LanIP           string `json:"lan_ip"`
	NumAdopted      int    `json:"num_adopted"`
	NumDisconnected int    `json:"num_disconnected"`
	NumGuest        int    `json:"num_guest"`
	NumPending      int    `json:"num_pending"`
	NumSw           int    `json:"num_sw"`
	NumUser         int    `json:"num_user"`
	RxBytesR        int64  `json:"rx_bytes-r"`
	Status          string `json:"status"`
	Subsystem       string `json:"subsystem"`
	TxBytesR        int64  `json:"tx_bytes-r"`
}

type VPN struct {
	Status    string `json:"status"`
	Subsystem string `json:"subsystem"`
}

type WAN struct {
	Gateways      []string `json:"gateways"`
	GwMac         string   `json:"gw_mac"`
	GwName        string   `json:"gw_name"`
	GwSystemStats struct {
		CPU   string `json:"cpu"`
		Mem   string `json:"mem"`
		Temps struct {
			BoardCPU string `json:"Board (CPU)"`
			BoardPHY string `json:"Board (PHY)"`
			CPU      string `json:"CPU"`
			PHY      string `json:"PHY"`
		} `json:"temps"`
		Uptime string `json:"uptime"`
	} `json:"gw_system-stats"`
	GwVersion       string   `json:"gw_version"`
	Nameservers     []string `json:"nameservers"`
	Netmask         string   `json:"netmask"`
	NumAdopted      int      `json:"num_adopted"`
	NumDisconnected int      `json:"num_disconnected"`
	NumGw           int      `json:"num_gw"`
	NumPending      int      `json:"num_pending"`
	NumSta          int      `json:"num_sta"`
	RxBytesR        int64    `json:"rx_bytes-r"`
	Status          string   `json:"status"`
	Subsystem       string   `json:"subsystem"`
	TxBytesR        int64    `json:"tx_bytes-r"`
	WanIP           string   `json:"wan_ip"`
}

type WLAN struct {
	NumAdopted      int    `json:"num_adopted"`
	NumAp           int    `json:"num_ap"`
	NumDisabled     int    `json:"num_disabled"`
	NumDisconnected int    `json:"num_disconnected"`
	NumGuest        int    `json:"num_guest"`
	NumPending      int    `json:"num_pending"`
	NumUser         int    `json:"num_user"`
	RxBytesR        int64  `json:"rx_bytes-r"`
	Status          string `json:"status"`
	Subsystem       string `json:"subsystem"`
	TxBytesR        int64  `json:"tx_bytes-r"`
}

type WWW struct {
	Drops            int     `json:"drops"`
	GwMac            string  `json:"gw_mac"`
	Latency          int     `json:"latency"`
	RxBytesR         int64   `json:"rx_bytes-r"`
	SpeedtestLastrun int     `json:"speedtest_lastrun"`
	SpeedtestPing    int     `json:"speedtest_ping"`
	SpeedtestStatus  string  `json:"speedtest_status"`
	Status           string  `json:"status"`
	Subsystem        string  `json:"subsystem"`
	TxBytesR         int     `json:"tx_bytes-r"`
	Uptime           int     `json:"uptime"`
	XputDown         float64 `json:"xput_down"`
	XputUp           float64 `json:"xput_up"`
}

type Health struct {
	LAN  LAN
	VPN  VPN
	WAN  WAN
	WLAN WLAN
	WWW  WWW
}

// Returns a slice of access points
func (u *Unifi) Health(site *Site) (Health, error) {

	// The health struct
	var health Health

	// Response[]UAP from controller
	var response struct {
		Data []json.RawMessage
		Meta meta
	}

	// Fetch the health data
	err := u.parse(site, "stat/health", nil, &response)

	// Loop thru the data objects to parse them with the corresponding struct
	for _, d := range response.Data {

		// unmarshal into a map to check the "subsystem" field
		var obj map[string]interface{}
		err := json.Unmarshal(d, &obj)
		if err != nil {
			return health, err
		}

		subsystem, ok := obj["subsystem"].(string)
		if !ok {
			return health, fmt.Errorf("Error on retrieving subsystem from raw Json")
		}

		switch subsystem {
		case "lan":
			var lan LAN

			err := json.Unmarshal(d, &lan)
			if err != nil {
				return health, err
			}
			// Unmarshal successful. Add to health object
			health.LAN = lan

		case "vpn":
			var vpn VPN

			err := json.Unmarshal(d, &vpn)
			if err != nil {
				return health, err
			}
			// Unmarshal successful. Add to health object
			health.VPN = vpn

		case "wlan":
			var wlan WLAN

			err := json.Unmarshal(d, &wlan)
			if err != nil {
				return health, err
			}
			// Unmarshal successful. Add to health object
			health.WLAN = wlan

		case "wan":
			var wan WAN

			err := json.Unmarshal(d, &wan)
			if err != nil {
				return health, err
			}
			// Unmarshal successful. Add to health object
			health.WAN = wan

		case "www":
			var www WWW

			err := json.Unmarshal(d, &www)
			if err != nil {
				return health, err
			}
			// Unmarshal successful. Add to health object
			health.WWW = www

		}
	}
	return health, err
}
