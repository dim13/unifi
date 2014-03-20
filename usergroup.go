package unifi

type UserGroup struct {
	Attr_hidden_id    string
	Attr_no_delete    bool
	DownRate_enabled  bool
	Name              string
	Qos_rate_max_down int
	Qos_rate_max_up   int
	UpRate_enabled    bool
}
