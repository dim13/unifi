// Copyright (c) 2014 Dimitri Sokolyuk. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

type StaMap map[string]Sta

// Station data
type Sta struct {
	u                *Unifi
	ApMac            string `json:"ap_mac"`
	AssocTime        int    //Timestamp
	AuthTime         int    //Timestamp
	Authorized       bool
	BssID            string
	Bytes            int
	BytesD           int `json:"bytes.d"`
	BytesR           int `json:"bytes.r"`
	Ccq              int
	Channel          int
	DhcpendTime      int `json:"dhcpend_time"`
	DhcpstartTime    int `json:"dhcpstart_time"`
	EssID            string
	FirstSeen        int //Timestamp
	Hostname         string
	Idletime         int
	IP               string
	Is11a            bool `json:"is_11a"`
	Is11ac           bool `json:"is_11ac"`
	Is11b            bool `json:"is_11b"`
	Is11n            bool `json:"is_11n"`
	IsGuest          bool `json:"is_guest"`
	LastSeen         int  //Timestamp
	Mac              string
	MapID            string `json:"map_id"`
	Noise            int
	Oui              string
	PowersaveEnabled bool `json:"powersave_enabled"`
	QosPolicyApplied bool `json:"qos_policy_applied"`
	Radio            string
	RoamCount        int `json:"roam_count"`
	Rssi             int
	RxBytes          int `json:"rx_bytes"`
	RxBytesD         int `json:"rx_bytes.d"`
	RxBytesR         int `json:"rx_bytes.r"`
	RxCrypts         int `json:"rx_crypts"`
	RxCryptsD        int `json:"rx_crytps.d"`
	RxCryptsR        int `json:"rx_crytps.r"`
	RxDropped        int `json:"rx_dropped"`
	RxDroppedD       int `json:"rx_dropped.d"`
	RxDroppedR       int `json:"rx_dropped.r"`
	RxErrors         int `json:"rx_errors"`
	RxErrorsD        int `json:"rx_errors.d"`
	RxErrorsR        int `json:"rx_errors.r"`
	RxFrags          int `json:"rx_frags"`
	RxFragsD         int `json:"rx_frags.d"`
	RxFragsR         int `json:"rx_frags.r"`
	RxPackets        int `json:"rx_packets"`
	RxPacketsD       int `json:"rx_packets.d"`
	RxPacketsR       int `json:"rx_packets.r"`
	RxRate           int `json:"rx_rate"`
	Signal           int
	State            int
	StateHt          bool `json:"state_ht"`
	StatePwrmgt      bool `json:"state_pwrmgt"`
	T                string
	TxBytes          int `json:"tx_bytes"`
	TxBytesD         int `json:"tx_bytes.d"`
	TxBytesR         int `json:"tx_bytes.r"`
	TxDropped        int `json:"tx_dropped"`
	TxDroppedD       int `json:"tx_dropped.d"`
	TxDroppedR       int `json:"tx_dropped.r"`
	TxErrors         int `json:"tx_errors"`
	TxErrorsD        int `json:"tx_errors.d"`
	TxErrorsR        int `json:"tx_errors.r"`
	TxPackets        int `json:"tx_packets"`
	TxPacketsD       int `json:"tx_packets.d"`
	TxPacketsR       int `json:"tx_packets.r"`
	TxPower          int `json:"tx_power"`
	TxRate           int `json:"tx_rate"`
	TxRetries        int `json:"tx_retries"`
	TxRetriesD       int `json:"tx_retries.d"`
	TxRetriesR       int `json:"tx_retries.r"`
	Uptime           int
	UserID           string `json:"user_id"`
}

// Returns a station name
func (s Sta) Name() string {
	if s.Hostname != "" {
		return s.Hostname
	}
	if s.IP != "" {
		return s.IP
	}
	return s.Mac
}

func (s Sta) Block() error {
	if s.u == nil {
		return ErrLoginFirst
	}
	return s.u.stacmd(s.Mac, "block-sta")
}

func (s Sta) UnBlock() error {
	if s.u == nil {
		return ErrLoginFirst
	}
	return s.u.stacmd(s.Mac, "unblock-sta")
}

func (s Sta) Disconnect() error {
	if s.u == nil {
		return ErrLoginFirst
	}
	return s.u.stacmd(s.Mac, "kick-sta")
}

func (s Sta) AuthorizeGuest(minutes int) error {
	if s.u == nil {
		return ErrLoginFirst
	}
	return s.u.stacmd(s.Mac, "authorize-guest", minutes)
}

func (s Sta) UnauthorizeGuest() error {
	if s.u == nil {
		return ErrLoginFirst
	}
	return s.u.stacmd(s.Mac, "unauthorize-guest")
}
