package core

import "strings"
import "reflect"
import "strconv"

// extract name of string
func extratName(v any) string {
	name, ok := v.(string)
	if ok && strings.Contains(name, "#") {
		name = name[1:]
		return name
	}
	return "undefined"
}

// return string type of variable
func Type(v any) string {
	return reflect.TypeOf(v).String()
}

// return real value if name is undefined
func IsUndefined(name string, a, b any) any {
	if name == "undefined" {
		return a
	}
	return b
}

// stop program if value no is type
func IsType(value any, types string) {
	if Type(value) != types {
		panic("value no is " + types)
	}
}

//
func getNameAndValue(values []any, types string) (name string, res any) {
	name = extratName(values[0])
	if len(values) > 1 {
		res = IsUndefined(name, values[0], values[1])
	} else {
		res = values[0]
	}
	IsType(res, types)
	return
}

//
func GetNameAndAllValues(values []any) (name string, res []any) {
	if len(values) > 0 {
		name = extratName(values[0])
		if name != "undefined" {
			res = values[1:]
		} else {
			res = values
		}
	}
	return
}

//
func toInt(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return number
}
