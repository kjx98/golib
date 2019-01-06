package timestamp

import (
	"testing"
	"time"
)

func TestTime32(t *testing.T) {
	t1 := time.Now()
	t32 := TimeT32(t1.Unix())
	t2 := t32.Time()
	if uint32(t1.Unix()) != uint32(t32) || t1.Unix() != t2.Unix() {
		t.Error("TimeT32 diff", t1.Unix(), t32.Unix())
	}
	t.Log(t1.Format("01-02 15:04:05"), t32)
}

func TestTimeHM(t *testing.T) {
	t1 := TimeHM(915)
	t2 := TimeHM(120)
	t3 := TimeHM(50)
	if t1.Add(t2) != TimeHM(1035) {
		t.Error("TimeHM add diff", t1, t2, t1.Add(t2))
	}
	if t1.Add(t3) != TimeHM(1005) {
		t.Error("TimeHM add diff", t1, t3, t1.Add(t3))
	}
	if t1.Sub(t2) != TimeHM(755) {
		t.Error("TimeHM sub diff", t1, t2, t1.Sub(t2))
	}
	if t1.Add(-t2) != TimeHM(755) {
		t.Error("TimeHM add diff", t1, -t2, t1.Add(-t2))
	}
}
