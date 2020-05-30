// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

import (
	"encoding/json"
	"fmt"
	"time"
)

// Access point data
type USW struct {
	u  *Unifi
	ID string `json:"_id"`
	//Uptime        int    `json:"_uptime"`
	Adopted       bool    `json:"adopted"`
	Bytes         float64 `json:"bytes"`
	Cfgversion    string  `json:"cfgversion"`
	ConfigNetwork struct {
		IP   string `json:"ip"`
		Type string `json:"type"`
	} `json:"config_network"`
	ConnectRequestIP string `json:"connect_request_ip"`
	//ConnectRequestPort int           `json:"connect_request_port"` //FIXME: Prior to 5.7.20 type string!
	ConsideredLostAt     int           `json:"considered_lost_at"`
	DeviceID             string        `json:"device_id"`
	DhcpServerTable      []interface{} `json:"dhcp_server_table"`
	Dot1XPortctrlEnabled bool          `json:"dot1x_portctrl_enabled"`
	DownlinkTable        []interface{} `json:"downlink_table"`
	EthernetTable        []struct {
		Mac     string `json:"mac"`
		Name    string `json:"name"`
		NumPort int    `json:"num_port,omitempty"`
	} `json:"ethernet_table"`
	FlowctrlEnabled    bool           `json:"flowctrl_enabled"`
	FwCaps             int            `json:"fw_caps"`
	GeneralTemperature int            `json:"general_temperature"`
	GuestNumSta        int            `json:"guest-num_sta"`
	HasFan             bool           `json:"has_fan"`
	HasTemperature     bool           `json:"has_temperature"`
	InformIP           string         `json:"inform_ip"`
	InformURL          string         `json:"inform_url"`
	IP                 string         `json:"ip"`
	JumboframeEnabled  bool           `json:"jumboframe_enabled"`
	KnownCfgversion    string         `json:"known_cfgversion"`
	LastSeen           int            `json:"last_seen"`
	LedOverride        string         `json:"led_override"`
	LicenseState       string         `json:"license_state"`
	Locating           bool           `json:"locating"`
	Mac                string         `json:"mac"`
	Model              string         `json:"model"`
	Name               string         `json:"name"`
	NextHeartbeatAt    int            `json:"next_heartbeat_at"`
	NumSta             int            `json:"num_sta"`
	Overheating        bool           `json:"overheating"`
	PortOverrides      []PortOverride `json:"port_overrides,omitempty"`
	PortTable          []Port         `json:"port_table"`
	RxBytes            int64          `json:"rx_bytes"`
	Serial             string         `json:"serial"`
	SiteID             string         `json:"site_id"`
	SSHSessionTable    []interface{}  `json:"ssh_session_table"`
	Stat               struct {
		Bytes    float64   `json:"bytes"`
		Datetime time.Time `json:"datetime"`
		Duration float64   `json:"duration"`
		O        string    `json:"o"`
		Oid      string    `json:"oid"`
		//Port1RxBytes   float64     `json:"port_1-rx_bytes"`
		//Port1RxPackets int       `json:"port_1-rx_packets"`
		//Port1TxBytes   int64     `json:"port_1-tx_bytes"`
		//Port1TxPackets int       `json:"port_1-tx_packets"`
		RxBroadcast float64 `json:"rx_broadcast"`
		RxBytes     float64 `json:"rx_bytes"`
		RxCrypts    float64 `json:"rx_crypts"`
		RxDropped   float64 `json:"rx_dropped"`
		RxErrors    float64 `json:"rx_errors"`
		RxFrags     float64 `json:"rx_frags"`
		RxMulticast float64 `json:"rx_multicast"`
		RxPackets   float64 `json:"rx_packets"`
		SiteID      string  `json:"site_id"`
		Sw          string  `json:"sw"`
		Time        int64   `json:"time"`
		TxBroadcast float64 `json:"tx_broadcast"`
		TxBytes     float64 `json:"tx_bytes"`
		TxDropped   float64 `json:"tx_dropped"`
		TxErrors    float64 `json:"tx_errors"`
		TxMulticast float64 `json:"tx_multicast"`
		TxPackets   float64 `json:"tx_packets"`
		TxRetries   float64 `json:"tx_retries"`
	} `json:"stat"`
	State       DevState `json:"state"`
	StpPriority string   `json:"stp_priority"`
	StpVersion  string   `json:"stp_version"`
	SysStats    struct {
	} `json:"sys_stats"`
	SystemStats struct {
	} `json:"system-stats"`
	TxBytes           int64  `json:"tx_bytes"`
	Type              string `json:"type"`
	Upgradable        bool   `json:"upgradable"`
	UpgradeState      int    `json:"upgrade_state"`
	UpgradeToFirmware string `json:"upgrade_to_firmware"`
	Uplink            struct {
		FullDuplex  bool   `json:"full_duplex"`
		IP          string `json:"ip"`
		Mac         string `json:"mac"`
		MaxSpeed    int    `json:"max_speed"`
		Media       string `json:"media"`
		Name        string `json:"name"`
		Netmask     string `json:"netmask"`
		NumPort     int    `json:"num_port"`
		PortIdx     int    `json:"port_idx"`
		RxBytes     int64  `json:"rx_bytes"`
		RxBytesR    int64  `json:"rx_bytes-r"`
		RxDropped   int64  `json:"rx_dropped"`
		RxErrors    int64  `json:"rx_errors"`
		RxMulticast int64  `json:"rx_multicast"`
		RxPackets   int64  `json:"rx_packets"`
		Speed       int    `json:"speed"`
		TxBytes     int64  `json:"tx_bytes"`
		TxBytesR    int64  `json:"tx_bytes-r"`
		TxDropped   int64  `json:"tx_dropped"`
		TxErrors    int64  `json:"tx_errors"`
		TxPackets   int64  `json:"tx_packets"`
		Type        string `json:"type"`
		Up          bool   `json:"up"`
	} `json:"uplink"`
	UplinkDepth            int    `json:"uplink_depth"`
	Uptime                 int    `json:"uptime"`
	UserNumSta             int    `json:"user-num_sta"`
	Version                string `json:"version"`
	VersionIncompatible    bool   `json:"version_incompatible"`
	XAuthkey               string `json:"x_authkey"`
	XFingerprint           string `json:"x_fingerprint"`
	XHasSSHHostkey         bool   `json:"x_has_ssh_hostkey"`
	XInformAuthkey         string `json:"x_inform_authkey"`
	XSSHHostkeyFingerprint string `json:"x_ssh_hostkey_fingerprint"`
}

