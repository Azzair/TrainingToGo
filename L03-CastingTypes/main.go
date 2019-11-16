package main

import (
	"encoding/binary"
	"fmt"
	"math"
)

func main() {
	var (
		a int32   = -3
		b float64 = 1.01
		c string  = "hello"
	)
	as := make([]byte, 4)
	bs := make([]byte, 8)

	binary.BigEndian.PutUint32(as, uint32(a))
	binary.BigEndian.PutUint64(bs, math.Float64bits(b))

	fmt.Printf("a int: %#x\n", as)
	fmt.Printf("b int: %#x\n", bs)
	fmt.Println("c int: ", []byte(c))
}
