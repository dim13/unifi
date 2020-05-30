// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

import (
	"encoding/json"
	"fmt"
	"time"
)

type ConfigNetwork struct {
	IP   string
	Type string
}

type DownlinkTable struct {
	/* FIXME */
}

type Port struct {
	AggregatedBy   interface{} `json:"aggregated_by"` // FIXME: type false or int (5.6.22)
	Autoneg        bool
	BytesR         int64  `json:"bytes-r"`
	Dot1xMode      string `json:"dot1x_mode"`
	Dot1xStatus    string `json:"dot1x_status"`
	Enable         bool
	FlowCtrl       bool `json:"flowctrl_tx"`
	FlowCtrlRx     bool `json:"flowctrl_rx"`
	FullDuplex     bool `json:"full_duplex"`
	IsUplink       bool `json:"is_uplink"`
	Jumbo          bool
	Masked         bool
	Media          string
	Name           string `json:"name"`
	OpMode         string `json:"op_mode"`
	PoeCaps        int    `json:"poe_caps"`
	PoeClass       string `json:"poe_class"`
	PoeCurrent     string `json:"poe_current"`
	PoeEnabled     bool   `json:"poe_enable"`
	PoeGood        bool   `json:"poe_good"`
	PoeMode        string `json:"poe_mode"`
	PoePower       string `json:"poe_power"`
	PoeVoltage     string `json:"poe_voltage"`
	PortIdx        int    `json:"port_idx"`
	PortPoe        bool   `json:"port_poe"`
	PortconfID     string `json:"portconf_id"`
	RxBroadcast    int64  `json:"rx_broadcast"`
	RxBytes        int64  `json:"rx_bytes"`
	RxBytesR       int64  `json:"rx_bytes-r"`
	RxDropped      int64  `json:"rx_dropped"`
	RxErrors       int64  `json:"rx_errors"`
	RxMulticast    int64  `json:"rx_multicast"`
	RxPackets      int64  `json:"rx_packets"`
	Speed          int
	SfpCompliance  string `json:"sfp_compliance"`
	SfpCurrent     string `json:"sfp_current"`
	SfpFound       bool   `json:"sfp_found"`
	SfpPart        string `json:"sfp_part"`
	SfpRevision    string `json:"sfp_rev"`
	SfpRxFault     bool   `json:"sfp_rxfault"`
	SfpRxPower     string `json:"sfp_rxpower"`
	SfpSerial      string `json:"sfp_serial"`
	SfpTemperature string `json:"sfp_temperature"`
	SfpTxFault     bool   `json:"sfp_txfault"`
	SfpTxPower     string `json:"sfp_txpower"`
	SfpVendor      string `json:"sfp_vendor"`
	SfpVoltage     string `json:"sfp_voltage"`
	StpPathcost    int64  `json:"stp_pathcost"`
	StpState       string `json:"stp_state"`
	TxBroadcast    int64  `json:"tx_broadcast"`
	TxMulticast    int64  `json:"tx_multicast"`
	TxBytes        int64  `json:"tx_bytes"`
	TxBytesR       int64  `json:"tx_bytes-r"`
	TxDropped      int64  `json:"tx_dropped"`
	TxErrors       int64  `json:"tx_errors"`
	TxPackets      int64  `json:"tx_packets"`
	Up             bool
}

// Unifi API Version 5.7.20
type Uplink struct {
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
}

type UplinkTable struct {
	/* FIXME */
}

type StaTable struct {
	AuthTime      int `json:"auth_time"`
	Authorized    bool
	Ccq           int
	DhcpendTime   int `json:"dhcpend_time"`
	DhcpstartTime int `json:"dhcpstart_time"`
	Hostname      string
	Idletime      int
	IP            string
	Is11b         bool `json:"is_11b"`
	Is11n         bool `json:"is_11n"`
	Mac           string
	Noise         int
	Rssi          int
	RxBytes       int `json:"rx_bytes"`
	RxPackets     int `json:"rx_packets"`
	RxRate        int `json:"rx_rate"`
	Signal        int
	State         int
	StateHt       bool `json:"state_ht"`
	StatePwrmgt   bool `json:"state_pwrmgt"`
	TxBytes       int  `json:"tx_bytes"`
	TxPackets     int  `json:"tx_packets"`
	TxPower       int  `json:"tx_power"`
	TxRate        int  `json:"tx_rate"`
	Uptime        int
}

type VwireTable struct {
	/* FIXME */
}

type Ethernet struct {
	Mac     string
	Name    string
	NumPort *int `json:"num_port"` // Pointer to check if absent
}

