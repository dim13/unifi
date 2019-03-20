package unifi

import (
	"strconv"
)

type fuzzyFloat float64

func (f *fuzzyFloat) UnmarshalJSON(b []byte) error {
	s := string(b)
	if len(s) == 0 {
		return nil
	}
	// quick'n'dirty unquote
	if s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	// catch empty strings and null
	if s == "" || s == "null" {
		return nil
	}
	// try to parse
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*f = fuzzyFloat(v)
	return nil
}
