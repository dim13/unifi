package unifi

import "log"

type ConfigNetwork struct {
	Ip   string
	Type string
}

type DownlinkTable struct {
	/* FIXME */
}

type RadioTable struct {
	Antenna_gain     int `json:"-"` // FIXME: buggy input, sometimes string and sometimes int
	Builtin_ant_gain int
	Builtin_antenna  bool
	Channel          string
	Ht               string
	Max_txpower      int
	Mode             string
	Name             string
	Radio            string
	Tx_power         string
	Tx_power_mode    string
}

type Uplink struct {
	Full_duplex  bool
	Ip           string
	Mac          string
	Name         string
	Num_port     int
	Rx_bytes     int
	Rx_dropped   int
	Rx_errors    int
	Rx_multicast int
	Rx_packets   int
	Speed        int
	Tx_bytes     int
	Tx_dropped   int
	Tx_errors    int
	Tx_packets   int
	Type         string
	Up           bool
}

type UplinkTable struct {
	/* FIXME */
}

type VapTable struct {
	Ap_mac      string
	Bssid       string
	Ccq         int
	Channel     int
	Essid       string
	Id          string
	Is_guest    bool
	Is_wep      bool
	Map_id      string
	Name        string
	Num_sta     int
	Radio       string
	Rx_bytes    int
	Rx_crypts   int
	Rx_dropped  int
	Rx_errors   int
	Rx_frags    int
	Rx_nwids    int
	Rx_packets  int
	Sta_table   []StaTable
	State       string
	T           string
	Tx_bytes    int
	Tx_dropped  int
	Tx_errors   int
	Tx_packets  int
	Tx_power    int
	Tx_retries  int
	Up          bool
	Usage       string
	Wlanconf_id string
}

type StaTable struct {
	Auth_time      int
	Authorized     bool
	Ccq            int
	Dhcpend_time   int
	Dhcpstart_time int
	Hostname       string
	Idletime       int
	Ip             string
	Is_11b         bool
	Is_11n         bool
	Mac            string
	Noise          int
	Rssi           int
	Rx_bytes       int
	Rx_packets     int
	Rx_rate        int
	Signal         int
	State          int
	State_ht       bool
	State_pwrmgt   bool
	Tx_bytes       int
	Tx_packets     int
	Tx_power       int
	Tx_rate        int
	Uptime         int
}

type VwireTable struct {
	/* FIXME */
}

type Stat struct {
	Ap    string
	Bytes int
	/*
		Na-num_sta  int
		Na-rx_bytes  int
		Na-rx_frags  int
		Na-rx_packets  int
		Na-time_delta  int
		Na-tx_bytes  int
		Na-tx_errors  int
		Na-tx_packets  int
		Na-tx_retries  int
		Ng-num_sta  int
		Ng-rx_bytes  int
		Ng-rx_frags  int
		Ng-rx_packets  int
		Ng-time_delta  int
		Ng-tx_bytes  int
		Ng-tx_errors  int
		Ng-tx_packets  int
		Ng-tx_retries  int
	*/
	Num_sta    int
	O          string
	Rx_bytes   int
	Rx_frags   int
	Rx_packets int
	Time_delta int
	Tx_bytes   int
	Tx_errors  int
	Tx_packets int
	Tx_retries int
	Type       string
	/*
		Uplink-rx_bytes  int
		Uplink-rx_packets  int
		Uplink-time_delta  int
		Uplink-tx_bytes  int
		Uplink-tx_packets  int
		User-na-num_sta  int
		User-na-rx_bytes  int
		User-na-rx_frags  int
		User-na-rx_packets  int
		User-na-time_delta  int
		User-na-tx_bytes  int
		User-na-tx_errors  int
		User-na-tx_packets  int
		User-na-tx_retries  int
		User-ng-num_sta  int
		User-ng-rx_bytes  int
		User-ng-rx_frags  int
		User-ng-rx_packets  int
		User-ng-time_delta  int
		User-ng-tx_bytes  int
		User-ng-tx_errors  int
		User-ng-tx_packets  int
		User-ng-tx_retries  int
		User-num_sta  int
		User-rx_bytes  int
		User-rx_frags  int
		User-rx_packets  int
		User-time_delta  int
		User-tx_bytes  int
		User-tx_errors  int
		User-tx_packets  int
		User-tx_retries  int
	*/
}

type ApsMap map[string]Aps

// Access point datA
type Aps struct {
	u                    *Unifi
	Adopted              bool
	Bytes                int
	Cfgversion           string
	Config_network       ConfigNetwork
	Connect_request_ip   string
	Connect_request_port string
	Considered_lost_at   int
	Downlink_table       []DownlinkTable
	Guest_num_sta        int `json:"guest-num_sta"`
	Guest_token          string
	Has_eth1             bool
	Has_poe_passthrough  bool
	Inform_authkey       string
	Inform_ip            string
	Inform_url           string
	Ip                   string
	Known_cfgversion     string
	Last_seen            int
	Locating             bool
	Locked               bool
	Mac                  string
	Map_id               string
	Model                string
	Na_channel           int    `json:"na-channel"`
	Na_eirp              int    `json:"na-eirp"`
	Na_extchannel        int    `json:"na-extchannel"`
	Na_gain              int    `json:"na-gain"`
	Na_state             string `json:"na-state"`
	Na_tx_power          string `json:"na-tx_power"`
	Name                 string
	Next_heartbeat_at    int
	Ng_channel           int    `json:"ng-channel"`
	Ng_eirp              int    `json:"ng-eirp"`
	Ng_extchannel        int    `json:"ng-extchannel"`
	Ng_gain              int    `json:"ng-gain"`
	Ng_state             string `json:"ng-state"`
	Ng_tx_power          string `json:"ng-tx_power"`
	Num_sta              int
	Radio_table          []RadioTable
	Rx_bytes             int
	Scanning             bool
	Serial               string
	Stat                 Stat
	State                int
	Tx_bytes             int
	Uplink               Uplink
	Uplink_table         []UplinkTable
	Uptime               int
	User_num_sta         int `json:"user-num_sta"`
	Vap_table            []VapTable
	Version              string
	VwireEnabled         bool
	Vwire_table          []VwireTable
	X                    int
	X_authkey            string
	X_fingerprint        string
	X_vwirekey           string
	Y                    int
}

// Reboot access point
func (a Aps) Restart() {
	if a.u == nil {
		log.Fatal("login first")
	}
	a.u.maccmd(a.Mac, "restart", "devmgr")
}
