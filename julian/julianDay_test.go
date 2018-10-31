package kta

import (
    "fmt"
    "testing"
	"time"
)

//测试获取数据
func TestNewJulianDay(t *testing.T) {
	//测试获取数据
	jd := NewJulianDay(2018,10,1)
	year, month, day, err := jd.GetYMD()
	if err != nil {
        t.Error(err)
    } else if year != 2018 || month != 10 || day != 1 {
		t.Error("YMD diff 20181001", year, month, day)
    }
	jdNew := newJDN(0,1,1)
	year, month, day = jdNew.CalcYMD()
	if year != 0 || month != 1 || day != 1 {
		t.Error("CalcYMD diff", year, month, day)
	} else {
		t.Log("JulianDayfor 01/01/0000", int(jdNew))
	}
	jdNew = JulianDay(0)
	year, month, day = jdNew.CalcYMD()
	t.Log("YMD for JulianDay(0)", year, month, day)
}

func TestParseJulianDay(t *testing.T) {
	dateStr := "20181003"
	jd := ParseJulianDay("YYYYMMDD", dateStr)
	jdStr := jd.String8()
	if dateStr != jdStr {
		t.Error("Date diff", dateStr, jdStr)
	}
	dateStr = "20081101"
	jd = ParseJulianDay("CCYYMMDD", dateStr)
	jdStr = jd.String8()
	if dateStr != jdStr {
		t.Error("Date diff", dateStr, jdStr)
	}
	jd = ParseJulianDay("YYYYMMDD", time.Now().Format("20060102"))
	fmt.Printf("Today's %s JulianDate %d weekday: %d\n", jd,int(jd), int(jd)%7)
}

func BenchmarkParseJulianDay( b *testing.B) {
	dateStr := "20181011"
	for i := 0; i< b.N; i++ {
		_ = ParseJulianDay("YYYYMMDD", dateStr)
	}
}

func BenchmarkJdOlsStr( b *testing.B) {
	dateStr := "20181011"
	jd := ParseJulianDay("YYYYMMDD", dateStr)
	for i := 0; i< b.N; i++ {
		_ = jd.String8()
	}
}

func BenchmarkJdStr( b *testing.B) {
	dateStr := "20181011"
	jd := ParseJulianDay("YYYYMMDD", dateStr)
	for i := 0; i< b.N; i++ {
		_ = jd.String()
	}
}

func BenchmarkNewJulianDay( b *testing.B) {
	for i := 0; i< b.N; i++ {
		_ = newJulianDay(2018,10,1)
	}
}

func BenchmarkNewJDN( b *testing.B) {
	for i := 0; i< b.N; i++ {
		_ = newJDN(2018,10,1)
	}
}

func BenchmarkGetYMD( b *testing.B) {
	jd := newJDN(2018, 10, 4)
	for i := 0; i< b.N; i++ {
		_,_,_,err := jd.GetYMD()
		if err != nil { break }
	}
}

func BenchmarkCalcYMD( b *testing.B) {
	jd := newJDN(2018, 10, 4)
	for i := 0; i< b.N; i++ {
		_,_,_ = jd.CalcYMD()
	}
}
