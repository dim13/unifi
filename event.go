// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

import (
	"encoding/json"
	"time"
)

type Event struct {
	RawEvent  json.RawMessage
	ID        string `json:"_id"`
	Datetime  time.Time
	Key       string `json:"key"`
	Message   string `json:"msg"`
	Timestamp int64  `json:"time"`
}

type EVT_AP_Lost_Contact struct {
	ID        string    `json:"_id"`
	Ap        string    `json:"ap"`
	ApName    string    `json:"ap_name"`
	Datetime  time.Time `json:"datetime"`
	Key       string    `json:"key"`
	Msg       string    `json:"msg"`
	SiteID    string    `json:"site_id"`
	Subsystem string    `json:"subsystem"`
	Time      int64     `json:"time"`
}

// EVT_SW_PoeDisconnect controller version 5.8.21+ / USW firmware 3.9.27+
type EVT_SW_PoeDisconnect struct {
	ID        string    `json:"_id"`
	Datetime  time.Time `json:"datetime"`
	Key       string    `json:"key"`
	Msg       string    `json:"msg"`
	Port      int       `json:"port"`
	SiteID    string    `json:"site_id"`
	Subsystem string    `json:"subsystem"`
	Sw        string    `json:"sw"`
	SwName    string    `json:"sw_name"`
	Time      int64     `json:"time"`
}

// Returns a slice of json RawDevices as received by the controller
func (u *Unifi) Events(site *Site) ([]Event, error) {

	var events []Event

	// Response from controller
	var response struct {
		Data []json.RawMessage
		Meta meta
	}
	err := u.parse(site, "stat/event", &response)

	for _, e := range response.Data {

		var event Event
		err := json.Unmarshal(e, &event)
		if err != nil {
			return nil, err
		}

		event.RawEvent = e
		events = append(events, event)
	}

	return events, err
}