const (
	OpModeSwitch    = "switch" // FIXME
	OpModeAggregate = "aggregate"
	OpModeMirror    = "mirror"
)

const (
	POEModeAuto   = "auto"
	POEModeOff    = "off"
	POEModePasv24 = "pasv24"
)

type PortOverride struct {
	// Bool and int fields use pointers and flagged with omitempty
	// Thus, absent fields received in query stay absent in response.
	// Additionally, if these fields are set they are contained in they
	// response, even if they are set to their default value (e.g. 0 or false)

	AggregateNumPorts        *int   `json:"aggregate_num_ports,omitempty"`
	Autoneg                  *bool  `json:"autoneg,omitempty"`
	FullDuplex               *bool  `json:"full_duplex,omitempty"`
	Isolation                *bool  `json:"isolation,omitempty"`
	LLDPMedEnabled           *bool  `json:"lldpmed_enabled,omitempty"`
	MirrorPortIdx            string `json:"mirror_port_idx,omitempty"`
	Name                     string `json:"name,omitempty"`
	OpMode                   string `json:"op_mode,omitempty"`  // Values (switching(absent)|aggregate\mirror )
	POEMode                  string `json:"poe_mode,omitempty"` // Values ('auto'|'off'|'pasv24')
	PortIdx                  int    `json:"port_idx"`
	PortconfID               string `json:"portconf_id"`
	Speed                    *int   `json:"speed,omitempty"`
	StormcontrolBcastEnabled *bool  `json:"stormctrl_bcast_enabled,omitempty"`
	StormcontrolBcastRate    *int   `json:"stormctrl_bcast_rate,omitempty"`
	StormcontrolMcastEnabled *bool  `json:"stormctrl_mcast_enabled,omitempty"`
	StormcontrolMcastRate    *int   `json:"stormctrl_mcast_rate,omitempty"`
	StormcontrolUcastEnabled *bool  `json:"stormctrl_ucast_enabled,omitempty"`
	StormcontrolUcastRate    *int   `json:"stormctrl_ucast_rate,omitempty"`
}

// API version 5.12.35
type Stat struct {
	Ap Ap `json:"ap"`
}

type Ap struct {
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
}

type SysStats struct {
	MemBuffer int    `json:"mem_buffer"`
	MemUsed   int    `json:"mem_used"`
	LoadAvg1  string `json:"loadavg_1"`
	LoadAvg5  string `json:"loadavg_5"`
	LoadAvg15 string `json:"loadavg_15"`
	MemTotal  int    `json:"mem_total"`
}

type RadioTableStats struct {
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
	TxPackets   int         `json:"tx_packets"`
	TxPower     int         `json:"tx_power"`
	TxRetries   int         `json:"tx_retries"`
	UserNumSta  int         `json:"user-num_sta"`
}

type SystemStats struct {
	SystemStats struct {
		CPU   string `json:"cpu"`
		Mem   string `json:"mem"`
		Temps struct {
			BoardCPU string `json:"Board (CPU)"`
			BoardPHY string `json:"Board (PHY)"`
			CPU      string `json:"CPU"`
			PHY      string `json:"PHY"`
		} `json:"temps"`
	} `json:"system-stats"`
}

type ConfigNetworkWan struct {
	DNS1    string `json:"dns1"`
	DNS2    string `json:"dns2"`
	Gateway string `json:"gateway"`
	IP      string `json:"ip"`
	Netmask string `json:"netmask"`
	Type    string `json:"type"`
}

type Wan struct {
	BytesR      int      `json:"bytes-r"`
	DNS         []string `json:"dns"`
	Enable      bool     `json:"enable"`
	FullDuplex  bool     `json:"full_duplex"`
	Gateway     string   `json:"gateway"`
	Ifname      string   `json:"ifname"`
	IP          string   `json:"ip"`
	Mac         string   `json:"mac"`
	MaxSpeed    int      `json:"max_speed"`
	Name        string   `json:"name"`
	Netmask     string   `json:"netmask"`
	RxBytes     int64    `json:"rx_bytes"`
	RxBytesR    int      `json:"rx_bytes-r"`
	RxDropped   int      `json:"rx_dropped"`
	RxErrors    int      `json:"rx_errors"`
	RxMulticast int      `json:"rx_multicast"`
	RxPackets   int      `json:"rx_packets"`
	Speed       int      `json:"speed"`
	TxBytes     int64    `json:"tx_bytes"`
	TxBytesR    int      `json:"tx_bytes-r"`
	TxDropped   int      `json:"tx_dropped"`
	TxErrors    int      `json:"tx_errors"`
	TxPackets   int      `json:"tx_packets"`
	Type        string   `json:"type"`
	Up          bool     `json:"up"`
}

