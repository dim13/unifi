// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

import (
	"encoding/json"
	"time"
)

// Access point data
type UAP struct {
	u  *Unifi
	ID string `json:"_id"`
	//Uptime        int           `json:"_uptime"`
	Adopted       bool          `json:"adopted"`
	AntennaTable  []interface{} `json:"antenna_table"`
	Bytes         int64         `json:"bytes"`
	BytesD        int64         `json:"bytes-d"`
	BytesR        int64         `json:"bytes-r"`
	Cfgversion    string        `json:"cfgversion"`
	ConfigNetwork struct {
		IP   string `json:"ip"`
		Type string `json:"type"`
	} `json:"config_network"`
	ConnectRequestIP string `json:"connect_request_ip"`
	//ConnectRequestPort int           `json:"connect_request_port"` //FIXME: Prior to 5.7.20 type string!
	ConsideredLostAt int           `json:"considered_lost_at"`
	DeviceID         string        `json:"device_id"`
	DownlinkTable    []interface{} `json:"downlink_table"`
	EthernetTable    []struct {
		Mac     string `json:"mac"`
		Name    string `json:"name"`
		NumPort int    `json:"num_port"`
	} `json:"ethernet_table"`
	FwCaps          int    `json:"fw_caps"`
	GuestNumSta     int    `json:"guest-num_sta"`
	GuestToken      string `json:"guest_token"`
	HasEth1         bool   `json:"has_eth1"`
	HasSpeaker      bool   `json:"has_speaker"`
	InformIP        string `json:"inform_ip"`
	InformURL       string `json:"inform_url"`
	IP              string `json:"ip"`
	Isolated        bool   `json:"isolated"`
	KnownCfgversion string `json:"known_cfgversion"`
	LastSeen        int    `json:"last_seen"`
	LedOverride     string `json:"led_override"`
	LicenseState    string `json:"license_state"`
	Locating        bool   `json:"locating"`

	Mac   string `json:"mac"`
	MapID string `json:"map_id"`
	Model string `json:"model"`

	Name            string        `json:"name"`
	NextHeartbeatAt int           `json:"next_heartbeat_at"`
	NumSta          int           `json:"num_sta"`
	PortStats       []interface{} `json:"port_stats"`
	PortTable       []Port        `json:"port_table,omitempty"`
	RadioTable      []struct {
		BuiltinAntGain     int    `json:"builtin_ant_gain"`
		BuiltinAntenna     bool   `json:"builtin_antenna"`
		CurrentAntennaGain int    `json:"current_antenna_gain"`
		MaxTxpower         int    `json:"max_txpower"`
		MinTxpower         int    `json:"min_txpower"`
		Name               string `json:"name"`
		Nss                int    `json:"nss"`
		Radio              string `json:"radio"`
		RadioCaps          int    `json:"radio_caps"`
		WlangroupID        string `json:"wlangroup_id"`
		HasDfs             bool   `json:"has_dfs,omitempty"`
		HasFccdfs          bool   `json:"has_fccdfs,omitempty"`
		Is11Ac             bool   `json:"is_11ac,omitempty"`
	} `json:"radio_table"`
	RadioTableStats []struct {
		AstBeXmit   interface{} `json:"ast_be_xmit"`
		AstCst      interface{} `json:"ast_cst"`
		AstTxto     interface{} `json:"ast_txto"`
		Channel     int         `json:"channel"`
		CuSelfRx    int         `json:"cu_self_rx"`
		CuSelfTx    int         `json:"cu_self_tx"`
		CuTotal     int         `json:"cu_total"`
		Extchannel  int         `json:"extchannel"`
		Gain        int         `json:"gain"`
		GuestNumSta int         `json:"guest-num_sta"`
		Name        string      `json:"name"`
		NumSta      int         `json:"num_sta"`
		Radio       string      `json:"radio"`
		State       string      `json:"state"`
		TxPackets   int64       `json:"tx_packets"`
		TxPower     int64       `json:"tx_power"`
		TxRetries   int64       `json:"tx_retries"`
		UserNumSta  int         `json:"user-num_sta"`
	} `json:"radio_table_stats"`

	RxBytes          int64         `json:"rx_bytes"`
	RxBytesD         int64         `json:"rx_bytes-d"`
	ScanRadioTable   []interface{} `json:"scan_radio_table"`
	Scanning         bool          `json:"scanning"`
	Serial           string        `json:"serial"`
	SiteID           string        `json:"site_id"`
	SpectrumScanning bool          `json:"spectrum_scanning"`
	SSHSessionTable  []interface{} `json:"ssh_session_table"`
	Stat             struct {
		Ap struct {
			Ap                                             string    `json:"ap"`
			Bytes                                          float64   `json:"bytes"`
			Datetime                                       time.Time `json:"datetime"`
			Duration                                       float64   `json:"duration"`
			GuestRxBytes                                   float64   `json:"guest-rx_bytes"`
			GuestRxCrypts                                  float64   `json:"guest-rx_crypts"`
			GuestRxDropped                                 float64   `json:"guest-rx_dropped"`
			GuestRxErrors                                  float64   `json:"guest-rx_errors"`
			GuestRxFrags                                   float64   `json:"guest-rx_frags"`
			GuestRxPackets                                 float64   `json:"guest-rx_packets"`
			GuestTxBytes                                   float64   `json:"guest-tx_bytes"`
			GuestTxDropped                                 float64   `json:"guest-tx_dropped"`
			GuestTxErrors                                  float64   `json:"guest-tx_errors"`
			GuestTxPackets                                 float64   `json:"guest-tx_packets"`
			GuestTxRetries                                 float64   `json:"guest-tx_retries"`
			O                                              string    `json:"o"`
			Oid                                            string    `json:"oid"`
			RxBytes                                        float64   `json:"rx_bytes"`
			RxCrypts                                       float64   `json:"rx_crypts"`
			RxDropped                                      float64   `json:"rx_dropped"`
			RxErrors                                       float64   `json:"rx_errors"`
			RxFrags                                        float64   `json:"rx_frags"`
			RxPackets                                      float64   `json:"rx_packets"`
			SiteID                                         string    `json:"site_id"`
			Time                                           int64     `json:"time"`
			TxBytes                                        float64   `json:"tx_bytes"`
			TxDropped                                      float64   `json:"tx_dropped"`
			TxErrors                                       float64   `json:"tx_errors"`
			TxPackets                                      float64   `json:"tx_packets"`
			TxRetries                                      float64   `json:"tx_retries"`
			UserRxBytes                                    float64   `json:"user-rx_bytes"`
			UserRxCrypts                                   float64   `json:"user-rx_crypts"`
			UserRxDropped                                  float64   `json:"user-rx_dropped"`
			UserRxErrors                                   float64   `json:"user-rx_errors"`
			UserRxFrags                                    float64   `json:"user-rx_frags"`
			UserRxPackets                                  float64   `json:"user-rx_packets"`
			UserTxBytes                                    float64   `json:"user-tx_bytes"`
			UserTxDropped                                  float64   `json:"user-tx_dropped"`
			UserTxErrors                                   float64   `json:"user-tx_errors"`
			UserTxPackets                                  float64   `json:"user-tx_packets"`
			UserTxRetries                                  float64   `json:"user-tx_retries"`
			UserWifi0Ath05Aa99001C9E77C011Dc149F7RxBytes   float64   `json:"user-wifi0-ath0-5aa99001c9e77c011dc149f7-rx_bytes"`
			UserWifi0Ath05Aa99001C9E77C011Dc149F7RxPackets float64   `json:"user-wifi0-ath0-5aa99001c9e77c011dc149f7-rx_packets"`
			UserWifi0Ath05Aa99001C9E77C011Dc149F7TxBytes   float64   `json:"user-wifi0-ath0-5aa99001c9e77c011dc149f7-tx_bytes"`
			UserWifi0Ath05Aa99001C9E77C011Dc149F7TxPackets float64   `json:"user-wifi0-ath0-5aa99001c9e77c011dc149f7-tx_packets"`
			UserWifi0RxBytes                               float64   `json:"user-wifi0-rx_bytes"`
			UserWifi0RxCrypts                              float64   `json:"user-wifi0-rx_crypts"`
			UserWifi0RxDropped                             float64   `json:"user-wifi0-rx_dropped"`
			UserWifi0RxErrors                              float64   `json:"user-wifi0-rx_errors"`
			UserWifi0RxFrags                               float64   `json:"user-wifi0-rx_frags"`
			UserWifi0RxPackets                             float64   `json:"user-wifi0-rx_packets"`
			UserWifi0TxBytes                               float64   `json:"user-wifi0-tx_bytes"`
			UserWifi0TxDropped                             float64   `json:"user-wifi0-tx_dropped"`
			UserWifi0TxErrors                              float64   `json:"user-wifi0-tx_errors"`
			UserWifi0TxPackets                             float64   `json:"user-wifi0-tx_packets"`
			UserWifi0TxRetries                             float64   `json:"user-wifi0-tx_retries"`
			UserWifi1Ath2NullRxBytes                       float64   `json:"user-wifi1-ath2-null-rx_bytes"`
			UserWifi1Ath2NullRxPackets                     float64   `json:"user-wifi1-ath2-null-rx_packets"`
			UserWifi1Ath2NullTxBytes                       float64   `json:"user-wifi1-ath2-null-tx_bytes"`
			UserWifi1Ath2NullTxPackets                     float64   `json:"user-wifi1-ath2-null-tx_packets"`
			UserWifi1RxBytes                               float64   `json:"user-wifi1-rx_bytes"`
			UserWifi1RxCrypts                              float64   `json:"user-wifi1-rx_crypts"`
			UserWifi1RxDropped                             float64   `json:"user-wifi1-rx_dropped"`
			UserWifi1RxErrors                              float64   `json:"user-wifi1-rx_errors"`
			UserWifi1RxFrags                               float64   `json:"user-wifi1-rx_frags"`
			UserWifi1RxPackets                             float64   `json:"user-wifi1-rx_packets"`
			UserWifi1TxBytes                               float64   `json:"user-wifi1-tx_bytes"`
			UserWifi1TxDropped                             float64   `json:"user-wifi1-tx_dropped"`
			UserWifi1TxErrors                              float64   `json:"user-wifi1-tx_errors"`
			UserWifi1TxPackets                             float64   `json:"user-wifi1-tx_packets"`
			UserWifi1TxRetries                             float64   `json:"user-wifi1-tx_retries"`
			Wifi0RxBytes                                   float64   `json:"wifi0-rx_bytes"`
			Wifi0RxCrypts                                  float64   `json:"wifi0-rx_crypts"`
			Wifi0RxDropped                                 float64   `json:"wifi0-rx_dropped"`
			Wifi0RxErrors                                  float64   `json:"wifi0-rx_errors"`
			Wifi0RxFrags                                   float64   `json:"wifi0-rx_frags"`
			Wifi0RxPackets                                 float64   `json:"wifi0-rx_packets"`
			Wifi0TxBytes                                   float64   `json:"wifi0-tx_bytes"`
			Wifi0TxDropped                                 float64   `json:"wifi0-tx_dropped"`
			Wifi0TxErrors                                  float64   `json:"wifi0-tx_errors"`
			Wifi0TxPackets                                 float64   `json:"wifi0-tx_packets"`
			Wifi0TxRetries                                 float64   `json:"wifi0-tx_retries"`
			Wifi1RxBytes                                   float64   `json:"wifi1-rx_bytes"`
			Wifi1RxCrypts                                  float64   `json:"wifi1-rx_crypts"`
			Wifi1RxDropped                                 float64   `json:"wifi1-rx_dropped"`
			Wifi1RxErrors                                  float64   `json:"wifi1-rx_errors"`
			Wifi1RxFrags                                   float64   `json:"wifi1-rx_frags"`
			Wifi1RxPackets                                 float64   `json:"wifi1-rx_packets"`
			Wifi1TxBytes                                   float64   `json:"wifi1-tx_bytes"`
			Wifi1TxDropped                                 float64   `json:"wifi1-tx_dropped"`
			Wifi1TxErrors                                  float64   `json:"wifi1-tx_errors"`
			Wifi1TxPackets                                 float64   `json:"wifi1-tx_packets"`
			Wifi1TxRetries                                 float64   `json:"wifi1-tx_retries"`
		} `json:"at"`
	} `json:"stat"`
	State    DevState `json:"state"`
	SysStats struct {
	} `json:"sys_stats"`
	SystemStats struct {
	} `json:"system-stats"`
	TxBytes           int64  `json:"tx_bytes"`
	TxBytesD          int64  `json:"tx_bytes-d"`
	Type              string `json:"type"`
	Upgradable        bool   `json:"upgradable"`
	UpgradeState      int    `json:"upgrade_state"`
	UpgradeToFirmware string `json:"upgrade_to_firmware"`
	Uplink            struct {
		FullDuplex  bool   `json:"full_duplex"`
		IP          string `json:"ip"`
		Mac         string `json:"mac"`
		MaxSpeed    int    `json:"max_speed"`
		Name        string `json:"name"`
		Netmask     string `json:"netmask"`
		NumPort     int    `json:"num_port"`
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
	UplinkTable []interface{} `json:"uplink_table"`
	Uptime      int           `json:"uptime"`
	UserNumSta  int           `json:"user-num_sta"`
	VapTable    []struct {
		ApMac      string `json:"ap_mac"`
		Bssid      string `json:"bssid"`
		Ccq        int    `json:"ccq"`
		Channel    int    `json:"channel"`
		Essid      string `json:"essid"`
		ID         string `json:"id"`
		IsGuest    bool   `json:"is_guest"`
		IsWep      bool   `json:"is_wep"`
		MapID      string `json:"map_id"`
		Name       string `json:"name"`
		NumSta     int    `json:"num_sta"`
		Radio      string `json:"radio"`
		RadioName  string `json:"radio_name"`
		RxBytes    int64  `json:"rx_bytes"`
		RxCrypts   int64  `json:"rx_crypts"`
		RxDropped  int64  `json:"rx_dropped"`
		RxErrors   int64  `json:"rx_errors"`
		RxFrags    int64  `json:"rx_frags"`
		RxNwids    int64  `json:"rx_nwids"`
		RxPackets  int64  `json:"rx_packets"`
		SiteID     string `json:"site_id"`
		State      string `json:"state"`
		T          string `json:"t"`
		TxBytes    int64  `json:"tx_bytes"`
		TxDropped  int64  `json:"tx_dropped"`
		TxErrors   int64  `json:"tx_errors"`
		TxPackets  int64  `json:"tx_packets"`
		TxPower    int64  `json:"tx_power"`
		TxRetries  int64  `json:"tx_retries"`
		Up         bool   `json:"up"`
		Usage      string `json:"usage"`
		WlanconfID string `json:"wlanconf_id"`
	} `json:"vap_table"`
	Version                string        `json:"version"`
	VersionIncompatible    bool          `json:"version_incompatible"`
	VwireEnabled           bool          `json:"vwireEnabled"`
	VwireTable             []interface{} `json:"vwire_table"`
	VwireVapTable          []interface{} `json:"vwire_vap_table"`
	WifiCaps               int           `json:"wifi_caps"`
	X                      fuzzyFloat    `json:"x"`
	XAuthkey               string        `json:"x_authkey"`
	XFingerprint           string        `json:"x_fingerprint"`
	XHasSSHHostkey         bool          `json:"x_has_ssh_hostkey"`
	XInformAuthkey         string        `json:"x_inform_authkey"`
	XSSHHostkeyFingerprint string        `json:"x_ssh_hostkey_fingerprint"`
	XVwirekey              string        `json:"x_vwirekey"`
	Y                      fuzzyFloat    `json:"y"`
}

type UAPmap map[string]UAP

// Returns a slice of access points
// Deprecated: Use UAPs instead
func (u *Unifi) Aps(site *Site) ([]UAP, error) {
	return u.UAPs(site)
}

// Returns a slice of access points
func (u *Unifi) UAPs(site *Site) ([]UAP, error) {

	// Devices
	var uaps []UAP

	rawDevices, err := u.RawDevices(site, "uap")

	for i, _ := range rawDevices {
		var uap UAP
		err := json.Unmarshal(rawDevices[i].Data, &uap)
		if err != nil {
			return uaps, err
		}
		// Set unifi pointer
		uap.u = u

		uaps = append(uaps, uap)
	}
	return uaps, err
}

// Returns a map of access points with mac as a key
func (u *Unifi) UAPMap(site *Site) (UAPmap, error) {
	aps, err := u.Aps(site)
	if err != nil {
		return nil, err
	}
	m := make(UAPmap)
	for _, a := range aps {
		m[a.Mac] = a
	}
	return m, nil
}

// Reboot access point
func (a UAP) Restart() error {
	if a.u == nil {
		return ErrLoginFirst
	}
	return a.u.devcmd(a.Mac, "restart")
}

func (a UAP) DeviceName() string {
	if a.Name != "" {
		return a.Name
	}
	// If no name is given, return mac as name
	return a.Mac
}

func (a UAP) ModelName() string {
	if m, ok := model[a.Model]; ok {
		return m
	}
	return a.Model
}
