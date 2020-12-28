package domain

type Unit struct {
	BaseUnit string
	Prefix   Prefix
}

// The Category enum helps to identify whether certain variables are commensurable.
type Category int

const (
	Liquid = iota
	Solid
	Approx
	Piece
)

// The Prefix is used to convert between units with numeric prefixes (e.g. gram and kilogram).
type Prefix float32

const (
	milli = 1 / 1000
	kilo  = 1000
)
