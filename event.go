// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

import (
	"encoding/json"
	"fmt"
	"time"
)

type EventFilter struct {
	Limit  int `json:"_limit"`
	Start  int `json:"_start"`
	Within int `json:"within"`
}

type RawEvent struct {
	Data json.RawMessage
	Key  string
}

type BasicEvent struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field	ID        string `json:"_id"`
	ID             string    `json:"_id"`
	Datetime       time.Time
	Key            string `json:"key"`
	Message        string `json:"msg"`
	Time           int64  `json:"time"`
}

type EVT_AD_Login struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Admin          string    `json:"admin"`
	Datetime       time.Time `json:"datetime"`
	IP             string    `json:"ip"`
	IsAdmin        bool      `json:"is_admin"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Time           int64     `json:"time"`
}

type EVT_AP_Adopted struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Admin          string    `json:"admin"`
	Ap             string    `json:"ap"`
	ApName         string    `json:"ap_name"`
	Datetime       time.Time `json:"datetime"`
	IsAdmin        bool      `json:"is_admin"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Time           int64     `json:"time"`
}

type EVT_AP_Connected struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Ap             string    `json:"ap"`
	ApName         string    `json:"ap_name"`
	Datetime       time.Time `json:"datetime"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Time           int64     `json:"time"`
}

type EVT_AP_DetectRogueAP struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Datetime       time.Time `json:"datetime"`
	Essid          string    `json:"essid"`
	Key            string    `json:"key"`
	Mac            string    `json:"mac"`
	Msg            string    `json:"msg"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Time           int64     `json:"time"`
}

type EVT_AP_DiscoveredPending struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Ap             string    `json:"ap"`
	ApName         string    `json:"ap_name"`
	Datetime       time.Time `json:"datetime"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Time           int64     `json:"time"`
}

type EVT_AP_Isolated struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Ap             string    `json:"ap"`
	Datetime       time.Time `json:"datetime"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Time           int64     `json:"time"`
}

type EVT_AP_Lost_Contact struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Ap             string    `json:"ap"`
	ApName         string    `json:"ap_name"`
	Datetime       time.Time `json:"datetime"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Time           int64     `json:"time"`
}

type EVT_AP_PossibleInterference struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Ap             string    `json:"ap"`
	Channel        string    `json:"channel"`
	Datetime       time.Time `json:"datetime"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	Radio          string    `json:"radio"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Time           int64     `json:"time"`
}

type EVT_AP_Restarted struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Admin          string    `json:"admin"`
	Ap             string    `json:"ap"`
	ApName         string    `json:"ap_name"`
	Datetime       time.Time `json:"datetime"`
	IsAdmin        bool      `json:"is_admin"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Time           int64     `json:"time"`
}

type EVT_AP_RestartedUnknown struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Ap             string    `json:"ap"`
	ApName         string    `json:"ap_name"`
	Datetime       time.Time `json:"datetime"`
	Duration       int       `json:"duration"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	NumSta         int       `json:"num_sta"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Time           int64     `json:"time"`
}

type EVT_AP_Upgraded struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Ap             string    `json:"ap"`
	ApName         string    `json:"ap_name"`
	Datetime       time.Time `json:"datetime"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Time           int64     `json:"time"`
	VersionFrom    string    `json:"version_from"`
	VersionTo      string    `json:"version_to"`
}

type EVT_AP_UpgradeScheduled struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Admin          string    `json:"admin"`
	Ap             string    `json:"ap"`
	ApName         string    `json:"ap_name"`
	Datetime       time.Time `json:"datetime"`
	IsAdmin        bool      `json:"is_admin"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Time           int64     `json:"time"`
}

type EVT_LU_Connected struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Datetime       time.Time `json:"datetime"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	Network        string    `json:"network"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Time           int64     `json:"time"`
	User           string    `json:"user"`
}

type EVT_LU_Disconnected struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Bytes          int       `json:"bytes"`
	Datetime       time.Time `json:"datetime"`
	Duration       int       `json:"duration"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	Network        string    `json:"network"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Time           int64     `json:"time"`
	User           string    `json:"user"`
}

type EVT_SW_Adopted struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Admin          string    `json:"admin"`
	Datetime       time.Time `json:"datetime"`
	IsAdmin        bool      `json:"is_admin"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Sw             string    `json:"sw"`
	SwName         string    `json:"sw_name"`
	Time           int64     `json:"time"`
}

type EVT_SW_Connected struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Datetime       time.Time `json:"datetime"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Sw             string    `json:"sw"`
	SwName         string    `json:"sw_name"`
	Time           int64     `json:"time"`
}

type EVT_SW_DiscoveredPending struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Datetime       time.Time `json:"datetime"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Sw             string    `json:"sw"`
	SwName         string    `json:"sw_name"`
	Time           int64     `json:"time"`
}

type EVT_SW_Lost_Contact struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Datetime       time.Time `json:"datetime"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Sw             string    `json:"sw"`
	SwName         string    `json:"sw_name"`
	Time           int64     `json:"time"`
}

