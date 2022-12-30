package core

import "fmt"

type boolean struct {
	data bool
	typ  string
	name string
}

func Bool(v ...any) *boolean {
	name, val := getNameAndValue(v, "bool")
	return &boolean{
		data: val.(bool),
		name: name,
		typ:  "bool",
	}
}

//
func (b *boolean) String(query ...string) string {
	return fmt.Sprint(b.data)
}

//
func (b *boolean) Int(query ...string) int {
	if b.data {
		return 1
	}
	return 0
}

//
func (b *boolean) Float(query ...string) float64 {
	if b.data {
		return 1
	}
	return 0
}
