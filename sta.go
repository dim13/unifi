// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

type StaMap map[string]Sta

// Station data
type Sta struct {
	u                *Unifi
	ID               string `json:"_id"`
	IsGuestByUsw     bool   `json:"_is_guest_by_usw,omitempty"`
	LastSeenByUsw    int    `json:"_last_seen_by_usw,omitempty"`
	UptimeByUsw      int    `json:"_uptime_by_usw,omitempty"`
	AssocTime        int    `json:"assoc_time"`
	FirstSeen        int64  `json:"first_seen"`
	IP               string `json:"ip"`
	IsGuest          bool   `json:"is_guest"`
	IsWired          bool   `json:"is_wired"`
	LastSeen         int64  `json:"last_seen"`
	LatestAssocTime  int    `json:"latest_assoc_time"`
	Mac              string `json:"mac"`
	Network          string `json:"network,omitempty"`
	NetworkID        string `json:"network_id,omitempty"`
	Oui              string `json:"oui"`
	SiteID           string `json:"site_id"`
	SwDepth          int    `json:"sw_depth,omitempty"`
	SwMac            string `json:"sw_mac,omitempty"`
	SwPort           int    `json:"sw_port,omitempty"`
	Uptime           int    `json:"uptime"`
	UserID           string `json:"user_id"`
	Hostname         string `json:"hostname,omitempty"`
	IsGuestByUap     bool   `json:"_is_guest_by_uap,omitempty"`
	LastSeenByUap    int    `json:"_last_seen_by_uap,omitempty"`
	RoamCount        int    `json:"roam_count,omitempty"`
	UptimeByUap      int    `json:"_uptime_by_uap,omitempty"`
	ApMac            string `json:"ap_mac,omitempty"`
	Authorized       bool   `json:"authorized,omitempty"`
	BSSID            string `json:"bssid,omitempty"`
	BytesR           int    `json:"bytes-r,omitempty"`
	Ccq              int    `json:"ccq,omitempty"`
	Channel          int    `json:"channel,omitempty"`
	ESSID            string `json:"essid,omitempty"`
	Idletime         int    `json:"idletime,omitempty"`
	Is11R            bool   `json:"is_11r,omitempty"`
	Noise            int    `json:"noise,omitempty"`
	PowersaveEnabled bool   `json:"powersave_enabled,omitempty"`
	QosPolicyApplied bool   `json:"qos_policy_applied,omitempty"`
	Radio            string `json:"radio,omitempty"`
	RadioProto       string `json:"radio_proto,omitempty"`
	Rssi             int    `json:"rssi,omitempty"`
	RxBytes          int64  `json:"rx_bytes,omitempty"`
	RxBytesR         int64  `json:"rx_bytes-r,omitempty"`
	RxPackets        int64  `json:"rx_packets,omitempty"`
	RxRate           int64  `json:"rx_rate,omitempty"`
	Signal           int    `json:"signal,omitempty"`
	TxBytes          int64  `json:"tx_bytes,omitempty"`
	TxBytesR         int64  `json:"tx_bytes-r,omitempty"`
	TxPackets        int64  `json:"tx_packets,omitempty"`
	TxPower          int64  `json:"tx_power,omitempty"`
	TxRate           int64  `json:"tx_rate,omitempty"`
	Vlan             int    `json:"vlan,omitempty"`
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
