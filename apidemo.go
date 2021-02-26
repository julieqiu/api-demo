package apidemo

import "errors"

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

var (
	ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
	ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
	ErrBufferFull        = errors.New("bufio: buffer full")
	ErrNegativeCount     = errors.New("bufio: negative count")
)
