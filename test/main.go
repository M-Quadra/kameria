package main

import (
	"fmt"

	"github.com/m_quadra/kameria"
)

type testModel struct {
	KeyName string `dbkey:"key1"`
	KeyUsr  string `dbkey:"key2"`
}

func main() {
	fmt.Println("kameria の test")

	ipv4Str := "192.168.0.1"
	ipv4Num, err := kameria.IPv4StringToInt(ipv4Str)
	fmt.Println(ipv4Str, "->", ipv4Num, err)
	ipv4Str, err = kameria.IPv4IntToString(ipv4Num)
	fmt.Println(ipv4Num, "->", ipv4Str, err)

	var model testModel
	model.KeyName = "name1"
	model.KeyUsr = "usr"

	needMap := map[string]interface{}{
		"key1": &model.KeyName,
		"key2": &model.KeyUsr,
	}
	fmt.Println("      needMap:", needMap)

	tagMap, err := kameria.StructTagMap("dbkey", &model)
	if kameria.HasError(err) {
		fmt.Println(".StructTagMap error")
	} else {
		fmt.Println(".StructTagMap:", tagMap)
	}
}
