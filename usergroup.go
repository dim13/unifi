package unifi

type UserGroup struct {
	AttrHiddenID    string `json:"attr_hidden_id"`
	AttrNoDelete    bool   `json:"attr_no_delete"`
	DownRateEnabled bool   `json:"downrate_enabled"`
	Name            string
	QosRateMaxDown  int  `json:"qos_rate_max_down"`
	QosRateMaxUp    int  `json:"qos_rate_max_up"`
	UpRateEnabled   bool `json:"uprate_enabled"`
}
