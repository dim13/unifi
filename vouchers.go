package unifi

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
	SiteId       string `json:"site_id"`
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

func (u *Unifi) Voucher() ([]Voucher, error) {
	var response struct {
		Data []Voucher
		Meta meta
	}
	err := u.parse("stat/voucher", &response)
	for i := range response.Data {
		response.Data[i].u = u
	}

	return response.Data, err
}

func (u *Unifi) VoucherMap() (VoucherMap, error) {
	vouch, err := u.Voucher()
	if err != nil {
		return nil, err
	}
	m := make(VoucherMap)
	for _, a := range vouch {
		m[a.Code] = a
	}
	return m, nil
}
