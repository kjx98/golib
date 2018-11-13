package timestamp

import (
    "testing"
	"time"
)

func TestDurationMs(t *testing.T) {
	//测试获取数据
	ot := time.Now()
	n1 := time.Unix(ot.Unix(), 0)
	min2 := DurationMs(120000)
	n2 := min2.Time(n1.Unix())
	dms, err := DiffMs(n1, n2)
	if err != nil {
        t.Error(err)
    } else if min2 != dms {
		t.Error("DiffMs diff", n1.Unix(), n2.Unix(), int32(dms))
    }
	t.Log("Now & plus 2min", n1, n2)
}

func TestDateTimeMs(t *testing.T) {
	t1 := time.Now()
	ms1 := ToDateTimeMs(t1)
	if t1.Unix() != ms1.Unix() || t1.Nanosecond()/1000000 != ms1.Millisecond() {
		t.Error("DateTime diff", t1.Unix(), ms1.Unix())
	}
}

func BenchmarkDurationMs( b *testing.B) {
	t1 := time.Now()
	t2 := time.Unix(t1.Unix()-1000, 0) 
	for i := 0; i< b.N; i++ {
		_,err := DiffMs(t1, t2)
		if err != nil {
			b.Fatal(err)
		}
	}
}
