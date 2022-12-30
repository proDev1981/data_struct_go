package core

import "strconv"

//import "fmt"

type tupla struct {
	name string
	data []any
	typ  string
}

func Tupla(args ...any) *tupla {
	name, values := GetNameAndAllValues(args)
	return &tupla{
		data: values,
		typ:  "tupla",
		name: name,
	}
}

// return len tupla
func (t *tupla) Len() int {
	return len(t.data)
}

//
func (t *tupla) Range() []any {
	return t.data
}

//
func (t *tupla) String(query ...string) string {
	return t.data[t.searchQuery(query)].(string)
}

//
func (t *tupla) Int(query ...string) int {
	return t.data[t.searchQuery(query)].(int)
}

//
func (t *tupla) Bool(query ...string) bool {
	return t.data[t.searchQuery(query)].(bool)
}

//
func (t *tupla) Struct(query ...string) *object {
	return t.data[t.searchQuery(query)].(*object)
}

//
func (t *tupla) Float(query ...string) float64 {
	return t.data[t.searchQuery(query)].(float64)
}

// foreach
func (t *tupla) ForEach(fn func(int, Value)) {
	for index, item := range t.data {
		fn(index, item.(Value))
	}
}

// filter
func (t *tupla) Filter(fn func(item Value) bool) *tupla {
	newtupla := &tupla{
		name: "filter_to_" + t.name,
		typ:  "tupla",
	}
	for _, item := range t.data {
		if fn(item.(Value)) {
			newtupla.data = append(newtupla.data, item)
		}
	}
	return newtupla
}

// search in item query
func (t *tupla) searchQuery(query []string) int {
	if len(query) > 0 {
		number, err := strconv.Atoi(query[0])
		if err != nil {
			for index, item := range t.data {
				switch item.(type) {
				case *integer:
					if item.(*integer).name == query[0] {
						return index
					}
				case *floater:
					if item.(*floater).name == query[0] {
						return index
					}
				case *str:
					if item.(*str).name == query[0] {
						return index
					}
				case *boolean:
					if item.(*boolean).name == query[0] {
						return index
					}
				case *tupla:
					if item.(*tupla).name == query[0] {
						return index
					}
				case *object:
					if item.(*object).name == query[0] {
						return index
					}
				default:
					return -1
				}

			}
		}
		return number
	}
	return -1
}
