package timestamp

import (
	"time"
)

type DurationUs	int64

type DateTimeUs	int64


// calc time.Time from baseTime plus DurationMs
func (deltaUs DurationUs) Time(baseTime time.Time) time.Time {
	deltaSec := int64(deltaUs/1000000)
	ns := int64(deltaUs%1000000) * 1000
	ns += int64(baseTime.Nanosecond())
	if ns >= 1000000000 {
		ns -= 1000000000
		deltaSec++
	}
	return time.Unix(baseTime.Unix()+deltaSec, ns)
}

// no consider overflow
func DiffUs(st, en time.Time) DurationUs {
	diffNs := en.Nanosecond() - st.Nanosecond()
	diffSec := en.Unix() - st.Unix()
	if diffNs < 0 {
		diffSec--
		diffNs += 1000000000
	}
	diffNs /= 1000
	diffSec *= 1000000
	return DurationUs(diffSec+int64(diffNs))
}

// no consider for overflow int64 datetimeMs, about 4.9e5 years
// Convert DateTimeMs
func (dtUs DateTimeUs) Time() time.Time {
	nanoSec := dtUs & 0xfffff
	return time.Unix(int64(dtUs >> 20), int64(nanoSec))
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
	return DateTimeUs(sec | int64(us & 0xfffff))
}
