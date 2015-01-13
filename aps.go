package unifi

type ConfigNetwork struct {
	IP   string
	Type string
}

type DownlinkTable struct {
	/* FIXME */
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
	FullDuplex  bool `json:"full_duplex"`
	IP          string
	Mac         string
	Name        string
	NumPort     int `json:"num_port"`
	RxBytes     int `json:"rx_bytes"`
	RxDropped   int `json:"rx_dropped"`
	RxErrors    int `json:"rx_errors"`
	RxMulticast int `json:"rx_multicast"`
	RxPackets   int `json:"rx_packets"`
	Speed       int
	TxBytes     int `json:"tx_bytes"`
	TxDropped   int `json:"tx_dropped"`
	TxErrors    int `json:"tx_errors"`
	TxPackets   int `json:"tx_packets"`
	Type        string
	Up          bool
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

type ApsMap map[string]Aps

// Access point datA
type Aps struct {
	u                  *Unifi
	Adopted            bool
	Bytes              int
	Cfgversion         string
	ConfigNetwork      ConfigNetwork   `json:"config_network"`
	ConnectRequestIP   string          `json:"connect_request_ip"`
	ConnectRequestPort string          `json:"connect_request_port"`
	ConsideredLostAt   int             `json:"considered_lost_at"`
	DownlinkTable      []DownlinkTable `json:"downlink_table"`
	GuestNumSta        int             `json:"guest-num_sta"`
	GuestToken         string          `json:"guest_token"`
	HasEth1            bool            `json:"has_eth1"`
	HasPoePassthrough  bool            `json:"has_poe_passthrough"`
	InformAuthkey      string          `json:"inform_authkey"`
	InformIP           string          `json:"inform_ip"`
	InformURL          string          `json:"inform_url"`
	IP                 string
	KnownCfgversion    string `json:"known_cfgversion"`
	LastSeen           int    `json:"last_seen"`
	Locating           bool
	Locked             bool
	Mac                string
	MapID              string `json:"map_id"`
	Model              string
	NaChannel          int    `json:"na-channel"`
	NaEirp             int    `json:"na-eirp"`
	NaExtchannel       int    `json:"na-extchannel"`
	NaGain             int    `json:"na-gain"`
	NaState            string `json:"na-state"`
	NaTxPower          string `json:"na-tx_power"`
	Name               string
	NextHeartbeatAt    int          `json:"next_heartbeat_at"`
	NgChannel          int          `json:"ng-channel"`
	NgEirp             int          `json:"ng-eirp"`
	NgExtchannel       int          `json:"ng-extchannel"`
	NgGain             int          `json:"ng-gain"`
	NgState            string       `json:"ng-state"`
	NgTxPower          string       `json:"ng-tx_power"`
	NumSta             int          `json:"num_sta"`
	RadioTable         []RadioTable `json:"radio_table"`
	RxBytes            int          `json:"rx_bytes"`
	Scanning           bool
	Serial             string
	Stat               Stat
	State              int
	TxBytes            int `json:"tx_bytes"`
	Uplink             Uplink
	UplinkTable        []UplinkTable `json:"uplink_table"`
	Uptime             int
	UserNumSta         int        `json:"user-num_sta"`
	VapTable           []VapTable `json:"vap_table"`
	Version            string
	VwireEnabled       bool
	VwireTable         []VwireTable `json:"vwire_table"`
	X                  float64
	XAuthkey           string `json:"x_authkey"`
	XFingerprint       string `json:"x_fingerprint"`
	XVwirekey          string `json:"x_vwirekey"`
	Y                  float64
}

// Reboot access point
func (a Aps) Restart() error {
	if a.u == nil {
		return ErrLoginFirst
	}
	return a.u.devcmd(a.Mac, "restart", 0)
}

var model = map[string]string{
	"BZ2":   "UniFi AP",
	"BZ2LR": "UniFi AP-LR",
	"U2S48": "UniFi AP",
	"U2L48": "UniFi AP-LR",
	"U2HSR": "UniFi AP-Outdoor+",
	"U2O":   "UniFi AP-Outdoor",
	"U5O":   "UniFi AP-Outdoor 5G",
	"U7P":   "UniFi AP-Pro",
	"U2M":   "UniFi AP-Mini",
	"p2N":   "PicoStation M2",
	"U7E":   "UniFi AP-AC",
	"U7O":   "UniFi AP-AC Outdoor",
	"U7Ev2": "UniFi AP-AC v2",
}

func (a Aps) ModelName() string {
	if m, ok := model[a.Model]; ok {
		return m
	}
	return "unknown"
}

var status = map[int]string{
	0:  "Disconnected",
	1:  "Connected",
	2:  "Pending",
	3:  "Disconnected",
	4:  "Upgrading",
	5:  "Provisioning",
	6:  "Heartbeat Missed",
	7:  "Adoping",
	8:  "Deleting",
	9:  "Managed By Others",
	10: "Adopt Failed",
	11: "Isolated",
	12: "Isolate Pending",
	13: "Wireless Adopting",
}

func (a Aps) Status() string {
	if s, ok := status[a.State]; ok {
		return s
	}
	return "unknown"
}
