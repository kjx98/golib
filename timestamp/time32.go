package timestamp

import (
	"time"
)

type TimeT32 uint32

// return int64 value of time seconds from 1970/1/1
func (timeV TimeT32) Unix() int64 {
	return int64(timeV)
}

// returns the local Time corresponding to the given Unix time, sec
//   seconds and nsec nanoseconds since January 1, 1970 UTC.
func (timeV TimeT32) Time() time.Time {
	return time.Unix(int64(timeV), 0)
}

func (timeV TimeT32) String() string {
	return timeV.Time().Format("01-02 15:04:05")
}

type TimeHM int32

// GetHM from Time
func GetTimeHM(t time.Time) TimeHM {
	res := int32(t.Hour())*100 + int32(t.Minute())
	return TimeHM(res)
}

// TimeHM Sub
// t1, t2 must >= 0
func (t1 TimeHM) Sub(t2 TimeHM) TimeHM {
	if (t1 % 100) >= (t2 % 100) {
		return t1 - t2
	} else {
		return (t1 - 40) - t2
	}
}

// TimeHM Add
// t1 must >= 0
func (t1 TimeHM) Add(t2 TimeHM) TimeHM {
	if t2 < 0 {
		return t1.Sub(-t2)
	}
	res := t1 + t2
	if res%100 >= 60 {
		res += 40
	}
	return res
}

func (t1 TimeHM) Second() int {
	res := int(t1 % 100)
	res += int(t1/100) * 60
	return res * 60
}
