package flightradar

import "reflect"

func getString(v interface{}) string {
	if reflect.TypeOf(v).String() == "string" {
		return v.(string)
	}
	return ""
}

func getInt64(v interface{}) int64 {
	if reflect.TypeOf(v).String() == "int64" {
		return v.(int64)
	}
	return 0
}

func getBoolean(v interface{}) bool {
	return v.(float64) == 1
}

func getUint(v interface{}) uint {
	if reflect.TypeOf(v).String() == "uint" {
		return v.(uint)
	}
	return 0
}
