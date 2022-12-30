package core

import "fmt"

type floater struct {
	data float64
	name string
	typ  string
}

func Float(value ...any) *floater {
	name, val := getNameAndValue(value, "float64")
	return &floater{
		data: val.(float64),
		name: name,
		typ:  Type(val),
	}
}

//
func (i *floater) String(query ...string) string {
	return fmt.Sprint(i.data)
}

//
func (i *floater) Int(query ...string) int {
	return int(i.data)
}

//
func (i *floater) Float(query ...string) float64 {
	return i.data
}

//
func (i *floater) Bool(query ...string) bool {
	return i.data == 0
}
