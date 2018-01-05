package go_helpers

import (
	"github.com/kataras/iris/core/errors"
	"reflect"
	"strconv"
)

//func GetIIMapValues(m map[int]int) []int {
//	vs := make([]int, len(m))
//	for _, v := range m {
//		vs = append(vs, v)
//	}
//	return vs
//}

//Removes duplicate values from an slice
func MakeSliceUnique(s interface{}) []interface{} {
	rt := make([]interface{}, 0)
	m := map[string]bool{}
	ok := false
	k := ""

	switch  t := s.(type) {
	case []bool:
		v := reflect.ValueOf(t)
		for i := 0; i < v.Len(); i++ {
			if v.Index(i).Bool() {
				k = "0"
			} else {
				k = "1"
			}
			if _, ok = m[k]; !ok {
				rt = append(rt, v)
				m[k] = true
			}
		}
	case []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64:
		v := reflect.ValueOf(t)
		for i, tt := 0, int64(0); i < v.Len(); i++ {
			tt = v.Index(i).Int()
			k = strconv.FormatInt(tt, 10)
			if _, ok = m[k]; !ok {
				rt = append(rt, tt)
				m[k] = true
			}
		}
	case []float32, []float64:
		v := reflect.ValueOf(t)
		for i, tt := 0, 0.0; i < v.Len(); i++ {
			tt = v.Index(i).Float()
			k = strconv.FormatFloat(tt, 'f', -1, 64)
			if _, ok = m[k]; !ok {
				rt = append(rt, tt)
				m[k] = true
			}
		}
	case []string:
		v := reflect.ValueOf(t)
		for i := 0; i < v.Len(); i++ {
			k = v.Index(i).String()
			if _, ok = m[k]; !ok {
				rt = append(rt, v)
				m[k] = true
			}
		}
	default:
		panic(errors.New("cannot support type to unique: " + reflect.TypeOf(s).String()))
	}

	return rt
}