type USWmap map[string]USW

// Returns a slice of switches
func (u *Unifi) USWs(site *Site) ([]USW, error) {

	// Devices
	var usws []USW

	rawDevices, err := u.RawDevices(site, "usw")

	for i, _ := range rawDevices {
		var usw USW
		err := json.Unmarshal(rawDevices[i].Data, &usw)
		if err != nil {
			return usws, err
		}
		// Set unifi pointer
		usw.u = u

		usws = append(usws, usw)
	}
	return usws, err
}

// Returns a map of access points with mac as a key
func (u *Unifi) USWmap(site *Site) (USWmap, error) {
	usws, err := u.USWs(site)
	if err != nil {
		return nil, err
	}
	m := make(USWmap)
	for _, s := range usws {
		m[s.Mac] = s
	}
	return m, nil
}

// Returns a USW pointer for USW with a given name
func (u *Unifi) USW(site *Site, name string) (*USW, error) {
	devices, err := u.USWs(site)
	if err != nil {
		return nil, err
	}
	for _, d := range devices {
		if name == d.DeviceName() {
			return &d, nil
		}
	}
	return nil, fmt.Errorf("No device with name: %s", name)
}

// Reboot access point
func (s USW) Restart() error {
	if s.u == nil {
		return ErrLoginFirst
	}
	return s.u.devcmd(s.Mac, "restart")
}

func (s USW) DeviceName() string {
	if s.Name != "" {
		return s.Name
	}
	// If no name is given, return mac as name
	return s.Mac
}

func (s USW) ModelName() string {
	if m, ok := model[s.Model]; ok {
		return m
	}
	return s.Model
}
