// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

import (
	"strconv"
	"time"
)

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	res, err := strconv.ParseInt(string(b), 0, 32)
	if err != nil {
		return err
	}
	*t = Timestamp(time.Unix(res, 0))
	return nil
}

func (t Timestamp) String() string {
	return time.Time(t).String()
}
