package unifi

type User struct {
	FirstSeen int `json:"first_seen"`
	Hostname  string
	IsGuest   bool `json:"is_guest"`
	LastSeen  int  `json:"last_seen"`
	Mac       string
	Oui       string
}
