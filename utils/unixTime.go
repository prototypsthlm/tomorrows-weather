package utils

import (
	"time"
)

func ParseUnixTimestamp(unixTimestamp int) time.Time {
	tm := time.Unix(int64(unixTimestamp), 0)
	return tm
}
