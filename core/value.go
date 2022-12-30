package core

type Value interface {
	String(...string) string
	Int(...string) int
	Bool(...string) bool
	Float(...string) float64
}
