// Copyright (c) 2014 Dimitri Sokolyuk. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

type ConfigNetwork struct {
	IP   string
	Type string
}

type DownlinkTable struct {
	/* FIXME */
}

type Port struct {
	AggregatedBy   bool `json:"aggregated_by"`
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
	PortconfId     string `json:"portconf_id"`
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
	TxBytes        int64  `json"tx_bytes"`
	TxBytesR       int64  `json:"tx_bytes-r"`
	TxDropped      int64  `json:"tx_dropped"`
	TxErrors       int64  `json:"tx_errors"`
	TxPackets      int64  `json:"tx_packets"`
	Up             bool
}

type RadioTable struct {
	AntennaGain    int  `json:"-"` // FIXME: buggy input, sometimes string and sometimes int
	BuiltinAntGain int  `json:"builtin_ant_gain"`
	BuiltinAntenna bool `json:"builtin_antenna"`
	Channel        string
	Ht             string
	MaxTxpower     int `json:"max_txpower"`
	Mode           string
	Name           string
	Radio          string
	TxPower        string `json:"tx_power"`
	TxPowerMode    string `json:"tx_power_mode"`
}

type Uplink struct {
	FullDuplex       bool `json:"full_duplex"`
	Ip               string
	LagMember        bool `json:"lag_member"`
	Mac              string
	MaxSpeed         int64 `json:"max_speed"`
	Name             string
	Netmask          string
	NumPort          int64 `json:"num_port"`
	RxBytes          int64 `json:"rx_bytes"`
	RxBytesR         int64 `json:"rx_bytes.r"`
	RxDropped        int64 `json:"rx_dropped"`
	RxErrors         int64 `json:"rx_errors"`
	RxMulticast      int64 `json:"rx_multicast"`
	RxPackets        int64 `json:"rx_packets"`
	Speed            int
	TxBytes          int64 `json:"tx_bytes"`
	TxDropped        int64 `json:"tx_dropped"`
	TxErrors         int64 `json:"tx_errors"`
	TxPackets        int64 `json:"tx_packets"`
	Type             string
	Up               bool
	UplinkMac        string `json:"uplink_mac"`
	UplinkRemotePort int64  `json:"uplink_remote_port"`
}

type UplinkTable struct {
	/* FIXME */
}

