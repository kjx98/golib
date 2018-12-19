package julian

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type JulianDay uint32

const (
	JULIAN_ADJUSTMENT = 1721425
)

var month_tot = []int{0, 0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334, 365}

/* Jan Feb Mar  Apr  May  Jun  Jul  Aug  Sep  Oct  Nov  Dec
31  28  31   30   31   30   31   31   30   31   30   31
*/

//	static c4julian
//
//	Returns an (int) day of the year starting from 1.
//	Ex.    Jan 1, returns  1
//
//	Returns  -1  if it is an illegal date.
//
func c4julian(year, month, day int) int {
	var bLeap bool

	bLeap = (year%4 == 0 && year%100 != 0) || year%400 == 0

	month_days := month_tot[month+1] - month_tot[month]
	if month == 2 && bLeap {
		month_days++
	}

	if year < 0 ||
		month < 1 || month > 12 ||
		day < 1 || day > month_days {
		return -1 //Illegal Date
	}

	res := month_tot[month] + day
	if month > 2 && bLeap {
		res++
	}
	return res
}

//	c4ytoj -  Calculates the number of days to the year
func c4ytoj(yr int) int {
	/*
		This calculation takes into account the fact that
			1)  Years divisible by 400 are always leap years.
			2)  Years divisible by 100 but not 400 are not leap years.
			3)  Otherwise, years divisible by four are leap years.

		Since we do not want to consider the current year, we will
		subtract the year by 1 before doing the calculation.
	*/

	yr--
	return yr*365 + yr/4 - yr/100 + yr/400
}

// newJulianDay
//	Calc Julian Day with year, month, day
//       year > 0 for AD
//		 year<=0, year-1 for BC
func NewJulianDay(years, months, day int) JulianDay {
	return newJDN(years, months, day)
}

func newJulianDay(years, months, day int) JulianDay {
	res := c4julian(years, months, day)
	if res < 0 {
		return JulianDay(0)
	}
	res += c4ytoj(years)
	res += JULIAN_ADJUSTMENT
	return JulianDay(res)
}

// newJDN
//	Calc Fast Julian Day with year, month, day
//       year > 0 for AD
//		 year<=0, year-1 for BC
func newJDN(year, month, day int) JulianDay {
	if month < 1 || month > 12 || day < 1 || day > 31 {
		return JulianDay(0)
	}
	res := (1461 * (year + 4800 + (month-14)/12)) / 4
	res += (367 * (month - 2 - 12*((month-14)/12))) / 12
	res -= (3 * ((year + 4900 + (month-14)/12) / 100)) / 4
	res += day - 32075
	return JulianDay(res)
}

//	ParseJulianDay
//	Converts from formatted Date Data long julian Date format
//		format like "CCYYMMDD", "CCYY-MM-DD"
//			C for century, Y for year, M month, D day
func ParseJulianDay(pic, date string) (res JulianDay) {
	if date == "" {
		return
	}
	day_count := 7
	month_count := 4
	year_count := 1
	century_count := -1
	var buffer [16]byte
	var ybuff [8]byte
	for i, c := range pic {
		switch c {
		case 'D':
			day_count++
			if day_count >= 10 {
				break
			}
			buffer[day_count] = date[i]
		case 'M':
			month_count++
			if month_count >= 7 {
				break
			}
			buffer[month_count] = date[i]
		case 'Y':
			if year_count <= 6 {
				ybuff[year_count] = date[i]
			}
			year_count++
			if year_count >= 4 {
				break
			}
			buffer[year_count] = date[i]
		case 'C':
			century_count++
			if century_count >= 2 {
				break
			}
			buffer[century_count] = date[i]
		}
	}
	// We assume always exist date chars
	if year_count >= 4 {
		copy(buffer[:], ybuff[1:5])
	} else {
		if century_count == -1 {
			copy(buffer[:2], []byte("19"))
		}
		if year_count == 1 {
			copy(buffer[2:4], []byte("01"))
		}
	}
	if month_count == 4 {
		copy(buffer[5:7], []byte("01"))
	}
	if day_count == 7 {
		copy(buffer[8:10], []byte("01"))
	}
	years, err := strconv.Atoi(string(buffer[:4]))
	if err != nil {
		return
	}
	months, err := strconv.Atoi(string(buffer[5:7]))
	if err != nil {
		return
	}
	day, err := strconv.Atoi(string(buffer[8:10]))
	if err != nil {
		return
	}
	res = NewJulianDay(years, months, day)
	return
}

//	c4mon_dy
//
//	Given the year and the day of the year, returns the
//	month and day of month.
func c4mon_dy(year, days int) (month, day int, err error) {
	is_leap := 0
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		is_leap = 1
	}
	if days < 59 {
		is_leap = 0
	}

	for i := 2; i <= 13; i++ {
		if days <= month_tot[i]+is_leap {
			i--
			month = i
			if i <= 2 {
				is_leap = 0
			}
			day = days - month_tot[i] - is_leap
			return
		}
	}
	err = errors.New("Invalid days")
	return
}

