package ddd

import "fmt"

type Domain interface {
	fmt.Stringer
	Parent() Domain
}
