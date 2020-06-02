/**
 * @Author: alienongwlx@gmail.com
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/5/29 14:30
 */

package misc

import (
	"errors"
	"reflect"
)

func SliceArray(size int, arr interface{}, beans interface{}) (error, interface{}) {
	slice := reflect.Indirect(reflect.ValueOf(arr))
	len := slice.Len()
	if slice.Kind() != reflect.Slice || len == 0 {
		return errors.New("need a slice"), nil
	}
	res := reflect.Indirect(reflect.ValueOf(beans))
	rType := res.Type()
	count := len/size + 1
	rSlice := reflect.MakeSlice(rType, 0, 0)
	for i := 0; i < count; i++ {
		if i*size >= len {
			continue
		}
		ty := slice.Type()
		temp := reflect.MakeSlice(ty, 0, 0)
		end := (i + 1) * size
		if end > len {
			end = len
		}
		for j := i * size; j < end; j++ {
			temp = reflect.Append(temp, slice.Index(j))
		}
		rSlice = reflect.Append(rSlice, reflect.ValueOf(temp.Interface()))
	}
	return nil, rSlice.Interface()
}
