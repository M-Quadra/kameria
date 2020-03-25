package kameria

import (
	"fmt"
	"testing"
)

func TestIPv4conv(t *testing.T) {
	fmt.Println("[kameria] IPv4conv test...")

	ipv4Str := "192.168.0.1"
	ipv4Num, err := IPv4StringToInt(ipv4Str)
	if err != nil || ipv4Num != 3232235521 {
		t.Fail()
	}

	fmt.Println(ipv4Str, "->", ipv4Num, err)

	ipv4Str, err = IPv4IntToString(ipv4Num)
	if err != nil || ipv4Str != "192.168.0.1" {
		t.Fail()
	}

	fmt.Println(ipv4Num, "->", ipv4Str, err)
}
