package timestamp

import (
	"time"
)

type DurationUs	int64

type DateTimeUs	struct {
	second	int64
	usecond	int
}


// calc time.Time from baseTime plus DurationMs
func (deltaMs DurationUs) Time(baseTime time.Time) time.Time {
	deltaSec := int64(deltaMs/1000000)
	ms := int64(deltaMs%1000000) * 1000
	ms += int64(baseTime.Nanosecond())
	if ms >= 1000000000 {
		ms -= 1000000000
		deltaSec++
	}
	return time.Unix(baseTime.Unix()+deltaSec, ms)
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
func (dtUs *DateTimeUs) Time() time.Time {
	nanoSec := dtUs.usecond * 1000
	return time.Unix(dtUs.second, int64(nanoSec))
}

// return seconds from 1970/1/1 UTC
func (dtUs *DateTimeUs) Unix() int64 {
	return dtUs.second
}

func (dtUs *DateTimeUs) Usecond() int {
	return dtUs.usecond
}

// convert time.Time to DateTimeMs
func ToDateTimeUs(dt time.Time) DateTimeUs {
	var dtUs=DateTimeUs{second: dt.Unix(), usecond: dt.Nanosecond()/1000}
	return dtUs
}
