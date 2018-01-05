package go_helpers

import (
	"strconv"
	"github.com/kataras/iris/core/errors"
	"reflect"
)

//func GetIIMapValues(m map[int]int) []int {
//	vs := make([]int, len(m))
//	for _, v := range m {
//		vs = append(vs, v)
//	}
//	return vs
//}

//Removes duplicate values from an slice
func MakeSliceUnique(s []interface{}) []interface{} {
	rt := make([]interface{}, 0)
	m := map[string]bool{}
	ok := false
	k := ""
	for _, v := range s {
		switch v.(type) {
		case int, int8, int16, int32, int64, uint8, uint16, uint32, uint64:
			k = strconv.FormatInt(int64(v.(int)), 10)
		case string:
			k = v.(string)
		default:
			panic(errors.New("cannot support type to unique: " + reflect.TypeOf(v).String()))
		}
		if _, ok = m[k]; !ok {
			rt = append(rt, v)
			m[k] = true
		}
	}
	return rt
}
