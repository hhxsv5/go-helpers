package go_helpers

import (
	"reflect"
	"strconv"
	"errors"
)

//func GetIIMapValues(m map[int]int) []int {
//	vs := make([]int, len(m))
//	for _, v := range m {
//		vs = append(vs, v)
//	}
//	return vs
//}

//Removes duplicate values from an slice
//Return a slice of interface{}, need type assertion like use a.(int)
func MakeSliceUnique(s interface{}) []interface{} {
	rt := make([]interface{}, 0)
	m := map[string]bool{}
	ok := false
	k := ""

	switch  s.(type) {
	case []bool:
		v := reflect.ValueOf(s)
		for i, tt := 0, false; i < v.Len(); i++ {
			tt = v.Index(i).Bool()
			if tt {
				k = "0"
			} else {
				k = "1"
			}
			if _, ok = m[k]; !ok {
				rt = append(rt, tt)
				m[k] = true
			}
		}
	case []int:
		v := reflect.ValueOf(s)
		for i, tt := 0, int64(0); i < v.Len(); i++ {
			tt = v.Index(i).Int()
			k = strconv.FormatInt(tt, 10)
			if _, ok = m[k]; !ok {
				rt = append(rt, int(tt))
				m[k] = true
			}
		}
	case []int8:
		v := reflect.ValueOf(s)
		for i, tt := 0, int64(0); i < v.Len(); i++ {
			tt = v.Index(i).Int()
			k = strconv.FormatInt(tt, 10)
			if _, ok = m[k]; !ok {
				rt = append(rt, int8(tt))
				m[k] = true
			}
		}
	case []int16:
		v := reflect.ValueOf(s)
		for i, tt := 0, int64(0); i < v.Len(); i++ {
			tt = v.Index(i).Int()
			k = strconv.FormatInt(tt, 10)
			if _, ok = m[k]; !ok {
				rt = append(rt, int16(tt))
				m[k] = true
			}
		}
	case []int32:
		v := reflect.ValueOf(s)
		for i, tt := 0, int64(0); i < v.Len(); i++ {
			tt = v.Index(i).Int()
			k = strconv.FormatInt(tt, 10)
			if _, ok = m[k]; !ok {
				rt = append(rt, int32(tt))
				m[k] = true
			}
		}
	case []int64:
		v := reflect.ValueOf(s)
		for i, tt := 0, int64(0); i < v.Len(); i++ {
			tt = v.Index(i).Int()
			k = strconv.FormatInt(tt, 10)
			if _, ok = m[k]; !ok {
				rt = append(rt, int64(tt))
				m[k] = true
			}
		}
	case []uint:
		v := reflect.ValueOf(s)
		for i, tt := 0, int64(0); i < v.Len(); i++ {
			tt = v.Index(i).Int()
			k = strconv.FormatInt(tt, 10)
			if _, ok = m[k]; !ok {
				rt = append(rt, uint(tt))
				m[k] = true
			}
		}
	case []uint8:
		v := reflect.ValueOf(s)
		for i, tt := 0, int64(0); i < v.Len(); i++ {
			tt = v.Index(i).Int()
			k = strconv.FormatInt(tt, 10)
			if _, ok = m[k]; !ok {
				rt = append(rt, uint8(tt))
				m[k] = true
			}
		}
	case []uint16:
		v := reflect.ValueOf(s)
		for i, tt := 0, int64(0); i < v.Len(); i++ {
			tt = v.Index(i).Int()
			k = strconv.FormatInt(tt, 10)
			if _, ok = m[k]; !ok {
				rt = append(rt, uint16(tt))
				m[k] = true
			}
		}
	case []uint32:
		v := reflect.ValueOf(s)
		for i, tt := 0, int64(0); i < v.Len(); i++ {
			tt = v.Index(i).Int()
			k = strconv.FormatInt(tt, 10)
			if _, ok = m[k]; !ok {
				rt = append(rt, uint32(tt))
				m[k] = true
			}
		}
	case []uint64:
		v := reflect.ValueOf(s)
		for i, tt := 0, int64(0); i < v.Len(); i++ {
			tt = v.Index(i).Int()
			k = strconv.FormatInt(tt, 10)
			if _, ok = m[k]; !ok {
				rt = append(rt, uint64(tt))
				m[k] = true
			}
		}
	case []float32:
		v := reflect.ValueOf(s)
		var (
			tt interface{}
			vv float32
		)
		for i := 0; i < v.Len(); i++ {
			tt = v.Index(i).Interface()
			vv = tt.(float32)
			k = strconv.FormatFloat(float64(vv), 'f', -1, 32)
			if _, ok = m[k]; !ok {
				rt = append(rt, vv)
				m[k] = true
			}
		}
	case []float64:
		v := reflect.ValueOf(s)
		var (
			tt interface{}
			vv float64
		)
		for i := 0; i < v.Len(); i++ {
			tt = v.Index(i).Interface()
			vv = tt.(float64)
			k = strconv.FormatFloat(vv, 'f', -1, 64)
			if _, ok = m[k]; !ok {
				rt = append(rt, vv)
				m[k] = true
			}
		}
	case []string:
		v := reflect.ValueOf(s)
		for i := 0; i < v.Len(); i++ {
			k = v.Index(i).String()
			if _, ok = m[k]; !ok {
				rt = append(rt, k)
				m[k] = true
			}
		}
	default:
		panic(errors.New("cannot support type to make unique: " + reflect.TypeOf(s).String()))
	}

	return rt
}
