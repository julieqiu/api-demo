package apidemo

type Field int

const (
	A Field = 1 << iota
	B
	C
)

type Geometry interface {
	Area() float64
	Perim(x, y int) float64
}
