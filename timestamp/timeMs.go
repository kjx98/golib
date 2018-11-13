package timestamp

import (
	"errors"
	"time"
)

type DurationMs	uint32

type DateTimeMs	int64


// calc time.Time from baseTime plus DurationMs
func (deltaMs DurationMs) Time(baseTime int64) time.Time {
	deltaSec := int64(deltaMs/1000)
	ms := int64(deltaMs%1000)
	return time.Unix(baseTime+deltaSec, ms*1000000)
}

func DiffMs(st, en time.Time) (DurationMs,error) {
	diffNs := en.Nanosecond() - st.Nanosecond()
	diffSec := en.Unix() - st.Unix()
	if diffNs < 0 {
		diffSec--
		diffNs += 1000000000
	}
	if diffSec > 4000000 || diffSec < -4000000 {
		return DurationMs(0), errors.New("Overflow")
	}
	diffNs /= 1000000
	diffSec *= 1000
	return DurationMs(diffSec+int64(diffNs)), nil
}

// no consider for overflow int64 datetimeMs, about 4.9e5 years
// Convert DateTimeMs
func (dtMs DateTimeMs) Time() time.Time {
	nanoSec := int64(dtMs%1000) * 1000000
	sec := int64(dtMs/1000)
	return time.Unix(sec, nanoSec)
}

// return seconds from 1970/1/1 UTC
func (dtMs DateTimeMs) Unix() int64 {
	return int64(dtMs)/1000
}

func (dtMs DateTimeMs) Millisecond() int {
	return int(dtMs%1000)
}

// convert time.Time to DateTimeMs
func ToDateTimeMs(dt time.Time) DateTimeMs {
	ms := dt.Nanosecond() / 1000000
	return DateTimeMs(dt.Unix()*1000 + int64(ms))
}