type RawDevice struct {
	Data json.RawMessage
	Type string
}

// Returns a slice of json RawDevices as received by the controller
func (u *Unifi) RawDevices(site *Site, filter string) ([]RawDevice, error) {

	var rawDevices []RawDevice

	// Response from controller
	var response struct {
		Data []json.RawMessage
		Meta meta
	}
	err := u.parse(site, "stat/device", nil, &response)

	for _, d := range response.Data {

		// unmarshal into a map to check the "type" field
		var obj map[string]interface{}
		err := json.Unmarshal(d, &obj)
		if err != nil {
			return nil, err
		}

		deviceType, ok := obj["type"].(string)
		if !ok {
			return nil, fmt.Errorf("Error on retrieving object type from raw Json")
		}

		switch filter {
		case "":
			var rd RawDevice
			rd.Type = deviceType
			rd.Data = d
			rawDevices = append(rawDevices, rd)
		default:
			// Filter is set. Only return devices of the given type
			if deviceType == filter {
				var rd RawDevice
				rd.Type = deviceType
				rd.Data = d
				rawDevices = append(rawDevices, rd)
			}
		}
	}

	return rawDevices, err
}

type Device struct {
	u  *Unifi
	ID string `json:"_id"`
	//Uptime             int           `json:"_uptime"`
	Adopted          bool          `json:"adopted"`
	AntennaTable     []interface{} `json:"antenna_table,omitempty"`
	Bytes            int           `json:"bytes"`
	BytesD           int           `json:"bytes-d,omitempty"`
	BytesR           int           `json:"bytes-r,omitempty"`
	Cfgversion       string        `json:"cfgversion"`
	ConfigNetwork    ConfigNetwork `json:"config_network"`
	ConnectRequestIP string        `json:"connect_request_ip"`
	//ConnectRequestPort int           `json:"connect_request_port"` //FIXME: Prior to 5.7.20 type string!
	ConsideredLostAt int           `json:"considered_lost_at"`
	DeviceID         string        `json:"device_id"`
	EthernetTable    []Ethernet    `json:"ethernet_table"`
	FwCaps           int           `json:"fw_caps"`
	GuestNumSta      int           `json:"guest-num_sta"`
	GuestToken       string        `json:"guest_token,omitempty"`
	HasEth1          bool          `json:"has_eth1,omitempty"`
	HasSpeaker       bool          `json:"has_speaker,omitempty"`
	InformIP         string        `json:"inform_ip"`
	InformURL        string        `json:"inform_url"`
	IP               string        `json:"ip"`
	Isolated         bool          `json:"isolated,omitempty"`
	KnownCfgversion  string        `json:"known_cfgversion"`
	LastSeen         int           `json:"last_seen"`
	LedOverride      string        `json:"led_override"`
	LicenseState     string        `json:"license_state"`
	Locating         bool          `json:"locating"`
	Mac              string        `json:"mac"`
	MapID            string        `json:"map_id,omitempty"`
	Model            string        `json:"model"`
	Name             string        `json:"name"`
	NextHeartbeatAt  int           `json:"next_heartbeat_at"`
	NumSta           int           `json:"num_sta"`
	PortStats        []interface{} `json:"port_stats,omitempty"`
	PortTable        []Port        `json:"port_table"`
	RadioTable       []struct {
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
	} `json:"radio_table,omitempty"`
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
		TxPackets   int         `json:"tx_packets"`
		TxPower     int         `json:"tx_power"`
		TxRetries   int         `json:"tx_retries"`
		UserNumSta  int         `json:"user-num_sta"`
	} `json:"radio_table_stats,omitempty"`
	RxBytes           int           `json:"rx_bytes"`
	RxBytesD          int           `json:"rx_bytes-d,omitempty"`
	ScanRadioTable    []interface{} `json:"scan_radio_table,omitempty"`
	Scanning          bool          `json:"scanning,omitempty"`
	Serial            string        `json:"serial"`
	SiteID            string        `json:"site_id"`
	SpectrumScanning  bool          `json:"spectrum_scanning,omitempty"`
	SSHSessionTable   []interface{} `json:"ssh_session_table,omitempty"`
	Stat              Stat          `json:"stat"`
	State             DevState      `json:"state"`
	SysStats          SysStats      `json:"sys_stats"`
	SystemStats       SystemStats   `json:"system-stats"`
	TxBytes           int           `json:"tx_bytes"`
	TxBytesD          int           `json:"tx_bytes-d,omitempty"`
	Type              string        `json:"type"`
	Upgradable        bool          `json:"upgradable"`
	UpgradeState      int           `json:"upgrade_state"`
	UpgradeToFirmware string        `json:"upgrade_to_firmware"`
	Uplink            Uplink        `json:"uplink"`
	UplinkTable       []interface{} `json:"uplink_table,omitempty"`
	Uptime            int           `json:"uptime"`
	UserNumSta        int           `json:"user-num_sta"`
	VapTable          []struct {
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
		RxBytes    int    `json:"rx_bytes"`
		RxCrypts   int    `json:"rx_crypts"`
		RxDropped  int    `json:"rx_dropped"`
		RxErrors   int    `json:"rx_errors"`
		RxFrags    int    `json:"rx_frags"`
		RxNwids    int    `json:"rx_nwids"`
		RxPackets  int    `json:"rx_packets"`
		SiteID     string `json:"site_id"`
		State      string `json:"state"`
		T          string `json:"t"`
		TxBytes    int    `json:"tx_bytes"`
		TxDropped  int    `json:"tx_dropped"`
		TxErrors   int    `json:"tx_errors"`
		TxPackets  int    `json:"tx_packets"`
		TxPower    int    `json:"tx_power"`
		TxRetries  int    `json:"tx_retries"`
		Up         bool   `json:"up"`
		Usage      string `json:"usage"`
		WlanconfID string `json:"wlanconf_id"`
	} `json:"vap_table,omitempty"`
	Version                string           `json:"version"`
	VersionIncompatible    bool             `json:"version_incompatible"`
	VwireEnabled           bool             `json:"vwireEnabled,omitempty"`
	VwireTable             []interface{}    `json:"vwire_table,omitempty"`
	VwireVapTable          []interface{}    `json:"vwire_vap_table,omitempty"`
	WifiCaps               int              `json:"wifi_caps,omitempty"`
	X                      fuzzyFloat       `json:"x,omitempty"`
	XAuthkey               string           `json:"x_authkey"`
	XFingerprint           string           `json:"x_fingerprint"`
	XHasSSHHostkey         bool             `json:"x_has_ssh_hostkey"`
	XInformAuthkey         string           `json:"x_inform_authkey"`
	XSSHHostkeyFingerprint string           `json:"x_ssh_hostkey_fingerprint"`
	XVwirekey              string           `json:"x_vwirekey,omitempty"`
	Y                      fuzzyFloat       `json:"y,omitempty"`
	DhcpServerTable        []interface{}    `json:"dhcp_server_table,omitempty"`
	Dot1XPortctrlEnabled   bool             `json:"dot1x_portctrl_enabled,omitempty"`
	FlowctrlEnabled        bool             `json:"flowctrl_enabled,omitempty"`
	GeneralTemperature     int              `json:"general_temperature,omitempty"`
	HasFan                 bool             `json:"has_fan,omitempty"`
	HasTemperature         bool             `json:"has_temperature,omitempty"`
	JumboframeEnabled      bool             `json:"jumboframe_enabled,omitempty"`
	Overheating            bool             `json:"overheating,omitempty"`
	StpPriority            string           `json:"stp_priority,omitempty"`
	StpVersion             string           `json:"stp_version,omitempty"`
	UplinkDepth            int              `json:"uplink_depth,omitempty"`
	ConfigNetworkWan       ConfigNetworkWan `json:"config_network_wan,omitempty"`
	ConfigNetworkWan2      ConfigNetworkWan `json:"config_network_wan2,omitempty"`
	NetworkTable           []interface{}    `json:"network_table,omitempty"`
	NumDesktop             int              `json:"num_desktop,omitempty"`
	NumHandheld            int              `json:"num_handheld,omitempty"`
	NumMobile              int              `json:"num_mobile,omitempty"`
	SpeedtestStatus        struct {
		Latency        float64 `json:"latency"` // NOTE: Prior to 5.8 type string
		Rundate        int     `json:"rundate"`
		Runtime        int     `json:"runtime"`
		StatusDownload int     `json:"status_download"`
		StatusPing     int     `json:"status_ping"`
		StatusSummary  int     `json:"status_summary"`
		StatusUpload   int     `json:"status_upload"`
		XputDownload   float64 `json:"xput_download"`
		XputUpload     float64 `json:"xput_upload"`
	} `json:"speedtest-status,omitempty"`
	SpeedtestStatusSaved bool           `json:"speedtest-status-saved,omitempty"`
	UsgCaps              int            `json:"usg_caps,omitempty"`
	Wan1                 Wan            `json:"wan1,omitempty"`
	Wan2                 Wan            `json:"waWann2,omitempty"`
	PortOverrides        []PortOverride `json:"port_overrides"`
}

