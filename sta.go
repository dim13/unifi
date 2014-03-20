package unifi

type Sta struct {
	Ap_mac             string
	Assoc_time         int //Timestamp
	Auth_time          int //Timestamp
	Authorized         bool
	Bssid              string
	Bytes              int
	Bytes_d            int `json:"bytes.d"`
	Bytes_r            int `json:"bytes.r"`
	Ccq                int
	Channel            int
	Dhcpend_time       int
	Dhcpstart_time     int
	Essid              string
	First_seen         int //Timestamp
	Hostname           string
	Idletime           int
	Ip                 string
	Is_11a             bool
	Is_11ac            bool
	Is_11b             bool
	Is_11n             bool
	Is_guest           bool
	Last_seen          int //Timestamp
	Mac                string
	Map_id             string
	Noise              int
	Oui                string
	Powersave_enabled  bool
	Qos_policy_applied bool
	Radio              string
	Roam_count         int
	Rssi               int
	Rx_bytes           int
	Rx_bytes_d         int `json:"rx_bytes.d"`
	Rx_bytes_r         int `json:"rx_bytes.r"`
	Rx_crypts          int
	Rx_crypts_d        int `json:"rx_crytps.d"`
	Rx_crypts_r        int `json:"rx_crytps.r"`
	Rx_dropped         int
	Rx_dropped_d       int `json:"rx_dropped.d"`
	Rx_dropped_r       int `json:"rx_dropped.r"`
	Rx_errors          int
	Rx_errors_d        int `json:"rx_errors.d"`
	Rx_errors_r        int `json:"rx_errors.r"`
	Rx_frags           int
	Rx_frags_d         int `json:"rx_frags.d"`
	Rx_frags_r         int `json:"rx_frags.r"`
	Rx_packets         int
	Rx_packets_d       int `json:"rx_packets.d"`
	Rx_packets_r       int `json:"rx_packets.r"`
	Rx_rate            int
	Signal             int
	State              int
	State_ht           bool
	State_pwrmgt       bool
	T                  string
	Tx_bytes           int
	Tx_bytes_d         int `json:"tx_bytes.d"`
	Tx_bytes_r         int `json:"tx_bytes.r"`
	Tx_dropped         int
	Tx_dropped_d       int `json:"tx_dropped.d"`
	Tx_dropped_r       int `json:"tx_dropped.r"`
	Tx_errors          int
	Tx_errors_d        int `json:"tx_errors.d"`
	Tx_errors_r        int `json:"tx_errors.r"`
	Tx_packets         int
	Tx_packets_d       int `json:"tx_packets.d"`
	Tx_packets_r       int `json:"tx_packets.r"`
	Tx_power           int
	Tx_rate            int
	Tx_retries         int
	Tx_retries_d       int `json:"tx_retries.d"`
	Tx_retries_r       int `json:"tx_retries.r"`
	Uptime             int
	User_id            string
}
