package timestamp

import (
	"time"
)

type timeT32 uint32

// return int64 value of time seconds from 1970/1/1
func (timeV timeT32) Time64() int64 {
	return int64(timeV)
}

// returns the local Time corresponding to the given Unix time, sec
//   seconds and nsec nanoseconds since January 1, 1970 UTC.
func (timeV timeT32) Time() time.Time {
	return time.Unix(int64(timeV), 0)
}

func (timeV timeT32) String() string {
	return timeV.Time().Format("01-02 15:04:05")
}