func getYearDays(tot_days int) (year, days int) {
	/* A dBASE III index file date is stored as a julian day.  That is the
	   number of days since the date  Jan 1, 4713 BC
	   Ex.  Jan 1, 1981 is  2444606
	*/

	if tot_days <= JULIAN_ADJUSTMENT {
		return
	}
	tot_days -= JULIAN_ADJUSTMENT
	// year = (int) ((double)tot_days/365.2425) + 1
	year = tot_days * 400 / 146097
	days = tot_days - c4ytoj(year)
	if days <= 0 {
		year--
		days = tot_days - c4ytoj(year)
	}
	max_days := 365
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		max_days++
	}
	if days > max_days {
		year++
		days -= max_days
	}
	return
}

//	GetYMD
//
//	Converts from the julian date format to year/month/day format
func (tot_days JulianDay) GetYMD() (year, month, day int, err error) {
	var days int
	year, days = getYearDays(int(tot_days))
	month, day, err = c4mon_dy(year, days)
	return
}

//	Date
//
//	Fast convert julian date to year, month, day
func (jDN JulianDay) Date() (y, m, d int) {
	j := int(jDN)
	f := j + 1401 + (((4*j+274277)/146097)*3)/4 - 38
	e := 4*f + 3
	g := (e % 1461) / 4
	h := 5*g + 2
	d = (h%153)/5 + 1
	m = (h/153+2)%12 + 1
	y = e/1461 - 4716 + (12+2-m)/12
	return
}

// return UTC time.Time
func (jDN JulianDay) UTC() time.Time {
	y, m, d := jDN.Date()
	if y >= 0 {
		return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
	}
	return time.Unix(0, 0).UTC()
}

// return YYYYMMDD base 10 uint32
func (jDN JulianDay) Uint32() uint32 {
	y, m, d := jDN.Date()
	if y >= 0 {
		res := y*10000 + m*100 + d
		return uint32(res)
	}
	return 0
}

// convert base 10 uint32 YYYYMMDD to julianDay
func FromUint32(days uint32) JulianDay {
	year := int(days / 10000)
	md := days % 10000
	mon := int(md / 100)
	mday := int(md % 100)
	return newJDN(year, mon, mday)
}

// return "YYYYMMDD" string
func (julianDay JulianDay) String8() string {
	y, m, d := julianDay.Date()
	if y >= 0 {
		res := y*10000 + m*100 + d
		return strconv.FormatInt(int64(res), 10)
	}
	y = -y
	res := y*10000 + m*100 + d
	return strconv.FormatInt(int64(-res), 10)
	/*
		if year, month, day, err:= julianDay.GetYMD(); err == nil {
			res := year*10000 + month*100 + day
			return strconv.FormatInt(int64(res), 10)
			//return fmt.Sprintf("%04d%02d%02d", year, month, day)
		} else {
			return "Invalid JulianDay"
		}
		return ""
	*/
}

func (jDN JulianDay) String() string {
	y, m, d := jDN.Date()
	if y >= 0 {
		return fmt.Sprintf("%04d-%02d-%02d", y, m, d)
		/*
			res := y * 10000 + m * 100 + d
			ss := strconv.FormatInt(int64(res), 10)
			return ss[:4]+"-"+ss[4:6]+"-"+ss[6:]
		*/
	}
	y = -y
	return fmt.Sprintf("BC %04d-%02d-%02d", y, m, d)
	/*
		res := y * 10000 + m * 100 + d
		ss := strconv.FormatInt(int64(res), 10)
		return "BC "+ss[:4]+"-"+ss[4:6]+"-"+ss[6:]
	*/
}

// FormatFrom "YYYY-MM-DD" to julianDay
func FromString(buffer string) JulianDay {
	years, _ := strconv.Atoi(string(buffer[:4]))
	months, _ := strconv.Atoi(string(buffer[5:7]))
	day, _ := strconv.Atoi(string(buffer[8:10]))
	return NewJulianDay(years, months, day)
}

// forward nDays
func (jDN JulianDay) Add(nDays int) JulianDay {
	return JulianDay(int(jDN) + nDays)
}


func (jDN JulianDay) Sub(j JulianDay) int {
	return int(jDN) - int(j)
}

// return Weekday time.Weekday
func (jDN JulianDay) Weekday() time.Weekday {
	res := uint32(jDN) + 1
	res %= 7
	return time.Weekday(res)
}

// return weekBase sunday julianDay
func (jDN JulianDay) Weekbase() JulianDay {
	res := uint32(jDN)+1
	res %= 7
	return JulianDay(uint32(jDN)-res)
}
