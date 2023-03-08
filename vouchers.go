// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"reflect"
)

type VoucherMap map[string]Voucher

type Voucher struct {
	u            *Unifi
	AdminName    string `json:"admin_name"`
	Code         string
	CreateTime   int `json:"create_time"`
	Duration     int
	ForHotspot   bool `json:"for_hotspot"`
	Note         string
	QosOverwrite bool `json:"qos_overwrite"`
	Quota        int
	SiteID       string `json:"site_id"`
	Used         int
}

type NewVoucher struct {
	Cmd          string `json:"cmd"`
	Expire       string `json:"expire"`
	ExpireNumber string `json:"expire_number"`
	ExpireUnit   string `json:"expire_unit"`
	N            string `json:"n"`
	Note         string `json:"note"`
	Quota        string `json:"quota"`
}

// Value with parameters for create New Voucher
var Nv NewVoucher

func (u *Unifi) Voucher(site *Site) ([]Voucher, error) {
	var response struct {
		Data []Voucher
		Meta meta
	}
	err := u.parse(site, "stat/voucher", nil, &response)
	for i := range response.Data {
		response.Data[i].u = u
	}

	return response.Data, err
}

func (u *Unifi) VoucherMap(site *Site) (VoucherMap, error) {
	vouch, err := u.Voucher(site)
	if err != nil {
		return nil, err
	}
	m := make(VoucherMap)
	for _, a := range vouch {
		m[a.Code] = a
	}
	return m, nil
}

//Functions creating new Vouchers

func (u *Unifi) NewVoucher(site *Site, nv NewVoucher) ([]Voucher, error) {
	var response struct {
		Data []Voucher
		Meta meta
	}

	Nv = nv
	err := u.parseNewVoucher(site, "cmd/hotspot", &response)
	return response.Data, err

}

func (u *Unifi) apicmdNewVoucher(site *Site, cmd string) ([]byte, error) {
	jsonData := Nv

	// Setup url
	cmdurl := u.apiURL
	cmdurl += fmt.Sprintf("s/%s/%s", site.Name, cmd)

	data, err := json.Marshal(jsonData)

	if err != nil {
		return nil, err
	}
	val := url.Values{"json": {string(data)}}

	resp, err := u.client.PostForm(cmdurl, val)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (u *Unifi) parseNewVoucher(site *Site, cmd string, v any) error {
	body, err := u.apicmdNewVoucher(site, cmd)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &v); err != nil {
		log.Println(cmd)
		log.Println(string(body))
		return err
	}
	m := reflect.ValueOf(v).Elem().FieldByName("Meta").Interface().(meta)
	if m.Rc != "ok" {
		return fmt.Errorf("bad request: %s", m.Rc)
	}
	return nil
}
