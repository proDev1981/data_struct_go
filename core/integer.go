package core

import "fmt"

type integer struct {
	data int
	name string
	typ  string
}

func Int(value ...any) *integer {
	name, val := getNameAndValue(value, "int")
	return &integer{
		data: val.(int),
		name: name,
		typ:  Type(val),
	}
}

//
func (i *integer) String(query ...string) string {
	return fmt.Sprint(i.data)
}

//
func (i *integer) Int(query ...string) int {
	return i.data
}

//
func (i *integer) Float(query ...string) float64 {
	return float64(i.data)
}

//
func (i *integer) Bool(quey ...string) bool {
	return i.data == 0
}
