package unifi

type WlanConf struct {
	Enabled        bool
	HideSsid       bool `json:"hide_ssid"`
	IsGuest        bool `json:"is_guest"`
	Name           string
	RadiusIP1      string `json:"radius_ip_1"`
	RadiusPort1    string `json:"radius_port_1"`
	Security       string
	UsergroupID    string `json:"usergroup_id"`
	Vlan           string
	VlanEnabled    bool   `json:"vlan_enabled"`
	WepIdx         string `json:"wep_idx"`
	WpaEnc         string `json:"wpa_enc"`
	WpaMode        string `json:"wpa_mode"`
	XPassphrase    string `json:"x_passphrase"`
	XRadiusSecret1 string `json:"x_radius_secret_1"`
	XWep           string `json:"x_wep"`
}
