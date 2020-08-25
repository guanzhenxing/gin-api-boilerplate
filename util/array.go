package util

import (
	"reflect"
)

//判断一个元素是否在数组中
func InArray(val interface{}, array interface{}) (exists bool) {
	exists = false

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				exists = true
				return
			}
		}
	}

	return
}

// 返回单元顺序相反的数组
func ArrayReverse(array interface{}) {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
			iTmp := reflect.ValueOf(s.Index(i).Interface())
			jTmp := reflect.ValueOf(s.Index(j).Interface())
			s.Index(i).Set(jTmp)
			s.Index(j).Set(iTmp)
		}
	}
}
