package unifi

type WlanConf struct {
	Enabled           bool
	Hide_ssid         bool
	Is_guest          bool
	Name              string
	Radius_ip_1       string
	Radius_port_1     string
	Security          string
	Usergroup_id      string
	Vlan              string
	Vlan_enabled      bool
	Wep_idx           string
	Wpa_enc           string
	Wpa_mode          string
	X_passphrase      string
	X_radius_secret_1 string
	X_wep             string
}
