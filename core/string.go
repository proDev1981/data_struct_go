package core

type str struct {
	data string
	typ  string
	name string
}

func String(v ...any) *str {
	name, val := getNameAndValue(v, "string")
	return &str{
		data: val.(string),
		name: name,
		typ:  "string",
	}
}

//
func (s *str) String(query ...string) string {
	return s.data
}

//
func (s *str) Int(query ...string) int {
	if s.Bool() {
		return 1
	}
	return 0
}

//
func (s *str) Float(query ...string) float64 {
	if s.Bool() {
		return 1
	}
	return 0
}

//
func (s *str) Bool(query ...string) bool {
	return s.data != ""
}
