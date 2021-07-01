package apidemo

import (
	"errors"
	"fmt"
)

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

type Name struct {
	Country, Organization, OrganizationalUnit []string
	Locality, Province                        []string
	StreetAddress, PostalCode                 []string
	SerialNumber, CommonName                  string
}

func foo() {
	fmt.Println("hi")
}

func Foo() {
	fmt.Println("hi2")
}

func FoO() {
	fmt.Println("hi3")
}

func FOO() {
	fmt.Println("hi4")
}
