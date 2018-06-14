// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

import (
	"encoding/json"
	"fmt"
)


type AlarmFilter struct {
	Limit  int `json:"_limit"`  
	Start  int `json:"_start"`
	Withcount bool `json:"withcount"`
	Archived bool `json:"archived"`
}

type RawAlarm struct {
	Data json.RawMessage
	Key  string
}

// Returns a slice of json RawDevices as received by the controller
func (u *Unifi) RawAlarms(site *Site, filter interface{}) ([]RawAlarm, error) {

	var rawAlarms []RawAlarm

	// Response from controller
	var response struct {
		Data []json.RawMessage
		Meta meta
	}
		
	err := u.parse(site, "stat/alarm", filter, &response)
	if err != nil {
		return nil, err
	}

	
	for _, d := range response.Data {

		// unmarshal into a map to check the "type" field
		var obj map[string]interface{}
		err := json.Unmarshal(d, &obj)
		if err != nil {
			return nil, err
		}

		alarmKey, ok := obj["key"].(string)
		if !ok {
			return nil, fmt.Errorf("Error on retrieving object type from raw Json")
		}

		var ra RawAlarm
		ra.Key = alarmKey
		ra.Data = d
		rawAlarms = append(rawAlarms, ra)
	}
	return rawAlarms, err
}

// Returns a slice of json RawDevices as received by the controller
func (u *Unifi) BasicAlarms(site *Site, filter interface{}) ([]BasicEvent, error) {

	var basicAlarms []BasicEvent

	rawAlarms, err := u.RawAlarms(site, filter) 

	for _, ra := range rawAlarms{

		var ba BasicEvent
		err := json.Unmarshal(ra.Data, &ba)
		if err != nil {
			return nil, err
		}

		basicAlarms = append(basicAlarms, ba)
	}

	return basicAlarms, err
}
