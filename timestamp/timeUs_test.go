package timestamp

import (
    "testing"
	"time"
)

func TestDurationUs(t *testing.T) {
	//测试获取数据
	n1 := time.Now()
	min2 := DurationUs(120000000)
	n2 := min2.Time(n1)
	dms := DiffUs(n1, n2)
    if min2 != dms {
		t.Error("DiffUs diff", n1.Unix(), n2.Unix(), int(dms))
    }
	t.Log("Now & plus 2min", n1, n2)
}

func TestDateTimeUs(t *testing.T) {
	t1 := time.Now()
	ms1 := ToDateTimeUs(t1)
	if t1.Unix() != ms1.Unix() || t1.Nanosecond()/1000 != ms1.Usecond() {
		t.Error("DateTime diff", t1.Unix(), ms1.Unix())
	}
}

func BenchmarkDurationUs( b *testing.B) {
	t1 := time.Now()
	t2 := time.Unix(t1.Unix()-1000, 0) 
	for i := 0; i< b.N; i++ {
		_ = DiffUs(t1, t2)
	}
}
