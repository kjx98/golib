package timestamp

import (
	"errors"
	"time"
)

type DurationMs uint32
type DateTimeMs int64

// calc time.Time from baseTime plus DurationMs
func (deltaMs DurationMs) Time(baseTime int64) time.Time {
	deltaSec := int64(deltaMs / 1000)
	ms := int64(deltaMs % 1000)
	return time.Unix(baseTime+deltaSec, ms*1e6)
}

// calc seconds of DurationMs
func (deltaMs DurationMs) Seconds() float64 {
	return float64(deltaMs) * 0.001
}

// calc ms duration
func DiffMs(st, en time.Time) (DurationMs, error) {
	diffMs := (en.Nanosecond() - st.Nanosecond()) / 1e6
	diffSec := en.Unix() - st.Unix()
	if diffMs < 0 {
		diffSec--
		diffMs += 1000
	}
	if diffSec > 4000000 || diffSec < -4000000 {
		return DurationMs(0), errors.New("Overflow")
	}
	diffSec *= 1000
	return DurationMs(diffSec + int64(diffMs)), nil
}

// no consider for overflow int64 datetimeMs, about 4.9e5 years
// Convert DateTimeMs
func (dtMs DateTimeMs) Time() time.Time {
	sec := int64(dtMs / 1000)
	ns := int64(dtMs%1000) * 1e6
	return time.Unix(sec, ns)
}

// return seconds from 1970/1/1 UTC
func (dtMs DateTimeMs) Unix() int64 {
	return int64(dtMs) / 1000
}

func (dtMs DateTimeMs) Millisecond() int {
	return int(dtMs % 1000)
}

// convert time.Time to DateTimeMs
func ToDateTimeMs(dt time.Time) DateTimeMs {
	ms := dt.Nanosecond() / 1e6
	return DateTimeMs(dt.Unix()*1000 + int64(ms))
}
