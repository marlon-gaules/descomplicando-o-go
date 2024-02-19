package main

import (
	"fmt"
	"math"
)

func main() {

	var a int8
	var b int16
	var c int32
	var d int64
	a = math.MaxInt8
	b = math.MaxInt16
	c = math.MaxInt32
	d = math.MaxInt64
	fmt.Println("Tamanho máximo int8 :: ", a)
	fmt.Println("Tamanho máximo int8 :: ", b)
	fmt.Println("Tamanho máximo int8 :: ", c)
	fmt.Println("Tamanho máximo int8 :: ", d)

	a = math.MinInt8
	b = math.MinInt16
	c = math.MinInt32
	d = math.MinInt64
	fmt.Println("Tamanho mínimo int8 :: ", a)
	fmt.Println("Tamanho mínimo int8 :: ", b)
	fmt.Println("Tamanho mínimo int8 :: ", c)
	fmt.Println("Tamanho mínimo int8 :: ", d)

	var i uint = 18446744073709551615
	fmt.Println(i)
	for j := 0; j < 10; j++ {
		i++
		fmt.Println(i)
	}

	fmt.Println("-----------------")

	var birds int32 = 5
	var bees int16 = 10
	fmt.Println(birds)
	fmt.Println(bees)
	animals := birds + int32(bees)
	fmt.Println(animals)

	var e float32
	e = 12.52
	fmt.Println(e)

	var f float32
	f = 125e10
	fmt.Println(f)

	var g float32
	g = 12e-2
	fmt.Println(g)

	h := complex(5, 2)
	fmt.Println(real(h))
	j := complex(5, 2)
	fmt.Println(imag(j))
	fmt.Println(j + complex(10, 5))

	k := "LinuxTips"
	fmt.Println(k)
	fmt.Println(len(k))

	var l bool
	l = true
	fmt.Println(l)
	if l == (1 == 1) {
		fmt.Println(l)
	}

	fmt.Println(1 == 2)
	fmt.Println(1 != 3)
	fmt.Println(1 < 3)
	fmt.Println(3 > 5)
	fmt.Println(3 >= 3)
	fmt.Println(3 <= 3)
}