type VapTable struct {
	ApMac      string `json:"ap_mac"`
	Bssid      string
	Ccq        int
	Channel    int
	Essid      string
	ID         string
	IsGuest    bool   `json:"is_guest"`
	IsWep      bool   `json:"is_wep"`
	MapID      string `json:"map_id"`
	Name       string
	NumSta     int `json:"num_sta"`
	Radio      string
	RxBytes    int        `json:"rx_bytes"`
	RxCrypts   int        `json:"rx_crypts"`
	RxDropped  int        `json:"rx_dropped"`
	RxErrors   int        `json:"rx_errors"`
	RxFrags    int        `json:"rx_frags"`
	RxNwids    int        `json:"rx_nwids"`
	RxPackets  int        `json:"rx_packets"`
	StaTable   []StaTable `json:"sta_table"`
	State      string
	T          string
	TxBytes    int `json:"tx_bytes"`
	TxDropped  int `json:"tx_dropped"`
	TxErrors   int `json:"tx_errors"`
	TxPackets  int `json:"tx_packets"`
	TxPower    int `json:"tx_power"`
	TxRetries  int `json:"tx_retries"`
	Up         bool
	Usage      string
	WlanconfID string `json:"wlanconf_id"`
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

type EthernetTable struct {
	Mac     string
	Name    string
	NumPort *int `json:"num_port"` // Pointer to check if absent
}

const (
	OPMODE_SWITCH    = "switch" // FIXME
	OPMODE_AGGREGATE = "aggregate"
	OPMODE_MIRROR    = "mirror"
	POEMODE_AUTO     = "auto"
	POEMODE_OFF      = "off"
	POEMODE_PASV24   = "pasv24"
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

type Stat struct {
	Ap    string
	Bytes int
	/*
		Na-numSta  int	`json:"na-num_sta"`
		Na-rxBytes  int	`json:"na-rx_bytes"`
		Na-rxFrags  int	`json:"na-rx_frags"`
		Na-rxPackets  int	`json:"na-rx_packets"`
		Na-timeDelta  int	`json:"na-time_delta"`
		Na-txBytes  int	`json:"na-tx_bytes"`
		Na-txErrors  int	`json:"na-tx_errors"`
		Na-txPackets  int	`json:"na-tx_packets"`
		Na-txRetries  int	`json:"na-tx_retries"`
		Ng-numSta  int	`json:"ng-num_sta"`
		Ng-rxBytes  int	`json:"ng-rx_bytes"`
		Ng-rxFrags  int	`json:"ng-rx_frags"`
		Ng-rxPackets  int	`json:"ng-rx_packets"`
		Ng-timeDelta  int	`json:"ng-time_delta"`
		Ng-txBytes  int	`json:"ng-tx_bytes"`
		Ng-txErrors  int	`json:"ng-tx_errors"`
		Ng-txPackets  int	`json:"ng-tx_packets"`
		Ng-txRetries  int	`json:"ng-tx_retries"`
	*/
	NumSta    int `json:"num_sta"`
	O         string
	RxBytes   int `json:"rx_bytes"`
	RxFrags   int `json:"rx_frags"`
	RxPackets int `json:"rx_packets"`
	TimeDelta int `json:"time_delta"`
	TxBytes   int `json:"tx_bytes"`
	TxErrors  int `json:"tx_errors"`
	TxPackets int `json:"tx_packets"`
	TxRetries int `json:"tx_retries"`
	Type      string
	/*
		Uplink-rxBytes  int	`json:"uplink-rx_bytes"`
		Uplink-rxPackets  int	`json:"uplink-rx_packets"`
		Uplink-timeDelta  int	`json:"uplink-time_delta"`
		Uplink-txBytes  int	`json:"uplink-tx_bytes"`
		Uplink-txPackets  int	`json:"uplink-tx_packets"`
		User-na-numSta  int	`json:"user-na-num_sta"`
		User-na-rxBytes  int	`json:"user-na-rx_bytes"`
		User-na-rxFrags  int	`json:"user-na-rx_frags"`
		User-na-rxPackets  int	`json:"user-na-rx_packets"`
		User-na-timeDelta  int	`json:"user-na-time_delta"`
		User-na-txBytes  int	`json:"user-na-tx_bytes"`
		User-na-txErrors  int	`json:"user-na-tx_errors"`
		User-na-txPackets  int	`json:"user-na-tx_packets"`
		User-na-txRetries  int	`json:"user-na-tx_retries"`
		User-ng-numSta  int	`json:"user-ng-num_sta"`
		User-ng-rxBytes  int	`json:"user-ng-rx_bytes"`
		User-ng-rxFrags  int	`json:"user-ng-rx_frags"`
		User-ng-rxPackets  int	`json:"user-ng-rx_packets"`
		User-ng-timeDelta  int	`json:"user-ng-time_delta"`
		User-ng-txBytes  int	`json:"user-ng-tx_bytes"`
		User-ng-txErrors  int	`json:"user-ng-tx_errors"`
		User-ng-txPackets  int	`json:"user-ng-tx_packets"`
		User-ng-txRetries  int	`json:"user-ng-tx_retries"`
		User-numSta  int	`json:"user-num_sta"`
		User-rxBytes  int	`json:"user-rx_bytes"`
		User-rxFrags  int	`json:"user-rx_frags"`
		User-rxPackets  int	`json:"user-rx_packets"`
		User-timeDelta  int	`json:"user-time_delta"`
		User-txBytes  int	`json:"user-tx_bytes"`
		User-txErrors  int	`json:"user-tx_errors"`
		User-txPackets  int	`json:"user-tx_packets"`
		User-txRetries  int	`json:"user-tx_retries"`
	*/
}

type SysStats struct {
	MemBuffer int    `json:"mem_buffer"`
	MemUsed   int    `json:"mem_used"`
	LoadAvg1  string `json:"loadavg_1"`
	LoadAvg5  string `json:"loadavg_5"`
	LoadAvg15 string `json:"loadavg_15"`
	MemTotal  int    `json:"mem_total"`
}

type UAPmap map[string]UAP

type GenericDevice struct {
	u                  *Unifi
	Adopted            bool
	BoardRev           int `json:"board_rev"`
	Bytes              int
	Cfgversion         string
	ConfigNetwork      ConfigNetwork `json:"config_network"`
	ConnectRequestIP   string        `json:"connect_request_ip"`
	ConnectRequestPort string        `json:"connect_request_port"`
	ConsideredLostAt   int           `json:"considered_lost_at"`
	// DHCPServerTable  // TODO
	DeviceID            string `json:"device_id"`
	Disabled            bool
	DownlinkTable       []DownlinkTable `json:"downlink_table"`
	FwCaps              int             `json:"fw_caps"`
	GuestNumSta         int             `json:"guest-num_sta"`
	InformURL           string          `json:"inform_url"`
	InformIP            string          `json:"inform_ip"`
	IP                  string
	KnownCfgversion     string `json:"known_cfgversion"`
	LastSeen            int    `json:"last_seen"`
	LedOverride         string `json:"led_override"`
	LicenseState        string `json:"license_state"`
	Locating            bool
	Mac                 string
	MgmtNetworkID       string `json:"mgmt_network_id"`
	Model               string
	Name                string
	NextHeartbeatAt     int            `json:"next_heartbeat_at"`
	NumSta              int            `json:"num_sta"`
	OutdoorModeOverride string         `json:"outdoor_mode_override"`
	PortOverrides       []PortOverride `json:"port_overrides"`
	PortTable           []Port         `json:"port_table"`
	RxBytes             int            `json:"rx_bytes"`
	Serial              string
	SiteId              string `json:"site_id"`
	State               int
	SysStats            SysStats `json:"sys_stats"` // Seems to be the replacement of system-stats
	TxBytes             int      `json:"tx_bytes"`
	UpgradeToFirmware   string   `json:"upgrade_to_firmware"` // Optional field, only available if Upgradeable = True
	Uplink              Uplink
	Uptime              int64
	_Uptime             int `json:_uptime` // Uptime value of system-stats, different than uptime
	Version             string
	Type                string `json:"type"`
}

type DeviceMap map[string]interface{}

type USW struct {
	*GenericDevice
	Dot1PortControlEnabled bool `json:"dot1x_portctrl_enabled"`
	//EthernetTable          EthernetTable `json:"ethernet_table"`
	FlowctrlEnabled    bool `json:"flowctrl_enabled"`
	FanLevel           int  `json:"fan_level"`
	GeneralTemperature *int `json:"general_temperature"`
	HasFan             bool `json:"has_fan"`
	JumboframeEnabled  bool `json:"jumboframe_enabled"`
	Overheating        bool
	StpVersion         string `json:"stp_version"`
	StpPriority        string `json:"stp_priority"`
}

// Access point data
type UAP struct {
	*GenericDevice
	// TODO Fix Fields (Generic or UAP only)
	BytesD            int    `json:"bytes-d"`
	BytesR            int    `json:"bytes-r"`
	GuestToken        string `json:"guest_token"`
	HasEth1           bool   `json:"has_eth1"`
	HasPoePassthrough bool   `json:"has_poe_passthrough"`
	HasSpeaker        bool   `json:"has_speaker"`
	InformAuthkey     string `json:"inform_authkey"`
	Isolated          bool
	Locked            bool
	MapID             string       `json:"map_id"`
	NaChannel         int          `json:"na-channel"`
	NaCuSelfRx        int          `json:"na_cu_self_rx"`
	NaCuSelfTx        int          `json:"na_cu_self_tx"`
	NaCuTotal         int          `"json:na-cu-total"`
	NaAstBeXmint      int          `json:"na_ast_be_xmit"`
	NaEirp            int          `json:"na-eirp"`
	NaExtchannel      int          `json:"na-extchannel"`
	NaGain            int          `json:"na-gain"`
	NaGuestNumSta     int          `json:"na-guest-num_sta"`
	NaNumSta          int          `json:"na-num_sta"`
	NaState           string       `json:"na-state"`
	NaTxPower         int          `json:"na-tx_power"`
	NaTxRetries       int          `json:"na_tx_retries"`
	NaTxPackets       int          `json:"na_tx_packets"`
	NgAstBeXmit       int          `json:"ng_ast_be_xmit"`
	NgChannel         int          `json:"ng-channel"`
	NgCuSelfTx        int          `json:"ng_cu_self_tx"`
	NgGuestNumSta     int          `json:"ng-guest-num_sta"`
	NgEirp            int          `json:"ng-eirp"`
	NgExtchannel      int          `json:"ng-extchannel"`
	NgGain            int          `json:"ng-gain"`
	NgNumSta          int          `json:"ng-num_sta"`
	NgState           string       `json:"ng-state"`
	NgTxPackets       int          `json:"ng_tx_packets"`
	NgTxPower         int          `json:"ng-tx_power"`
	NgUserNumSta      int          `json:"ng-user-num-sta"`
	RadioTable        []RadioTable `json:"radio_table"`
	Rollupgrade       bool         `json:"rollupgrade"`
	RxBytesD          int          `json:"rx_bytes-d"`
	Scanning          bool
	SpectrumScanning  bool `spectrum_scanning`
	TxBytesD          int  `json:"tx_bytes-d"`
	Type              string
	Upgradeable       bool
	UplinkTable       []UplinkTable `json:"uplink_table"`
	UserNumSta        int           `json:"user-num_sta"`
	VapTable          []VapTable    `json:"vap_table"`
	VwireEnabled      bool
	VwireTable        []VwireTable `json:"vwire_table"`
	WifiCaps          int          `json:"wifi_caps"`
	WlanGroupIdNa     string       `json:"wlangroup_id_na"`
	WlanGroupIdNg     string       `json:"wlangroup_id_na"`
	XAuthkey          string       `json:"x_authkey"`
	XFingerprint      string       `json:"x_fingerprint"`
	XVwirekey         string       `json:"x_vwirekey"`
	_Id               string       `json:"_id"`
}

var model = map[string]string{
	// APs
	"BZ2":     "UniFi AP",
	"BZ2LR":   "UniFi AP-LR",
	"U2S48":   "UniFi AP",
	"U2L48":   "UniFi AP-LR",
	"U2HSR":   "UniFi AP-Outdoor+",
	"U2O":     "UniFi AP-Outdoor",
	"U5O":     "UniFi AP-Outdoor 5G",
	"U7P":     "UniFi AP-Pro",
	"U2M":     "UniFi AP-Mini",
	"p2N":     "PicoStation M2",
	"U7E":     "UniFi AP-AC",
	"U7O":     "UniFi AP-AC Outdoor",
	"U7Ev2":   "UniFi AP-AC v2",
	"U7HD":    "UniFi UAP-AC-HD",
	"U7MSH":   "UniFi UAP-AC-M",
	"U7MSH??": "UniFi UAP-AC-M-PRO",
	"U7PG2":   "UniFi UAP-AC-PRO",
	// Switches
	"US8P150":  "UniFi US-8-150W",
	"US16P150": "UniFi US-16-150W",
	"US24P250": "UniFi US-24-250W",
	"USXG":     "Ubiquiti UniFi US-16-XG",
}

// Reboot access point
func (a UAP) Restart() error {
	if a.u == nil {
		return ErrLoginFirst
	}
	return a.u.devcmd(a.Mac, "restart")
}

// Reboot switch
func (s USW) Restart() error {
	if s.u == nil {
		return ErrLoginFirst
	}
	return s.u.devcmd(s.Mac, "restart")
}

func (a UAP) ModelName() string {
	if m, ok := model[a.Model]; ok {
		return m
	}
	return "unknown"
}

func (s USW) ModelName() string {
	if m, ok := model[s.Model]; ok {
		return m
	}
	return s.Model
}

func (a UAP) DeviceName() string {
	if a.Name != "" {
		return a.Name
	}
	// If no name is given, return mac as name
	return a.Mac
}

func (s USW) DeviceName() string {
	if s.Name != "" {
		return s.Name
	}
	// If no name is given, return mac as name
	return s.Mac
}

const DISCONNECTED = 0
const CONNECTED = 1
const PENDING = 2
const DISCONNECTING = 3
const UPGRADING = 4
const PROVISIONING = 5
const HEARTBEATMISSING = 6
const ADOPTING = 7
const DELETING = 8
const MANAGEDBYOTHERS = 9
const ADOPTIONFAILED = 10
const ISOLATED = 11
const ISOLATIONPENDING = 12
const ADOPTINGWIRELESS = 13

var status = map[int]string{
	0:  "Disconnected",
	1:  "Connected",
	2:  "Pending",
	3:  "Disconnnecting", // ?
	4:  "Upgrading",
	5:  "Provisioning",
	6:  "Heartbeat Missing",
	7:  "Adoping",
	8:  "Deleting",
	9:  "Managed By Others",
	10: "Adoption Failed",
	11: "Isolated",
	12: "Isolate Pending",
	13: "Adopting (Wireless)",
}

func (a UAP) Status() string {
	if s, ok := status[a.State]; ok {
		return s
	}
	return "unknown"
}

func (a USW) Status() string {
	if s, ok := status[a.State]; ok {
		return s
	}
	return "unknown"
}

func (a UAP) SetU(u *Unifi) {
	if u != nil {
		a.u = u
	}
}