type DeviceMap map[string]Device

// Returns a slice of devices
func (u *Unifi) Devices(site *Site, filter string) ([]Device, error) {

	// Devices
	var devices []Device

	rawDevices, err := u.RawDevices(site, filter)

	for i, _ := range rawDevices {
		var device *Device
		err := json.Unmarshal(rawDevices[i].Data, &device)
		if err != nil {
			return devices, err
		}

		// Set unifi pointer
		device.u = u

		devices = append(devices, *device)
	}
	return devices, err
}

// Returns a map of access points with mac as a key
func (u *Unifi) DeviceMap(site *Site) (DeviceMap, error) {
	devices, err := u.Devices(site, "")
	if err != nil {
		return nil, err
	}
	m := make(DeviceMap)
	for _, d := range devices {
		m[d.Mac] = d
	}
	return m, nil
}

func (d Device) DeviceName() string {
	if d.Name != "" {
		return d.Name
	}
	// If no name is given, return mac as name
	return d.Mac
}

func (d Device) ModelName() string {
	if m, ok := model[d.Model]; ok {
		return m
	}
	return "unknown " + d.Model
}

var model = map[string]string{
	"BZ2":      "UniFi AP",
	"BZ2LR":    "UniFi AP-LR",
	"U2HSR":    "UniFi AP-Outdoor+",
	"U2IW":     "UniFi AP-In Wall",
	"U2L48":    "UniFi AP-LR",
	"U2M":      "UniFi AP-Mini",
	"U2O":      "UniFi AP-Outdoor",
	"U2S48":    "UniFi AP",
	"U5O":      "UniFi AP-Outdoor 5G",
	"U7E":      "UniFi AP-AC",
	"U7EDU":    "UniFi AP-AC-EDU",
	"U7Ev2":    "UniFi AP-AC v2",
	"U7HD":     "UniFi UAP-AC-HD",
	"U7LO":     "UniFi AP-AC-Pro-Outdoor",
	"U7LR":     "UniFi AP-AC-LR",
	"U7LT":     "UniFi AP-AC-Lite",
	"U7MP":     "UniFi UAP-AC-M-PRO",
	"U7MSH":    "UniFi UAP-AC-M",
	"U7O":      "UniFi AP-AC Outdoor",
	"U7P":      "UniFi AP-Pro",
	"U7PC":     "UniFi AP-AC-Pico",
	"U7PG2":    "UniFi AP-AC-Pro Gen2",
	"UGW3":     "UniFi Security Gateway",
	"UGW4":     "UniFi Security Gateway-Pro",
	"UP4":      "UniFi Phone-X",
	"UP5":      "UniFi Phone",
	"UP5c":     "UniFi Phone",
	"UP5t":     "UniFi Phone-Pro",
	"UP5tc":    "UniFi Phone-Pro",
	"UP7":      "UniFi Phone-Executive",
	"UP7c":     "UniFi Phone-Executive",
	"US16P150": "UniFi Switch 16 POE-150W",
	"US24":     "UniFi Switch 24",
	"US24P250": "UniFi Switch 24 POE-250W",
	"US24P500": "UniFi Switch 24 POE-500W",
	"US48":     "UniFi Switch 48",
	"US48P500": "UniFi Switch 48 POE-500W",
	"US48P750": "UniFi Switch 48 POE-750W",
	"US8P60":   "UniFi Switch 8 POE-60W",
	"US8P150":  "UniFi Switch 8 POE-150W",
	"USXG":     "Ubiquiti UniFi US-16-XG",
	"p2N":      "PicoStation M2",
}

type DevState int

//go:generate stringer -type=DevState

const (
	Disconnected     DevState = 0
	Connected        DevState = 1
	Pending          DevState = 2
	Disconnecting    DevState = 3
	Upgrading        DevState = 4
	Provisioning     DevState = 5
	HeartbeatMissing DevState = 6
	Adopting         DevState = 7
	Deleting         DevState = 8
	ManagedByOthers  DevState = 9
	AdoptionFailed   DevState = 10
	Isolated         DevState = 11
	IsolationPending DevState = 12
	AdoptingWireless DevState = 13
)
