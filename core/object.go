package core

type object struct {
	name string
	data map[string]any
	typ  string
}

// constructor
func Struct(args ...any) *object {

	obj := &object{
		data: make(map[string]any),
		typ:  "object",
		name: extratName(args[0]),
	}
	for _, item := range args {
		obj.data[getName(item)] = item
	}
	return obj
}

//
func (o *object) String(name ...string) string {
	if len(name) > 0 {
		return o.data[name[0]].(*str).data
	}
	return ""
}

//
func (o *object) Int(name ...string) int {
	if len(name) > 0 {
		return o.data[name[0]].(*integer).data
	}
	return 0
}

//
func (o *object) Float(name ...string) float64 {
	if len(name) > 0 {
		return o.data[name[0]].(*floater).data
	}
	return 0.0
}

//
func (o *object) Bool(name ...string) bool {
	if len(name) > 0 {
		return o.data[name[0]].(*boolean).data
	}
	return false
}

//
func (o *object) Struct(name ...string) *object {
	if len(name) > 0 {
		return o.data[name[0]].(*object)
	}
	return nil
}

/*// assets //*/
// return name item
func getName(item any) (name string) {
	switch item.(type) {
	case *integer:
		return item.(*integer).name
	case *floater:
		return item.(*floater).name
	case *str:
		return item.(*str).name
	case *boolean:
		return item.(*boolean).name
	case *tupla:
		return item.(*tupla).name
	case *object:
		return item.(*object).name
	default:
		return "undefined"
	}

}
