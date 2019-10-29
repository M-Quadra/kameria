package main

import (
	"fmt"

	"github.com/m_quadra/kameria"
)

func main() {
	fmt.Println("kameria の test")

	ipv4Str := "192.168.0.1"
	ipv4Num, err := kameria.IPv4StringToInt(ipv4Str)
	fmt.Println(ipv4Str, "->", ipv4Num, err)
	ipv4Str, err = kameria.IPv4IntToString(ipv4Num)
	fmt.Println(ipv4Num, "->", ipv4Str, err)

}
