package timestamp

import (
	"time"
)

type DurationUs int64
type DateTimeUs int64

// calc time.Time from baseTime plus DurationMs
func (deltaUs DurationUs) Time(baseTime time.Time) time.Time {
	deltaSec := int64(deltaUs / 1e6)
	ns := int64(deltaUs%1e6) * 1000
	ns += int64(baseTime.Nanosecond())
	if ns >= 1e9 {
		ns -= 1e9
		deltaSec++
	}
	return time.Unix(baseTime.Unix()+deltaSec, ns)
}

// calc seconds of DurationUs
func (deltaUs DurationUs) Seconds() float64 {
	return float64(deltaUs) * 0.000001
}

// no consider overflow
func DiffUs(st, en time.Time) DurationUs {
	diffUs := (en.Nanosecond() - st.Nanosecond()) / 1000
	diffSec := en.Unix() - st.Unix()
	if diffUs < 0 {
		diffSec--
		diffUs += 1e6
	}
	diffSec *= 1e6
	return DurationUs(diffSec + int64(diffUs))
}

// no consider for overflow int64 datetimeMs, about 4.9e5 years
// Convert DateTimeMs
func (dtUs DateTimeUs) Time() time.Time {
	nanoSec := dtUs & 0xfffff
	sec := dtUs >> 20
	return time.Unix(int64(sec), int64(nanoSec))
}

// return seconds from 1970/1/1 UTC
func (dtUs DateTimeUs) Unix() int64 {
	return int64(dtUs) >> 20
}

func (dtUs DateTimeUs) Usecond() int {
	return int(dtUs & 0xfffff)
}

// convert time.Time to DateTimeMs
func ToDateTimeUs(dt time.Time) DateTimeUs {
	sec := dt.Unix() << 20
	us := dt.Nanosecond() / 1000
	return DateTimeUs(sec + int64(us&0xfffff))
}
