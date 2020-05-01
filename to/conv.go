package to

import (
	"math"
	"reflect"
	"strconv"
)

func String(v interface{}) string {
	switch v.(type) {
	case string:
		return reflect.ValueOf(v).String()
	case int8, int16, int32, int, int64:
		iv := reflect.ValueOf(v).Int()
		return strconv.FormatInt(iv, 10)
	case uint8, uint16, uint32, uint, uint64:
		iv := reflect.ValueOf(v).Uint()
		return strconv.FormatUint(iv, 10)
	case float32:
		iv := reflect.ValueOf(v).Float()
		return strconv.FormatFloat(iv, 'g', -1, 32)
	case float64:
		iv := reflect.ValueOf(v).Float()
		return strconv.FormatFloat(iv, 'g', -1, 64)

	}
	return reflect.ValueOf(v).String()
}

func Int64(v interface{}) int64 {
	switch v.(type) {
	case int8, int16, int32, int, int64:
		return reflect.ValueOf(v).Int()
	case uint8, uint16, uint32, uint, uint64:
		return int64(reflect.ValueOf(v).Uint())
	case float32, float64:
		return int64(reflect.ValueOf(v).Float())
	case string:
		ss := reflect.ValueOf(v).String()
		res, _ := strconv.ParseInt(ss, 10, 64)
		return res
	}
	return 0
}

func Int(v interface{}) int {
	return int(Int64(v))
}

func Uint64(v interface{}) uint64 {
	switch v.(type) {
	case int8, int16, int32, int, int64:
		return uint64(reflect.ValueOf(v).Int())
	case uint8, uint16, uint32, uint, uint64:
		return reflect.ValueOf(v).Uint()
	case float32, float64:
		return uint64(reflect.ValueOf(v).Float())
	case string:
		ss := reflect.ValueOf(v).String()
		res, _ := strconv.ParseUint(ss, 10, 64)
		return res
	}
	return 0
}

func Double(v interface{}) float64 {
	switch v.(type) {
	case int8, int16, int32, int, int64:
		return float64(reflect.ValueOf(v).Int())
	case uint8, uint16, uint32, uint, uint64:
		return float64(reflect.ValueOf(v).Uint())
	case float32, float64:
		return reflect.ValueOf(v).Float()
	case string:
		ss := reflect.ValueOf(v).String()
		res, _ := strconv.ParseFloat(ss, 64)
		return res
	case nil:
		return math.NaN()
	}
	return 0
}