// EVT_SW_PoeDisconnect controller version 5.8.21+ / USW firmware 3.9.27+
type EVT_SW_PoeDisconnect struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Datetime       time.Time `json:"datetime"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	Port           int       `json:"port"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Sw             string    `json:"sw"`
	SwName         string    `json:"sw_name"`
	Time           int64     `json:"time"`
}

type EVT_SW_RestartedUnknown struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Datetime       time.Time `json:"datetime"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Sw             string    `json:"sw"`
	SwName         string    `json:"sw_name"`
	Time           int64     `json:"time"`
}

type EVT_SW_Upgraded struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Datetime       time.Time `json:"datetime"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Sw             string    `json:"sw"`
	SwName         string    `json:"sw_name"`
	Time           int64     `json:"time"`
	VersionFrom    string    `json:"version_from"`
	VersionTo      string    `json:"version_to"`
}

type EVT_SW_UpgradeScheduled struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Admin          string    `json:"admin"`
	Datetime       time.Time `json:"datetime"`
	IsAdmin        bool      `json:"is_admin"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Sw             string    `json:"sw"`
	SwName         string    `json:"sw_name"`
	Time           int64     `json:"time"`
}

type EVT_WU_Connected struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Ap             string    `json:"ap"`
	Channel        string    `json:"channel"`
	Datetime       time.Time `json:"datetime"`
	Hostname       string    `json:"hostname"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	Radio          string    `json:"radio"`
	SiteID         string    `json:"site_id"`
	Ssid           string    `json:"ssid"`
	Subsystem      string    `json:"subsystem"`
	Time           int64     `json:"time"`
	User           string    `json:"user"`
}

type EVT_WU_Disconnected struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Ap             string    `json:"ap"`
	Bytes          int       `json:"bytes"`
	Datetime       time.Time `json:"datetime"`
	Duration       int       `json:"duration"`
	Hostname       string    `json:"hostname"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	SiteID         string    `json:"site_id"`
	Ssid           string    `json:"ssid"`
	Subsystem      string    `json:"subsystem"`
	Time           int64     `json:"time"`
	User           string    `json:"user"`
}

type EVT_WU_Roam struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	ApFrom         string    `json:"ap_from"`
	ApTo           string    `json:"ap_to"`
	Channel        string    `json:"channel"`
	ChannelFrom    string    `json:"channel_from"`
	ChannelTo      string    `json:"channel_to"`
	Datetime       time.Time `json:"datetime"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	Radio          string    `json:"radio"`
	RadioFrom      string    `json:"radio_from"`
	RadioTo        string    `json:"radio_to"`
	SiteID         string    `json:"site_id"`
	Ssid           string    `json:"ssid"`
	Subsystem      string    `json:"subsystem"`
	Time           int64     `json:"time"`
	User           string    `json:"user"`
}

type EVT_WU_RoamRadio struct {
	Archived       *bool     `json:"archived,omitempty"`         // Alarm field
	HandledAdminID string    `json:"handled_admin_id,omitempty"` // Alarm field
	HandledTime    time.Time `json:"handled_time,omitempty"`     // Alarm field
	ID             string    `json:"_id"`
	Ap             string    `json:"ap"`
	ChannelFrom    string    `json:"channel_from"`
	ChannelTo      string    `json:"channel_to"`
	Datetime       time.Time `json:"datetime"`
	Key            string    `json:"key"`
	Msg            string    `json:"msg"`
	RadioFrom      string    `json:"radio_from"`
	RadioTo        string    `json:"radio_to"`
	SiteID         string    `json:"site_id"`
	Subsystem      string    `json:"subsystem"`
	Time           int64     `json:"time"`
	User           string    `json:"user"`
}

// Returns a slice of json RawDevices as received by the controller
func (u *Unifi) RawEvents(site *Site, filter any) ([]RawEvent, error) {

	var rawEvents []RawEvent

	// Response from controller
	var response struct {
		Data []json.RawMessage
		Meta meta
	}

	err := u.parse(site, "stat/event", filter, &response)
	if err != nil {
		return nil, err
	}

	for _, d := range response.Data {

		// unmarshal into a map to check the "type" field
		var obj map[string]any
		err := json.Unmarshal(d, &obj)
		if err != nil {
			return nil, err
		}

		eventKey, ok := obj["key"].(string)
		if !ok {
			return nil, fmt.Errorf("error on retrieving object type from raw json")
		}

		var re RawEvent
		re.Key = eventKey
		re.Data = d
		rawEvents = append(rawEvents, re)
	}
	return rawEvents, err
}

// Returns a slice of json RawDevices as received by the controller
func (u *Unifi) BasicEvents(site *Site, filter any) ([]BasicEvent, error) {

	var basicEvents []BasicEvent

	rawEvents, err := u.RawEvents(site, filter)

	for _, re := range rawEvents {

		var be BasicEvent
		err := json.Unmarshal(re.Data, &be)
		if err != nil {
			return nil, err
		}

		basicEvents = append(basicEvents, be)
	}

	return basicEvents, err
}
